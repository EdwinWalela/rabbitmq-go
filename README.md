# RabbitMQ Demo

Message brokering using RabbitMQ 

`producer` generates messages and publishes them to the `TestQueue`

`consumer` consumes messages in the `TestQueue`


## Setup

- Run `docker run -d --hostname my-rabbit --name rabbit-demo -p 15672:15672 -p 5672:5672 rabbitmq:3-management`

- Access management console on: `http://localhost:15672`