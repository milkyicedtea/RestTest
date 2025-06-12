# Build

FROM debian:bookworm-slim AS builder
WORKDIR /app

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
RUN /vcpkg/vcpkg install crow:x64-linux-dynamic fmt:x64-linux-dynamic glaze:x64-linux-dynamic

# copy source
COPY . .

# build project
RUN cmake -B build -S . \
    -DCMAKE_TOOLCHAIN_FILE=/vcpkg/scripts/buildsystems/vcpkg.cmake \
    -DCMAKE_BUILD_TYPE=Release \
    -DVCPKG_TARGET_TRIPLE=x64-linux-dynamic \
    -DIN_DOCKER=ON \
    && cmake --build build


# Runtime
FROM debian:bookworm-slim

# install runtime dependencies
RUN apt-get update && apt-get install -y \
    libssl3 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app


# copy built executable and deps
COPY --from=builder /app/build/crow-resttest .

RUN chmod +x ./crow-resttest

EXPOSE 8000

CMD ["./crow-resttest"]