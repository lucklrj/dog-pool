<?php
$m = new Memcached();
$m->addServer('127.0.0.1','11211');
//$m->set("lrj",date("Y-m-d H:i:s"),10000);

$max=100;
$i=0;
while($i<$max){
	echo $m->get("lrj");
	$i++;
}


