package decoder

import (
	"encoding/json"
	"fmt"
	decoder "netflow-catcher-service/domain/decoder"
	helper "netflow-catcher-service/package/helper/contract"
	"netflow-catcher-service/package/helper/logger"
	"os"

	"github.com/fln/nf9packet"
)

type implementation struct {
	logger helper.Logger
}

func (i *implementation) RestructureData(template *nf9packet.TemplateRecord, records []nf9packet.FlowDataRecord) {

	for _, r := range records {
		item := map[string]interface{}{}
		for i := range r.Values {
			item[template.Fields[i].Name()] = template.Fields[i].DataToString(r.Values[i])
		}

		jsonData, _ := json.MarshalIndent(item, "", " ")
		fmt.Print(string(jsonData))
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
