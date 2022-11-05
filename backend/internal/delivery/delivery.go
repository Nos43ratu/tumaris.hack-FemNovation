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

	api := router.Group("/api")

	api.POST("/sign-in", h.SignIn)
	api.OPTIONS("/sign-in", h.OK)

	api.GET("/refresh", h.Refresh)
	api.OPTIONS("/refresh", h.OK)

	router.GET("/categories/:category_id/products/:product_id", h.CreateOrUpdateProduct)
	router.DELETE("/categories/:category_id/products/:product_id")
	router.POST("/categories/:category_id/products/:product_id")
	router.GET("/categories/:category_id")

	api.GET("/categories/:category_id/products/:product_id")
	api.DELETE("/categories/:category_id/products/:product_id")
	api.POST("/categories/:category_id/products/:product_id")
	api.GET("/categories/:category_id")

	api.GET("/products")

	api.POST("/categories/:category_id/products/:product_id/order")
	api.GET("/orders/:order_id")

	return router
}
