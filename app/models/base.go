package models

import (
	"app_todo/config"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

var Db *sql.DB

var err error

/*
//定数名を宣言
const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)
*/

//テーブルの作成
func init() {

	url := os.Getenv("DATABASE URL")
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}

	// 	//sqlを開く
	// 	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	// 	//エラーハンドリング
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	//コマンドの作成 テーブルネームのユーザーを代入
	// 	/*ユーザーテーブルがなければテーブルを作成
	// 	AUTOINCREMENT　自動増分
	// 	NOT NULL null値の禁止 / UNIQUE重複の禁止
	// 	created_at DATETIME　型の指定
	// 	*/

	// 	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		uuid STRING NOT NULL UNIQUE,
	// 		name STRING,
	// 		email STRING,
	// 		password STRING,
	// 		created_at DATETIME)`, tableNameUser)

	// 	//コマンドの実施
	// 	Db.Exec(cmdU)

	// 	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		content TEXT,
	// 		user_id INTEGER,
	// 		created_at DATETIME)`, tableNameTodo)

	// 	Db.Exec(cmdT)

	// 	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		uuid STRING NOT NULL UNIQUE,
	// 		email STRING,
	// 		user_id INTEGER,
	// 		created_at DATETIME)`, tableNameSession)

	// 	Db.Exec(cmdS)

}

//uuidの作成　返り値uuidobj
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

//posswordの保存
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
