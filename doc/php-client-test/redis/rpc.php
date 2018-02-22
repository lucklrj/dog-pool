<?php
require(realpath(__dir__)."/../lib.php");
$run_num = 1;
$client = new JsonRPC("127.0.0.1", 5555);
    
for($i=0;$i<$run_num;$i++){
    $r = $client->Call("Redis.Set",array("Key"=>"lrj",'Value'=>date("Y-m-d H:i:s"),"LeftTime"=>3600),$i);
    $r = $client->Call("Redis.Get",array("Key"=>"lrj"),$i);
    $r = $client->Call("Redis.Delete",array("Key"=>"lrj"),$i);
}
