package models

import (
	"time"
)

type MealFoodsRequest struct {
	MealTypeId   int64     `json:"meal_type_id"`
	Date         time.Time `json:"date"`
	FoodId       *int64    `json:"food_id"`
	CustomFoodId *int64    `json:"custom_food_id"`
	Qty          *int16    `json:"qty"`
	// FoodName        *string   `json:"food_name"`
	// FoodDescription *string   `json:"food_description"`
	// FoodKcal        *float64  `json:"food_kcal"`

	FoodName        *string  `json:"food_name"`
	FoodDescription *string  `json:"food_description"`
	FoodKcal        *float64 `json:"food_kcal"`
	FoodCarbs       *float64 `json:"food_carbs"`
	FoodProteins    *float64 `json:"food_proteins"`
	FoodFats        *float64 `json:"food_fats"`

	// Food *BasicFoodCreate `json:"food"`
}

type MealFoodsResponse struct {
	Id           int64  `json:"id"`
	MealId       int64  `json:"meal_id"`
	FoodId       *int64 `json:"food_id"`
	CustomFoodId *int64 `json:"custom_food_id"`
	Qty          *int16 `json:"qty"`
}

type DailyMealFoods struct {
	Id          int64    `json:"id"`
	MealTypeId  int64    `json:"meal_type_id"`
	Qty         int16    `json:"qty"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Kcal        *float64 `json:"kcal"`
	Carbs       *float64 `json:"carbs"`
	Proteins    *float64 `json:"proteins"`
	Fats        *float64 `json:"fats"`
}
