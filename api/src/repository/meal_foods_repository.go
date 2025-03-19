package repository

import (
	"toyfit/config"
	"toyfit/src/models"

	_ "github.com/lib/pq"
)

type MealFoodsRepository struct{}

// func (r *MealFoodsRepository) GetAll(loggedUserId int) ([]models.ResponseTransaction, error) {
// 	sqlStatement := `
// 		SELECT
// 			id,
// 			to_char(date, 'YYYY-MM-DD HH24:MI:SS') AS date,
// 			amount,
// 			type_id,
// 			description
// 		FROM transactions
// 		WHERE user_id = $1
// 		ORDER BY date DESC, id DESC
// 	`
// 	var data []models.ResponseTransaction
// 	rows, err := config.DB.Query(sqlStatement, loggedUserId)
// 	if err != nil {
// 		return data, err
// 	}
// 	for rows.Next() {
// 		var transaction models.ResponseTransaction
// 		err := rows.Scan(
// 			&transaction.Id,
// 			&transaction.Date,
// 			&transaction.Amount,
// 			&transaction.TypeId,
// 			&transaction.Description,
// 		)
// 		if err != nil {
// 			return data, err
// 		}
// 		data = append(data, transaction)
// 	}
// 	return data, err
// }

func (r *MealFoodsRepository) Create(loggedUserId int64, model models.MealFoodsRequest) (int64, error) {
	sqlStatement := `
		insert into meal_foods (created_by, meal_type_id, date, food_id, custom_food_id, qty) 
		values ($1, $2, $3, $4, $5, $6) returning id
	`
	var id int64
	err := config.DB.QueryRow(
		sqlStatement,
		loggedUserId,
		model.MealTypeId,
		model.Date,
		model.FoodId,
		model.CustomFoodId,
		model.Qty,
	).Scan(&id)

	return id, err
}
