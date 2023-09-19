package main

import (
	"log"
	"net"
	userDelivery "writesend/UserMicroservice/internal/user/delivery"
	userRep "writesend/UserMicroservice/internal/user/repository/postgres"
	userUsecase "writesend/UserMicroservice/internal/user/usecase"
	user "writesend/UserMicroservice/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var testCfgPg = postgres.Config{DSN: "host=localhost user=postgres password=postgres port=13080"}

var prodCfgPg = postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}

func main() {
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.New(prodCfgPg),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	userDB := userRep.New(db)
	userUC := userUsecase.New(userDB)
	user.RegisterUsersServer(server, userDelivery.New(userUC))

	log.Println("starting server at :8084")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
