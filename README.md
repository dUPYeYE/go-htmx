# GO Backend API with HTMX and TailwindCSS
Boilerplate for a GO backend API with HTMX and TailwindCSS

## Features
- GO Backend API
- HTMX
- TailwindCSS
- Docker

## Instructions
1. Clone the repository
2. For local use:
    - Populate the `.env` file with the necessary environment variables based on .env.example
    - Run the following command to initialize the database:
        ```bash
        ./scripts/migrate_up.sh
        ```
    - Run the following command to start the server:
        ```bash
        go run .
        ```
    - Access the server at `http://localhost:<PORT>`
3. For Docker use:
    - Populate the `.env` file with the necessary environment variables based on .env.example
    - Run the following command to initialize the database:
        ```bash
        ./scripts/migrate_up.sh
        ```
    - Run the following command to start the server:
        ```bash
        docker-compose up
        ```
    - Access the server at `http://localhost:<PORT>`

## API Endpoints
    `api/users`
    TODO: Add more API endpoints

## Checks
![Tests and Stylechecks](https://github.com/dUPYeYE/go-htmx/actions/workflows/ci.yml/badge.svg)

## Author
- dupp
