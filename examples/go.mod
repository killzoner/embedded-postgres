module github.com/fergusstrange/embedded-postgres/examples

go 1.18

replace github.com/fergusstrange/embedded-postgres => ../

require (
	github.com/fergusstrange/embedded-postgres v0.0.0
	github.com/georgysavva/scany/v2 v2.0.0
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jackc/pgx/v5 v5.2.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/pressly/goose/v3 v3.0.1
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gofrs/uuid v4.3.1+incompatible // indirect
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/puddle/v2 v2.1.2 // indirect
	github.com/lib/pq v1.10.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/sync v0.0.0-20220923202941-7f9b1623fab7 // indirect
	golang.org/x/text v0.3.8 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
