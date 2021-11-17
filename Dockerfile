FROM golang AS build
WORKDIR ${GOPATH}
COPY greeter.go src/
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/greeter src/greeter.go

FROM scratch
EXPOSE 8080
COPY --from=build /go/bin/greeter /greeter
ENTRYPOINT ["/greeterrrr"]