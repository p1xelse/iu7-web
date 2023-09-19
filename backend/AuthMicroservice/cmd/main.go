package main

import (
	"log"
	"net"
	authDelivery "writesend/AuthMicroservice/internal/auth/delivery"
	authRep "writesend/AuthMicroservice/internal/auth/repository/redis"
	authUsecase "writesend/AuthMicroservice/internal/auth/usecase"
	auth "writesend/AuthMicroservice/proto"

	"github.com/go-redis/redis"
	"google.golang.org/grpc"
)

// var testCfgRedis = &redis.Options{Addr: ":6379", Password: "ws_redis_password"}

var prodCfgRedis = &redis.Options{Addr: "redis:6379", Password: "ws_redis_password"}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	redisClient := redis.NewClient(prodCfgRedis)

	err = redisClient.Ping().Err()
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	authDB := authRep.New(redisClient)
	authUC := authUsecase.New(authDB)
	auth.RegisterAuthServer(server, authDelivery.New(authUC))

	log.Println("starting server at :8081")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
