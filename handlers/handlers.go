package handlers

import (
	"1_helloapi/utils"
	"net/http"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"message": "Welcome to my Go Backend API",
	}

	utils.JSONResponse(w, response)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"message": "Hello from Go backend!",
	}

	utils.JSONResponse(w, response)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"status": "ok",
	}

	utils.JSONResponse(w, response)
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"server_time": time.Now().Format(time.RFC3339),
	}

	utils.JSONResponse(w, response)
}
