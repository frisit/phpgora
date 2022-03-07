<?php

require_once __DIR__ . '/vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;


// producer - отправитель
// если увеличить connection timeout, то всё работает
$connection = new AMQPStreamConnection(
    'rabbitmq',
    5672,
    'guest',
    'guest',
    '/',
    false,
    'AMQPLAIN',
    null,
    'en_US',
    30, //Connection Timeout
    30
);
$channel = $connection->channel();

$channel->queue_declare('hello', false, false, false, false);

$msgText = 'This is message from index.php. Hi, man!';


$msgBlock = [
    'query' => $msgText,
    'consumer' => 'query prowiders'
];

$msg = new AMQPMessage(json_encode($msgText));
$channel->basic_publish($msg, '', 'hello');

echo "<b>Message is sent: </b>";
echo $msgText;

$channel->close();
$connection->close();
