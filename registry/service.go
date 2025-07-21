package registry

import (
	decoder "netflow-catcher-service/domain/decoder"
	DecoderService "netflow-catcher-service/service/decoder"
)

type Service struct {
	Decoder decoder.Service
}

func ServiceRegistry() *Service {
	services := &Service{
		Decoder: DecoderService.NewDecoder(),
	}

	return services
}
