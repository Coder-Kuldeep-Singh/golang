package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"html/template"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
)

var root = flag.String("root", ".", "file system path")

func main() {
	http.HandleFunc("/blue/", blueHandler)
	http.HandleFunc("/red/", redHandler)
	http.Handle("/", http.FileServer(http.Dir(*root)))
	log.Println("Listening on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
func blueHandler(response http.ResponseWriter, req *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	writeImage(response, &img)
}
func redHandler(response http.ResponseWriter, req *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{255, 0, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	WriteImageWithTemplate(response, &img)
}

var ImageTemplate string = `
<!Doctype html>
<html lang="en"><head></head>
<body><img src="data:image/jpeg;base64,{{.Image}}"></body></html>`

//Writeimagewithtemplate encodes an image 'img' in jpeg
//format and writes it into ResponseWriter using a template

func WriteImageWithTemplate(response http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Fatalln("unable to encode image.")
	}
	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	if tmpl, err := template.New("image").Parse(ImageTemplate); err != nil {
		log.Fatalln("unable to parse image template.")
	} else {
		data := map[string]interface{}{"image": str}
		if err = tmpl.Execute(response, data); err != nil {
			log.Println("unable to execute template.")
		}
	}
}

//writeImage encodes an image 'img' in jpeg format and
//writes it into ResponseWriter
func writeImage(response http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	response.Header().Set("Content-Type", "image/jpeg")
	response.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := response.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
