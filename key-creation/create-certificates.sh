#!/usr/bin/env bash


openssl genrsa > ca.key
openssl req -new -config ca.conf -new -key ca.key -out ca.csr
openssl x509 -req -days 1095 -in ca.csr -signkey ca.key -out ca.crt 

openssl genrsa > admissioncontroller.key
openssl req -new -config admissioncontroller.conf -new -key admissioncontroller.key -out admissioncontroller.csr
openssl x509 -req -days 360 -in admissioncontroller.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out admissioncontroller.crt -sha256

openssl genrsa > client.key
openssl req -new -config clientcsr.conf -key client.key -out client.csr
openssl x509 -req -days 360 -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -sha256


