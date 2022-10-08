FROM golang:1.18-alpine

WORKDIR /view

COPY ./view/public/ public/
COPY ./view/templates/ templates/
COPY ./view/main.go .
COPY ./view/go.mod .
COPY ./view/go.sum .
#RUN go mod tidy

# COPY view/ view/
RUN go mod download


RUN go build -o /frontend main.go

EXPOSE 3000
CMD ["/frontend"]