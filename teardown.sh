#!/usr/bin/env bash
kubectl delete -f admissioncontroller-config.yaml
kubectl delete -f admissioncontroller.yaml


kubectl delete secret ca-secret
kubectl delete secret admissioncontroller
