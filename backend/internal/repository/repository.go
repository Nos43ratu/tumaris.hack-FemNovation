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
	-----END RSA PRIVATE KEY-----
	`

	p2 := `-----BEGIN PUBLIC KEY-----
	MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA1zcCymk9xFUGPUMITzUM
	tBohhjqPS5LtaB6Jg6KjwLig2BWCHIJyYysjmBVNVY5weYTJx+4E5wgMZfdjKr0y
	9T+qhmZJUh4vU52IKbytYGl0FOB8uzbsnnQyygwc80OlNi5kpoUffzNptMPbMCEH
	4pZ5lRDy4dH+NWUEFs9soVkw53HQOMBGoWxMOQyb6Jr/BebJhZ+bwg9yLLP+A8A3
	Y0NRRK1QzomuRrqXhtDeN8hdtuWMpv34bKvwUcvh2l4nOvsU9vTuhWRg+m3cPagh
	b50veCxJoYDsuHyn/Csg7EoUv19NSZz5ezCqbVDQr30HqXSg0MH0EobVK0mFluc5
	zlc2hrsAJ+9K7faHAnr/f8jcyt15p9QMHj1+9OL+XbpKRw++pfnQTGyHIRCWeVOG
	bmuAZjq6hln/ZaxWlx5CEFZjhKRpt8ATDN4ecCF6b/p3SUevo5YkHSY2VpF1aZ6X
	HGReB6p2TBFV2+TP1pfk1TYJ9u4gZWFPtXmVtDMoxpZ5fN1+LkAS1bfPZXveYaBV
	fxmbV9ZYfWc9lrwe1JLaOqx9NsKDZbDStqeniJUJAAQV53Dx2zjs4HeHwWP2uAeR
	64SrMQBx/2y5HKTC42yK05r7FILYR0S3jmkcgBwQhSGVs4MW+Zz4jg2n9yOyzlfI
	tnJB1XL1MPJB5nnAw0bbmpcCAwEAAQ==
	-----END PUBLIC KEY-----
	`

	return &Repository{
		Auth:     auth.NewAuthRepo(logger, sqlite, sqliteTimeout),
		Products: products.NewProductsRepo(logger, sqlite, sqliteTimeout),
		Order:    order.NewOrderRepo(logger, db, dbTimeout),
		Token:    token.NewTokenRepo(logger, sqlite, sqliteTimeout, []byte(p1), []byte(p2)),
	}
}
