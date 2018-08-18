package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.Writer = &Writer{ResponseWriter: c.Writer, store: make([]byte, 0)}
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func middleware(c *gin.Context) {
	w := &Writer{ResponseWriter: c.Writer, store: make([]byte, 0)}
	c.Writer = w
	c.Next()
	fmt.Println(string(w.store))
}

type Writer struct {
	gin.ResponseWriter
	store []byte
}

func (w *Writer) Write(b []byte) (int, error) {
	w.store = b
	return w.ResponseWriter.Write(b)
}
