package main

import (
	"bytes"
	"encoding/base64"
	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"
)

func requestImg(url string) image.Image {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(res.Body)
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(bytes.NewReader(buffer.Bytes()))
	if err != nil {
		panic(err)
	}

	return img
}

func resizeImg(img image.Image, width uint, height uint) image.Image {

	return resize.Resize(width, height, img, resize.Lanczos3)

}

func convertWebp(img image.Image) []byte {
	var buf bytes.Buffer
	err := webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: 90})
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func Main(args map[string]interface{}) map[string]interface{} {
	url, _ := args["url"].(string)
	width, _ := strconv.Atoi(args["width"].(string))
	height, _ := strconv.Atoi(args["height"].(string))

	img := requestImg(url)

	resized := resizeImg(img, uint(width), uint(height))

	wp := convertWebp(resized)

	encodedImg := base64.StdEncoding.EncodeToString(wp)

	msg := make(map[string]interface{})
	msg["body"] = encodedImg
	msg["headers"] = map[string]interface{}{
		"Content-Type": "image/webp",
	}
	return msg
}
