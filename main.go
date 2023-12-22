package main

import (
	"TableTru/api/controller"
	"TableTru/api/repository"
	"TableTru/api/routes"
	"TableTru/api/service"
	"TableTru/infrastructure"
	"TableTru/models"
)

func init() {
    infrastructure.LoadEnv()
}

func main() {

    router := infrastructure.NewGinRouter() //router has been initialized and configured
    db := infrastructure.NewDatabase() // databse has been initialized and configured
	
    CategoryRepository := repository.NewCategoryRepository(db)
    CategoryService := service.NewCategoryService(CategoryRepository)
    CategoryController := controller.NewCategoryController(CategoryService)
    CategoryRoute := routes.NewCategoryRoute(CategoryController, router)
    CategoryRoute.Setup()
    db.DB.AutoMigrate(&models.Category{})

    LocationRepository := repository.NewLocationRepository(db)
    LocationService := service.NewLocationService(LocationRepository)
    LocationController := controller.NewLocationController(LocationService)
    LocationRoute := routes.NewLocationRoute(LocationController, router)
    LocationRoute.Setup()
    db.DB.AutoMigrate(&models.Store{})

    PromotionRepository := repository.NewPromotionRepository(db)
    PromotionService := service.NewPromotionService(PromotionRepository)
    PromotionController := controller.NewPromotionController(PromotionService)
    PromotionRoute := routes.NewPromotionRoute(PromotionController, router)
    PromotionRoute.Setup()
    db.DB.AutoMigrate(&models.Store{})

    PromotionCodeRepository := repository.NewPromotionCodeRepository(db)
    PromotionCodeService := service.NewPromotionCodeService(PromotionCodeRepository)
    PromotionCodeController := controller.NewPromotionCodeController(PromotionCodeService)
    PromotionCodeRoute := routes.NewPromotionCodeRoute(PromotionCodeController, router)
    PromotionCodeRoute.Setup()
    db.DB.AutoMigrate(&models.Store{})

    RatingRepository := repository.NewRatingRepository(db)
    RatingService := service.NewRatingService(RatingRepository)
    RatingController := controller.NewRatingController(RatingService)
    RatingRoute := routes.NewRatingRoute(RatingController, router)
    RatingRoute.Setup()
    db.DB.AutoMigrate(&models.Store{})

    ReviewRepository := repository.NewReviewRepository(db)
    ReviewService := service.NewReviewService(ReviewRepository)
    ReviewController := controller.NewReviewController(ReviewService)
    ReviewRoute := routes.NewReviewRoute(ReviewController, router)
    ReviewRoute.Setup()
    db.DB.AutoMigrate(&models.Store{})

    ReviewImageRepository := repository.NewReviewImageRepository(db)
    ReviewImageService := service.NewReviewImageService(ReviewImageRepository)
    ReviewImageController := controller.NewReviewImageController(ReviewImageService)
    ReviewImageRoute := routes.NewReviewImageRoute(ReviewImageController, router)
    ReviewImageRoute.Setup()
    db.DB.AutoMigrate(&models.Store{})
    

    StoreRepository := repository.NewStoreRepository(db)
    StoreService := service.NewStoreService(StoreRepository)
    StoreController := controller.NewStoreController(StoreService)
    StoreRoute := routes.NewStoreRoute(StoreController, router)
    StoreRoute.Setup()
    db.DB.AutoMigrate(&models.Store{})

    StoreImageRepository := repository.NewStoreImageRepository(db)
    StoreImageService := service.NewStoreImageService(StoreImageRepository)
    StoreImageController := controller.NewStoreImageController(StoreImageService)
    StoreImageRoute := routes.NewStoreImageRoute(StoreImageController, router)
    StoreImageRoute.Setup()
    db.DB.AutoMigrate(&models.Store{})

    TableBookingRepository := repository.NewTableBookingRepository(db)
    TableBookingService := service.NewTableBookingService(TableBookingRepository)
    TableBookingController := controller.NewTableBookingController(TableBookingService)
    TableBookingRoute := routes.NewTableBookingRoute(TableBookingController, router)
    TableBookingRoute.Setup()
    db.DB.AutoMigrate(&models.Store{})

	userRepository := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepository)
    userController := controller.NewUserController(userService)
    userRoute := routes.NewUserRoute(userController, router)
    userRoute.Setup()
    db.DB.AutoMigrate(&models.User{})

    
    router.Gin.Run(":8000") //server started on 8000 port
}