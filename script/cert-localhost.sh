#!/bin/bash
#
# generate self cert for localhost development
#
# Reference:
# - https://gist.github.com/cecilemuller/9492b848eb8fe46d462abeb26656c4f8
# - https://stackoverflow.com/questions/7580508/getting-chrome-to-accept-self-signed-localhost-certificate
#

DOMAIN=localhost
OUTPUT_PATH="deploy/cert/localhost"

# change following config when need
CAKEY="${OUTPUT_PATH}/ca.key"
CAPEM="${OUTPUT_PATH}/ca.pem"
CACRT="${OUTPUT_PATH}/ca.crt"
PRIKEY="${OUTPUT_PATH}/${DOMAIN}.key"
CSR="${OUTPUT_PATH}/${DOMAIN}.csr"
CRT="${OUTPUT_PATH}/${DOMAIN}.crt"
CONFIEXT="${OUTPUT_PATH}/${DOMAIN}.ext"

# various settings
COUNTRY="CN"
STATE="HK"
LOCATION="Local"
ORGANIZATION="Localhost"
ORGANIZATION_UNIT="Local Unit"
CN="localhost.com"
EMAIL="someone@localhost.com"
CA_DAYS=1095
CERT_DAYS=365

# subject for root certificate
SUBJECT_CA="\
/C=${COUNTRY}/ST=${STATE}/L=${LOCATION}\
/CN=${CN}-CA/emailAddress=${EMAIL}"
SUBJECT="\
/C=${COUNTRY}/ST=${STATE}/L=${LOCATION}\
/O=${ORGANIZATION}\
/CN=${CN}/emailAddress=${EMAIL}"

# generate Certificate Authority (ca) first
openssl req -x509 -new -nodes -sha256 -newkey rsa:2048 -keyout $CAKEY -out $CAPEM -subj $SUBJECT_CA
openssl x509 -outform pem -in $CAPEM -out $CACRT -days $CA_DAYS
# use the ca to create server cert and private key
openssl req -new -nodes -sha256 -newkey rsa:2048 -keyout $PRIKEY -out $CSR -subj $SUBJECT
openssl x509 -req -sha256 -CA $CAPEM -CAkey $CAKEY -CAcreateserial -extfile $CONFIEXT -in $CSR -out $CRT -days $CERT_DAYS

# view the cert
# openssl req -in $CSR -noout -text
# openssl x509 -in $CRT -noout -text

# verify
openssl verify -CAfile $CAPEM -verify_hostname localhost $CRT
openssl verify -CAfile $CAPEM -verify_ip 127.0.0.1 $CRT
