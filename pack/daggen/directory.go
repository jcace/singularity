package daggen

import (
	"bufio"
	"bytes"
	"context"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-ipfs-blockstore"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	uio "github.com/ipfs/go-unixfs/io"
	"github.com/ipld/go-car"
	"github.com/ipld/go-car/util"
	"github.com/klauspost/compress/zstd"
	"github.com/pkg/errors"
	"io"
)

var encoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
var decoder, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(0))

func ResolveDirectoryTree(currentID uint64,
	dirCache map[uint64]*DirectoryData,
	childrenCache map[uint64][]uint64,
) (*format.Link, error) {
	current, ok := dirCache[currentID]
	if !ok {
		return nil, errors.Errorf("no directory data for current %d", currentID)
	}
	children, _ := childrenCache[currentID]

	for _, child := range children {
		link, err := ResolveDirectoryTree(child, dirCache, childrenCache)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve child %d", child)
		}
		err = current.AddItem(link.Name, link.Cid, link.Size)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to add child %d to directory", child)
		}
	}

	node, err := current.dir.GetNode()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Node from directory")
	}
	size, err := node.Size()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get size from Node")
	}
	current.Node = node
	return &format.Link{
		Name: current.Directory.Name,
		Size: size,
		Cid:  node.Cid(),
	}, nil
}

type DirectoryData struct {
	Directory model.Directory
	Node      format.Node
	dir       uio.Directory
	bstore    blockstore.Blockstore
}

func NewDirectoryData() DirectoryData {
	ds := datastore.NewMapDatastore()
	bs := blockstore.NewBlockstore(ds)
	bs.HashOnRead(false)
	dagServ := merkledag.NewDAGService(blockservice.New(bs, nil))
	dir := uio.NewDirectory(dagServ)
	dir.SetCidBuilder(merkledag.V1CidPrefix())
	node, _ := dir.GetNode()
	return DirectoryData{
		dir:    dir,
		bstore: bs,
		Node:   node,
	}
}

func (d *DirectoryData) AddItem(name string, c cid.Cid, length uint64) error {
	return d.dir.AddChild(context.Background(), name, NewDummyNode(length, c))
}

func (d *DirectoryData) AddItemFromLinks(name string, links []format.Link) (cid.Cid, error) {
	ctx := context.Background()
	blks, node, err := pack.AssembleItemFromLinks(links)
	if err != nil {
		return cid.Undef, errors.Wrap(err, "failed to assemble item from links")
	}
	err = d.dir.AddChild(ctx, name, node)
	if err != nil {
		return cid.Undef, errors.Wrap(err, "failed to add child to directory")
	}
	err = d.bstore.PutMany(ctx, blks)
	if err != nil {
		return cid.Undef, errors.Wrap(err, "failed to put blocks into blockstore")
	}
	return node.Cid(), nil
}

func (d *DirectoryData) MarshalBinary() ([]byte, error) {
	d.bstore.HashOnRead(false)
	buf := &bytes.Buffer{}
	root, err := d.dir.GetNode()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get root Node")
	}
	_, err = pack.WriteCarHeader(buf, root.Cid())
	if err != nil {
		return nil, errors.Wrap(err, "failed to write CAR header")
	}
	ctx := context.Background()
	err = d.bstore.DeleteBlock(ctx, d.Node.Cid())
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete old Node from blockstore")
	}
	d.Node = root
	err = d.bstore.Put(ctx, root)
	if err != nil {
		return nil, errors.Wrap(err, "failed to put root Node into blockstore")
	}
	ch, err := d.bstore.AllKeysChan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all keys from blockstore")
	}
	for k := range ch {
		data, err := d.bstore.Get(ctx, k)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get data from blockstore")
		}
		_, err = pack.WriteCarBlock(buf, data)
		if err != nil {
			return nil, errors.Wrap(err, "failed to write CAR block")
		}
	}
	return encoder.EncodeAll(buf.Bytes(), make([]byte, 0, len(buf.Bytes()))), nil
}

func (d *DirectoryData) UnmarshallBinary(data []byte) error {
	ds := datastore.NewMapDatastore()
	bs := blockstore.NewBlockstore(ds)
	bs.HashOnRead(false)
	dagServ := merkledag.NewDAGService(blockservice.New(bs, nil))
	if len(data) == 0 {
		dir := uio.NewDirectory(dagServ)
		dir.SetCidBuilder(merkledag.V1CidPrefix())
		node, err := dir.GetNode()
		if err != nil {
			return errors.Wrap(err, "failed to get Node from directory")
		}
		*d = DirectoryData{
			dir:    dir,
			bstore: bs,
			Node:   node,
		}
		return nil
	}

	ctx := context.Background()
	decoded, err := decoder.DecodeAll(data, nil)
	if err != nil {
		return errors.Wrap(err, "failed to decode data")
	}
	reader := bufio.NewReader(bytes.NewReader(decoded))
	ch, err := car.ReadHeader(reader)
	if err != nil {
		return errors.Wrap(err, "failed to read CAR header")
	}
	dirCID := ch.Roots[0]
	for {
		c, data, err := util.ReadNode(reader)
		if err != nil && err != io.EOF {
			return errors.Wrap(err, "failed to read CAR block")
		}
		if err == io.EOF {
			break
		}
		blk, _ := blocks.NewBlockWithCid(data, c)
		err = bs.Put(ctx, blk)
		if err != nil {
			return errors.Wrap(err, "failed to put data into blockstore")
		}
	}
	dirNode, err := dagServ.Get(ctx, dirCID)
	if err != nil {
		return errors.Wrap(err, "failed to get root Node")
	}
	dir, err := uio.NewDirectoryFromNode(dagServ, dirNode)
	if err != nil {
		return errors.Wrap(err, "failed to create directory from Node")
	}
	dir.SetCidBuilder(merkledag.V1CidPrefix())
	*d = DirectoryData{
		dir:    dir,
		bstore: bs,
		Node:   dirNode,
	}
	return nil
}
