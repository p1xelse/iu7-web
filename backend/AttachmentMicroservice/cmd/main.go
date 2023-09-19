package main

import (
	"log"
	"net"

	attachmentDelivery "writesend/AttachmentMicroservice/internal/attachment/delivery"
	attachmentRep "writesend/AttachmentMicroservice/internal/attachment/repository/postgres"
	attachmentUsecase "writesend/AttachmentMicroservice/internal/attachment/usecase"
	attachment "writesend/AttachmentMicroservice/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var testCfgPg = postgres.Config{DSN: "host=localhost user=postgres password=postgres port=13080"}

var prodCfgPg = postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.New(prodCfgPg),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	attachmentDB := attachmentRep.NewAttachmentRepository(db)
	attachmentUC := attachmentUsecase.New(attachmentDB)
	attachment.RegisterAttachmentsServer(server, attachmentDelivery.New(attachmentUC))

	log.Println("starting server at :8082")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
