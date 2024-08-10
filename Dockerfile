FROM golang:alpine AS build-stage

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /sso-jwt

FROM gcr.io/distroless/static-debian12:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /sso-jwt /sso-jwt
COPY --from=build-stage /app/api.config api.config
COPY --from=build-stage /app/privateKey privateKey

ENTRYPOINT [ "/sso-jwt" ]