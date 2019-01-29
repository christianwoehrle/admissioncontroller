package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/api/core/v1"
)

var (
	tlsCertFile string
	tlsKeyFile  string
)

type auditLogEntry struct {
	Time      string `json:"time"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Operation string `json:"operation"`
	User      string `json:"user"`
}

func main() {
	flag.StringVar(&tlsCertFile, "tls-cert", "/etc/admission-controller/tls/cert.pem", "TLS certificate file.")
	flag.StringVar(&tlsKeyFile, "tls-key", "/etc/admission-controller/tls/key.pem", "TLS key file.")
	flag.Parse()

	startLog := &auditLogEntry{
		Time:      time.Now().Format(time.RFC3339),
		Kind:      "Pod",
		Name:      "AufitLog",
		Namespace: "default",
		Operation: "CREATE",
		User:      "NA",
	}
	startLogJson, _ := json.Marshal(startLog)
	fmt.Println(string(startLogJson))

	http.HandleFunc("/", admissionReviewHandler)
	s := http.Server{
		Addr: ":8443",
		TLSConfig: &tls.Config{
			ClientAuth: tls.NoClientCert,
		},
	}

	fmt.Println("Starting WebServer 8443  ------")
	log.Fatal(s.ListenAndServeTLS(tlsCertFile, tlsKeyFile))
}

func admissionReviewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nCalled admissionReviewHandler")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//log.Println(string(data))

	ar := v1beta1.AdmissionReview{}
	if err := json.Unmarshal(data, &ar); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ale := &auditLogEntry{
		Time:      time.Now().Format(time.RFC3339),
		Kind:      ar.Request.Kind.Kind,
		Name:      ar.Request.Name,
		Namespace: ar.Request.Namespace,
		Operation: string(ar.Request.Operation),
		User:      ar.Request.UserInfo.Username,
	}
	//fmt.Printf("Kind: %s\nName: %s\nNamespace: %s\nUser: %s\n", ar.Request.Kind, ar.Request.Name, ar.Request.Namespace, ar.Request.UserInfo.Username)
	aleJson, _ := json.Marshal(ale)
	fmt.Println(string(aleJson))

	admissionResponse := v1beta1.AdmissionResponse{Allowed: false}

	ar = v1beta1.AdmissionReview{
		Response: &admissionResponse,
	}
	ar.Response.Allowed = true

	data, err = json.Marshal(ar)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func showPod(ar v1beta1.AdmissionReview) {

	pod := v1.Pod{}
	if err := json.Unmarshal(ar.Request.Object.Raw, &pod); err != nil {
		log.Println(err)
		//w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("Tried to start ", pod.Name, pod.Spec.NodeName)

	for _, container := range pod.Spec.Containers {
		fmt.Printf("Container: %s\n", container.Name)
	}

	for k, v := range pod.Labels {
		fmt.Printf("Label: (%s,%s)\n", k, v)
	}

}
