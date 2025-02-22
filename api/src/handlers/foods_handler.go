package handlers

import (
	"encoding/json"
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
			kcal,
			carbs,
			proteins,
			fats
		FROM foods
		ORDER BY name DESC
	`
	var data []models.FoodResponse
	rows, err := config.DB.Query(query)
	if err != nil {
		// return data, err
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for rows.Next() {
		var food models.FoodResponse
		err := rows.Scan(
			&food.Id,
			&food.Name,
			&food.Kcal,
			&food.Carbs,
			&food.Proteins,
			&food.Fats,
		)
		if err != nil {
			// return data, err
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data = append(data, food)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rows)
}

// func (h *FoodsHandler) Create(w http.ResponseWriter, r *http.Request) {
// 	loggedUserId := r.Context().Value(keys.LoggedUserId).(int)

// 	var reqTransaction models.RequestedTransaction
// 	decoder := json.NewDecoder(r.Body)

// 	if err := decoder.Decode(&reqTransaction); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	id, err := h.TransactionsRepository.Create(loggedUserId, reqTransaction)
// 	if err != nil {
// 		fmt.Printf("Error %v", err)
// 		http.Error(w, "Could create the transaction", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(id)
// }
