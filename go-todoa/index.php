<?php



static $items = [];

if (isset($_GET['id']) && isset($_GET['title']) && isset($_GET['ready'])){
    $items[] = [
        'id'    => $_GET['id'],
        'title' => $_GET['title'],
        'ready' => $_GET['ready'],
    ];
}else{
    $json = file_get_contents(__DIR__.'/main.json');
    echo $json;
    // echo json_encode($items);
    die;
}