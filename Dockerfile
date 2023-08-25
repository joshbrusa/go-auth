FROM golang AS build

WORKDIR /code

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -C ./src -o ../bin/go-auth

FROM alpine

COPY --from=build /code/bin/go-auth .

CMD [ "/go-auth" ]