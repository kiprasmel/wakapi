package mocks

import (
	"github.com/muety/wakapi/models"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) GetUserById(s string) (*models.User, error) {
	args := m.Called(s)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserServiceMock) GetUserByKey(s string) (*models.User, error) {
	args := m.Called(s)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserServiceMock) GetUserByEmail(s string) (*models.User, error) {
	args := m.Called(s)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserServiceMock) GetUserByResetToken(s string) (*models.User, error) {
	args := m.Called(s)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserServiceMock) GetAll() ([]*models.User, error) {
	args := m.Called()
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *UserServiceMock) GetAllMapped() (map[string]*models.User, error) {
	args := m.Called()
	return args.Get(0).(map[string]*models.User), args.Error(1)
}

func (m *UserServiceMock) GetMany(s []string) ([]*models.User, error) {
	args := m.Called(s)
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *UserServiceMock) GetManyMapped(s []string) (map[string]*models.User, error) {
	args := m.Called()
	return args.Get(0).(map[string]*models.User), args.Error(1)
}

func (m *UserServiceMock) GetAllByLeaderboard(b bool) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *UserServiceMock) GetAllByReports(b bool) ([]*models.User, error) {
	args := m.Called(b)
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *UserServiceMock) GetUserByStripeCustomerId(s string) (*models.User, error) {
	args := m.Called(s)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserServiceMock) GetActive(b bool) ([]*models.User, error) {
	args := m.Called(b)
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *UserServiceMock) Count() (int64, error) {
	args := m.Called()
	return int64(args.Int(0)), args.Error(1)
}

func (m *UserServiceMock) CountCurrentlyOnline() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *UserServiceMock) CreateOrGet(signup *models.Signup, isAdmin bool) (*models.User, bool, error) {
	args := m.Called(signup, isAdmin)
	return args.Get(0).(*models.User), args.Bool(1), args.Error(2)
}

func (m *UserServiceMock) Update(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserServiceMock) Delete(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserServiceMock) ResetApiKey(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserServiceMock) ToggleBadges(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserServiceMock) SetWakatimeApiCredentials(user *models.User, s1, s2 string) (*models.User, error) {
	args := m.Called(user, s1, s2)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserServiceMock) GenerateResetToken(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserServiceMock) ChangeUserId(user *models.User, s1 string) (*models.User, error) {
	args := m.Called(user, s1)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserServiceMock) FlushCache() {
	m.Called()
}

func (m *UserServiceMock) FlushUserCache(s string) {
	m.Called(s)
}
