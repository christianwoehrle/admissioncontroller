package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/api/core/v1"
)

var (
	tlsCertFile string
	tlsKeyFile  string
)

func main() {
	flag.StringVar(&tlsCertFile, "tls-cert", "/etc/admission-controller/tls/cert.pem", "TLS certificate file.")
	flag.StringVar(&tlsKeyFile, "tls-key", "/etc/admission-controller/tls/key.pem", "TLS key file.")
	flag.Parse()

	fmt.Println("Start WebServer")
	http.HandleFunc("/", admissionReviewHandler)
	s := http.Server{
		Addr: ":8443",
		TLSConfig: &tls.Config{
			ClientAuth: tls.NoClientCert,
		},
	}

	fmt.Println("Starting WebServer 8443")
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

	log.Println(string(data))

	ar := v1beta1.AdmissionReview{}
	if err := json.Unmarshal(data, &ar); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pod := v1.Pod{}
	if err := json.Unmarshal(ar.Request.Object.Raw, &pod); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("Tried to start ", pod.Name, pod.Spec.NodeName)

	admissionResponse := v1beta1.AdmissionResponse{Allowed: false}
	for _, container := range pod.Spec.Containers {
		fmt.Printf("Container: %s\n", container.Name)
	}

	for k, v := range pod.Labels {
		fmt.Printf("Label: (%s,%s)\n", k, v)
	}

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
