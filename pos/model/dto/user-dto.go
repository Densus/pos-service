package dto

//type userType string
//
//const (
//	CUSTOMER userType = "CUSTOMER"
//	MERCHANT userType = "MERCHANT"
//	OUTLET userType = "OUTLET"
//)
//func (ut *userType) Scan(value interface{}) error {
//	*ut = userType(value.([]byte))
//	return nil
//}
//
//func (ut userType) Value() (driver.Value, error) {
//	return string(ut), nil
//}

type RegisterDTO struct {
	UserName string `form:"user_name" binding:"required" validate:"required,gte=3,lte=80"`
	FullName string `form:"full_name" binding:"required" validate:"required,gte=3,lte=80"`
	Email string `form:"email" binding:"required" validate:"required,email,lte=80"`
	Password string `form:"password" binding:"required" validate:"required,gte=8,lte=80"`
	Role string `form:"role" binding:"required"`
}

type LoginDTO struct {
	Email string `form:"email" binding:"required" validate:"required,email,lte=80"`
	Password string `form:"password" binding:"required" validate:"required,gte=8,lte=80"`
}

type UpdateUserDTO struct {
	ID uint32 `form:"id"`
	UserName string `form:"user_name" binding:"required" validate:"required,gte=3,lte=80"`
	FullName string `form:"full_name" binding:"required" validate:"required,gte=3,lte=80"`
	Email string `form:"email" binding:"required" validate:"required,email,lte=80"`
	Password string `form:"password" binding:"required" validate:"required,gte=8,lte=80"`
	Role string `form:"role"`
}