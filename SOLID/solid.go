package main

import "fmt"

// Single Responsibility Principle:
// Each struct has only a single responsibility.
// User struct represents a user in the system.
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// UserService struct defines a service for managing users.
type UserService struct {
	users []User
}

// AddUser adds a new user to the service.
func (s *UserService) AddUser(u User) {
	s.users = append(s.users, u)
}

// GetUserByID returns the user with the given ID.
func (s *UserService) GetUserByID(id int) User {
	for _, u := range s.users {
		if u.ID == id {
			return u
		}
	}
	return User{}
}

// Open/Closed Principle:
// The UserService is open for extension, but closed for modification.
// We can add new functionality by implementing new interfaces,
// rather than modifying the existing UserService.
// UserRepository defines the interface for a user repository.
type UserRepository interface {
	SaveUser(u User) error
	FindUserByID(id int) (User, error)
}

// UserRepositoryImpl is a concrete implementation of the UserRepository interface.
// It uses the UserService to store and retrieve users.
type UserRepositoryImpl struct {
	userService *UserService
}

// SaveUser saves a user to the repository.
func (r *UserRepositoryImpl) SaveUser(u User) error {
	r.userService.AddUser(u)
	return nil
}

// FindUserByID finds a user with the given ID in the repository.
func (r *UserRepositoryImpl) FindUserByID(id int) (User, error) {
	return r.userService.GetUserByID(id), nil
}

// Liskov Substitution Principle:
// The UserRepositoryImpl should be substitutable for the UserRepository interface.
// This means that we should be able to use either the interface or the concrete implementation
// without knowing which one we are using.
// UserController is a controller for managing users.
// It uses a UserRepository to store and retrieve users.
type UserController struct {
	repository UserRepository
}

// NewUserController creates a new UserController.
func NewUserController(r UserRepository) *UserController {
	return &UserController{repository: r}
}

// CreateUser creates a new user.
func (c *UserController) CreateUser(u User) error {
	return c.repository.SaveUser(u)
}

// GetUserByID gets the user with the given ID.
func (c *UserController) GetUserByID(id int) (User, error) {
	return c.repository.FindUserByID(id)
}
func solid() {
	// Dependency Inversion Principle:
	// The UserController depends on the UserRepository interface,
	// rather than on the concrete UserRepositoryImpl.
	// This allows us to use any implementation of the UserRepository interface with the UserController.

	userService := &UserService{}
	repository := &UserRepositoryImpl{userService: userService}
	controller := NewUserController(repository)
	user := User{ID: 1, FirstName: "John", LastName: "Doe"}
	controller.CreateUser(user)
	retrievedUser, _ := controller.GetUserByID(1)
	fmt.Println(retrievedUser)
}
