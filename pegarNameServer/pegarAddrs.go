package pegarNameServer

import (
	"net"
)

func PegarNameServer(url string) ([]string, error) {
	servidores, err := net.LookupNS(url)
	namesServers := []string{}

	if err != nil {
		return namesServers, err
	}

	for _, servidor := range servidores {
		namesServers = append(namesServers, servidor.Host)
	}

	return namesServers, nil
}
