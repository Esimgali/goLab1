package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request структура для представления JSON-запроса
type Request struct {
	Message string `json:"message"`
}

// Response структура для представления JSON-ответа
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	// Чтение данных из тела POST-запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Преобразование JSON-запроса в структуру
	var requestJSON Request
	err = json.Unmarshal(body, &requestJSON)
	if err != nil || requestJSON.Message == "" {
		// Ошибка в преобразовании JSON
		responseError := Response{
			Status:  http.StatusBadRequest,
			Message: "Некорректное JSON-сообщение",
		}
		sendJSONResponse(w, responseError)
		return
	}

	// Вывод данных в консоль
	fmt.Printf("Received POST request with message: %s\n", requestJSON.Message)

	// Создание JSON-ответа
	response := Response{
		Status:  http.StatusOK,
		Message: "Данные успешно приняты",
	}

	// Отправка JSON-ответа клиенту
	sendJSONResponse(w, response)
}

// Функция для отправки JSON-ответа клиенту
func sendJSONResponse(w http.ResponseWriter, response Response) {
	// Преобразование структуры в JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	// Отправка JSON-ответа клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	w.Write(responseJSON)
}

func main() {
	// Установка обработчика для POST-запросов на эндпоинт "/post"
	http.HandleFunc("/post", handlePostRequest)

	// Запуск веб-сервера на порту 8080
	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
