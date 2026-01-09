FROM golang:1.25

WORKDIR /workspace

# Install common tools
RUN apt-get update && apt-get install -y \
    git \
    curl \
    vim \
    && rm -rf /var/lib/apt/lists/*

CMD ["sleep", "infinity"]
