const amqp = require('amqp-connection-manager');
require('dotenv').config();
(async () => {
  try {
    const { RABBITMQ_URL } = process.env;
    const connection = await amqp.connect([RABBITMQ_URL]);
    const queueName = 'test';

    const channel = await connection.createChannel({
      json: true,
      setup: function (channel) {
        return channel.assertQueue(queueName, { durable: true });
      },
    });

    await channel.sendToQueue(queueName, {
      hello: 'world',
    });
  } catch (error) {
    console.error(error);
  }
})();
