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
	*repository.MealFoodsRepository
}

func (h *MealsHandler) Create(w http.ResponseWriter, r *http.Request) {
	loggedUserId := r.Context().Value(keys.LoggedUserId).(int64)

	var mealRequest models.MealWithFoodsRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&mealRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.MealsRepository.Create(loggedUserId, mealRequest)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not create the meal", http.StatusInternalServerError)
		return
	}

	var mealFoodsRequest = *mealRequest.MealFoodsRequest

	for _, foodRequest := range mealFoodsRequest {
		foodRequest.MealId = id // Asignar el ID de la comida creada
		_, err := h.MealFoodsRepository.Create(loggedUserId, foodRequest)
		if err != nil {
			fmt.Printf("Error creating meal food %v\n", err)
			http.Error(w, "Could not create the meal food", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}
