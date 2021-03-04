package main

import (
	"github.com/nfnt/resize"
	"image/png"
	"image/jpeg"
	"os"
	"fmt"
	"net/http"
)
func main() {
	Imagetype()
	// ResizeJPG()
}
//Imagetype ...
func  Imagetype() {

	// maximize CPU usage for maximum performance
	// runtime.GOMAXPROCS(runtime.NumCPU())

	// open the uploaded file
	f := "imgs/golang.jpeg"
	filename := "work.png"
fmt.Println("stage0")
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("error with the image")
	}
fmt.Println("stage1")
	buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err = file.Read(buff)

	if err != nil {
					fmt.Println(err)
					os.Exit(1)
	}

	filetype := http.DetectContentType(buff)

	fmt.Println(filetype)

	switch filetype {
	case "image/jpeg", "image/jpg":
					ResizeJPG(f, filename)

	case "image/png":
					ResizePng(f, filename)

	default:
					fmt.Println("Wrong file format")
	}

}
//ResizePng ...
func ResizePng(f, filepath string){
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("error opening png")
	}
	// decode png into image.Image
	
	fmt.Println("stage 3")
	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("error decoding png")
	}

	// resize to width 60 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(500, 0, img, resize.Lanczos3)
	// return out, nil
	out, err := os.Create("./work.png")
	if err != nil {
		 	fmt.Println("error decoding png")
	}
	defer out.Close()

	// write new image to file
	f1 := png.Encode(out, m)
		if f1 != nil {
		fmt.Println("error encoding png")
	}
	png.Encode(out, m)
}
//ResizeJPG ...
func ResizeJPG(f, filepath string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("error opening jpeg")
	}
	// decode png into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("error decoding jpeg")
	}

	// resize to width 60 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(500, 0, img, resize.Lanczos3)
	out, err := os.Create("./work.png")
	if err != nil {
			fmt.Println("error decoding jpeg")
	}
	defer out.Close()

	// write new image to file
	f1 := jpeg.Encode(out, m, nil)
		if f1 != nil {
			fmt.Println("error encoding jpeg")
	}
	jpeg.Encode(out, m, nil)
}