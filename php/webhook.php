<?php

use GuzzleHttp\Client;
use Monolog\Logger;
use Monolog\Handler\StreamHandler;
use Monolog\Formatter\LineFormatter;

require 'vendor/autoload.php';

// Env setting.
$config = parse_ini_file('config.ini');
$logger = new Logger('Webhook');
$handler = new StreamHandler($config['log_file'], Logger::INFO);
$formatter = new LineFormatter(null, null, false, true);
$handler->setFormatter($formatter);
$logger->pushHandler($handler);

/**
 * Summary.
 *
 * Send message based on Webhook.
 *
 * @since 1.0.0
 *
 * @see null
 * @link https://github.com/ExinOne/webhook-samples
 *
 * @param object $config Required. Base configurations.
 * @param object $logger Required. Logging to file.
 * @param string $data Required. Message content.
 * @param string $category Required. Message category. Default PLAIN_TEXT.
 * @return null.
 */
function send_mixin($config, $logger, $data, $category = 'PLAIN_TEXT')
{
    if (empty($data)) {
        return;
    }

    $client = new Client();

    try {
        $response = $client->post($config['webhook_url'].$config['access_token'],
            [
                'headers' => [
                    'Content-Type' => 'application/json',
                ],
                'body'  => json_encode([
                    'category' => $category,
                    'data'     => $data,
                ]),
            ]);

        $body = $response->getBody();
        $status = json_decode($body)->success;
        $logger->info($body);

        if ($status == "true"){
            $logger->info("Send mixin Successfully.");
        }
        else{
            $logger->error("Send mixin failed.");
        }

    } catch (\Exception $e) {
        $logger->error($e.getMessage());
        return;
    }
}

// call send mixin function
send_mixin($config, $logger, 'Hello World')

?>