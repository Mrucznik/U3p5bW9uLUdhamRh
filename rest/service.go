package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CreateRequest struct {
	Id int32 `json:"id,omitempty"`
	// URL Address.
	Url string `json:"url,omitempty"`
}

type IdRequest struct {
	Id int32 `json:"id,omitempty"`
}

func listUrls(writer http.ResponseWriter, request *http.Request) {
	result, err := service.Get()
	if err != nil {
		http.Error(writer, http.StatusText(400), 400)
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(writer, http.StatusText(500), 500)
	}
	writer.Write(jsonData)
}

func createUrl(writer http.ResponseWriter, request *http.Request) {
	var data CreateRequest
	request.Body = http.MaxBytesReader(writer, request.Body, 1000000) //max 1MB

	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := service.Create(data.Url, data.Id)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	writer.Write([]byte(strconv.Itoa(int(id))))
}

func deleteUrl(writer http.ResponseWriter, request *http.Request) {
	var data IdRequest
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.Delete(data.Id)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	writer.Write([]byte{})
}

func listHistory(writer http.ResponseWriter, request *http.Request) {
	var data IdRequest
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.History(data.Id)
	if err != nil {
		http.Error(writer, http.StatusText(400), 400)
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(writer, http.StatusText(500), 500)
	}
	writer.Write(jsonData)
}
