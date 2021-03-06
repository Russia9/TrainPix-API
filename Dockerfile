FROM alpine:3.11.3 AS build
RUN apk update
RUN apk upgrade
RUN apk add --no-cache go
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=1 GOOS=linux go build trainpix-api

FROM alpine:3.11.3
LABEL maintainer="russia9@russia9.dev"
WORKDIR /app
COPY --from=build /app/trainpix-api /app/trainpix-api
CMD ["/app/trainpix-api"]
