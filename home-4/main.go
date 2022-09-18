package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

// directoryContent -.
func directoryContent(w http.ResponseWriter, r *http.Request, f http.File) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprintf(w, "<pre>\n")
	_, err = fmt.Fprintf(w, `<table style="border: 1px solid black;">`)
	_, err = fmt.Fprintf(w, `<tr><th style="border: 1px solid black;">name</th><th style="border: 1px solid black;">ext</th><th style="border: 1px solid black;">size</th></tr>`)
	//ext := r.URL.Query().Get("ext")
	ext := r.FormValue("ext")

	for {
		dirs, err := f.Readdir(100)
		if err != nil || len(dirs) == 0 {
			break
		}

		for _, d := range dirs {
			name := d.Name()
			if d.IsDir() {
				name += "/"
			}
			ne := strings.Split(name, ".")

			if len(ne) > 1 && ne[0] != "" && ne[1] != "" {
				if ext != "" {
					if ne[1] == ext {
						_, err = fmt.Fprintf(w, `<tr><td style="border: 1px solid black;">%s</td><td style="border: 1px solid black;">%s</td><td style="border: 1px solid black;">%d</td></tr>`, ne[0], ne[1], d.Size())
					}

				} else {
					_, err = fmt.Fprintf(w, `<tr><td  style="border: 1px solid black;">%s</td><td style="border: 1px solid black;">%s</td><td style="border: 1px solid black;">%d</td></tr>`, ne[0], ne[1], d.Size())
				}

			}
		}
	}
	_, err = fmt.Fprintf(w, "</table>")
	_, err = fmt.Fprintf(w, "</pre>\n")
	if err != nil {
		fmt.Println("error writing files info: ", err)
	}
}

// serveFile -.
func serveFile(w http.ResponseWriter, r *http.Request, fs http.FileSystem, name string) error {
	f, err := fs.Open(name)
	if err != nil {
		fmt.Println("error opening file: ", name, err)
		return err
	}
	defer func(f http.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("error closing file: ", name, err)
			//return err
		}
	}(f)

	d, err1 := f.Stat()
	if err1 != nil {
		fmt.Println(err1)
		return err
	}

	if d.IsDir() {
		directoryContent(w, r, f)
		return err
	}
	return nil
}

// fileHandler -.
type fileHandler struct {
	http.FileSystem
}

// FileServer -.
func FileServer(fs http.FileSystem) http.Handler {
	return &fileHandler{fs}
}

// ServeHTTP -.
func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ServeHTTP Method: ", r.Method)
	err := serveFile(w, r, f, path.Clean(r.URL.Path))
	if err != nil {
		fmt.Println("error serving files: ", err)
	}
}

type UploadHandler struct {
	UploadDir string
	HostAddr  string
}

func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("Unable to read file: ", err)
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file) // считали файл
	if err != nil {
		log.Println("Unable to read file content: ", err)
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}

	filePath := h.UploadDir + "/" + header.Filename // выяснили куда писать
	log.Println("filePath: ", filePath, h.HostAddr)
	err = ioutil.WriteFile(filePath, data, 0777) // записали файл
	if err != nil {
		log.Println("Unable to save file", err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "File %s has been successfully uploaded", header.Filename)
	fileLink := h.HostAddr + "/" + header.Filename
	fmt.Fprintln(w, fileLink)
}

func main() {

	mux := http.NewServeMux()
	dir := http.Dir("./files")
	fmt.Println("server started at port 8082")

	uploadHandler := &UploadHandler{
		UploadDir: "files",
		HostAddr:  ":63342",
	}

	mux.Handle("/files", uploadHandler)
	mux.Handle("/", FileServer(dir))

	srv := &http.Server{
		Addr:         ":8082",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	srv.ListenAndServe()
}
