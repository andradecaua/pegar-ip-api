package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pegar-ip-ou-addr/model"
	"strings"
)

func IpController(resWriter http.ResponseWriter, req *http.Request) {
	if strings.ToLower(req.Method) != "get" {
		resWriter.Write([]byte(fmt.Sprintf("NO METHOD %s", req.Method)))
	}
	url := req.URL.Query().Get("url")
	urlTransform := strings.Split(url, "'")
	url = urlTransform[1]

	if ips, err := model.PegarIp(url); err != nil {
		resWriter.WriteHeader(403)
		resWriter.Write([]byte(err.Error()))
	} else {

		if ipsJSON, err := json.Marshal(ips); err != nil {
			resWriter.WriteHeader(400)
			io.WriteString(resWriter, "Impossível converte para json")
		} else {
			resWriter.WriteHeader(200)
			resWriter.Header().Set("ContentType", "application/json")
			resWriter.Write(ipsJSON)
		}
	}
}
