package service

type (
	IUser interface {
	}

	User struct{}
)

func NewUser() IUser {
	return &User{}
}
