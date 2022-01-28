# Go-Blockchain
`Go-Blockchain` is a simple Blockchain prototype using Go

## Getting Started

### Project setup
1. Get all dependencies using `Go Mod`
   ```bash
   go mod tidy -v
   ```

### Run the project
Build and run api endpoint using
```bash
go run cmd/<main.go
```

### Update wire after modifying service
1. move to core folder
   ```bash
   cd internal/app/
   ```
2. trigger wire
   ```bash
   wire
   ```
