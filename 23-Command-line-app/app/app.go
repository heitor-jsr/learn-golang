package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// o retorno dessa função é o cli.App. o ponteiro indica que deve retornar a aplicação em si.

//vai retronar a app de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search for IPs and server names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}
	// define os comandos da aplicação. ele é do tipo slice do pacote cli.Command.
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search ips from internet addres",
			// flags tb possuem um tipo especifico: são slices de cli.Flags
			Flags:  flags,
			Action: serachIps,
		},
		{
			Name:  "servers",
			Usage: "Search server names from internet addres",
			// flags tb possuem um tipo especifico: são slices de cli.Flags
			Flags:  flags,
			Action: serachServers,
		},
	}
	return app
}

func serachIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func serachServers(c *cli.Context) {
	host := c.String("host")
	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
