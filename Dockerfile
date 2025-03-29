
# Go bazaviy image
FROM golang:1.21-alpine

# Ishchi katalogni o‘rnatish
WORKDIR /app

# Modul fayllarni nusxalash va o‘rnatish
COPY go.mod go.sum ./
RUN go mod tidy

# Loyihani nusxalash va build qilish
COPY . .
RUN go build -o main ./cmd/main.go

# Portni ochish
EXPOSE 8080

# Serverni ishga tushirish
CMD ["./main"]

