package main

import (
	"PegarIpOuAddres/pegarIp"
	"PegarIpOuAddres/pegarNameServer"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {

	server := &http.Server{
		Addr: ":3000",
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public/teste.html"))))
	http.HandleFunc("/v1/get-ip", handleGetIp)
	http.HandleFunc("/v1/get-nameservers", handleGetNameServers)

	log.Fatal(server.ListenAndServe(), nil)
}

func handleGetNameServers(resWriter http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		resWriter.WriteHeader(500)
	}

	url := req.URL.Query().Get("url")
	urlTransform := strings.Split(url, "'")
	url = urlTransform[1]

	nameServers, err := pegarNameServer.PegarNameServer(url)

	if err != nil {
		resWriter.WriteHeader(403)
		io.WriteString(resWriter, err.Error())
	} else {
		data, errConverter := json.Marshal(nameServers)

		if errConverter != nil {
			resWriter.WriteHeader(400)
			io.WriteString(resWriter, "Não fo possível converter para JSON")
		} else {
			resWriter.WriteHeader(200)
			resWriter.Header().Set("ContentType", "application/json")
			resWriter.Write(data)
		}

	}
}

func handleGetIp(resWriter http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		resWriter.WriteHeader(500)
	}

	url := req.URL.Query().Get("url")
	urlTransform := strings.Split(url, "'")
	url = urlTransform[1]
	ips, err := pegarIp.PegarIp(url)

	if err != nil {

		resWriter.WriteHeader(403)
		resWriter.Write([]byte(err.Error()))

	} else {

		data, errMarshal := json.Marshal(ips)

		if errMarshal != nil {
			resWriter.WriteHeader(400)
			io.WriteString(resWriter, "Impossível converte para json")
		}

		resWriter.WriteHeader(200)
		resWriter.Header().Set("ContentType", "application/json")
		resWriter.Write(data)
	}
}
