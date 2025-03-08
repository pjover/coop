package ports

type PdfService interface {
	DeliveryNote(id int) (string, error)
	DraftDeliveryNotes() (string, error)
}
