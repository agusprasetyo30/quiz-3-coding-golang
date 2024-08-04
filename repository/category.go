package repository

import (
	"database/sql"
	"errors"
	"quiz-3/model"
	"time"
)

func GetAllCategory(db *sql.DB) (result []model.Category, err error) {
	sql := "SELECT * FROM categories"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var category model.Category

		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
		if err != nil {
			return
		}

		result = append(result, category)
	}

	return
}

func GetCategory(db *sql.DB, category model.Category) (*model.Category, error) {
	sql := "SELECT * FROM categories WHERE id = $1"

	errs := db.QueryRow(sql, category.ID).Scan(&category.ID,
		&category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)

	if errs != nil {
		return nil, errs
	}

	return &category, nil
}

func InsertCategory(db *sql.DB, category model.Category) error {
	sql := "INSERT INTO categories (name, created_at, created_by) VALUES ($1,$2,$3) RETURNING *"
	_, errs := db.Exec(sql, category.Name, time.Now(), category.CreatedBy)

	return errs
}

func UpdateCategory(db *sql.DB, category model.Category) error {
	sql := "UPDATE categories SET name = $1, modified_at = $2, modified_by = $3 WHERE id = $4"

	errs := db.QueryRow(sql, category.Name, time.Now(), category.ModifiedBy, category.ID)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category model.Category) error {
	// return errs.Err()

	sql := "DELETE FROM categories WHERE id = $1"

	result, err1 := db.Exec(sql, category.ID)
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

func GetBookByCategory(db *sql.DB, id int) (result []model.Book, err error) {
	sql := "SELECT b.id, b.category_id, c.*, b.description, b.image_url, b.release_year, b.price, b.total_page, b.thickness, b.created_at, b.created_by, b.modified_at, b.modified_by FROM books b INNER JOIN categories c ON b.category_id = c.id WHERE b.category_id = $1"

	rows, err := db.Query(sql, id)
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
