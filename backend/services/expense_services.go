package services

import (
	"context"
	"expense-tracker/config"
	"expense-tracker/models"

	"github.com/jackc/pgx/v5"
)

func GetAllExpenses(category string) ([]models.Expense, error) {
	var rows pgx.Rows
	var err error

	if category != "" {
		rows, err = config.DB.Query(context.Background(),
			"SELECT id, amount, description, category, created_at FROM expenses WHERE category = $1 ORDER BY created_at DESC",
			category)
	} else {
		rows, err = config.DB.Query(context.Background(),
			"SELECT id, amount, description, category, created_at FROM expenses ORDER BY created_at DESC")
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var e models.Expense
		err := rows.Scan(&e.ID, &e.Amount, &e.Description, &e.Category, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, e)
	}

	return expenses, nil
}

func CreateExpense(e models.Expense) error {
	_, err := config.DB.Exec(context.Background(),
		"INSERT INTO expenses (amount, description, category) VALUES ($1, $2, $3)",
		e.Amount, e.Description, e.Category)

	return err
}

func UpdateExpense(id int64, e models.Expense) error {
	_, err := config.DB.Exec(context.Background(),
		"UPDATE expenses SET amount=$1, description=$2, category=$3 WHERE id=$4",
		e.Amount, e.Description, e.Category, id)

	return err
}

func DeleteExpense(id int64) error {
	_, err := config.DB.Exec(context.Background(),
		"DELETE FROM expenses WHERE id=$1", id)

	return err
}
