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
