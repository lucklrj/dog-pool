<?php
require(realpath(__dir__)."/../lib.php");
$run_num = 100;
$client = new JsonRPC("127.0.0.1", 5555);
for($i=1;$i<=$run_num;$i++){
    //$r = $client->Call("Arith.Muliply",array('A'=>3,'B'=>2),$i);Cconfig
    //print_r($r);
    
    $r = $client->Call("Mysql.Query",array('sql'=>"select id,num,last_update_time from r where id < ?",'database'=>'r',"bind"=>[3]),$i);
    print_r($r);

    // $r = $client->Call("Mysql.Query",array('sql'=>"insert into r(num,last_update_time)values(?,?);",'database'=>'r',"bind"=>[0,time()]),$i);
    // print_r($r);

    // $r = $client->Call("Mysql.Query",array('sql'=>"update r set num=num+?",'database'=>'r',"bind"=>[1]),$i);
    // print_r($r);

}



