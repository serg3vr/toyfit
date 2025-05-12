package models

import (
// "time"
)

type BasicFoodCreate struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Kcal        *float64 `json:"kcal"`
	Carbs       *float64 `json:"carbs"`
	Proteins    *float64 `json:"proteins"`
	Fats        *float64 `json:"fats"`
}

type FoodResponse struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Kcal        float64  `json:"kcal"`
	Carbs       *float64 `json:"carbs"`
	Proteins    *float64 `json:"proteins"`
	Fats        *float64 `json:"fats"`
	Sodium      *float64 `json:"sodium"`
}
