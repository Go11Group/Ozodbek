package crudusers

import (
	model "Projects/N11/modelUser"
	"database/sql"
)

type CrudUsersRepo struct {
	DB *sql.DB
}

func NewCrudUsersRepo(DB *sql.DB) *CrudUsersRepo {
	return &CrudUsersRepo{DB}
}

func (i *CrudUsersRepo) Insertuser(username, email, password string) ([]model.Users, error) {
	tr, err := i.DB.Begin()

	if err != nil {
		return nil, err
	}

	defer tr.Commit()

	_, err = i.DB.Exec("insert into users(username,email,password) values($1,$2,$3)",
		username, email, password)

	if err != nil {
		return nil, err
	}

	return []model.Users{}, nil
}

func (u *CrudUsersRepo) Readuser() ([]model.Users, error) {
	tr, err := u.DB.Begin()

	if err != nil {
		return nil, err
	}

	defer tr.Commit()

	row, err := u.DB.Query(`select * from users`)

	if err != nil {
		return nil, err
	}

	var models []model.Users
	for row.Next() {
		var model model.Users
		err = row.Scan(&model.Id, &model.Username, &model.Email, &model.Password)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}
	return models, nil
}

func (up *CrudUsersRepo) UpdateUser(password string, id int) ([]model.Users, error) {
	tr, err := up.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	_, err = up.DB.Exec("update users set password=$1 where id=$2", password, id)
	if err != nil {
		return nil, err
	}
	return []model.Users{}, nil
}

func (d *CrudUsersRepo) Deleteuser(id int) ([]model.Users, error) {
	tr, err := d.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	_, err = d.DB.Exec("delete from users where id=$1", id)
	if err != nil {
		return nil, err
	}
	return []model.Users{}, nil
}