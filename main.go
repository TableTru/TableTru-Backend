package main

import (
	"TableTru/api/controller"
	"TableTru/api/repository"
	"TableTru/api/routes"
	"TableTru/api/service"
	"TableTru/infrastructure"
	"TableTru/models"
	// "time"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {

	router := infrastructure.NewGinRouter() //router has been initialized and configured
	db := infrastructure.NewDatabase()      // databse has been initialized and configured

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
	db.DB.AutoMigrate(&models.Location{})

	PromotionRepository := repository.NewPromotionRepository(db)
	PromotionService := service.NewPromotionService(PromotionRepository)
	PromotionController := controller.NewPromotionController(PromotionService)
	PromotionRoute := routes.NewPromotionRoute(PromotionController, router)
	PromotionRoute.Setup()
	db.DB.AutoMigrate(&models.Promotion{})

	PromotionCodeRepository := repository.NewPromotionCodeRepository(db)
	PromotionCodeService := service.NewPromotionCodeService(PromotionCodeRepository)
	PromotionCodeController := controller.NewPromotionCodeController(PromotionCodeService)
	PromotionCodeRoute := routes.NewPromotionCodeRoute(PromotionCodeController, router)
	PromotionCodeRoute.Setup()
	db.DB.AutoMigrate(&models.PromotionCode{})

	ReviewRepository := repository.NewReviewRepository(db)
	ReviewService := service.NewReviewService(ReviewRepository)
	ReviewController := controller.NewReviewController(ReviewService)
	ReviewRoute := routes.NewReviewRoute(ReviewController, router)
	ReviewRoute.Setup()
	db.DB.AutoMigrate(&models.Review{})

	StoreRepository := repository.NewStoreRepository(db)
	StoreService := service.NewStoreService(StoreRepository)
	StoreController := controller.NewStoreController(StoreService)
	StoreRoute := routes.NewStoreRoute(StoreController, router)
	StoreRoute.Setup()
	db.DB.AutoMigrate(&models.Store{})

	OpenTimeRepository := repository.NewOpenTimeRepository(db)
	OpenTimeService := service.NewOpenTimeService(OpenTimeRepository)
	OpenTimeController := controller.NewOpenTimeController(OpenTimeService)
	OpenTimeRoute := routes.NewOpenTimeRoute(OpenTimeController, router)
	OpenTimeRoute.Setup()
	db.DB.AutoMigrate(&models.OpenTime{})

	StoreImageRepository := repository.NewStoreImageRepository(db)
	StoreImageService := service.NewStoreImageService(StoreImageRepository)
	StoreImageController := controller.NewStoreImageController(StoreImageService)
	StoreImageRoute := routes.NewStoreImageRoute(StoreImageController, router)
	StoreImageRoute.Setup()
	db.DB.AutoMigrate(&models.StoreImage{})

	TableBookingRepository := repository.NewTableBookingRepository(db)
	TableBookingService := service.NewTableBookingService(TableBookingRepository)
	TableBookingController := controller.NewTableBookingController(TableBookingService)
	TableBookingRoute := routes.NewTableBookingRoute(TableBookingController, router)
	TableBookingRoute.Setup()
	db.DB.AutoMigrate(&models.TableBooking{})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()
	db.DB.AutoMigrate(&models.User{})

	//Seed Data

	categories := []models.Category{
		{Name: "ไทย", ImageName: "https://images.unsplash.com/photo-1554054204-b2f70b09d031?q=80&w=2022&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"},
		{Name: "นานาชาติ", ImageName: "https://images.unsplash.com/photo-1485921325833-c519f76c4927?q=80&w=1964&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"},
		{Name: "ญิ่ปุ่น", ImageName: "https://images.unsplash.com/photo-1580822184713-fc5400e7fe10?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"},
		{Name: "จีน", ImageName: "https://images.unsplash.com/photo-1544601284-7fe39c93d4d4?q=80&w=1654&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"},
		{Name: "อิตาเลี่ยน", ImageName: "https://images.unsplash.com/photo-1627042633145-b780d842ba45?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"},
		{Name: "ฟิวชั่น", ImageName: "https://api2.krua.co/wp-content/uploads/2022/06/ArticlePic_1670x1095-02-8-734x1024.jpg"},
	}

	locations := []models.Location{
		{Name: "สีลม", ImageName: "https://cdn-images.prod.thinkofliving.com/wp-content/uploads/1/2021/11/03150807/Silom_Skyline-1.jpg"},
		{Name: "สาทร", ImageName: "https://cockpit-images.s3.ap-southeast-1.amazonaws.com/2022-04-08/1649393269-DSCF8944.webp"},
		{Name: "พร้อมพงษ์", ImageName: "https://park.co.th/wp-content/uploads/2020/03/rsz_shutterstock_788969065-768x512.jpg"},
		{Name: "ราชเทวี", ImageName: "https://origin.co.th/wp-content/uploads/2019/08/light-rails-459064_1920-696x464.jpg"},
		{Name: "ห้าแยกลาดพร้าว", ImageName: "https://static.estopolis.com/article/591e58ce15f02071ef23a882_591e81b315f02071ef23a8c9.jpg"},
		{Name: "อโศก", ImageName: "https://cdn-cms.pgimgs.com/areainsider/2019/05/Asoke_09.jpg"},
        {Name: "อื่นๆ", ImageName: "https://twitter.com/AungporNapat/status/1546092996197658624/photo/1"},
	}

	stores := []models.Store{
		{CategoryID: 1 , Name: "store1", Location: "16/9 ถ. หอวัง แขวงจตุจักร เขตจตุจักร กรุงเทพมหานคร 10900 ประเทศไทย"},
		{CategoryID: 1 ,Name: "store2", Location: "Lat Yao, Chatuchak, Bangkok 10900"},
	}

    users := []models.User{
		{Username: "user1", Password: "por1111" , Status: "user"},
        {Username: "user2", Password: "por1111" , Status: "user"},
        {Username: "user3", Password: "por1111" , Status: "user"},
	}

	CategoryRepository.SeedData(categories)
	LocationRepository.SeedData(locations)
	StoreRepository.SeedData(stores)
    userRepository.SeedData(users)

	router.Gin.Run(":8000") //server started on 8000 port
}
