package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var requestJSON Request
	err = json.Unmarshal(body, &requestJSON)
	if err != nil || requestJSON.Message == "" {
		responseError := Response{
			Status:  http.StatusBadRequest,
			Message: "Некорректное JSON-сообщение",
		}
		sendJSONResponse(w, responseError)
		return
	}

	fmt.Printf("Received POST request with message: %s\n", requestJSON.Message)

	response := Response{
		Status:  http.StatusOK,
		Message: "Данные успешно приняты",
	}

	sendJSONResponse(w, response)
}

func sendJSONResponse(w http.ResponseWriter, response Response) {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	w.Write(responseJSON)
}

func main() {
	http.HandleFunc("/post", handlePostRequest)

	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
