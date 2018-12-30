# Admissioncontroller

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

