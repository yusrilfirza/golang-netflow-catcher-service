package registry

import decoder "netflow-catcher-service/domain/decoder"

func Registry(port string) {
	infrastructure := InfrastructureRegistry()
	service := ServiceRegistry()

	addr, err := infrastructure.ResolveUDPAddr(port)
	if err != nil {
		panic(err)
	}

	con, err := infrastructure.ListenUDP(addr)
	if err != nil {
		panic(err)
	}

	data := make([]byte, 8960)
	cache := make(decoder.TemplateCache)

	for {
		length, remote, err := con.ReadFrom(data)
		if err != nil {
			panic(err)
		}

		service.Decoder.PacketDump(remote.String(), data[:length], cache)
	}
}
