#!/bin/bash
#
# Script to go mod init or update.
#

DIR="$(dirname "${0}")"
source ${DIR}/global.sh

# either 0 argument
NUM_ARGS=0

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script to go mod init or update."
	echo
	echo "Usage: $SCRIPT_NAME"
	echo "Options:"
	echo " -h                           This help message"
	echo
}

# go mod init or update
GoModInitOrUpdate () {
	local output_dir=$1
	local service=$2
	if [ -d ${output_dir}/${OUTPUT_FILE_PREFIX}-${service}-go ]; then
		cd ${output_dir}/${OUTPUT_FILE_PREFIX}-${service}-go
		echo "go mod \"${service}\""
		[ ! -f go.mod ] && go mod init github.com/chingkamhing/${OUTPUT_FILE_PREFIX}-${service}-go
		go get -u .
		cd -
	fi
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
	GoModInitOrUpdate $OUTPUT_DIR $service
done
