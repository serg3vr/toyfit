package models

import (
	"time"
)

type MealRequest struct {
	UserId     int64      `json:"user_id"`
	MealTypeId int64      `json:"meal_type_id"`
	Date       *time.Time `json:"date"`
}

type MealResponse struct {
	Id           int64     `json:"id"`
	UserId       int64     `json:"user_id"`
	MealTypeId   int64     `json:"meal_type_id"`
	Date         time.Time `json:"date"`
	MealTypeName string    `json:"meal_type_name"`
}

type MealWithFoodsRequest struct {
	MealRequest
	MealFoodsRequest *[]MealFoodsRequest `json:"meal_foods_request"`
}

type MealWithFoodsResponse struct {
	MealResponse
	MealFoodsResponse *[]MealFoodsResponse `json:"meal_foods_response"`
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
