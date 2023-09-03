package auth

type PhoneNumber struct {
	Id          int    `json:"-" db:"id"`
	UserId      int    `json:"-" db:"user_id" validate:"required"`
	Phone       string `json:"phone" validate:"required,min=12,max=12"`
	IsFax       bool   `json:"is_fax" validate:"required"`
	Description string `json:"description" validate:"max=255"`
}
