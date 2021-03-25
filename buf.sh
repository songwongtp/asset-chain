# Substitute BIN for your bin directory.
# Substitute VERSION for the current released version.
# Substitute BINARY_NAME for "buf", "protoc-gen-buf-breaking", or "protoc-gen-buf-lint".
BIN="/usr/local/bin" && \
VERSION="0.40.0" && \
BINARY_NAME="protoc-gen-buf-lint" && \
  curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/${BINARY_NAME}-$(uname -s)-$(uname -m)" \
    -o "${BIN}/${BINARY_NAME}" && \
  chmod +x "${BIN}/${BINARY_NAME}"