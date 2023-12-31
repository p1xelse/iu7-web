package postgres_test

import (
	"regexp"
	"testing"
	"time"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	chatRep "writesend/ChatMicroservice/internal/chat/repository/postgres"
	"writesend/ChatMicroservice/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestRepositoryCreateDialog(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	gdb.Logger.LogMode(logger.Info)

	mock.ExpectBegin()

	var mockDialog models.Dialog
	err = faker.FakeData(&mockDialog)
	assert.NoError(t, err)

	mockDialog.Id = 1

	var mockId uint64 = 1

	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "chat" ("user_id1","user_id2","id") VALUES ($1,$2,$3) RETURNING "id"`)).WithArgs(
		mockDialog.UserId1, mockDialog.UserId2, mockDialog.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockId))

	mock.ExpectCommit()

	repository := chatRep.NewChatRepository(gdb)

	err = repository.CreateDialog(&mockDialog)
	require.NoError(t, err)
	assert.Equal(t, mockId, mockDialog.Id)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepositoryCreateMessage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	gdb.Logger.LogMode(logger.Info)

	mock.ExpectBegin()

	mockMessage := models.Message{
		ID:         1,
		DialogID:   1,
		SenderID:   1,
		ReceiverID: 1,
		Body:       "body",
	}

	mockMessage.Attachments = make([]models.Attachment, 2)

	mockMessage.Attachments[0] = models.Attachment{
		ID:      1,
		AttLink: "link1",
	}

	mockMessage.Attachments[1] = models.Attachment{
		ID:      2,
		AttLink: "link2",
	}

	mockMessage.CreatedAt = time.Now()

	var mockId uint64 = 1

	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "message" ("chat_id","sender_id","receiver_id","text","created_at","id") `+
			`VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).WithArgs(
		mockMessage.DialogID, mockMessage.SenderID, mockMessage.ReceiverID, mockMessage.Body,
		mockMessage.CreatedAt, mockMessage.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockId))

	mock.ExpectCommit()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "message_attachments" ("message_id","att_id") VALUES ($1,$2),($3,$4)`)).
		WithArgs(mockMessage.ID, mockMessage.Attachments[0].ID, mockMessage.ID, mockMessage.Attachments[1].ID).
		WillReturnResult(sqlmock.NewResult(int64(1), 1))

	mock.ExpectCommit()

	repository := chatRep.NewChatRepository(gdb)

	err = repository.CreateMessage(&mockMessage)
	require.NoError(t, err)
	assert.Equal(t, mockId, mockMessage.ID)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepositorySelectDialog(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	gdb.Logger.LogMode(logger.Info)

	mockDialog := models.Dialog{
		Id:      1,
		UserId1: 1,
		UserId2: 2,
	}

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "chat" WHERE id = $1 LIMIT 1`)).WithArgs(mockDialog.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id1", "user_id2"}).
			AddRow(mockDialog.Id, mockDialog.UserId1, mockDialog.UserId2))

	repository := chatRep.NewChatRepository(gdb)

	actualUser, err := repository.SelectDialog(mockDialog.Id)
	require.NoError(t, err)
	assert.Equal(t, mockDialog, *actualUser)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepositorySelectDialogByUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	gdb.Logger.LogMode(logger.Info)

	mockDialog := models.Dialog{
		Id:      1,
		UserId1: 1,
		UserId2: 2,
	}

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "chat" WHERE (user_id1 = $1 AND user_id2 = $2) OR `+
			`(user_id1 = $3 AND user_id2 = $4) LIMIT 1`)).WithArgs(mockDialog.UserId1,
		mockDialog.UserId2, mockDialog.UserId2, mockDialog.UserId1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id1", "user_id2"}).
			AddRow(mockDialog.Id, mockDialog.UserId1, mockDialog.UserId2))

	repository := chatRep.NewChatRepository(gdb)

	actualUser, err := repository.SelectDialogByUsers(mockDialog.UserId1, mockDialog.UserId2)
	require.NoError(t, err)
	assert.Equal(t, mockDialog, *actualUser)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepositorySelectAllDialogs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	gdb.Logger.LogMode(logger.Info)

	mockDialogs := make([]models.Dialog, 0, 10)
	err = faker.FakeData(&mockDialogs)
	assert.NoError(t, err)

	var userId uint64 = 1

	rows := sqlmock.NewRows([]string{"id", "user_id1", "user_id2"})

	for i := range mockDialogs {
		mockDialogs[i].Messages = nil
		rows.AddRow(mockDialogs[i].Id, mockDialogs[i].UserId1, mockDialogs[i].UserId2)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "chat"."id","chat"."user_id1","chat"."user_id2" ` +
		`FROM "chat" WHERE user_id1 = $1 OR user_id2 = $2`)).
		WillReturnRows(rows)

	repository := chatRep.NewChatRepository(gdb)

	actualDialogs, err := repository.SelectAllDialogs(userId)
	require.NoError(t, err)
	assert.Equal(t, mockDialogs, actualDialogs)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepositorySelectMessages(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	gdb.Logger.LogMode(logger.Info)

	mockMessages := make([]models.Message, 0, 10)
	err = faker.FakeData(&mockMessages)
	assert.NoError(t, err)

	for idx := range mockMessages {
		mockMessages[idx].Attachments = nil
	}

	var chatId uint64 = 1

	rows := sqlmock.NewRows([]string{"id", "chat_id", "sender_id", "receiver_id", "text", "created_at", "sticker_id"})

	for _, mockMessage := range mockMessages {
		rows.AddRow(mockMessage.ID, mockMessage.DialogID, mockMessage.SenderID,
			mockMessage.ReceiverID, mockMessage.Body, mockMessage.CreatedAt, mockMessage.StickerID)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "message" WHERE chat_id = $1 ORDER BY id`)).
		WillReturnRows(rows)

	repository := chatRep.NewChatRepository(gdb)

	actualMessages, err := repository.SelectMessages(chatId)
	require.NoError(t, err)
	assert.Equal(t, mockMessages, actualMessages)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
