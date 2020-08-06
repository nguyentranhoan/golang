package Media

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// create an encoded url

func EncodedURL(filename string) string{
	y := GetYear()
	m := GetMonth()
	d := GetDay()
	nginxSrcProxy := GetNginxSrc()
	url := nginxSrcProxy + y + "/" + m + "/" + d + "/" + filename
	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(url))
	return encodedURL
}

func Signature(filename string) {

	CreateDir()

	// check key, salt

	var keyBin, saltBin []byte
	var err error

	// define key & salt

	key := GetKey()
	salt := GetSalt()

	if keyBin, err = hex.DecodeString(key); err != nil {
		log.Fatal("Key expected to be hex-encoded string")
	}

	if saltBin, err = hex.DecodeString(salt); err != nil {
		log.Fatal("Salt expected to be hex-encoded string")
	}
	// define attribute of resized images
	resize := GetResize()
	gravity := GetGravity()
	enlarge := GetEnlarge()
	extension := GetExtension()

	// create an encode url
	encodedURL := EncodedURL(filename)


	// resize source images to create a needed sizes, then save it.
	for _, sizedt := range size {
		//fmt.Println(sizedt.height)
		width, _ := strconv.ParseInt(sizedt.width, 10, 32)
		height, _ := strconv.ParseInt(sizedt.height, 10, 32)
		path := fmt.Sprintf("/%s/%d/%d/%s/%s/%s.%s", resize, width, height, gravity, enlarge, encodedURL, extension)
		mac := hmac.New(sha256.New, keyBin)
		mac.Write(saltBin)
		mac.Write([]byte(path))
		signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

		proxyServer := GetImgProxyServer()
		finalPath := proxyServer + signature + path

		response, e := http.Get(finalPath)
		if e != nil {
			log.Fatal(e)
		}
		defer response.Body.Close()

		//open a file for writing
		// save resized image
		pathFile := "./resources/cache"

		fileDate := "/" + GetYear() + "/" + GetMonth() + "/" + GetDay() + "/"

		resizedImage := "/" + sizedt.width + "x" + sizedt.height
		resizedImageName := filename + ".png"
		resizedImagePath := pathFile + resizedImage + fileDate + resizedImageName

		file, err := os.Create(resizedImagePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Success!")
	}
}