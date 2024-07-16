# Go-C CSV Processor

Biblioteca para processar arquivos CSV aplicando filtros e selecionando colunas conforme especificado. Ela é implementada em Go e pode ser utilizada por uma aplicação em C.

## Índice

1. [Instalação](#instalação)
2. [Uso](#uso)
    1. [Funções Exportadas](#funções-exportadas)
    2. [Exemplo de Uso](#exemplo-de-uso)
3. [Estrutura do Projeto](#estrutura-do-projeto)

## Instalação

### Pré-requisitos

Certifique-se de ter os seguintes softwares instalados em seu sistema:

1. **Go**: Você pode baixar e instalar a partir do [site oficial do Go](https://golang.org/dl/).
2. **GCC**: O compilador GCC é necessário para compilar a aplicação C. Em sistemas baseados em Debian, você pode instalar com:

    ```bash
    sudo apt update
    sudo apt install build-essential
    ```
### Passos de Instalação

1. Clone o repositório
    ```bash
    git clone https://github.com/CarlosDanielss/go-c-csv-processor.git
    ```

2. Compile a biblioteca compartilhada

    ```bash
    go build -o libcsv.so -buildmode=c-shared main.go csv_processor.go
    ```

3. Copie os arquivos gerados para a pasta do projeto onde será utilizado

## Uso

### Funções Exportadas

A biblioteca exporta as seguintes funções para serem usadas em C:

```c
/**
 * Processa os dados do CSV aplicando filtros e selecionando colunas.
 *
 * @param csv Os dados do CSV a serem processados.
 * @param selectedColumns As colunas a serem selecionadas dos dados CSV.
 * @param rowFilterDefinitions Os filtros a serem aplicados aos dados CSV.
 *
 * @return void
 */
void processCsv(const char[], const char[], const char[]);

/**
 * Processa o arquivo CSV aplicando filtros e selecionando colunas.
 *
 * @param csvFilePath O caminho para o arquivo CSV a ser processado.
 * @param selectedColumns As colunas a serem selecionadas dos dados CSV.
 * @param rowFilterDefinitions Os filtros a serem aplicados aos dados CSV.
 *
 * @return void
 */
void processCsvFile(const char[], const char[], const char[]);
```
### Exemplo de Uso
```c
const char csv[] = "header1,header2,header3\n1,2,3\n4,5,6\n7,8,9";
processCsv(csv, "header1,header3", "header1>1\nheader3<8");

// saída
// header1,header3
// 4,6

const char csv_file[] = "path/to/csv_file.csv";
processCsvFile(csv_file, "header1,header3", "header1>1\nheader3<8");

// saída
// header1,header3
// 4,6
```

## Estrutura do Projeto
```plaintext
csv-processor/
│
├── main.go            // Implementação da biblioteca Go
└── csv_processor.go  // Lógica de processamento do CSV
```
