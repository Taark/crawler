package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taark/crawler/src/models"
	"net/http"
)

type Scan func([]string) []*models.Target

type server struct {
	scan Scan
	port string
}

func New(opts ...Option) *server {
	o := newOptions(opts...)
	return &server{
		scan: o.Scan,
		port: o.Port,
	}
}

func (s *server) crawlerHandler(c *gin.Context) {
	var urls []string
	if err := c.ShouldBindJSON(&urls); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, s.scan(urls))
}

func (s *server) Run() error {

	r := gin.Default()

	r.POST("/crawler", s.crawlerHandler)
	return r.Run(fmt.Sprintf(":%s", s.port))
}
