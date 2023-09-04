package auth

type PhoneNumber struct {
	Id          int    `json:"id" db:"id"`
	UserId      int    `json:"user_id" db:"user_id" validate:"required"`
	Phone       string `json:"phone" validate:"required,min=12,max=12"`
	IsFax       bool   `json:"is_fax"`
	Description string `json:"description" validate:"max=255"`
}

type UpdatingPhoneNumber struct {
	Id          int    `json:"id" db:"id" validate:"required"`
	UserId      int    `json:"user_id" db:"user_id" validate:"required"`
	Phone       string `json:"phone" validate:"required,min=12,max=12"`
	IsFax       bool   `json:"is_fax"`
	Description string `json:"description" validate:"required,max=255"`
}
