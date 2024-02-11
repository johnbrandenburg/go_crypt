package go_crypt

import (
	// "fmt"
	"log"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	client := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    os.Getenv("ALPACA_KEY"),
		APISecret: os.Getenv("ALPACA_SECRET"),
		BaseURL:   os.Getenv("ALPACA_URL"),
	})

	r.GET("/account", func(c *gin.Context) {
		acct, err := client.GetAccount()
		if err != nil {
			panic(err)
		}

		c.JSON(200, *acct)
	})
	r.Run(":8181") // listen and serve on 0.0.0.0:8080
}
