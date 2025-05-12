package handlers

import (
	"encoding/json"
	"fmt"
	// "fmt"
	"net/http"
	"toyfit/config"
	"toyfit/src/models"

	// "toyfit/src/repository"

	// keys "toyfit/src/lib"

	_ "github.com/lib/pq"
)

type FoodsHandler struct {
	//*repository.TransactionsRepository
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
