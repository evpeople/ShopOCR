<template>
  <div>
    <input type="file" ref="fileInput" @change="handleFileInputChange">
    <button @click="uploadFile">上传文件</button>
  </div>
</template>

<script>

import OCRService from "../services/ocr.service"
export default {
  data() {
    return {
      file: null
    }
  },
  
  methods: {
    handleFileInputChange(event) {
      this.file = event.target.files[0]
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
        console.log(response.data)
      } catch (error) {
        console.error('上传文件失败', error)
      }
    }
  }
}
</script>
