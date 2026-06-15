param (
    $command
)

if (-not $command)  {
    $command = "start"
}

$ProjectRoot = "${PSScriptRoot}/.."

$env:AMBULANCE_API_ENVIRONMENT="Development"
$env:AMBULANCE_API_PORT="8080"
$env:AMBULANCE_API_DB_USERNAME="root"
$env:AMBULANCE_API_DB_PASSWORD="neUhaDnes"

function postgres {
    docker compose --file ${ProjectRoot}/deployments/docker-compose/compose.yaml $args
}



switch ($command) {
    "start" {
        try{
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
        docker run --rm -ti -v ${ProjectRoot}:/local openapitools/openapi-generator-cli generate -c /local/scripts/generator-cfg.yaml
    }
    "docker" {
         docker build -t mariia610/ambulance-webapi:local-build -f ${ProjectRoot}/build/docker/Dockerfile .
   }
    default {
        throw "Unknown command: $command"
    }
}