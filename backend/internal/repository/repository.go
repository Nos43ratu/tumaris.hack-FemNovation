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

	p1 := `-----BEGIN PRIVATE KEY-----
	MIIEowIBAAKCAQEAn3sCj1h0ZS8muwog8WaAh99nLk3CbnqqrGjeMDoED+9PveQd
	RH4U74l0RQbmLFJ8EFxvzU9vBkvBpl21X3mc5CqULivoiKSYPvnYcMQdtcQDOCJF
	bNWKJzdS+ceXIaIb3aSQhXeFIWm8qy+eD+wogJ973PtMqSQ9HpFeGKk322vATl+A
	YyZBTHAuUxFTP8+P6NPB3db2JjEeAkKo7yNnejze3pRxDshZ8RWU7Ri5BFUpIH4b
	EtvEQSAM0/6QcB15dYejY+jRnAzn4YAxi8BGM1HW8V8kPN32uPF0jB57iVN6pOH+
	igN5c605qlIWwkKCkv5Y1PSHU1WqVOo4tYvzvwIDAQABAoIBAB6DT5MTjnmJkvR4
	VKM7RF46vSgb0paGu77u9YvunbZayDwBuCPRp7nI58kJL+LNCVWsCewTRZ/DwNJO
	3ZPd1cnF58IJvpzeOQ7biLDSpQiQ7zi/Pd874AnuyN5ndKAHvyMiCFCw9HTXITTz
	kF9lhL5PRUibgChgeZ0unP1E6x3861MLkHNT6QaryYDIum5BSmlr8c3FGRrwy/Vh
	xqu2CjihnvsfR0rFIivjIcyu2Nc80Z1dxHRLn6HaKDhRS3H1toDRs2k7Gypcn1T/
	5VhDUxfmdXi0fRxCqgujCjBfzU8XIUwaXDWvrNgVBpZzvawLLkoCXD32w5STf3mN
	+YreWFECgYEA0wGy25IpPBwnV1OVm+cx1ek+1j1jK6MjKKW1g88G3UFcvrbfwPp8
	w/9a5rrcsaesyRUplg00693UavkNejEN5kVyOPTnQSKsAjEZKvhFbhlKJSbd7mI6
	XtyHvOFxHN7drEhnDfhnv6Lmdl84U0dKHXHW3hpvOWyglnOwU24XrqcCgYEAwXyi
	cipswF+x60J8HsFL0rp3G6bYOK1w1i6j+CimwnnNc93/Zyvy9bI/rbr9dmeL6vkx
	U1kSN1BfszMMSUgpzOoNNCXnusmecC45NFigXK44gJWbXg8SnF3pbMqrdQYe0yJH
	xSRx11eraWk6d80wkutdSql5FZDQeaylHrvZjSkCgYBNS1cAdccjHNEQfS7VwFgS
	GvIIus4EIyty7VjTotfJ3vKhbg4C+/8OMRFUaekELv7tXhIwxD+5kzYHxZm5RBTq
	YFAaKyQ3SQMnfmLQycdtMFhnyZgMhfEJ5NDYUOEFiWI8hP0zMeaXXY38mJ7iar/s
	j1H9ZEOQDom480IJHdARCwKBgQCnmQ4dViJdgIbdAfoBqe3D20DQ8vTYI2RaIoxl
	qNDKJDeS5fkVpTwPFxeRpSdoKMlKABB54e2lfDlpQA6ka53JofNN+IeBhKN5CJEE
	hsBhh4fp2a17DBNS3wYS9nsOIv8QY+FHIwp5Pbd812mN3G721Gk7AxGfXJE41BMz
	TARmAQKBgDWm1/3Ms8nTrVYm2DUHD+zViK/A9fc7Xa61npuCTjEtwoyhonuUaEOy
	cEj8XGG34s7ZhyXMMqWgD0GBnIbf4rHIGSo5t7j12xzyLbtCWFVkTxaeHCthPXL5
	nxMgwBTLJEN892ZOKfFz8wC7AQ3uLjTeG2kHVBHLjRKspfZtdUdB
	-----END PRIVATE KEY-----
	`

	p2 := `-----BEGIN PUBLIC KEY-----
	MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAn3sCj1h0ZS8muwog8WaA
	h99nLk3CbnqqrGjeMDoED+9PveQdRH4U74l0RQbmLFJ8EFxvzU9vBkvBpl21X3mc
	5CqULivoiKSYPvnYcMQdtcQDOCJFbNWKJzdS+ceXIaIb3aSQhXeFIWm8qy+eD+wo
	gJ973PtMqSQ9HpFeGKk322vATl+AYyZBTHAuUxFTP8+P6NPB3db2JjEeAkKo7yNn
	ejze3pRxDshZ8RWU7Ri5BFUpIH4bEtvEQSAM0/6QcB15dYejY+jRnAzn4YAxi8BG
	M1HW8V8kPN32uPF0jB57iVN6pOH+igN5c605qlIWwkKCkv5Y1PSHU1WqVOo4tYvz
	vwIDAQAB
	-----END PUBLIC KEY-----
	`

	return &Repository{
		Auth:     auth.NewAuthRepo(logger, sqlite, sqliteTimeout),
		Products: products.NewProductsRepo(logger, sqlite, sqliteTimeout),
		Order:    order.NewOrderRepo(logger, db, dbTimeout),
		Token:    token.NewTokenRepo(logger, sqlite, sqliteTimeout, []byte(p1), []byte(p2)),
	}
}
