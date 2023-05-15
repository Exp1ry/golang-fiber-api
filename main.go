package main

import (
	"fmt"
	"os"

	"api.ainvest.com/controller/api/routes"
	"api.ainvest.com/controller/db"
	cryptoBrokers "api.ainvest.com/controller/pkg/crypto"
	forexBrokers "api.ainvest.com/controller/pkg/forex"
	nftBroker "api.ainvest.com/controller/pkg/nft"
	stockBroker "api.ainvest.com/controller/pkg/stocks"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/logger"

)

func HandleErr(err error) error{
if err!= nil {
	return err
}
return nil
}


func main(){

err := godotenv.Load()
if err!= nil {
	panic(err)
}

port := os.Getenv("PORT")

app := fiber.New()
app.Use(cors.New())
app.Use(logger.New())
api := app.Group("/api/v1")

db, cancel, err := db.ConnectToDB()
if err!= nil {
	fmt.Errorf("Error connecting to DB, %e\n",err)
}

forexCol := db.Collection("forexandstocks")
stocksCol := db.Collection("stocksBrokers")
cryptoCol := db.Collection("cryptoBrokers")
nftCol:= db.Collection("nftBrokers")


cryptoRepo := cryptoBrokers.NewRepo(cryptoCol)
nftRepo := nftBroker.NewRepo(nftCol)
forexRepo := forexBrokers.NewRepo(forexCol)
stocksRepo := stockBroker.NewRepo(stocksCol)

cryptoService := cryptoBrokers.NewService(cryptoRepo)
nftService := nftBroker.NewService(nftRepo)
forexService := forexBrokers.NewService(forexRepo)
stockService := stockBroker.NewService(stocksRepo)


routes.CryptoRoutes(api, cryptoService)
routes.NftRoutes(api, nftService)
routes.ForexRoutes(api, forexService)
routes.StockRoutes(api, stockService)


app.Listen(port)


defer cancel()
}