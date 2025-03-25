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

type MealFoodsHandler struct {
	*repository.MealFoodsRepository
	*repository.FoodsRepository
}

func (h *MealFoodsHandler) Create(w http.ResponseWriter, r *http.Request) {
	loggedUserId := r.Context().Value(keys.LoggedUserId).(int64)

	var mealFoodRequest models.MealFoodsRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&mealFoodRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mealFoodRequest.FoodName != nil && *mealFoodRequest.FoodName != "" {
		var dbFoodCreate models.DBFoodCreate
		dbFoodCreate.Name = *mealFoodRequest.FoodName
		dbFoodCreate.Kcal = 1

		dbFoodCreateId, err := h.FoodsRepository.Create(loggedUserId, dbFoodCreate)

		if err != nil {
			fmt.Printf("Error %v\n", err)
			http.Error(w, "Could not create the food", http.StatusNotFound)
			return
		}
		mealFoodRequest.FoodId = &dbFoodCreateId
	}

	id, err := h.MealFoodsRepository.Create(loggedUserId, mealFoodRequest)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not create the meal food", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}
