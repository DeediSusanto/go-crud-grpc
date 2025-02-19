 Struktur Direktori Proyek (Setelah Implementasi 

pb\                         # Protobuf files
│   ├── lead.proto              # File utama Protobuf
│   ├── lead.pb.go              # Dihasilkan oleh protoc (gRPC service definitions)
│   ├── lead_grpc.pb.go         # Dihasilkan oleh protoc (gRPC client/server interfaces)
│
│── internal\                    # Business logic & data layer
│   ├── repository\              # Repository layer untuk DB access
│   │   ├── lead_repository.go   # Repository untuk Lead
│   ├── service\                 # Service layer untuk business logic
│   │   ├── lead_service.go      # Service untuk Lead
│
│── handler\                     # Layer handler untuk HTTP (REST API)
│   ├── lead_handler.go          # Handler untuk Lead API
│
│── model\                       # Struktur database (GORM Models)
│   ├── lead.go                  # Model Lead untuk ORM
│
│── config\                      # Konfigurasi sistem
│   ├── database.go              # Koneksi ke MySQL
│
│── server\                      # Server dan entry point aplikasi
│   ├── main.go                  # Entry point aplikasi (Gin & gRPC server)
│
│── go.mod                        # Go module dependencies
│── go.sum                        # Checksum dependencies
│── README.md                     # Dokumentasi proyek


 Penjelasan Struktur yang Diperbarui
✅ pb/lead.proto → File definisi Protobuf utama.
✅ pb/lead.pb.go & pb/lead_grpc.pb.go → File yang dihasilkan oleh protoc.
✅ internal/repository/lead_repository.go → Fungsi untuk query database (CRUD).
✅ internal/service/lead_service.go → Business logic Lead sebelum diteruskan ke repository.
✅ handler/lead_handler.go → API handler untuk REST menggunakan Gin.
✅ model/lead.go → Model Lead untuk ORM menggunakan GORM.
✅ config/database.go → File konfigurasi database MySQL.
✅ server/main.go → Entry point untuk menjalankan gRPC & REST API server.