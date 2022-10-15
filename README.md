# Card Manager project with Clean architecture

### Description

This project aims to control corporate cards that are used internally.

### Repos References

- [go-clean-arch](https://github.com/bxcodec/go-clean-arch/tree/master/article)
- [gofiber/recipes](https://github.com/gofiber/recipes/tree/master/clean-architecture)

### Stack

* Golang
* Fiber
* Gorm
* SQLite
* Docker

### Task List

- [ ] Each employee has a maximum of two cards.
- [ ] This report must be generated for each card type.
- [ ] Only authorized user can handle this data
- [ ] To register a user capable of handling this data, it will be necessary to send an key that is generated internally.
- [ ] The app must be able to receive an CSV file and save those data in the database.
- [ ] This app must be able to generate a CSV report with the data from these cards- [ ] and the information of the employee who has it.
- [ ] The CSV file must contain this layout: [ SERIAL, CPF, VALUE, NAME ]