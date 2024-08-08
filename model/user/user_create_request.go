package user

type UserCreateRequest struct {
	UserName string `validate:"required,min=1,max=255"`
	Password string `validate:"required,min=1,max=255"`
}
