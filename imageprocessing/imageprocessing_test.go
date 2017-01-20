package imageprocessing

import "testing"

func TestNew(t *testing.T) {
	var card = Card{
		Title:    "Hello, World!",
		ImageSrc: "",
		Dest:     "file.pdf",
	}

	if err := New(card); err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}
