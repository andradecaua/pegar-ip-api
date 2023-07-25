package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pegar-ip-ou-addr/model"
	"strings"
)

type response struct {
	Status  int
	Message string
}

func NameServerController(resWriter http.ResponseWriter, req *http.Request) {
	if strings.ToLower(req.Method) != "get" {
		res := response{
			Status:  500,
			Message: "Metodo inacessível para esta rota",
		}
		resJSON, err := json.Marshal(res)
		if err != nil {
			resWriter.Header().Set("Content-Type", "application/json")
			resWriter.WriteHeader(500)
			resWriter.Write([]byte(""))
		}
		resWriter.Write(resJSON)
	}

	url := req.URL.Query().Get("url")
	urlTransform := strings.Split(url, "'")
	url = urlTransform[1]

	nameServers, err := model.PegarNameServer(url)

	if err != nil {

		var res response = response{
			Status:  403,
			Message: "Ouve um erro pegar o name server",
		}
		resJSON, errConvertJSON := json.Marshal(res)

		if errConvertJSON != nil {
			resWriter.WriteHeader(500)
			resWriter.Write([]byte("Ouve um erro ao converter para JSON sua resposta"))
		}

		resWriter.WriteHeader(403)
		resWriter.Write(resJSON)

	} else {
		dataJSON, errConverter := json.Marshal(nameServers)

		if errConverter != nil {
			resWriter.WriteHeader(400)
			io.WriteString(resWriter, "Não fo possível converter para JSON")
		} else {
			resWriter.WriteHeader(200)
			resWriter.Header().Set("ContentType", "application/json")
			resWriter.Write(dataJSON)
		}

	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Print("teste")
		}
	}()
}
