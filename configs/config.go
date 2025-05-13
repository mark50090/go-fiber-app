package configs

import (
	"fmt"
	"strings"

	// "time"

	v "github.com/spf13/viper"
)

type Config struct {
	XThreads        int
	SmallXThreads   int
	BudgetYearStart int
	App             App
	Database        Database
	// Credential      Credential
	// Signature       Signature
}

type App struct {
	Name                  string
	Port                  string
	Prefork               bool
	DisableStartupMessage bool
	Env                   string
	Cors                  Cors
}

type Cors struct {
	AllowOrigins string
}

type Database struct {
	Connection string
	DbName     string
}

// type Credential struct {
// 	Fdh Fdh
// }

// type Fdh struct {
// 	Url       string
// 	HeaderKey string
// 	Token     string
// }

// type Signature struct {
// 	AesKey     string
// 	PrivateKey string
// 	PublicKey  string
// }

func NewConfig() *Config {

	v.SetConfigFile(".env")                            // ใช้ .env
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // เปลี่ยน . เป็น _
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("error reading config file: ", err)
	}

	var cfg Config
	cfg.XThreads = v.GetInt("APP_XTHREADS")
	cfg.SmallXThreads = v.GetInt("APP_SMALL_XTHREADS")
	cfg.BudgetYearStart = v.GetInt("APP_BUDGET_YEAR_START")
	cfg.App = App{
		Name:                  v.GetString("APP_NAME"),
		Port:                  v.GetString("APP_PORT"),
		Prefork:               v.GetBool("APP_PREFORK"),
		DisableStartupMessage: v.GetBool("APP_DISABLE_STARTUP_MESSAGE"),
		Env:                   v.GetString("APP_ENV"),
		Cors: Cors{
			AllowOrigins: v.GetString("APP_CORS_ALLOW_ORIGINS"),
		},
	}
	cfg.Database = Database{
		Connection: v.GetString("DATABASE_URI"),
		DbName:     v.GetString("DATABASE_DBNAME"),
	}

	// println(cfg.Database)
	// ไม่มีการตั้งค่า Redis cache
	// cfg.Redis = Redis{
	// 	Connection: v.GetString("REDIS_HOST"),
	// 	TTL:        v.GetDuration("REDIS_TTL"),
	// 	PrefixKey:  v.GetString("REDIS_PREFIX_KEY"),
	// }

	// cfg.Credential = Credential{
	// 	Fdh: Fdh{
	// 		Url:       v.GetString("FDH_URL"),
	// 		HeaderKey: v.GetString("FDH_HEADER_KEY"),
	// 		Token:     v.GetString("FDH_TOKEN"),
	// 	},
	// }

	// cfg.Signature = Signature{
	// 	AesKey:     v.GetString("SIGNATURE_AES_KEY"),
	// 	PrivateKey: v.GetString("SIGNATURE_PRIVATE_KEY"),
	// 	PublicKey:  v.GetString("SIGNATURE_PUBLIC_KEY"),
	// }

	return &cfg
}
