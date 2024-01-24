# REST API-сервис для публикаций объявлений


Для разворачивания сервиса локально необходимо клонировать репозиторий:
```
git clone https://github.com/KazakNi/advertisements_api.git
```

Перейти в директорию проекта

```
cd advertisements_api
```
Заполнить файл с переменными окружения

Запустить docker-compose

```
docker-compose up
```
Проек запущен, Вам будут доступны по локальному адресу http://127.0.0.1:8080 следующие эндпоинты:

- GET /advertisements HTTP/1.1
  
  | Параметры | Описание |
  | ------ | ------ |
  | page | _int_ пагинация |
  | price | _"ASC"/"DESC"_ сортировка |
  | date |  _"ASC"/"DESC"_ сортировка |

  Response body:

```json
  {
    "items": [{
              "name": "adv_name",
              "description": "some description",
              "price": 100,
              "photos": "links"
              }],
    "next_page": "page_num_next"
  }

```

- GET /advertisements/{id} HTTP/1.1

  Response body:

```json
{
      "name": "adv_name",
      "description": "some description",
      "price": 100,
      "photos": "links"
}
```
 
- POST /advertisements HTTP/1.1
  
    Headers:
  
    Authorization: Bearer \<your token\>

    Request body parameters:
  
  | Параметры | Тип |
  | ------ | ------ |
  | name | _string_ 
  | description | _string_|
  | price |  _int_|
  | photos |  _string_|

   Response body:

```json
  {
    "Id": "created_id",
    "status_code": 201
  }

```



- POST auth//signup HTTP/1.1
   Request body parameters:
  
  | Параметры | Тип |
  | ------ | ------ |
  | name | _string_ 
  | email | _string_|
  | password |  _int_|

  Response body:

```json
  {
    "Id": "created_id",
    "status_code": 201
  }

```

- POST auth//signin HTTP/1.1
  
  Request body parameters:
  
  | Параметры | Тип |
  | ------ | ------ |
  | name | _string_ 
  | email | _string_|
  | password |  _int_|
  
 Response: 302, Redirect to /advertisements
