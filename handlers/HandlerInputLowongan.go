package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "karierku.com/backend/model"
)

func HandlerInputLowongan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	db, err := model.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	payload := struct {
		ImageUrl string `json:"image_url"`
		Company string `json:"company"`
		Position string `json:"description"`
		Requirement string `json:"requirement"`
		Link string `json:"link"`
	}{}

	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	statement := `
	INSERT INTO lowongan (image_url, company, position, description, requirement, link)
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err = db.Exec(statement, payload.ImageUrl, payload.Company, payload.Position, payload.Requirement, payload.Link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}