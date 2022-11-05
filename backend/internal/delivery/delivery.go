package delivery

import (
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"tumaris.hack-FemNovation/backend/internal/service"
)

type Handler struct {
	service *service.Service
	logger  *zap.SugaredLogger

	tokenInfo *Token
}

type Token struct {
	Pepper     string
	Access     *Access
	Refresh    *Refresh
	Domain     string
	Secure     bool
	HttpOnly   bool
	PrivateKey string
	PublicKey  string
}

type Access struct {
	TTL time.Duration
}

type Refresh struct {
	TTL time.Duration
}

func NewHandler(services *service.Service, sugar *zap.SugaredLogger) *Handler {
	token := &Token{
		Pepper: "fjdskljdsfldsfdsjldsjflie4r",
		Access: &Access{
			TTL: 15 * time.Hour,
		},
		Refresh: &Refresh{
			TTL: 15 * time.Hour,
		},
		Domain:     "kustoma.shop",
		Secure:     false,
		HttpOnly:   true,
		PrivateKey: "",
		PublicKey:  "",
	}

	return &Handler{
		service:   services,
		logger:    sugar,
		tokenInfo: token,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/sign-in", h.SignIn)
	router.OPTIONS("/sign-in", h.OK)

	router.GET("/refresh", h.Refresh)
	router.OPTIONS("/refresh", h.OK)

	router.GET("/sign-out", h.SignOut)
	router.OPTIONS("/sign-out", h.OK)

	router.GET("/categories/:category_id/products/:product_id")
	router.DELETE("/categories/:category_id/products/:product_id")
	router.POST("/categories/:category_id/products/:product_id")
	router.GET("/categories/:category_id")

	router.GET("/products")

	router.POST("/categories/:category_id/products/:product_id/order")
	router.GET("/orders/:order_id")

	return router
}
