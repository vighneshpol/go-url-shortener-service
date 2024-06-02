# Go URL Shortener Service

    A simple URL shortener service written in Go. This project provides a RESTful API for shortening URLs, retrieving the original URLs, and tracking the top domains that have been shortened the most. It features a clean code architecture, comprehensive unit tests, and includes a metrics API to monitor the most frequently shortened domains.

## Features

- Shorten long URLs into compact, shareable links.
- Retrieve the original URL using the shortened code.
- Metrics API to track and display the top 3 most frequently shortened domains.
- Easy-to-read code with clear structure and descriptive names.
- Comprehensive unit tests to ensure code reliability.

## Technologies Used

- Go programming language
- Standard library `net/http` for handling HTTP requests
- Docker for containerization

## Prerequisites

- Go 1.14 or higher installed on your system
- Docker installed (optional, for containerization)

## Endpoints

1. Shorten URL: POST /shorten

    Request Body: {"url": "YOUR_LONG_URL"}

    Response: {"shortened_url": "SHORTENED_URL"}

2. Redirect to Original URL: GET /SHORT_CODE

    Redirects to the original URL associated with the short code.

3. Metrics API: GET /metrics

    Retrieves the top 3 domains with the highest counts.

## Installation locally

1. Clone the repository:

    `git clone https://github.com/vighneshpol/go-url-shortener-service.git`

2. Navigate to the project directory:

    `cd go-url-shortener-service`

3. Initialize Go Module(Optional as i have push go.mod already)

    `go mod init`

4. Run the main file

    `go run main.go`

5. The server will start running on

    `http://localhost:8080`

follow below testing section for the testing purpose and to see metrics

### Testing:

You can test the program by making requests using the terminal. Here are some examples:

- To shorten a URL: Run the following command in your terminal:

    `Invoke-RestMethod -Method POST -Uri "http://localhost:8080/shorten" -ContentType "application/json" -Body '{"url": "https://www.udemy.com/"}'`

This command will shorten the URL "https://www.udemy.com/" and return a shortened link.

- To check metrics of shortened URLs: Use the following command:

    `Invoke-RestMethod -Method GET -Uri "http://localhost:8080/metrics"`

This command retrieves metrics about the most frequently shortened URLs, including the top three domains.

### Using Docker (Optional, for easier Setup):

If you prefer using Docker, follow these steps:

1. Pull the Docker image (if not built locally):

    `docker pull vigupol/golangproject`

2. Run the Docker container:

    `docker run -p 8080:8080 vigupol/golangproject:latest`

3. To check if the container is running, execute parallely on the new terminal window:

    `docker ps`

Now you can use the program the same way as before.

---

With these instructions, you should be able to use the URL Shortener Service without any problems. If you have any questions or need help, feel free to reach out!

