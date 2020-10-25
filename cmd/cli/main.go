package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Avimitin/GoExFetcher/internal/browser"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`
		ExFetcher is a fetcher for getting exhentai comic's information.

		Usage:
		
			exfetcher [arguments]

		Exp:

			exfetcher https://exhentai.org/g/1761703/00b9540533/ https://exhentai.org/g/1761748/882d6987a6/

		Help:
			List all the comic link as the arguments, program will fetch their infomation one by one.
		`)
		return
	}
	args := os.Args[1:]
	for _, arg := range args {
		result := browser.Reg(arg)
		if result == "" {
			log.Println("[Input Error]Unknown websites")
			os.Exit(-1)
		}
		GAT := browser.GetGIDAndToken(result)
		if GAT == nil {
			log.Println("[Input Error]Unknown websites")
			os.Exit(-1)
		}
		resp, err := browser.Browser(GAT[0], GAT[1])
		if err != nil {
			log.Printf("[Error]%s", err.Error())
		}
		for _, gmd := range resp.Gmd {
			fmt.Printf("%s", browser.Beautify(&gmd))
		}
	}
}
