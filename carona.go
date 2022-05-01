package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Carona struct {
	Origem  string
	Destino string
}

const API = "AIzaSyDlMhRPBkMyJMQVNcPq1RgSHVJLrtOa7wk"

type Route struct {
	Rows   []Row  "json:rows"
	Status string "json:status"
}
type Row struct {
	Elements []Element "json:elements"
}
type Element struct {
	Distance Distance "json:distance"
	Duration Duration "json:duration"
}
type Distance struct {
	Text  string "json:text"
	Value int    "json:value"
}
type Duration struct {
	Text  string "json:text"
	Value int    "json:value"
}

func GetJson(origem string, destino string, ch chan []byte) {
	url := "https://maps.google.com.br/maps/api/distancematrix/json?origins=" + url.QueryEscape(origem) + "&destinations=" + url.QueryEscape(destino) + "&key=AIzaSyDlMhRPBkMyJMQVNcPq1RgSHVJLrtOa7wk"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	ch <- body
}

func serveCarona(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writingSync.Lock()
	programIsRunning = true
	writingSync.Unlock()
	enableCors(&writer)
	var carona Carona
	err := json.NewDecoder(request.Body).Decode(&carona)
	if err != nil {
		fmt.Println(err.Error())
		var responseData Carona
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}
	var teste_uber [14]Carona
	teste_uber[0] = Carona{"sobradinho, DF", "asa sul"}
	teste_uber[1] = Carona{"asa sul", "Universidade de Brasilia"}
	teste_uber[2] = Carona{"paranoa, DF", "nucleo bandeirante, DF"}
	teste_uber[3] = Carona{"planaltina, DF", "asa sul"}
	teste_uber[4] = Carona{"lago norte", "valparaiso, GO"}
	teste_uber[5] = Carona{"Memorial JK", "Coco Bambu Lago Sul"}
	teste_uber[6] = Carona{"Floresta Nacional de Brasilia", "Hospital de Base do Distrito Federal"}
	teste_uber[7] = Carona{"JK Shopping", "asa sul"}
	teste_uber[8] = Carona{"Praça do Relogio de Taguatinga", "Universidade de Brasilia"}
	teste_uber[9] = Carona{"paranoa, DF", "JK Shopping"}
	teste_uber[10] = Carona{"planaltina, DF", "Floresta Nacional de Brasilia"}
	teste_uber[11] = Carona{"Universidade de Brasilia", "valparaiso, GO"}
	teste_uber[12] = Carona{"Memorial JK", "Coco Bambu Lago Sul"}
	teste_uber[13] = Carona{"Floresta Nacional de Brasilia", "Hospital de Base do Distrito Federal"}

	var result1 Route
	var result3 Route
	var result2 Route
	var LimiteHour int = 1000000000

	var melhor Carona
	ch2 := make(chan []byte)
	ch1 := make(chan []byte)
	ch3 := make(chan []byte)

	go GetJson(carona.Origem, carona.Destino, ch2)
	route2 := <-ch2
	json.Unmarshal(route2, &result2)
	for i := 0; i < 14; i++ {
		go GetJson(teste_uber[i].Origem, carona.Origem, ch1)
		go GetJson(carona.Destino, teste_uber[i].Destino, ch3)
		route1 := <-ch1
		json.Unmarshal(route1, &result1)
		route3 := <-ch3
		json.Unmarshal(route3, &result3)

		var percurso int = (result1.Rows[0].Elements[0].Duration.Value +
			result2.Rows[0].Elements[0].Duration.Value +
			result3.Rows[0].Elements[0].Duration.Value)
		if percurso <= LimiteHour {
			melhor = Carona{teste_uber[i].Origem, carona.Destino}
			LimiteHour = percurso
		}

	}
	json.NewEncoder(writer).Encode(melhor)
	writingSync.Lock()
	programIsRunning = false
	writingSync.Unlock()
}
