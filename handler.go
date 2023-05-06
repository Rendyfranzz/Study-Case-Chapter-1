package main

import (
	"crud/helper"
	"net/http"
)

type Map map[string]interface{}

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"id":    1,
			"title": "Kemeja",
			"stok":  1000,
		},
		{
			"id":           2,
			"nama_product": "Celana",
			"stok":         10000,
		},
		{
			"id":           1,
			"nama_product": "Sepatu",
			"stok":         500,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)

}
