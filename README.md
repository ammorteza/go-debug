# go-debug
Debugging golang code on docker container with VSCode (delve)

## Requirements
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Instalation
### Clone
Cloning go-debug repository.
```shell
clone https://github.com/ammorteza/go-debug.git
```

### Configure debugging with VSCode
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "remotePath": "/app",
            "port": 2345,
            "showLog": true,
            "trace": "verbose"
          }
    ]
}
```

### Build
Build and run go-debug services with Docker Compose.
```shell 
docker-compose up
```
After running `docker compose up`, make sure all containers run correctly.

## Running
### API Endpoint
Send request in development mode.
> http://localhost:8090/rest/user
```json
{
	"user" : "Morteza Amzajerdi"
}
```


Send request in debugging mode.
> http://localhost:9999/rest/user
```json
{
	"user" : "Morteza Amzajerdi"
}
```
