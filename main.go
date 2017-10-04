package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/page.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func CityHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'cities':'Toronto, Montreal, Ottawa'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}


func handlerCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	t, err := template.ParseFiles("csv-data/csco-daily.full.csv")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}


func main() {
	port := os.Getenv("WEB_SERVER_PORT")
	if port == "" {
		port = ":8000"
	}
	fmt.Printf("Listening on port %s!\n", port)

	http.HandleFunc("/", handler)
	http.HandleFunc("/cities.json", CityHandler)

	http.HandleFunc("/get-csv-data", handlerCSV)

	fs := http.FileServer(http.Dir("charts"))
	http.Handle("/charts/", http.StripPrefix("/charts/", fs))

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
