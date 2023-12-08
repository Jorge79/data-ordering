package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Utilização: <arquivo_de_dados.csv>")
		return
	}

	arquivoEntrada := os.Args[1]

	if err := processarArquivo(arquivoEntrada); err != nil {
		fmt.Println("Erro ao processar o arquivo:", err)
	}
}

func processarArquivo(arquivoEntrada string) error {
	registros, cabecalho, err := lerCSV(arquivoEntrada)
	if err != nil {
		return err
	}

	if len(registros) > 0 {
		registros = registros[1:]
	}

	ordenarPorColuna(registros, 0)
	arquivoSaidaNome := "ordenadoNome.csv"
	err = escreverCSV(arquivoSaidaNome, registros, cabecalho)
	if err != nil {
		return err
	}

	ordenarPorColuna(registros, 1)
	arquivoSaidaIdade := "ordenadoIdade.csv"
	err = escreverCSV(arquivoSaidaIdade, registros, cabecalho)
	if err != nil {
		return err
	}

	ordenarPorColuna(registros, 2)
	arquivoSaidaPontos := "ordenadoPontos.csv"
	err = escreverCSV(arquivoSaidaPontos, registros, cabecalho)
	if err != nil {
		return err
	}

	fmt.Printf("Registros ordenados por nome, idade e pontos. Verifique os arquivos: %s, %s e %s\n", arquivoSaidaNome, arquivoSaidaIdade, arquivoSaidaPontos)

	return nil
}

func lerCSV(nomeArquivo string) ([][]string, []string, error) {
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		return nil, nil, err
	}
	defer arquivo.Close()

	leitorCSV := csv.NewReader(arquivo)

	cabecalho, err := leitorCSV.Read()
	if err != nil {
		return nil, nil, err
	}

	registros, err := leitorCSV.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	return registros, cabecalho, nil
}

func escreverCSV(nomeArquivo string, registros [][]string, cabecalho []string) error {
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	escritorCSV := csv.NewWriter(arquivo)
	defer escritorCSV.Flush()

	if err := escritorCSV.Write(cabecalho); err != nil {
		return err
	}

	err = escritorCSV.WriteAll(registros)
	if err != nil {
		return err
	}

	return nil
}

func ordenarPorColuna(registros [][]string, coluna int) {
	sort.Slice(registros, func(i, j int) bool {
		switch coluna {
		case 0:
			return registros[i][coluna] < registros[j][coluna]
		case 1:
			idadeI, _ := strconv.Atoi(registros[i][coluna])
			idadeJ, _ := strconv.Atoi(registros[j][coluna])
			return idadeI < idadeJ
		case 2:
			pontosI, _ := strconv.Atoi(registros[i][coluna])
			pontosJ, _ := strconv.Atoi(registros[j][coluna])
			return pontosI < pontosJ
		default:
			return false
		}
	})
}
