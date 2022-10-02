package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/itinao/go-sample/app/models"
	"github.com/itinao/go-sample/config"
)

var templates = template.Must(template.ParseFiles("app/views/index.html"))

func viewIndexHandler(w http.ResponseWriter, r *http.Request) {
	df, _ := models.GetAllTodo(100)
	err := templates.ExecuteTemplate(w, "index.html", df.Todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func APIError(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONError{Error: errMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

var apiValidPath = regexp.MustCompile("^/api/todo/$")

func apiMakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := apiValidPath.FindStringSubmatch(r.URL.Path)
		if len(m) == 0 {
			APIError(w, "Not found", http.StatusNotFound)
			return
		}
		fn(w, r)
	}
}

func apiTodoHandler(w http.ResponseWriter, r *http.Request) {
	strLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strLimit)
	if strLimit == "" || err != nil {
		APIError(w, "No limit param", http.StatusBadRequest)
		return
	}

	if limit < 0 || limit > 1000 {
		APIError(w, "Upper and lower limit error [0 ~ 1000]", http.StatusBadRequest)
		return
	}

	df, _ := models.GetAllTodo(limit)

	js, err := json.Marshal(df)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func StartWebServer() error {
	log.Printf("start web server localhost:%d", config.Config.Port)
	http.HandleFunc("/", viewIndexHandler)
	http.HandleFunc("/api/todo/", apiMakeHandler(apiTodoHandler))
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
