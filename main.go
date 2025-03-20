package main

import (
	"fmt"
	"log"

	"github.com/deanFL0/blog_api_go/api/routes"
	"github.com/deanFL0/blog_api_go/pkg/article"
	"github.com/deanFL0/blog_api_go/pkg/entities"
	users "github.com/deanFL0/blog_api_go/pkg/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database settings
const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "admin"
	dbname   = "blog_go"
)

// Database connection
func databaseConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := databaseConnection()
	if err != nil {
		log.Fatal("Database connection error.\n", err)
	}
	fmt.Println("Database connection success")
	db.AutoMigrate(&entities.Article{})
	db.AutoMigrate(&entities.User{})

	articleRepo := article.NewRepo(db)
	articleService := article.NewService(articleRepo)

	userRepo := users.NewRepo(db)
	userService := users.NewService(userRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to Blog!"))
	})
	api := app.Group("/api")
	routes.ArticleRouter(api, articleService)
	routes.UserRouter(api, userService)
	log.Fatal(app.Listen(":8080"))
}
