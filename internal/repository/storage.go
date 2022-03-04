package repository

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"refactoring/internal/model"
	"strconv"
	"time"
)

type User struct {
	CreatedAt   time.Time `json:"created_at"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
}

type UserList map[string]User

type UserStore struct {
	filename  string   `json:"-"`
	Increment int      `json:"increment"`
	List      UserList `json:"list"`
}

func New(filename string) (*UserStore, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, model.FileNotFound
	}
	s := new(UserStore)
	_ = json.Unmarshal(f, s)
	s.filename = filename

	return s, nil
}

func (s *UserStore) save() {
	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(s.filename, b, fs.ModePerm)
}

func (s *UserStore) SearchUsers() *UserList {
	return &s.List
}

func (s *UserStore) GetByID(id string) (User, error) {
	if _, ok := s.List[id]; !ok {
		return User{}, model.UserNotFound
	}

	return s.List[id], nil
}

func (s *UserStore) UpdateUser(id string, request model.UpdateUserRequest) error {
	user, err := s.GetByID(id)
	if err != nil {
		return err
	}

	user.DisplayName = request.DisplayName
	s.List[id] = user

	s.save()
	return nil
}

func (s *UserStore) CreateUser(request model.CreateUserRequest) (string, error) {
	s.Increment++
	user := User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}

	id := strconv.Itoa(s.Increment)
	s.List[id] = user

	s.save()
	return id, nil
}

func (s *UserStore) DeleteUser(id string) error {
	if _, ok := s.List[id]; !ok {
		return model.UserNotFound
	}

	delete(s.List, id)
	s.save()

	return nil
}
