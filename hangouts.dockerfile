FROM golang:latest

# Install Task
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d
# Install bun
RUN apt-get update && apt-get install -y zip && apt-get clean && rm -rf /var/lib/apt/lists/*
RUN curl -fsSL https://bun.sh/install | bash
ENV PATH="/root/.bun/bin:$PATH"

# Copy the backend into the container
COPY ./hangouts/ ./hangouts/
COPY ./api-docs/ ./api-docs/
COPY ./openapi.json ./openapi.json

# Set the container's working directory
WORKDIR /go/hangouts/

# Build and run the backend server
CMD ["sh", "-c", "task build --force && task run"]
