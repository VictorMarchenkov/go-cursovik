FROM golang:1.18-alpine

WORKDIR /app

#COPY ./app/database/ database/
#COPY ./app/model/ model/
#COPY ./app/server/ server/
#COPY ./app/utils/ utils/

COPY ./app ./

#COPY ./app/.env .
#COPY ./app/main.go .

#COPY ./app/go.mod .
#COPY ./app/go.sum .

COPY ./app/go.* .
COPY ./app/ .

RUN go mod tidy

 # COPY view/ view/
 RUN go mod download


RUN go build -o /server main.go

EXPOSE 8080
CMD ["/server"]