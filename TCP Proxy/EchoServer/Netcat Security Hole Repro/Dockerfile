FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY "devil_script" .
RUN chmod +x /app/devil_script
EXPOSE 20080

CMD ["/app/devil_script"]