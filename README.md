# go-kvstore

A lightweight, file-based key-value store written in Go — with in-memory indexing for fast lookups and thread-safe operations.

---

## 🚀 Features

- ✅ Persistent append-only storage (`data.db`)
- ✅ In-memory indexing for O(1) key lookups
- ✅ Thread-safe `Set()` and `Get()` using sync locks
- ✅ Clean modular Go project using packages
- ✅ Built from scratch as a learning project

---

## 📁 Project Structure
go-kvstore/
├── main.go # Entry point that demonstrates usage
├── go.mod # Go module file
├── storage/ # Storage logic
│ ├── index.go # Loads offsets into memory
│ └── storage.go # Set and Get logic with file I/O



---

## 💻 How to Run

Make sure [Go](https://go.dev/doc/install) is installed.

```bash
go run main.go

Expected output:
Got value for 'name': Mussy

🧠 How It Works
-Set(key, value) appends key=value to a file and stores the offset in memory
-Get(key) uses the in-memory map to seek directly to that line and read the value
-Thread safety is handled using sync.RWMutex
-File operations use os.OpenFile and Seek to control access

🛠️ Built With
-Go (1.22+)
-Standard Library only — no frameworks
-VS Code (on Windows)

🌱 Future Ideas
-Support Delete() or Update()
-Add a CLI interface (e.g., ./kvstore get name)
-Build a REST API with net/http
-Write unit tests and benchmarks
-Encode values in JSON or binary

👩‍💻 Author
Chaya M. Goldstein
Computer Science student at Touro University
GitHub Profile

📄 License
MIT – free to use, modify, and share.

