# Personal Daily Diary Application for Go(lang)

## Requirements
- Install Docker

## Usage
Run docker compose and build `docker-compose up --build -d` (only needed 1st time)
Next time you can just run `docker-compose up -d`

Docker compose will run 3 containers: `app`, `mysql`, and 'redis' container

### API
#### localhost:8080/register
* `POST` : Register user
```
body:
{
    "username": "lala",
    "password": "P@ssword!21",
    "email": "lala@email.com",
    "name": "Lala Lolo",
    "birthday": "2021-01-12"
}
```

#### localhost:8080/login
* `POST` : login with response token
```
body:
{
    "username": "lala",
    "password": "P@ssword!21"
}

response:
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTA4OTMwNDgsInVzZXJuYW1lIjoibGFsYSIsInV1aWQiOiJkMTAwMGFmZS01NGQ1LTQzZjItOTc4Yy1jOTFhMjZkMTQ5NTkifQ.23tqRv7q4JQpayEvzLLpHZulziEew4bXi27P2eplV6Y"
}
```

#### localhost:8080/diary/{year}/{quarter}
* `GET` : get diary content
```
header:
Authorization: Bearer {token}
```

#### localhost:8080/diary
* `POST` : create diary
```
header:
Authorization: Bearer {token}

body:
{
    "content": "This is my diary",
    "date": "2021-07-19"
}
```

#### localhost:8080/logout
* `DELETE` : logout, to make token invalid
```
header:
Authorization: Bearer {token}
```
  
### Stopping
- To stop run `docker-compose stop`
- To stop and remove container `docker-compose down`

## Logs
- To see logs and errors, run `docker-compose logs -f`

## Diagram
![Diagram](http://www.plantuml.com/plantuml/png/SoWkIImgAStDuKf9B4bCIYnELL3ohGmEz55II2nMI4dYIilC0Geb5XGeQBZdvoJcfMk2PO02I3g2KbCoYy7YHI0M5nT8lQuTK3-41MH291nIyrA0bW40)
