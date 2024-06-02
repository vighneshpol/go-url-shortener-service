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


## Usage
1. Initialize Go Module(Optional as i have push go.mod already)
`go mod init`

2. Run the main file
`go run main.go`

3. The server will start running on
`http://localhost:8080`

4. Now to test the URL go to terminal and try to hit different POST methods 
`Invoke-RestMethod -Method POST -Uri "http://localhost:8080/shorten" -ContentType "application/json" -Body '{"url": "https://www.udemy.com/"}'`
`Invoke-RestMethod -Method POST -Uri "http://localhost:8080/shorten" -ContentType "application/json" -Body '{"url": "https://www.udemy.com/pricing/"}'`

5. Now to check metrics use the below GET command 
`Invoke-RestMethod -Method GET -Uri "http://localhost:8080/metrics"`

## To run the application using Docker:

1. Pull the Docker image (if not built locally):
    `docker pull <registry_url>/<image_name>`

2. Build the Docker image (if not pulled):
    `docker build -t go-url-shortener-service .`

3. Run the Docker container:
    `docker run -p 8080:8080 go-url-shortener-service`

4. To check if the container is running, execute parallely on the new terminal window:
    `docker ps`

5. Now to test the URL go to terminal and try to hit different POST methods 
    `Invoke-RestMethod -Method POST -Uri "http://localhost:8080/shorten" -ContentType "application/json" -Body '{"url": "https://www.udemy.com/"}'`
    `Invoke-RestMethod -Method POST -Uri "http://localhost:8080/shorten" -ContentType "application/json" -Body '{"url": "https://www.udemy.com/pricing/"}'`

5. Now to check metrics use the below GET command 
    `Invoke-RestMethod -Method GET -Uri "http://localhost:8080/metrics"`
