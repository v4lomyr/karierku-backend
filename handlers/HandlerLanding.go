package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"karierku.com/backend/model"
)

func HandlerLanding(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	db, err := model.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query(`
	SELECT id, image_url, company, position, description, requirement, link
	FROM lowongan
	ORDER BY id DESC
	FETCH FIRST 4 ROWS ONLY
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