package user

import (
	"simple-crud/internal/domain"
	"simple-crud/pkg/xlogger"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{db: db}
}

func (r *mysqlUserRepository) Create(user domain.User) error {

	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func  (r *mysqlUserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *mysqlUserRepository) Update(id int, user domain.User) error {
xlogger.Logger.Debug().Msgf("Updating user with ID: %d, Name: %s", id, user.Name)
	if err := r.db.Model(&domain.User{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *mysqlUserRepository) Delete(id int) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.User{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *mysqlUserRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}