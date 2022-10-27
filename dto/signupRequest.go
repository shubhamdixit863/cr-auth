package dto

import "cr-auth/domain"

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (s *SignupRequest) ToDomain() *domain.User {

	return &domain.User{
		Username: s.Username,
		Email:    s.Email,
		Password: s.Password,
		Name:     s.Name,
	}

}
