# Go RESTful API Test Task for Jungle Consulting

## Run using docker-compose

```shell
git clone https://github.com/Masedko/test-task-jungle-consulting.git

cd test-task-jungle-consulting

mv .env.example .env  # This will work on Linux and MacOS, for Windows simply use: ren .env.example .env

docker build -t test-task-jungle-consulting:1.0 -f Dockerfile .

docker-compose up -d
```

## Access Swagger docs

```shell
http://localhost:8080/
```
