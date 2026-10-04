FROM docker.io/library/golang:1.27-bookworm

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        git \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /workspace

ENV CGO_ENABLED=0

CMD ["bash"]