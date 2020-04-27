FROM alpine
WORKDIR /app
COPY server .env ./
EXPOSE 8100
ENTRYPOINT /app/server