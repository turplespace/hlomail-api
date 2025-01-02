package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/turplespace/hlomailapi/internal/models"
	"github.com/turplespace/hlomailapi/internal/services"
)

// ContactFormHandler handles contact form submissions
func HandlePostContactFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req models.EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" || req.Number == "" || req.Message == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	to := os.Getenv("CONTACT_FORM_RECIPIENT")
	subject := "Contact From : " + req.Name
	body := "You have received a new message from your website's contact form:\n\n" +
		"Name: " + req.Name + "\n" +
		"Email: " + req.Email + "\n\n" +
		"Number: " + req.Number + "\n\n" +
		"Message:\n" + req.Message

	if err := services.SendEmail(to, subject, body); err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		log.Println("Failed to send email:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Message sent successfully",
	})
}
