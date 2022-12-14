package delivery

import (
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"tumaris.hack-FemNovation/backend/internal/service"

	"github.com/gin-contrib/cors"
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
		Pepper: "67072341-eb28-4174-a01f-baf72c40b966",
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

	router.Use(cors.New(cors.Config{
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowOrigins: []string{"*"},
	}))

	api := router.Group("/api")

	api.Use(cors.New(cors.Config{
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowOrigins: []string{"*"},
	}))

	api.POST("/sign-in", h.SignIn)
	api.OPTIONS("/sign-in", h.OK)

	api.GET("/refresh", h.Refresh)
	api.OPTIONS("/refresh", h.OK)

	api.GET("/sign-out", h.SignOut)
	api.OPTIONS("/sign-out", h.OK)

	api.POST("/categories/:category_id/products", h.CreateOrUpdateProduct)
	api.POST("/categories/:category_id/products/:product_id", h.CreateOrUpdateProduct)
	api.DELETE("/categories/:category_id/products/:product_id", h.DeleteProduct)
	api.GET("/categories/:category_id/products/:product_id", h.GetProduct)
	api.GET("/categories/:category_id", h.GetProducts)
	api.GET("/categories", h.GetProducts)

	api.GET("/products", h.GetProducts)

	api.POST("/orders", h.CreateOrder)
	api.POST("/orders/:order_id", h.UpdateOrder)
	api.GET("/users/:user_id/orders", h.GetAllOrders)
	api.GET("/orders/:order_id", h.GetOrder)

	api.GET("/users", h.GetUser)

	return router
}
