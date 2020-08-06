package Media

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

// process upload
func upload(w http.ResponseWriter, r *http.Request) {
	//Saved in the server memory with maxMemory size
	r.ParseMultipartForm(32 << 20)

	//Handle file upload
	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		response(w,"Lỗi upload",err)
		return
	}
	defer file.Close()

	//Make Path to save file
	var folder string = "resources/sources/" + GetYear() + "/" + GetMonth() + "/" + GetDay() + "/"
	fmt.Print(folder)
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, os.ModePerm)
	}


	fileName := get_img_name()

	//Save file to storage
	f, err := os.OpenFile(folder + fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		response(w,"Lỗi Server",err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	//Response thumb image
	if thumb(folder, fileName) {
		http.ServeFile(w, r, path.Join("thumb", fileName))
		os.Remove("./thumb/" +fileName)


		// Process Resize & Save in XXX
		Signature(get_img_name())

		return
	}
	response(w,"Lỗi trong quá trình thumbnails",nil)
	return
}

//handle request
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/upload", upload).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func get_img_name() string{
	return file_name + ".jpg"
}
