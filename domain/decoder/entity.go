package decoder

import "github.com/fln/nf9packet"

type TemplateCache map[string]*nf9packet.TemplateRecord

type NetflowRecord struct {
	LastSwitched      uint64
	FirstSwitched     uint64
	InPkts            uint64
	InBytes           uint64
	InputSNMP         uint32
	OutputSNMP        uint32
	IPv4SrcAddr       string
	IPv4DstAddr       string
	Protocol          uint8
	SrcTOS            uint8
	L4SrcPort         uint16
	L4DstPort         uint16
	IPv4NextHop       string
	DstMask           uint8
	SrcMask           uint8
	TCPFlags          uint8
	SamplingInterval  uint32
	SamplingAlgorithm uint8
	InDstMAC          string
	InSrcMAC          string
	OutDstMAC         string
	OutSrcMAC         string
	UnknownType225    interface{}
	UnknownType226    interface{}
	UnknownType227    interface{}
	UnknownType228    interface{}
}
