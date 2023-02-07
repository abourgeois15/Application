FROM golang:1.16-alpine

WORKDIR /api

COPY api.exe ./
COPY conf.yaml ./

EXPOSE 8080

CMD ["./api.exe"]