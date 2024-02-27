<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/MvArqAv.jpeg" alt="Project logo"></a>
</p>

<h3 align="center">Plata: сurrency quotation</h3>



---


## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Built Using](#built_using)
- [Authors](#authors)

## 🧐 About <a name = "about"></a>

Test assignment for the Golang developer position.
The project is a small API with three endpoints.
The functionality is described in the swagger documentation.

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.
### Prerequisites

What things you need to install the software and how to install them.

1 - You need to install docker/docker-compose
[Docker](https://docs.docker.com/compose/install/)


### Clonning and running

The first thing to do is to clone the project

```
git clone https://github.com/ShunAkuma/Plata.git
```
Second, go to the root directory of the project and start it.

```
docker-compose up
or
make start
```

If the project started successfully ->  [swagger UI](http://localhost:8080/swagger/index.html)


## ⛏️ Built Using <a name = "built_using"></a>

- [Redis](https://redis.io/docs/connect/clients/go/) - Database
- [GIN](https://github.com/gin-gonic/gin) - Web Framework
- [Docker] - Containerization

## ✍️ Authors <a name = "authors"></a>

- [@Shun](https://github.com/ShunAkuma) - Initial work
