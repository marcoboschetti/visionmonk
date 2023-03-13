package entities

type CatalogProduct struct {
	entityMetadata

	Title           string `json:"title"`
	Description     string `json:"description"`
	Barcode         string `json:"barcode"`
	ImageBase64     string `json:"image_base_64"`
	IsGlobalCatalog bool   `json:"is_global_catalog"`
	ShopID          string `json:"shop_id" pg:"product_shop_id"`

	Category string `json:"category"`
	Brand    string `json:"brand"`
	Color    string `json:"color"`
	Size     string `json:"size"`
}

func (c *CatalogProduct) Equals(o CatalogProduct) bool {
	return c.Title == o.Title &&
		c.Description == o.Description &&
		c.Barcode == o.Barcode &&
		c.ImageBase64 == o.ImageBase64 &&
		c.Category == o.Category &&
		c.Brand == o.Brand &&
		c.Color == o.Color &&
		c.Size == o.Size
}

type ShopProduct struct {
	shopEntityMetadata

	CatalogProductID string          `json:"catalog_product_id"`
	CatalogProduct   *CatalogProduct `json:"catalog_product"`
	SKU              string          `json:"sku"`
	Inventory        int             `json:"inventory"`
	PriceCts         int             `json:"price_cts"`
}
