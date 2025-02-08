package userRepository

func (r *UserRepository) GetUsers() ([]User, error) {
	rows, queryError := r.DB.Query("SELECT * FROM Users")
	if queryError != nil {
		panic(queryError.Error())
	}

	var users []User

	for rows.Next() {
		var user User
		if error := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.Password); error != nil {
			return nil, error
		}
		users = append(users, user)
	}

	return users, nil
}
