package templates

var Dockerfile = `# Code generated by {{ .Generator }} {{ .Version }} ({{ .Build }}); DO NOT EDIT

FROM golang:1.15.2-alpine3.12 AS compile
WORKDIR /go/src/app
COPY . .
ENV CGO_ENABLED 0
RUN go build -ldflags "-s -w -X=main.build={{ .Build }} -X=main.name={{ .Name }} -X=main.version={{ .Version }}" -o /go/src/app/app main.go

FROM scratch
COPY --from=compile /go/src/app/app /app
COPY --from=compile /go/src/app/web/static /web/static
CMD ["/app"]
`
