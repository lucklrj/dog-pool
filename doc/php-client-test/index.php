<?php
$time_start = microtime_float();

$type="memcache";
$max=1;

for($i=0;$i<=$max;$i++){
	exec("php -f ./".$type."/php-ext.php");
}

$time_end = microtime_float();
$time = $time_end - $time_start;
echo $time;
echo "\r";

$time_start = microtime_float();
for($i=0;$i<=$max;$i++){
	exec("php -f ./".$type."/rpc.php");
}
$time_end = microtime_float();
$time = $time_end - $time_start;
echo $time;
echo "\r";


function microtime_float()
{
    list($usec, $sec) = explode(" ", microtime());
    return ((float)$usec + (float)$sec);
}