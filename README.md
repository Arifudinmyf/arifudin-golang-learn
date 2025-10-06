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

# Jalankan Project
```bash
go run main.go
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
#### a. build vendor tanpa run
```bash
go build -mod=vendor
```
#### b. Build dan running projec
```bash
go run -mod=vendor main.go
```

# Install Global Migration CLI
```bash
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Buat Migrasi database
##### membuat dua file SQL migrasi baru dengan timestamp otomatis di folder migrations/, siap untuk kamu isi dengan query CREATE dan DROP.
```bash
migrate create -ext sql -dir migrations create_users_table
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