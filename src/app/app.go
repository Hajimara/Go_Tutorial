package main

import (
	"log"
	"net/http"
)

// http.ResponseWriter 브라우저에 전달할 응답을 수정하기 위한 값, Request 전달받은 요청값
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, GO!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
