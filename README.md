# Stress Test

<p align="center">
  <img src="https://blog.golang.org/gopher/gopher.png" alt="">
</p>


Stress Test é uma aplicação desenvolvida em Go, cujo objetivo é realizar testes de carga em um serviço web. 

Esta aplicação é um CLI que recebe três parâmetros, conforme descrito a seguir.

## Parâmetros

É obrigatório enviar três parâmetros para realizar o teste de carga:

- **--url**
  - URL a receber a carga através do método GET
  
- **--requests**
  - Número inteiro positivo que representa a quantidade de chamadas a realizar.
  
- **--concurrency**
  - Número inteiro positivo que indica a quantidade de processos paralelos a serem utilizados no processamento para realizar chamadas simultâneas.
  
## Índice

- [Instalação](#instalação)
- [Como Executar](#como-executar)
- [Informações geradas](#informações-geradas)
- [Contato](#contato)
- [Agradecimento](#agradecimento)

## Instalação

Para clonar o Stress Test:

```sh
git clone https://github.com/gilbertom/go-cli-stress-test
cd go-cli-stress-test
```

## Como Executar

Criar imagem Docker:
```sh
docker build -t stress-test-img .
```

Criar e Executar Container

Exemplo de teste de carga no site http://google.com.br com 200 chamadas distribuídas em 10 processos paralelos.

```sh
docker run --rm stress-test-img ./stress-test-bin run \
--url=http://google.com.br \
--requests=200 \
--concurrency=10
```

## Informações geradas

  Exemplo de informações geradas em caso de sucesso:
  ```
  Iniciando stress test...
  Parâmetros recebidos:
    URL        : http://google.com.br
    Requests   : 200
    Concurrency: 10

  Tempo total gasto na execução                : 13.903783319s
  Quantidade total de requests realizados      : 200
  Quantidade de requests com status HTTP 200   : 127
  Distribuição de outros códigos de status HTTP:
    Status 200: 127
    Status 500: 73
  ```

## Contato
Para entrar em contato com o desenvolvedor deste projeto:
[gilbertomakiyama@gmail.com](mailto:gilbertomakiyama@gmail.com)

## Agradecimento
Gostaria de expressar minha imensa gratidão a todo o time do curso de Pós-Graduação em Go Avançado da FullCycle pelo empenho, dedicação e excelência no ensino. Suas contribuições foram fundamentais para o meu desenvolvimento e sucesso. Muito obrigado!