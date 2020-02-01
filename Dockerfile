FROM alpine
LABEL maintainer="russia9@russia9.dev"
WORKDIR /app
COPY . .
RUN apk add --no-cache go
RUN CGO_ENABLED=1 go build trainpix-api
CMD ./trainpix-api

