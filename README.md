# basic-rest-api

## Basic REST API build with Go

## Endpoints:
> GET
---
`/users` -> Get all users
<br>
`/users/:id` -> Get user by id

> POST
---
`/users/:id` -> Add a new user

> DELETE
---
`/users/:id` -> Remove an existent user

> PATCH
---
`/users/:id` -> Modify an existent user

<br>

## Steps to build:

1. Install dependencies
```bash
$ go get .
```

2. Build executable
```bash
$ go build
```

3. Run server
<details>
<summary>On Windows</summary>

```bash
$ ./restApi.exe
```
</details>
<details>
<summary>On Linux/MacOS</summary>

```bash
$ ./restApi
```
</details>