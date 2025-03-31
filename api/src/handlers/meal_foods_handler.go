package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"toyfit/config"
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
	date := r.URL.Query().Get("date")

	query := `
		select
			mf.id,
			mf.meal_type_id, 
			f.name 
		from meal_foods mf
		join foods f on f.id = mf.food_id
		where 
			date(date) = $1
		order by mf.meal_type_id, mf.created_at
	`
	// var data []models.DailyMealFoods
	data := make([]models.DailyMealFoods, 0)
	rows, err := config.DB.Query(query, date)
	if err != nil {
		// return data, err
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for rows.Next() {
		var model models.DailyMealFoods
		err := rows.Scan(
			&model.Id,
			&model.MealTypeId,
			&model.Name,
		)
		if err != nil {
			// return data, err
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data = append(data, model)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
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
