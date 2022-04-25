FROM --platform=${BUILDPLATFORM} golang:1.17 as build

COPY ./ /build

ARG TARGETOS
ARG TARGETARCH
RUN cd /build/; CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /bin/fetch-inventory ./cmd/fetch-inventory/;

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
ENV PATH=/bin
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /bin/fetch-inventory /bin/fetch-inventory

CMD ["/bin/fetch-inventory"]