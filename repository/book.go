package repository

import (
	"database/sql"
	"errors"
	"quiz-3/model"
	"time"
)

func GetAllBook(db *sql.DB) (result []model.Book, err error) {
	sql := "SELECT b.id, b.category_id, c.*, b.description, b.image_url, b.release_year, b.price, b.total_page, b.thickness, b.created_at, b.created_by, b.modified_at, b.modified_by FROM books b INNER JOIN categories c ON b.category_id = c.id"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var book model.Book

		err = rows.Scan(&book.ID, &book.CategoryID, &book.Category.ID, &book.Category.Name, &book.Category.Name, &book.Category.CreatedBy, &book.Category.CreatedAt, &book.Category.ModifiedBy,
			&book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)

		if err != nil {
			return
		}

		result = append(result, book)
	}

	return
}

func InsertBook(db *sql.DB, book model.Book) error {
	sql := "INSERT INTO books (category_id, description, image_url, release_year, price, total_page, thickness, created_at, created_by) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING *"
	_, errs := db.Exec(sql, book.CategoryID, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, time.Now(), book.CreatedBy)

	return errs
}

func GetBook(db *sql.DB, book model.Book) (*model.Book, error) {
	sql := "SELECT b.id, b.category_id, c.*, b.description, b.image_url, b.release_year, b.price, b.total_page, b.thickness, b.created_at, b.created_by, b.modified_at, b.modified_by FROM books b INNER JOIN categories c ON b.category_id = c.id WHERE b.id = $1"

	errs := db.QueryRow(sql, book.ID).Scan(&book.ID, &book.CategoryID, &book.Category.ID, &book.Category.Name, &book.Category.Name, &book.Category.CreatedBy, &book.Category.CreatedAt, &book.Category.ModifiedBy,
		&book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)

	if errs != nil {
		return nil, errs
	}

	return &book, nil
}

func UpdateBook(db *sql.DB, book model.Book) error {
	sql := "UPDATE books SET category_id = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, modified_at = $8, modified_by = $9 WHERE id = $10 RETURNING *"

	_, errs := db.Exec(sql, book.CategoryID, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, time.Now(), book.ModifiedBy, book.ID)

	return errs
}

func DeleteBook(db *sql.DB, book model.Book) error {
	sql := "DELETE FROM books WHERE id = $1"

	result, err1 := db.Exec(sql, book.ID)
	if err1 != nil {
		return err1
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("category not found")
	}

	return nil
}
