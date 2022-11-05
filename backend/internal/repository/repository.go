package repository

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"tumaris.hack-FemNovation/backend/internal/repository/auth"
	"tumaris.hack-FemNovation/backend/internal/repository/order"
	"tumaris.hack-FemNovation/backend/internal/repository/products"
	"tumaris.hack-FemNovation/backend/internal/repository/token"
)

type Repository struct {
	Auth     auth.Auth
	Products products.Products
	Token    token.Token
	Order    order.Order
}

func New(db *pgxpool.Pool, sqlite *sql.DB, sqliteTimeout time.Duration, logger *zap.SugaredLogger) *Repository {
	dbTimeout := 10 * time.Second

	p1 := `-----BEGIN RSA PRIVATE KEY-----
	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
	-----END RSA PRIVATE KEY-----`

	p2 := `-----BEGIN PUBLIC KEY-----
	MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAwdji+YqZLwQef6tg6d7G
	gorGen+EcjSGAdhm3L8WT8ULZdmL6Llil6VpidQdE5GcxnLqQA8SIwC9JLGTpMy3
	Nc5vXE19e8vcXuhqZTVNC/2VqhI0/HtU2Bmqmi11WIXQSkgDT3BUTGfJpDJqLKwA
	66XhcklB5GbPsqq788KpsuZhqKQc+1GFqE6vnAlb2+Txz9RQg/11gCzUdJQQjZZH
	6RvlqFyEcl7LGUKaqHfRPAqEjTTZLgFTalGYwKr+62O8Ww+sZ+02sQDiGDXFYpcS
	oPld8kZd873WC0kBodEaS2FIeAT7f3LctR7Cjcwg+ibtA2eaTJJdYsHUvbNYXHJ3
	7AGoDMCvun5G1nuu4ze+6KO/4vLLCsET7CaQy/+LzQ6/sRr14SFgZTDxH1Zmz1C1
	2qtx4zWAfdzP/yqL1JQRJXXzK9dTj+3GnW+uOZPGtL5ImQPhwAD5xcoCBB2LI5qi
	IAeASonx1Oqg8IdV5sdcI/MF+JohoN04/OxpWmJceTLONIalB6tfBAGMw1UiHkVx
	1ou40lyR7l4kofyee0GJLXkOb1nQFsuEe7ut7TRtGprA4j+YAhrf6M9lqSBg4+6x
	+euoputZ7p0G0ZWxUYNjEjG/OoTauIQE23dKJQimLCCdKZlHAchd+FzDRAC3NApJ
	nmOfaxgYUyuc3taA1ro3IqsCAwEAAQ==
	-----END PUBLIC KEY-----`

	return &Repository{
		Auth:     auth.NewAuthRepo(logger, sqlite, sqliteTimeout),
		Products: products.NewProductsRepo(logger, sqlite, sqliteTimeout),
		Order:    order.NewOrderRepo(logger, db, dbTimeout),
		Token:    token.NewTokenRepo(logger, sqlite, sqliteTimeout, []byte(p1), []byte(p2)),
	}
}
