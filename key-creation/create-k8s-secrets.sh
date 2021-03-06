#!/usr/bin/env bash

cd $(dirname $0)

kubectl create secret tls admissioncontroller --key admissioncontroller.key --cert admissioncontroller.crt

kubectl create secret generic ca-secret --from-file=ca.crt=ca.crt 
