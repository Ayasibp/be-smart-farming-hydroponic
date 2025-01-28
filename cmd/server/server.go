package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/constant"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/handler"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/middleware"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/routes"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	dbstore "github.com/Ayasibp/be-smart-farming-hydroponic/internal/store/db"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/hasher"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/tokenprovider"
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

	db := dbstore.Get()

	hasher := hasher.NewBcrypt(10)

	accountRepo := repository.NewAuthRepository(db)
	profileRepo := repository.NewProfileRepository(db)

	accountService := service.NewAccountService(service.AccountServiceConfig{
		AccountRepo: accountRepo,
		ProfileRepo: profileRepo,
		Hasher:      hasher,
	})

	profileService := service.NewProfileService(service.ProfileServiceConfig{
		ProfileRepo:profileRepo,
	})


	accountHandler := handler.NewAccountHandler(handler.AccountHandlerConfig{
		AccountService: accountService,
	})
	profileHandler:= handler.NewProfileHandler(handler.ProfileHandlerConfig{
		ProfileService: profileService,
	})

	handlers = routes.Handlers{
		Account: accountHandler,
		Profile: profileHandler,
	}
	return
}
