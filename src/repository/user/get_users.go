package userRepository

import "pam/src/domain/entity"

func (r *UserRepository) GetUsers() ([]entity.User, error) {
	rows, queryError := r.DB.Query("SELECT * FROM Users")
	if queryError != nil {
		panic(queryError.Error())
	}

	var users []entity.User

	for rows.Next() {
		var user entity.User
		if error := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.Password); error != nil {
			return nil, error
		}
		users = append(users, user)
	}

	return users, nil
}
