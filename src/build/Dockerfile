###############################
##### IMAGE TO GOLANG BUILD ###
FROM  golang:1.22-alpine AS build

LABEL org.opencontainers.image.authors="RAFAEL TOMELIN"
LABEL name="Rafael Tomelin"
LABEL email="[rafael.tomelin@gmail.com]"

WORKDIR /app

RUN adduser --disabled-password --no-create-home -u 1001 nonroot

ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY ./src/go.mod .
COPY ./src/go.sum .

RUN go mod download

COPY ./src /app/

RUN go build -ldflags='-s -w' -o /app/kube /app/cmd/main.go

# # ###############################
# # ##### IMAGE EXECUTE BINARIO ###
FROM scratch

LABEL org.opencontainers.image.authors="RAFAEL TOMELIN"
LABEL name="Rafael Tomelin"
LABEL email="[rafael.tomelin@gmail.com]"

WORKDIR /app

COPY --from=build /etc/passwd /etc/passwd

USER nonroot
COPY --chmod=0755 --from=build /app/kube /app/

CMD ["/app/kube"]
