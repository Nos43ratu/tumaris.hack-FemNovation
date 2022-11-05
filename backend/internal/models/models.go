package models

import "time"

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

type Product struct {
	ProductID int `json:"id"`
	ShopID int `json:"shop_id"`
	Name    string `json:"name"`
	Description string `json:"description"`
	Sizes     []string `json:"sizes,omitempty"`
	Colors []int `json:"colors,omitempty"`
	Weight float64 `json:"weight"`
	Price float64 `json:"price"`
	Rating float64 `json:"rating"`
	CategoryID int `json:"category_id"`
}
type Order struct {
	ID           int       `json:"id,omitempty"`
	Status       int       `json:"status,omitempty"`
	ClientID     int       `json:"client_id,omitempty"`
	ShopID       int       `json:"shop_id,omitempty"`
	ProductID    int       `json:"product_id,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	PayedAt      time.Time `json:"payed_at,omitempty"`
	PackedAt     time.Time `json:"packed_at,omitempty"`
	DeliveredAt  time.Time `json:"delivered_at,omitempty"`
	CancelReason string    `json:"cancel_reason,omitempty"`
}
