package repo

import (
	"errors"
	"github.com/bangnh1/golang-training/10/models"
	"time"
)

/* Lấy danh sách user */
func GetAllUser() (users []models.User, err error) {
	err = DB.Model(&users).Select()
	if err != nil {
		return nil, err
	}

	return users, nil
}

/* Lấy thông tin user */
func GetUserById(id string) (user *models.User, err error) {
	user = &models.User{
		Id: id,
	}

	err = DB.Model(user).WherePK().Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

/* Tạo user */
func CreateUser(req *models.CreateUser) (user *models.User, err error) {
	if req.FullName == "" {
		return nil, errors.New("tên không được để trống")
	}

	user = &models.User{
		Id:        NewID(),
		FullName:  req.FullName,
		Email:     req.Email,
		Phone:     req.Phone,
		Age:       req.Age,
		Sex:       req.Sex,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = DB.Model(user).WherePK().Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return user, nil
}

/* Cập nhật thông tin user */
func UpdateUser(id string, req *models.CreateUser) (user *models.User, err error) {
	if req.FullName == "" {
		return nil, errors.New("tên không được để trống")
	}

	user = &models.User{
		Id:        id,
		FullName:  req.FullName,
		Email:     req.Email,
		Phone:     req.Phone,
		Age:       req.Age,
		Sex:       req.Sex,
		UpdatedAt: time.Now(),
	}

	_, err = DB.Model(user).Column("full_name", "phone", "email", "age", "sex", "updated_at").Returning("*").WherePK().Update()
	if err != nil {
		return nil, err
	}

	return user, nil
}

/* Xóa user */
func DeleteUser(id string) (err error) {
	user := models.User{
		Id: id,
	}

	_, err = DB.Model(&user).WherePK().Delete()
	if err != nil {
		return err
	}

	return nil
}
