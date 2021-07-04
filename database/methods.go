package database

import "github.com/siddhanthpx/phonebook/models"

func RegisterUser(user models.User) error {
	return DB.Create(&user).Error
}

func VerifyUser(phone string) (models.User, error) {

	var user models.User
	if err := DB.Where("phone_number = ?", phone).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetUser(id string) (models.User, error) {
	var user models.User

	if err := DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
