func resizeImg(img image.Image, width uint, height uint) image.Image {
	return resize.Resize(width, height, img, resize.Lanczos3)
}

func Main(args map[string]interface{}) map[string]interface{} {
    imgData := args["image"].(string) // Assuming you're passing the image data as a string

    // Decode the base64 encoded image data
    decodedImg, err := base64.StdEncoding.DecodeString(imgData)
    if err != nil {
        // Handle error
        return nil
    }

    // Decode the image
    img, _, err := image.Decode(bytes.NewReader(decodedImg))
    if err != nil {
        // Handle error
        return nil
    }

    resized := resizeImg(img, uint(width), uint(height))

    encodedImg := base64.StdEncoding.EncodeToString(wp)

    msg := make(map[string]interface{})
    msg["body"] = encodedImg
    msg["headers"] = map[string]interface{}{
        "Content-Type": "image/webp",
    }
    return msg
}
    