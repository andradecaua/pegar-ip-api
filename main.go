package main

import (
	"log"
	"net/http"
	"pegar-ip-ou-addr/controller"
)

func main() {

	server := &http.Server{
		Addr: ":80",
	}

	http.HandleFunc("/", controller.IndexControler)
	http.HandleFunc("/v1/get-ip", controller.IpController)
	http.HandleFunc("/v1/get-nameservers", controller.NameServerController)

	log.Fatal(server.ListenAndServe(), nil)

}
