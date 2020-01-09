package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func Start(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/chunk", chunkGenerator)
	http.HandleFunc("/pdf", pdfServer)
	http.HandleFunc("/pdf2", pdfServer2)
	http.HandleFunc("/payload", payload)

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./static"))))

	log.Println("Listening on " + port)
	http.ListenAndServe(":"+port, nil)
}

func chunkGenerator(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	w.Header().Set("X-Content-Type-Options", "nosniff")
	for i := 1; i <= 100; i++ {

		fmt.Fprintf(w, "Chunk #%d\n", i)
		flusher.Flush() // Trigger "chunked" encoding and send a chunk...
		time.Sleep(500 * time.Millisecond)
	}
}

func pdfServer(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./static/foo.pdf")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	defer f.Close()

	//Set header
	w.Header().Set("Content-type", "application/pdf")

	//Stream to response
	if _, err := io.Copy(w, f); err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
}

type Payload struct {
	MineType     string `json:"mineType"`
	DocumentData string `json:"documentData"`
}

func pdfServer2(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./static/foo.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pdfData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print(err)
	}

	pdfBase := base64.URLEncoding.EncodeToString(pdfData)
	data := &Payload{
		MineType:     "application/pdf",
		DocumentData: string(pdfBase),
	}

	PDFJson, err := json.Marshal(data)
	if err != nil {
		fmt.Print(err)
	}

	//Set header
	w.Header().Set("Content-type", "application/json")

	_, err = w.Write(PDFJson)
	if err != nil {
		fmt.Print(err)
	}
}

func payload(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 77399)
	w.Header().Add("Content-Length", fmt.Sprintf("%d", len(b)))
	w.Write(b)
}
