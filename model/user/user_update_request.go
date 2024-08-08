package user

type UserUpdateRequest struct {
	IDUser   uint   `validate:"required"`
	UserName string `validate:"required,min=1,max=255"`
}
