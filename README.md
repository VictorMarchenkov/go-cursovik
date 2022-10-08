# Goly, The Go URL Shortener

This is the repo for the video found [here](https://youtu.be/bTLQT7W12dQ)

---
## PostgreSQL Docker Image
I used postgres:14 and can run it
```bash
$ docker run --name name-of-container -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=test -d postgres:14
$ docker run --name auth-psql -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=test -d postgres:14
```
You can name the container anything you want, or not name it all. `--name` flag is really used to run docker commands easier, rather than using the randomly generated UUID.

---
:zap: Happy Coding!


# What is used...

[Networking with standalone containers](https://docs.docker.com/network/network-tutorial-standalone/)

# Video conspect


# My insides

1) не забыть сделать commit (TablePlus)
2) не срабатываеи "New" на пустой ДБ
3) В docker-compose все запустил. Однако не работает из-за cors ? (не проходим через nginx) Надо настроить клиенту headers