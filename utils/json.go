package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(res http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Json parsing error!\n%v", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(code)
	res.Write(data)
}

func RespondWithError(res http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Internal Service Error!, %v", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	RespondWithJson(res, code, errResponse{Error: msg})
}
