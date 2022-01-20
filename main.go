package main

import (
	"flag"
	"log"
	"os"
	"vendas/config"
	"vendas/controller"
)

func init() {
	var createFlag bool
	flag.BoolVar(&createFlag, "c", false, "create an yaml config file")
	flag.Parse()

	if createFlag {
		if err := config.NewYaml("config.yaml"); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
}

func main() {
	controller, err := controller.InitializeController("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(controller.ListenAndServe())
}
