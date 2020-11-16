# challenge-api
 API that is able to evaluate arbitrary logical expressions defined by the user.

## Instructions

For run the project use command on project root **make run** or **docker-compose up**. Then a postgres database up with initial schema and a rest api made with golang run on port **9011**. 

For run test use command **make test**

## Endpoints
 
### signin

This service returns token to use on others services. Below an example using the user creates on database startup
>Token for saved user **test**, on database, is eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.frTc-4CSqX_UBTNcEug-94lqKq1ReG8ZcDG33WE1T8E

```json
POST http://localhost:9011/signin
Content-Type: application/json

{
    "username":"test",
    "secret": "secret"
}

``` 

Response

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.frTc-4CSqX_UBTNcEug-94lqKq1ReG8ZcDG33WE1T8E"
}
```

### Create an expression

Service to add new expression on database.

```
POST http://localhost:9011/expressions
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.frTc-4CSqX_UBTNcEug-94lqKq1ReG8ZcDG33WE1T8E

{
    "expression": "((x OR y) AND (z OR k) OR j)"
}
```

### update an expression

Update an existing expression on database

```
POST http://localhost:9011/expressions
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjN9.PZLMJBT9OIVG2qgp9hQr685oVYFgRgWpcSPmNcw6y7M

{
    "id":1,
    "expression": "x OR z"
}
```

### List expressions

List all expressions saved

```
GET http://localhost:9011/expressions
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.frTc-4CSqX_UBTNcEug-94lqKq1ReG8ZcDG33WE1T8E
```

### Evaluate an expression

Evaluate an expression. following below the uri format

> /evaluate/{expression_id}?x=1&y=0&z=1

where expression_id is the id of the expression saved on database and all query params is variables used in expression. Their values can be 1 for true and 0 for false

```
GET http://localhost:9011/evaluate/2?x=0&y=0&z=1&k=0&j=1
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.frTc-4CSqX_UBTNcEug-94lqKq1ReG8ZcDG33WE1T8E
```

### Delete expression

Delete an expression saved before by create or update service.

```
DELETE  http://localhost:9011/expressions/1
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.frTc-4CSqX_UBTNcEug-94lqKq1ReG8ZcDG33WE1T8E

```
