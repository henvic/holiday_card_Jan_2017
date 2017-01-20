package imageprocessing

import "github.com/jung-kurt/gofpdf"

// Card for processing
type Card struct {
	Title    string
	ImageSrc string
	Dest     string
}

const template = "frame.png"

// New creates a new card
// frame from https://www.prestophoto.com/designer/template/800
func New(card Card) (err error) {
	pdf := gofpdf.New("L", "mm", "A6", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, card.Title)

	pdf.Image(template, 0, 0, 0, 115, false, "", 0, "")

	if card.ImageSrc != "" {
		pdf.Image(card.ImageSrc, 10, 10, 30, 0, false, "", 0, "")
	}

	err = pdf.OutputFileAndClose(card.Dest)
	return err
}
