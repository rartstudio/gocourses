package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rartstudio/gocourses/common"
	"github.com/rartstudio/gocourses/initializers"
	v1 "github.com/rartstudio/gocourses/routes/v1"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, customValidator *common.CustomValidator, config *initializers.Config, db *gorm.DB, mail *gomail.Dialer, redis *redis.Client) {
	v1.SetupRoutesAuthV1(app, customValidator, config, db, mail, redis)
	v1.SetupRoutesUserV1(app, customValidator, config, db);
}