package imageprocessing

import "testing"

func TestNew(t *testing.T) {
	var card = Card{
		Title:    "Hello, World!",
		Text:     "I am Go's Gopher.",
		ImageSrc: "fancygopher.jpg",
		Dest:     "file.pdf",
	}

	if err := New(card); err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}
