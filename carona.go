package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
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
	URL := "https://maps.google.com.br/maps/api/distancematrix/json?origins=" + url.QueryEscape(origem) + "&destinations=" + url.QueryEscape(destino) + "&key=" + API
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, URL, nil)

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	ch <- body
}
func GetJsonSemgo(origem string, destino string) []byte {
	URL := "https://maps.google.com.br/maps/api/distancematrix/json?origins=" + url.QueryEscape(origem) + "&destinations=" + url.QueryEscape(destino) + "&key=" + API
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, URL, nil)

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
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
	var teste_uber [61]Carona
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
	teste_uber[13] = Carona{"aguas claras,brasilia - DF", "Universidade de Brasilia"}
	teste_uber[14] = Carona{"paranoa, DF", "Hospital de Base do Distrito Federal"}
	teste_uber[15] = Carona{"Floresta Nacional de Brasilia", "aguas claras,brasilia - DF"}
	teste_uber[16] = Carona{"JK Shopping", "aguas claras,brasilia - DFl"}
	teste_uber[17] = Carona{"aguas claras,brasilia - DF", "valparaiso, GO"}
	teste_uber[18] = Carona{"valparaiso, GO", "Hospital de Base do Distrito Federal"}
	teste_uber[19] = Carona{"Universidade de Brasilia", "Hospital de Base do Distrito Federal"}
	teste_uber[20] = Carona{"Ginásio Nilson Nelson", "Hospital de Base do Distrito Federal"}
	teste_uber[21] = Carona{"Universidade de Brasilia", "Ginásio Nilson Nelson"}
	teste_uber[22] = Carona{"Memorial JK", "Ginásio Nilson Nelson"}
	teste_uber[23] = Carona{"nucleo bandeirante, DF", "Hospital de Base do Distrito Federal"}
	teste_uber[24] = Carona{"nucleo bandeirante, DF", "Ginásio Nilson Nelson"}
	teste_uber[25] = Carona{"Nicolândia", "Hospital de Base do Distrito Federal"}
	teste_uber[26] = Carona{"Universidade de Brasilia", "Nicolândia"}
	teste_uber[27] = Carona{"Memorial JK", "Nicolândia"}
	teste_uber[28] = Carona{"Nicolândia, DF", "Hospital de Base do Distrito Federal"}
	teste_uber[29] = Carona{"nucleo bandeirante, DF", "Nicolândia"}
	teste_uber[30] = Carona{"Zoológico de Brasília", "asa sul"}
	teste_uber[31] = Carona{"asa sul", "Zoológico de Brasília"}
	teste_uber[32] = Carona{"paranoa, DF", "Zoológico de Brasília"}
	teste_uber[33] = Carona{"planaltina, DF", "Zoológico de Brasília"}
	teste_uber[34] = Carona{"Zoológico de Brasília", "valparaiso, GO"}
	teste_uber[35] = Carona{"Zoológico de Brasília", "Coco Bambu Lago Sul"}
	teste_uber[36] = Carona{"Zoológico de Brasília", "Hospital de Base do Distrito Federal"}
	teste_uber[37] = Carona{"JK Shopping", "Zoológico de Brasília"}
	teste_uber[38] = Carona{"Praça do Relogio de Taguatinga", "Zoológico de Brasília"}
	teste_uber[39] = Carona{"taguatinga", "Zoológico de Brasília"}
	teste_uber[40] = Carona{"Zoológico de Brasília", "Floresta Nacional de Brasilia"}
	teste_uber[41] = Carona{"Universidade de Brasilia", "Parque Deck Sul"}
	teste_uber[42] = Carona{"Parque Deck Sul", "Coco Bambu Lago Sul"}
	teste_uber[43] = Carona{"aguas claras,brasilia - DF", "Parque Deck Sul"}
	teste_uber[44] = Carona{"paranoa, DF", "Parque Deck Sul"}
	teste_uber[45] = Carona{"Floresta Nacional de Brasilia", "Parque Deck Sul"}
	teste_uber[46] = Carona{"JK Shopping", "Parque Deck Sul"}
	teste_uber[47] = Carona{"Parque Deck Sul", "valparaiso, GO"}
	teste_uber[48] = Carona{"Base Aérea de Brasília", "Hospital de Base do Distrito Federal"}
	teste_uber[49] = Carona{"Universidade de Brasilia", "Base Aérea de Brasília"}
	teste_uber[50] = Carona{"Ginásio Nilson Nelson", "Base Aérea de Brasília"}
	teste_uber[51] = Carona{"Base Aérea de Brasília", "Nicolândia"}
	teste_uber[52] = Carona{"Memorial JK", "Base Aérea de Brasília"}
	teste_uber[53] = Carona{"nucleo bandeirante, DF", "Base Aérea de Brasília"}
	teste_uber[54] = Carona{"Base Aérea de Brasília", "nucleo bandeirante, DF"}
	teste_uber[55] = Carona{"Nicolândia", "Hospital Brasília"}
	teste_uber[56] = Carona{"Universidade de Brasilia", "Hospital Brasília"}
	teste_uber[57] = Carona{"Memorial JK", "Hospital Brasília"}
	teste_uber[58] = Carona{"Hospital Brasília", "Hospital de Base do Distrito Federal"}
	teste_uber[59] = Carona{"nucleo bandeirante, DF", "Hospital Brasília"}
	teste_uber[60] = Carona{"Asa Norte,Brasilia - DF", "itapoâ,Brasilia - DF"}

	var result1 Route
	var result3 Route
	var result2 Route
	LimiteHour := 100000000
	var melhor Carona
	ch2 := make(chan []byte)
	ch1 := make(chan []byte)
	ch3 := make(chan []byte)

	go GetJson(carona.Origem, carona.Destino, ch2)
	route2 := <-ch2
	/*route2 := GetJsonSemgo(carona.Origem, carona.Destino)*/
	//o tempo em media utilizando as goroutines é 2x mais rapido do fazendo sem elas
	//fiz uma outra função para demonstrar na sala de aula.
	err = json.Unmarshal(route2, &result2)
	if err != nil {
		return
	}
	for i := 0; i < 61; i++ {
		go GetJson(teste_uber[i].Origem, carona.Origem, ch1)
		go GetJson(carona.Destino, teste_uber[i].Destino, ch3)
		route1 := <-ch1
		/*route1 := GetJsonSemgo(teste_uber[i].Origem, carona.Origem)*/
		err = json.Unmarshal(route1, &result1)
		if err != nil {
			return
		}
		route3 := <-ch3
		/*route3 := GetJsonSemgo(carona.Destino, teste_uber[i].Destino)*/
		err = json.Unmarshal(route3, &result3)
		if err != nil {
			return
		}
		var percurso int = (result1.Rows[0].Elements[0].Duration.Value +
			result2.Rows[0].Elements[0].Duration.Value +
			result3.Rows[0].Elements[0].Duration.Value)
		if percurso <= LimiteHour {
			melhor = Carona{teste_uber[i].Origem, carona.Destino}
			LimiteHour = percurso
		}

	}
	err = json.NewEncoder(writer).Encode(melhor)
	if err != nil {
		return
	}
	writingSync.Lock()
	programIsRunning = false
	writingSync.Unlock()
}
