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

To test the AdmissionController, just start a new pod and check that the admissioncontroller logs the events.



## New certificates


You can use the certficates that are already prepared in the directory ```key-creation```.

The Files```ca[key|crt]``` form up the certificate authority that is used to sign the server certificate ```admissioncontroller.[key|crt]``` and the client certificate ```client.[key|crt]```.

Or you can create new certificates with script ```key-creation/create-certificates.sh```.
If you want to use new certificates, change the files ```*.conf``` to your linking.

One important thing to keep in mind is that the common name (CN) of the admissioncontroller-certificate must match
the DNS Name of the service of the admission controller.

I.e. the field ```commonName``` in File ```key-creation/admissioncontroller.conf``` could be ```admissioncontroller1.default.svc```
and the field ```name``` in file ```admissioncontroller.yaml``` could be like that.

```
kind: Service
metadata:
  name: admissioncontroller

``` 

Keep in mind that for every servicename a DNS-Name with the suffix ```<namespace>.svc``` is created, so in the namespace ```default``` the dns name
becomes ```admissioncontroller.default.svc``` which matched the common name.





If you create a new certificate authorization, you have to update the file ```admissioncontroller-config.yaml```.
The field ```caBundle``` has to contain the base64 encoded server certificate from ```ca.crt```

You get this with 
```
base64 key-creation/ca.crt| tr -d "\n"
```

