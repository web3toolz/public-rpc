FROM --platform=linux/amd64 golang:1.21-alpine3.18 as builder

WORKDIR /home/app

COPY . /home/app

RUN go mod download
RUN go build -o /public-rpc ./cmd/cli


FROM --platform=linux/amd64 alpine:3.18.4

RUN mkdir /home/app
WORKDIR /home/app

RUN apk --no-cache add gcompat tini
COPY --from=builder /public-rpc /home/app/public-rpc

# expose default port
EXPOSE 8000

ENTRYPOINT ["/sbin/tini", "--"]
CMD [ "/home/app/public-rpc", "run" ]