FROM oven/bun:latest

WORKDIR /app

COPY . .

RUN bun i

EXPOSE 8080

CMD ["bun", "src/index.ts"]