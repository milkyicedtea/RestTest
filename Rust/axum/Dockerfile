# Build
FROM lukemathwalker/cargo-chef:latest-rust-1 AS chef
WORKDIR /app

# Prepare recipe
FROM chef AS planner
COPY . .
RUN cargo chef prepare --recipe-path recipe.json

# Build dependencies
FROM chef AS builder
COPY --from=planner /app/recipe.json recipe.json
RUN cargo chef cook --release --recipe-path recipe.json

# Build application
COPY . .
RUN cargo build --release && cp target/release/axum-resttest /app/axum-resttest

# Create smaller image for exec
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary and .env to runtime image
COPY --from=builder /app/axum-resttest .
COPY .env .

RUN chmod +x ./axum-resttest

EXPOSE 8080

CMD ["./axum-resttest"]