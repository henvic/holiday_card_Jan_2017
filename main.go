package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/SDGophers/holiday_card_Jan_2017/cardtemplate"
	"github.com/SDGophers/holiday_card_Jan_2017/imageprocessing"
)

var (
	infile    string
	outfile   string
	title     string
	message   string
	usersFile string
)

// initializes cmd line flags
func init() {
	flag.StringVar(&infile, "i", "", "The image filename")
	flag.StringVar(&outfile, "o", "", "The output filename")
	flag.StringVar(&title, "t", "", "The title for the card")
	flag.StringVar(&message, "m", "", "The message for the card")
	flag.StringVar(&usersFile, "u", "", "The user's file")
}

func genCard(t cardtemplate.Template, file string, user string) {
	if _, e := os.Stat(infile); e != nil {
		fmt.Println("File does not exist")
	}

	var text, errt = t.GetSubject(user)

	if errt != nil {
		panic(errt)
	}

	var card = imageprocessing.Card{
		Title:    title,
		Text:     text,
		ImageSrc: infile,
		Dest:     file,
	}

	if err := imageprocessing.New(card); err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}

	if infile == "" {
		println("Please specify a file name (-i)")
		os.Exit(1)
		return
	}

	if message == "" {
		println("Please specify a message, use quotes (-m)")
		os.Exit(1)
	}

	if usersFile != "" {
		f, err := os.Open(usersFile)
		if err != nil {
			println("error opening file ", err)
			os.Exit(1)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		cntr := 0

		var cardTemplate = cardtemplate.Template{
			Subject: message,
		}

		for scanner.Scan() {
			name := fmt.Sprintf("card%d.pdf", cntr)
			genCard(cardTemplate, name, scanner.Text())
			cntr++
		}
	}
}
