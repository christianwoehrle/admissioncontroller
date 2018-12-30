#!/usr/bin/env bash
THIS_SCRIPT=$(readlink -f "$0")
SCRIPTDIR=$(dirname "${THIS_SCRIPT}")
cd $SCRIPTDIR
./key-creation/create-k8s-secrets.sh

kubectl apply -f admissioncontroller-config.yaml
kubectl apply -f admissioncontroller.yaml
