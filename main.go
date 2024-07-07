package main

import (
	"fmt"
	"golangestudo/model"
	"time"
)

func main() {
	fmt.Println("Iniciando...")

	endereco := model.Endereco{
		Rua:    "Alcides de Souza",
		Numero: 112,
		Cidade: "Valença",
	}
	fmt.Println(endereco)

	pessoa := model.Pessoa{
		Nome:             "João Pedro",
		Endereco:         endereco,
		DataDeNascimento: time.Date(1998, 8, 10, 0, 0, 0, 0, time.Local),
	}
	fmt.Println(pessoa)

	pessoa.CalculaIdade()
	fmt.Println(pessoa.Idade)

	automovelMoto := model.Automovel{
		Ano:    2022,
		Placa:  "XPTO",
		Modelo: "CG",
	}
	moto := model.Moto{
		Automovel:   automovelMoto,
		Cilindradas: 125,
	}
	fmt.Println(moto)
	fmt.Println(moto.Ano)
}
