package entity

type ChangePassword struct {
	PasswordNew     string `valid:"required"`
	PasswordConfirm string `valid:"required"`
	PasswordOld     string `valid:"required"`
}