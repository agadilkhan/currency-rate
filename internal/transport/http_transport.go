package transport

import (
	"encoding/json"
	"fmt"
	"github.com/agadilkhan/currency-rate/internal/entity"
	"net/http"
)

type HttpTransport struct {
	Host string
}

func New(host string) *HttpTransport {
	return &HttpTransport{
		host,
	}
}

func (t *HttpTransport) GetCurrencies() (*[]entity.Currency, error) {
	resp, err := http.Get(t.Host)
	if err != nil {
		return nil, fmt.Errorf("failed to Get err: %v", err)
	}

	defer resp.Body.Close()

	data := struct {
		Rates map[string]float64 `json:"rates"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to Decode err: %v", err)
	}

	rates := data.Rates

	var res = make([]entity.Currency, 0, len(rates))
	for key, val := range rates {
		res = append(res, entity.Currency{
			Code: key,
			Rate: val,
		})
	}

	return &res, nil
}
