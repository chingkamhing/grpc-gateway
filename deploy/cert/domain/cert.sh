#!/bin/bash
#
# generate self cert for domain host
#
# Reference:
# - https://adfinis.com/en/blog/openssl-x509-certificates/
#

DOMAIN=example.com
OUTPUT_PATH="deploy/cert/domain"

# generate Certificate Authority (ca) first
openssl req -x509 -new -nodes -sha256 -newkey rsa:4096 -keyout ${OUTPUT_PATH}/ca.key -out ${OUTPUT_PATH}/ca.csr -subj "/C=HK/CN=Example-Root-CA"
openssl x509 -outform pem -in ${OUTPUT_PATH}/ca.csr -out ${OUTPUT_PATH}/ca.pem -days 1095
# use the ca to create server cert and private key
openssl req -new -nodes -sha256 -newkey rsa:4096 -keyout ${OUTPUT_PATH}/${DOMAIN}.key -out ${OUTPUT_PATH}/${DOMAIN}.csr -subj "/C=HK/ST=YourState/L=YourCity/O=Example-Certificates/CN=${DOMAIN}"
openssl x509 -req -sha256 -CA ${OUTPUT_PATH}/ca.pem -CAkey ${OUTPUT_PATH}/ca.key -CAcreateserial -extfile ${OUTPUT_PATH}/${DOMAIN}.ext -in ${OUTPUT_PATH}/${DOMAIN}.csr -out ${OUTPUT_PATH}/${DOMAIN}.pem -days 365

# view the cert
openssl req -in ${OUTPUT_PATH}/${DOMAIN}.csr -noout -text
openssl x509 -in ${OUTPUT_PATH}/${DOMAIN}.pem -noout -text
