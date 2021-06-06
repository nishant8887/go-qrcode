package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/nishant8887/go-qrcode"
)

const commandTitle = `qrgenerate - Generate a QR code
https://github.com/nishant8887/go-qrcode

Flags:`

const commandUsage = `Usage:
1. Generate a QR code image.
	qrgenerate -o hello.png "Hello"
2. Specify an error correction level.
	qrgenerate -e 3 -o hello.jpg "Hello"
`

func main() {
	outFile := flag.String("o", "", "filename for image output (with .png or .jpg extension)")
	errorCorrectionLevel := flag.Int("e", 1, "error correction level L=1, M=2, Q=3, H=4")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, commandTitle)
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, commandUsage)
	}
	flag.Parse()

	if len(flag.Args()) == 0 || *outFile == "" || (*errorCorrectionLevel < 1 || *errorCorrectionLevel > 4) {
		flag.Usage()
		os.Exit(1)
		return
	}

	extension := filepath.Ext(*outFile)
	if extension != ".png" && extension != ".jpg" {
		flag.Usage()
		os.Exit(1)
		return
	}

	content := flag.Args()[0]

	ecl := qrcode.L
	switch *errorCorrectionLevel {
	case 1:
		ecl = qrcode.L
	case 2:
		ecl = qrcode.M
	case 3:
		ecl = qrcode.Q
	case 4:
		ecl = qrcode.H
	}

	code, err := qrcode.New(content, ecl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	imgFile, err := os.Create(*outFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	if extension == ".png" {
		err = png.Encode(imgFile, code.Image())
	} else {
		err = jpeg.Encode(imgFile, code.Image(), nil)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}
