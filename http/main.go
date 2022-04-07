//https://tutorialedge.net/golang/go-file-upload-tutorial/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("error retriving the file from form file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded file %v\n", handler.Filename)
	fmt.Printf("file size %v\n", handler.Size)
	fmt.Printf("MIME header %v\n", handler.Header)

	// tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	// if err != nil {
	//     fmt.Println(err)
	// }
	// defer tempFile.Close()

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

}
func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}
func main() {
	fmt.Println("go file upload tutorial")
	setupRoutes()
}
