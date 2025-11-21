# Setup & Development

---

1. Package structure
2. Define Primary Ports, Secondary Ports
3. Define Models & Data Structure
4. Proto Contract
5. Generate Proto File
6. Define Ports & Models
7. Implement Primary & Secondary Adapter
8. Implement Core Business Logic
9. Testing

# Cara Generate Protobuf

---

1. Di repo ini, menggunakan libprotoc 26.1, untuk cek versi protoc yang terinstall bisa menggunakan command berikut
   ```shell
   protoc --version
   ```
2. Definisikan spec proto di folder `lib/protos/v1/xxxx/xxx.proto`
3. Misal untuk wallet di `lib/protos/v1/wallet/wallet.proto`
4. Lalu jalankan command berikut
   ```shell
    protoc -I . --proto_path=. --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --openapiv2_out ./lib/protos/openapiv2 --openapiv2_opt logtostderr=true ./lib/protos/v1/wallet/wallet.proto
   ```
5. File \*.go akan tergenerate sesuai lokasi folder proto yang kita define, dan akan melakukan generate ke spec open API di folder `lib/protos/openapiv2/lib/protos/v1/wallet/wallet.swagger.json`
6. Isi dari makefile, pada dasarnya merupakan command di atas yang lebih di rapikan
7. Untuk generate melalui makefile, misal untuk proto wallet yang ada di folder `lib/protos/v1/wallet/wallet.proto` bisa menggunakan command berikut
   ```shell
      MODULE=v1/wallet make generate
   ```
8. Untuk generate file go di semua proto, bisa langsung jalankan command berikut
   ```shell
   make generate
   ```

9. Untuk membuat schema baru, contoh commandnya sebagai berikut
   ```shell
   go run -mod=mod entgo.io/ent/cmd/ent new EntityName --target internal/membership/ent/schema
   ```

10. Untuk generate schema yang sudah dibuat, contoh commandnya sebagai berikut
   ```shell
   go generate ./internal/membership/ent
   ```
11. Untuk generate migration, contoh commandnya sebagai berikut
   ```shell
   atlas migrate diff migration_name --dir "file://internal/membership/ent/migrate/migrations" --to "ent://internal/membership/ent/schema" --dev-url "postgres://postgres:postgres@172.28.35.4:5432/postgres?search_path=public&sslmode=disable"
   ```

12. Untuk apply migration, contoh commandnya sebagai berikut
   ```shell
   atlas migrate apply --dir "file://internal/membership/ent/migrate/migrations" --url "postgres://postgres:postgres@172.28.35.4:5432/postgres?search_path=public&sslmode=disable"
   ```