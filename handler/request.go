package handler

import (
	"encoding/json"
	"fixer-api-client/enitity"
	"fmt"
	"html/template"
	"net/http"
)

const (
	Endpoint  string = "http://data.fixer.io/api/"
	Latest    string = "latest/"
	Symbols   string = "symbols/"
	AccessKey string = "?access_key="
)

type RequestHandler struct {
	tmpl *template.Template
}

func NewRequestHandler(tmpl *template.Template) *RequestHandler {
	return &RequestHandler{tmpl: tmpl}
}
func (rh *RequestHandler) Latest(w http.ResponseWriter, r *http.Request) {
	base := r.FormValue("base")
	symbols := r.FormValue("symbols")
	res, err := http.Get(Endpoint + Latest + AccessKey + queryfy("base", base) + queryfy("symbols", symbols))
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer res.Body.Close()
	l := res.ContentLength
	data := make([]byte, l)
	_, err = res.Body.Read(data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	exchangeReq := enitity.ExchangeRequest{}
	err = json.Unmarshal(data, &exchangeReq)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	err = rh.tmpl.ExecuteTemplate(w, "response.layout", exchangeReq)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}
func (rh *RequestHandler) Historical(w http.ResponseWriter, r *http.Request) {
	date := r.FormValue("day")
	base := r.FormValue("base")
	symbols := r.FormValue("symbols")
	fmt.Print(date)
	res, err := http.Get(Endpoint + date + "/" + AccessKey + queryfy("base", base) + queryfy("symbols", symbols))
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer res.Body.Close()
	l := res.ContentLength
	data := make([]byte, l)
	_, err = res.Body.Read(data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	exchangeReq := enitity.ExchangeRequest{}
	err = json.Unmarshal(data, &exchangeReq)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	rh.tmpl.ExecuteTemplate(w, "response.layout", exchangeReq)
}

func (rh *RequestHandler) Home(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.Write([]byte("404 page not found"))
		return
	}
	res, err := http.Get(Endpoint + Symbols + AccessKey)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer res.Body.Close()
	l := res.ContentLength
	data := make([]byte, l)
	res.Body.Read(data)
	symbolsRequest := enitity.SymbolsRequest{}
	err = json.Unmarshal(data, &symbolsRequest)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	err = rh.tmpl.ExecuteTemplate(writer, "request.layout", symbolsRequest)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}
func queryfy(name, value string) string {
	return "?" + name + "=" + value
}
