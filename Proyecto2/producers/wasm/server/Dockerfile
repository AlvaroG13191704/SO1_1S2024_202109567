# Stage 1: Build the Rust application
FROM ubuntu:20.04 as builder

# Set noninteractive installation
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=America

# Install dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    openssl libssl-dev \
    pkg-config \
    python \
    valgrind \
    zlib1g-dev

# Install Rust and cargo
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y

# Add cargo to PATH
ENV PATH="/root/.cargo/bin:${PATH}"

# Install cmake
RUN apt-get update && apt-get install -y cmake

# Copy your Rust project code
COPY . /app

# Build your Rust API
WORKDIR /app
RUN cargo build --release

# Stage 2: Build the runtime image
FROM ubuntu:20.04

# Set noninteractive installation
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=America

# Install runtime dependencies
RUN apt-get update && apt-get install -y \
    openssl libssl-dev \
    zlib1g

# Copy the built application from the builder stage
COPY --from=builder /app/target/release/api /app/api

# Set the command to run your application
CMD ["/app/api"]
