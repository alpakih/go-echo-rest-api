package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"go-echo-rest-api/app/helpers"
	"go-echo-rest-api/app/http/controllers"
	"go-echo-rest-api/app/http/controllers/backend"
	auth "go-echo-rest-api/app/http/middleware"
	"html/template"

	"go-echo-rest-api/db"
)

func New() *echo.Echo {

	e := echo.New()
	database := db.New()
	db.AutoMigrate(database)

	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Validator = NewValidator()

	t := &Template{
		templates: template.Must(template.ParseGlob("resources/views/*.html")),
	}

	e.Renderer = t

	jwtMiddleware := auth.JWT(helpers.JWTSecret)

	v1 := e.Group("/api")
	v1.POST("/register", controllers.NewCustomerController(database).Store)
	v1.POST("/login", controllers.NewCustomerController(database).Login)

	v2 := v1.Group("/customers", jwtMiddleware)
	v2.GET("", controllers.NewCustomerController(database).FindAll)
	v2.GET("/current", controllers.NewCustomerController(database).CurrentUser)
	v2.GET("/:id", controllers.NewCustomerController(database).GetByID)
	v2.PUT("/:id", controllers.NewCustomerController(database).Update)
	v2.DELETE("/:id", controllers.NewCustomerController(database).Destroy)

	e.GET("/hello", backend.Hello)
	return e
}
