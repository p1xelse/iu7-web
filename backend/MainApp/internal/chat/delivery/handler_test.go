package delivery_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	chatDelivery "writesend/MainApp/internal/chat/delivery"
	"writesend/MainApp/internal/chat/usecase/mocks"
	"writesend/MainApp/models"
)

type TestCaseGetAllDialogs struct {
	ArgDataContext uint64
	Error          error
	StatusCode     int
}

type TestCaseGetDialog struct {
	ArgData    string
	Error      error
	StatusCode int
}

type TestCaseGetDialogByUsers struct {
	ArgDataParam   string
	ArgDataContext uint64
	Error          error
	StatusCode     int
}

type TestCaseSendMessage struct {
	ArgDataBody    string
	ArgDataContext uint64
	Error          error
	StatusCode     int
}

type TestCaseWsChatHandler struct {
	ArgDataParam string
	Error        error
	StatusCode   int
}

func TestDeliveryGetDialog(t *testing.T) {
	var dialog models.Dialog
	err := faker.FakeData(&dialog)
	assert.NoError(t, err)
	dialog.Id = 1

	var dialogNotFound models.Dialog
	err = faker.FakeData(&dialogNotFound)
	assert.NoError(t, err)
	dialogNotFound.Id = 2

	var dialogInternalErr models.Dialog
	err = faker.FakeData(&dialogInternalErr)
	assert.NoError(t, err)
	dialogInternalErr.Id = 3

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("SelectDialog", dialog.Id).Return(&dialog, nil)
	mockUCase.On("SelectDialog", dialogNotFound.Id).Return(nil, models.ErrNotFound)
	mockUCase.On("SelectDialog", dialogInternalErr.Id).Return(nil, models.ErrInternalServerError)

	handler := chatDelivery.Delivery{
		ChatUC: mockUCase,
	}

	dialogIdBadRequest := "jgsv"

	e := echo.New()
	chatDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCaseGetDialog{
		"success": {
			ArgData:    strconv.Itoa(int(dialog.Id)),
			Error:      nil,
			StatusCode: http.StatusOK,
		},
		"bad_request": {
			ArgData: dialogIdBadRequest,
			Error: &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"not_found": {
			ArgData: strconv.Itoa(int(dialogNotFound.Id)),
			Error: &echo.HTTPError{
				Code:    http.StatusNotFound,
				Message: models.ErrNotFound.Error(),
			},
		},
		"internal_error": {
			ArgData: strconv.Itoa(int(dialogInternalErr.Id)),
			Error: &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/chat/:id", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/chat/:id")
			c.SetParamNames("id")
			c.SetParamValues(test.ArgData)

			err := handler.GetDialog(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestDeliveryGetDialogByUsers(t *testing.T) {
	var dialog models.Dialog
	err := faker.FakeData(&dialog)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("SelectDialogByUsers", dialog.UserId1, dialog.UserId2).Return(&dialog, nil)
	handler := chatDelivery.Delivery{
		ChatUC: mockUCase,
	}

	userIdBadRequest := "jgsv"

	e := echo.New()
	chatDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCaseGetDialogByUsers{
		"success": {
			ArgDataParam:   strconv.Itoa(int(dialog.UserId2)),
			ArgDataContext: dialog.UserId1,
			Error:          nil,
			StatusCode:     http.StatusOK,
		},
		"bad_request": {
			ArgDataParam:   userIdBadRequest,
			ArgDataContext: dialog.UserId1,
			Error: &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"invalid_context": {
			ArgDataParam: "2",
			Error: &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/chat/user/:id", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/chat/user/:id")
			c.SetParamNames("id")
			c.SetParamValues(test.ArgDataParam)
			if name != "invalid_context" {
				c.Set("user_id", test.ArgDataContext)
			}

			err := handler.GetDialogByUsers(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestDeliveryGetAllDialogs(t *testing.T) {
	dialogs := make([]models.Dialog, 0, 10)
	err := faker.FakeData(&dialogs)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	var userId uint64 = 1

	mockUCase.On("SelectAllDialogs", userId).Return(dialogs, nil)
	handler := chatDelivery.Delivery{
		ChatUC: mockUCase,
	}

	e := echo.New()
	chatDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCaseGetAllDialogs{
		"success": {
			ArgDataContext: userId,
			Error:          nil,
			StatusCode:     http.StatusOK,
		},
		"internal_error": {
			Error: &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/chat", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/chat")
			if name != "internal_error" {
				c.Set("user_id", test.ArgDataContext)
			}

			err := handler.GetAllDialogs(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestDeliverySendMessage(t *testing.T) {
	var message models.Message
	err := faker.FakeData(&message)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("SendMessage", mock.AnythingOfType("*models.Message")).Return(nil)
	handler := chatDelivery.Delivery{
		ChatUC: mockUCase,
	}

	jsonMessage, err := json.Marshal(message)
	assert.NoError(t, err)

	e := echo.New()
	chatDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCaseSendMessage{
		"success": {
			ArgDataBody:    string(jsonMessage),
			ArgDataContext: message.SenderID,
			Error:          nil,
			StatusCode:     http.StatusOK,
		},
		"bad_request": {
			ArgDataBody: "aaa",
			Error: &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"internal_error": {
			Error: &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/chat/send_message", strings.NewReader(test.ArgDataBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/chat/send_message")
			if name != "internal_error" {
				c.Set("user_id", test.ArgDataContext)
			}

			err := handler.SendMessage(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestDeliveryWsChatHandler(t *testing.T) {
	mockUCase := mocks.NewUseCaseI(t)

	handler := chatDelivery.Delivery{
		ChatUC: mockUCase,
	}

	e := echo.New()
	chatDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCaseWsChatHandler{
		"success": {
			ArgDataParam: "1",
			Error:        nil,
			StatusCode:   http.StatusBadRequest,
		},
		"bad_request": {
			ArgDataParam: "skdvb",
			Error: &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/ws/:roomId", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/ws/:roomId")
			c.SetParamNames("roomId")
			c.SetParamValues(test.ArgDataParam)

			err := handler.WsChatHandler(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}
