package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/wduartebr/race_request/internal/entity"
)

/*
Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://cdn.apicep.com/file/apicep/" + cep + ".json

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.
*/
func main() {
	var viacep entity.ViaCep
	var cndcpe entity.CdnApiCep

	urlviaCep := viacep.GetUrl("24800181")
	urlcdnCep := cndcpe.GetUrl("24800-181")

	var chv = make(chan string)
	var chc = make(chan string)

	go func() {
		chv <- requestApi(urlviaCep)
		chc <- requestApi(urlcdnCep)

	}()

	/* 	go func() {
		chc <- requestApi(urlcdnCep)

	}() */

	select {
	case ret := <-chv:
		fmt.Printf("Url: %s \n Body: %s", urlviaCep, ret)

	case ret := <-chc:
		fmt.Printf("Url: %s \n Body: %s", urlcdnCep, ret)

	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}

}

func requestApi(url string) string {
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	resp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	return string(resp)

}
