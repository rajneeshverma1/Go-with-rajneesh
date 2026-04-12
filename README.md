# Go with Rajneesh

## Learning Journey

This repository documents my journey of learning Go (Golang) from scratch.

### Goal
Complete Go fundamentals in 5 days and build a strong foundation for backend development.

---

## Progress Tracker

### Day 1 - ✅ Completed
**Date:** April 12, 2026

**What I Learned:**
- Set up Go development environment
- Understood Go project structure and best practices
- Created a minimal, scalable Go project with clean architecture
- Learned about Go modules (go.mod)
- Understood the purpose of `cmd/` and `internal/` directories

**What I Built:**
- Created a properly structured Go project
- Implemented a simple "Hello, World" application
- Set up Git repository and pushed to GitHub

**Key Concepts:**
- Go module initialization
- Package organization
- Clean folder structure for scalability
- Entry point separation from business logic

---

### Day 2 - 🔄 Upcoming
**Target Topics:**
- Variables and data types
- Control structures (if/else, switch, for loops)
- Functions and multiple return values

### Day 3 - 🔄 Upcoming
**Target Topics:**
- Arrays, slices, and maps
- Structs and methods
- Pointers basics

### Day 4 - 🔄 Upcoming
**Target Topics:**
- Interfaces
- Error handling
- File I/O operations

### Day 5 - 🔄 Upcoming
**Target Topics:**
- Goroutines and channels (concurrency basics)
- Building a small project to consolidate learning
- Best practices and next steps

---

## How to Run This Project

### Prerequisites
- Go 1.21 or higher installed
- Git (optional, for version control)

### Running the Code
```bash
go run cmd/app/main.go
```

### Building an Executable
```bash
go build -o myapp cmd/app/main.go
```

### Running the Built Executable (Windows)
```bash
.\myapp.exe
```

---

## Project Structure

```
Go-with-rajneesh/
├── cmd/
│   └── app/
│       └── main.go          # Application entry point
├── internal/
│   └── app/
│       └── app.go           # Core application logic
├── go.mod                   # Go module definition
├── .gitignore              # Git ignore rules
└── README.md               # This file
```

---

## Resources I'm Using

- Official Go Documentation: https://go.dev/doc/
- A Tour of Go: https://go.dev/tour/
- Go by Example: https://gobyexample.com/

---

## Next Steps

After completing the 5-day learning sprint:
- Build REST APIs with Go
- Learn database integration
- Explore web frameworks (Gin, Echo, or Fiber)
- Start building real-world projects

---

**"Every expert was once a beginner. Day 1 of many!"** 🚀
