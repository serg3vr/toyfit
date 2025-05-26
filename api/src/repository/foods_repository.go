package repository

import (
	"toyfit/config"
	"toyfit/src/models"

	_ "github.com/lib/pq"
)

type FoodsRepository struct{}

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

func (r *FoodsRepository) Create(loggedUserId int64, model models.FoodRequest) (int64, error) {
	sqlStatement := `
		insert into foods (created_by, name, description, kcal, carbs, proteins, fats, sodium) 
		values ($1, $2, $3, $4, $5, $6, $7, $8) returning id
	`
	var id int64
	err := config.DB.QueryRow(
		sqlStatement,
		loggedUserId,
		model.Name,
		model.Description,
		model.Kcal,
		model.Carbs,
		model.Proteins,
		model.Fats,
		model.Sodium,
	).Scan(&id)

	return id, err
}

func (r *FoodsRepository) Update(loggedUserId int64, model models.FoodRequest) (int64, error) {
	sqlStatement := `
		update foods set 
			updated_by = $1,
			name = $3, 
			description = $4,
			kcal = $5,
			carbs = $6,
			proteins = $7,
			fats = $8,
			sodium = $9
		where id = $2
		returning id
	`
	var id int64
	err := config.DB.QueryRow(
		sqlStatement,
		loggedUserId,
		model.Id,
		model.Name,
		model.Description,
		model.Kcal,
		model.Carbs,
		model.Proteins,
		model.Fats,
		model.Sodium,
	).Scan(&id)

	return id, err
}
