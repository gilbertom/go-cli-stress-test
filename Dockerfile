# Use a imagem base do Go
FROM golang:1.22.2

# Define o diretório de trabalho
WORKDIR /app

# Copie o código fonte para o container
COPY . .

# Baixe as dependências e compile o código
RUN go mod download

# Compile o binário Go
RUN go build -o stress-test-bin .

# Comando para executar o aplicativo
CMD ["./stress-test-bin"]
