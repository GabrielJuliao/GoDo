package main

import (
	"fmt"
	"log"
	"os"
)

func isArgsValid(args []string) bool {
	for _, arg := range args {
		if arg == "godo" {
			log.Fatal("Cannot call godo inside it self.")
		}
	}
	return true
}

func main() {
	fmt.Printf(
		"\n%s\n\n",
		b64Decoder("4pSM4pSA4pSQ4pSM4pSA4pSQ4pSM4pSs4pSQ4pSM4pSA4pSQDQrilIIg4pSs4pSCIOKUgiDilILilILilIIg4pSCDQrilJTilIDilJjilJTilIDilJjilIDilLTilJjilJTilIDilJggdjAuMC4xDQpodHRwczovL2dpdGh1Yi5jb20vR2FicmllbEp1bGlhby9nb2Rv"),
	)
	appArguments := os.Args[1:]
	if len(appArguments) > 0 && isArgsValid(appArguments) && appArguments[0] != "-h" && appArguments[0] != "--help" {
		execMacro(appArguments)
	} else {
		printUsage()
	}
}
