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
	table          = "book"
	layoutDateTime = "2006-01-02 15:04:05"
)

// Insert Book

func Insert(ctx context.Context, book models.Book) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"INSERT INTO %v (title, description, image_url, release_year, price, total_page, thickness) values('%v','%v','%v',%v, '%v','%v','%v', NOW(), NOW())",
		table,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_year,
		book.Price,
		book.Total_page,
		book.Thickness)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Get All

func GetAll(ctx context.Context) ([]models.Book, error) {
	var bookAll []models.Book
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Book
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&book.ID,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&book.Thickness,
			&createdAt,
			&updatedAt,
			&book.CategoryId); err != nil {
			return nil, err
		}

		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		bookAll = append(bookAll, book)
	}
	return bookAll, nil
}

// Update

func Update(ctx context.Context, book models.Book, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set title = '%s', description = '%s', image_url = '%s', release_year = %d, price = '%s', total_page = '%s', thickness = '%s',category_id = %d ,updated_at = NOW() where id = %d",
		table,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_year,
		book.Price,
		book.Total_page,
		book.Thickness,
		book.ID,
		book.CategoryId)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

// Delete

func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

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
