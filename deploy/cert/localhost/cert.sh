#!/bin/bash
#
# generate self cert for localhost development
#
# Reference:
# - https://gist.github.com/cecilemuller/9492b848eb8fe46d462abeb26656c4f8
#

DOMAIN=localhost
OUTPUT_PATH="deploy/cert/localhost"

# generate Certificate Authority (ca) first
openssl req -x509 -new -nodes -sha256 -newkey rsa:2048 -keyout ${OUTPUT_PATH}/ca.key -out ${OUTPUT_PATH}/ca.csr -subj "/C=HK/CN=Example-Root-CA"
openssl x509 -outform pem -in ${OUTPUT_PATH}/ca.csr -out ${OUTPUT_PATH}/ca.pem -days 1095
# use the ca to create server cert and private key
openssl req -new -nodes -sha256 -newkey rsa:2048 -keyout ${OUTPUT_PATH}/${DOMAIN}.key -out ${OUTPUT_PATH}/${DOMAIN}.csr -subj "/C=HK/ST=YourState/L=YourCity/O=Example-Certificates/CN=localhost.local"
openssl x509 -req -sha256 -CA ${OUTPUT_PATH}/ca.pem -CAkey ${OUTPUT_PATH}/ca.key -CAcreateserial -extfile ${OUTPUT_PATH}/${DOMAIN}.ext -in ${OUTPUT_PATH}/${DOMAIN}.csr -out ${OUTPUT_PATH}/${DOMAIN}.pem -days 365

# view the cert
openssl req -in ${OUTPUT_PATH}/${DOMAIN}.csr -noout -text
openssl x509 -in ${OUTPUT_PATH}/${DOMAIN}.pem -noout -text
