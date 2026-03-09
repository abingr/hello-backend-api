package utils

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Contest-Type", "applicaion/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}
