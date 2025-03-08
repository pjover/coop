package cfg

import (
	"github.com/pjover/coop/domain/ports"
	"github.com/pjover/coop/domain/services/pdf"
	"log"
	"sync"
)

type DiContainer struct {
	configService ports.ConfigService
	pdfService    ports.PdfService
}

// Singleton instance of DiContainer
var (
	instance *DiContainer
	once     sync.Once
)

func DI() *DiContainer {
	once.Do(func() {
		log.Print("Initializing dependency injection container...")
		var configService = NewConfigService()
		instance = &DiContainer{
			configService: configService,
			pdfService:    pdf.NewPdfService(configService),
		}
	})
	return instance
}

func (di DiContainer) ConfigService() ports.ConfigService {
	return di.configService
}

func (di DiContainer) PdfService() ports.PdfService {
	return di.pdfService
}
