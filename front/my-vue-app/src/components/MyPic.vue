<template>
  <div class="my-wrapper" v-cloak>
      <!-- 头部 -->
      <div class="my-header row aCenter">
          <!-- canvas背景 -->
          <canvas id="canvas"></canvas>
      </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, nextTick, onUpdated } from 'vue'
import { string } from 'yup'
export default defineComponent({
    name: 'MyPic',
	props: {
		imageSB:String,
	},
    setup(props) {
        // 创建canvas动画
//         const oninitCanvas = (img: HTMLImageElement) => {
//             nextTick(() => {
//               const canvas: any = document.getElementById('canvas'),
//                     ctx = canvas.getContext('2d'),
//                     height: number = img.height,
//                     width: number = img.width,
//                     canvasAny: any = { width: width, height: height },
//                     requestAnimFrame = (function() {        // 波浪自执行函数
//                     })()
//               let step: number = 0
//               // 动起来
//               const draw = function() {
// 				console.log("be call ed ")
//                 //   ctx.clearRect(0, 0, canvasAny.width, canvasAny.height)
// 				  ctx.drawImage(img, 0, 0);
// 				  console.log("have drawed image ");
				  
//               }
//               draw()
//             })
//         }
		onUpdated(() => {
			console.log("onUpdated");
			var canvas: any = document.getElementById("canvas");
			var ctx = canvas.getContext("2d")
			const img = new Image()
			img.onload = function () {
				canvas.width = img.width
				canvas.height = img.height
				ctx.drawImage(img, 0, 0)//绘制图片
				console.log("image onload");
				// 绘制红框
				ctx.strokeStyle = 'red';
				ctx.lineWidth = 3;
				ctx.strokeRect(100, 200, 200, 400);
			}
			if (props.imageSB!=undefined) {
				img.src = props.imageSB;
			}
		}),
		onMounted(() => {
			var canvas:any = document.getElementById("canvas");  
			var ctx=canvas.getContext("2d")
			const img = new Image()
			img.onload = function () {
				canvas.width = img.width
				canvas.height = img.height
				ctx.drawImage(img, 0, 0)//绘制图片
				console.log("image onload");
// 绘制红框
				ctx.strokeStyle = 'red';
				ctx.lineWidth = 3;
				ctx.strokeRect(100, 200, 200, 400);
			}
			if (props.imageSB!=undefined) {
				img.src = props.imageSB;
			}
        })
    }
})
</script>

<style scoped>
/* .my-header{height: 230px;background-color: #6689e2;position: relative;} */
/* #canvas {position: absolute; left: 0; right: 0; top: 0; width: 100%; height: 230px;} */
</style>