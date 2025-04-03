FROM debian:stable-slim

COPY driven /bin/driven

CMD ["/bin/driven"]
