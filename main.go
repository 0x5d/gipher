package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gosuri/uiprogress"
)

var delay *int

var (
	magicPNG = []byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a}
	magicJPG = []byte{0xff, 0xd8, 0xff}
)

func main() {
	if len(os.Args) < 3 {
		showUsage()
		return
	}
	delay = flag.Int("d", 25, "The delay between images.")
	flag.Parse()
	args := flag.Args()
	paths := args[:len(args)-1]
	outPath := args[len(args)-1]
	var filenames []string
	for _, arg := range paths {
		fmt.Printf("Scanning %s\n", arg)
		var err error
		filenames, err = getImageFilenames(arg)
		if err != nil {
			panic(err)
		}
	}
	err := generateGIF(filenames, outPath)
	if err != nil {
		panic(err)
	}
}

func getImageFilenames(dirName string) ([]string, error) {
	files, err := ioutil.ReadDir(dirName)
	var filenames []string
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		isFileSupported, err := fileSupported(filepath.Join(dirName, file.Name()))
		if err != nil {
			return nil, err
		}
		if !file.IsDir() && isFileSupported {
			filenames = append(filenames, filepath.Join(dirName, file.Name()))
		}

	}
	return filenames, nil
}

func generateGIF(filenames []string, outPath string) error {
	fmt.Printf("Generating GIF in %s\n", outPath)
	uiprogress.Start()
	bar := uiprogress.AddBar(len(filenames)).AppendCompleted()
	anim := gif.GIF{LoopCount: len(filenames)}
	for _, filename := range filenames {
		reader, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer reader.Close()
		img, _, err := image.Decode(reader)
		if err != nil {
			return err
		}
		bounds := img.Bounds()
		drawer := draw.FloydSteinberg
		palettedImg := image.NewPaletted(bounds, palette.Plan9)
		drawer.Draw(palettedImg, bounds, img, image.ZP)
		anim.Image = append(anim.Image, palettedImg)
		anim.Delay = append(anim.Delay, *delay)
		bar.Incr()
	}
	file, err := os.Create(outPath)
	defer file.Close()
	if err != nil {
		return err
	}
	encodeErr := gif.EncodeAll(file, &anim)
	if encodeErr != nil {
		return encodeErr
	}
	return nil
}

func fileSupported(filename string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	fileReader := bufio.NewReader(file)
	headerBytes, err := fileReader.Peek(8)
	if err != nil {
		return false, err
	}
	switch {
	case bytes.Equal(headerBytes, magicPNG):
		return true, nil
	case bytes.Equal(headerBytes[:len(magicJPG)], magicJPG):
		return true, nil
	}
	return false, nil
}

func showUsage() {
	usage := "Usage: gipher [OPTIONS] [in...] out\n\n" +
		"A small GIF generator made with Go.\n\n" +
		"Options:\n\n" +
		"  -d=25\t\tThe delay between frames (in 100s of a second)."
	fmt.Println(usage)
}
