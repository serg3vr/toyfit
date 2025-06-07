package repository

import (
	"toyfit/config"
	"toyfit/src/models"

	_ "github.com/lib/pq"
)

type FoodsRepository struct{}

func (r *FoodsRepository) GetAll() ([]models.FoodResponse, error) {
	query := `
		SELECT
			id,
			name,
			description,
			kcal,
			carbs,
			proteins,
			fats,
			sodium
		FROM foods
		ORDER BY name DESC
	`
	var data []models.FoodResponse
	rows, err := config.DB.Query(query)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		var model models.FoodResponse
		err := rows.Scan(
			&model.Id,
			&model.Name,
			&model.Description,
			&model.Kcal,
			&model.Carbs,
			&model.Proteins,
			&model.Fats,
			&model.Sodium,
		)
		if err != nil {
			return data, err
		}
		data = append(data, model)
	}

	return data, err
}

func (r *FoodsRepository) GetById(id int64) (models.FoodResponse, error) {
	query := `
		SELECT
			id,
			name,
			description,
			kcal,
			carbs,
			proteins,
			fats,
			sodium
		FROM foods
		WHERE id = $1
		ORDER BY name DESC
	`
	var model models.FoodResponse
	err := config.DB.QueryRow(query, id).Scan(
		&model.Id,
		&model.Name,
		&model.Description,
		&model.Kcal,
		&model.Carbs,
		&model.Proteins,
		&model.Fats,
		&model.Sodium,
	)

	if err != nil {
		return model, err
	}

	return model, err
}

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
