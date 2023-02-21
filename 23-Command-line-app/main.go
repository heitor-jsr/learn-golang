package main

import (
	"command-line/app"
	"log"
	"os"
)

// para rodar a aplicação, vc precisa executar o metodo Run dela, com os argumentos de os.Args, que vai dizer ao go para reconhecer os comandos do seu sistema operacional pela linha de comando

func main() {
	application := app.Gerar()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
