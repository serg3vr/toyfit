package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	keys "toyfit/src/lib"
	"toyfit/src/models"
	"toyfit/src/repository"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

type FoodsHandler struct {
	*repository.FoodsRepository
	*repository.MealFoodsRepository
}

func (h *FoodsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.FoodsRepository.GetAll()

	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not get the foods", http.StatusInternalServerError)
		return
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

func (h *FoodsHandler) Update(w http.ResponseWriter, r *http.Request) {
	loggerUserId := r.Context().Value(keys.LoggedUserId).(int64)

	str := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(str, 10, 64)

	_, err := h.FoodsRepository.GetById(id)

	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not found the food", http.StatusBadRequest)
		return
	}

	var food models.FoodRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&food); err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not decode the food", http.StatusBadRequest)
		return
	}

	food.Id = &id

	id, err = h.FoodsRepository.Update(loggerUserId, food)

	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not update the food", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func (h *FoodsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// loggerUserId := r.Context().Value(keys.LoggedUserId).(int64)

	str := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(str, 10, 64)

	_, err := h.FoodsRepository.GetById(id)

	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not found the food", http.StatusBadRequest)
		return
	}

	// Obtener meal foods que contengan esta food si existe mandar error 400

	mealFoods, err := h.MealFoodsRepository.FindByFood(id)

	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not found the food", http.StatusBadRequest)
		return
	}

	if len(mealFoods) > 0 {
		http.Error(w, "Food has depending meal foods", http.StatusBadRequest)
		return
	}

	id, err = h.FoodsRepository.Delete(id)

	if err != nil {
		fmt.Printf("Error %v\n", err)
		http.Error(w, "Could not delete the food", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}
