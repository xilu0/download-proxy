package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/download", func(c *gin.Context) {
		address, _ := c.GetQuery("address")
		data, err := download(address)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
		}
		c.Data(http.StatusOK, "text/plain", data)
	})
	return r
}

func download(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}

var port int

func init() {
	const (
		defaultPort = 8080
	)
	flag.IntVar(&port, "port", defaultPort, "server port")
	flag.IntVar(&port, "p", defaultPort, "server port shorthand")
}

func main() {
	flag.Parse()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + fmt.Sprint(port))
}
