<?php
$run_num=100;
$r = new Redis();
//连接
$r->connect('127.0.0.1', 6379);
for($i=1;$i<=$run_num;$i++){
	$r->set("lrj",date("Y-m-d H:i:s"));
	$r->get("lrj");
	$r->delete("lrj");
}
