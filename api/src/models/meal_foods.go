package models

import (
// "time"
)

type MealFoodsRequest struct {
	MealId       int64  `json:"meal_id"`
	FoodId       *int64 `json:"food_id"`
	CustomFoodId *int64 `json:"custom_food_id"`
	Qty          *int16 `json:"qty"`
}

type MealFoodsResponse struct {
	Id           int64  `json:"id"`
	MealId       int64  `json:"meal_id"`
	FoodId       *int64 `json:"food_id"`
	CustomFoodId *int64 `json:"custom_food_id"`
	Qty          *int16 `json:"qty"`
}

// type ResponseTransaction struct {
// 	Id          int64   `json:"id"`
// 	Date        string  `json:"date"`
// 	Amount      float64 `json:"amount"`
// 	TypeId      int64   `json:"type"`
// 	Description *string `json:"description"`
// }

// type DBTransaction struct {
// 	Id          int64      `json:"id"`
// 	CreatedAt   time.Time  `json:"created_at"`
// 	CreatedBy   int        `json:"created_by"`
// 	UpdatedAt   *time.Time `json:"updated_at"`
// 	UpdatedBy   *int       `json:"updated_by"`
// 	UserId      int64      `json:"user_id"`
// 	Date        string     `json:"date"`
// 	Amount      float64    `json:"amount"`
// 	TypeId      int64      `json:"type"`
// 	Description *string    `json:"description"`
// }
