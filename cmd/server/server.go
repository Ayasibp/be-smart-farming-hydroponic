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
	farmRepo := repository.NewFarmRepository(db)
	systemUnitRepo := repository.NewSystemUnitRepository(db)
	growthHistRepo := repository.NewGrowthHistRepository(db)
	systemLogRepo := repository.NewSystemLogRepository(db)
	superAccountRepo := repository.NewSuperAccountRepository(db)
	unitIdRepo := repository.NewUnitIdRepository(db)
	tankTransRepo := repository.NewTankTransRepository(db)

	accountService := service.NewAccountService(service.AccountServiceConfig{
		AccountRepo: accountRepo,
		ProfileRepo: profileRepo,
		Hasher:      hasher,
	})
	superAccountService := service.NewSuperAccountService(service.SuperAccountServiceConfig{
		SuperAccountRepo: superAccountRepo,
		Hasher:           hasher,
	})
	profileService := service.NewProfileService(service.ProfileServiceConfig{
		ProfileRepo: profileRepo,
		AccountRepo: accountRepo,
	})
	farmService := service.NewFarmService(service.FarmServiceConfig{
		FarmRepo:    farmRepo,
		ProfileRepo: profileRepo,
	})
	systemUnitService := service.NewSystemUnitService(service.SystemUnitServiceConfig{
		SystemUnitRepo: systemUnitRepo,
		FarmRepo:       farmRepo,
		UnitKeyRepo:    unitIdRepo,
	})
	growthHistService := service.NewGrowthHistService(service.GrowthHistServiceConfig{
		GrowthHistRepo: growthHistRepo,
		FarmRepo:       farmRepo,
		SystemUnitRepo: systemUnitRepo,
	})
	tankTransService := service.NewTankTransService(service.TankTransServiceConfig{
		TankTransRepo: tankTransRepo,
		FarmRepo:       farmRepo,
		SystemUnitRepo: systemUnitRepo,
	})
	systemLogService := service.NewSystemLogService(service.SystemLogServiceConfig{
		SystemLogRepo: systemLogRepo,
	})
	unitIdService := service.NewUnitIdService(service.UnitIdServiceConfig{
		UnitIdRepo: unitIdRepo,
	})

	accountHandler := handler.NewAccountHandler(handler.AccountHandlerConfig{
		AccountService:   accountService,
		SystemLogService: systemLogService,
	})
	profileHandler := handler.NewProfileHandler(handler.ProfileHandlerConfig{
		ProfileService:   profileService,
		SystemLogService: systemLogService,
	})
	farmHandler := handler.NewFarmHandler(handler.FarmHandlerConfig{
		FarmService:      farmService,
		SystemLogService: systemLogService,
	})
	systemUnitHandler := handler.NewSystemUnitHandler(handler.SystemUnitHandlerConfig{
		SystemUnitService: systemUnitService,
		SystemLogService:  systemLogService,
	})
	growthHistHandler := handler.NewGrowthHistHandler(handler.GrowthHistHandlerConfig{
		GrowthHistService: growthHistService,
		SystemLogService:  systemLogService,
	})
	superAccountHandler := handler.NewSuperAccountHandler(handler.SuperAccountHandlerConfig{
		SuperAccountService: superAccountService,
		SystemLogService:    systemLogService,
	})
	tankTransHandler := handler.NewTankTransHandler(handler.TankTransHandlerConfig{
		TankTransService: tankTransService,
		SystemLogService:  systemLogService,
	})
	unitIdHandler := handler.NewUnitIdHandler(handler.UnitIdHandlerConfig{
		UnitIdService:    unitIdService,
		SystemLogService: systemLogService,
	})

	handlers = routes.Handlers{
		Account:      accountHandler,
		Profile:      profileHandler,
		Farm:         farmHandler,
		SystemUnit:   systemUnitHandler,
		GrowthHist:   growthHistHandler,
		SuperAccount: superAccountHandler,
		UnitId:       unitIdHandler,
		TankTrans:tankTransHandler,
	}
	return
}
