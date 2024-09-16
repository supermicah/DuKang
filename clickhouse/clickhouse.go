package main

import (
	"context"
	"fmt"

	"github.com/uptrace/go-clickhouse/ch"
)

func main() {
	ctx := context.Background()
	db := ch.Connect(
		// clickhouse://<user>:<password>@<host>:<port>/<database>?sslmode=disable
		ch.WithDSN("clickhouse://default:Sjpclickhouse05.@10.10.10.239:9000/default?sslmode=disable"),
	)
	res, err := db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.RowsAffected())

	var num int
	err = db.QueryRowContext(ctx, "SELECT 1").Scan(&num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}
