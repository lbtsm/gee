package main

import (
	"log"
	"net/http"

	"github.com/lbtsm/gee"
)

func main() {
	g := gee.New()
	g.Get("/", func(c *gee.Context) {
		c.Status(http.StatusOK)
		_, _ = c.String("get ok")
	})
	g.Post("/hello", func(c *gee.Context) {
		c.Status(http.StatusOK)
		_, _ = c.String("post ok")
	})
	group := g.Group("v1")
	group.Get("/api/hello", func(ctx *gee.Context) {
		_, _ = ctx.String("hello")
	})

	err := g.Run(":8081")
	if err != nil {
		log.Println("run err", err)
	}
}
