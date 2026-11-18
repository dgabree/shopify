package shopify

import (
	"github.com/ugp-io/shopify/models"
)

type ThemeServiceOp struct {
	client *Client
}

type ThemeService interface {
	CreateTheme(models.CreateThemeRequest) (*models.ThemeResponse, error)
}

func (s *ThemeServiceOp) CreateTheme(req models.CreateThemeRequest) (*models.ThemeResponse, error) {

	var resp models.ThemeResponse

	return &resp, nil
}
