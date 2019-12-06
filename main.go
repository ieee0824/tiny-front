package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ieee0824/getenv"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := flag.String("p", getenv.String("TF_PORT", "8080"), "-p \"any port\" or set TF_PORT")
	dir := flag.String("d", getenv.String("TF_STATIC_DIR", "./static"), "-d \"any dir\" or set TF_STATIC_DIR")
	flag.Parse()
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{fmt.Sprintf("http://localhost:%s", *port)}

	r.Use(cors.New(config))

	r.StaticFS("/", http.Dir(*dir))

	if err := r.Run(fmt.Sprintf(":%s", *port)); err != nil {
		log.Fatalln(err)
	}
}
