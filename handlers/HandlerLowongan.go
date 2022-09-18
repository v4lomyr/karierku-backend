package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "karierku.com/backend/model"
)

type lowongan struct {
	Id int
	ImageUrl string
	Position string
	Company string
	Detail detail
}

type detail struct {
	Description string
	Requirement string
	Link string
}

func HandlerLowongan(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	db, err := model.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	rows, err := db.Query(`
	SELECT id, image_url, company, position, description, requirement, link
	FROM lowongan
	`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	defer db.Close()

	var data []lowongan
	for rows.Next() {
		var each = lowongan{}
		var err = rows.Scan(
			&each.Id,
			&each.ImageUrl,
			&each.Company,
			&each.Position,
			&each.Detail.Description,
			&each.Detail.Requirement,
			&each.Detail.Link,
		)

		if err != nil {
			fmt.Println(err.Error())
		}

		data = append(data, each)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	response := struct {
		ListOfVacancies []lowongan
	}{
		ListOfVacancies: data,
	}
	
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandlerRekomendasiLowongan(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "*")
	
	if r.Method != "GET" {
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
		RecommendationRole string
	}{}

	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := db.Query(`
		select lowongan.id, lowongan.image_url, lowongan.company, lowongan.position, lowongan.description, lowongan.requirement, lowongan.link
		from lowongan
		join roles on lowongan.suitable = roles.id
		where roles.role = $1
	`, payload.RecommendationRole)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var data []lowongan

	for rows.Next() {
		var each = lowongan{}
		var err = rows.Scan(
			&each.Id,
			&each.ImageUrl,
			&each.Company,
			&each.Position,
			&each.Detail.Description,
			&each.Detail.Requirement,
			&each.Detail.Link,
		)
		
		if err != nil {
			fmt.Println(err.Error())
		}

		data = append(data, each)
	}

	response := struct{
		ListOfVacancies []lowongan
	}{
		ListOfVacancies: data,
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}