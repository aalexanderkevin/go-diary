# Personal Daily Diary Application for Go(lang)

## Requirements
- Install Docker

## Usage
Run docker compose and build `docker-compose up --build -d` (only needed 1st time)
Next time you can just run `docker-compose up -d`

Docker compose will run 3 containers: `app`, `mysql`, and `redis` container

### API
#### localhost:8080/api/v1/register
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

#### localhost:8080/api/v1/login
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

#### localhost:8080/api/v1/diary/{year}/{quarter}
* `GET` : get diary content
```
header:
Authorization: Bearer {token}

response:
[
  {
    "content": "this is content",
    "date": "2021-05-19"
  }
]
```

#### localhost:8080/api/v1/diary
* `POST` : create diary
```
header:
Authorization: Bearer {token}

body:
{
    "content": "This is my diary",
    "date": "2021-07-19"
}

reponse:
{
    "status": "Diary added"
}
```

#### localhost:8080/api/v1/logout
* `POST` : logout, to make token invalid
```
header:
Authorization: Bearer {token}

response:
{
    "status": "Success"
}
```
  
### Stopping
- To stop run `docker-compose stop`
- To stop and remove container `docker-compose down`

## Logs
- To see logs and errors, run `docker-compose logs -f`

## Diagram
![Diagram](http://www.plantuml.com/plantuml/png/SoWkIImgAStDuKf9B4bCIYnELL3ohGmEz55II2nMI4dYIilC0Geb5XGeQBZdvoJcfMk2PO02I3g2KbCoYy7YHI0M5nT8lQuTK3-41MH291nIyrA0bW40)

![ERDiagram](http://www.plantuml.com/plantuml/png/PLDXQzim4FtkNt5pFpI6hEas-iKOqv2r3BRH4S8OZ18Kwn8VaILFEYyrO_zz9oSEPCqNIVVUlVToX-y3AyzTOv9hw6pbcWOj0zS8XYp21eqx06sXTzsrH-W2sHq8hUqOFHo8Qr3WaSagq1HQetfH2dkohVAg0TqIVBYzJTvet4R1bTeOd5ZLu5HZg3AeZ0e1OW5KW3GMKoWAcAocLu-FS_bbK9QYUOZkaiWP3LldeVNKAD_37gFZpYDQ2MO4YMXdwLU70YYhvQ7HXyfv733hU6qxOyPGezHQf2Ol6HIRygMORSeHFuaDqE2G3swHD3pFnBFoGEqYeOlH7TVeWTKkWFgjoCGPv58lg2nc90onLDK-NpnzULgMe-eK3jA6fpz5qdkw3iEIi-R6d4gLqoWoas5aYqKMTu2O2uNRV3QJjTYuv1lPt6oOE8QRs7xrrNbCQj883NBjT3W-0ekpuxx5a8WMMYQcU3_khmlefv368ytWJ_BjyyKCOps6pE3xVBduD5_-lBjzJbgnFJYlpyNtTtVF2JQAp0KlTbvoDkIvrgg7Z_FLG_9tJA5932N6S8uv7bz7MPE-J58vosWPLWy_Llzdruen6utYxYm2EsZ4tQDHJCw6cbf16cNx8sHVONEjGIJtfqy73EqnYLaWVPV2sn9IcyAG6t98q-GUhPRz_0C0)
