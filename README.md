# go-bsb-service

The application runs as an HTTP server at port 8080. It provides the following RESTful endpoints:

* `GET: /status`: health check endpoint which returns 200 - Status OK

## Development environment - Visual Studio Code

1. Install `Go (By Go Team @ Google)` extention.
2. Open `Command Palette` > select `Go: Install/Update Tools` > select `dlv` and click on `OK`.
3. Open `Command Palette` > select `Debug: Open launch.json`, and use this configuration to get started.
4. Thats it. You can now add a breakpoint and start debugging.

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Run",
      "type": "go",
      "request": "launch",
      "port": 8080,
      "mode": "auto",
      "program": "${workspaceFolder}/main.go",
      "env": {
        "SERVICE_NAME": "go-bsb-service",
        "STAGE_NAME": "dev",
        "SERVER_PORT": 8080,
        "VERSION": "v1",
        "LOG_LEVEL": "debug"
      },
      "args": []
    },
    {
      "name": "Debug",
      "type": "go",
      "request": "launch",
      "port": 8080,
      "mode": "debug",
      "program": "${workspaceFolder}/main.go",
      "env": {
        "SERVICE_NAME": "go-bsb-service",
        "STAGE_NAME": "dev",
        "SERVER_PORT": 8080,
        "VERSION": "v1",
        "LOG_LEVEL": "debug"
      },
      "args": []
    }
  ]
}
```
## Common commands

Run unit tests - `make test`

Compile the service - `make build`

Lint the service - `make lint`

[TODO] Run service with env config - `CONFIG=prod make run`
