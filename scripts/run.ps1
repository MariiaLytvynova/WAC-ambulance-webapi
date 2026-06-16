param (
    $command
)

if (-not $command) {
    $command = "start"
}

$ProjectRoot = "${PSScriptRoot}/.."

# ----------------------------
# Environment variables (LOCAL DEV)
# ----------------------------
$env:AMBULANCE_API_ENVIRONMENT = "Development"
$env:AMBULANCE_API_PORT = "8080"

$env:POSTGRES_HOST = "localhost"
$env:POSTGRES_PORT = "5432"
$env:POSTGRES_USER = "root"
$env:POSTGRES_PASSWORD = "neUhaDnes"
$env:POSTGRES_DB = "ambulance"
$env:POSTGRES_SSLMODE = "disable"

# ----------------------------
# Docker compose helper
# ----------------------------
function postgres {
    docker compose --file ${ProjectRoot}/deployments/docker-compose/compose.yaml $args
}

# ----------------------------
# Commands
# ----------------------------
switch ($command) {

    "start" {
        try {
            postgres up --detach

            go run ${ProjectRoot}/cmd/ambulance-api-service

        } finally {
            postgres down
        }
    }

    "postgres" {
        postgres up
    }

    "openapi" {
        docker run --rm -ti `
            -v ${ProjectRoot}:/local `
            openapitools/openapi-generator-cli `
            generate -c /local/scripts/generator-cfg.yaml
    }

    "docker" {
        docker build `
            -t mariia610/ambulance-webapi:local-build `
            -f ${ProjectRoot}/build/docker/Dockerfile .
    }

    default {
        throw "Unknown command: $command"
    }
}