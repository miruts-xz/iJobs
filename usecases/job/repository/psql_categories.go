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
		err = rows.Scan(&category.ID, &category.Name, &category.Descr, &category.Image)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, category)
	}

	return ctgs, nil
}

//Category returns a category with a given id
func (cri *CategoryRepositoryImpl) Category(id int) (entity.Category, error) {

	row := cri.conn.QueryRow("SELECT * FROM categories WHERE id = $1", id)

	c := entity.Category{}

	err := row.Scan(&c.ID, &c.Name, &c.Descr, &c.Image)
	if err != nil {
		return c, err
	}

	return c, nil
}

//UpdateCategory updates a given object with a new data
func (cri *CategoryRepositoryImpl) UpdateCategory(c *entity.Category) (*entity.Category, error) {
	_, err := cri.conn.Exec("UPDATE categories SET name=$1,descr=$2, image=$3 WHERE id=$4", c.Name, c.Descr, c.Image, c.ID)
	if err != nil {
		return c, errors.New("Update has failed")
	}

	return c, nil
}

//DeleteCategory removes a category from a database by its id
func (cri *CategoryRepositoryImpl) DeleteCategory(id int) (entity.Category, error) {
	var category entity.Category
	_, err := cri.conn.Exec("DELETE FROM categories WHERE id=$1", id)
	if err != nil {
		return category, errors.New("Delete has failed")
	}

	return category, nil
}

//StoreCategory stores new category information to database
func (cri *CategoryRepositoryImpl) StoreCategory(c *entity.Category) (*entity.Category, error) {

	_, err := cri.conn.Exec("INSERT INTO categories (name,descr,image) values($1, $2, $3)", c.Name, c.Descr, c.Image)
	if err != nil {
		return c, errors.New("Insertion has failed")
	}

	return c, nil
}
