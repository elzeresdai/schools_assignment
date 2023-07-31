# School Assigment

School managment system

## Installation

```bash
make build
```

## Usage

for starting application with **--port** flag
```bash
make jumpin
go run cmd/main.go --port 8090
```
without **--port** flag it will starts on 8080 by default

##Endpoints

>
> -  Create schools using the **POST /schools endpoint**, providing the school name and address in the request body.
> - Add students to schools using the **POST /schools/:schoolID/students** endpoint, providing the student details in the request body.
>- Get list of all schools using the **GET /schools** endpoint.
>- Get a list of all students in the order they were added using the **GET /students/all** endpoint.
>- Get a list of students for a specific school using the **GET /schools/:schoolID/students** endpoint.