package cmd

import (
	"fmt"

	stress_test "github.com/gilbertom/go-cli-stress-test/cmd/stress-test"
	"github.com/spf13/cobra"
)

// stressCmd represents the stress command
var stressCmd = &cobra.Command{
	Use:   "run",
	Short: "Utilize este comando para executar o stress test de uma API",
	Long: `Este comando é utilizado para executar o stress test de uma API.
	O usuário deve fornecer a URL da API, a quantidade de requisições que deseja realizar e a 
	quantidade de requisições simultâneas.
	
	Ao final do teste, o comando exibirá:
	- O tempo total de execução
	- A quantidade de requisições realizadas
	- A quantidade de requisições com status HTTP 200	
	- Distribuição de outros status HTTP (como 429, 500, etc).`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Iniciando stress test...")
		cmd.Println("Parâmetros recebidos:")
		cmd.Println("  URL        :", url)
		cmd.Println("  Requests   :", requests)
		cmd.Println("  Concurrency:", concurrency)
		cmd.Println("")

		stress_test.Run(url, requests, concurrency)
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if url == "" {
			return fmt.Errorf("O valor de url não pode ser vazio")
		}

		if requests <= 0 {
			return fmt.Errorf("O valor de requests deve ser maior que zero")
		}

		if concurrency <= 0 {
			return fmt.Errorf("O valor de concurrency deve ser maior que zero")
		}

		return nil
	},
}

var (
		url = ""
		requests = 0
		concurrency = 0
)

func init() {
	rootCmd.AddCommand(stressCmd)

	stressCmd.Flags().StringVarP(&url, "url", "u", "", "URL da API a ser testada")
	stressCmd.MarkFlagRequired("url")

	stressCmd.Flags().IntVarP(&requests, "requests", "r", 1, "Quantidade de requisições a serem realizadas")
	stressCmd.MarkFlagRequired("requests")

	stressCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "Quantidade de requisições simultâneas")
	stressCmd.MarkFlagRequired("concurrency")
}
