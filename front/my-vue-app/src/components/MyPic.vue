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
interface Pos{
	  x: number
  y: number
}
interface Reply{
  content: string
  probability: number
  pos: Array<Pos> 
}
export default defineComponent({
    name: 'MyPic',
	props: {
		imageSB:String,
		replies:Array<Reply>
	},
	setup(props) {
		const draw = (props) => { 
			nextTick(() => {
			var canvas:any = document.getElementById("canvas");  
			var ctx=canvas.getContext("2d")
			const img = new Image()
			img.onload = function () {
				canvas.width = img.width
				canvas.height = img.height
				ctx.drawImage(img, 0, 0)//绘制图片
				console.log("image onload");
// 绘制红框
				props.replies.forEach((reply: Reply) => {
					console.log(reply);
					
					ctx.beginPath();
					ctx.strokeStyle = "red";
					ctx.lineWidth = 1;
					ctx.moveTo(reply.pos[0].x, reply.pos[0].y);
					console.log(reply.pos[0].x, reply.pos[0].y);
					
					for (let i = 1; i < reply.pos.length; i++) {
						ctx.lineTo(reply.pos[i].x, reply.pos[i].y);
					}
					ctx.closePath();
					ctx.stroke();
					ctx.fillStyle = "red";
					ctx.font = "15px Arial";
					ctx.fillText(reply.content, reply.pos[0].x, reply.pos[0].y);
					ctx.fillText(reply.probability, reply.pos[1].x, reply.pos[1].y);

				});
			}
			if (props.imageSB!=undefined) {
				img.src = props.imageSB;
			}
			})
		}
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
			
			draw(props)
		}),
			onMounted(() => {
			
			draw(props)
        })
    }
})
</script>

<style scoped>
/* .my-header{height: 230px;background-color: #6689e2;position: relative;} */
/* #canvas {position: absolute; left: 0; right: 0; top: 0; width: 100%; height: 230px;} */
</style>