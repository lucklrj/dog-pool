<?php
$run_num = 1;
$dbh = new PDO('mysql:host=localhost;dbname=r', 'root', '123asd123'); 
	
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
}


?>


