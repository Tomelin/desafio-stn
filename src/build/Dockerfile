###############################
##### IMAGE TO GOLANG BUILD ###
FROM  golang:1.22-alpine AS build

LABEL org.opencontainers.image.authors="RAFAEL TOMELIN"
LABEL name="Rafael Tomelin"
LABEL email="[rafael.tomelin@gmail.com]"

WORKDIR /app

COPY ./src /app/src

ENV CGO_ENABLED=0

RUN cd /app/src ; go build -ldflags='-s -w' -o /app/kube /app/src/cmd/main.go

RUN adduser --disabled-password --no-create-home -u 1001 nonroot

# # ###############################
# # ##### IMAGE EXECUTE BINARIO ###
FROM scratch

LABEL org.opencontainers.image.authors="RAFAEL TOMELIN"
LABEL name="Rafael Tomelin"
LABEL email="[rafael.tomelin@gmail.com]"

WORKDIR /src

COPY --from=build /etc/passwd /etc/passwd

USER nonroot
COPY --chmod=0755 --from=build /app/kube /src/

CMD ["/src/kube"]


