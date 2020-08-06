package Media

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"os"
	"runtime"
)

// thumgb image
func thumb(folder string, fileName string) bool {
	// use all CPU cores for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	// load images and make 100x100 thumbnails of them
	img, err := imaging.Open(folder + fileName)
	if err != nil {
		panic(err)
	}
	var thumbnails image.Image
	thumbnails = imaging.Thumbnail(img, 100, 100, imaging.CatmullRom)

	// create a new blank image
	dst := imaging.New(100, 100, color.NRGBA{0, 0, 0, 0})

	// paste thumbnails into the new image side by side
	dst = imaging.Paste(dst, thumbnails, image.Pt(0, 0))

	// save the combined image to file
	var folderThumb string = "./thumb/"
	if _, err := os.Stat(folderThumb); os.IsNotExist(err) {
		os.MkdirAll(folderThumb, os.ModePerm)
	}
	fmt.Println("xong")
	err = imaging.Save(dst, "thumb/"+fileName)
	if err != nil {
		return false
	}
	return true
}
