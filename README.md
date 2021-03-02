# How to run this project locally?

- You must install docker
- You have to define your own credentials config/config.yaml

$ docker build -t ingilizce-kelime-go .

$ docker run -p 8080:8080 -d ingilizce-kelime-go

or if you want to interactive mode

$ docker run -p 8080:8080 -it ingilizce-kelime-go

TODO
- [ ] Telegram Implementation

- [ ] Redis

- [ ] Elastic Search

- [ ] RabbitMQ

- [ ] Maybe PostgreSQL Implementation for a specific branch
- [ ] Maybe Mongo Implementation

```
 git log --all --graph --decorate --oneline
```