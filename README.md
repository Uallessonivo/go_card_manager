# Card Manager project with Clean architecture

## Description

This project aims to control corporate cards that are used internally.

### References

- [go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [gofiber/recipes](https://github.com/gofiber/recipes/tree/master/clean-architecture)

### Stack

- Golang
- Fiber
- Gorm
- Postgres
- Docker

### Task List

- [X] Each employee has a maximum of two cards.
- [ ] Only authorized user can handle this data.
- [X] Cards can only be registered if the employee (owner) exists.
- [X] To register a user capable of handling this data, it will be necessary to send a key that is generated internally.
- [X] The app must be able to receive an CSV file and save those data in the database.
- [X] This app must be able to generate a CSV report with the data from these cards and the information of the employee who has it.
- [X] This report must be generated for each card type.
- [X] The CSV file must contain this layout: [ SERIAL, CPF, VALUE, NAME ]
