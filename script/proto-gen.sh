#!/bin/bash
#
# Script to protoc generate all proto source files.
#

DIR="$(dirname "${0}")"
source ${DIR}/global.sh

PROTO_DIR="proto"
OUTPUT_FILE_PREFIX="tm2-proto"
OUTPUT_LANG_FILE="proto-output"
IS_COMMON_JS="no"
COMMON_JS="import_style=commonjs"
NUM_ARGS=0
DEBUG=""

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script to protoc generate all proto source files."
	echo
	echo "Usage: $SCRIPT_NAME"
	echo "Options:"
	echo " -c                           For js language, output CommonJS imports"
	echo " -h                           This help message"
	echo
}

# Generate proto files
GenerateProto () {
	local src_dir=$1
	local service=$2
	local output_root_dir=$3
	[ -d ${output_root_dir} ] || mkdir -p ${output_root_dir}
	# loop for different proto output languages
	while read -r lang; do
		local output_dir=${output_root_dir}/${OUTPUT_FILE_PREFIX}-${service}-${lang}
		case "$lang" in
		"go")
			local lang_opts="--go_out ${output_root_dir} --go_opt paths=source_relative --go-grpc_out ${output_root_dir} --go-grpc_opt paths=source_relative"
			local option_gateway="--grpc-gateway_out ${output_root_dir} --grpc-gateway_opt paths=source_relative --openapiv2_out ${output_dir} --openapiv2_opt allow_merge=true,merge_file_name=${service}"
			local lang_proc="[ -f ${output_dir}/${service}_struct.pb.go ] && protoc-go-inject-tag -input ${output_dir}/${service}_struct.pb.go"
			;;
		"js")
			local js_opt=$([ "$IS_COMMON_JS" == "yes" ] && echo "$COMMON_JS" || echo "library=${service}/${service}" )
			local lang_opts="--js_out ${js_opt},binary:${output_root_dir}"
			local lang_proc=""
			;;
		*)
			# unsupport output language, skip
			continue
			;;
		esac
		echo "====> Generating proto files for service \"$service\" to language \"$lang\""
		for file in ${src_dir}/${service}/*.proto; do
			[ -d ${output_dir} ] || mkdir -p ${output_dir}
			IsFileGrpcGateway $file
			if [ "$?" -eq "0" ]; then
				local file_lang_opts="$lang_opts $option_gateway"
			else
				local file_lang_opts="$lang_opts"
			fi
			$DEBUG protoc -I ${src_dir} -I ${GOPATH}/src/github.com/googleapis/googleapis -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
				$file_lang_opts \
				$file && \
				$DEBUG mv -f ${output_root_dir}/${service}/* ${output_dir}/ && \
				$DEBUG rmdir ${output_root_dir}/${service} && \
				$DEBUG eval $lang_proc
		done
	done < "$src_dir/$service/$OUTPUT_LANG_FILE"
}

# Check if the service file contain grpc gateway option
IsFileGrpcGateway () {
	local file=$1
	grep "option.*(google.api.http)" $file > /dev/null
	return $?
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
