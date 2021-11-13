#!/bin/bash
#
# Script file to use curl command to get user info by user id.
#

URL="http://localhost"
PORT="8000"
ENDPOINT="v1/user"
TOKEN="1234567890abcdefg"
OPTS=""
NUM_ARGS=1

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script file to use curl command to get user info by user id."
	echo
	echo "Usage: $SCRIPT_NAME [id]"
	echo "Options:"
	echo " -k                           Allow https insecure connection"
	echo " -u  [url]                    TM2 API server URL"
	echo " -p  [port]                   TM2 API server port number"
	echo " -v                           Verbose output"
	echo " -h                           This help message"
	echo
}

# Parse input argument(s)
while [ "${1:0:1}" == "-" ]; do
	OPT=${1:1:1}
	case "$OPT" in
	"k")
		OPTS="$OPTS -k"
		;;
	"u")
		URL=$2
		shift
		;;
	"p")
		PORT=$2
		shift
		;;
	"v")
		OPTS="$OPTS -v"
		;;
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

ID=$1

# trim URL trailing "/"
if [ "$PORT" = "" ]; then
	URL="$(echo -e "${URL}" | sed -e 's/\/*$//')"
else
	URL="$(echo -e "${URL}:${PORT}" | sed -e 's/\/*$//')"
fi

# perform curl to get the response
curl $OPTS \
	-H "Accept: application/json" \
	-H "Authorization: Bearer ${TOKEN}" \
	${URL}/${ENDPOINT}/${ID}
