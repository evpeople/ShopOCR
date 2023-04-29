<template>
  <div class="item-container" :style="statusBack">
    <div class="i-summary">“<router-link :to="{
        name: 'seek-item',
        params: {
          id: this.requester
        }
    }">{{ text }}</router-link></div>
    <div class="i-responser i-details"><i class="el-icon-user-solid"></i>用户{{ responser }}</div>
    <div class="i-create-date i-details">发布时间：{{ createDate }}</div>
    <div class="i-edit-date i-details">最后修改：{{ editDate }}</div>
    <div class="i-status">
      <el-tag size="mini" type="info" effect="dark">{{ offerStatus }}</el-tag>
      <el-button type="plain" size="mini" icon="el-icon-edit" circle style="margin-left: 50px;" @click="isUpload = true" :disabled="buttonDisabled"></el-button>
      <el-button type="danger" size="mini" icon="el-icon-delete" circle style="margin-left: 5px;" @click="deleteOffer" :disabled="buttonDisabled"></el-button>
    </div>
    <el-dialog title="修改请品鉴" :visible.sync="isUpload" width="30%">
      <el-form ref="editOfferForm" label-width="100px" :model="editOfferForm">
          <el-form-item label="品鉴描述">
              <el-input v-model="editOfferForm.offerMsg" autocomplete="off" type="textarea" :rows="7"></el-input>
          </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
          <el-button @click="cancelForm">取 消</el-button>
          <el-button type="primary" @click="submitForm()">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'OfferResponseItem',
  data() {
    return {
      // itemInfo:{
      //   text: ' 日食记日式豆乳锅底不辣火锅底料调味料关东煮料包家庭汤料包炖菜汤 日食记日式豆乳锅底不辣火锅底料调味料关东煮料包家庭汤料包炖菜汤',
      //   responser: '王小明',
      //   createDate: '2022-12-24',
      //   editDate: '2022-12-25',
      //   status: '待接受',
      //   requester: ''
      // },
      isUpload: false,
      editOfferForm: {
        offerMsg: ""
      }
    }
  },
  props: {
    text: {
      type:String,
      required:true
    },
    responser: {
      type:Number,
      required:true
    },
    createDate: {
      type:String,
      required:true
    },
    editDate: {
      type:String,
      required:true
    },
    status: {
      type:Number,
      required:true
    },
    requester: {
      type:Number,
      required:true
    }
  },
  computed: {
    offerStatus() {
      var test = parseInt(this.status)
      switch (test) {
        case 3:
            return "取消";
        case 2:
            return "待接受";
        case 1:
            return "同意";
        case 0:
            return "拒绝";
        default:
            // console.log('test值为'+test)
            // console.log('type值为'+this.type)
            return "发布错误";
      }
    },
    statusBack() {
      switch (parseInt(this.status)) {
        case 3:
            return { 'background': 'rgb(241, 241, 241)' };
        case 2:
            return { 'background': '#fff' };
        case 1:
            return { 'background': 'rgb(234, 255, 238)' };
        case 0:
            return { 'background': 'rgb(255, 225, 225)' };
        default:
            return "家常味道";
      }
    },
    buttonDisabled() {
      if (this.status == 1) {
        return true
      } else {
        return false
      }
    }
  },
  methods: {
    classifyDate(date) {
        var year = date.getFullYear(), //获取时间，年
            month = date.getMonth() + 1, //获取时间 月 月是0-11，正常显示月份要+1
            dates = date.getDate(), //获取时间 日
            // day = date.getDay(), //判断今天是礼拜几 0-6 0是星期日
            hours = date.getHours(), //获取 小时
            minutes = date.getMinutes(), //获取 分钟
            seconds = date.getSeconds(), //获取 秒钟
            // arr = ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'];
        hour = hours < 10 ? '0' + hours : hours; //三元表达式，如果小时小于10的话 前面加0 如05点
        minutes = minutes < 10 ? '0' + minutes : minutes; //分钟三元
        seconds = seconds < 10 ? '0' + seconds : seconds; //秒钟三元
        return String(year+'/'+month+'/'+dates+' '+hour+':'+minutes+':'+seconds)
        // console.log('今天是：' + year + '年' + month + '月' + dates + '日 ' + arr[day] + ' ' + hours + ':' + minutes + ':' + seconds);
    },
    cancelForm() {
      this.isUpload = false;
    },
    submitForm() {
      let nowDate = this.classifyDate(new Date());
      let nowDate_toString = nowDate.toString();
      let config = {
          headers: {'Content-Type': 'application/json'}
      }; 
      this.$axios.patch(`/v1/response/please`, {
        "response_id": 3341,
        "request_id": this.requester,
        "offer_people_id": parseInt(this.responser),
        "offer_msg": this.editOfferForm.offerMsg,
        "offer_begin_time": this.createDate,
        "offer_change_time": nowDate_toString,
        "offer_state": this.status,
      }, config).then((response)=>{
          if (response.data === 1){
              this.$message({
                  showClose: true,
                  message: '您的账号状态错误！',
                  type: 'error',
              });
              this.$bus.$emit("changeStatus",{status: false, uname:''});
          }
          else {
              console.log("提交成功！")
              console.log(response.data)
              this.$message({
                  showClose: true,
                  message: '提交成功',
                  type: 'success',
              });
              this.isUpload = false;
              this.cancelForm();
              // location.reload()
              // this.$router.back();
          }
      });
    },
    deleteOffer() {
      let config = {
          headers: {'Content-Type': 'application/json'}
      }; 
      this.$axios.delete(`/v1/response/please`, {
        "response_id": 3341,
        "request_id": this.requester,
        "offer_people_id": parseInt(this.responser),
        "offer_msg": this.text,
        "offer_begin_time": this.createDate,
        "offer_change_time": this.editDate,
        "offer_state": this.status,
      }, config).then((response)=>{
          if (response.data === 1){
              this.$message({
                  showClose: true,
                  message: '您的账号状态错误！',
                  type: 'error',
              });
              this.$bus.$emit("changeStatus",{status: false, uname:''});
          }
          else {
              console.log("对请品鉴删除成功！")
              console.log(response.data)
              this.$message({
                  showClose: true,
                  message: '删除成功',
                  type: 'success',
              });
              // this.$router.back();
          }
      });
    }
  },
}
</script>

<style scoped>
a {
  text-decoration: none;
  color: #777;
  transition: 0.1s ease;
}

a:hover {
  color: #ee6a8b;
}

.item-container {
  width: 160px;
  /* position: absolute; */
  z-index: 1;
  /* left: 0; */
  /* top: 0; */
  background: #fff;
  border: 1px solid #fff;
  padding: 10px 10px 10px;
  transition: .1s ease;
}

.item-container:hover {
  z-index: 2;
  border: 1px solid rgb(224, 224, 224);
  box-shadow: 0 0px 10px rgb(65 62 101 / 15%);
}

.i-summary {
  width: 150px !important;
  height: 150px;
  padding: 10px 5px 10px;
  margin-bottom: 10px;
  position: relative;
  overflow:hidden;
  font-size: 16px;
  font-style: italic;
  line-height: 26px;
  color: #333;
  text-align: left;
  background: #eee;
  border-radius: 10px;
}

.i-summary:after {
    content: "...”"; position: absolute; bottom: 0; right: 10px; padding-left: 40px;
    background: -webkit-linear-gradient(left, transparent, #eee 55%);
    background: -o-linear-gradient(right, transparent, #eee 55%);
    background: -moz-linear-gradient(right, transparent, #eee 55%);
    background: linear-gradient(to right, transparent, #eee 55%);
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

.i-status {
  text-align: left;
}

</style>