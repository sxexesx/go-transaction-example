package service

import (
	"context"

	"github.com/sxexesx/go-transaction-example/repository"
)

type WareHouse struct {
	repo *repository.Repo
}

func NewWareHouse(repo *repository.Repo) WareHouse {
	return WareHouse{repo: repo}
}

func (w *WareHouse) AddNew(ctx context.Context) error {
	i := Item{
		ItemID:      100,
		Name:        "test-item",
		SellerID:    100000,
		Description: "test-item",
	}

	v := Sku{
		SkuID:       1,
		ItemID:      100,
		Name:        "test-sku",
		Description: "test-sku",
	}
	if err := w.repo.AddNewProduct(ctx, toRepoItem(i), toReposSku(v)); err != nil {
		return err
	}
	return nil
}

func toRepoItem(i Item) repository.Item {
	return repository.Item{
		ID:          nil,
		ItemID:      i.ItemID,
		Name:        i.Name,
		SellerID:    i.SellerID,
		Description: i.Description,
	}
}

func toReposSku(s Sku) repository.Sku {
	return repository.Sku{
		ID:          nil,
		SkuID:       s.SkuID,
		ItemID:      s.ItemID,
		Name:        s.Name,
		Description: s.Description,
	}
}
