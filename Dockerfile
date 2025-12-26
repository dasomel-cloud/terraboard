FROM golang:1.23 AS builder
LABEL maintainer="dasomell@gmail.com"
WORKDIR /opt/build
COPY . .
RUN make build

FROM node:22 AS node-builder
WORKDIR /opt/build
COPY static/terraboard-vuejs ./terraboard-vuejs
WORKDIR /opt/build/terraboard-vuejs
RUN yarn install
RUN yarn run build

FROM scratch
WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /opt/build/terraboard /
COPY --from=node-builder /opt/build/terraboard-vuejs/dist /static
ENTRYPOINT ["/terraboard"]
CMD [""]
