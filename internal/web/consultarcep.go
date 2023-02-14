package web

import (
	"io"
	"net/http"
)

func Consultar(cep string, icep InterfaceCep) ([]byte, error) {
	url := icep.GetUrl(cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// err = json.NewDecoder(resp.Body).Decode(icep)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
