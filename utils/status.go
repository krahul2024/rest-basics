package utils

import "net/http"

func ReportStatus(res http.ResponseWriter, req *http.Request) {
	RespondWithJson(
		res, 200,
		status{Message: "Server is running without any errors!"})
}

type status struct {
	Message string `json:"message"`
}
