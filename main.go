package main

import "rpc/service"

func main() {
	service.ClientInit("kcp", "localhost")
}
