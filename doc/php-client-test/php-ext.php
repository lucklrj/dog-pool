
<?php
$run_num =10;

$dbh = new PDO('mysql:host=localhost;dbname=r', 'root', '123asd123'); 
	
$m = new Memcached();
$m->addServer('127.0.0.1','11211');

$r = new Redis();
$r->connect('127.0.0.1', 6379);

for($i=1;$i<=$run_num;$i++){
	$dbh->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION); 
	$dbh->exec('set names utf8');


	/*查询*/
	$sql = "select id,num,last_update_time from r where id < ?"; 
	$stmt = $dbh->prepare($sql); 
	$stmt->execute(array(3)); 
	print_r( $stmt->fetchAll(PDO::FETCH_ASSOC));


	$sql = "insert into r(num,last_update_time)values(?,?)";
	$stmt = $dbh->prepare($sql);  
	$stmt->execute(array(0,time())); 
	echo $dbh->lastinsertid(); 
	echo "\r";
	/*修改*/
	$sql = "update r set num=num+?";
	$stmt = $dbh->prepare($sql); 
	$stmt->execute(array(1)); 
	echo $stmt->rowCount();


	$m->set("lrj",date("Y-m-d H:i:s"),3600);
	$m->set("lrj2","你好，刘仁俊",3600);
	$m->getMulti(["lrj","lrj2"]); 
	$m->flush(); 
	$m->getMulti(["lrj","lrj2"]); 

	$r->set("lrj",date("Y-m-d H:i:s"));
	$r->get("lrj");
	$r->delete("lrj");
}

