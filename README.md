RabbitMQ Training

```sh
// run consumer
cd consumer
sudo docker build -t training/consume .
sudo docker run training/consume


// run producer
cd ../producer
node index.js
```