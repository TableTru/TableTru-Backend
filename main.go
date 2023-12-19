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
	
	userRepository := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepository)
    userController := controller.NewUserController(userService)
    userRoute := routes.NewUserRoute(userController, router)
    userRoute.Setup()
    db.DB.AutoMigrate(&models.User{})

    CategoryRepository := repository.NewCategoryRepository(db)
    CategoryService := service.NewCategoryService(CategoryRepository)
    CategoryController := controller.NewCategoryController(CategoryService)
    CategoryRoute := routes.NewCategoryRoute(CategoryController, router)
    CategoryRoute.Setup()
    db.DB.AutoMigrate(&models.Category{})

    StoreRepository := repository.NewStoreRepository(db)
    StoreService := service.NewStoreService(StoreRepository)
    StoreController := controller.NewStoreController(StoreService)
    StoreRoute := routes.NewStoreRoute(StoreController, router)
    StoreRoute.Setup()
    db.DB.AutoMigrate(&models.Store{})

    
    router.Gin.Run(":8000") //server started on 8000 port
}