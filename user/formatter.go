package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"string"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"toke"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         int(user.ID),
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
	return formatter
}
