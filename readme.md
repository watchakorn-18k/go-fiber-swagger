<div align="center">

# go-fiber-swagger

![img](docs/image.png)

</div>

## Install golang package

```bash
go mod tidy
```

# Set Environment `.env`
```
DATABASE_NAME= <database in mongodb>
MONGODB_URI= <database in mongodb>
PORT=1818

JWT_SECRET_KEY=Test
JWT_REFESH_SECRET_KEY=Test
```

# Start APP

```sh
go run . || go run main.go
```
- go to (localhost:1818/api/users/docs)


## Air

```sh
go install github.com/cosmtrek/air@latest
```

```sh
air
```

# Podman

```
podman build -t fiber-test .
```

```
podman run --rm -it -p 3000:3000 fiber-test
```

# Swagger
- can edit in [docs/swagger.json](docs/swagger.json)
