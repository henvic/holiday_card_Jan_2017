package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/godinezj/holiday_card_Jan_2017/imageprocessing"
)

var (
	infile  string
	outfile string
	title   string
	message string
)

// initializes cmd line flags
func init() {
	flag.StringVar(&infile, "i", "", "The image filename")
	flag.StringVar(&outfile, "o", "", "The output filename")
	flag.StringVar(&title, "t", "", "The title for the card")
	flag.StringVar(&message, "m", "", "The message for the card")
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}

	if infile == "" {
		fmt.Println("Please specify a file name (-i)")
		return
	}

	if message == "" {
		fmt.Println("Please specify a message, use quotes (-m)")
		return
	}

	// fmt.Printf("%v\n", infile)
	fmt.Printf("%v\n", title)

	if _, e := os.Stat(infile); e != nil {
		fmt.Println("File does not exist")
	}

	var card = imageprocessing.Card{
		Title:    title,
		ImageSrc: fmt.Sprintf("../%s", infile),
		Dest:     outfile,
	}

	if err := imageprocessing.New(card); err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
}
