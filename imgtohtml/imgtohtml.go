package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {

	file, _ := os.Open("./img.jpg")
	img, _ := jpeg.Decode(file)
	file.Close()

	m := resize.Resize(160, 0, img, resize.Lanczos3)

	fileresized := "./img_resized.jpg"
	out, err := os.Create(fileresized)

	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	imgfile, err := os.Open("./img_resized.jpg")

	if err != nil {
		fmt.Println("img.jpg file not found!")
		os.Exit(1)
	}

	defer imgfile.Close()

	imgCfg, _, err := image.DecodeConfig(imgfile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	width := imgCfg.Width
	height := imgCfg.Height

	imgfile.Seek(0, 0)

	img, _, err = image.Decode(imgfile)

	fmt.Print(`<html>
		<head>
		<title>Seasons greetings!</title>		
		<style type="text/css">
		body {background-color: black;}

		.table{
			width:960px;
			border: 0px;
		}
		td{
			width: 3px;
			height: 3px;
			border: 0px;
		}
		div {
			height: 500px;
			-webkit-align-content: center;
			align-content: center;
		}
		</style>
		</head>
		<body backgroundcolour>
		<div style="height:80px;">&nbsp;</div>
		<div align="center">
		<img src="/find/step0001/greetings.gif"><br/>
		<table class="imgtable" border="0" cellpadding="0" cellspacing="0" style="margin-left: auto; margin-right: auto; font-size:0px;">
			   `)

	message := "/find/2.7182818284590452353602874713527/"
	msgcnt := 0
	for y := 0; y < height; y++ {
		fmt.Print(`
			<tr>
			`)
		for x := 0; x < width; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r := r1 / 256
			g := g1 / 256
			b := b1 / 256
			if y > 50 && msgcnt < len(message) {
				fmt.Printf(`<td bgcolor="#%2x%2x%2x">%s</td>
					`, r, g, b, string(message[msgcnt]))
				msgcnt++
			} else {
				fmt.Printf(`<td bgcolor="#%2x%2x%2x">&nbsp;</td>
					`, r, g, b)
			}
		}
		fmt.Print(`
			</tr>
			`)
	}
	fmt.Print(`</table><br/>
		<img src="/find/step0001/selection.gif"><br/>		
		</div>
		</body>
</html>`)

}
