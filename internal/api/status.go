package api

import (
	"context"
	"encoding/json"

	"dsgo/internal/client"
)

type Model struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Balance struct {
	Currency string `json:"currency"`
	Total    string `json:"total"`
	Granted  string `json:"granted"`
	ToppedUp string `json:"topped_up"`
}

type Status struct {
	OK       bool      `json:"ok"`
	Models   []Model   `json:"models"`
	Balances []Balance `json:"balances"`
}

func LoadStatus(ctx context.Context, c *client.Client) (Status, error) {
	models, err := listModels(ctx, c)
	if err != nil {
		return Status{}, err
	}
	ok, balances, err := balance(ctx, c)
	if err != nil {
		return Status{}, err
	}
	return Status{OK: ok, Models: models, Balances: balances}, nil
}

func listModels(ctx context.Context, c *client.Client) ([]Model, error) {
	data, err := c.Get(ctx, "/models")
	if err != nil {
		return nil, err
	}
	var payload struct {
		Data []Model `json:"data"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, err
	}
	if payload.Data == nil {
		payload.Data = []Model{}
	}
	return payload.Data, nil
}

func balance(ctx context.Context, c *client.Client) (bool, []Balance, error) {
	data, err := c.Get(ctx, "/user/balance")
	if err != nil {
		return false, nil, err
	}
	var payload struct {
		IsAvailable bool `json:"is_available"`
		Infos       []struct {
			Currency        string `json:"currency"`
			TotalBalance    string `json:"total_balance"`
			GrantedBalance  string `json:"granted_balance"`
			ToppedUpBalance string `json:"topped_up_balance"`
		} `json:"balance_infos"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return false, nil, err
	}
	out := make([]Balance, 0, len(payload.Infos))
	for _, info := range payload.Infos {
		out = append(out, Balance{
			Currency: info.Currency,
			Total:    info.TotalBalance,
			Granted:  info.GrantedBalance,
			ToppedUp: info.ToppedUpBalance,
		})
	}
	return payload.IsAvailable, out, nil
}
