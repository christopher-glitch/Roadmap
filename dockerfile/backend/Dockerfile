FROM golang:latest

WORKDIR /app/go

RUN go install github.com/cosmtrek/air@latest && \
	go install golang.org/x/tools/gopls@latest && \
	go install github.com/ramya-rao-a/go-outline@latest && \
	go install golang.org/x/tools/cmd/goimports@latest && \
	go install github.com/josharian/impl@latest && \
	go install github.com/go-delve/delve/cmd/dlv@latest && \
	go install github.com/cweill/gotests/...@latest && \
	go install github.com/haya14busa/goplay/cmd/goplay@latest && \
	go install github.com/fatih/gomodifytags@latest && \
	go install go.uber.org/mock/mockgen@latest && \
	go install github.com/onsi/ginkgo/v2/ginkgo@latest && \
	go install github.com/onsi/gomega/...@latest && \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2

COPY ./ ./


RUN go mod tidy 
EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
