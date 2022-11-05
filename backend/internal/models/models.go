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
	ID           int    `json:"id,omitempty"`
	Status       int    `json:"status,omitempty"`
	ClientID     int    `json:"client_id,omitempty"`
	ShopID       int    `json:"shop_id,omitempty"`
	ProductID    int    `json:"product_id,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	PayedAt      string `json:"payed_at,omitempty"`
	PackedAt     string `json:"packed_at,omitempty"`
	DeliveredAt  string `json:"delivered_at,omitempty"`
	CancelReason string `json:"cancel_reason,omitempty"`
}
