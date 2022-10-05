package handlers

import (
	"encoding/json"
	"net/http"
)

type result struct {
	RecommendationMajor string
	RecomendationHardSkills []string
}

type question struct {
	Question string
	Options []option
	IsLast bool
}

type option struct {
	Id int
	Text string
}

func HandlerQuiz(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	payload := struct {
		Id int `json:"id"`
		IsLast bool `json:"isLast"`
	}{}
	
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	if payload.IsLast {
		response := generateResult(payload.Id)
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else { 
		response := generateQuestion(payload.Id)			
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func generateQuestion(answerId int) (question){
	switch answerId {
	case 0:
		question := question{
			Question: "Apakah kamu suka melakukan coding ?",
			Options: []option{
				{Id: 1, Text: "Ya"},
				{Id: 2, Text: "Tidak"},
			},
			IsLast: false,
		}
		return question
	
	case 1:
		question := question{
			Question: "Yang mana yang lebih kamu sukai ?",
			Options: []option{
				{Id: 3, Text: "Visual"},
				{Id: 4, Text: "Data"},
			},
			IsLast: true,
		}
		return question

	case 2:
		question := question{
			Question: "Yang mana yang lebih kamu sukai ?",
			Options: []option{
				{Id: 5, Text: "Visual"},
				{Id: 6, Text: "Data"},
			},
			IsLast: true,
		}
		return question

	default:
		return question{}
	}
}

func generateResult(answerId int) (result) {
	switch answerId{
	case 3:
		res := result{
			RecommendationMajor: "Front-End Engineer",
			RecomendationHardSkills: []string {"HTML", "CSS", "JavaScript"},
		}
		return res
	case 4:
		res := result{
			RecommendationMajor: "Back-End Engineer",
			RecomendationHardSkills: []string {"DBMS", "API Engineering"},
		}
		return res
	case 5:
		res := result{
			RecommendationMajor: "UI/UX Designer",
			RecomendationHardSkills: []string {"UI Design", "UX Design"},
		}
		return res
	case 6:
		res := result{
			RecommendationMajor: "Data Analyst",
			RecomendationHardSkills: []string {"Data Visualization", "Data Analysis"},
		}
		return res
	default:
		return result{}
	}
}
