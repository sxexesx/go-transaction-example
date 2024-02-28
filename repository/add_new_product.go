package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Item struct {
	ID          *int64 `db:"id"`
	ItemID      int64  `db:"item_id"`
	Name        string `db:"name"`
	SellerID    int64  `db:"seller_id"`
	Description string `db:"description"`
}

type Sku struct {
	ID          *int64 `db:"id"`
	SkuID       int64  `db:"sku_id"`
	ItemID      int64  `db:"item_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

func (r *Repo) AddNewProduct(ctx context.Context, item Item, sku Sku) error {
	return r.WithTx(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		return addNewProductTx(ctx, tx, item, sku)
	})
}

func addNewProductTx(ctx context.Context, tx *sqlx.Tx, item Item, sku Sku) error {
	var itemDB Item
	err := tx.GetContext(ctx, &itemDB, insertItemSQL, item.ItemID, item.Name, item.SellerID, item.Description)
	if err != nil {
		return err
	}
	var skuDB Sku
	err = tx.GetContext(ctx, &skuDB, insertSkuSQL, sku.SkuID, sku.ItemID, sku.Name, sku.Description)
	if err != nil {
		return err
	}

	return nil
}

const insertItemSQL = `insert into item(item_id, name, seller_id, description)
values ($1, $2, $3, $4)
returning id, item_id, name, seller_id, description`

const insertSkuSQL = `insert into sku(sku_id, item_id, name, description)
values ($1, $2, $3, $4)
returning id, sku_id, item_id, name, description`
