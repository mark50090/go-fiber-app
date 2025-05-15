package servers

import (
	"fmt"
	"go-fiber-app/configs"
	"go-fiber-app/core/controller"
	"go-fiber-app/core/repositories"
	services "go-fiber-app/core/service"
	"go-fiber-app/pkg"

	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	App  *fiber.App
	MgDb *mongo.Database
	cfg  *configs.Config
	// cache pkg.Cache
	log pkg.AppLog
}

func NewServer(config *configs.Config, client *mongo.Database, log pkg.AppLog) *server {
	fiberCfg := fiber.Config{
		AppName: config.App.Name,
		// EnableTrustedProxyCheck: config.App.EnableTrustedProxyCheck,
		ProxyHeader: "X-Forwarded-For",
		// Prefork:                 true,
		ReadBufferSize:  1024 * 1024 * 15,
		WriteBufferSize: 1024 * 1024 * 15,
	}
	// if len(config.App.TrustedProxies) > 0 && config.App.TrustedProxies[0] != "" {
	// 	fiberCfg.TrustedProxies = config.App.TrustedProxies
	// }

	return &server{
		App:  fiber.New(fiberCfg),
		MgDb: client,
		cfg:  config,
		// cache: cache,
		log: log,
	}
}

func (s *server) Start() {

	s.App.Use(s.Cors(s.cfg))
	s.App.Use(s.middlewareLogger)
	s.App.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))

	api := s.App.Group("api")
	v1 := api.Group("v1")
	v1.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	unitCostRepo := repositories.NewUnitCostRepository(s.MgDb)
	unitCostSrv := services.NewUnitCostService(*s.cfg, unitCostRepo)
	unitCostCtl := controller.NewHandlerUnitCost(unitCostSrv)
	s.registerRoutesUnitCost(v1.Group("unit-cost"), unitCostCtl)

	siaMarkRepo := repositories.NewHintRepository(s.MgDb)
	siaMarkSrv := services.NewHintService(*s.cfg, siaMarkRepo)
	siaMarkController := controller.NewHandlerHint(siaMarkSrv)
	s.siaMarkControllerRoute(v1.Group("siamark"), siaMarkController)

	// unitCommentRepo := repositories.NewUnitCommentRepo(s.MgDb)
	// unitCommentSrv := services.NewUnitCommentService(unitCommentRepo)
	// unitCommentCtl := controller.NewHandlerUnitComment(unitCommentSrv)
	// s.registerRoutesUnitComment(v1.Group("unit-comment"), unitCommentCtl)

	// sumRepo := repositories.NewSummaryRepository(s.MgDb)
	// sumSrv := services.NewSummarizeService(s.log, *s.cfg, sumRepo, unitCostRepo)
	// sumCtl := controller.NewControllerSchedule(sumSrv)
	// s.registerRoutesSummary(v1.Group("schedules"), sumCtl)

	s.App.Get("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(
			map[string]interface{}{
				"message": "Not Found",
				"code":    fiber.StatusNotFound,
			},
		)
	})

	fiberConnURL := fmt.Sprintf(":%s", s.cfg.App.Port)
	if strings.ToLower(s.cfg.App.Env) == "development" {
		//# เมื่อทำการ Run แบบ Developer Mode

		fiberConnURL = "localhost:" + s.cfg.App.Port
	} else {
		//# เมื่อต้องการ Run แบบ Production Mode
		fiberConnURL = ":" + s.cfg.App.Port
	}
	if err := s.App.Listen(fiberConnURL); err != nil {
		panic(err.Error())
	}
}
func (s *server) registerRoutesUnitCost(router fiber.Router, controller *controller.UnitCostController) {
	// router.Get("/example", controller.GetExample)
	router.Get("/table", controller.GetUnitCostSummary)
	router.Get("/dropdown", controller.GetInsclDetail)

}

// func (s *server) registerRoutesSummary(router fiber.Router, controller *controller.SummarizeController) {
// 	// router.Post("/example", controller.Example)
// 	router.Post("/summarize-unit-cost-diary", controller.SummarizeUnitCostDiary)
// 	// router.Post("/summarize-unit-cost-hospital-diary", controller.SummarizeUnitCostHospitalDiary)
// }

// func (s *server) registerRoutesUnitComment(router fiber.Router, controller *controller.UnitCommentController) {
// 	router.Post("/", controller.CreateComment)
// 	router.Get("/", controller.GetComment)
// }

func (s *server) siaMarkControllerRoute(router fiber.Router, controller *controller.HintController) {
	router.Get("/transaction", controller.GetTransaction)
	router.Post("/excel", controller.ReportExcelRegistrationV3)
}
