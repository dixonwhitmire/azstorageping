# azstorageping

Connects to an azure storage account by access key and lists container contents.

## Quickstart

To run from source:
```
make run account=<account name> key=<access key> container=<container name>

2023/04/21 15:23:23 azstorageping.pingAccount: connected to https://<some account>.blob.core.windows.net
2023/04/21 15:23:23 azstorageping.pingAccount: checking access to container <container name>
2023/04/21 15:23:23 azstorageping.pingAccount: listing blobs
2023/04/21 15:23:23 sample-file
2023/04/21 15:23:23 some-demographic.csv
2023/04/21 15:23:23 flat-file.dat
2023/04/21 15:23:23 azstorageping.main: ping complete
```

To build:
```
make build
```

The output executable is created in target/azstorageping