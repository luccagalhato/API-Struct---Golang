package controller

import (
	"log"
	"net/http"
	"vendas/config"
	"vendas/database"
	"vendas/server"
)

//Linx ...
type linx database.SQLStr

// Controller ...
type Controller struct {
	linx   *linx
	Server *http.Server
}

//ListenAndServe ...
func (c *Controller) ListenAndServe() error {
	return c.Server.ListenAndServe()
}

//InitializeController ...
func InitializeController(filePath string) (*Controller, error) {
	c := &Controller{}

	confChan := make(chan config.Config)

	firstRead := make(chan error)
	go func() {
		conf := <-confChan
		l, err := database.NewSQL(&conf.Linx)
		if err != nil {
			firstRead <- err
			return
		}
		c.linx = (*linx)(l)
		if err := (*database.SQLStr)(c.linx).Connect(); err != nil {
			firstRead <- err
			return
		}
		c.Server = server.NewServer(conf, c.Controllers())

		firstRead <- nil

		for conf := range confChan {
			if err := (*database.SQLStr)(c.linx).UpdateConfig(&conf.Linx); err != nil {
				log.Fatal(err)
			}
			c.Server = server.NewServer(conf, c.Controllers())
		}
	}()
	_, err := config.LoadYaml(filePath, func(conf config.Config) {
		confChan <- conf
	})
	if err != nil {
		return nil, err
	}
	if err := <-firstRead; err != nil {
		return nil, err
	}
	return c, nil
}
