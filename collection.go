package shopify

import (
	"fmt"

	"github.com/ugp-io/shopify/models"
)

type CollectionServiceOp struct {
	client *Client
}

type CollectionService interface {
	CreateCollection(models.CreateCollectionRequest) (*models.CollectionResponse, error)
}

func (s *CollectionServiceOp) CreateCollection(req models.CreateCollectionRequest) (*models.CollectionResponse, error) {

	var sources []map[string]any
	for _, colSource := range req.Collection.Sources {

		var conditions []map[string]any
		for _, conditionFields := range colSource.Conditions {
			for field, condition := range conditionFields {

				switch value := condition.Value.(type) {
				case string:
					conditions = append(conditions, map[string]any{
						field: map[string]any{
							"matchType": condition.MatchType,
							"relation":  condition.Relation,
							"value":     condition.Value,
						}})
				case []string:
					conditions = append(conditions, map[string]any{
						field: map[string]any{
							"matchType": condition.MatchType,
							"relation":  condition.Relation,
							"values":    condition.Value,
						}})
				default:
					fmt.Printf("%T\n", value)
				}

			}
		}

		sources = append(sources, map[string]any{
			"source": map[string]any{
				"title": req.Collection.Title,
				"inclusion": map[string]any{
					"matchType":  colSource.MatchType,
					"conditions": conditions,
				},
			},
		})
	}

	input := map[string]any{
		"title":   req.Collection.Title,
		"sources": sources,
	}

	body := map[string]any{
		"query": `
		mutation CreateCollection($collection: CollectionCreateInput!) {
  collectionCreate(collection: $collection) {
    collection {
      id
      title
      handle
    }
    userErrors {
      field
      message
    }
  }
}
  `,
		"variables": map[string]any{"collection": input},
	}

	// bolB, _ := json.Marshal(body)
	// fmt.Println()
	// fmt.Println(string(bolB))
	// fmt.Println()

	var resp models.CollectionResponse
	if err := s.client.Request("POST", fmt.Sprintf(baseURL, s.client.Store), body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
