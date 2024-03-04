package viacep

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/felipezschornack/golang-api-temperature/internal/erro"
	"github.com/felipezschornack/golang-api-temperature/internal/util"
)

// For more information, please see https://viacep.com.br
type ViaCEP struct {
	Cep         string `json:"cep,omitempty"`
	Logradouro  string `json:"logradouro,omitempty"`
	Complemento string `json:"complemento,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf,omitempty"`
	Ibge        string `json:"ibge,omitempty"`
	Gia         string `json:"gia,omitempty"`
	Ddd         string `json:"ddd,omitempty"`
	Siafi       string `json:"siafi,omitempty"`
}

type ErroViaCEP struct {
	Erro string `json:"erro"`
}

func BuscaCep(zipcode string) (*ViaCEP, *erro.Erro) {
	if zipcode == "" {
		return nil, erro.New(http.StatusBadRequest, "invalid zipcode")
	}

	zipcode, err := util.FormatZipCode(zipcode)
	if err != nil {
		return nil, erro.New(http.StatusBadRequest, "invalid zipcode")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", url.QueryEscape(zipcode)))
	if err != nil {
		return nil, erro.New(http.StatusBadRequest, err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, erro.New(http.StatusInternalServerError, err.Error())
		}

		var data ViaCEP
		err = json.Unmarshal(body, &data)
		if err == nil {
			if data.Localidade != "" {
				//util.PrintDataAsJson("ViaCEP API", data)
				return &data, nil
			} else {
				return nil, erro.New(http.StatusUnprocessableEntity, "invalid zipcode")
			}
		} else {
			var erroViaCep ErroViaCEP // se for codigo de erro da API (pode ser 200 e erro == true)
			err = json.Unmarshal(body, &erroViaCep)
			if err != nil {
				return nil, erro.New(http.StatusInternalServerError, err.Error())
			} else {
				return nil, erro.New(http.StatusUnprocessableEntity, "invalid zipcode")
			}
		}
	} else if resp.StatusCode == 404 {
		return nil, erro.New(http.StatusNotFound, "can not find zipcode")
	} else {
		return nil, erro.New(http.StatusBadRequest, "erro ao fazer requisicao")
	}
}
