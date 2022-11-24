package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	a := 0.33
	//crgamos las 2 imagenes
	imgPath := "img1.jpg"
	f, err := os.Open(imgPath)
	check(err)

	img, _, err := image.Decode(f)

	imgPath2 := "img2.jpg"
	f2, err2 := os.Open(imgPath2)
	check(err2)

	img2, _, err := image.Decode(f2)

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	start := time.Now()
	// loop though all the x

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			pixel := img.At(x, y)
			pixel2 := img2.At(x, y)

			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			originalColor2 := color.RGBAModel.Convert(pixel2).(color.RGBA)

			// Offset colors a little, adjust it to your taste
			r := float64(originalColor.R)
			g := float64(originalColor.G)
			b := float64(originalColor.B)

			r2 := float64(originalColor2.R)
			g2 := float64(originalColor2.G)
			b2 := float64(originalColor2.B)

			// sumamos las 2 iamgenes para combinarlas
			r3 := uint8((r + r2) * a)
			g3 := uint8((g + g2) * a)
			b3 := uint8((b + b2) * a)

			if r3 > 255 || g3 > 255 || b3 > 255 {
				r3 = 255
				g3 = 255
				b3 = 255
			}

			if r3 < 0 || g3 < 0 || b3 < 0 {
				r3 = 0
				g3 = 0
				b3 = 0
			}

			c := color.RGBA{
				R: r3, G: g3, B: b3, A: originalColor.A,
			}
			wImg.Set(x, y, c)
		}
	}

	elapsed := time.Since(start)
	log.Printf("tiempo de ejcucion: %s", elapsed)

	ext := filepath.Ext(imgPath)
	name := strings.TrimSuffix(filepath.Base(imgPath), ext)
	newImagePath := fmt.Sprintf("%s/%s_ejercicio1sec%s", filepath.Dir(imgPath), name, ext)
	fg, err := os.Create(newImagePath)
	defer fg.Close()
	check(err)
	err = jpeg.Encode(fg, wImg, nil)
	check(err)
}
