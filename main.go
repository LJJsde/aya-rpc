package main

import "rpc/service"

func main() {
	service.ServerInit("kcp", "localhost")
}
