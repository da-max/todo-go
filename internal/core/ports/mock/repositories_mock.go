// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/core/ports/repositories.go
//
// Generated by this command:
//
//	mockgen -source=./internal/core/ports/repositories.go -destination=./internal/core/ports/mock/repositories_mock
//
// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"
	time "time"

	domain "github.com/Da-max/todo-go/internal/core/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockTodoRepository is a mock of TodoRepository interface.
type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

// MockTodoRepositoryMockRecorder is the mock recorder for MockTodoRepository.
type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

// NewMockTodoRepository creates a new mock instance.
func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockTodoRepository) Get(id string) (*domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTodoRepositoryMockRecorder) Get(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTodoRepository)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockTodoRepository) GetAll() ([]*domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockTodoRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTodoRepository)(nil).GetAll))
}

// GetByUserId mocks base method.
func (m *MockTodoRepository) GetByUserId(userId string) ([]*domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserId", userId)
	ret0, _ := ret[0].([]*domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserId indicates an expected call of GetByUserId.
func (mr *MockTodoRepositoryMockRecorder) GetByUserId(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserId", reflect.TypeOf((*MockTodoRepository)(nil).GetByUserId), userId)
}

// Remove mocks base method.
func (m *MockTodoRepository) Remove(todo *domain.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockTodoRepositoryMockRecorder) Remove(todo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockTodoRepository)(nil).Remove), todo)
}

// Save mocks base method.
func (m *MockTodoRepository) Save(todo *domain.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockTodoRepositoryMockRecorder) Save(todo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTodoRepository)(nil).Save), todo)
}

// MockAuthRepository is a mock of AuthRepository interface.
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository.
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance.
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

// CheckPassword mocks base method.
func (m *MockAuthRepository) CheckPassword(user *domain.User, password string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPassword", user, password)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPassword indicates an expected call of CheckPassword.
func (mr *MockAuthRepositoryMockRecorder) CheckPassword(user, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPassword", reflect.TypeOf((*MockAuthRepository)(nil).CheckPassword), user, password)
}

// GeneratePassword mocks base method.
func (m *MockAuthRepository) GeneratePassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeneratePassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeneratePassword indicates an expected call of GeneratePassword.
func (mr *MockAuthRepositoryMockRecorder) GeneratePassword(password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeneratePassword", reflect.TypeOf((*MockAuthRepository)(nil).GeneratePassword), password)
}

// GenerateTokens mocks base method.
func (m *MockAuthRepository) GenerateTokens(user *domain.User, expiresIn time.Duration) (*domain.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTokens", user, expiresIn)
	ret0, _ := ret[0].(*domain.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTokens indicates an expected call of GenerateTokens.
func (mr *MockAuthRepositoryMockRecorder) GenerateTokens(user, expiresIn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTokens", reflect.TypeOf((*MockAuthRepository)(nil).GenerateTokens), user, expiresIn)
}

// GetCurrentUser mocks base method.
func (m *MockAuthRepository) GetCurrentUser(token *domain.Token) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentUser", token)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentUser indicates an expected call of GetCurrentUser.
func (mr *MockAuthRepositoryMockRecorder) GetCurrentUser(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentUser", reflect.TypeOf((*MockAuthRepository)(nil).GetCurrentUser), token)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockUserRepository) Get(id string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserRepositoryMockRecorder) Get(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserRepository)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockUserRepository) GetAll() ([]*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUserRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUserRepository)(nil).GetAll))
}

// GetByUsername mocks base method.
func (m *MockUserRepository) GetByUsername(username string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockUserRepositoryMockRecorder) GetByUsername(username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockUserRepository)(nil).GetByUsername), username)
}

// Remove mocks base method.
func (m *MockUserRepository) Remove(user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockUserRepositoryMockRecorder) Remove(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockUserRepository)(nil).Remove), user)
}

// Save mocks base method.
func (m *MockUserRepository) Save(user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockUserRepositoryMockRecorder) Save(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserRepository)(nil).Save), user)
}

// MockMessageRepository is a mock of MessageRepository interface.
type MockMessageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRepositoryMockRecorder
}

// MockMessageRepositoryMockRecorder is the mock recorder for MockMessageRepository.
type MockMessageRepositoryMockRecorder struct {
	mock *MockMessageRepository
}

// NewMockMessageRepository creates a new mock instance.
func NewMockMessageRepository(ctrl *gomock.Controller) *MockMessageRepository {
	mock := &MockMessageRepository{ctrl: ctrl}
	mock.recorder = &MockMessageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageRepository) EXPECT() *MockMessageRepositoryMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockMessageRepository) SendMessage(messageType domain.MessageType, subject string, to []string, args ...any) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []any{messageType, subject, to}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessage", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockMessageRepositoryMockRecorder) SendMessage(messageType, subject, to any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{messageType, subject, to}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockMessageRepository)(nil).SendMessage), varargs...)
}
