# see https://hub.docker.com/_/golang/
FROM golang:1.17-alpine
# define homework
WORKDIR /app
# define default environments variables
ENV DB_HOST=localhost \
    DB_USER=postgres \
    DB_PASSWORD=postgres \
    DB_NAME=postgres \
    DB_PORT=5432 \
    DB_SSLMODE=disable \
    VIA_CEP_URL=https://viacep.com.br/ws \
    SECRET=golang
# send module dependencies to container
COPY go.mod ./
COPY go.sum ./
# go get dependencies
RUN go mod download
# copy all file .go to container
COPY *.go ./
# prepare compiled program
RUN go build -o /program
# define entrtpoint of application
CMD [ "/program" ]

