# API REST in GOLANG
* [Installation](#installation)
* [Dependencies](#dependencies)
* [API List](#api_list)
    * [Get people by id](#get_people)
    * [Get all peoples](#get_all_people)
    * [Create new people](#create_people)
    * [Update people](#update_people)
    * [Delete people](#delete_people)
* [Improvement](improvement)

# Installation
Project recovery:

    git clone https://github.com/tiyodev/api-rest-go-v1.git

Run http server in local:

    go run main.go

# Dependencies
In this project I use two external dependencies
  * [Gorilla mux](https://github.com/gorilla/mux)
  * [GORM](https://github.com/jinzhu/gorm) 

Installation of Gorilla mux
  
    go get -u github.com/gorilla/mux

Installation of GORM

    go get -u github.com/jinzhu/gorm

# API List

## Get people by id
### Request
`GET /people/{id}`

    curl -i -H 'Accept: application/json' http://localhost:8080/people/1

### Response
    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Thu, 19 Sep 2019 06:55:16 GMT
    Transfer-Encoding: chunked

    {"name":"Luke Skywalker","height":"172","mass":"77","hair_color":"blond","skin_color":"fair","eye_color":"blue","birth_year":"19BBY","gender":"male","homeworld":{ ... },"homeworldId":1,"films":[ ... ],"species":[ ... ],"vehicles":[ ... ],"starships":[ ... ], "created":"2014-12-09T13:50:51.644000Z","edited":"2014-12-20T21:17:56.891000Z","url":1,"id":1}

## Get all peoples
### Request
`GET /peoples?limit=2&offset=5`

    curl -i -H 'Accept: application/json' http://localhost:8080/peoples

  or

    curl -i -H 'Accept: application/json' "http://localhost:8080/peoples?limit=2&offset=1"

### Response
    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Thu, 19 Sep 2019 07:05:39 GMT
    Transfer-Encoding: chunked

    [{"name":"Luke Skywalker","height":"172","mass":"77","hair_color":"blond","skin_color":"fair","eye_color":"blue","birth_year":"19BBY","gender":"male","homeworld":{ ... },"homeworldId":1,"films":[ ... ],"species":[ ... ],"vehicles":[ ... ],"starships":[ ... ], "created":"2014-12-09T13:50:51.644000Z","edited":"2014-12-20T21:17:56.891000Z","url":1,"id":1}, { ... }]

## Create new people
### Request
`POST /people`

    curl -i -d "{\"name\":\"yohannBis\",\"homeworld_id\":57}" -X POST -H "Content-Type: application/json"  http://localhost:8080/people

### Response
    HTTP/1.1 201 Created
    Content-Type: application/json
    Date: Thu, 19 Sep 2019 07:53:57 GMT
    Content-Length: 615

    {"name":"yohannBis","height":"unknown","mass":"unknown","hair_color":"unknown","skin_color":"unknown","eye_color":"unknown","birth_year":"unknown","gender":"na","homeworld":{...},"homeworld_id":57,"films":null,"species":null,"vehicles":null,"starships":null,"created":"2019-09-19 09:53:57.9153238 +0200 CEST m=+125.634429001","edited":"2019-09-19 09:53:57.9153238 +0200 CEST m=+125.634429001","url":89,"id":89}

### Body

| params name | params value | default value | required | 
| :---------- |:-------------| :--------| :--------|
| name | string | unknown | false |
| homeworld_id | int |  | true |
| height | string | unknown | false |
| mass | string | unknown | false |
| hair_color | string | unknown | false |
| skin_color | string | unknown | false |
| eye_color | string | unknown | false |
| birth_year  | string | unknown | false |
| gender  | string | na | false |

## Update people
### Request
`PUT /people/{id}`

    curl -i -d "{\"name\":\"Yohann\",\"homeworld_id\":50,\"mass\":\"32\"}" -X PUT -H "Content-Type: application/json"  http://localhost:8080/people/89

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Thu, 19 Sep 2019 09:11:18 GMT
    Content-Length: 726

    {"name":"Yohann","height":"103","mass":"32","hair_color":"unknown","skin_color":"unknown","eye_color":"unknown","birth_year":"unknown","gender":"na","homeworld":{...},"homeworld_id":50,"films":[],"species":[],"vehicles":[],"starships":[],"created":"2019-09-19 10:10:42.9312583 +0200 CEST m=+15.307038501","edited":"2019-09-19 11:01:33.5114171 
    +0200 CEST m=+14.079350801","url":89,"id":89

### Body

| params name | params value | required | 
| :---------- |:-------------| :--------|
| name | string | false |
| homeworld_id | int | false |
| height | string | false |
| mass | string | false |
| hair_color | string | false |
| skin_color | string | false |
| eye_color | string | false |
| birth_year  | string | false |
| gender  | string | false |

## Delete people
### Request
`DELETE /people/{id}`

    curl -i -X DELETE http://localhost:8080/people/89

### Response
    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Thu, 19 Sep 2019 09:28:11 GMT
    Content-Length: 5

    "ok"

# Improvement

Add unit tests.

Perform better error management.

Check that dependencies exist when creating or modifying an object.
For exemple, when we create a People object, we don't check if the parameter homeworld_id corresponding to an existing object in database.