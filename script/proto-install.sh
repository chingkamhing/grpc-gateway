#!/bin/bash
#
# Script to install various libraries to generate the proto files.
#

# https://github.com/protocolbuffers/protobuf/releases/
PROTOC_VER="3.19.1"
PROTOC_ZIP="protoc-${PROTOC_VER}-linux-x86_64.zip"
INSTALL_DIR="${HOME}/.local"

# either 0 argument
NUM_ARGS=0

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script to install various libraries to generate the proto files."
	echo
	echo "Usage: $SCRIPT_NAME"
	echo "Options:"
	echo " -h                           This help message"
	echo
}

# Parse input argument(s)
while [ "${1:0:1}" == "-" ]; do
	OPT=${1:1:1}
	case "$OPT" in
	"h")
		Usage
		exit
		;;
	esac
	shift
done

if [ "$#" -ne "$NUM_ARGS" ]; then
    echo "Invalid parameter!"
	Usage
	exit 1
fi

# install protoc
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/${PROTOC_ZIP}
unzip -o $PROTOC_ZIP -d ${INSTALL_DIR} bin/protoc
unzip -o $PROTOC_ZIP -d ${INSTALL_DIR}/bin 'include/*'
rm -f $PROTOC_ZIP

# install protocol buffer libraries
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install github.com/favadi/protoc-go-inject-tag
