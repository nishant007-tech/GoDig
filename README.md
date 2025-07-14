# GoDig

GoDig is a production-grade Go microservice scaffold that demonstrates how to wire up complex dependency graphs using Uber’s [Dig](https://pkg.go.dev/go.uber.org/dig). It provides a battle-tested starting point for building maintainable, testable, and scalable Go services suitable for teams of all sizes.

## Features

* **Config & Logging**: Centralized configuration provider and a pluggable logger.
* **Multi-DB Support**: Named instances for Postgres (`primaryDB`) and MySQL (`secondaryDB`).
* **Layered Architecture**: Separation of Repository → Service → Handler layers, each wired via `dig.In`.
* **Plugin System**: Grouped providers (`dig.Group`) so you can drop in new plugins (e.g. payment, analytics) without touching core code.
* **Advanced Dig Usage**:

  * **Named dependencies** (`dig.Name`) for multiple implementations of the same type.
  * **Optional injection** (`optional:"true"`) for non-critical components.
  * **Dependency groups** (`dig.Group`) for slice injection.

## Folder Structure

```
GoDig/
├── go.mod
├── cmd/
│   └── server/
│       └── main.go
└── internal/
    ├── config/
    │   └── config.go
    ├── logger/
    │   └── logger.go
    ├── database/
    │   ├── mysql.go
    │   └── postgres.go
    ├── repository/
    │   └── user_repo.go
    ├── service/
    │   └── user_service.go
    ├── handler/
    │   └── user_handler.go
    └── plugin/
        ├── plugin.go
        ├── payment.go
        └── analytics.go
```

## Getting Started

### Prerequisites

* Go 1.21+
* Git

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/yourusername/GoDig.git
   cd GoDig
   ```

2. **Download dependencies**

   ```bash
   go mod download
   ```

3. **Run the server**

   ```bash
   go run ./cmd/server
   ```

   The server will start on port `8080`. You can access the users endpoint:

   ```bash
   curl http://localhost:8080/users
   ```

## Usage

* **Config**: Modify `internal/config/config.go` to load from environment or files.
* **Logger**: Replace `internal/logger/logger.go` with your preferred logging library.
* **Database**: Update `internal/database/*.go` for real DB connections.
* **Plugins**: Add new plugin constructors under `internal/plugin` and register them with `dig.Group("plugins")` in `main.go`.

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -m 'Add new feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
