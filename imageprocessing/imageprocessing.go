package imageprocessing

import "github.com/jung-kurt/gofpdf"

// Card for processing
type Card struct {
	Title    string
	ImageSrc string
	Dest     string
}

// New creates a new card
func New(card Card) (err error) {
	pdf := gofpdf.New("L", "mm", "A6", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, card.Title)

	var options = gofpdf.ImageOptions{
		ReadDpi: true,
	}

	if card.ImageSrc != "" {
		pdf.ImageOptions(card.ImageSrc, 0, 0, 0, -1, false, options, 1, "")
	}

	err = pdf.OutputFileAndClose(card.Dest)
	return err
}
