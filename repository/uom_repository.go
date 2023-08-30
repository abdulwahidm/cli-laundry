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
	DeleteById(id string) // DELETE FROM apa WHERE id = ?
}

type uomRepository struct {
	db *sql.DB
}

// DeleteById implements UomRepository.
func (u *uomRepository) DeleteById(id string) {
	panic("unimplemented")
}

// FindAll implements UomRepository.
func (u *uomRepository) FindAll() ([]model.Uom, error) {
	panic("unimplemented")
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

// Save implements UomRepository.
func (u *uomRepository) Save(uom model.Uom) error {
	_, err := u.db.Exec("INSERT INTO uom VALUES ($1, $2)", uom.Id, uom.Name)
	if err != nil {
		return err
	}
	return nil
}

// Update implements UomRepository.
func (u *uomRepository) Update(uom model.Uom) error {
	panic("unimplemented")
}

func NewUomRepository(db *sql.DB) UomRepository {
	return &uomRepository{db: db}
}
