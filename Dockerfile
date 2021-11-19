FROM golang AS build
WORKDIR /app
COPY go.mod go.sum *.go ./
RUN go mod download
RUN go test -v
RUN CGO_ENABLED=0 GOOS=linux go build -o /greeter

FROM scratch
EXPOSE 8080
COPY --from=build /greeter /greeter
ENTRYPOINT ["/greeter"]