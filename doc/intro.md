**客户端**

```
<?php
class JsonRPC
{
    private $conn;

    function __construct($host, $port) {
        $this->conn = fsockopen($host, $port, $errno, $errstr, 3);
        if (!$this->conn) {
            return false;
        }
    }

    public function Call($method, $params,$id) {
        if ( !$this->conn ) {
            return false;
        }
        $err = fwrite($this->conn, json_encode(array(
                'method' => $method,
                'params' => array($params),
                'id'     => $id,
            ))."\n");
        if ($err === false)
            return false;
        stream_set_timeout($this->conn, 0, 3000);
        $line = fgets($this->conn);
        if ($line === false) {
            return NULL;
        }
        return json_decode($line,true);
    }
}
for($i=1;$i<=4;$i++){
    $client = new JsonRPC("127.0.0.1", 1234);
    $r = $client->Call("Arith.Muliply",array('A'=>3,'B'=>2),$i);
    print_r($r);
}
```


