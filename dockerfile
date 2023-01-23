FROM golang:1.16-alpine AS build

WORKDIR /app

# download the required Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
#COPY *.go ./
COPY . ./

RUN ls

RUN go build -o /main .

EXPOSE 8080

CMD [ "/main" ]