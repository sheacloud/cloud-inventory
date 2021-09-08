FROM golang:1.17-alpine as build

COPY ./ /build

RUN cd /build/; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOBIN=/bin/ go install ./cmd/inventory-scraper/;

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
ENV PATH=/bin
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /bin/inventory-scraper /bin/inventory-scraper

CMD ["/bin/inventory-scraper"]