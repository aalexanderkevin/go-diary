package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

// Connect to db mysql
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:"+os.Getenv("MYSQL_PORT")+")/tap_talk")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// Connect to redis
func Redis() *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})
	_, err := r.Ping().Result()
	if err != nil {
		panic(err)
	}

	return r
}

// GET UUID from Redis
func FetchUUID(username string) (string, error) {
	result, err := Redis().Get(username).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}
