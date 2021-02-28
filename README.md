# How to run this project locally?

- You must install docker
- You have to define your own credentials config/config.yaml

$ docker build -t ingilizce-kelime-go .

$ docker run -p 8080:8080 -d ingilizce-kelime-go

or if you want to interactive mode

$ docker run -p 8080:8080 -it ingilizce-kelime-go

TODO
- [ ] Telegram Implementation

- [ ] Mongo Implementation

- [ ] Redis

- [ ] Elastic Search

- [ ] RabbitMQ

- [ ] PostgreSQL Implementation for a specific branch

```
 git log --all --graph --decorate --oneline
```