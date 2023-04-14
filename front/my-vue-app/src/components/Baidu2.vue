<template>
<head>
    <meta charset="utf-8">
    <title>全景图</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="initial-scale=1.0, user-scalable=no">
    <meta http-equiv="X-UA-Compatible" content="IE=Edge">
	</head>
	<body>
    <div id="panorama" style="width:100%;height:100%"></div>
	</body>
</template>
<script >
// script type="text/javascript" src="https://api.map.baidu.com/api?v=1.0&type=webgl&ak=ZienfAK0hlDGMlqcjYeepy8PXpPIaFA6"> 

var panorama = BMapGL.Panorama('panorama',{
    maxZoom:21,
});
panorama.setPosition(new BMapGL.Point(116.403850, 39.913795));
var panoramaService = new BMapGL.PanoramaService();
console.log(panoramaService);
panoramaService.getPanoramaByLocation(new BMapGL.Point(116.403850, 39.913795), function (res) {
    console.log('位置全景信息:', res);
});
</script>

<style scoped>  
        html{height:100%}  
        body{height:100%;margin:0px;padding:0px}  
        #container{height:100%}  
</style>  