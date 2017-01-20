package imageprocessing

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// Card for processing
type Card struct {
	Title string
	Text  string

	ImageSrc string
	Dest     string
}

const frame = "../frame.png"

// New creates a new card
// frame from https://www.prestophoto.com/designer/template/800
func New(card Card) (err error) {
	pdf := gofpdf.New("L", "mm", "A6", "")
	pdf.AddPage()

	pdf.Image(frame, 0, 0, 0, 115, false, "", 0, "")

	if card.ImageSrc != "" {
		pdf.Image(card.ImageSrc, 40, 30, 30, 0, false, "", 0, "")
	}

	pdf.SetFont("Arial", "B", 14)
	pdf.WriteAligned(0, 60, strings.Repeat(" ", 45)+card.Title, "")
	pdf.Ln(20)
	pdf.SetFont("Arial", "I", 9)
	pdf.WriteAligned(0, 35, strings.Repeat(" ", 70)+card.Text, "L")

	err = pdf.OutputFileAndClose(card.Dest)
	return err
}
