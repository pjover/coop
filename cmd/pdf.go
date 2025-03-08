package cmd

import (
	"github.com/pjover/coop/cfg"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	pdfCmd.AddCommand(deliveryNoteCmd)
	rootCmd.AddCommand(pdfCmd)
}

var pdfCmd = &cobra.Command{
	Use:   "pdf",
	Short: "PDF related commands",
}

var deliveryNoteCmd = &cobra.Command{
	Use:   "deliveryNote",
	Short: "Create PDFs for delivery notes and update status to `published`",
	Long:  "Create PDFs for delivery notes (all draft or by ID) and update status to `published`",
	Example: `  pdf dn       Create PDF for all draft delivery notes,
  pdf dn 666   Create PDF for delivery note 666`,
	Annotations: map[string]string{"PDF": "PDF commands"},
	Aliases:     []string{"dn"},
	Args:        cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var pdfService = cfg.DI().PdfService()
		if len(args) > 0 {
			id, err := ParseInteger(args[0], "id")
			if err != nil {
				return err
			}
			msg, err := pdfService.DeliveryNote(id)
			if err != nil {
				log.Fatal(err)
				return err
			}
			log.Print(msg)
		} else {
			msg, err := pdfService.DraftDeliveryNotes()
			if err != nil {
				log.Fatal(err)
				return err
			}
			log.Print(msg)
		}
		return nil
	},
}
