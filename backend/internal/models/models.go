package models

import (
	"github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

type UserInfo struct {
	ID        int         `json:"id,omitempty"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone_number"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Password  string      `json:"password"`
	Role      string      `json:"role,omitempty"`
	AboutMe   string      `json:"about_me"`
	Instagram string      `json:"instagram"`
	Rating    float64     `json:"rating"`
	ShopID    int         `json:"shop_id"`
	CreatedAt pq.NullTime `json:"created_at"`
	UpdatedAt pq.NullTime `json:"updated_at"`
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
	ProductID   int      `json:"id"`
	ShopID      int      `json:"shop_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Sizes       []string `json:"sizes,omitempty"`
	Colors      []int    `json:"colors,omitempty"`
	Weight      float32  `json:"weight"`
	Price       float32  `json:"price"`
	Rating      float32  `json:"rating"`
	CategoryID  int      `json:"category_id"`
}

type Order struct {
	ID           int         `json:"id,omitempty"`
	Status       int         `json:"status"`
	ClientID     int         `json:"client_id,omitempty"`
	ShopID       int         `json:"shop_id,omitempty"`
	ProductID    int         `json:"product_id,omitempty"`
	CreatedAt    pq.NullTime `json:"created_at,omitempty"`
	PayedAt      pq.NullTime `json:"payed_at,omitempty"`
	PackedAt     pq.NullTime `json:"packed_at,omitempty"`
	DeliveredAt  pq.NullTime `json:"delivered_at,omitempty"`
	CancelReason string      `json:"cancel_reason"`
	Products     *Product    `json:"products,omitempty"`
}
