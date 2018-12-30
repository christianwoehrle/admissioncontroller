# Admissioncontroller
[![GoDoc](https://godoc.org/github.com/christianwoehrle/admissioncontroller?status.svg)](https://godoc.org/github.com/christianwoehrle/iadmissioncontroller)
[![CircleCI](https://img.shields.io/circleci/project/github/christianwoehrle/admissioncontroller.png)](https://circleci.com/gh/christianwoehrle/admissioncontroller)
[![Go Report Card](https://goreportcard.com/badge/github.com/christianwoehrle/admissioncontroller)](https://goreportcard.com/report/github.com/christianwoehrle/admissioncontroller)



Just a simple Admissioncontroller with the necessary secrets and config so that the ApiServer can actually call the Controller without TLS-Errors.

YOu can start everything with the command 

```
./setup.sh
```

and check that the AdmissionController receives call about new pod with

```
kubectl logs -f $(kubectl get po | grep admissioncontroller | gawk '{print $1}')
```

Just start a new pod and check that the admissioncontroller logs the events.





You can create new certificates with script ```key-creation/create-certificates.sh``` or just use whatś in the directory ```key-creation```.

The Files```ca[key|crt]``` form up the certificate authority that is used to sign the server certificate ```admissioncontroller.[key|crt]``` and the client certificate ```client.[key|crt]```.

If you create new certificates, you have to update the file ```webhook-config.yaml```.
The field ```caBundle``` has to contain the base64 encoded server certificate from ```admissioncontroller.crt```

You get this with 
```
base64 key-creation/admissioncontroller.crt| tr -d "\n"
```

