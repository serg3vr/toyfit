package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"toyfit/src/models"
	"toyfit/src/repository"

	keys "toyfit/src/lib"

	_ "github.com/lib/pq"
)

type MealsHandler struct {
	*repository.MealsRepository
}

func (h *MealsHandler) Create(w http.ResponseWriter, r *http.Request) {
	loggedUserId := r.Context().Value(keys.LoggedUserId).(int)

	var mealRequest models.MealRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&mealRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.MealsRepository.Create(loggedUserId, mealRequest)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not create the meal", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}
