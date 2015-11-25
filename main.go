package main

import (
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
	"strings"
)

var delay = 1

func main() {
	if len(os.Args) < 3 {
		showUsage()
		return
	}
	args := os.Args[1 : len(os.Args)-1]
	outPath := os.Args[len(os.Args)-1]
	var fileNames []string
	for _, arg := range args {
		var err error
		fileNames, err = getImageFileNamesRec(arg)
		if err != nil {
			panic(err)
		}
	}
	err := generateGIF(fileNames, outPath)
	if err != nil {
		panic(err)
	}
}

func getImageFileNamesRec(dirName string) ([]string, error) {
	files, err := ioutil.ReadDir(dirName)
	var fileNames []string
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			childrenFiles, err := getImageFileNamesRec(filepath.Join(dirName, file.Name()))
			if err != nil {
				return nil, err
			}
			fileNames = append(fileNames, childrenFiles...)
		} else if formatSupported(file.Name()) {
			fileNames = append(fileNames, filepath.Join(dirName, file.Name()))
		}

	}
	return fileNames, nil
}

func generateGIF(fileNames []string, outPath string) error {
	anim := gif.GIF{LoopCount: len(fileNames)}
	for _, fileName := range fileNames {
		reader, err := os.Open(fileName)
		if err != nil {
			panic(err)
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
		anim.Delay = append(anim.Delay, delay)
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

func formatSupported(fileName string) bool {
	return strings.HasSuffix(fileName, ".png") ||
		strings.HasSuffix(fileName, ".jpg") ||
		strings.HasSuffix(fileName, ".jpeg")
}

func showUsage() {
	fmt.Println("Usage: gipher <folder paths> <output path>")
}
