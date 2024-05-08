package user

import "go-technical-test-bankina/src/entity"

type UserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserDetailFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatUser(user entity.User, token string) UserFormatter {
	formatUser := UserFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return formatUser
}

func FormatUserDetail(user entity.User) UserDetailFormatter {
	formatUser := UserDetailFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return formatUser
}

func FormatUsers(users []entity.User) []UserDetailFormatter {
	usersFormatter := []UserDetailFormatter{}

	for _, user := range users {
		UserFormatter := FormatUserDetail(user)
		usersFormatter = append(usersFormatter, UserFormatter)
	}

	return usersFormatter
}
