<template>
  <div class="details-container">
    <div class="quest-banner">
      <div class="quest-header">
        <button class="quest-back" @click="toBack"><i class="el-icon-back"></i>返回</button>
        <div class="quest-title" :title="itemInfo.title">{{ itemInfo.id }}</div>
      </div>
      <div class="quest-main">
        <div class="quest-image">
          <el-image 
            :src="require('/static/data/'+ url)" 
            :preview-src-list="srcList">
          </el-image>
        </div>
        <div class="quest-info">
          <el-descriptions :column="3">
            <template slot="title">
              <span>{{ itemInfo.title }}&emsp;</span>
              <el-tag size="small">{{ tasteStatus }}</el-tag>
            </template>
            <template slot="extra">
              <el-button type="primary" round style="margin-right: 60px;" @click="isUpload = true" v-show="checkButtonAuth"><i class="el-icon-plus"></i>我来请品鉴！</el-button>
              <el-button type="primary" round style="margin-right: 10px;" @click="isEdit = true" v-show="!checkButtonAuth" :disabled="haveResponse"><i class="el-icon-edit"></i>修改信息</el-button>
              <el-button type="danger" round style="margin-right: 60px;" @click="deleteSeek" v-show="!checkButtonAuth" :disabled="haveResponse"><i class="el-icon-delete"></i>删除</el-button>
              <el-dialog title="提交请品鉴" :visible.sync="isUpload" width="30%">
                <el-form ref="upOfferForm" label-width="100px" :model="upOfferForm">
                    <el-form-item label="品鉴描述">
                        <el-input v-model="upOfferForm.offerMsg" autocomplete="off" type="textarea" :rows="7"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="cancelForm">取 消</el-button>
                    <el-button type="primary" @click="submitForm()">确 定</el-button>
                </div>
              </el-dialog>
              <el-dialog title="修改寻味道" :visible.sync="isEdit" width="50%">
                <el-form :model="editForm" ref="editForm">
                  <el-form-item label="味道类型" label-width="250px">
                    <el-col :span="11">
                      <el-select v-model="editForm.type" placeholder="请选择味道类型">
                          <el-option label="家乡小吃" value="1"></el-option>
                          <el-option label="地方特色" value="2"></el-option>
                          <el-option label="香辣味" value="3"></el-option>
                          <el-option label="甜酸味" value="4"></el-option>
                          <el-option label="绝一味菜" value="5"></el-option>
                      </el-select>
                    </el-col>
                  </el-form-item>
                  <el-form-item label="主题名称" label-width="250px">
                    <el-col :span="17">
                      <el-input v-model="editForm.title" autocomplete="off"></el-input>
                    </el-col>
                  </el-form-item>
                  <el-form-item label="可接受最高单价" label-width="250px">
                    <el-col :span="17">
                      <el-input v-model="editForm.price" autocomplete="off"></el-input>
                    </el-col>
                  </el-form-item>
                  <el-form-item label="截止时间" label-width="250px">
                      <el-col :span="11">
                          <el-date-picker type="date" placeholder="选择截止日期" v-model="editForm.endDate" style="width: 100%;"></el-date-picker>
                      </el-col>
                  </el-form-item>
                  <el-form-item label="描述图片" label-width="250px">
                    <el-col :span="17">
                      <el-upload
                          class="upload-demo"
                          drag
                          action="https://jsonplaceholder.typicode.com/posts/"
                          multiple>
                          <i class="el-icon-upload"></i>
                          <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                          <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
                      </el-upload>
                    </el-col>
                  </el-form-item>
                  <el-form-item label="具体描述" label-width="250px">
                    <el-col :span="18">
                      <el-input v-model="editForm.descrip" autocomplete="off" type="textarea" :rows="5"></el-input>
                    </el-col>
                  </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="cancelForm">取 消</el-button>
                    <el-button type="primary" @click="submitEditForm()">确 定</el-button>
                </div>
              </el-dialog>
            </template>
              <el-descriptions-item label="发布人">用户{{ itemInfo.uploader }}</el-descriptions-item>
              <el-descriptions-item label="最高接受单价">{{ itemInfo.price }} 元</el-descriptions-item>
              <el-descriptions-item label="味道类型"><el-tag size="small" type="warning">{{ tasteType }}</el-tag></el-descriptions-item>
              <el-descriptions-item label="发布时间">{{ itemInfo.startDate }}</el-descriptions-item>
              <el-descriptions-item label="修改时间">{{ itemInfo.editDate }}</el-descriptions-item>
              <el-descriptions-item label="截止时间">{{ itemInfo.endDate }}</el-descriptions-item>
              <el-descriptions-item label="具体描述">{{ itemInfo.descrip }}</el-descriptions-item>
          </el-descriptions>
        </div>
      </div>
    </div>
    <div class="quest-response-main">
      <ul>
        <li v-for="(item, index) in this.offerInfo" :key="index">
          <el-descriptions :column="3">
            <template slot="title">
              <span>{{ item.id }}&emsp;</span>
              <el-tag size="small">{{ offerStatus(item.status) }}</el-tag>
            </template>
            <template slot="extra">
              <el-button-group style="margin-right: 60px;" v-show="!checkButtonAuth">
                <el-button type="primary" icon="el-icon-check" @click="agreePayment(item.id, 1)" :disabled="haveCompleted"></el-button>
                <el-button type="danger" icon="el-icon-close" @click="agreePayment(item.id, 0)" :disabled="haveCompleted"></el-button>
              </el-button-group>
            </template>
              <el-descriptions-item label="响应人">用户{{ item.responser }}</el-descriptions-item>
              <el-descriptions-item label="发布时间">{{ item.createDate }}</el-descriptions-item>
              <el-descriptions-item label="修改时间">{{ item.editDate }}</el-descriptions-item>
              <el-descriptions-item label="具体描述">{{ item.text }}</el-descriptions-item>
          </el-descriptions>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SeekItemDetails',
  data() {
    return {
      // reset: true,
      isUpload: false,
      isEdit: false,
      upOfferForm: {
        offerMsg: ''
      },
      editForm: {
        type: '',
        title: '',
        descrip: '',
        price: '',
        endDate: '',
        image: '',
      },
      itemInfo:{
        id: 0,
        uploader: '未知用户',
        type: '地方特色',
        title: '便利店711日式关东煮汤料炖菜火锅',
        imgSrc: '',
        descrip: '无描述信息',
        price: '?',
        startDate: '2022-02-24',
        editDate: '2022-12-24',
        endDate: '2022-12-25',
        status: '待响应'
      },
      offerInfo:[],
    }
  },
  computed: {
    // itemVisible() {
    //   return this.reset;
    // },
    checkButtonAuth() {
      var uid = parseInt(this.$cookies.get("sk_uid"))
      console.log('用户id是'+uid+',页面发布者id是'+this.itemInfo.uploader)
      if(uid == this.itemInfo.uploader) {
        return false
      } else {
        return true
      }
    },
    haveResponse() {
      if (this.offerInfo == false) {
        return false
      } else {
        return true
      }
    },
    haveCompleted() {
      if (parseInt(this.itemInfo.status) === 0) {
        return true
      } else {
        return false
      }
    },
    url() {
      if (this.itemInfo.imgSrc === '') {
        return 'default-sample.jpeg'
      } else {
        return this.itemInfo.imgSrc
      }
    },
    srcList() {
      return [
        require('/static/data/' + this.itemInfo.imgSrc)
      ]
    },
    tasteType() {
      var test = parseInt(this.itemInfo.type)
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
    tasteStatus() {
      var test = parseInt(this.itemInfo.status)
      switch (test) {
        case 3:
            return "到期未达成";
        case 2:
            return "已取消";
        case 1:
            return "待响应";
        case 0:
            return "已完成";
        default:
            // console.log('test值为'+test)
            // console.log('type值为'+this.type)
            return "发布错误";
      }
    },
  },
  methods: {
    offerStatus(val) {
      var test = parseInt(val)
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
    catchRecv(item) {
      this.itemInfo.id = item.request_id
      this.itemInfo.uploader = item.user_id
      this.itemInfo.type = item.taste_type
      this.itemInfo.title = item.topic_name
      this.itemInfo.imgSrc = item.topic_photo
      this.itemInfo.descrip = item.topic_intro
      this.itemInfo.price = item.topic_price
      this.itemInfo.startDate = item.topic_begin_date
      this.itemInfo.editDate = item.topic_change_date
      this.itemInfo.endDate = item.topic_end_date
      this.itemInfo.status = item.topic_state
      this.editForm.type = item.taste_type
      this.editForm.title = item.topic_name
      this.editForm.descrip = item.topic_intro
      this.editForm.price = item.topic_price
      this.editForm.endDate = item.topic_end_date
      this.editForm.image = item.topic_photo
      this.offerInfo.splice(0);
      for (var key in item.response) {
        var resp = item.response[key];
        console.log(resp)
        this.offerInfo.push({
          id: resp.ID,
          text: resp.offer_msg,
          responser: resp.offer_people_id,
          createDate: resp.offer_begin_time,
          editDate: resp.offer_change_time,
          status: resp.offer_state,
        })
      }
    },
    toBack() {
      this.$router.back()
    },
    resetForm() {
      this.upOfferForm.offerMsg = "";
    },
    submitForm() {
      let nowDate = this.classifyDate(new Date());
      let nowDate_toString = nowDate.toString();
      let config = {
          headers: {'Content-Type': 'application/json'}
      }; 
      this.$axios.post(`/v1/response/please`, {
        "response_id": 3341,
        "request_id": this.itemInfo.id,
        "offer_people_id": parseInt(this.$cookies.get("sk_uid")),
        "offer_msg": this.upOfferForm.offerMsg,
        "offer_begin_time": nowDate_toString,
        "offer_change_time": nowDate_toString,
        "offer_state": 2,
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
              this.resetForm();
              this.cancelForm();
              // location.reload()
              this.$router.back();
          }
      });
    },
    submitEditForm() {
      let nowDate = this.classifyDate(new Date());
      let nowDate_toString = nowDate.toString();
      let endDate_toString = this.editForm.endDate.toString();
      let config = {
          headers: {'Content-Type': 'application/json'}
      }; 
      this.$axios.patch(`/v1/user/request`, {
        "request_id": 3340,
        "user_id": parseInt(this.$cookies.get("sk_uid")),
        "taste_type": this.editForm.type,
        "topic_name": this.editForm.title,
        "topic_intro": this.editForm.descrip,
        "topic_photo": this.editForm.image,
        "topic_begin_date": this.itemInfo.startDate,
        "topic_end_date": endDate_toString,
        "topic_change_date": nowDate_toString,
        "topic_state": "1",
        "topic_price": parseFloat(this.editForm.price),
        "topic_location": {
          "topic_country": localStorage.getItem("sk_locate1"),
          "topic_province": localStorage.getItem("sk_locate2"),
          "topic_city": localStorage.getItem("sk_locate3"),
        },
        "response_id": 0,
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
              this.isEdit = false;
              this.cancelForm();
              this.$router.back();
          }
      });
    },
    cancelForm() {
      this.isUpload = false;
      this.isEdit = false;
    },
    agreePayment(id, agree) {
      let config = {
          headers: {'Content-Type': 'application/json'}
      }; 
      this.$axios.post(`/v1/user/request/choose`, {
        "request_id": parseInt(this.itemInfo.id),
        "respoense_id": parseInt(id),
        "agree": parseInt(agree)
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
              console.log("对请品鉴操作成功！")
              console.log(response.data)
              this.$message({
                  showClose: true,
                  message: '操作成功',
                  type: 'success',
              });
              this.$router.back();
          }
      });
    },
    deleteSeek() {
      let config = {
          headers: {'Content-Type': 'application/json'}
      }; 
      this.$axios.delete(`/v1/user/request`, {
        "id": parseInt(this.itemInfo.id),
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
              console.log("对寻味道删除成功！")
              console.log(response.data)
              this.$message({
                  showClose: true,
                  message: '删除成功',
                  type: 'success',
              });
              this.$router.back();
          }
      });
    }
  },
  mounted() {
    this.$bus.$on("itemDetailsRecvInfo", (data) => {
      console.log('details页面收到信息')
      this.catchRecv(data)
    })
    this.$bus.$on("itemDetailsRecvInfoFromAll", (data) => {
      console.log('details页面收到All上的信息')
      this.catchRecv(data)
    })
  },
  activated() {
    console.log('我真的发了消息了！——seek item details')
    this.$bus.$emit("itemDetailsNeedInfo", this.$route.params.id)
  },
  beforeDestroy() {
    this.$bus.$off("itemDetailsRecvInfo")
    this.$bus.$off("itemDetailsRecvInfoFromAll")
  },
}
</script>

