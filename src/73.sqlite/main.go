package main
import (
	"database/sql"
	"log"
	//"fmt"
	_"github.com/mattn/go-sqlite3"
)

//database + sqlite3

var Db *sql.DB

type Person struct {
	Name string
	Age int
}

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	//* テーブル作成
	// cmd := `CREATE TABLE IF NOT EXISTS person(
	// 					name STRING,
	// 					age INT)`
	//_, err := Db.Exec(cmd)
	
	//* データの追加
	// cmd := "INSERT INTO person( name, age) VALUES (?, ?)"
	// _, err := Db.Exec(cmd, "tarou", 20)

	//* データの更新
	// cmd := "UPDATE person SET age = ? WHERE name =?"
	// _, err := Db.Exec(cmd, 25, "tarou")

	//* データの取得
	// cmd := "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err := Db.Exec(cmd, "hanako", 19)

	//* 特定のデータの取得
	// cmd := "SELECT * FROM person where age = ?"
	//? //QueryRow 1レコード取得
	// row := Db.QueryRow(cmd, 25)
	// var p Person
	// err := row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	//* 複数データの取得
	// cmd := "SELECT * FROM person"
	//? Queryは条件に合うものを全て取得
	// rows, _ := Db.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next(){
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }

	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	//* データの削除
	cmd := "DELETE FROM person WHERE name = ?"
	_, err := Db.Exec(cmd, "hanako")
	if err != nil {
		log.Fatalln(err)
	}
}