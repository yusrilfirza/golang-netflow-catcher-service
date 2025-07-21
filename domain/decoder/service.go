package decoder

import "github.com/fln/nf9packet"

type Service interface {
	RestructureData(template *nf9packet.TemplateRecord, records []nf9packet.FlowDataRecord)

	PacketDump(addr string, data []byte, cache TemplateCache)
}
