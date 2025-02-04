# Build

FROM ubuntu:24.04 AS builder

# build dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    ninja-build \
    git \
    curl \
    zip \
    unzip \
    tar \
    pkg-config \
    libssl-dev \
    && rm -rf /var/lib/apt/lists/*

# vcpkg for c++ packages
RUN git clone https://github.com/Microsoft/vcpkg.git /vcpkg \
    && /vcpkg/bootstrap-vcpkg.sh

# install c++ packages
RUN /vcpkg/vcpkg install crow:x64-linux-dynamic fmt:x64-linux-dynamic

# copy source
WORKDIR /app
COPY . .

# build project
RUN cmake -B build -S . \
    -DCMAKE_TOOLCHAIN_FILE=/vcpkg/scripts/buildsystems/vcpkg.cmake \
    -DCMAKE_BUILD_TYPE=Release \
    && cmake --build build


# Runtime
FROM ubuntu:24.04

# install runtime dependencies
RUN apt-get update && apt-get install -y \
    libssl3 \
    && rm -rf /var/lib/apt/lists/*

# create non-root user
RUN useradd -m -s /bin/bash appuser

# create directories for assets and config
RUN mkdir -p /app/assets /app/config \
    && chown -R appuser:appuser /app

# copy built executable
COPY --from=builder /app/build/CrowAPI /app/
COPY --chown=appuser:appuser assets/ /app/assets

USER appuser
WORKDIR /app

EXPOSE 8000

CMD ["./CrowAPI"]