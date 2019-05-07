package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/robertkrimen/otto"
	"html/template"
	"log"
	"net/http"
	"time"
)

var templates = template.Must(template.ParseFiles("src/index.html"))

type State struct {
	count int `datastore:"value,noindex" json:"value"`
}

type IndexPage struct {
	HTML  template.HTML
	State template.JS
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/state", HandleState)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:7070",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Print("listening on port :7070")
	log.Fatal(srv.ListenAndServe())

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderedHTML, renderedState, err := render(&State{count: 0})
	if err != nil {
		panic(err)
	}

	_ = templates.ExecuteTemplate(w, "src/index.html", IndexPage{
		HTML:  template.HTML(renderedHTML),
		State: template.JS(renderedState),
	})

}

func HandleState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"count": 0})
}

func render(state *State) (string, string, error) {
	stateJson, err := json.Marshal(&state)

	if err != nil {
		return "", "", err
	}

	var renderResult otto.Value
	renderResult, err = callRenderJS(string(stateJson))

	var renderedHTML, renderedState otto.Value
	renderedHTML, err = renderResult.Object().Get("html")
	if err != nil {
		return "", "", err
	}
	renderedState, err = renderResult.Object().Get("state")
	if err != nil {
		return "", "", err
	}

	return renderedHTML.String(), renderedState.String(), nil

}

func callRenderJS(stateJSON string) (otto.Value, error) {
	vm := otto.New()

	var v, renderJS otto.Value
	script, err := vm.Compile("src/server.js", nil)
  	if err != nil {
		return v, err
	}
	v, err = vm.Run(script)
	if err != nil {
		return v, err
	}
	v, err = vm.Get("server")
	if err != nil {
		return v, err
	}
	renderJS, err = v.Object().Get("render")
	if err != nil {
		return v, err
	}

	return renderJS.Call(otto.NullValue(), stateJSON)
}
