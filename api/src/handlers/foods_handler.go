package handlers

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"toyfit/config"
	keys "toyfit/src/lib"
	"toyfit/src/models"
	"toyfit/src/repository"
)

type FoodsHandler struct {
	*repository.FoodsRepository
}

func (h *FoodsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT
			id,
			name,
			description,
			kcal,
			carbs,
			proteins,
			fats,
			sodium
		FROM foods
		ORDER BY name DESC
	`
	var data []models.FoodResponse
	rows, err := config.DB.Query(query)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not get the foods", http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var model models.FoodResponse
		err := rows.Scan(
			&model.Id,
			&model.Name,
			&model.Description,
			&model.Kcal,
			&model.Carbs,
			&model.Proteins,
			&model.Fats,
			&model.Sodium,
		)
		if err != nil {
			fmt.Printf("Error %v\n", err)
			http.Error(w, "Could not get the foods", http.StatusInternalServerError)
			return
		}
		data = append(data, model)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *FoodsHandler) Create(w http.ResponseWriter, r *http.Request) {
	loggerUserId := r.Context().Value(keys.LoggedUserId).(int64)

	var food models.FoodRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&food); err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not create the food", http.StatusBadRequest)
		return
	}

	id, err := h.FoodsRepository.Create(loggerUserId, food)

	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not create the food", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}
