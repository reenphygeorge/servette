package ports

import (
	"errors"
	"net"
	"strconv"
)

func GetFreePort(startPort, endPort int) (string, error) {
	for port := startPort; port <= endPort; port++ {
		address := "localhost:" + strconv.Itoa(port)

		ln, err := net.Listen("tcp", address)
		if err == nil {
			defer ln.Close()
			return strconv.Itoa(port), nil
		}
	}

	return "", errors.New("no free port available in the specified range")
}
