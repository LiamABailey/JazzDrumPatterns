FROM golang:1.18-alpine

WORKDIR "jdp"

COPY go.mod ./
COPY go.sum ./
COPY api ./api
COPY assets/beatimages/kick-hihat ./assets/beatimages/kick-hihat
COPY assets/beatimages/ride-snare ./assets/beatimages/ride-snare
COPY assets/beatimages/assetembedding.go ./assets/beatimages/assetembedding.go
COPY assets/beatimages/go.mod ./assets/beatimages/go.mod
COPY cmd ./cmd 
COPY internal ./internal

WORKDIR ./cmd 
RUN go build 

EXPOSE 8050
CMD ["./cmd"]