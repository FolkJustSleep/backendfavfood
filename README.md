# Get started

```bash
go mod tidy
```

# Set Environtment

```
PORT= <yourPort>

DB_HOST=<your-db-host>
DB_PORT=<your-db-port>
DB_USER=<your-db-user>
DB_PASSWORD=<your-db-password>
DB_NAME=<your-db-name>

JWT_SECRET=<your-jwt-secret>
```

# Run App

Using swagger first:

If using macOS

```
export PATH=$PATH:$HOME/go/bin
swag init
```

If using other os you can just:

```
swag init
```

And then start the app

```
go run .
```

# Docker Run

Build Docker image

```bash
docker-compose up --build
```

Press ctrl c or

```bash
docker-compose down
```
To stopped docker
