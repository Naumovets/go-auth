# Система аутентификации и авторизации go-auth
Проект предоставляет базовую реализацию сервиса аутентификации и авторизации на Go, используя gRPC и http сервер (grpc-gateway).
Это будет отличной отправной точкой для быстрой реализации вашего собственнного сервиса аутентификации и авторизации.

## Как запустить проект
  - Установите docker и docker-compose 
  - Скачайте репозиторий:
    ```
    git clone https://github.com/Naumovets/go-auth.git
    ```
  - Зайдите в папку:
    ```
    cd go-auth
    ```
  - Создайте файл с переменными окружения .env:
    ```
    # .env
    PG_NAME=YOUR_NAME_DATABASE
    PG_USER=YOUR_USERNAME
    PG_PASSWORD=YOUR_PASSWORD
    PG_HOST=CONTAINER_NAME_OF_YOUR_DB # посмотрите имя контейнера с бд в docker-compose.yml
    PG_PORT=YOUR_PORT_DB
    HTTP_PORT=YOUR_HTTP_SERVER_PORT
    GRPC_PORT=YOUR_GRPC_SERVER_PORT 
    ```
  - Создайте файл с переменными окружения .auth.env
    ```
    # .auth.env
    # придумайте / сгенерируйте секреты для шифрования refresh и access токенов
    REFRESH_TOKEN_SECRET=SECRET_REFRESH
    ACCESS_TOKEN_SECRET=SECRET_ACCESS
    ```
  - Запустите docker-compose:
    ```
    docker compose up
    ```
