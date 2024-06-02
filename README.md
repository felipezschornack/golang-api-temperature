# Sistema de temperatura por CEP
## Objetivo
A partir de um cep informado no formato NNNNNNNN ou NNNNN-NNN, apresentar os dados de temperatura em Celsius, Fahrenheit e Kelvin.

## Executando a aplicação utilizando Docker
Com o [Docker](https://www.docker.com/) instalado em sua estação de trabalho, execute o comando:
```
docker compose up -d
```

## Acessando a aplicação Localmente
Depois de instalar e executar a aplicação via Docker, a aplicação estará disponível para uso [aqui](http://localhost:8080/weather?zipcode=99999999).

## Acessando a aplicação no Google Cloud Run
[API no Google Cloud Run](https://golang-api-temperature-teuygslxla-uc.a.run.app/weather?zipcode=99999999)
