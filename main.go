package main

import (
	"encoding/json"
	"fmt"
	"goCms/models"
	"goCms/services"
	"net/http"
)

func main() {

	service := services.Service{}
	service.GetConfig()
	// mysql := clients.MysqlClient{}
	// mysql.Connect()

	fmt.Println("Welcome to use the simple golang CMS system.")
	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		respondent := models.Respondent{ID: 1, Name: "hello world."}
		w.Write([]byte("hello"))
		data, err := json.Marshal(respondent)
		if err == nil {
			w.Write(data)
		}
	})
	fmt.Printf("starting server on port:%d...", 9999)
	http.ListenAndServe("127.0.0.1:9999", nil)

}
