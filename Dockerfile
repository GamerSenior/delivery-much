################################
# Parte 1 - Build do projeto
################################
FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# Copia e baixa dependencias do projeto
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copia o código para o container
COPY . .

# Realiza build da aplicação
RUN go build -o main .

WORKDIR /dist

# Copia binário do build para a pasta de destribuição
RUN cp /build/main .

#################################
# PARTE 2 - Controi imagem menor 
#################################
FROM scratch

COPY --from=builder /dist/main /

ENTRYPOINT ["/main"]
