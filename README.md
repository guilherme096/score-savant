# Score-Savant

An app that acts as a database of football players, clubs, leagues and countries with detailed information about each entity.

## Context

Project developed under the Databases course at University of Aveiro.

## Features

1. **Player Registration**:
   - Insert and remove new players in the database.
   - Associate players with clubs, leagues, countries, and contracts.
2. **Player Search**:
   - Search players by different attributes (name, nationality, age, club, position, etc.).
   - Pagination of search results.
3. **Club and League Search**:
   - Obtain detailed information about clubs and leagues, including the number of players, average salaries, and market value.
4. **Attribute Management**:
   - Calculate player role ratings according to their attributes and positions.
5. **Additional Features**:
   - Mark players as favorites.
   - Obtain random players.

## Stack

- [HTMX](https://htmx.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [Go](https://golang.org/)
  - [Echo](https://echo.labstack.com/)
  - [Templ](https://templ.guide/)
  - [Air](https://github.com/air-verse/air)
- [Microsoft SQL Server](https://www.microsoft.com/en-us/sql-server)
- [Docker & Docker Compose](https://www.docker.com/)

## How to run the project:

This repository works only as an archive for the code. Since the database is hosted int IEETA's servers, it is not possible to have
access to the database without exposing the credentials. However, the code can be run locally but it will not be able to connect to the database.

**Note:** The docker must be running to run the project.

```bash
make buid
make run
```

**Home Page:** [http://localhost:8080](http://localhost:8080)

## Grade

This project was graded with **16**/20.

## Authors

- [Guilherme Rosa](https://github.com/guilherme096)
- [João Roldão](https://github.com/JohnnyBoiR04)
