package shopify

import (
	"fmt"

	"github.com/ugp-io/shopify/models"
)

type ProductServiceOp struct {
	client *Client
}

type ProductService interface {
	CreateProduct(models.CreateProductRequest) (*models.ProductResponse, error)
}

func (s *ProductServiceOp) CreateProduct(req models.CreateProductRequest) (*models.ProductResponse, error) {

	input := map[string]any{
		"title":          req.Product.Title,
		"tags":           req.Product.Tags,
		"vendor":         req.Product.Vendor,
		"productOptions": req.Product.ProductOptions,
		"variants":       req.Product.Variants,
		"files":          req.Product.Files,
	}

	body := map[string]any{
		"query": `
      mutation createProductAsynchronous($productSet: ProductSetInput!, $synchronous: Boolean!) {
        productSet(synchronous: $synchronous, input: $productSet) {
          product {
            id
          }
          productSetOperation {
            id
            status
            userErrors {
              code
              field
              message
            }
          }
          userErrors {
            code
            field
            message
          }
        }
      }
		`,
		"variables": map[string]any{
			"synchronous": true,
			"productSet":  input,
		},
	}

	// bolB, _ := json.Marshal(body)
	// fmt.Println()
	// fmt.Println(string(bolB))
	// fmt.Println()

	var resp models.ProductResponse
	if err := s.client.Request("POST", fmt.Sprintf(baseURL, s.client.Store), body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
