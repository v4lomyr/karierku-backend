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
			Question: "dummy question 1",
			Options: []option{
				{Id: 1, Text: "dummy option 1"},
				{Id: 2, Text: "dummy option 2"},
			},
			IsLast: false,
		}
		return question
	
	case 1:
		question := question{
			Question: "dummy question 2",
			Options: []option{
				{Id: 3, Text: "dummy option 1"},
				{Id: 4, Text: "dummy option 2"},
			},
			IsLast: false,
		}
		return question

	case 2:
		question := question{
			Question: "dummy question 2",
			Options: []option{
				{Id: 3, Text: "dummy option 1"},
				{Id: 4, Text: "dummy option 2"},
			},
			IsLast: false,
		}
		return question

	case 3, 4:
		question := question{
			Question: "dummy question 3",
			Options: []option{
				{Id: 5, Text: "dummy option 1"},
				{Id: 6, Text: "dummy option 2"},
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
	case 5:
		res := result{
			RecommendationMajor: "Front-End Engineer",
			RecomendationHardSkills: []string {"Design", "HTML", "CSS", "JavaScript"},
		}
		return res
	case 6:
		res := result{
			RecommendationMajor: "Back-End Engineer",
			RecomendationHardSkills: []string {"DBMS", "API Engineering"},
		}
		return res
	default:
		return result{}
	}
}
