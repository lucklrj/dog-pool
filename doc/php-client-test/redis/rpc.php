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
        //stream_set_timeout($this->conn, 0, 3000);
        $line = fgets($this->conn);
        if ($line === false) {
            echo "没有获取到数据";
            return NULL;
        }
        return json_decode($line,true);
    }
}
$run_num = 1;

for($i=0;$i<$run_num;$i++){
    $client = new JsonRPC("127.0.0.1", 5555);
    $r = $client->Call("Redis.Set",array("Key"=>"lrj",'Value'=>date("Y-m-d H:i:s"),"LeftTime"=>3600),$i);
    print_r($r);
    $r = $client->Call("Redis.Get",array("Key"=>"lrj"),$i);
    print_r($r);
    $r = $client->Call("Redis.Delete",array("Key"=>"lrj"),$i);
    print_r($r);
    

    /*
    $r = $client->Call("Memcache.Get",array("Key"=>"lrj"),$i);
    print_r($r);
    $r = $client->Call("Memcache.Delete",array("Key"=>"lrj"),$i);
    print_r($r);
    $r = $client->Call("Memcache.FlushAll","",$i);
    print_r($r);
    */

    
    //$r = $client->Call("Mysql.Query",array('sql'=>"update r set num=num+?",'database'=>'test',"bind"=>[1]),$i);
    //print_r($r);
    // $r = $client->Call("Mysql.Query",array('sql'=>"insert into xx",'database'=>'xx'),$i);
    // print_r($r);
}




