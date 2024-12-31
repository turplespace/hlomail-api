package models

// EmailRequest represents the structure of the contact form data
type EmailRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Number  string `json:"number"`
	Message string `json:"message"`
}
