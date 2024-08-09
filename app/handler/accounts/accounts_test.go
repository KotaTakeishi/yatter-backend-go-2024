package accounts

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountUsecase struct {
	mock.Mock
}

func (m *MockAccountUsecase) FindByUsername(ctx context.Context, username string) (*usecase.GetAccountDTO, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*usecase.GetAccountDTO), args.Error(1)
}

func (m *MockAccountUsecase) Create(ctx context.Context, username, password string) (*usecase.CreateAccountDTO, error) {
	args := m.Called(ctx, username, password)
	return args.Get(0).(*usecase.CreateAccountDTO), args.Error(1)
}

func (m *MockAccountUsecase) Update(ctx context.Context, id int64, display_name, note, avatar, header *string) (*usecase.UpdateAccountDTO, error) {
	args := m.Called(ctx, id, display_name, note, avatar, header)
	return args.Get(0).(*usecase.UpdateAccountDTO), args.Error(1)
}

func (m *MockAccountUsecase) Follow(ctx context.Context, followerID, followeeID int64) (*usecase.FollowAccountDTO, error) {
	args := m.Called(ctx, followerID, followeeID)
	return args.Get(0).(*usecase.FollowAccountDTO), args.Error(1)
}

func TestFindByUsername(t *testing.T) {
	mockUsecase := new(MockAccountUsecase)
	handler := &handler{accountUsecase: mockUsecase}

	account := &object.Account{
		Username: "testuser",
	}

	mockUsecase.On("FindByUsername", mock.Anything, "testuser").Return(&usecase.GetAccountDTO{Account: account}, nil)

	req := httptest.NewRequest(http.MethodGet, "/accounts/testuser", nil)
	req = req.WithContext(context.Background())
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Get("/accounts/{username}", handler.FindByUsername)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var responseAccount object.Account
	err := json.NewDecoder(rr.Body).Decode(&responseAccount)
	assert.NoError(t, err)
	assert.Equal(t, "testuser", responseAccount.Username)

	mockUsecase.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockUsecase := new(MockAccountUsecase)
	handler := &handler{accountUsecase: mockUsecase}

	reqBody := AddRequest{
		Username: "testuser",
		Password: "testpass",
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/v1/accounts", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	mockUsecase.On("Create", mock.Anything, "testuser", "testpass").Return(&usecase.CreateAccountDTO{
		Account: &object.Account{
			Username: "testuser",
		},
	}, nil)

	handler.Create(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp object.Account
	err := json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, "testuser", resp.Username)

	mockUsecase.AssertExpectations(t)
}
