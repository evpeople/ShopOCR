<template>
  <div>
    <input type="file" ref="fileInput" @change="handleFileInputChange">
    <button @click="uploadFile2">上传文件</button>
    <img id="my-img" v-if="imageUrl" :src="imageUrl" alt="上传的图片">
    <my-pic v-if="replies" :image-s-b="imageUrl" :replies="replies"></my-pic>
  </div>
</template>

<script>

import OCRService from "../services/ocr.service"
import UploadService from "../services/upload.service"
import  MyPic from "./MyPic.vue";
export default {
  components: {
    MyPic
  },
  data() {
    return {
      file: null,
      imageUrl:null,
      replies:null
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
    // async uploadFile() {
    //   try {
    //     const base64Data = await this.encodeFileToBase64(this.file)
    //     const response = await OCRService.postImageData({
    //       base64: base64Data
    //     })
    //     console.log(JSON.stringify(response.data))
    //     console.log(response.data)
    //     this.replies=response.data.Replies
    //   } catch (error) {
    //     console.error('上传文件失败', error)
    //   }
    // },
    async uploadFile2() {
      try {
        // const formData = new FormData();
        // formData.append('image', this.file); // `image`是后端接口接收文件的键名，`file`是文件对象
        const response22= await  UploadService.uploadImg({file:this.file})
        console.log(response22);
        const response = await OCRService.postImageURL({
          url : response22.data.name,
          base64:""
        })
        console.log(JSON.stringify(response.data))
        console.log(response.data)
        this.replies=response.data.Replies
      } catch (error) {
        console.error('上传文件失败', error)
      }
    },
  },
}
</script>
