package main

import (
	"log"
	"net/http"

	"github.com/youtube/google-oauth/config"
	"github.com/youtube/google-oauth/controller"
)

func main() {
	// load configs
	//	config.LoadEnv()
	config.LoadConfig()

	// create a router
	fs := http.FileServer(http.Dir("temp"))
	http.Handle("/", fs)

	// define routes
	http.HandleFunc("/google_login", controller.GoogleLogin)
	http.HandleFunc("/google_callback", controller.GoogleCallback)
	http.HandleFunc("/fb_login", controller.FbLogin)
	http.HandleFunc("/fb_callback", controller.FbCallback)
	http.HandleFunc("/git_login", controller.GitLogin)

	//run server
	log.Println("started server on :: http://localhost:8494/")
	if oops := http.ListenAndServe(":8494", nil); oops != nil {
		log.Fatal(oops)
	}
}
