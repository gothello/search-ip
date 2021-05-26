package main

import (
	"fmt"
	"net"
	"net/http"
)

func last(s string, b byte) int {
	i := len(s)

	for i--; i > 0; i-- {
		if s[i] == b {
			break
		}
	}

	return i
}

func splitHostZone(s string) (host, zone string) {
	if i := last(s, '%'); i > 0 {
		host, zone = s[:i], s[i+1:]
		return
	}

	host = s
	return
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	addr := r.Context().Value(http.LocalAddrContextKey)
	serverAddr, serverPort, err := net.SplitHostPort(
		fmt.Sprintf("%v", addr),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	serverIP, serverZone := splitHostZone(serverAddr)
	fmt.Println("server ip:", serverIP, "port:", serverPort)
	if serverZone != "" {
		fmt.Println("zone:", serverZone)
	}

	fmt.Println(net.JoinHostPort(serverIP, serverPort))

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=")

	clientAddr, clientPort, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	clientIP, clientZone := splitHostZone(clientAddr)
	fmt.Println("client ip:", clientIP, "port:", clientPort)
	fmt.Println("zone:", clientZone)
	fmt.Println(net.JoinHostPort(clientIP, clientPort))
}
func main() {

	r := http.NewServeMux()

	r.HandleFunc("/", handleMain)

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}

}
