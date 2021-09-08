FROM golang:1.17-alpine as build

COPY ./ /build

RUN cd /build/; CGO_ENABLED=1 GOBIN=/bin/ go install ./cmd/inventory-scraper/;

FROM alpine as prod

COPY --from=build /bin/inventory-scraper /bin/inventory-scraper

CMD ["/bin/inventory-scraper"]