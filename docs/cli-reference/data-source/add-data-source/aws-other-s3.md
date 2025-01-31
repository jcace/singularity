# AWS / other S3

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add s3 - Amazon S3 Compliant Storage Providers including AWS, Alibaba, Ceph, China Mobile, Cloudflare, ArvanCloud, DigitalOcean, Dreamhost, Huawei OBS, IBM COS, IDrive e2, IONOS Cloud, Liara, Lyve Cloud, Minio, Netease, RackCorp, Scaleway, SeaweedFS, StackPath, Storj, Tencent COS, Qiniu and Wasabi

USAGE:
   singularity datasource add s3 [command options] <dataset_name> <source_path>

DESCRIPTION:
   --s3-provider
      Choose your S3 provider.

      Examples:
         | AWS          | Amazon Web Services (AWS) S3
         | Alibaba      | Alibaba Cloud Object Storage System (OSS) formerly Aliyun
         | Ceph         | Ceph Object Storage
         | ChinaMobile  | China Mobile Ecloud Elastic Object Storage (EOS)
         | Cloudflare   | Cloudflare R2 Storage
         | ArvanCloud   | Arvan Cloud Object Storage (AOS)
         | DigitalOcean | DigitalOcean Spaces
         | Dreamhost    | Dreamhost DreamObjects
         | HuaweiOBS    | Huawei Object Storage Service
         | IBMCOS       | IBM COS S3
         | IDrive       | IDrive e2
         | IONOS        | IONOS Cloud
         | LyveCloud    | Seagate Lyve Cloud
         | Liara        | Liara Object Storage
         | Minio        | Minio Object Storage
         | Netease      | Netease Object Storage (NOS)
         | RackCorp     | RackCorp Object Storage
         | Scaleway     | Scaleway Object Storage
         | SeaweedFS    | SeaweedFS S3
         | StackPath    | StackPath Object Storage
         | Storj        | Storj (S3 Compatible Gateway)
         | TencentCOS   | Tencent Cloud Object Storage (COS)
         | Wasabi       | Wasabi Object Storage
         | Qiniu        | Qiniu Object Storage (Kodo)
         | Other        | Any other S3 compatible provider

   --s3-leave-parts-on-error
      [Provider] - AWS
         If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.

         It should be set to true for resuming uploads across different sessions.

         WARNING: Storing parts of an incomplete multipart upload counts towards space usage on S3 and will add additional costs if not cleaned up.


   --s3-list-version
      Version of ListObjects to use: 1,2 or 0 for auto.

      When S3 originally launched it only provided the ListObjects call to
      enumerate objects in a bucket.

      However in May 2016 the ListObjectsV2 call was introduced. This is
      much higher performance and should be used if at all possible.

      If set to the default, 0, rclone will guess according to the provider
      set which list objects method to call. If it guesses wrong, then it
      may be set manually here.


   --s3-memory-pool-flush-time
      How often internal memory buffer pools will be flushed.

      Uploads which requires additional buffers (f.e multipart) will use memory pool for allocations.
      This option controls how often unused buffers will be removed from the pool.

   --s3-download-url
      Custom endpoint for downloads.
      This is usually set to a CloudFront CDN URL as AWS S3 offers
      cheaper egress for data downloaded through the CloudFront network.

   --s3-sts-endpoint
      [Provider] - AWS
         Endpoint for STS.

         Leave blank if using AWS to use the default endpoint for the region.

   --s3-secret-access-key
      AWS Secret Access Key (password).

      Leave blank for anonymous access or runtime credentials.

   --s3-sse-customer-key
      [Provider] - AWS, Ceph, ChinaMobile, Minio
         To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.

         Alternatively you can provide --sse-customer-key-base64.

         Examples:
            | <unset> | None

   --s3-storage-class
      [Provider] - AWS
         The storage class to use when storing new objects in S3.

         Examples:
            | <unset>             | Default
            | STANDARD            | Standard storage class
            | REDUCED_REDUNDANCY  | Reduced redundancy storage class
            | STANDARD_IA         | Standard Infrequent Access storage class
            | ONEZONE_IA          | One Zone Infrequent Access storage class
            | GLACIER             | Glacier storage class
            | DEEP_ARCHIVE        | Glacier Deep Archive storage class
            | INTELLIGENT_TIERING | Intelligent-Tiering storage class
            | GLACIER_IR          | Glacier Instant Retrieval storage class

      [Provider] - Alibaba
         The storage class to use when storing new objects in OSS.

         Examples:
            | <unset>     | Default
            | STANDARD    | Standard storage class
            | GLACIER     | Archive storage mode
            | STANDARD_IA | Infrequent access storage mode

      [Provider] - ChinaMobile
         The storage class to use when storing new objects in ChinaMobile.

         Examples:
            | <unset>     | Default
            | STANDARD    | Standard storage class
            | GLACIER     | Archive storage mode
            | STANDARD_IA | Infrequent access storage mode

      [Provider] - Liara
         The storage class to use when storing new objects in Liara

         Examples:
            | STANDARD | Standard storage class

      [Provider] - ArvanCloud
         The storage class to use when storing new objects in ArvanCloud.

         Examples:
            | STANDARD | Standard storage class

      [Provider] - TencentCOS
         The storage class to use when storing new objects in Tencent COS.

         Examples:
            | <unset>     | Default
            | STANDARD    | Standard storage class
            | ARCHIVE     | Archive storage mode
            | STANDARD_IA | Infrequent access storage mode

      [Provider] - Scaleway
         The storage class to use when storing new objects in S3.

         Examples:
            | <unset>  | Default.
            | STANDARD | The Standard class for any upload.
                       | Suitable for on-demand content like streaming or CDN.
            | GLACIER  | Archived storage.
                       | Prices are lower, but it needs to be restored first to be accessed.

      [Provider] - Qiniu
         The storage class to use when storing new objects in Qiniu.

         Examples:
            | STANDARD     | Standard storage class
            | LINE         | Infrequent access storage mode
            | GLACIER      | Archive storage mode
            | DEEP_ARCHIVE | Deep archive storage mode

   --s3-upload-concurrency
      Concurrency for multipart uploads.

      This is the number of chunks of the same file that are uploaded
      concurrently.

      If you are uploading small numbers of large files over high-speed links
      and these uploads do not fully utilize your bandwidth, then increasing
      this may help to speed up the transfers.

   --s3-v2-auth
      If true use v2 authentication.

      If this is false (the default) then rclone will use v4 authentication.
      If it is set then rclone will use v2 authentication.

      Use this only if v4 signatures don't work, e.g. pre Jewel/v10 CEPH.

   --s3-no-head-object
      If set, do not do HEAD before GET when getting objects.

   --s3-memory-pool-use-mmap
      Whether to use mmap buffers in internal memory pool.

   --s3-disable-http2
      Disable usage of http2 for S3 backends.

      There is currently an unsolved issue with the s3 (specifically minio) backend
      and HTTP/2.  HTTP/2 is enabled by default for the s3 backend but can be
      disabled here.  When the issue is solved this flag will be removed.

      See: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631



   --s3-location-constraint
      [Provider] - AWS
         Location constraint - must be set to match the Region.

         Used when creating buckets only.

         Examples:
            | <unset>        | Empty for US Region, Northern Virginia, or Pacific Northwest
            | us-east-2      | US East (Ohio) Region
            | us-west-1      | US West (Northern California) Region
            | us-west-2      | US West (Oregon) Region
            | ca-central-1   | Canada (Central) Region
            | eu-west-1      | EU (Ireland) Region
            | eu-west-2      | EU (London) Region
            | eu-west-3      | EU (Paris) Region
            | eu-north-1     | EU (Stockholm) Region
            | eu-south-1     | EU (Milan) Region
            | EU             | EU Region
            | ap-southeast-1 | Asia Pacific (Singapore) Region
            | ap-southeast-2 | Asia Pacific (Sydney) Region
            | ap-northeast-1 | Asia Pacific (Tokyo) Region
            | ap-northeast-2 | Asia Pacific (Seoul) Region
            | ap-northeast-3 | Asia Pacific (Osaka-Local) Region
            | ap-south-1     | Asia Pacific (Mumbai) Region
            | ap-east-1      | Asia Pacific (Hong Kong) Region
            | sa-east-1      | South America (Sao Paulo) Region
            | me-south-1     | Middle East (Bahrain) Region
            | af-south-1     | Africa (Cape Town) Region
            | cn-north-1     | China (Beijing) Region
            | cn-northwest-1 | China (Ningxia) Region
            | us-gov-east-1  | AWS GovCloud (US-East) Region
            | us-gov-west-1  | AWS GovCloud (US) Region

      [Provider] - ChinaMobile
         Location constraint - must match endpoint.

         Used when creating buckets only.

         Examples:
            | wuxi1      | East China (Suzhou)
            | jinan1     | East China (Jinan)
            | ningbo1    | East China (Hangzhou)
            | shanghai1  | East China (Shanghai-1)
            | zhengzhou1 | Central China (Zhengzhou)
            | hunan1     | Central China (Changsha-1)
            | zhuzhou1   | Central China (Changsha-2)
            | guangzhou1 | South China (Guangzhou-2)
            | dongguan1  | South China (Guangzhou-3)
            | beijing1   | North China (Beijing-1)
            | beijing2   | North China (Beijing-2)
            | beijing4   | North China (Beijing-3)
            | huhehaote1 | North China (Huhehaote)
            | chengdu1   | Southwest China (Chengdu)
            | chongqing1 | Southwest China (Chongqing)
            | guiyang1   | Southwest China (Guiyang)
            | xian1      | Nouthwest China (Xian)
            | yunnan     | Yunnan China (Kunming)
            | yunnan2    | Yunnan China (Kunming-2)
            | tianjin1   | Tianjin China (Tianjin)
            | jilin1     | Jilin China (Changchun)
            | hubei1     | Hubei China (Xiangyan)
            | jiangxi1   | Jiangxi China (Nanchang)
            | gansu1     | Gansu China (Lanzhou)
            | shanxi1    | Shanxi China (Taiyuan)
            | liaoning1  | Liaoning China (Shenyang)
            | hebei1     | Hebei China (Shijiazhuang)
            | fujian1    | Fujian China (Xiamen)
            | guangxi1   | Guangxi China (Nanning)
            | anhui1     | Anhui China (Huainan)

      [Provider] - ArvanCloud
         Location constraint - must match endpoint.

         Used when creating buckets only.

         Examples:
            | ir-thr-at1 | Tehran Iran (Asiatech)
            | ir-tbz-sh1 | Tabriz Iran (Shahriar)

      [Provider] - IBMCOS
         Location constraint - must match endpoint when using IBM Cloud Public.

         For on-prem COS, do not make a selection from this list, hit enter.

         Examples:
            | us-standard       | US Cross Region Standard
            | us-vault          | US Cross Region Vault
            | us-cold           | US Cross Region Cold
            | us-flex           | US Cross Region Flex
            | us-east-standard  | US East Region Standard
            | us-east-vault     | US East Region Vault
            | us-east-cold      | US East Region Cold
            | us-east-flex      | US East Region Flex
            | us-south-standard | US South Region Standard
            | us-south-vault    | US South Region Vault
            | us-south-cold     | US South Region Cold
            | us-south-flex     | US South Region Flex
            | eu-standard       | EU Cross Region Standard
            | eu-vault          | EU Cross Region Vault
            | eu-cold           | EU Cross Region Cold
            | eu-flex           | EU Cross Region Flex
            | eu-gb-standard    | Great Britain Standard
            | eu-gb-vault       | Great Britain Vault
            | eu-gb-cold        | Great Britain Cold
            | eu-gb-flex        | Great Britain Flex
            | ap-standard       | APAC Standard
            | ap-vault          | APAC Vault
            | ap-cold           | APAC Cold
            | ap-flex           | APAC Flex
            | mel01-standard    | Melbourne Standard
            | mel01-vault       | Melbourne Vault
            | mel01-cold        | Melbourne Cold
            | mel01-flex        | Melbourne Flex
            | tor01-standard    | Toronto Standard
            | tor01-vault       | Toronto Vault
            | tor01-cold        | Toronto Cold
            | tor01-flex        | Toronto Flex

      [Provider] - RackCorp
         Location constraint - the location where your bucket will be located and your data stored.


         Examples:
            | global    | Global CDN Region
            | au        | Australia (All locations)
            | au-nsw    | NSW (Australia) Region
            | au-qld    | QLD (Australia) Region
            | au-vic    | VIC (Australia) Region
            | au-wa     | Perth (Australia) Region
            | ph        | Manila (Philippines) Region
            | th        | Bangkok (Thailand) Region
            | hk        | HK (Hong Kong) Region
            | mn        | Ulaanbaatar (Mongolia) Region
            | kg        | Bishkek (Kyrgyzstan) Region
            | id        | Jakarta (Indonesia) Region
            | jp        | Tokyo (Japan) Region
            | sg        | SG (Singapore) Region
            | de        | Frankfurt (Germany) Region
            | us        | USA (AnyCast) Region
            | us-east-1 | New York (USA) Region
            | us-west-1 | Freemont (USA) Region
            | nz        | Auckland (New Zealand) Region

      [Provider] - Qiniu
         Location constraint - must be set to match the Region.

         Used when creating buckets only.

         Examples:
            | cn-east-1      | East China Region 1
            | cn-east-2      | East China Region 2
            | cn-north-1     | North China Region 1
            | cn-south-1     | South China Region 1
            | us-north-1     | North America Region 1
            | ap-southeast-1 | Southeast Asia Region 1
            | ap-northeast-1 | Northeast Asia Region 1

      [Provider] - Ceph, Minio
         Location constraint - must be set to match the Region.

         Leave blank if not sure. Used when creating buckets only.

   --s3-sse-customer-algorithm
      [Provider] - AWS, Ceph, ChinaMobile, Minio
         If using SSE-C, the server-side encryption algorithm used when storing this object in S3.

         Examples:
            | <unset> | None
            | AES256  | AES256

   --s3-sse-customer-key-md5
      [Provider] - AWS, Ceph, ChinaMobile, Minio
         If using SSE-C you may provide the secret encryption key MD5 checksum (optional).

         If you leave it blank, this is calculated automatically from the sse_customer_key provided.


         Examples:
            | <unset> | None

   --s3-max-upload-parts
      Maximum number of parts in a multipart upload.

      This option defines the maximum number of multipart chunks to use
      when doing a multipart upload.

      This can be useful if a service does not support the AWS S3
      specification of 10,000 chunks.

      Rclone will automatically increase the chunk size when uploading a
      large file of a known size to stay below this number of chunks limit.


   --s3-copy-cutoff
      Cutoff for switching to multipart copy.

      Any files larger than this that need to be server-side copied will be
      copied in chunks of this size.

      The minimum is 0 and the maximum is 5 GiB.

   --s3-profile
      Profile to use in the shared credentials file.

      If env_auth = true then rclone can use a shared credentials file. This
      variable controls which profile is used in that file.

      If empty it will default to the environment variable "AWS_PROFILE" or
      "default" if that environment variable is also not set.


   --s3-use-accelerate-endpoint
      [Provider] - AWS
         If true use the AWS S3 accelerated endpoint.

         See: [AWS S3 Transfer acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration-examples.html)

   --s3-decompress
      If set this will decompress gzip encoded objects.

      It is possible to upload objects to S3 with "Content-Encoding: gzip"
      set. Normally rclone will download these files as compressed objects.

      If this flag is set then rclone will decompress these files with
      "Content-Encoding: gzip" as they are received. This means that rclone
      can't check the size and hash but the file contents will be decompressed.


   --s3-region
      [Provider] - AWS
         Region to connect to.

         Examples:
            | us-east-1      | The default endpoint - a good choice if you are unsure.
                             | US Region, Northern Virginia, or Pacific Northwest.
                             | Leave location constraint empty.
            | us-east-2      | US East (Ohio) Region.
                             | Needs location constraint us-east-2.
            | us-west-1      | US West (Northern California) Region.
                             | Needs location constraint us-west-1.
            | us-west-2      | US West (Oregon) Region.
                             | Needs location constraint us-west-2.
            | ca-central-1   | Canada (Central) Region.
                             | Needs location constraint ca-central-1.
            | eu-west-1      | EU (Ireland) Region.
                             | Needs location constraint EU or eu-west-1.
            | eu-west-2      | EU (London) Region.
                             | Needs location constraint eu-west-2.
            | eu-west-3      | EU (Paris) Region.
                             | Needs location constraint eu-west-3.
            | eu-north-1     | EU (Stockholm) Region.
                             | Needs location constraint eu-north-1.
            | eu-south-1     | EU (Milan) Region.
                             | Needs location constraint eu-south-1.
            | eu-central-1   | EU (Frankfurt) Region.
                             | Needs location constraint eu-central-1.
            | ap-southeast-1 | Asia Pacific (Singapore) Region.
                             | Needs location constraint ap-southeast-1.
            | ap-southeast-2 | Asia Pacific (Sydney) Region.
                             | Needs location constraint ap-southeast-2.
            | ap-northeast-1 | Asia Pacific (Tokyo) Region.
                             | Needs location constraint ap-northeast-1.
            | ap-northeast-2 | Asia Pacific (Seoul).
                             | Needs location constraint ap-northeast-2.
            | ap-northeast-3 | Asia Pacific (Osaka-Local).
                             | Needs location constraint ap-northeast-3.
            | ap-south-1     | Asia Pacific (Mumbai).
                             | Needs location constraint ap-south-1.
            | ap-east-1      | Asia Pacific (Hong Kong) Region.
                             | Needs location constraint ap-east-1.
            | sa-east-1      | South America (Sao Paulo) Region.
                             | Needs location constraint sa-east-1.
            | me-south-1     | Middle East (Bahrain) Region.
                             | Needs location constraint me-south-1.
            | af-south-1     | Africa (Cape Town) Region.
                             | Needs location constraint af-south-1.
            | cn-north-1     | China (Beijing) Region.
                             | Needs location constraint cn-north-1.
            | cn-northwest-1 | China (Ningxia) Region.
                             | Needs location constraint cn-northwest-1.
            | us-gov-east-1  | AWS GovCloud (US-East) Region.
                             | Needs location constraint us-gov-east-1.
            | us-gov-west-1  | AWS GovCloud (US) Region.
                             | Needs location constraint us-gov-west-1.

      [Provider] - RackCorp
         region - the location where your bucket will be created and your data stored.


         Examples:
            | global    | Global CDN (All locations) Region
            | au        | Australia (All states)
            | au-nsw    | NSW (Australia) Region
            | au-qld    | QLD (Australia) Region
            | au-vic    | VIC (Australia) Region
            | au-wa     | Perth (Australia) Region
            | ph        | Manila (Philippines) Region
            | th        | Bangkok (Thailand) Region
            | hk        | HK (Hong Kong) Region
            | mn        | Ulaanbaatar (Mongolia) Region
            | kg        | Bishkek (Kyrgyzstan) Region
            | id        | Jakarta (Indonesia) Region
            | jp        | Tokyo (Japan) Region
            | sg        | SG (Singapore) Region
            | de        | Frankfurt (Germany) Region
            | us        | USA (AnyCast) Region
            | us-east-1 | New York (USA) Region
            | us-west-1 | Freemont (USA) Region
            | nz        | Auckland (New Zealand) Region

      [Provider] - Scaleway
         Region to connect to.

         Examples:
            | nl-ams | Amsterdam, The Netherlands
            | fr-par | Paris, France
            | pl-waw | Warsaw, Poland

      [Provider] - HuaweiOBS
         Region to connect to. - the location where your bucket will be created and your data stored. Need bo be same with your endpoint.


         Examples:
            | af-south-1     | AF-Johannesburg
            | ap-southeast-2 | AP-Bangkok
            | ap-southeast-3 | AP-Singapore
            | cn-east-3      | CN East-Shanghai1
            | cn-east-2      | CN East-Shanghai2
            | cn-north-1     | CN North-Beijing1
            | cn-north-4     | CN North-Beijing4
            | cn-south-1     | CN South-Guangzhou
            | ap-southeast-1 | CN-Hong Kong
            | sa-argentina-1 | LA-Buenos Aires1
            | sa-peru-1      | LA-Lima1
            | na-mexico-1    | LA-Mexico City1
            | sa-chile-1     | LA-Santiago2
            | sa-brazil-1    | LA-Sao Paulo1
            | ru-northwest-2 | RU-Moscow2

      [Provider] - Cloudflare
         Region to connect to.

         Examples:
            | auto | R2 buckets are automatically distributed across Cloudflare's data centers for low latency.

      [Provider] - Qiniu
         Region to connect to.

         Examples:
            | cn-east-1      | The default endpoint - a good choice if you are unsure.
                             | East China Region 1.
                             | Needs location constraint cn-east-1.
            | cn-east-2      | East China Region 2.
                             | Needs location constraint cn-east-2.
            | cn-north-1     | North China Region 1.
                             | Needs location constraint cn-north-1.
            | cn-south-1     | South China Region 1.
                             | Needs location constraint cn-south-1.
            | us-north-1     | North America Region.
                             | Needs location constraint us-north-1.
            | ap-southeast-1 | Southeast Asia Region 1.
                             | Needs location constraint ap-southeast-1.
            | ap-northeast-1 | Northeast Asia Region 1.
                             | Needs location constraint ap-northeast-1.

      [Provider] - IONOS
         Region where your bucket will be created and your data stored.


         Examples:
            | de           | Frankfurt, Germany
            | eu-central-2 | Berlin, Germany
            | eu-south-2   | Logrono, Spain

      [Provider] - Ceph, StackPath, Minio, IBMCOS
         Region to connect to.

         Leave blank if you are using an S3 clone and you don't have a region.

         Examples:
            | <unset>            | Use this if unsure.
                                 | Will use v4 signatures and an empty region.
            | other-v2-signature | Use this only if v4 signatures don't work.
                                 | E.g. pre Jewel/v10 CEPH.

   --s3-bucket-acl
      Canned ACL used when creating buckets.

      For more info visit https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      Note that this ACL is applied when only when creating buckets.  If it
      isn't set then "acl" is used instead.

      If the "acl" and "bucket_acl" are empty strings then no X-Amz-Acl:
      header is added and the default (private) will be used.


      Examples:
         | private            | Owner gets FULL_CONTROL.
                              | No one else has access rights (default).
         | public-read        | Owner gets FULL_CONTROL.
                              | The AllUsers group gets READ access.
         | public-read-write  | Owner gets FULL_CONTROL.
                              | The AllUsers group gets READ and WRITE access.
                              | Granting this on a bucket is generally not recommended.
         | authenticated-read | Owner gets FULL_CONTROL.
                              | The AuthenticatedUsers group gets READ access.

   --s3-session-token
      An AWS session token.

   --s3-list-chunk
      Size of listing chunk (response list for each ListObject S3 request).

      This option is also known as "MaxKeys", "max-items", or "page-size" from the AWS S3 specification.
      Most services truncate the response list to 1000 objects even if requested more than that.
      In AWS S3 this is a global maximum and cannot be changed, see [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html).
      In Ceph, this can be increased with the "rgw list buckets max chunk" option.


   --s3-no-head
      If set, don't HEAD uploaded objects to check integrity.

      This can be useful when trying to minimise the number of transactions
      rclone does.

      Setting it means that if rclone receives a 200 OK message after
      uploading an object with PUT then it will assume that it got uploaded
      properly.

      In particular it will assume:

      - the metadata, including modtime, storage class and content type was as uploaded
      - the size was as uploaded

      It reads the following items from the response for a single part PUT:

      - the MD5SUM
      - The uploaded date

      For multipart uploads these items aren't read.

      If an source object of unknown length is uploaded then rclone **will** do a
      HEAD request.

      Setting this flag increases the chance for undetected upload failures,
      in particular an incorrect size, so it isn't recommended for normal
      operation. In practice the chance of an undetected upload failure is
      very small even with this flag.


   --s3-use-presigned-request
      Whether to use a presigned request or PutObject for single part uploads

      If this is false rclone will use PutObject from the AWS SDK to upload
      an object.

      Versions of rclone < 1.59 use presigned requests to upload a single
      part object and setting this flag to true will re-enable that
      functionality. This shouldn't be necessary except in exceptional
      circumstances or for testing.


   --s3-versions
      Include old versions in directory listings.

   --s3-env-auth
      Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).

      Only applies if access_key_id and secret_access_key is blank.

      Examples:
         | false | Enter AWS credentials in the next step.
         | true  | Get AWS credentials from the environment (env vars or IAM).

   --s3-access-key-id
      AWS Access Key ID.

      Leave blank for anonymous access or runtime credentials.

   --s3-server-side-encryption
      [Provider] - AWS, Ceph, ChinaMobile, Minio
         The server-side encryption algorithm used when storing this object in S3.

         Examples:
            | <unset> | None
            | AES256  | AES256
            | aws:kms | aws:kms

   --s3-upload-cutoff
      Cutoff for switching to chunked upload.

      Any files larger than this will be uploaded in chunks of chunk_size.
      The minimum is 0 and the maximum is 5 GiB.

   --s3-shared-credentials-file
      Path to the shared credentials file.

      If env_auth = true then rclone can use a shared credentials file.

      If this variable is empty rclone will look for the
      "AWS_SHARED_CREDENTIALS_FILE" env variable. If the env value is empty
      it will default to the current user's home directory.

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"


   --s3-no-check-bucket
      If set, don't attempt to check the bucket exists or create it.

      This can be useful when trying to minimise the number of transactions
      rclone does if you know the bucket exists already.

      It can also be needed if the user you are using does not have bucket
      creation permissions. Before v1.52.0 this would have passed silently
      due to a bug.


   --s3-use-multipart-etag
      Whether to use ETag in multipart uploads for verification

      This should be true, false or left unset to use the default for the provider.


   --s3-might-gzip
      Set this if the backend might gzip objects.

      Normally providers will not alter objects when they are downloaded. If
      an object was not uploaded with `Content-Encoding: gzip` then it won't
      be set on download.

      However some providers may gzip objects even if they weren't uploaded
      with `Content-Encoding: gzip` (eg Cloudflare).

      A symptom of this would be receiving errors like

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      If you set this flag and rclone downloads an object with
      Content-Encoding: gzip set and chunked transfer encoding, then rclone
      will decompress the object on the fly.

      If this is set to unset (the default) then rclone will choose
      according to the provider setting what to apply, but you can override
      rclone's choice here.


   --s3-endpoint
      [Provider] - AWS
         Endpoint for S3 API.

         Leave blank if using AWS to use the default endpoint for the region.

      [Provider] - ChinaMobile
         Endpoint for China Mobile Ecloud Elastic Object Storage (EOS) API.

         Examples:
            | eos-wuxi-1.cmecloud.cn      | The default endpoint - a good choice if you are unsure.
                                          | East China (Suzhou)
            | eos-jinan-1.cmecloud.cn     | East China (Jinan)
            | eos-ningbo-1.cmecloud.cn    | East China (Hangzhou)
            | eos-shanghai-1.cmecloud.cn  | East China (Shanghai-1)
            | eos-zhengzhou-1.cmecloud.cn | Central China (Zhengzhou)
            | eos-hunan-1.cmecloud.cn     | Central China (Changsha-1)
            | eos-zhuzhou-1.cmecloud.cn   | Central China (Changsha-2)
            | eos-guangzhou-1.cmecloud.cn | South China (Guangzhou-2)
            | eos-dongguan-1.cmecloud.cn  | South China (Guangzhou-3)
            | eos-beijing-1.cmecloud.cn   | North China (Beijing-1)
            | eos-beijing-2.cmecloud.cn   | North China (Beijing-2)
            | eos-beijing-4.cmecloud.cn   | North China (Beijing-3)
            | eos-huhehaote-1.cmecloud.cn | North China (Huhehaote)
            | eos-chengdu-1.cmecloud.cn   | Southwest China (Chengdu)
            | eos-chongqing-1.cmecloud.cn | Southwest China (Chongqing)
            | eos-guiyang-1.cmecloud.cn   | Southwest China (Guiyang)
            | eos-xian-1.cmecloud.cn      | Nouthwest China (Xian)
            | eos-yunnan.cmecloud.cn      | Yunnan China (Kunming)
            | eos-yunnan-2.cmecloud.cn    | Yunnan China (Kunming-2)
            | eos-tianjin-1.cmecloud.cn   | Tianjin China (Tianjin)
            | eos-jilin-1.cmecloud.cn     | Jilin China (Changchun)
            | eos-hubei-1.cmecloud.cn     | Hubei China (Xiangyan)
            | eos-jiangxi-1.cmecloud.cn   | Jiangxi China (Nanchang)
            | eos-gansu-1.cmecloud.cn     | Gansu China (Lanzhou)
            | eos-shanxi-1.cmecloud.cn    | Shanxi China (Taiyuan)
            | eos-liaoning-1.cmecloud.cn  | Liaoning China (Shenyang)
            | eos-hebei-1.cmecloud.cn     | Hebei China (Shijiazhuang)
            | eos-fujian-1.cmecloud.cn    | Fujian China (Xiamen)
            | eos-guangxi-1.cmecloud.cn   | Guangxi China (Nanning)
            | eos-anhui-1.cmecloud.cn     | Anhui China (Huainan)

      [Provider] - ArvanCloud
         Endpoint for Arvan Cloud Object Storage (AOS) API.

         Examples:
            | s3.ir-thr-at1.arvanstorage.com | The default endpoint - a good choice if you are unsure.
                                             | Tehran Iran (Asiatech)
            | s3.ir-tbz-sh1.arvanstorage.com | Tabriz Iran (Shahriar)

      [Provider] - IBMCOS
         Endpoint for IBM COS S3 API.

         Specify if using an IBM COS On Premise.

         Examples:
            | s3.us.cloud-object-storage.appdomain.cloud               | US Cross Region Endpoint
            | s3.dal.us.cloud-object-storage.appdomain.cloud           | US Cross Region Dallas Endpoint
            | s3.wdc.us.cloud-object-storage.appdomain.cloud           | US Cross Region Washington DC Endpoint
            | s3.sjc.us.cloud-object-storage.appdomain.cloud           | US Cross Region San Jose Endpoint
            | s3.private.us.cloud-object-storage.appdomain.cloud       | US Cross Region Private Endpoint
            | s3.private.dal.us.cloud-object-storage.appdomain.cloud   | US Cross Region Dallas Private Endpoint
            | s3.private.wdc.us.cloud-object-storage.appdomain.cloud   | US Cross Region Washington DC Private Endpoint
            | s3.private.sjc.us.cloud-object-storage.appdomain.cloud   | US Cross Region San Jose Private Endpoint
            | s3.us-east.cloud-object-storage.appdomain.cloud          | US Region East Endpoint
            | s3.private.us-east.cloud-object-storage.appdomain.cloud  | US Region East Private Endpoint
            | s3.us-south.cloud-object-storage.appdomain.cloud         | US Region South Endpoint
            | s3.private.us-south.cloud-object-storage.appdomain.cloud | US Region South Private Endpoint
            | s3.eu.cloud-object-storage.appdomain.cloud               | EU Cross Region Endpoint
            | s3.fra.eu.cloud-object-storage.appdomain.cloud           | EU Cross Region Frankfurt Endpoint
            | s3.mil.eu.cloud-object-storage.appdomain.cloud           | EU Cross Region Milan Endpoint
            | s3.ams.eu.cloud-object-storage.appdomain.cloud           | EU Cross Region Amsterdam Endpoint
            | s3.private.eu.cloud-object-storage.appdomain.cloud       | EU Cross Region Private Endpoint
            | s3.private.fra.eu.cloud-object-storage.appdomain.cloud   | EU Cross Region Frankfurt Private Endpoint
            | s3.private.mil.eu.cloud-object-storage.appdomain.cloud   | EU Cross Region Milan Private Endpoint
            | s3.private.ams.eu.cloud-object-storage.appdomain.cloud   | EU Cross Region Amsterdam Private Endpoint
            | s3.eu-gb.cloud-object-storage.appdomain.cloud            | Great Britain Endpoint
            | s3.private.eu-gb.cloud-object-storage.appdomain.cloud    | Great Britain Private Endpoint
            | s3.eu-de.cloud-object-storage.appdomain.cloud            | EU Region DE Endpoint
            | s3.private.eu-de.cloud-object-storage.appdomain.cloud    | EU Region DE Private Endpoint
            | s3.ap.cloud-object-storage.appdomain.cloud               | APAC Cross Regional Endpoint
            | s3.tok.ap.cloud-object-storage.appdomain.cloud           | APAC Cross Regional Tokyo Endpoint
            | s3.hkg.ap.cloud-object-storage.appdomain.cloud           | APAC Cross Regional HongKong Endpoint
            | s3.seo.ap.cloud-object-storage.appdomain.cloud           | APAC Cross Regional Seoul Endpoint
            | s3.private.ap.cloud-object-storage.appdomain.cloud       | APAC Cross Regional Private Endpoint
            | s3.private.tok.ap.cloud-object-storage.appdomain.cloud   | APAC Cross Regional Tokyo Private Endpoint
            | s3.private.hkg.ap.cloud-object-storage.appdomain.cloud   | APAC Cross Regional HongKong Private Endpoint
            | s3.private.seo.ap.cloud-object-storage.appdomain.cloud   | APAC Cross Regional Seoul Private Endpoint
            | s3.jp-tok.cloud-object-storage.appdomain.cloud           | APAC Region Japan Endpoint
            | s3.private.jp-tok.cloud-object-storage.appdomain.cloud   | APAC Region Japan Private Endpoint
            | s3.au-syd.cloud-object-storage.appdomain.cloud           | APAC Region Australia Endpoint
            | s3.private.au-syd.cloud-object-storage.appdomain.cloud   | APAC Region Australia Private Endpoint
            | s3.ams03.cloud-object-storage.appdomain.cloud            | Amsterdam Single Site Endpoint
            | s3.private.ams03.cloud-object-storage.appdomain.cloud    | Amsterdam Single Site Private Endpoint
            | s3.che01.cloud-object-storage.appdomain.cloud            | Chennai Single Site Endpoint
            | s3.private.che01.cloud-object-storage.appdomain.cloud    | Chennai Single Site Private Endpoint
            | s3.mel01.cloud-object-storage.appdomain.cloud            | Melbourne Single Site Endpoint
            | s3.private.mel01.cloud-object-storage.appdomain.cloud    | Melbourne Single Site Private Endpoint
            | s3.osl01.cloud-object-storage.appdomain.cloud            | Oslo Single Site Endpoint
            | s3.private.osl01.cloud-object-storage.appdomain.cloud    | Oslo Single Site Private Endpoint
            | s3.tor01.cloud-object-storage.appdomain.cloud            | Toronto Single Site Endpoint
            | s3.private.tor01.cloud-object-storage.appdomain.cloud    | Toronto Single Site Private Endpoint
            | s3.seo01.cloud-object-storage.appdomain.cloud            | Seoul Single Site Endpoint
            | s3.private.seo01.cloud-object-storage.appdomain.cloud    | Seoul Single Site Private Endpoint
            | s3.mon01.cloud-object-storage.appdomain.cloud            | Montreal Single Site Endpoint
            | s3.private.mon01.cloud-object-storage.appdomain.cloud    | Montreal Single Site Private Endpoint
            | s3.mex01.cloud-object-storage.appdomain.cloud            | Mexico Single Site Endpoint
            | s3.private.mex01.cloud-object-storage.appdomain.cloud    | Mexico Single Site Private Endpoint
            | s3.sjc04.cloud-object-storage.appdomain.cloud            | San Jose Single Site Endpoint
            | s3.private.sjc04.cloud-object-storage.appdomain.cloud    | San Jose Single Site Private Endpoint
            | s3.mil01.cloud-object-storage.appdomain.cloud            | Milan Single Site Endpoint
            | s3.private.mil01.cloud-object-storage.appdomain.cloud    | Milan Single Site Private Endpoint
            | s3.hkg02.cloud-object-storage.appdomain.cloud            | Hong Kong Single Site Endpoint
            | s3.private.hkg02.cloud-object-storage.appdomain.cloud    | Hong Kong Single Site Private Endpoint
            | s3.par01.cloud-object-storage.appdomain.cloud            | Paris Single Site Endpoint
            | s3.private.par01.cloud-object-storage.appdomain.cloud    | Paris Single Site Private Endpoint
            | s3.sng01.cloud-object-storage.appdomain.cloud            | Singapore Single Site Endpoint
            | s3.private.sng01.cloud-object-storage.appdomain.cloud    | Singapore Single Site Private Endpoint

      [Provider] - IONOS
         Endpoint for IONOS S3 Object Storage.

         Specify the endpoint from the same region.

         Examples:
            | s3-eu-central-1.ionoscloud.com | Frankfurt, Germany
            | s3-eu-central-2.ionoscloud.com | Berlin, Germany
            | s3-eu-south-2.ionoscloud.com   | Logrono, Spain

      [Provider] - Liara
         Endpoint for Liara Object Storage API.

         Examples:
            | storage.iran.liara.space | The default endpoint
                                       | Iran

      [Provider] - Alibaba
         Endpoint for OSS API.

         Examples:
            | oss-accelerate.aliyuncs.com          | Global Accelerate
            | oss-accelerate-overseas.aliyuncs.com | Global Accelerate (outside mainland China)
            | oss-cn-hangzhou.aliyuncs.com         | East China 1 (Hangzhou)
            | oss-cn-shanghai.aliyuncs.com         | East China 2 (Shanghai)
            | oss-cn-qingdao.aliyuncs.com          | North China 1 (Qingdao)
            | oss-cn-beijing.aliyuncs.com          | North China 2 (Beijing)
            | oss-cn-zhangjiakou.aliyuncs.com      | North China 3 (Zhangjiakou)
            | oss-cn-huhehaote.aliyuncs.com        | North China 5 (Hohhot)
            | oss-cn-wulanchabu.aliyuncs.com       | North China 6 (Ulanqab)
            | oss-cn-shenzhen.aliyuncs.com         | South China 1 (Shenzhen)
            | oss-cn-heyuan.aliyuncs.com           | South China 2 (Heyuan)
            | oss-cn-guangzhou.aliyuncs.com        | South China 3 (Guangzhou)
            | oss-cn-chengdu.aliyuncs.com          | West China 1 (Chengdu)
            | oss-cn-hongkong.aliyuncs.com         | Hong Kong (Hong Kong)
            | oss-us-west-1.aliyuncs.com           | US West 1 (Silicon Valley)
            | oss-us-east-1.aliyuncs.com           | US East 1 (Virginia)
            | oss-ap-southeast-1.aliyuncs.com      | Southeast Asia Southeast 1 (Singapore)
            | oss-ap-southeast-2.aliyuncs.com      | Asia Pacific Southeast 2 (Sydney)
            | oss-ap-southeast-3.aliyuncs.com      | Southeast Asia Southeast 3 (Kuala Lumpur)
            | oss-ap-southeast-5.aliyuncs.com      | Asia Pacific Southeast 5 (Jakarta)
            | oss-ap-northeast-1.aliyuncs.com      | Asia Pacific Northeast 1 (Japan)
            | oss-ap-south-1.aliyuncs.com          | Asia Pacific South 1 (Mumbai)
            | oss-eu-central-1.aliyuncs.com        | Central Europe 1 (Frankfurt)
            | oss-eu-west-1.aliyuncs.com           | West Europe (London)
            | oss-me-east-1.aliyuncs.com           | Middle East 1 (Dubai)

      [Provider] - HuaweiOBS
         Endpoint for OBS API.

         Examples:
            | obs.af-south-1.myhuaweicloud.com     | AF-Johannesburg
            | obs.ap-southeast-2.myhuaweicloud.com | AP-Bangkok
            | obs.ap-southeast-3.myhuaweicloud.com | AP-Singapore
            | obs.cn-east-3.myhuaweicloud.com      | CN East-Shanghai1
            | obs.cn-east-2.myhuaweicloud.com      | CN East-Shanghai2
            | obs.cn-north-1.myhuaweicloud.com     | CN North-Beijing1
            | obs.cn-north-4.myhuaweicloud.com     | CN North-Beijing4
            | obs.cn-south-1.myhuaweicloud.com     | CN South-Guangzhou
            | obs.ap-southeast-1.myhuaweicloud.com | CN-Hong Kong
            | obs.sa-argentina-1.myhuaweicloud.com | LA-Buenos Aires1
            | obs.sa-peru-1.myhuaweicloud.com      | LA-Lima1
            | obs.na-mexico-1.myhuaweicloud.com    | LA-Mexico City1
            | obs.sa-chile-1.myhuaweicloud.com     | LA-Santiago2
            | obs.sa-brazil-1.myhuaweicloud.com    | LA-Sao Paulo1
            | obs.ru-northwest-2.myhuaweicloud.com | RU-Moscow2

      [Provider] - Scaleway
         Endpoint for Scaleway Object Storage.

         Examples:
            | s3.nl-ams.scw.cloud | Amsterdam Endpoint
            | s3.fr-par.scw.cloud | Paris Endpoint
            | s3.pl-waw.scw.cloud | Warsaw Endpoint

      [Provider] - StackPath
         Endpoint for StackPath Object Storage.

         Examples:
            | s3.us-east-2.stackpathstorage.com    | US East Endpoint
            | s3.us-west-1.stackpathstorage.com    | US West Endpoint
            | s3.eu-central-1.stackpathstorage.com | EU Endpoint

      [Provider] - Storj
         Endpoint for Storj Gateway.

         Examples:
            | gateway.storjshare.io | Global Hosted Gateway

      [Provider] - TencentCOS
         Endpoint for Tencent COS API.

         Examples:
            | cos.ap-beijing.myqcloud.com       | Beijing Region
            | cos.ap-nanjing.myqcloud.com       | Nanjing Region
            | cos.ap-shanghai.myqcloud.com      | Shanghai Region
            | cos.ap-guangzhou.myqcloud.com     | Guangzhou Region
            | cos.ap-nanjing.myqcloud.com       | Nanjing Region
            | cos.ap-chengdu.myqcloud.com       | Chengdu Region
            | cos.ap-chongqing.myqcloud.com     | Chongqing Region
            | cos.ap-hongkong.myqcloud.com      | Hong Kong (China) Region
            | cos.ap-singapore.myqcloud.com     | Singapore Region
            | cos.ap-mumbai.myqcloud.com        | Mumbai Region
            | cos.ap-seoul.myqcloud.com         | Seoul Region
            | cos.ap-bangkok.myqcloud.com       | Bangkok Region
            | cos.ap-tokyo.myqcloud.com         | Tokyo Region
            | cos.na-siliconvalley.myqcloud.com | Silicon Valley Region
            | cos.na-ashburn.myqcloud.com       | Virginia Region
            | cos.na-toronto.myqcloud.com       | Toronto Region
            | cos.eu-frankfurt.myqcloud.com     | Frankfurt Region
            | cos.eu-moscow.myqcloud.com        | Moscow Region
            | cos.accelerate.myqcloud.com       | Use Tencent COS Accelerate Endpoint

      [Provider] - RackCorp
         Endpoint for RackCorp Object Storage.

         Examples:
            | s3.rackcorp.com           | Global (AnyCast) Endpoint
            | au.s3.rackcorp.com        | Australia (Anycast) Endpoint
            | au-nsw.s3.rackcorp.com    | Sydney (Australia) Endpoint
            | au-qld.s3.rackcorp.com    | Brisbane (Australia) Endpoint
            | au-vic.s3.rackcorp.com    | Melbourne (Australia) Endpoint
            | au-wa.s3.rackcorp.com     | Perth (Australia) Endpoint
            | ph.s3.rackcorp.com        | Manila (Philippines) Endpoint
            | th.s3.rackcorp.com        | Bangkok (Thailand) Endpoint
            | hk.s3.rackcorp.com        | HK (Hong Kong) Endpoint
            | mn.s3.rackcorp.com        | Ulaanbaatar (Mongolia) Endpoint
            | kg.s3.rackcorp.com        | Bishkek (Kyrgyzstan) Endpoint
            | id.s3.rackcorp.com        | Jakarta (Indonesia) Endpoint
            | jp.s3.rackcorp.com        | Tokyo (Japan) Endpoint
            | sg.s3.rackcorp.com        | SG (Singapore) Endpoint
            | de.s3.rackcorp.com        | Frankfurt (Germany) Endpoint
            | us.s3.rackcorp.com        | USA (AnyCast) Endpoint
            | us-east-1.s3.rackcorp.com | New York (USA) Endpoint
            | us-west-1.s3.rackcorp.com | Freemont (USA) Endpoint
            | nz.s3.rackcorp.com        | Auckland (New Zealand) Endpoint

      [Provider] - Qiniu
         Endpoint for Qiniu Object Storage.

         Examples:
            | s3-cn-east-1.qiniucs.com      | East China Endpoint 1
            | s3-cn-east-2.qiniucs.com      | East China Endpoint 2
            | s3-cn-north-1.qiniucs.com     | North China Endpoint 1
            | s3-cn-south-1.qiniucs.com     | South China Endpoint 1
            | s3-us-north-1.qiniucs.com     | North America Endpoint 1
            | s3-ap-southeast-1.qiniucs.com | Southeast Asia Endpoint 1
            | s3-ap-northeast-1.qiniucs.com | Northeast Asia Endpoint 1

      [Provider] - Ceph, Cloudflare, Minio
         Endpoint for S3 API.

         Required when using an S3 clone.

         Examples:
            | objects-us-east-1.dream.io              | Dream Objects endpoint
            | syd1.digitaloceanspaces.com             | DigitalOcean Spaces Sydney 1
            | sfo3.digitaloceanspaces.com             | DigitalOcean Spaces San Francisco 3
            | fra1.digitaloceanspaces.com             | DigitalOcean Spaces Frankfurt 1
            | nyc3.digitaloceanspaces.com             | DigitalOcean Spaces New York 3
            | ams3.digitaloceanspaces.com             | DigitalOcean Spaces Amsterdam 3
            | sgp1.digitaloceanspaces.com             | DigitalOcean Spaces Singapore 1
            | localhost:8333                          | SeaweedFS S3 localhost
            | s3.us-east-1.lyvecloud.seagate.com      | Seagate Lyve Cloud US East 1 (Virginia)
            | s3.us-west-1.lyvecloud.seagate.com      | Seagate Lyve Cloud US West 1 (California)
            | s3.ap-southeast-1.lyvecloud.seagate.com | Seagate Lyve Cloud AP Southeast 1 (Singapore)
            | s3.wasabisys.com                        | Wasabi US East 1 (N. Virginia)
            | s3.us-east-2.wasabisys.com              | Wasabi US East 2 (N. Virginia)
            | s3.us-central-1.wasabisys.com           | Wasabi US Central 1 (Texas)
            | s3.us-west-1.wasabisys.com              | Wasabi US West 1 (Oregon)
            | s3.ca-central-1.wasabisys.com           | Wasabi CA Central 1 (Toronto)
            | s3.eu-central-1.wasabisys.com           | Wasabi EU Central 1 (Amsterdam)
            | s3.eu-central-2.wasabisys.com           | Wasabi EU Central 2 (Frankfurt)
            | s3.eu-west-1.wasabisys.com              | Wasabi EU West 1 (London)
            | s3.eu-west-2.wasabisys.com              | Wasabi EU West 2 (Paris)
            | s3.ap-northeast-1.wasabisys.com         | Wasabi AP Northeast 1 (Tokyo) endpoint
            | s3.ap-northeast-2.wasabisys.com         | Wasabi AP Northeast 2 (Osaka) endpoint
            | s3.ap-southeast-1.wasabisys.com         | Wasabi AP Southeast 1 (Singapore)
            | s3.ap-southeast-2.wasabisys.com         | Wasabi AP Southeast 2 (Sydney)
            | storage.iran.liara.space                | Liara Iran endpoint
            | s3.ir-thr-at1.arvanstorage.com          | ArvanCloud Tehran Iran (Asiatech) endpoint

   --s3-requester-pays
      [Provider] - AWS
         Enables requester pays option when interacting with S3 bucket.

   --s3-sse-kms-key-id
      [Provider] - AWS, Ceph, Minio
         If using KMS ID you must provide the ARN of Key.

         Examples:
            | <unset>                 | None
            | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --s3-sse-customer-key-base64
      [Provider] - AWS, Ceph, ChinaMobile, Minio
         If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.

         Alternatively you can provide --sse-customer-key.

         Examples:
            | <unset> | None

   --s3-chunk-size
      Chunk size to use for uploading.

      When uploading files larger than upload_cutoff or files with unknown
      size (e.g. from "rclone rcat" or uploaded with "rclone mount" or google
      photos or google docs) they will be uploaded as multipart uploads
      using this chunk size.

      Note that "--s3-upload-concurrency" chunks of this size are buffered
      in memory per transfer.

      If you are transferring large files over high-speed links and you have
      enough memory, then increasing this will speed up the transfers.

      Rclone will automatically increase the chunk size when uploading a
      large file of known size to stay below the 10,000 chunks limit.

      Files of unknown size are uploaded with the configured
      chunk_size. Since the default chunk size is 5 MiB and there can be at
      most 10,000 chunks, this means that by default the maximum size of
      a file you can stream upload is 48 GiB.  If you wish to stream upload
      larger files then you will need to increase chunk_size.

      Increasing the chunk size decreases the accuracy of the progress
      statistics displayed with "-P" flag. Rclone treats chunk as sent when
      it's buffered by the AWS SDK, when in fact it may still be uploading.
      A bigger chunk size means a bigger AWS SDK buffer and progress
      reporting more deviating from the truth.


   --s3-disable-checksum
      Don't store MD5 checksum with object metadata.

      Normally rclone will calculate the MD5 checksum of the input before
      uploading it so it can add it to metadata on the object. This is great
      for data integrity checking but can cause long delays for large files
      to start uploading.

   --s3-no-system-metadata
      Suppress setting and reading of system metadata

   --s3-acl
      [Provider] - RackCorp, HuaweiOBS, Qiniu, AWS, Scaleway, IONOS, ArvanCloud, StackPath, Ceph, Alibaba, TencentCOS, Minio, ChinaMobile, IBMCOS, Liara
         Canned ACL used when creating buckets and storing or copying objects.

         This ACL is used for creating objects and if bucket_acl isn't set, for creating buckets too.

         For more info visit https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

         Note that this ACL is applied when server-side copying objects as S3
         doesn't copy the ACL from the source but rather writes a fresh one.

         If the acl is an empty string then no X-Amz-Acl: header is added and
         the default (private) will be used.


         Examples:
            | default                   | Owner gets Full_CONTROL.
                                        | No one else has access rights (default).
            | private                   | Owner gets FULL_CONTROL.
                                        | No one else has access rights (default).
            | public-read               | Owner gets FULL_CONTROL.
                                        | The AllUsers group gets READ access.
            | public-read-write         | Owner gets FULL_CONTROL.
                                        | The AllUsers group gets READ and WRITE access.
                                        | Granting this on a bucket is generally not recommended.
            | authenticated-read        | Owner gets FULL_CONTROL.
                                        | The AuthenticatedUsers group gets READ access.
            | bucket-owner-read         | Object owner gets FULL_CONTROL.
                                        | Bucket owner gets READ access.
                                        | If you specify this canned ACL when creating a bucket, Amazon S3 ignores it.
            | bucket-owner-full-control | Both the object owner and the bucket owner get FULL_CONTROL over the object.
                                        | If you specify this canned ACL when creating a bucket, Amazon S3 ignores it.
            | private                   | Owner gets FULL_CONTROL.
                                        | No one else has access rights (default).
                                        | This acl is available on IBM Cloud (Infra), IBM Cloud (Storage), On-Premise COS.
            | public-read               | Owner gets FULL_CONTROL.
                                        | The AllUsers group gets READ access.
                                        | This acl is available on IBM Cloud (Infra), IBM Cloud (Storage), On-Premise IBM COS.
            | public-read-write         | Owner gets FULL_CONTROL.
                                        | The AllUsers group gets READ and WRITE access.
                                        | This acl is available on IBM Cloud (Infra), On-Premise IBM COS.
            | authenticated-read        | Owner gets FULL_CONTROL.
                                        | The AuthenticatedUsers group gets READ access.
                                        | Not supported on Buckets.
                                        | This acl is available on IBM Cloud (Infra) and On-Premise IBM COS.

   --s3-force-path-style
      If true use path style access if false use virtual hosted style.

      If this is true (the default) then rclone will use path style access,
      if false then rclone will use virtual path style. See [the AWS S3
      docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      for more info.

      Some providers (e.g. AWS, Aliyun OSS, Netease COS, or Tencent COS) require this set to
      false - rclone will do this automatically based on the provider
      setting.

   --s3-list-url-encode
      Whether to url encode listings: true/false/unset

      Some providers support URL encoding listings and where this is
      available this is more reliable when using control characters in file
      names. If this is set to unset (the default) then rclone will choose
      according to the provider setting what to apply, but you can override
      rclone's choice here.


   --s3-encoding
      The encoding for the backend.

      See the [encoding section in the overview](/overview/#encoding) for more info.

   --s3-version-at
      Show file versions as they were at the specified time.

      The parameter should be a date, "2006-01-02", datetime "2006-01-02
      15:04:05" or a duration for that long ago, eg "100d" or "1h".

      Note that when using this no file write operations are permitted,
      so you can't upload files or delete them.

      See [the time option docs](/docs/#time-option) for valid formats.



OPTIONS:
   --help, -h                         show help
   --s3-access-key-id value           AWS Access Key ID. [$S3_ACCESS_KEY_ID]
   --s3-acl value                     Canned ACL used when creating buckets and storing or copying objects. [$S3_ACL]
   --s3-endpoint value                Endpoint for S3 API. [$S3_ENDPOINT]
   --s3-env-auth value                Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars). (default: "false") [$S3_ENV_AUTH]
   --s3-location-constraint value     Location constraint - must be set to match the Region. [$S3_LOCATION_CONSTRAINT]
   --s3-provider value                Choose your S3 provider. [$S3_PROVIDER]
   --s3-region value                  Region to connect to. [$S3_REGION]
   --s3-secret-access-key value       AWS Secret Access Key (password). [$S3_SECRET_ACCESS_KEY]
   --s3-server-side-encryption value  The server-side encryption algorithm used when storing this object in S3. [$S3_SERVER_SIDE_ENCRYPTION]
   --s3-sse-kms-key-id value          If using KMS ID you must provide the ARN of Key. [$S3_SSE_KMS_KEY_ID]
   --s3-storage-class value           The storage class to use when storing new objects in S3. [$S3_STORAGE_CLASS]

   Advanced Options

   --s3-bucket-acl value               Canned ACL used when creating buckets. [$S3_BUCKET_ACL]
   --s3-chunk-size value               Chunk size to use for uploading. (default: "5Mi") [$S3_CHUNK_SIZE]
   --s3-copy-cutoff value              Cutoff for switching to multipart copy. (default: "4.656Gi") [$S3_COPY_CUTOFF]
   --s3-decompress value               If set this will decompress gzip encoded objects. (default: "false") [$S3_DECOMPRESS]
   --s3-disable-checksum value         Don't store MD5 checksum with object metadata. (default: "false") [$S3_DISABLE_CHECKSUM]
   --s3-disable-http2 value            Disable usage of http2 for S3 backends. (default: "false") [$S3_DISABLE_HTTP2]
   --s3-download-url value             Custom endpoint for downloads. [$S3_DOWNLOAD_URL]
   --s3-encoding value                 The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$S3_ENCODING]
   --s3-force-path-style value         If true use path style access if false use virtual hosted style. (default: "true") [$S3_FORCE_PATH_STYLE]
   --s3-leave-parts-on-error value     If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery. (default: "false") [$S3_LEAVE_PARTS_ON_ERROR]
   --s3-list-chunk value               Size of listing chunk (response list for each ListObject S3 request). (default: "1000") [$S3_LIST_CHUNK]
   --s3-list-url-encode value          Whether to url encode listings: true/false/unset (default: "unset") [$S3_LIST_URL_ENCODE]
   --s3-list-version value             Version of ListObjects to use: 1,2 or 0 for auto. (default: "0") [$S3_LIST_VERSION]
   --s3-max-upload-parts value         Maximum number of parts in a multipart upload. (default: "10000") [$S3_MAX_UPLOAD_PARTS]
   --s3-memory-pool-flush-time value   How often internal memory buffer pools will be flushed. (default: "1m0s") [$S3_MEMORY_POOL_FLUSH_TIME]
   --s3-memory-pool-use-mmap value     Whether to use mmap buffers in internal memory pool. (default: "false") [$S3_MEMORY_POOL_USE_MMAP]
   --s3-might-gzip value               Set this if the backend might gzip objects. (default: "unset") [$S3_MIGHT_GZIP]
   --s3-no-check-bucket value          If set, don't attempt to check the bucket exists or create it. (default: "false") [$S3_NO_CHECK_BUCKET]
   --s3-no-head value                  If set, don't HEAD uploaded objects to check integrity. (default: "false") [$S3_NO_HEAD]
   --s3-no-head-object value           If set, do not do HEAD before GET when getting objects. (default: "false") [$S3_NO_HEAD_OBJECT]
   --s3-no-system-metadata value       Suppress setting and reading of system metadata (default: "false") [$S3_NO_SYSTEM_METADATA]
   --s3-profile value                  Profile to use in the shared credentials file. [$S3_PROFILE]
   --s3-requester-pays value           Enables requester pays option when interacting with S3 bucket. (default: "false") [$S3_REQUESTER_PAYS]
   --s3-session-token value            An AWS session token. [$S3_SESSION_TOKEN]
   --s3-shared-credentials-file value  Path to the shared credentials file. [$S3_SHARED_CREDENTIALS_FILE]
   --s3-sse-customer-algorithm value   If using SSE-C, the server-side encryption algorithm used when storing this object in S3. [$S3_SSE_CUSTOMER_ALGORITHM]
   --s3-sse-customer-key value         To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data. [$S3_SSE_CUSTOMER_KEY]
   --s3-sse-customer-key-base64 value  If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data. [$S3_SSE_CUSTOMER_KEY_BASE64]
   --s3-sse-customer-key-md5 value     If using SSE-C you may provide the secret encryption key MD5 checksum (optional). [$S3_SSE_CUSTOMER_KEY_MD5]
   --s3-sts-endpoint value             Endpoint for STS. [$S3_STS_ENDPOINT]
   --s3-upload-concurrency value       Concurrency for multipart uploads. (default: "4") [$S3_UPLOAD_CONCURRENCY]
   --s3-upload-cutoff value            Cutoff for switching to chunked upload. (default: "200Mi") [$S3_UPLOAD_CUTOFF]
   --s3-use-accelerate-endpoint value  If true use the AWS S3 accelerated endpoint. (default: "false") [$S3_USE_ACCELERATE_ENDPOINT]
   --s3-use-multipart-etag value       Whether to use ETag in multipart uploads for verification (default: "unset") [$S3_USE_MULTIPART_ETAG]
   --s3-use-presigned-request value    Whether to use a presigned request or PutObject for single part uploads (default: "false") [$S3_USE_PRESIGNED_REQUEST]
   --s3-v2-auth value                  If true use v2 authentication. (default: "false") [$S3_V2_AUTH]
   --s3-version-at value               Show file versions as they were at the specified time. (default: "off") [$S3_VERSION_AT]
   --s3-versions value                 Include old versions in directory listings. (default: "false") [$S3_VERSIONS]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)

```
{% endcode %}
