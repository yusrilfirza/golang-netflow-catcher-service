package main

import "netflow-catcher-service/registry"

func main() {
	var port = "9995"
	registry.Registry(port)
}
