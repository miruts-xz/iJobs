package repository

import (
	"database/sql"
	"errors"
	"github.com/miruts/iJobs/entity"
)

// CategoryRepositoryImpl implements the menu.CategoryRepository interface
type CategoryRepositoryImpl struct {
	conn *sql.DB
}

// NewCategoryRepositoryImpl will create an object of PsqlCategoryRepository
func NewCategoryRepositoryImpl(Conn *sql.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{conn: Conn}
}

// Categories returns all cateogories from the database
func (cri *CategoryRepositoryImpl) Categories() ([]entity.Category, error) {

	rows, err := cri.conn.Query("SELECT * FROM job_categories;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Category{}
	for rows.Next() {
		category := entity.Category{}
		err = rows.Scan(&category.ID, &category.Name, &category.Desc, &category.Image)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, category)
	}

	return ctgs, nil
}

// Category returns a category with a given id
func (cri *CategoryRepositoryImpl) Category(id int) (entity.Category, error) {

	row := cri.conn.QueryRow("SELECT * FROM job_categories WHERE id = $1", id)

	c := entity.Category{}

	err := row.Scan(&c.ID, &c.Name, &c.Desc, &c.Image)
	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateCategory updates a given object with a new data
func (cri *CategoryRepositoryImpl) UpdateCategory(c entity.Category) error {

	_, err := cri.conn.Exec("UPDATE job_categories SET name=$1,short_desc=$2, image=$3 WHERE id=$4", c.Name, c.Desc, c.Image, c.ID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteCategory removes a category from a database by its id
func (cri *CategoryRepositoryImpl) DeleteCategory(id int) error {

	_, err := cri.conn.Exec("DELETE FROM job_categories WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreCategory stores new category information to database
func (cri *CategoryRepositoryImpl) StoreCategory(c entity.Category) error {

	_, err := cri.conn.Exec("INSERT INTO job_categories (name,short_desc,image) values($1, $2, $3)", c.Name, c.Desc, c.Image)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
