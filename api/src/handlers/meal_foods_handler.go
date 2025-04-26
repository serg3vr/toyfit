package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	// "toyfit/config"
	"toyfit/src/models"
	"toyfit/src/repository"

	keys "toyfit/src/lib"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

type MealFoodsHandler struct {
	*repository.MealFoodsRepository
	*repository.FoodsRepository
}

func (h *MealFoodsHandler) GetDaily(w http.ResponseWriter, r *http.Request) {
	loggedUserId := r.Context().Value(keys.LoggedUserId).(int64)
	timeZone := r.Context().Value(keys.TimeZone).(string)
	date := r.URL.Query().Get("date")

	data, err := h.MealFoodsRepository.GetDaily(loggedUserId, timeZone, date)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not get daily meal foods", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *MealFoodsHandler) Create(w http.ResponseWriter, r *http.Request) {
	loggedUserId := r.Context().Value(keys.LoggedUserId).(int64)

	var mealFood models.MealFoodsRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&mealFood); err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not create the meal food", http.StatusNotFound)
		return
	}

	if mealFood.FoodName != nil && *mealFood.FoodName != "" {
		var food models.BasicFoodCreate
		food.Name = *mealFood.FoodName
		food.Description = *mealFood.FoodDescription
		food.Kcal = *mealFood.FoodKcal

		createdFoodId, err := h.FoodsRepository.Create(loggedUserId, food)

		if err != nil {
			fmt.Printf("Error %v\n", err)
			http.Error(w, "Could not create the food", http.StatusNotFound)
			return
		}
		mealFood.FoodId = &createdFoodId
	}

	id, err := h.MealFoodsRepository.Create(loggedUserId, mealFood)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not create the meal food", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func (h *MealFoodsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	loggedUserId := r.Context().Value(keys.LoggedUserId).(int64)

	str := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(str, 10, 64)

	_, err := h.MealFoodsRepository.Delete(loggedUserId, id)

	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not delete the meal food", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}
