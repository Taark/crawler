package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taark/crawler/src/models"
	"net/http"
)

/*
тип Scan это функция
на вход принимает массив строк
возвращает массив указателей на models.Target
*/
type Scan func([]string) []*models.Target

//типа server это структура, фактически это класс
type server struct {
	scan Scan
	port string
}

/*
функция-инициализатор
на вход принимает неограниченное кол-во функций-модификаторов
в файле main.go в New() передается порт и реализация функции Scan
*/
func New(opts ...Option) *server {
	o := newOptions(opts...)
	return &server{
		scan: o.Scan,
		port: o.Port,
	}
}

//обработчик запроса
func (s *server) crawlerHandler(c *gin.Context) {
	var urls []string
	//перевод json в массив строк
	if err := c.ShouldBindJSON(&urls); err != nil {
		// если не удалось распарсить json - возвращается ошибка и 400 ощибка
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, s.scan(urls))
}

// запуск сервера
func (s *server) Run() error {
	r := gin.Default()

	//установка роута. Метод: POST, path: /crawler
	r.POST("/crawler", s.crawlerHandler)
	return r.Run(fmt.Sprintf(":%s", s.port))
}
