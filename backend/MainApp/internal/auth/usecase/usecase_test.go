package usecase_test

import (
	"strconv"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	authMocks "writesend/MainApp/internal/auth/repository/mocks"
	authUsecase "writesend/MainApp/internal/auth/usecase"
	userMocks "writesend/MainApp/internal/user/repository/mocks"
	"writesend/MainApp/models"
)

type TestCaseSignUp struct {
	ArgData     *models.User
	ExpectedRes uint64
	Error       error
}

type TestCaseSignIn struct {
	ArgData           models.UserSignIn
	ExpectedResUser   *models.User
	ExpectedResCookie uint64
	Error             error
}

type TestCaseDeleteCookie struct {
	ArgData string
	Error   error
}

type TestCaseAuth struct {
	ArgData  string
	Expected models.User
	Error    error
}

func TestUsecaseSignUp(t *testing.T) {
	var mockUserSuccess models.User
	err := faker.FakeData(&mockUserSuccess)
	assert.NoError(t, err)

	var mockUserConflictNickName models.User
	err = faker.FakeData(&mockUserConflictNickName)
	assert.NoError(t, err)

	var mockUserConflictEmail models.User
	err = faker.FakeData(&mockUserConflictEmail)
	assert.NoError(t, err)

	mockAuthRepo := authMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SelectUserByNickName", mockUserSuccess.NickName).Return(nil, status.Error(codes.NotFound, models.ErrNotFound.Error()))
	mockUserRepo.On("SelectUserByEmail", mockUserSuccess.Email).Return(nil, status.Error(codes.NotFound, models.ErrNotFound.Error()))
	mockUserRepo.On("CreateUser", &mockUserSuccess).Return(nil)
	mockAuthRepo.On("CreateCookie", mock.AnythingOfType("*models.Cookie")).Return(nil)

	mockUserRepo.On("SelectUserByNickName", mockUserConflictNickName.NickName).Return(&mockUserConflictNickName, nil)

	mockUserRepo.On("SelectUserByNickName", mockUserConflictEmail.NickName).Return(nil, status.Error(codes.NotFound, models.ErrNotFound.Error()))
	mockUserRepo.On("SelectUserByEmail", mockUserConflictEmail.Email).Return(&mockUserConflictEmail, nil)

	useCase := authUsecase.New(mockAuthRepo, mockUserRepo)

	cases := map[string]TestCaseSignUp{
		"success": {
			ArgData:     &mockUserSuccess,
			ExpectedRes: mockUserSuccess.Id,
			Error:       nil,
		},
		"conflict_nickname": {
			ArgData: &mockUserConflictNickName,
			Error:   models.ErrConflictNickname,
		},
		"conflict_email": {
			ArgData: &mockUserConflictEmail,
			Error:   models.ErrConflictEmail,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			cookie, err := useCase.SignUp(test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, cookie.UserId)
			}
		})
	}
	mockUserRepo.AssertExpectations(t)
	mockAuthRepo.AssertExpectations(t)
}

func TestUsecaseSignIn(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	var mockUserSignIn models.UserSignIn
	mockUserSignIn.Email = mockUser.Email
	mockUserSignIn.Password = mockUser.Password

	var mockUserSignInInvalidPassword models.UserSignIn
	err = faker.FakeData(&mockUserSignInInvalidPassword.Email)
	assert.NoError(t, err)
	mockUserSignInInvalidPassword.Password = "dfvdf"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mockUser.Password), 8)
	assert.NoError(t, err)

	mockUser.Password = string(hashedPassword)

	mockAuthRepo := authMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserFail := mockUser
	mockUserRepo.On("SelectUserByEmail", mockUserSignInInvalidPassword.Email).Return(&mockUserFail, nil)

	mockUserRepo.On("SelectUserByEmail", mockUserSignIn.Email).Return(&mockUser, nil)
	mockAuthRepo.On("CreateCookie", mock.AnythingOfType("*models.Cookie")).Return(nil)

	useCase := authUsecase.New(mockAuthRepo, mockUserRepo)

	expectedUser := mockUser
	expectedUser.Password = ""
	cases := map[string]TestCaseSignIn{
		"success": {
			ArgData:           mockUserSignIn,
			ExpectedResUser:   &expectedUser,
			ExpectedResCookie: mockUser.Id,
			Error:             nil,
		},
		"invalid_password": {
			ArgData: mockUserSignInInvalidPassword,
			Error:   models.ErrInvalidPassword,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			user, _, err := useCase.SignIn(test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedResUser, user)
			}
		})
	}
	mockUserRepo.AssertExpectations(t)
	mockAuthRepo.AssertExpectations(t)
}

func TestUsecaseDeleteCookie(t *testing.T) {
	var cookie models.Cookie
	err := faker.FakeData(&cookie)
	assert.NoError(t, err)

	var cookieGetFail models.Cookie
	err = faker.FakeData(&cookieGetFail)
	assert.NoError(t, err)

	var cookieDeleteFail models.Cookie
	err = faker.FakeData(&cookieDeleteFail)
	assert.NoError(t, err)

	mockAuthRepo := authMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockAuthRepo.On("GetCookie", cookie.SessionToken).Return(strconv.Itoa(int(cookie.UserId)), nil)
	mockAuthRepo.On("DeleteCookie", cookie.SessionToken).Return(nil)

	mockAuthRepo.On("GetCookie", cookieGetFail.SessionToken).Return("", models.ErrNotFound)

	mockAuthRepo.On("GetCookie", cookieDeleteFail.SessionToken).Return(strconv.Itoa(int(cookieDeleteFail.UserId)), nil)
	mockAuthRepo.On("DeleteCookie", cookieDeleteFail.SessionToken).Return(models.ErrInternalServerError)

	useCase := authUsecase.New(mockAuthRepo, mockUserRepo)

	cases := map[string]TestCaseDeleteCookie{
		"success": {
			ArgData: cookie.SessionToken,
			Error:   nil,
		},
		"fail_get": {
			ArgData: cookieGetFail.SessionToken,
			Error:   models.ErrNotFound,
		},
		"fail_delete": {
			ArgData: cookieDeleteFail.SessionToken,
			Error:   models.ErrInternalServerError,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.DeleteCookie(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockAuthRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseAuth(t *testing.T) {
	var cookie models.Cookie
	err := faker.FakeData(&cookie)
	assert.NoError(t, err)

	mockAuthRepo := authMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	var user models.User
	err = faker.FakeData(&user)
	assert.NoError(t, err)

	user.Id = cookie.UserId

	mockAuthRepo.On("GetCookie", cookie.SessionToken).Return(strconv.Itoa(int(cookie.UserId)), nil)
	mockUserRepo.On("SelectUserById", cookie.UserId).Return(&user, nil)

	user.Password = ""

	useCase := authUsecase.New(mockAuthRepo, mockUserRepo)

	cases := map[string]TestCaseAuth{
		"success": {
			ArgData:  cookie.SessionToken,
			Expected: user,
			Error:    nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			gotUser, err := useCase.Auth(test.ArgData)
			require.Equal(t, test.Error, err)
			if err == nil {
				assert.Equal(t, &test.Expected, gotUser)
			}
		})
	}
	mockAuthRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}
