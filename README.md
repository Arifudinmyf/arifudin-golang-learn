# arifudin-golang-learn

# Membuat project baru
```bash
go mod init nama-folder-project
```

# Menambah Dependency wajib
```bash
go get github.com/gofiber/fiber/v2
```

# tambah dependency dengan menggunakan 
```bash
go get nama-dependency
```

# Menambah/menghapus dependency di go.mod
```bash
go mod tidy
```

# Mengunduh semua dependency tanpa compile.
```bash
go mod download
```

# Check Migrate
```bash
ls $(go env GOPATH)/bin | grep migrate

```

# Running migrate 
```bash
migrate -path db/migrations -database "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable" up
```
# atau jika menggunakan vendor
```bash
    #### a. build vendor tanpa run

    go build -mod=vendor ./cmd/app
    #### b. Build dan running projec

    go run -mod=vendor ./cmd/app/main.go
```

# Jalankan Project
```bash
go run main.go
```

# Menyalin semua dependency ke folder /vendor
#### ketika deploy atau run offline bisa berjalan
```bash
go mod vendor
```

# Menyalin semua dependency ke folder /vendor
```bash
project-root/
├── go.mod
├── go.sum
├── vendor/
│   ├── github.com/
│   │   ├── gin-gonic/
│   │   ├── golang-migrate/
│   │   └── ...
│   └── modules.txt
└── main.go
```
```bash
| Kapan digunakan                                          | Kenapa penting                                                            |
| ------------------------------------------------------   | ------------------------------------------------------------------------- |
| ✅ Saat **build di environment tanpa internet**         | Supaya semua dependency sudah tersedia secara lokal di folder `vendor/`.  |
| ✅ Saat **CI/CD atau Docker build**                     | Agar tidak perlu `go mod download` setiap kali build → build lebih cepat. |
| ✅ Saat ingin **mengunci dependency** ke versi tertentu | Versi library disalin langsung → tidak akan berubah tanpa update manual.  |
| ✅ Saat **code review atau distribusi source**          | Orang lain bisa build tanpa butuh akses internet ke repo dependency.      |

```

# Install Global Migration CLI
```bash
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
## Export go Path
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## Buat Migrasi database
##### membuat dua file SQL migrasi baru dengan timestamp otomatis di folder migrations/, siap untuk kamu isi dengan query CREATE dan DROP.
```bash
migrate create -ext sql -dir db/migrations create_users_table
```
```bash
migrate -version
grep confluent go.mod

go clean -modcache
```

# CRUD Golang + Fiber + PostgreSQL + Confluent Kafka + golang-migrate

Proyek contoh ini berisi scaffold lengkap untuk membuat CRUD (Users) menggunakan:
- Go (mod)
- Fiber (web framework)
- PostgreSQL (database)
- Confluent Kafka (producer - menggunakan `confluent-kafka-go`)
- golang-migrate (migration tool)


Struktur proyek:

```bash
crud-fiber-postgres-kafka/
├── go.mod
├── .env.example
├── docker-compose.yml
├── migrations/
│ ├── 000001_create_users_table.up.sql
│ └── 000001_create_users_table.down.sql
├── cmd/
│ └── app/
│ └── main.go
├── internal/
│ ├── config/
│ │ └── config.go
│ ├── db/
│ │ └── postgres.go
│ ├── models/
│ │ └── user.go
│ ├── repository/
│ │ └── user_repo.go
│ ├── handlers/
│ │ └── user_handler.go
│ └── kafka/
│ ├── producer.go
│ └── consumer.go
└── README.md
```

## Buat folder
```bash
mkdir -p arifudin-golang-learn/{cmd,internal/{config,db,models,repository,handlers,kafka},migrations}
```
## Buat file
```bash
touch .env.example docker-compose.yml

touch internal/config/config.go
touch internal/db/postgres.go
touch internal/models/user.go
touch internal/repository/user_repo.go
touch internal/handlers/user_handler.go
touch internal/kafka/producer.go
touch internal/kafka/consumer.go

touch cmd/app/main.go
```