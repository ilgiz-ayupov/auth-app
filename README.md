# Auth App

## Структура проекта

- ### /cmd

  Содержит основную точку входа программы

- ### /configs

  Содержит конфигурацию проекта

- ### /db

  Содержит текстовый файл базы данных sqlite (database.sqlite), а также папку миграций базы данных

- ### /pkg

  Содержит внешние пакеты, которые могут быть импортированы другими программами.

  - #### /handler

    Содержит Обработчики запросов, Middleware, функции для ответа

  - #### /service

    Содержит файлы Логики приложения

  - #### /repository

    Содержит файлы которые взаимойствуют с Базами данных

- ### Файлы phoneNumber.go, server.go, user.go

  Содержат глобальные структуры приложения

## Маршруты приложения

- POST /user/register - Регистрация пользователя - Принимает данные из POST формы.
  ```
    login string
    password string
    name string
    age int
  ```
- POST /user/auth - Авторицация пользователя - Принимает JSON данные.
  ```JSON
  {
      "login": "UserLogin",
      "password": "UserPassword",
  }
  ```
- GET /user/:name - Получить пользователя - Возвращает JSON данные.
  ```JSON
  {
      "id": 1,
      "name": "UserName",
      "age": 25
  }
  ```
- POST /user/phone - Добавить номер телефона - Принимает JSON данные.
  ```JSON
  {
      "phone": "998901234567",
      "description": "Description",
      "is_fax": false
  }
  ```
- GET /user/phone?q="998" - Поиск номера телефона - Возвращает JSON данные.
  ```JSON
  [
      {
          "user_id": 1,
          "phone": "998901234567",
          "description": "Description",
          "is_fax": false
      },
      {
          "user_id": 2,
          "phone": "998889876543",
          "description": "Description",
          "is_fax": false
      }
  ]
  ```
- PUT /user/phone - Изменить номер телефона - Принимает JSON данные.
  ```JSON
  {
      "id": 1,
      "phone": "998901234567",
      "description": "Description",
      "is_fax": false
  }
  ```
- DELETE /user/phone/:id - Удалить номер телефона
