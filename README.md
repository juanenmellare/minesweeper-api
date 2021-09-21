# Minesweeper API
## About
Minesweeper (a small warship designed to remove or detonate naval mines) is an antic and famous video game released in 1992 for entertain people in their PCs, more than just a game was also designed to help users acclimate to a mouse.
This project tries to adapt the game to an API RESTful version.

## For Contributors
### Requirements
- Golang 1.16
- Docker

### Format Code
```bash
make format
```

### Run Tests
```bash
make tests
```

```bash
make tests-report // Create a html report of coverage in the root project folder.
```

### Run API with Postgres (docker + live reload)
```bash
make docker-up
```