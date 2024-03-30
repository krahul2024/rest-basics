package utils

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
)

func GetParams(req *http.Request) *RequestParams {
	params := &RequestParams{
		Method:        req.Method,
		URL:           req.URL.String(),
		Host:          req.Host,
		Protocol:      req.Proto,
		Headers:       req.Header,
		ContentLength: int(req.ContentLength),
		RemoteAddress: req.RemoteAddr,
		Form:          req.Form,
		PostForm:      req.Form,
		MultipartForm: req.MultipartForm,
	}
	return params
}

type RequestParams struct {
	Method        string          `json:"method"`
	URL           string          `json:"url"`
	Host          string          `json:"host"`
	Protocol      string          `json:"protocol"`
	Headers       http.Header     `json:"headers"`
	ContentLength int             `json:"contentLength"`
	RemoteAddress string          `json:"remoteAddress"`
	Form          url.Values      `json:"form"`
	PostForm      url.Values      `json:"postForm"`
	MultipartForm *multipart.Form `json:"multipartForm"`
}

func PrettyJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error formatting the JSON:", err)
		return
	}
	fmt.Printf("%s\n", jsonData)
}
