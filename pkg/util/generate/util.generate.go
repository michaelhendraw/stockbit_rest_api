package generate

import (
	"log"
	"net/url"
	"strings"
)

// CleanGRPCAddress use to append url scheme for gRPC address
// some url in config does not include url scheme but some include, so this function will return the url address with scheme
func CleanGRPCAddress(addr string) (host string) {
	if addr == "" {
		log.Println("func cleanGRPCAddress Empty addr")
		return
	}

	if !strings.HasPrefix(addr, "grpc://") {
		addr = "grpc://" + addr
	}
	u, err := url.Parse(addr)
	if err != nil {
		log.Println("func cleanGRPCAddress unexpected hostname:", addr, err)
		return
	}

	if u.Scheme != "grpc" && u.Scheme != "" {
		log.Println("func cleanGRPCAddress only grpc is supported, got", u.Scheme)
		return
	}
	host = u.Host
	return
}
