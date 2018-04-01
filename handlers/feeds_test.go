package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedFeedService struct {
	mock.Mock
}

func (m *MockedFeedService) CreateNew(eventType string, payload string) {
	m.Called(eventType, payload)
}

func Test_CanAcceptANewFeed(t *testing.T) {
	mockFeedService := new(MockedFeedService)
	mockFeedService.On("CreateNew", "invoice", "boo").Return()

	handler := Create(mockFeedService)
	body := strings.NewReader("{\"payload\":\"boo\"}")
	req, _ := http.NewRequest("POST", "/feeds", body)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, "", w.Body.String())

	mockFeedService.AssertExpectations(t)
}
