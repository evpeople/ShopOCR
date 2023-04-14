<template>
  <div>
    <BMap v-bind="$attrs" enableScrollWheelZoom ref="map" :center="point" @initd="handleInitd" @click="handleClick" mapStyleId="0f3219e982947931ae2893345940df80">
        <BPanoramaControl />
    <BPanoramaCoverageLayer />
      <template v-if="!isLoading && !isEmpty">
        <BMarker :position="point" v-on:click="show2 = !show2" enableClicking></BMarker>
        <BInfoWindow v-model="show2" :position="point" title="停车场信息" enableAutoPan enableCloseOnClick>
        <h2>天安门</h2>
        <div class="infoWindow-content">
          <p> 
        {{ result?.address }}
          </p>
          <img width="139" height="104" src="" alt="" />
        </div>
      </BInfoWindow>
        <!-- <BLabel
          style="color: #333; font-size: 9px"
          :position="result.point"
          :content="`地址: ${result?.address} 所属商圈:${result?.business} 最匹配地点: ${
            result?.surroundingPois[0]?.title || '无'
          }`"
        /> -->
         <!-- <BInfoWindow v-model="show2" :position="result.point" title="图文组合排版">
      </BInfoWindow> -->
      </template>
    </BMap>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue'
import { usePointGeocoder, PointGeocoderResult } from 'vue3-baidu-map-gl'
import { BPanoramaCoverageLayer } from 'vue3-baidu-map-gl'
import { BPanoramaControl } from 'vue3-baidu-map-gl'

  import { BInfoWindow } from 'vue3-baidu-map-gl'
  const map = ref()
  const { get, result, isLoading, isEmpty } = usePointGeocoder<PointGeocoderResult>()
  const point = ref({ lng: 116.30793520652882, lat: 40.05861561613348 })
  const markerPoint = point
const show2 = ref<boolean>(false)
let _map
function handleInitd({ map}) {
  get(point.value)
  _map = map
	console.log(_map);
    console.log(map);
    
  }
  function handleClick(e) {
    markerPoint.value = e.latlng
    if (show2.value==false) {
    show2.value = true
    } else {
      console.log("shbi");
    show2.value = false 
      
    }
	console.log(markerPoint.value);
	console.log(show2.value)
    get(e.latlng)
	console.log("lng ",e.latlng.lng,"lat ",e.latlng.lat)
	
  }
</script>

<style scoped>
  .state {
    margin-top: 15px;
  }
  .state span {
    margin-right: 25px;
  }
    .infoWindow-content {
    display: flex;
    justify-content: space-between;
    padding: 10px;
  }
  .infoWindow-content p {
    margin: 0;
    line-height: initial;
  }

</style>
