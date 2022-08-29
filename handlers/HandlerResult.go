package handlers

import (
	// "encoding/json"
	"net/http"
)

// type result struct {
// 	RecommendationRole string
// 	RecomendationHardSkills []string
// }

func HandlerResult(w http.ResponseWriter, r *http.Request){
// 	if r.Method != "POST" {
// 		http.Error(w, "", http.StatusBadRequest)
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	payload := struct {
// 		Id int `json:"id"`
// 	}{}

// 	if err := decoder.Decode(&payload); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	var data result

// 	switch payload.Id{
// 	case 3:
// 		data = result{
// 			RecommendationRole: "Dummy Recommendation Role 1",
// 			RecomendationHardSkills: []string {"Dummy Hardskill 1", "Dummy Hardskill 2"},
// 		}

// 	case 4:
// 		data = result{
// 			RecommendationRole: "Dummy Recommendation Role 2",
// 			RecomendationHardSkills: []string {"Dummy Hardskill 1", "Dummy Hardskill 2"},
// 		}
// 	}

// 	w.Header().Set("Content-Type", "application/json")
	
// 	err := json.NewEncoder(w).Encode(data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
}