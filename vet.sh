#!/bin/sh
PROTOBUF_VERSION=3.3.0
PROTOC_FILENAME=protoc-${PROTOBUF_VERSION}-linux-x86_64.zip

curl -OL https://github.com/google/protobuf/releases/download/v${PROTOBUF_VERSION}/${PROTOC_FILENAME}

# Unzip
unzip ${PROTOC_FILENAME} -d protoc3

# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/

# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/

# Optional: change owner
sudo chown $USER /usr/local/bin/protoc
sudo chown -R $USER /usr/local/include/google