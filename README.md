# HloMailAPI

This is a simple Go backend for handling contact form submissions and sending the form data via email using SMTP.

## Features

- Accepts contact form data via a POST API.
- Validates all required fields.
- Sends the contact form data as an email to a recipient.
- Supports configuration via an `.env` file.

## Prerequisites

- Go 1.18+ installed on your system.
- An SMTP server for sending emails (e.g., Gmail SMTP).

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/contact-form-backend.git
   cd contact-form-backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the project root:
   ```env
   # SMTP Configuration
   SMTP_HOST=smtp.gmail.com
   SMTP_PORT=587
   SMTP_USER=your-email@gmail.com
   SMTP_PASS=your-email-password

   # Contact Form Recipient Email
   CONTACT_FORM_RECIPIENT=recipient@example.com

   # Server Configuration
   PORT=8080

   # Allowed Domains for CORS
   ALLOWED_DOMAINS=https://example.com,https://another-example.com
   ```

   > Replace `your-email@gmail.com`, `your-email-password`, and `recipient@example.com` with your own credentials and recipient email.

4. Run the application:
   ```bash
   go run main.go
   ```

   The server will start on the specified port (default is `8080`).

## API Endpoint

### POST `/api/contact`

Handles contact form submissions and sends the data via email.

#### Request

- **Method**: `POST`
- **Content-Type**: `application/json`
- **Body**:
  ```json
  {
    "name": "John Doe",
    "email": "johndoe@example.com",
    "number": "1234567890",
    "message": "Hello, I am interested in your services."
  }
  ```

#### Response

- **200 OK**: Message sent successfully.
- **400 Bad Request**: Invalid request body or missing fields.
- **500 Internal Server Error**: Failed to send email.

#### Example cURL Request

```bash
curl -X POST http://localhost:8080/api/contact \
-H "Content-Type: application/json" \
-d '{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "number": "1234567890",
  "message": "Hello, I am interested in your services."
}'
```

## Environment Variables

The `.env` file is used to configure the application. Below are the available variables:

| Variable                 | Description                              | Default    |
|--------------------------|------------------------------------------|------------|
| `SMTP_HOST`              | SMTP server host                        | Required   |
| `SMTP_PORT`              | SMTP server port                        | Required   |
| `SMTP_USER`              | SMTP server username                    | Required   |
| `SMTP_PASS`              | SMTP server password                    | Required   |
| `CONTACT_FORM_RECIPIENT` | Email to receive contact form submissions | Required   |
| `PORT`                   | Server port                             | `8080`     |
| `ALLOWED_DOMAINS`        | Comma-separated list of allowed CORS domains | None |

## Logging

The server logs:
- Startup messages, including the port it's running on.
- Errors encountered when processing requests or sending emails.

## Contributing

Contributions are welcome! Please fork this repository and create a pull request for any improvements or bug fixes.

