comando uteis dev

docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management

#docker run -d -p 8080:8080 go-app

#docker build -t stark-bank .



docker build -t go-app --build-arg MODE=dev .
docker build -t go-app --build-arg MODE=prod .


docker-compose up --build -d

go mod tidy


depois do git clone 
baixar git checkout rabbitmq