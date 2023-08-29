# Note Service

Golang сервис, предоставляющий REST API интрефейс с методами:
* Регистрация пользователя
* Аутентификация пользователя
* Создание заметки
* Получение списка заметок
* Получение заметки по id

также при создание заметки происходит исправление орфографических ошибок 
(интеграция с сервисом Яндекс.Спеллер)

## Аутентификация
* JWT Token

# Start
1) Создать .env файл (по аналогии с .env.example)
2) `make` or `docker-compose up --build`

# REST API

## Регистрация пользователя

`POST /api/v1/auth/sing-up`

### Request
```json
{
  "name": "name",
  "username": "username",
  "password": "password"
}
```

## Аутентификация пользователя

`POST /api/v1/auth/login`

### Request
```json
{
  "username": "username",
  "password": "password"
}
```

### Response
```json
{
  "access": "access_token"
}
```
 
## Создание заметки

`POST /api/v1/notes`

### Request
```
Header:
    Authorization: Bearer access_token
```

```json
{
  "title": "title",
  "text": "text"
}
```

## Получение списка заметок

`GET /api/v1/notes`

### Request
```
Header:
    Authorization: Bearer access_token
```

### Response
```json
[
  {
    "id": 1,
    "title": "title",
    "text": "text"
  },
  {
    "id": 2,
    "title": "title",
    "text": "text"
  }
]
```

## Получение заметки по id

`GET /api/v1/notes/{id}`

### Request
```
Header:
    Authorization: Bearer access_token
```

### Response
```json
{
  "id": 1,
  "title": "title",
  "text": "text"
}
```

# Тестирование 
Протестировать API можно с помощью Postman коллекции `Notes Service.postman_collection.json` 