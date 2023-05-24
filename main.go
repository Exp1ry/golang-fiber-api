package main

import (
	"fmt"
	"os"

	"api.ainvest.com/controller/api/routes"
	"api.ainvest.com/controller/db"
	cryptoBrokers "api.ainvest.com/controller/pkg/crypto"
	forexBrokers "api.ainvest.com/controller/pkg/forex"
	nftBroker "api.ainvest.com/controller/pkg/nft"
	"api.ainvest.com/controller/pkg/users"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
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

crypto := api.Group("/crypto")
nft := api.Group("/nft")
stocks := api.Group("/stocks")
admin := api.Group("/admin")
user := api.Group("/users")
// user.Use(middleware.ValidateToken())
// crypto.Use(middleware.ValidateToken())
// nft.Use(middleware.ValidateToken())
// stocks.Use(middleware.ValidateToken())


db, cancel, err := db.ConnectToDB()
if err!= nil {
	fmt.Errorf("Error connecting to DB, %e\n",err)
}

forexCol := db.Collection("forexandstocks")
cryptoCol := db.Collection("cryptos")
nftCol:= db.Collection("nfts")
usersCol := db.Collection("users")
adminCol := db.Collection("admins")

cryptoRepo := cryptoBrokers.NewRepo(cryptoCol)
nftRepo := nftBroker.NewRepo(nftCol)
forexRepo := forexBrokers.NewRepo(forexCol)
usersRepo := users.NewRepo(usersCol, adminCol)

cryptoService := cryptoBrokers.NewService(cryptoRepo)
nftService := nftBroker.NewService(nftRepo)
forexService := forexBrokers.NewService(forexRepo)
usersService := users.NewService(usersRepo)

routes.CryptoRoutes(crypto, cryptoService)
routes.NftRoutes(nft, nftService)
routes.ForexRoutes(stocks, forexService)
routes.UserRoutes(user, usersService)
routes.AdminRoutes(admin, usersService)


app.Listen(port)


defer cancel()
}