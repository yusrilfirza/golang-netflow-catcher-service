package registry

import (
	"flag"
	"net"
	helperContract "netflow-catcher-service/package/helper/contract"
)

var dumpJSON bool

type implementation struct{}

func (i *implementation) ResolveUDPAddr(port string) (*net.UDPAddr, error) {
	listenAddr := flag.String("listen", ":"+port, "Address to listen for NetFlow v9 packets.")
	flag.BoolVar(&dumpJSON, "json", false, "Dump packet in JSON instead of plain text.")
	flag.Parse()

	return net.ResolveUDPAddr("udp", *listenAddr)
}

func (i *implementation) ListenUDP(address *net.UDPAddr) (*net.UDPConn, error) {
	return net.ListenUDP("udp", address)
}

func InfrastructureRegistry() helperContract.Intfrastructure {
	return &implementation{}
}
