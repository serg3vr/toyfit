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
		insert into meal_foods (created_by, user_id, meal_type_id, date, food_id, custom_food_id, qty) 
		values ($1, $2, $3, $4, $5, $6, $7) returning id
	`
	var id int64
	err := config.DB.QueryRow(
		sqlStatement,
		loggedUserId,
		loggedUserId,
		model.MealTypeId,
		model.Date,
		model.FoodId,
		model.CustomFoodId,
		model.Qty,
	).Scan(&id)

	return id, err
}

func (r *MealFoodsRepository) Delete(loggedUserId int64, id int64) (int64, error) {
	sqlStatement := `
		delete from meal_foods where user_id = $1 and id = $2
	`
	result, err := config.DB.Exec(sqlStatement, loggedUserId, id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	return rowsAffected, err
}

func (r *MealFoodsRepository) GetDaily(loggedUserId int64, timeZone string, date string) ([]models.DailyMealFoods, error) {
	query := `
		select
			mf.id,
			mf.meal_type_id, 
			f.name 
		from meal_foods mf
		join foods f on f.id = mf.food_id
		where 
			user_id = $1
			and date(date at time zone '` + timeZone + `')  = $2
		order by mf.meal_type_id, mf.created_at
	`
	// var data []models.DailyMealFoods
	data := make([]models.DailyMealFoods, 0)
	rows, err := config.DB.Query(query, loggedUserId, date)
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var model models.DailyMealFoods
		err := rows.Scan(
			&model.Id,
			&model.MealTypeId,
			&model.Name,
		)
		if err != nil {
			return data, err
		}
		data = append(data, model)
	}
	return data, err
}
