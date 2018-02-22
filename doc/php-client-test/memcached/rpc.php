<?php
require(realpath(__dir__)."/../lib.php");
$run_num = 1;
$client = new JsonRPC("127.0.0.1", 5555);
for($i=0;$i<$run_num;$i++){
    
    $r = $client->Call("Memcache.Set",array("Key"=>"lrj",'Value'=>date("Y-m-d H:i:s"),"LeftTime"=>3600),$i);
    $r = $client->Call("Memcache.Set",array("Key"=>"lrj2",'Value'=>"你好，刘仁俊","LeftTime"=>3600),$i);
    $r = $client->Call("Memcache.GetMulti",array("Key"=>"lrj,lrj2",'Value'=>"你好，刘仁俊","LeftTime"=>3600),$i);
    $r = $client->Call("Memcache.FlushAll","",$i);
    $r = $client->Call("Memcache.GetMulti",array("Key"=>"lrj,lrj2",'Value'=>"你好，刘仁俊","LeftTime"=>3600),$i);
    
    /*
    $r = $client->Call("Memcache.Get",array("Key"=>"lrj"),$i);
    print_r($r);
    $r = $client->Call("Memcache.Delete",array("Key"=>"lrj"),$i);
    print_r($r);

    */

    
    //$r = $client->Call("Mysql.Query",array('sql'=>"update r set num=num+?",'database'=>'test',"bind"=>[1]),$i);
    //print_r($r);
    // $r = $client->Call("Mysql.Query",array('sql'=>"insert into xx",'database'=>'xx'),$i);
    // print_r($r);
}



