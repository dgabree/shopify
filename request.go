package shopify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = `https://%v.myshopify.com/admin/api/2026-10/graphql.json`

type Client struct {
	accessToken  *string
	expiresAt    *time.Time
	ClientID     string
	ClientSecret string
	Store        string
	Collection   CollectionService
	Product      ProductService
}

func NewClient(clientID, clientSecret, store string) *Client {
	c := &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Store:        store,
	}

	c.Collection = &CollectionServiceOp{client: c}
	c.Product = &ProductServiceOp{client: c}

	return c
}

func (c *Client) Request(method string, url string, body interface{}, v interface{}) error {
	// c.accessToken = new(string)
	// *c.accessToken = "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"
	if c.accessToken == nil || c.expiresAt == nil || time.Now().After(*c.expiresAt) {
		if err := c.GetAccessToken(); err != nil {
			return fmt.Errorf("failed to get access token: %w", err)
		}
	}

	var bodyReader io.Reader
	if body != nil {
		requestJson, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewBuffer(requestJson)
	}

	httpReq, errNewRequest := http.NewRequest(method, url, bodyReader)
	if errNewRequest != nil {
		return errNewRequest
	}
	httpReq.Header.Set("X-Shopify-Access-Token", *c.accessToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	// fmt.Println("Resp:", string(respBody))
	// fmt.Println(url)
	decoder := json.NewDecoder(bytes.NewReader(respBody))
	errDecode := decoder.Decode(&v)
	if errDecode != nil {
		return errDecode
	}

	return nil
}

func (c *Client) GetAccessToken() error {

	body := map[string]string{
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
		"grant_type":    "client_credentials",
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(`https://%v.myshopify.com/admin/oauth/access_token`, c.Store), bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token request returned %d: %s", resp.StatusCode, string(respBody))
	}

	tokenResp := make(map[string]interface{})
	if err := json.Unmarshal(respBody, &tokenResp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if accessToken, ok := tokenResp["access_token"].(string); ok {
		// fmt.Println("accessToken:", accessToken)
		expiration := time.Now().Add(time.Duration(tokenResp["expires_in"].(float64)) * time.Second)
		c.expiresAt = &expiration
		c.accessToken = &accessToken
	}

	return nil
}
