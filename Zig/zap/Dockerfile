FROM debian:bookworm-slim AS builder

# Install dependencies
RUN apt update && \
    apt install -y --no-install-recommends \
        curl \
        xz-utils \
        bash \
        build-essential \
        linux-headers-amd64 \
        ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Install Zig
# Install Zig
RUN curl -L -o zig.tar.xz https://ziglang.org/download/0.14.1/zig-x86_64-linux-0.14.1.tar.xz && \
    tar -xJf zig.tar.xz && \
    mv zig-x86_64-linux-0.14.1 /opt/zig && \
    ln -s /opt/zig/zig /usr/local/bin/zig && \
    rm zig.tar.xz

RUN zig version

# Set working directory
WORKDIR /app

# Copy source files
COPY . .

RUN zig fetch --save "git+https://github.com/zigzap/zap#v0.10.1"

# Build the application
RUN zig build -Doptimize=ReleaseFast -Dcpu=baseline

# Runtime image
FROM debian:bookworm-slim

# Install runtime dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends libc6 && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary from builder
COPY --from=builder /app/zig-out/bin/zap-resttest /usr/local/bin/zap-resttest

# Expose port
EXPOSE 8080

# Run the application
CMD ["zap-resttest"]