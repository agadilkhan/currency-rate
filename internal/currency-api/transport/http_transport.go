package transport

import (
	"encoding/json"
	"fmt"
	"github.com/agadilkhan/currency-rate/internal/currency-api/entity"
	"net/http"
)

type HttpTransport struct {
	Host string
}

func (t *HttpTransport) GetCurrencies() (*[]entity.Currency, error) {
	resp, err := http.Get(t.Host)
	if err != nil {
		return nil, fmt.Errorf("failed to Get err: %v", err)
	}

	defer resp.Body.Close()

	var data map[string]any
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to Decode err: %v", err)
	}

	rates, ok := data["rates"].([]entity.Currency)
	if !ok {
		return nil, fmt.Errorf("cannot convert to arr of Currencies")
	}

	return &rates, nil
}
