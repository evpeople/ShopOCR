<template>
  <div class="item-container" :style="statusBack">
    <div class="i-image"><router-link
    :to="{
        name: 'seek-item',
        params: {
          id: this.id
        }
    }"><img :src="require('/static/data/'+ url)" alt="暂无图片"></router-link></div>
    <div class="i-price"><span class="left-style">接受最高单价:</span><span class="right-style"><em>￥{{ price }}</em></span></div>
    <div class="i-title" :title="title">{{ title }}</div>
    <div class="i-uploader i-details"><i class="el-icon-user"></i>用户{{ uploader }}</div>
    <div class="i-start-date i-details">发布时间：{{ startDate }}</div>
    <div class="i-end-date i-details">截止时间：{{ endDate }}</div>
    <div class="i-type"><el-tag size="mini" type="warning">{{ tasteType }}</el-tag></div>
  </div>
</template>

<script>
export default {
  name: 'SeekQuestItem',
  data() {
    return {
      // itemInfo:{
      //   imgSrc: '@/assets/sample.jpg',
      //   price: '45',
      //   type: '地方特色',
      //   title: '便利店711日式关东煮汤料炖菜火锅',
      //   uploader: '王小明',
      //   startDate: '2022-02-24',
      //   endDate: '2022-12-25',
      // },
    }
  },
  props: {
    id: {
      type:Number,
      required:true
    },
    imgSrc: {
      type:String,
      default:''
    },
    price: {
      type:Number,
      required:true
    },
    type: {
      type:String,
      required:true
    },
    title: {
      type:String,
      required:true
    },
    uploader: {
      type:Number,
      required:true
    },
    startDate: {
      type:String,
      required:true
    },
    endDate: {
      type:String,
      required:true
    },
    status: {
      type:String,
      default: '1'
    }
  },
  computed: {
    url() {
      if (this.imgSrc === '') {
        return 'default-sample.jpeg'
      } else {
        return this.imgSrc
      }
      
    },
    tasteType() {
      var test = parseInt(this.type)
      switch (test) {
        case 5:
            return "绝一味菜";
        case 4:
            return "甜酸味";
        case 3:
            return "香辣味";
        case 2:
            return "地方特色";
        case 1:
            return "家乡小吃";
        default:
            // console.log('test值为'+test)
            // console.log('type值为'+this.type)
            return "家常味道";
      }
    },
    statusBack() {
      switch (parseInt(this.status)) {
        case 3:
            return { 'background': 'rgb(255, 225, 225)' };
        case 2:
            return { 'background': 'rgb(241, 241, 241)' };
        case 1:
            return { 'background': '#fff' };
        case 0:
            return { 'background': 'rgb(234, 255, 238)' };
        default:
            return "家常味道";
      }
    }
  }
}
</script>

<style scoped>
.item-container {
  width: 160px;
  /* position: absolute; */
  z-index: 1;
  /* left: 0; */
  /* top: 0; */
  /* background: #fff; */
  border: 1px solid #fff;
  padding: 10px 10px 10px;
  transition: .1s ease;
}

.item-container:hover {
  z-index: 2;
  border: 1px solid rgb(224, 224, 224);
  box-shadow: 0 0px 10px rgb(65 62 101 / 15%);
}

.i-image img {
  width: 160px;
  height: 160px;
  aspect-ratio: auto 160 / 160;
}

.left-style {
  display: block;
  float: left;
  position: absolute;
  background: #e4393c;
  color: #fff;
  font-size: 10px;
  width: 85px;
  height: 20px;
  line-height: 20px;
  font-family: Helvetica, 'Hiragino Sans GB', 'Microsoft Yahei', '微软雅黑', Arial, sans-serif;
  border-radius: 5px;
  margin-top: 5px;
}

.right-style {
  display: block;
  float: right;
  position: absolute;
  right: 0;
  color: #e4393c;
  font-size: 24px;
}

.i-price {
  height: 33px;
}

.i-title {
  height: 20px;
  width: 160px !important;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  color: #666;
  text-align: left;
}

.i-details {
  height: 20px;
  width: 160px !important;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  color: #999;
  text-align: left;
}

.i-type {
  text-align: left;
}

</style>