#!/bin/bash
#
# Script file to use curl command to create a new user.
#

URL="http://localhost"
PORT="8000"
ENDPOINT="v1/user"
TOKEN="my-token-1234567890"
FILE_COOKIE=".cookie"
OPTS=""
NUM_ARGS=0

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script file to use curl command to create a new user."
	echo
	echo "Usage: $SCRIPT_NAME"
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

# trim URL trailing "/"
if [ "$PORT" = "" ]; then
	URL="$(echo -e "${URL}" | sed -e 's/\/*$//')"
else
	URL="$(echo -e "${URL}:${PORT}" | sed -e 's/\/*$//')"
fi

# perform curl to get the response
curl $OPTS -d "{ \
		\"value\":{ \
			\"user\":{ \
				\"userID\":\"1\", \
				\"name\":\"John Dupoint\", \
				\"gender\":\"MALE\", \
				\"email\":\"alanpoon@email.com\", \
				\"phoneNumber\":\"1234-5678\", \
				\"address\":\"John Dupoint - address\", \
				\"dateOfBirth\":946684800 \
			}, \
			\"company\":{ \
				\"companyID\":\"1\", \
				\"name\":\"Picasa Art Ltd.\", \
				\"phoneNumber\":\"1-800-1234-5678\", \
				\"address\":\"info@picasaart.com\", \
				\"bussiness\":\"art\" \
			} \
		} \
	}" \
	-H "Accept: application/json" \
	-H "Authorization: Bearer ${TOKEN}" \
	${URL}/${ENDPOINT}
