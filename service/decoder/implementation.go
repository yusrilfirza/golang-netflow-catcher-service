package decoder

import (
	"fmt"
	decoder "netflow-catcher-service/domain/decoder"
	helper "netflow-catcher-service/package/helper/contract"
	"netflow-catcher-service/package/helper/logger"
	"os"
	"strconv"

	"github.com/fln/nf9packet"
)

type implementation struct {
	logger helper.Logger
}

type templateCache map[string]*nf9packet.TemplateRecord

func (i *implementation) RestructureData(template *nf9packet.TemplateRecord, records []nf9packet.FlowDataRecord) {
	fmt.Printf("|")
	for _, f := range template.Fields {
		fmt.Printf(" %s |", f.Name())
	}
	fmt.Printf("\n")

	for _, r := range records {
		fmt.Printf("|")
		for i := range r.Values {
			colWidth := len(template.Fields[i].Name())
			fmt.Printf(" %"+strconv.Itoa(colWidth)+"s |", template.Fields[i].DataToString(r.Values[i]))
		}
		fmt.Printf("\n")
	}
}

func (i *implementation) PacketDump(addr string, data []byte, cache decoder.TemplateCache) {
	p, err := nf9packet.Decode(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	templateList := p.TemplateRecords()
	flowSets := p.DataFlowSets()

	for _, t := range templateList {
		templateKey := fmt.Sprintf("%s|%b|%v", addr, p.SourceId, t.TemplateId)
		cache[templateKey] = t
	}

	for _, set := range flowSets {
		templateKey := fmt.Sprintf("%s|%b|%v", addr, p.SourceId, set.Id)
		template, ok := cache[templateKey]
		if !ok {
			// We do not have template for this Data FlowSet yet
			continue
		}

		records := template.DecodeFlowSet(&set)
		if records == nil {
			// Error in decoding Data FlowSet
			continue
		}
		i.RestructureData(template, records)
	}
}

func NewDecoder() decoder.Service {
	return &implementation{
		logger: logger.NewLogger(),
	}
}
