package main

import (
	"encoding/json"
	"log"
	"net/http"
        "fmt"
        "os"
        "html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {  
        t, err := template.ParseFiles("templates/page.html")
        if err != nil {
             panic(err)
        }
        t.Execute(w, nil)
}

func CityHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'cities':'San Francisco, Amsterdam, Berlin, New York','Tokyo'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

funcxx main() {
        port := os.Getenv("WEB_SERVER_PORT")
        if port=="" {
            port = ":8000"
        }
        fmt.Printf("Listening on port %s!\n",port)
	http.HandleFunc("/cities.json", CityHandler)
        http.HandleFunc("/", handler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
