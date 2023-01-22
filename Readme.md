# Go Fiber Blog Server

A simple Go blog server using the Fiber web framework and GORM as an ORM.

## Prerequisites

- Go 1.15+
- PostgreSQL
- Air (for development)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/go-fiber-server.git
```

2. Install the dependencies:

```bash
go mod download
```

3. Create a `.env` file in the root of the project and set the following environment variables:

```md
DB_USER=<your_db_user>
DB_PASSWORD=<your_db_password>
DB_NAME=<your_db_name>
DB_PORT=<your_db_port>
DB_HOST=<your_db_ip(localhost)>
PORT=<server_port>
JWT_SECRET=<jwt_secret_string>
```

4. Run the migrations:

```bash
go run cmd/main/main.go
```

## Development

For development, we recommend using [Air](https://github.com/cosmtrek/air) to automatically reload the server on file changes.

1. Install Air:

```bash
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

or

```bash
go install github.com/cosmtrek/air@latest
```

2. Run the server with Air:

```bash
air
```

## Usage

Make requests under `http://localhost:<port>/api/v1/<some_path>` to the API endpoints defined in the `routes` package.

## Database

This server uses GORM as an ORM. You can find the models in the `models` package and the database configuration in `database.go`.

## Contributions

All contributions are welcome. Please open an issue or a pull request if you want to suggest a change or fix a bug.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

This project uses the [Fiber](https://gofiber.io/) web framework and [GORM](https://gorm.io/) as an ORM.
