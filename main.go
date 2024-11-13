package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/boil"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	_ "github.com/davecgh/go-spew/spew"

	"github.com/mamoruuji/dynamic_novel_api/config"
	"github.com/mamoruuji/dynamic_novel_api/handler"
)

func NullStringToEmptyString(str null.String) string {
	if !str.Valid {
		return ""
	}
	return str.String
}

var db boil.ContextExecutor

func main() {
	db, err := sql.Open("postgres", "host=db port=5432 user=postgres password=pass dbname=dynamic_novel sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	config.SetDB(db)

	mux := handler.Server()

	err = http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != nil {
		log.Fatalf("failed to listen(tcp, :8080)")
	}

}
