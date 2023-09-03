package auth

type User struct {
	Id       int    `json:"-"`
	Login    string `json:"login" binding:"required" validate:"required,min=3,max=30"`
	Password string `json:"password" binding:"required" validate:"required,min=8,max=16"`
	Name     string `json:"name" binding:"required" validate:"required,min=3,max=30"`
	Age      int    `json:"age" binding:"required" validate:"required,gte=0,lte=130"`
}

type UserAuthFields struct {
	Login    string `json:"login" binding:"required" validate:"required,min=3,max=30"`
	Password string `json:"password" binding:"required" validate:"required,min=8,max=16"`
}

type UserProfile struct {
	Id   int    `json:"id" binding:"required" validate:"required"`
	Name string `json:"name" binding:"required" validate:"required,min=3,max=30"`
	Age  int    `json:"age" binding:"required" validate:"required,gte=0,lte=130"`
}

type UserTokenClaims struct {
	Login  string `json:"login"`
	UserId int    `json:"user_id"`
}
