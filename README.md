# Minesweeper API
## About
Minesweeper (a small warship designed to remove or detonate naval mines) is an antic and famous video game released in 1992 for entertain people in their PCs, more than just a game was also designed to help users acclimate to a mouse.
This project tries to adapt the game to an API RESTful version.

## Routes
**Base URL:** 
[https://juanenmellare-minesweeper-api.herokuapp.com](https://juanenmellare-minesweeper-api.herokuapp.com)


**HealtCheck:** 
| Method      | Path        | Description |
| ----------- | ----------- | ------------|
| GET         | /ping       | Should return 'pong' text is the application is ok |

**Game:** 
| Ref | Method      | Path                                       | Description             | Params & Body                                                       |
| --- | ----------- | ------------------------------------------ | ----------------------- | ------------------------------------------------------------------- |
| 1   | POST        | /v1/games                                  | Create Game             | body: `minesQuantity` (amount of mines), `height` (columns height), `width` (rows width)|
| 2   | GET         | /v1/games/:uuid                            | Get Game                     | params: `uuid` (game uuid)                                    |
| 3   | PUT         | /v1/games/:uuid/field/:field-uuid/show     | Show field                   | params: `uuid` (game uuid), `field-uuid` (field uuid to show) |
| 4   | PUT         | /v1/games/:uuid/field/:field-uuid/flag     | Mark field as mine           | params: `uuid` (game uuid), `field-uuid` (field uuid to show) |
| 5   | PUT         | /v1/games/:uuid/field/:field-uuid/hide     | Restore from mark            | params: `uuid` (game uuid), `field-uuid` (field uuid to show) |
| 6   | PUT         | /v1/games/:uuid/field/:field-uuid/question | Mark field as possible mine  | params: `uuid` (game uuid), `field-uuid` (field uuid to show) |

## Rules

- Before create a game define your minefield height (min '3' & max '16'), width (min '3' & max '30') and mines quantity (min 1 & max height * width), after this procced to call route 1.
- Once we've created our game we should keep the UUID for further operations or if just want to check the data with the route 2.
- Reveal a field value with route 3, but be aware if the value is a mine the game will be over.
- Mark a field as mine (route 4) or as possible mine (route 6), once you have found all the mines and mark them with route 4 you win!.
- Is the game is over the timer will stop and you won't be able to execute any other action.

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
