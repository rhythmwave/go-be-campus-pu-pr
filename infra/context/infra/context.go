package infra

import (
	"github.com/go-redis/redis/v8"
	"github.com/sccicitb/pupr-backend/config"
	middleware "github.com/sccicitb/pupr-backend/infra/base_middleware"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/infra/file"
	mailer "github.com/sccicitb/pupr-backend/infra/mail"
	"github.com/sccicitb/pupr-backend/infra/otp"
	"github.com/sccicitb/pupr-backend/notification"
	"github.com/sccicitb/pupr-backend/utils"
)

// InfraCtx struct for infra context configuration
type InfraCtx struct {
	DB                   *db.DB
	DBLog                *db.DB
	Config               *config.Config
	RedisClient          *redis.Client
	Jwt                  middleware.JWTInterface
	Mail                 mailer.MailInterface
	NotificationTemplate notification.NotificationTemplate
	Otp                  otp.OTPInterface
	Storage              file.FileCtx
	Notification         utils.NotificationInterface
}
