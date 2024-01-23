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

GET /advertisements
GET /advertisements/{id}
POST /advertisements
POST auth//signup
POST auth//signin
