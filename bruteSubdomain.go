package main

// Author: @anakein
import (
	"fmt"
	"net"
)

const (
	domain string = ".uber.com"
)

var Encontrados []string

func main() {
	func() {
		listaPalavras := []string{"www", "admin", "help", "vpn", "anakein", "love", "people"}
		for _, subs := range listaPalavras {
			subdomains := subs
			meudomain := subdomains + domain
			cname, _ := net.LookupCNAME(meudomain)
			if len(cname) > 0 {
				Encontrados = append(Encontrados, cname)
				fmt.Println(meudomain)
			}
		}
		fmt.Println("")
		fmt.Println("Encontrados: ", len(Encontrados))
		/// --------------
		// add depois bypass WAF
		// add depois time sleep
		// add parse wordlist /// ler uma wordlist e passar como subdomain
		// add menu
		// add a generate of wordlist based in source from site.
	}()
}
