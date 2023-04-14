package pegarIp

import (
	"net"
)

// PegarIp busca o ip relacionado a urlInformada
func PegarIp(url string) ([]string, error) {
	netips, err := net.LookupIP(url)

	if noIp := []string{}; err != nil {
		return noIp, err
	}

	ips := []string{}

	for _, ip := range netips {
		ips = append([]string{}, ip.String())
	}

	return ips, err
}
