<template>
  <div id="footer-container">
    <div class="info-container">
      <div class="footerImage-container">
        <img alt="图片加载失败" class="footerIcon" src="@/assets/logo-white.png" style="width: 45px; height: 45px;">
        <p class="footerText">&emsp;好味道交易中介平台</p>
      </div>
    </div>
    <div class="info-container">
      <div class="footerImage-container">
        <img alt="图片加载失败" class="footerIcon" src="@/assets/icon-bupt.png">
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ComFooter",
}
</script>

<style scoped>
#footer-container{
  width: 100%;
  height: 100px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: row;
  background-color: #d01d40;
}
.info-container{
  width: 40%;
  height: 90%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: row;
}
.footerImage-container{
  width: 40%;
  height: 60%;
  display: flex;
  justify-content: center;
  align-items: center;
}
.footerIcon{
  width: 70%;
  object-fit: contain;
}
.footerText{
  color: white;
  font-weight: bold;
  font-family: 幼圆;
}
</style>