package user_domain

type User struct {
	Uuid string
	Email UserEmail
	Name UserName
	Password UserPassword
}

type NewUserProps struct {
	UserEmail UserEmail
	UserName UserName
	UserPassword UserPassword
}

func NewUserCreate(props NewUserProps) (User, error) {
	return User{
		Email: props.UserEmail,
		Name: props.UserName,
		Password: props.UserPassword,
	}, nil
}