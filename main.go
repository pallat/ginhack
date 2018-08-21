package main

import (
	"log"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func main() {
	r := gin.Default()
	r.GET("/ping", Compose(handler, uuid.NewV4().String()))
	r.Run() // listen and serve on 0.0.0.0:8080
}

func handler(c *Context) {
	c.JSON(200, gin.H{
		"id": c.UUID(),
	})
}

func Compose(h HandlerFunc, id string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer = &Writer{ResponseWriter: c.Writer, id: id}
		c.Set("id", id)

		ctx := &Context{c}

		log.Println(id)
		h(ctx)
	}
}

type HandlerFunc func(*Context)
type Context struct {
	*gin.Context
}

func (c *Context) UUID() string {
	if uid, ok := c.Get("id"); ok {
		return uid.(string)
	}
	return ""
}

type Writer struct {
	id string
	gin.ResponseWriter
}

func (w *Writer) Write(b []byte) (int, error) {
	defer log.Println(w.id, string(b))
	return w.ResponseWriter.Write(b)
}
