package main

import (
	"writesend/MainApp/internal/middleware"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"writesend/MainApp/cmd/server"
	_ "writesend/MainApp/docs"
	_attachmentDelivery "writesend/MainApp/internal/attachment/delivery"
	attachmentsRepository "writesend/MainApp/internal/attachment/repository/microservice"
	attachmentUsecase "writesend/MainApp/internal/attachment/usecase"
	_authDelivery "writesend/MainApp/internal/auth/delivery"
	authRep "writesend/MainApp/internal/auth/repository/microservice"
	authUseCase "writesend/MainApp/internal/auth/usecase"
	_chatDelivery "writesend/MainApp/internal/chat/delivery"
	chatRep "writesend/MainApp/internal/chat/repository/microservice"
	chatUseCase "writesend/MainApp/internal/chat/usecase"
	_communitiesDelivery "writesend/MainApp/internal/communities/delivery"
	communitiesRep "writesend/MainApp/internal/communities/repository/postgres"
	communitiesUseCase "writesend/MainApp/internal/communities/usecase"
	_friendsDelivery "writesend/MainApp/internal/friends/delivery"
	friendsRep "writesend/MainApp/internal/friends/repository/microservice"
	friendsUseCase "writesend/MainApp/internal/friends/usecase"
	_postsDelivery "writesend/MainApp/internal/post/delivery"
	postsRep "writesend/MainApp/internal/post/repository/postgres"
	postsUsecase "writesend/MainApp/internal/post/usecase"
	_stickersDelivery "writesend/MainApp/internal/stickers/delivery"
	stickersRep "writesend/MainApp/internal/stickers/repository/postgres"
	stickersUseCase "writesend/MainApp/internal/stickers/usecase"
	_usersDelivery "writesend/MainApp/internal/user/delivery"
	usersRep "writesend/MainApp/internal/user/repository/microservice"
	usersUseCase "writesend/MainApp/internal/user/usecase"
	attachment "writesend/MainApp/proto/attachment"
	auth "writesend/MainApp/proto/auth"
	chat "writesend/MainApp/proto/chat"
	user "writesend/MainApp/proto/user"

	"github.com/labstack/echo-contrib/prometheus"
)

// @title WS Swagger API
// @version 1.0
// @host 89.208.197.127:8080

// var testCfgPg = postgres.Config{DSN: "host=localhost user=postgres password=postgres port=13080"}

var prodCfgPg = postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}

func main() {
	db, err := gorm.Open(postgres.New(prodCfgPg),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Info("postgres connect success")

	grpcConnAuth, err := grpc.Dial(
		"auth_mvs:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnAuth.Close()
	authManager := auth.NewAuthClient(grpcConnAuth)

	grpcConnAttachment, err := grpc.Dial(
		"attachment_mvs:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnAttachment.Close()
	attachmentManager := attachment.NewAttachmentsClient(grpcConnAttachment)

	grpcConnChat, err := grpc.Dial(
		"chat_mvs:8083",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnChat.Close()
	chatManager := chat.NewChatClient(grpcConnChat)

	grpcConnUser, err := grpc.Dial(
		"user_mvs:8084",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnUser.Close()
	userManager := user.NewUsersClient(grpcConnUser)

	postDB := postsRep.NewPostRepository(db)
	authDB := authRep.New(authManager)
	usersDB := usersRep.New(userManager)
	friendsDB := friendsRep.New(userManager)
	attachmentDB := attachmentsRepository.New(attachmentManager)
	chatDB := chatRep.New(chatManager)
	communitiesDb := communitiesRep.NewCommunitiesRepository(db)
	stickersDb := stickersRep.New(db)

	postsUC := postsUsecase.NewPostUsecase(postDB, attachmentDB, usersDB)
	authUC := authUseCase.New(authDB, usersDB)
	usersUC := usersUseCase.New(usersDB)
	friendsUC := friendsUseCase.New(friendsDB, usersDB)
	attachmentUC := attachmentUsecase.NewAttachmentUsecase(attachmentDB)
	chatUC := chatUseCase.New(chatDB, attachmentDB)
	communitiesUC := communitiesUseCase.New(communitiesDb)
	stickersUC := stickersUseCase.NewStickerUsecase(stickersDb)

	e := echo.New()

	e.Logger.SetHeader(`time=${time_rfc3339} level=${level} prefix=${prefix} ` +
		`file=${short_file} line=${line} message:`)
	e.Logger.SetLevel(log.INFO)

	p := prometheus.NewPrometheus("echo", nil)
	p.MetricsPath = "/prometheus"
	p.SetMetricsPath(e)
	p.Use(e)

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://writesend.online", "http://89.208.197.127", "http://localhost"},
		AllowHeaders:     []string{"Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"X-CSRF-Token"},
	}))

	e.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: `time=${time_custom} remote_ip=${remote_ip} ` +
			`host=${host} method=${method} uri=${uri} user_agent=${user_agent} ` +
			`status=${status} error="${error}" ` +
			`bytes_in=${bytes_in} bytes_out=${bytes_out}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	e.Use(echoMiddleware.Recover())

	authMiddleware := middleware.NewMiddleware(authUC)
	e.Use(authMiddleware.Auth)
	e.Use(authMiddleware.CSRF)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	_postsDelivery.NewDelivery(e, postsUC)
	_authDelivery.NewDelivery(e, authUC)
	_usersDelivery.NewDelivery(e, usersUC)
	_attachmentDelivery.NewDelivery(e, attachmentUC)
	_friendsDelivery.NewDelivery(e, friendsUC)
	_chatDelivery.NewDelivery(e, chatUC)
	_communitiesDelivery.NewDelivery(e, communitiesUC)
	_stickersDelivery.NewDelivery(e, stickersUC)

	s := server.NewServer(e)
	if err := s.Start(); err != nil {
		e.Logger.Fatal(err)
	}
}
