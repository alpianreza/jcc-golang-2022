package books

import (
	"Quiz-3/config"
	"Quiz-3/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	tableCategory = "Category"
)

// Insert Category

func InsertCategory(ctx context.Context, Category models.Category) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"INSERT INTO %v (name) values('%v') NOW(), NOW())",
		tableCategory,
		Category.Name,
	)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Get All

func GetAllCategory(ctx context.Context) ([]models.Category, error) {
	var CategoryAll []models.Category
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v", tableCategory)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var Category models.Category
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&Category.ID,
			&Category.Name,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		Category.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		Category.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		CategoryAll = append(CategoryAll, Category)
	}
	return CategoryAll, nil
}

// Update

func UpdateCategory(ctx context.Context, Category models.Category, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set name = '%s', updated_at = NOW() where id = %d",
		tableCategory,
		Category.Name,
		Category.ID)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

// Delete

func DeleteCategory(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", tableCategory, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id not found")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
