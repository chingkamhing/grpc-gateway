#!/bin/bash
#
# Script to protoc generate all proto source files.
#

source global.sh

# either 0 argument
NUM_ARGS=0

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script to protoc generate all proto source files."
	echo
	echo "Usage: $SCRIPT_NAME"
	echo "Options:"
	echo " -h                           This help message"
	echo
}

# Generate proto files
GenerateProto () {
	local src_dir=$1
	local service=$2
	local output_dir=$3
	echo "Generating proto files for service \"$service\""
	[ -d ${output_dir}/${service} ] || mkdir -p ${output_dir}/${service}
	IsServiceGateway $service
	if [ $_RESULT == true ]; then
		local option_gateway="--grpc-gateway_out ${output_dir} --grpc-gateway_opt paths=source_relative --swagger_out ${output_dir} --swagger_opt allow_merge=true,merge_file_name=${service}"
	fi
    for file in *.proto; do
		[ -d ${output_dir}/tm2-proto-${service}-go ] || mkdir -p ${output_dir}/tm2-proto-${service}-go
		protoc -I ${src_dir} -I ${GOPATH}/src/github.com/googleapis/googleapis -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
			--go_out ${output_dir} --go_opt paths=source_relative \
			--go-grpc_out ${output_dir} --go-grpc_opt paths=source_relative \
			${option_gateway} \
			${src_dir}/${service}/${file} && \
			mv ${output_dir}/${service}/* ${output_dir}/tm2-proto-${service}-go/ && \
			rmdir ${output_dir}/${service} && \
			protoc-go-inject-tag -input ${output_dir}/tm2-proto-${service}-go/${service}_struct.pb.go
	done
}

# Check if the service is a gateway
IsServiceGateway () {
	local service=$1
	for gateway in "${GATEWAY[@]}"; do
		if [ "$service" == "$gateway" ]; then
			_RESULT=true
			return
		fi
	done
	_RESULT=false
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

# generate services proto files
for service in "${SERVICES[@]}"; do
	GenerateProto $PROTO_DIR $service $OUTPUT_DIR
done
