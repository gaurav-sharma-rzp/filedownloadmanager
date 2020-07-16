package main

import (
	"fmt"
	"net/http"
	"log"
	downloadserver "github.com/test/filedownloadmanager/internal/downloader"
	"github.com/test/filedownloadmanager/rpc/downloader"
)

type timeHandler struct {
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Before")
	//h.ServeHTTP(w, r) // call original
	w.Write([]byte("Application up and running!"))
	log.Println("After")
}
//
//func logger(h http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Println("Before")
//		//h.ServeHTTP(w, r) // call original
//		w.Write([]byte("Hi "))
//		log.Println("After")
//	})
//}

func main() {
	server := &downloadserver.Server{} // implements Downloader interface
	twirpHandler := downloader.NewDownloaderServer(server, nil)

	mux := http.NewServeMux()
	//twirpHandler.PathPrefix()
	fmt.Println(twirpHandler.PathPrefix())
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	//mux.Handle(twirpHandler.PathPrefix()+"/hat/", twirpHandler)
	//mux.Handle("/hat", twirpHandler)
	//th:= timeHandler{}
	//mux.Handle("/health", &th)

	http.ListenAndServe(":8080", mux)
	//http.ListenAndServe(":8080", twirpHandler)
}
