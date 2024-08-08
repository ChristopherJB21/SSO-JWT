package user

type UserUpdatePasswordRequest struct {
	IDUser      uint   `validate:"required"`
	NewPassword string `validate:"required,min=1,max=255"`
}
