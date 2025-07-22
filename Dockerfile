FROM golang:1.19.0-alpine AS builder

LABEL stage=server-intermediate

WORKDIR /go/src/github.com/infobloxopen/protoc-gen-atlas-query-validate
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/usr/bin/protoc-gen-atlas-query-validate main.go

FROM infoblox/atlas-gentool:latest AS runner

COPY --from=builder /out/usr/bin/protoc-gen-atlas-query-validate /usr/bin/protoc-gen-atlas-query-validate
COPY --from=builder /go/src/github.com/infobloxopen/protoc-gen-atlas-query-validate/options/*.proto /go/src/github.com/infobloxopen/protoc-gen-atlas-query-validate/options/

WORKDIR /go/src
