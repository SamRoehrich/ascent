package main

import (
	"flag"
	"log"
	"net/http"

	"engine"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "hi")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	if r.URL.Path != "/" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "home.html")
}

func main() {

	flag.Parse()

	hub := engine.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		engine.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Listen and serve: ", err)
	}

	log.Println("No errors, server running")
}
