package models

type User struct {
	ID       int    `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

type Tokens struct {
	Access  string `json:"access_token,omitempty"`
	Refresh string `json:"refresh_token,omitempty"`
}

type Session struct {
	ID            int
	UserID        int
	RefreshToken  string
	AccessToken   string
	RefreshExpire int64
}

type Order struct {
}
