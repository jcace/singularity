# Enterprise File Fabric

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add filefabric - Enterprise File Fabric

USAGE:
   singularity datasource add filefabric [command options] <dataset_name> <source_path>

DESCRIPTION:
   --filefabric-url
      URL of the Enterprise File Fabric to connect to.

      Examples:
         | https://storagemadeeasy.com       | Storage Made Easy US
         | https://eu.storagemadeeasy.com    | Storage Made Easy EU
         | https://yourfabric.smestorage.com | Connect to your Enterprise File Fabric

   --filefabric-root-folder-id
      ID of the root folder.

      Leave blank normally.

      Fill in to make rclone start with directory of a given ID.


   --filefabric-permanent-token
      Permanent Authentication Token.

      A Permanent Authentication Token can be created in the Enterprise File
      Fabric, on the users Dashboard under Security, there is an entry
      you'll see called "My Authentication Tokens". Click the Manage button
      to create one.

      These tokens are normally valid for several years.

      For more info see: https://docs.storagemadeeasy.com/organisationcloud/api-tokens


   --filefabric-token
      Session Token.

      This is a session token which rclone caches in the config file. It is
      usually valid for 1 hour.

      Don't set this value - rclone will set it automatically.


   --filefabric-token-expiry
      Token expiry time.

      Don't set this value - rclone will set it automatically.


   --filefabric-version
      Version read from the file fabric.

      Don't set this value - rclone will set it automatically.


   --filefabric-encoding
      The encoding for the backend.

      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --filefabric-permanent-token value  Permanent Authentication Token. [$FILEFABRIC_PERMANENT_TOKEN]
   --filefabric-root-folder-id value   ID of the root folder. [$FILEFABRIC_ROOT_FOLDER_ID]
   --filefabric-url value              URL of the Enterprise File Fabric to connect to. [$FILEFABRIC_URL]
   --help, -h                          show help

   Advanced Options

   --filefabric-encoding value      The encoding for the backend. (default: "Slash,Del,Ctl,InvalidUtf8,Dot") [$FILEFABRIC_ENCODING]
   --filefabric-token value         Session Token. [$FILEFABRIC_TOKEN]
   --filefabric-token-expiry value  Token expiry time. [$FILEFABRIC_TOKEN_EXPIRY]
   --filefabric-version value       Version read from the file fabric. [$FILEFABRIC_VERSION]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
```
{% endcode %}