<style scpoed>
.details-container {
  width: 100%;
  height: auto;
  min-height: 580px;
}

.quest-banner {
  background: #fff;
  border: 1px solid #f8f8f8;
  border-top: none;
  position: relative;
  width: 1310px;
  height: 360px;
  margin: auto;
  display: flex;
  flex-direction: column;
  box-shadow: 0 2px 10px 0 rgb(0 0 0 / 10%);
}

.quest-header {
  width: 100%;
  height: 50px;
  border-bottom: 1px solid #f8f8f8;
  display: flex;
  flex-direction: row;
}

.quest-main {
  width: 100%;
  height: 310px;
  background: #fff;
  display: flex;
  flex-direction: row;
}

.quest-header button {
  outline: none;
  border: none;
  background: transparent;
  cursor: pointer;
  transition: .15s ease;
}

.quest-header button:hover {
  opacity: 0.5;
}

.quest-header .quest-back {
  border-right: 1px solid #f8f8f8;
  font-size: 18px;
  color: #d01d40;
  padding: 0 15px 0 15px;
}

.quest-header .quest-title {
  width: 1205px;
  height: 50px;
  padding: 10px;
  font-size: 18px;
  line-height: 30px;
  color: #333;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-align: left;
}

.quest-image {
  padding: 10px;
  width: 385px;
  height: 285px;
  overflow: hidden;
}

.quest-image .el-image {
  height: 285px;
}

.quest-main .el-descriptions {
  width: 925px;
  padding-top: 20px;
}

.details-container .quest-response-main ul {
  display: block;
  list-style: none;
  list-style-type: disc;
  margin: auto;
  width: 1275px;
}

.details-container .quest-response-main li {
  width: 1310px;
  float: left;
  position: relative;
  z-index: 1;
  display: inline;
  margin: 5px 0 5px -38px;
  border: #f8f8f8 1px solid;
  box-shadow: 0 2px 10px 0 rgb(0 0 0 / 10%);
}

.details-container .quest-response-main li .el-descriptions {
  margin: 20px 0 20px 180px;
}

.quest-response-main {
  background: #fff;
}
</style>