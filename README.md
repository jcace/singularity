# Singularity
The new pure-go implementation of Singularity provides everything you need to onboard your, or your client's data to Filecoin network

## Current Status
This project is currently in active development. Below are the current feature list and their status.
|   |   |
|---|---|
| ![Stable](https://img.shields.io/badge/-Stable-brightgreen) | Feature is stable and ready for production use |
| ![Beta](https://img.shields.io/badge/-Beta-blue) | Feature is in beta and may still contain bugs |
| ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Feature is in alpha and should not be used in production |
| ![WIP](https://img.shields.io/badge/-WIP-yellow) | Feature is currently being worked on and is not usable |
| ![Planned](https://img.shields.io/badge/-Planned-lightgrey) | Feature is planned but not yet implemented |

| Category | Feature | Status                                                      | Description |
| --- | --- |-------------------------------------------------------------| --- |
| Data Source | File System | ![Alpha](https://img.shields.io/badge/-Alpha-orange)        | Support for preparing data on local file system |
| Data Source | Nginx File Browser | ![Alpha](https://img.shields.io/badge/-Alpha-orange)                                                            | Support for preparing data from Nginx directory listing service with autoindex turned on |
| Data Source | S3 Compatible | ![Alpha](https://img.shields.io/badge/-Alpha-orange)            | Support for preparing data from S3 compatible storage service |
| Data Prep | Create Dataset | ![Beta](https://img.shields.io/badge/-Beta-blue)            | CLI tool for creating dataset |
| Data Prep | Add Data Source | ![Beta](https://img.shields.io/badge/-Beta-blue)            | CLI tool for adding data sources to existing dataset |
| Data Prep | Inline Preparation | ![Alpha](https://img.shields.io/badge/-Alpha-orange)        | Support for inline preparation. No need to export CAR files |
| Data Prep | Upload API | ![Alpha](https://img.shields.io/badge/-Alpha-orange)            | Support for manually upload files via API |
| Data Prep | Push API | ![Alpha](https://img.shields.io/badge/-Alpha-orange)            | Support for manually queue a new item with an item path via API |
| Data Prep | Dag regen | ![Planned](https://img.shields.io/badge/-Planned-lightgrey) | Regenerate the unixfs dag to update the root CID of the whole dataset |
| Data Prep | Basic Encryption | ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Support for basic encryption with asynmmtric keys |
| Data Prep | Custom Encryption | ![Planned](https://img.shields.io/badge/-Planned-lightgrey) | Support for custom encryption with user providing encryption tools |
| Content Provider | HTTP piece | ![Alpha](https://img.shields.io/badge/-Alpha-orange)            | Support for HTTP piece download (CAR file download) |
| Content Provider | IPFS gateway | ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Support for IPFS Gatway compliant retrieval |
| Content Provider | Bitswap | ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Support for Bitswap retrieval (IPFS interop) |
| Content Provider | Graphsync | ![Planned](https://img.shields.io/badge/-Planned-lightgrey) | Support for Graphsync retrieval |
| Content Provider | Metadata API | ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Allow CAR file distribution from the original data owner |
| Content Provider | Donwload Client | ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Support downloading CAR file from the original data owner with the help of Metadata API |
| Deal Making | Deal Tracking | ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Track deal status |
| Deal Making | Deal Scheduler | ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Feature parity with [js-singularity](https://github.com/tech-greedy/singularity/tree/main#deal-replication) deal replication |
| Deal Making | Spade API | ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Support for compatible Spade API for storage provider to self proposal deals |
| Deal Making | Wallet Management | ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Support for wallet management |
| Deal Making | Remote Signer | ![Planned](https://img.shields.io/badge/-Planned-lightgrey) | Support for remote signer |
| Utilities | Benchmark | ![Planned](https://img.shields.io/badge/-Planned-lightgrey) | Support for benchmarking data preparation |
| Utilities | Monitoring | ![Planned](https://img.shields.io/badge/-Planned-lightgrey) | Support for monitoring data preparation and deal making |
| Dashboard | Dataset Explorer | ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Support for exploring dataset, folder by folder |
| Dashboard | Dataset Download | ![Planned](https://img.shields.io/badge/-Planned-lightgrey) | Support for downloading dataset directly on the browser |
| Dashboard | Deal Explorer | ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Support for exploring deal proposals |
| Dashboard | Piece Explorer | ![Alpha](https://img.shields.io/badge/-Alpha-orange) | Support for exploring by piece CIDs and check distribution |
| Dashboard | Provider View | ![Planned](https://img.shields.io/badge/-Planned-lightgrey) | Support for checking how fast providers are consuming the deals |


## Related projects
- [js-singularity](https://github.com/tech-greedy/singularity) -
The predecessor that was implemented in NodeJS
- [js-singularity-import-boost](https://github.com/tech-greedy/singularity-import) -
Automatically import deals to boost for Filecoin storage providers
- [js-singularity-browser](https://github.com/tech-greedy/singularity-browser) -
A next.js app for browsing singularity made deals
- [go-generate-car](https://github.com/tech-greedy/generate-car) -
The internal tool used by `js-singularity` to generate car files as well as commp
- [go-generate-ipld-car](https://github.com/tech-greedy/generate-car#generate-ipld-car) -
The internal tool used by `js-singularity` to regenerate the CAR that captures the unixfs dag of the dataset.

