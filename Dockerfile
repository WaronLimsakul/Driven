FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

COPY driven /usr/bin/driven
COPY static /static

CMD ["driven"]
