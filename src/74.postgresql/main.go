package main
import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Person struct {
	Name string
	Age int
}

func main(){
	Db, _ = sql.Open("postgres", "user= dbname=test_db password=password")
	defer Db.Close()

	//* C

	//* R

	//* U

	//* D
}