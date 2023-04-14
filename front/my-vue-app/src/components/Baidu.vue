<template>
  <div>
    <BMap v-bind="$attrs" enableScrollWheelZoom ref="map" :center="point" @initd="handleInitd" @click="handleClick"
      mapStyleId="0f3219e982947931ae2893345940df80">
      <template v-if="!isLoading && !isEmpty">
        <BMarker :position="point" enableClicking @click="handleClick2">
        </BMarker>
        <BInfoWindow v-show="show2" :position="point" title="result?.address" enableAutoPan enableCloseOnClick>
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
      </template>
    </BMap>
  </div>
</template>

<script lang="ts" >
import { ref } from 'vue'
import { usePointGeocoder, PointGeocoderResult } from 'vue3-baidu-map-gl'

import { BInfoWindow } from 'vue3-baidu-map-gl'
import { defineComponent } from 'vue'
import { onMounted } from 'vue'
import { ImgHTMLAttributes } from 'vue'
export default defineComponent({
  setup(props, context) {
    // 在这里声明数据，或者编写函数并在这里执行它
    const map = ref()
    const imgP = ref<ImgHTMLAttributes>()
    const { get, result, isLoading, isEmpty } = usePointGeocoder<PointGeocoderResult>()
    const point = ref({ lng: 116.30793520652882, lat: 40.05861561613348 })
    const imgBASEURL = "https://api.map.baidu.com/panorama/v2?width=139&height=104&location="
    const lastBaseURL = "&fov=180&ak=Np5CcMkThnumd5L4MpQ01LjMxvrjUyxr"
    var imgURL = ""
    var d = "ds"


    const show2 = ref<boolean>(true)
    // onMounted(() => {
    //   // console.log("before", imgP.value.alt);
    //   imgP.value.src = imgURL
    //   imgP.value.alt = d
    //   console.log(imgP.value.alt);
    // })



    return {
      map,
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
  methods: {

    // 在这里声明方法
    handleInitd() {
      const { get, result, isLoading, isEmpty } = usePointGeocoder<PointGeocoderResult>()
      const point = ref({ lng: 116.30793520652882, lat: 40.05861561613348 })
      get(point.value)
      console.log("called inited ");
      console.log(isEmpty.value, isLoading.value);

    },
    handleClick2(e) {
      const show2 = ref<boolean>(true)
      console.log("hsssadas");

      if (show2.value == false) {
        show2.value = true
      } else if (show2.value == true) {
        show2.value = false

      }
    },
    handleClick(e) {
      const imgBASEURL = "https://api.map.baidu.com/panorama/v2?width=139&height=104&location="
      const lastBaseURL = "&fov=180&ak=Np5CcMkThnumd5L4MpQ01LjMxvrjUyxr"
      var imgURL = ""
      var d = "ds"
      const show2 = ref<boolean>(true)
      const { get, result, isLoading, isEmpty } = usePointGeocoder<PointGeocoderResult>()
      const point = ref({ lng: 116.30793520652882, lat: 40.05861561613348 })
      const markerPoint = point
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


      console.log(imgURL);

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
