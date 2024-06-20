# Project Manager

The Project Manager app is a cool tool designed to help you and your team keep track of your projects and tasks. With this app, you can easily create, view, edit, and delete projects, giving you a central place to manage all your work. Each project can have multiple tasks, so you can keep an eye on every little detail, track progress, and see who’s responsible for what. The user-friendly interface and solid backend make sure everything runs smoothly, helping you plan and get things done more efficiently.

This project is also a fun way to dive into building servers with Go. By using Go and SQLite, the app is lightweight and easy to set up, perfect for experimenting and learning. We’ve set up a migration system to make sure your database tables and indexes are always up-to-date. It’s a great project for anyone looking to get hands-on experience with Go, learn more about server-side development, and create something useful at the same time.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.22^
- SQLite 3.36^
- Make

## MakeFile

build the application

```bash
make build
```

run the application

```bash
make run
```

live reload the application (for dev purposes)

```bash
make watch
```

run the test suite

```bash
make test
```

clean up binary from the last build and delete the sqlite database

```bash
make clean
```
