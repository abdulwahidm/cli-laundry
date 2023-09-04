package repository

import (
	"database/sql"
	"go-laundry/model"
)

type UomRepository interface {
	Save(uom model.Uom) error              // INSERT
	FindById(id string) (model.Uom, error) // SELECT by id
	FindAll() ([]model.Uom, error)         // SELECT *
	Update(uom model.Uom) error
	FindByName(name string) ([]model.Uom, error)
	DeleteById(id string) error // DELETE FROM table WHERE id = ?
}

type uomRepository struct {
	db *sql.DB
}

// FindByName implements UomRepository.
func (u *uomRepository) FindByName(name string) ([]model.Uom, error) {
	query := ("SELECT * FROM uom WHERE name ILIKE '%" + name + "%';")
	// fmt.Println(query)
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	var uoms []model.Uom
	for rows.Next() {
		uom := model.Uom{}
		err := rows.Scan(
			&uom.Id,
			&uom.Name,
		)
		if err != nil {
			return nil, err
		}
		uoms = append(uoms, uom)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return uoms, nil
}

// Save implements UomRepository.
func (u *uomRepository) Save(uom model.Uom) error {
	_, err := u.db.Exec("INSERT INTO uom VALUES ($1, $2)", uom.Id, uom.Name)
	if err != nil {
		return err
	}
	return nil
}

func (u *uomRepository) DeleteById(id string) error {
	_, err := u.db.Exec("DELETE FROM uom WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements UomRepository.
func (u *uomRepository) FindAll() ([]model.Uom, error) {
	rows, err := u.db.Query("SELECT * FROM uom")
	if err != nil {
		return nil, err
	}

	var uoms []model.Uom
	for rows.Next() {
		var uom model.Uom
		err := rows.Scan(&uom.Id, &uom.Name)
		if err != nil {
			return nil, err
		}
		uoms = append(uoms, uom)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return uoms, nil
}

// FindById implements UomRepository.
func (u *uomRepository) FindById(id string) (model.Uom, error) {
	row := u.db.QueryRow("SELECT id, name FROM uom WHERE id = $1", id)
	var uom model.Uom
	err := row.Scan(&uom.Id, &uom.Name)
	if err != nil {
		return model.Uom{}, err
	}
	return uom, nil
}

// Update implements UomRepository.
func (u *uomRepository) Update(uom model.Uom) error {
	// panic("unimplemented")
	_, err := u.db.Exec("UPDATE uom SET name = $2 WHERE id = $1", uom.Id, uom.Name)
	if err != nil {
		return err
	}
	return nil
}

func NewUomRepository(db *sql.DB) UomRepository {
	return &uomRepository{db: db}
}
