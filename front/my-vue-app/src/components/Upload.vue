<template>
  <div>
    <input type="file" ref="fileInput" @change="handleFileInputChange">
    <button @click="uploadFile">上传文件</button>
    <img id="my-img" v-if="imageUrl" :src="imageUrl" alt="上传的图片">
    <my-pic v-if="imageUrl" :image-s-b="imageUrl"></my-pic>
  </div>
</template>

<script>

import OCRService from "../services/ocr.service"
import  MyPic from "./MyPic.vue";
export default {
  components: {
    MyPic
  },
  data() {
    return {
      file: null,
      imageUrl:null,
    }
  },
  methods: {
    handleFileInputChange(event) {
      this.file = event.target.files[0]
      const reader = new FileReader();
      reader.onload = () => {
        this.imageUrl = reader.result;
      };
      reader.readAsDataURL(this.file);
    },
    encodeFileToBase64(file) {
      return new Promise((resolve, reject) => {
        const reader = new FileReader()
        reader.readAsDataURL(file)
        reader.onload = () => {
          resolve(reader.result.split(',')[1])
        }
        reader.onerror = (error) => {
          reject(error)
        }
      })
    },
    async uploadFile() {
      try {
        const base64Data = await this.encodeFileToBase64(this.file)
        const response = await OCRService.postImageData({
          base64: base64Data
        })
        console.log(JSON.stringify(response.data))
        console.log(response.data)
      } catch (error) {
        console.error('上传文件失败', error)
      }
    },
  },
}
</script>
