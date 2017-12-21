package main

//(0,0), (255,0), (0,175) and (255,175) are the bottom left-, bottom right-, top left- and top right-corners.
import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
)

func main() {
	imgfile, err := os.Open("./img.jpg")

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

	fmt.Print(`10 REM (c) RedSocks Security
20 PRINT "https://redsocks.eu/"
30 GO TO 20
40 REM https://lets.g0.rs/find/step0002/
50 REM ---------------------------------
60 REM Try GO TO there v
70 BRIGHT 1: PAPER 7: INK 2`)

	imgfile.Seek(0, 0)

	// get the image
	img, _, err := image.Decode(imgfile)
	drawing := false
	drawwidth := 0
	linecount := 101
	fmt.Print("\n100 ")
	itemcount := 0
	hasadded := false
	for y := 0; y < height; y++ {
		drawing = false
		for x := 0; x < width; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			if r1+g1+b1 > 0 { //A very simple cutoff
				if !drawing {
					drawing = true
					drawwidth = 0
					fmt.Print(fmt.Sprintf("PLOT %d,%d:", x, 175-y))
					hasadded = true
				} else {
					drawwidth++
				}
			} else {
				if drawing {
					if drawwidth > 0 {
						fmt.Print(fmt.Sprintf("DRAW %d,%d:", drawwidth, 0))
					}
					drawing = false
					hasadded = true
				}
			}
			if y == 130 && x == 0 {
				fmt.Print("\n", linecount, " BRIGHT 0: INK 4")
				linecount++
				itemcount = 0
				fmt.Print("\n", linecount, " ")
				linecount++
				hasadded = false
			} else if hasadded {
				itemcount++
				if itemcount > 15 {
					fmt.Print("\n", linecount, " ")
					linecount++
					itemcount = 0
				}
				hasadded = false
			}

		}
	}
	linecount++
	fmt.Print("\n", linecount, " BRIGHT 0:INK 0:PAPER 7")

}
