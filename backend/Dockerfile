FROM golang:1.21-bullseye as build

RUN apt update && apt install -y libmagickwand-dev \
&& rm -rf /var/cache/apt

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app /app/
RUN go build -o db /app/cmd/db/

FROM debian:bullseye-slim as deploy

WORKDIR /app

RUN apt update && apt install -y libmagickwand-dev \
&& rm -rf /var/cache/apt

RUN sed -i -E 's/rights="none" pattern="PDF"/rights="read|write" pattern="PDF"/g' /etc/ImageMagick-6/policy.xml

COPY --from=build /app/app /app/app
COPY --from=build /app/db /app/db
COPY migration /app/migration
COPY scripts /app/scripts
COPY configs /app/configs

RUN chmod +x /app/scripts/entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/app/scripts/entrypoint.sh"]