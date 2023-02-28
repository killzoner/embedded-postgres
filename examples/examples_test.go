package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	"database/sql"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/georgysavva/scany/v2/pgxscan"

	prevpgtype "github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	// _ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
)

type (
	Beer struct {
		Id       int                  `db:"id"`
		Name     string               `db:"name"`
		Consumed bool                 `db:"consumed"`
		Rating   float64              `db:"rating"`
		Tags     pgtype.Array[string] `db:"tags"`
	}
	BeerBis struct {
		Id       int                  `db:"id"`
		Name     string               `db:"name"`
		Consumed bool                 `db:"consumed"`
		Rating   float64              `db:"rating"`
		Tags     prevpgtype.TextArray `db:"tags"`
	}
)

func Test_ScanOne(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	t.Cleanup(cancel)

	connString := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	connect := func() (*sql.DB, error) {
		db, err := sql.Open("postgres", connString)
		return db, err
	}

	database := embeddedpostgres.NewDatabase(
		embeddedpostgres.DefaultConfig().Logger(nil),
	)
	err := database.Start()
	require.NoError(t, err)

	t.Cleanup(func() {
		err := database.Stop()
		require.NoError(t, err)
	})

	db, err := connect()
	require.NoError(t, err)

	err = goose.Up(db, "./migrations")
	require.NoError(t, err)

	t.Cleanup(func() {
		err = db.Close()
		require.NoError(t, err)
	})

	time.Sleep(1000)

	t.Run("sqlx", func(t *testing.T) {
		t.Parallel()
		// unsupported Scan, storing driver.Value type string into type *pgtype.FlatArray[string]

		connect := func() (*sqlx.DB, error) {
			db, err := sqlx.Open("pgx/v5", connString)
			return db, err
		}

		db, err := connect()
		require.NoError(t, err)

		rows, err := db.Queryx("SELECT * from beer_catalogue")
		require.NoError(t, err)
		require.NoError(t, rows.Err())

		beers := make([]Beer, 0)
		// beers := make([]BeerBis, 0)
		// pgv5 with pgtype.FlatArray now errors

		for rows.Next() {
			var b Beer
			// var b BeerBis
			err := rows.StructScan(&b)
			require.NoError(t, err)

			beers = append(beers, b)
		}

		fmt.Println(beers[0].Tags)
		require.Equal(t, 1, len(beers))
	})

	t.Run("dbscan", func(t *testing.T) {
		t.Parallel()
		// unsupported Scan, storing driver.Value type string into type *pgtype.FlatArray[string]

		connect := func() (*sql.DB, error) {
			db, err := sql.Open("pgx/v5", connString)
			return db, err
		}

		db, err := connect()
		require.NoError(t, err)

		rows, err := db.Query("SELECT * from beer_catalogue")
		require.NoError(t, err)
		require.NoError(t, rows.Err())

		api, err := dbscan.NewAPI()
		require.NoError(t, err)

		beers := make([]Beer, 0)

		for rows.Next() {
			var b Beer
			err := api.ScanRow(&b, rows)
			require.NoError(t, err)

			beers = append(beers, b)
		}

		fmt.Println(beers)
		require.Equal(t, 1, len(beers))
	})

	t.Run("pgxscan", func(t *testing.T) {
		t.Parallel()

		db, err := pgx.Connect(ctx, connString)
		require.NoError(t, err)

		rows, err := db.Query(ctx, "SELECT * from beer_catalogue")
		require.NoError(t, err)
		require.NoError(t, rows.Err())

		dbscanapi, err := dbscan.NewAPI()
		require.NoError(t, err)

		api, err := pgxscan.NewAPI(dbscanapi)
		require.NoError(t, err)

		beers := make([]Beer, 0)

		for rows.Next() {
			var b Beer
			err := api.ScanRow(&b, rows)
			require.NoError(t, err)

			beers = append(beers, b)
		}

		fmt.Println(beers)
		require.Equal(t, 1, len(beers))
	})
}
