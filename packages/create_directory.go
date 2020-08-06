package Media


import (
	"os"
	"path/filepath"
)
var size []SizeImage = updateSize()
func CreateDir()  {
	//size := updateSize()
	y := GetYear()
	m := GetMonth()
	d := GetDay()

	for _, sizedt:=range size{
		//fmt.Println(sizedt.height)
		width := sizedt.width
		height := sizedt.height
		destDir := GetDestDir()
		path := destDir + width + "x" + height + "/" + y + "/" + m + "/" + d
		newPath := filepath.Join(".", path)
		os.MkdirAll(newPath, os.ModePerm)
	}
}
