package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// WalletClient handles communication with wallet-related APIs.
type WalletClient struct {
	internalAddress string
	externalAddress string
	net             *http.Client
}

// NewWallet creates a new instance of WalletClient.
func NewWallet(internalAddress, externalAddress string) *WalletClient {
	return &WalletClient{
		internalAddress: internalAddress,
		externalAddress: externalAddress,
		net:             &http.Client{Timeout: time.Minute * 10},
	}
}

// SetNet sets a custom HTTP client for the WalletClient.
func (c *WalletClient) SetNet(net *http.Client) *WalletClient {
	c.net = net
	return c
}

// UpdateWalletBalanceRequest represents the request payload for updating wallet balance.
type UpdateWalletBalanceRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

// Error represents an error response from the API.
type Error struct {
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

// DecreaseWalletBalance decreases the wallet balance for a specific user.
func (c *WalletClient) DecreaseWalletBalance(ctx context.Context, req *UpdateWalletBalanceRequest) error {
	return c.callAPI(ctx, "/wallet/decrease", http.MethodPatch, req)
}

// IncreaseWalletBalance increases the wallet balance for a specific user.
func (c *WalletClient) IncreaseWalletBalance(ctx context.Context, req *UpdateWalletBalanceRequest) error {
	return c.callAPI(ctx, "/wallet/increase", http.MethodPatch, req)
}

// callAPI performs a generic API call.
func (c *WalletClient) callAPI(ctx context.Context, url string, method string, body any) error {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s%s", c.externalAddress, url), bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.net.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr Error
		if err = json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return fmt.Errorf("unknown error: %w", err)
		}
		return &apiErr
	}

	return nil
}
