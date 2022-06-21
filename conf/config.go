package conf

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	Server struct {
		Addr string `yaml:"addr"`
	}
	Mysql struct {
		Dsn string `yaml:"dsn"`
	}
	Redis struct {
		Addr       string        `yaml:"addr"`
		Database   int           `yaml:"database"`
		Password   string        `yaml:"password"`
		Expiration time.Duration `yaml:"expiration"`
	}
}

var (
	Cfg            *Config
	DB             *gorm.DB
	Rdb            *redis.Client
	configFilename = flag.String("config_filename", "config.yml", "Config Filename")
)

func init() {
	flag.Parse()
	cfg, err := NewConfig(*configFilename)
	if err != nil {
		log.Fatal(err)
	}
	Cfg = cfg
	fmt.Println(cfg)
	// init redis client
	err = initRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	// init mysql db session
	err = initMysqlDB()
	if err != nil {
		log.Fatal(err)
	}
}

func NewConfig(filename string) (cfg *Config, err error) {
	yml, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	cfg = new(Config)
	err = yaml.Unmarshal(yml, cfg)
	return
}

// init redis client
func initRedisClient() (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Cfg.Redis.Addr,
		Password: Cfg.Redis.Password,
		DB:       Cfg.Redis.Database,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return
	}
	Rdb = rdb
	return
}

// init mysql db session
func initMysqlDB() (err error) {
	dsn := Cfg.Mysql.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	DB = db
	return
}
