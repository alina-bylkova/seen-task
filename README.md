# seen-task

Seen-task is the implementation of REST service. 
It allows to provide information about video recipients and corresponding events.
Additionally it supports creation of new recipients and events.

## Table of Contents
1. [Getting Started](#getting-started)
2. [Dependencies](#dependencies)
3. [Environment variables](#environment-variables)
4. [Secrets](#secrets)
5. [Database](#database)
6. [Run the server](#run-the-server)
7. [Authentication](#authentication)
8. [Testing](#testing)

## Getting started

- Install `golang`, `git`, `docker` and `docker-compose`.

## Dependencies

To fetch dependencies, run:

```
$ go mod download
```

If things start getting messy, run:

```
$ go mod tidy
```
## Environment variables

| Env variable name          | Default                       | Description                                          |
|----------------------------|-------------------------------|------------------------------------------------------|
| server_address             | :8080                         | Server address                                       |
| db_user                    |                               | User to the database                                 |
| db_password                |                               | Password for the database access                     |
| db_host                    | localhost                     | IP to the database host                              |
| db_port                    | 5432                          | Database port                                        |
| db_name                    | seen                          | Name of the database                                 |
| max_connection_pool        | 3                             | Maximum number of open database connections          |
| max_connection_timeout     | 1 minute                      | Maximum amount of time a connection may be reused    |
| auth_user                  |                               | User for basic authentication                        |
| auth_password              |                               | Password for basic authentication                    |

## Secrets

The following secret variables are stored in the `secrets.env` file which is typically should be excluded from the repo (here is just for the demo):

- db_user
- db_password
- auth_user
- auth_password

This will overwrite environment variables.

## Database
PostgreSql is used for storing data.
Gorm is used to perform queries and map database records to structures.

You can run database in the container by typing the following command in the terminal:

```
$ docker-compose up
```

Database will use the port `5432`.

You can open a admin interface of the db by going to the [Admin Interface](http://localhost:9000).

Login credentials:

- System: **PostgreSQL**
- Server: **postgre**
- Username: **root**
- Password: **root**
- Database: **seen**

## Run the server

To run the server, open the second terminal and run:

```
$ go run main.go
```

It will be running at http://localhost:8080

## Endpoints

Get all recipients (any query argument will be used for searching in the database):

```
GET api/recipients?name=&email=&phone=
```

Get specific recipient by id:

```
GET api/recipients/:id
```

Create new recipient:

```
POST api/recipients
```

Create new event:
```
POST api/events
```

## Authentication

Basic auth is used for authentication.
Use basic auth with username: `seen` and password: `pass`.
Or simply pass this token `c2VlbjpwYXNz` to the Authentication header when sending requests to the server.

## Testing

To test everything, run:

```
$ go test ./...
```

To test everything with coverage info, run:
```
$ go test ./... -cover
```