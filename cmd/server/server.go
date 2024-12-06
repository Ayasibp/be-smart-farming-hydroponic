package server

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/constant"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv(constant.EnvKeyEnv)

	if env != "prod" {
		err := godotenv.Load()

		if err != nil {
			log.Println("error loading env", err)
			log.Fatalln("error loading env", err)
		}
	}

	handlers, middlewares := prepare()

	srv := gin.Default()

	srv.Use(middleware.CORS())

	routes.Build(srv, handlers, middlewares)

	srv.Static("/docs", "./internal/swaggerui")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := srv.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Println("Error running gin server: ", err)
		log.Fatalln("Error running gin server: ", err)
	}

}

func prepare() (handlers routes.Handlers, middlewares routes.Middlewares) {
	appName := os.Getenv(constant.EnvKeyAppName)
	jwtSecret := os.Getenv(constant.EnvKeyJWTSecret)
	refreshTokenDurationStr := os.Getenv(constant.EnvKeyRefreshTokenDuration)

	accessTokenDurationStr := os.Getenv(constant.EnvKeyAccessTokenDuration)

	refreshTokenDuration, err := strconv.Atoi(refreshTokenDurationStr)

	if err != nil {
		log.Fatalln("error creating handlers and middleware", err)
	}

	accessTokenDuration, err := strconv.Atoi(accessTokenDurationStr)
	if err != nil {
		log.Fatalln("error creating handlers and middlewares", err)
	}

	jwtProvider := tokenprovider.NewJWT(appName, jwtSecret, refreshTokenDuration, accessTokenDuration)

	middlewares = routes.Middlewares{
		Auth: middleware.CreateAuth(jwtProvider),
	}

	// db := dbstore.Get()

	handlers = routes.Handlers{}
	return

}
