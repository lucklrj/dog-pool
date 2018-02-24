<?php
$m = new Memcached();
$m->addServer('127.0.0.1','11211');
$m->set("lrj",date("Y-m-d H:i:s"),10000);

$max=1;
$i=0;
while($i<$max){
	$m->set("lrj",date("Y-m-d H:i:s"),3600);
	$m->set("lrj2","你好，刘仁俊",3600);
	$m->getMulti(["lrj","lrj2"]); 
	$m->flush(); 
	$m->getMulti(["lrj","lrj2"]); 
	$i++;
}


