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
$client = new JsonRPC("127.0.0.1", 5555);
for($i=1;$i<=$run_num;$i++){
    //$r = $client->Call("Arith.Muliply",array('A'=>3,'B'=>2),$i);Cconfig
    //print_r($r);
    
    $r = $client->Call("Mysql.Query",array('sql'=>"select fitter_begin,fitter_end from mxj_costs where id < ?",'database'=>'r',"bind"=>[3]),$i);
    print_r($r);

    $r = $client->Call("Mysql.Query",array('sql'=>"insert into r(num,last_update_time)values(0,".time().");",'database'=>'r',"bind"=>[]),$i);
    print_r($r);

    //$r = $client->Call("Mysql.Query",array('sql'=>"update r set num=num+?",'database'=>'test',"bind"=>[1]),$i);
    //print_r($r);
    // $r = $client->Call("Mysql.Query",array('sql'=>"insert into xx",'database'=>'xx'),$i);
    // print_r($r);
}



