<template>
  <div>
    <BMap v-bind="$attrs" enableScrollWheelZoom ref="map" :center="point" @initd="handleInitd" @click="handleClick"
      mapStyleId="0f3219e982947931ae2893345940df80">
      <!-- <div v-if="!isLoading && !isEmpty"> -->
      <BMarker :position="point" enableClicking @click="handleClick2">
      </BMarker>
      <BInfoWindow v-model="show2" :position="point" title="dashabi" enableAutoPan>
        <h2>天安门</h2>
        <div class="infoWindow-content">
          <p>
            {{ result?.address }}
          </p>
          <img width="139" height="104" ref="imgP" alt="rfdfs" />
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
    </BMap>
  </div>
</template>

<script lang="ts" >
import { ref } from 'vue'
import { usePointGeocoder, PointGeocoderResult } from 'vue3-baidu-map-gl'

import { BInfoWindow } from 'vue3-baidu-map-gl'
import { defineComponent } from 'vue'
import { onMounted } from 'vue'
export default defineComponent({
  setup(props, context) {
    // 在这里声明数据，或者编写函数并在这里执行它
    const map = ref()
    const imgP = ref<HTMLImageElement>()
    const { get, result, isLoading, isEmpty } = usePointGeocoder<PointGeocoderResult>()
    const point = ref({ lng: 116.30793520652882, lat: 40.05861561613348 })
    const markerPoint = point
    const show2 = ref<boolean>(true)
    const imgBASEURL = "https://api.map.baidu.com/panorama/v2?width=139&height=104&location="
    const lastBaseURL = "&fov=180&ak=Np5CcMkThnumd5L4MpQ01LjMxvrjUyxr"
    var imgURL = ""
    var d = "ds"
    onMounted(() => {
      imgP.value.alt = "dsadasdasadsa"
      console.log(imgP.value);

    })
    function handleInitd() {
      get(point.value)
    }
    function handleClick2(e) {
      console.log("hsssadas");

      if (show2.value == false) {
        show2.value = true
      } else if (show2.value == true) {
        show2.value = false

      }
    }
    async function handleClick(e) {
      markerPoint.value = e.latlng
      // if (show2.value == false) {
      //   show2.value = true
      //   console.log("22shbi");
      // } else if (show2.value == true) {
      //   console.log("shbi");
      //   show2.value = false

      // }
      console.log(markerPoint.value);
      console.log(show2.value)
      get(e.latlng)
      console.log("lng ", e.latlng.lng, "lat ", e.latlng.lat)
      d = "lng " + e.latlng.lng, "lat " + e.latlng.lat
      console.log(d);

      imgURL = imgBASEURL + e.latlng.lng + "," + e.latlng.lat + lastBaseURL

      console.log("before", imgP.value.alt);
      imgP.value.src = imgURL
      imgP.value.alt = d
      console.log(imgP.value.alt);

      console.log(imgURL);

    }


    return {
      map,
      handleClick,
      handleClick2,
      handleInitd,
      show2,
      point,
      result,
      isLoading,
      isEmpty,
      imgURL,
      imgBASEURL,
      d,
      imgP
      // 需要给 `<template />` 用的数据或函数，在这里 `return` 出去
    }
  },
})



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
