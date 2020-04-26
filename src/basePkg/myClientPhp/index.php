<?php
$fq = stream_socket_client("tcp://127.0.0.1:8082", $errno, $errstr, 3); // 连接go的RPC server
if (!$fq) {
    echo "$errstr ($errno) <br />\n";
    return;
}
$arr = [
    "method" => "UserService.GetUserName",
    "params" => [102],
    "id" => 0,
];
fwrite($fp, json_encode($arr)."\n");

echo $fgets($fq)

fclose($fq)