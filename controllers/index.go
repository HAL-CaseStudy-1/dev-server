package controllers

import (
	"encoding/json"
	"net/http"
)

type IndexJSON struct {
	Live bool `json:"live"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(IndexJSON{
		Live: true,
	})

	if err != nil {
		panic(err)
	}
}
