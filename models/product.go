package models

type CreateProductRequest struct {
	Product *Product
}

type ProductResponse struct{}

type Product struct {
	Handle         *string
	Title          *string
	Status         *string
	Vendor         *string
	Tags           *[]string
	Files          *[]ProductFile
	ProductOptions []ProductOption
	Variants       []ProductVariant
	// Category            *string         `json:"category"`
	// ClaimOwnership      *ClaimOwnership `json:"claimOwnership"`
	// Collections         []string        `json:"collections"`
	// CombinedListingRole *string         `json:"combinedListingRole"`
	// DescriptionHtml     string          `json:"descriptionHtml"`
	// GiftCard bool     `json:"giftCard"`
	// Metafields     []string        `json:"metafields"`
	// ProductType    string          `json:"productType"`
	// RedirectNewHandle bool            `json:"redirectNewHandle"`
	// RequiresSellingPlan bool            `json:"requiresSellingPlan"`
	// SEO                 *SEO            `json:"seo"`
	// TemplateSuffix      *string         `json:"templateSuffix"`
}

type ProductOption struct {
	Name     *string              `json:"name"`
	Position *int                 `json:"position"`
	Values   []ProductOptionValue `json:"values"`
}

type ProductVariant struct {
	ID           *int64               `json:"id"`
	Position     *int                 `json:"position"`
	Price        *string              `json:"price"`
	Published    *bool                `json:"published"`
	SKU          *string              `json:"sku"`
	Taxable      *bool                `json:"taxable"`
	OptionValues []ProductOptionValue `json:"optionValues"`
	// RequiresComponents bool                 `json:"requiresComponents"`
	// ShowUnitPrice      bool                 `json:"showUnitPrice"`
	// TaxCode              *string              `json:"taxCode"`
	// UnitPriceMeasurement *string              `json:"unitPriceMeasurement"`
	// Barcode              *string              `json:"barcode"`
	// Barcodes             []string             `json:"barcodes"`
	// CompareAtPrice       string               `json:"compareAtPrice"`
	// File                 *string              `json:"file"`
	// InventoryItem        InventoryItem        `json:"inventoryItem"`
	// InventoryPolicy      string               `json:"inventoryPolicy"`
	// InventoryQuantities  []string             `json:"inventoryQuantities"`
	// Metafields           []string             `json:"metafields"`
}

type ProductFile struct {
	Src *string `json:"originalSource"`
}

type ProductOptionValue struct {
	OptionName *string `json:"optionName,omitempty"`
	Name       *string `json:"name"`
}
