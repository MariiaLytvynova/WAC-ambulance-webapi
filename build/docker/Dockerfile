# use specific versions of images
FROM openapitools/openapi-generator-cli:latest AS api

WORKDIR /local

COPY api api
COPY scripts scripts
COPY .openapi-generator-ignore .openapi-generator-ignore

RUN docker-entrypoint.sh generate -c /local/scripts/generator-cfg.yaml

# not used normally but redefine entrypoint for the case of checking this stage results
ENTRYPOINT ["bash"]

############################################


FROM golang:latest AS build

WORKDIR /app

# download dependencies - low frequency of changes
COPY go.mod go.sum ./
RUN go mod download

# copy sources - higher frequency of changes
COPY internal/ internal/
COPY cmd/ cmd/
COPY --from=api /local/ ./

# ensure tests are passing
RUN go test ./...

# create executable - ambulance-webapi-srv
# we want to use scratch image so setting
# the build options in the way that will link all dependencies statically
RUN CGO_ENABLED=0 GOOS=linux \
      go build \
      -ldflags="-w -s" \
      -installsuffix 'static' \
      -o ./ambulance-webapi-srv ./cmd/ambulance-api-service

############################################

# See: https://github.com/opencontainers/image-spec/blob/main/annotations.md
FROM scratch
LABEL org.opencontainers.image.authors="Your Name"
LABEL org.opencontainers.image.title="Ambulance Waiting List WebAPI Service"
LABEL org.opencontainers.image.description="WEB API for managing ambulance waiting lists"

# Application configuration
ENV AMBULANCE_API_ENVIRONMENT=production
ENV AMBULANCE_API_PORT=8080

# PostgreSQL configuration
ENV POSTGRES_HOST=postgres-db
ENV POSTGRES_PORT=5432
ENV POSTGRES_DB=ambulance
ENV POSTGRES_USER=root
ENV POSTGRES_PASSWORD=
ENV POSTGRES_SSLMODE=disable

# Copy compiled application
COPY --from=build /app/ambulance-webapi-srv ./

# Default application port
EXPOSE 8080

# Start application
ENTRYPOINT ["./ambulance-webapi-srv"]
