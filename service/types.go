package service

type Item struct {
	ItemID      int64
	Name        string
	SellerID    int64
	Description string
}

type Sku struct {
	SkuID       int64
	ItemID      int64
	Name        string
	Description string
}
