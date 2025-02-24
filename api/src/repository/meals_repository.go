package repository

import (
	"toyfit/config"
	"toyfit/src/models"

	_ "github.com/lib/pq"
)

type MealsRepository struct{}

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

func (r *MealsRepository) Create(loggedUserId int, model models.MealRequest) (int, error) {
	sqlStatement := `
		insert into meals (created_by, user_id, meal_type_id, date) 
		values ($1, $2, $3, $4) returning id

	`
	var id int
	err := config.DB.QueryRow(
		sqlStatement,
		loggedUserId,
		model.UserId,
		model.MealTypeId,
		model.Date,
	).Scan(&id)

	return id, err
}
