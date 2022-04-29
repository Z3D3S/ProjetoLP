package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/julienschmidt/sse"
	"html/template"
	"net/http"
	"time"
)

type TimeDataInput struct {
	Name string
	Time string
}

func getTime(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data TimeDataInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data.Name)
	fmt.Println(data.Time)
}

func streamTime(timer *sse.Streamer) {
	fmt.Println("Streaming time started")
	for serviceIsRunning {
		timer.SendString("", "time", time.Now().String())
		time.Sleep(1 * time.Millisecond)
	}
}

func serveHomepage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	enableCors(&writer)
	writingSync.Lock()
	programIsRunning = true
	writingSync.Unlock()

	var homepage TimeDataInput
	homepage.Time = time.Now().String()
	tmpl := template.Must(template.ParseFiles("html/homepage.html"))
	_ = tmpl.Execute(writer, homepage)
	writingSync.Lock()
	programIsRunning = false
	writingSync.Unlock()
}
