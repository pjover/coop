package pdf

import (
	"fmt"
	"github.com/pjover/coop/domain/ports"
	"log"
)

type pdfService struct {
	configService ports.ConfigService
}

func NewPdfService(configService ports.ConfigService) ports.PdfService {
	return pdfService{
		configService: configService,
	}
}

func (p pdfService) DeliveryNote(id int) (string, error) {
	log.Printf("Building PDF for delivery note with id %d ...", id)
	filename := "DeliveryNote-001.pdf"
	msg := fmt.Sprintf("Succesfully built delivery note: %s", filename)
	return msg, nil
}

func (p pdfService) DraftDeliveryNotes() (string, error) {
	log.Print("Building PDFs for draft delivery notes ...")
	filenames := [...]string{"DeliveryNote-001.pdf", "DeliveryNote-002.pdf"}
	msg := fmt.Sprintf("Succesfully built delivery notes: %s", filenames)
	return msg, nil
}
