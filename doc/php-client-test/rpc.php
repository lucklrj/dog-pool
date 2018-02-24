<?php
require(realpath(__dir__)."/lib.php");
$run_num = 10;
$client = new JsonRPC("127.0.0.1", 5555);
for($i=1;$i<=$run_num;$i++){
    //$r = $client->Call("Arith.Muliply",array('A'=>3,'B'=>2),$i);Cconfig
    //print_r($r);
    
    $r = $client->Call("Mysql.Query",array('sql'=>"select id,num,last_update_time from r where id < ?",'database'=>'r',"bind"=>[3]),$i);
    print_r($r);

    $r = $client->Call("Mysql.Query",array('sql'=>"insert into r(num,last_update_time)values(?,?);",'database'=>'r',"bind"=>[0,time()]),$i);
    print_r($r);

    $r = $client->Call("Mysql.Query",array('sql'=>"update r set num=num+?",'database'=>'r',"bind"=>[1]),$i);
    print_r($r);

    $r = $client->Call("Memcache.Set",array("Key"=>"lrj",'Value'=>date("Y-m-d H:i:s"),"LeftTime"=>3600),$i);
    $r = $client->Call("Memcache.Set",array("Key"=>"lrj2",'Value'=>"你好，刘仁俊","LeftTime"=>3600),$i);
    $r = $client->Call("Memcache.GetMulti",array("Key"=>"lrj,lrj2",'Value'=>"你好，刘仁俊","LeftTime"=>3600),$i);
    $r = $client->Call("Memcache.FlushAll","",$i);
    $r = $client->Call("Memcache.GetMulti",array("Key"=>"lrj,lrj2",'Value'=>"你好，刘仁俊","LeftTime"=>3600),$i);

    $r = $client->Call("Redis.Set",array("Key"=>"lrj",'Value'=>date("Y-m-d H:i:s"),"LeftTime"=>3600),$i);
    $r = $client->Call("Redis.Get",array("Key"=>"lrj"),$i);
    $r = $client->Call("Redis.Delete",array("Key"=>"lrj"),$i);
}



