package main

import (
	"context"

	"github.com/sxexesx/go-transaction-example/postgre"
	"github.com/sxexesx/go-transaction-example/repository"
	"github.com/sxexesx/go-transaction-example/service"
)

func main() {
	db, err := postgre.NewDB()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	r := repository.NewRepo(db)

	house := service.NewWareHouse(r)
	err = house.AddNew(ctx)
	if err != nil {
		panic(err)
	}
}
