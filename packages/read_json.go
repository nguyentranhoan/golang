package Media

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DestDir string `json:"dest_dir"`
	Key string `json:"key"`
	Salt string `json:"salt"`
	Resize string `json:"resize"`
	Gravity string `json:"gravity"`
	Enlarge string `json:"enlarge"`
	Extension string `json:"extension"`
	Nginx_src string `json:"nginx_src"`
	Img_proxy_server string `json:"img_proxy_server"`
}

func LoadConfiguration(file string) Config {
	var cf Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&cf)
	return cf
}

var config Config = LoadConfiguration("./config.json")

func GetDestDir()  string{
	return config.DestDir
}

func GetKey()  string{
	return config.Key
}

func GetSalt()  string{
	return config.Salt
}

func GetResize()  string{
	return config.Resize
}

func GetGravity()  string{
	return config.Gravity
}

func GetEnlarge()  string{
	return config.Enlarge
}

func GetExtension()  string{
	return config.Extension
}

func GetNginxSrc()  string{
	return config.Nginx_src
}

func GetImgProxyServer()  string{
	return config.Img_proxy_server
}

func main() {
	fmt.Println(GetDestDir(), GetEnlarge(), GetExtension(), GetGravity())
}