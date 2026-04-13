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

### Day 2 - ✅ Completed
**Date:** April 13, 2026

**What I Learned:**
- Variables and data types (string, int, bool, float)
- Control structures: if/else, else if, switch statements
- For loops (traditional, while-style, infinite loops)
- Functions with single and multiple return values
- Arrays (fixed size) vs Slices (dynamic size)
- Maps (key-value pairs)
- Range keyword for iteration
- Short variable declaration (:=) vs var keyword

**What I Built:**
- Implemented comprehensive examples of all basic Go concepts
- Created helper functions (add, multiply, divide)
- Demonstrated slice operations (append, slicing, indexing)
- Practiced map creation and iteration
- Used range to iterate over slices and maps

**Key Concepts:**
- Go has only one loop construct: for (but it's versatile)
- Functions can return multiple values
- Slices are more flexible than arrays
- Maps are reference types (like dictionaries in other languages)
- Type inference with := makes code cleaner

---

### Day 3 - 🔄 Upcoming
**Target Topics:**
- Structs and methods
- Pointers basics
- Interfaces introduction

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
