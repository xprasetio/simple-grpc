# Example gRPC Service

Proyek ini adalah implementasi layanan  menggunakan gRPC dengan bahasa Go. Terdapat beberapa modul utama seperti chat, user, order, dan product yang di-generate dari file proto.

## Struktur Direktori

- `main.go` : Entry point aplikasi utama.
- `grpcclient/` : Client gRPC untuk mengakses layanan.
- `pb/` : File hasil generate dari proto (protobuf dan gRPC).
  - `chat/` : Protobuf dan gRPC untuk chat.
  - `ecommerce/order/` : Protobuf untuk order.
  - `ecommerce/product/` : Protobuf untuk product.
  - `user/` : Protobuf dan gRPC untuk user.
- `proto/` : File proto untuk mendefinisikan service dan message.
  - `chat/` : Definisi proto chat.
  - `ecommerce/order/` : Definisi proto order.
  - `ecommerce/product/` : Definisi proto product.
  - `user/` : Definisi proto user.

## Menjalankan Project

1. **Install dependencies**
	```bash
	go mod tidy
	```
2. **Generate file protobuf (jika ada perubahan pada proto)**
	```bash
	protoc --go_out=. --go-grpc_out=. proto/<path-to-proto-file>
	```
3. **Jalankan server atau client**
	```bash
	go run main.go
	# atau
	go run grpcclient/main.go
	```

## Dependencies
- Go (>=1.18)
- gRPC
- Protocol Buffers

## Kontribusi
Silakan buat pull request atau issue jika ingin berkontribusi atau menemukan bug.

## Lisensi
Proyek ini menggunakan lisensi MIT.

## install awal pembuatan
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc-gen-go --version
protoc-gen-go-grpc --version

protoc --go_out=./pb --go-grpc_out=./pb --proto_path=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative chat/chat.proto
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative user/user.proto

protoc --go_out=./pb --go-grpc_out=./pb --proto_path=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ecommerce/product/product.proto

