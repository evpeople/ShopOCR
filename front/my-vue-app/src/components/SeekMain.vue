<template>
<div>
    <div class="seek-header">
      <span class="left-style">寻味列表</span>
      <el-button type="danger" icon="el-icon-edit" class="right-style" @click="upload = true">+发布寻味道</el-button>
      <el-drawer
        title="发布新的寻味道"
        :before-close="cancelForm"
        :visible.sync="upload"
        direction="rtl"
        custom-class="upload-drawer"
        ref="drawer"
        size="45%"
        >
        <div class="drawer-content">
            <el-form :model="uploadForm" ref="uploadForm">
                <el-form-item label="味道类型" label-width="250px">
                  <el-col :span="11">
                    <el-select v-model="uploadForm.type" placeholder="请选择味道类型">
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
                    <el-input v-model="uploadForm.title" autocomplete="off"></el-input>
                  </el-col>
                </el-form-item>
                <el-form-item label="可接受最高单价" label-width="250px">
                  <el-col :span="17">
                    <el-input v-model="uploadForm.price" autocomplete="off"></el-input>
                  </el-col>
                </el-form-item>
                <el-form-item label="截止时间" label-width="250px">
                    <el-col :span="11">
                        <el-date-picker type="date" placeholder="选择截止日期" v-model="uploadForm.endDate" style="width: 100%;"></el-date-picker>
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
                    <el-input v-model="uploadForm.descrip" autocomplete="off" type="textarea" :rows="5"></el-input>
                  </el-col>
                </el-form-item>
            </el-form>
            <div slot="footer" class="drawer-footer">
                <el-button @click="cancelForm">取 消</el-button>
                <el-button type="primary" @click="submitForm()">确 定</el-button>
            </div>
        </div>
      </el-drawer>
    </div>
    <div class="seek-container">
    <div class="item-display" v-show="totalItem > 0">
      <ul v-if="itemVisible">
        <li v-for="(item, index) in this.historyList.slice(this.offset, this.offend)" :key="index">
          <SeekQuestItem :id="item.request_id" :imgSrc="item.topic_photo" :price="item.topic_price" :type="item.taste_type" 
          :title="item.topic_name" :uploader="item.user_id" :startDate="item.topic_begin_date" :endDate="item.topic_end_date"
          :status="item.topic_state"></SeekQuestItem>
        </li>
      </ul>
      <div v-if="!itemVisible" class="load-container">
        <HistoryLoad/>
      </div>
    </div>
    <el-empty description="暂无记录" v-show="totalItem == 0" style="height: 1000px;"></el-empty>
    <div class="pagination">
      <el-pagination
        @current-change="handleCurrentChange"
        :current-page.sync="currentPage"
        :page-size="pageSize"
        layout="prev, pager, next, jumper"
        :total="totalItem">
      </el-pagination>
    </div>
  </div>
</div>
</template>

<script>
import SeekQuestItem from '@/components/SeekQuestItem.vue'
import HistoryLoad from '@/components/Load.vue'

export default {
  name: 'DetailsSeek',
  components: {
    SeekQuestItem,
    HistoryLoad
  },
  data() {
    return {
      totalPage: 0,
      totalItem: 0,
      currentPage: 1,
      pageSize: 21,
      historyList: [],
      reset: true,
      upload: false,
      uploadForm: {
        type: '',
        title: '',
        descrip: '',
        price: '',
        endDate: '',
        image: '',
      },
      type: '',
      keyWord: ''
    };
  },
  computed: {
    itemVisible() {
      return this.reset;
    },
    offset() {
      return (this.currentPage-1) * this.pageSize
    },
    offend() {
      return this.currentPage * this.pageSize
    },
  },
  methods: {
    handleCurrentChange(val) {
      this.reset = false;
      // console.log(this.currentPage)
      console.log(val)
      // this.downloadHistory(this.currentPage,this.pageSize);
      this.reset = true;
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
    resetForm() {
      this.uploadForm.type = "";
      this.uploadForm.title = "";
      this.uploadForm.descrip = "";
      this.uploadForm.price = "";
      this.uploadForm.endDate = "";
      this.uploadForm.image = "";
      this.$refs.uploadForm.clearValidate();
    },
    submitForm() {
      let nowDate = this.classifyDate(new Date());
      let nowDate_toString = nowDate.toString();
      let endDate_toString = this.classifyDate(this.uploadForm.endDate).toString();
      let config = {
          headers: {'Content-Type': 'application/json'}
      }; 
      this.$axios.post(`/v1/user/request`, {
        "request_id": 3340,
        "user_id": parseInt(this.$cookies.get("sk_uid")),
        "taste_type": this.uploadForm.type,
        "topic_name": this.uploadForm.title,
        "topic_intro": this.uploadForm.descrip,
        "topic_photo": this.uploadForm.image,
        "topic_begin_date": nowDate_toString,
        "topic_end_date": endDate_toString,
        "topic_change_date": nowDate_toString,
        "topic_state": "1",
        "topic_price": parseFloat(this.uploadForm.price),
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
              this.upload = false;
              this.resetForm();
              this.cancelForm();
              this.$router.push({
                  name: 'blank'
              })
          }
      });
    },
    cancelForm() {
      this.upload = false;
    },
    //下载当前页面历史数据
    downloadHistory(currentPage, pageCount){
      let config = {
          headers: {'Content-Type': 'application/json'}
      };
      let currentOffset = (currentPage-1) * pageCount;
      let country = localStorage.getItem("sk_locate1")
      let province = localStorage.getItem("sk_locate2")
      let city = localStorage.getItem("sk_locate3")
      this.$axios.get("/v1/response/location", { params: { offset: currentOffset, limit: pageCount, country: country, province: province, city: city } }, config).then((response) => {
        if (response.data === 1){
          this.$message({
            showClose: true,
            message: '登录状态错误！请重新登录。',
            type: 'error'
          });
          // this.$bus.$emit("changeStatus",{status: false, uname:''});
        }
        else {
          console.log('地区寻味道：');
          console.log(response.data.data);
          var totalPage = Math.ceil(response.data.total / pageCount);
          this.loadHistory(response.data.data, totalPage, response.data.total)
        }
      })
    },
    //携带筛选的版本
    downloadHistoryWithFilter(currentPage, pageCount){
      let config = {
          headers: {'Content-Type': 'application/json'}
      };
      let currentOffset = (currentPage-1) * pageCount;
      let country = localStorage.getItem("sk_locate1")
      let province = localStorage.getItem("sk_locate2")
      let city = localStorage.getItem("sk_locate3")
      this.$axios.get("/v1/response/location", { params: { offset: currentOffset, limit: pageCount, country: country, province: province, city: city, type: this.type, keyWord: this.keyWord } }, config).then((response) => {
        if (response.data === 1){
          this.$message({
            showClose: true,
            message: '登录状态错误！请重新登录。',
            type: 'error'
          });
          // this.$bus.$emit("changeStatus",{status: false, uname:''});
        }
        else {
          console.log('地区寻味道(筛选启用)：');
          console.log(response.data.data);
          var totalPage = Math.ceil(response.data.total / pageCount);
          this.loadHistory(response.data.data, totalPage, response.data.total)
        }
      })
    },
    //载入下载的历史数据
    loadHistory(data,totalPage,totalItem){
      // console.log(totalPage);
      // for (var key in data) {
      //   var item = data[key];
      //   console.log(item);
      // }
      this.totalPage = totalPage;
      this.totalItem = totalItem;
      this.historyList.splice(0);
      for (var key in data) {
        var item = data[key];
        console.log(item)
        this.historyList.push({
          request_id: item.request_id,
          user_id: item.user_id,
          taste_type: item.taste_type,
          topic_name: item.topic_name,
          topic_intro: item.topic_intro,
          topic_photo: item.topic_photo,
          topic_begin_date: item.topic_begin_date,
          topic_end_date: item.topic_end_date,
          topic_change_date: item.topic_change_date,
          topic_state: item.topic_state,
          topic_price: item.topic_price,
          topic_location: item.topic_location,
          response: item.response,
        })
      }
      // // console.log(this.historyList);
      this.reset = true;//重建组件
      // console.log(this.itemVisible+'后');
    },
    sendItemInfo(id) {
      for (var key in this.historyList) {
        var item = this.historyList[key]
        console.log(item)
        if(item.request_id == id) {
          console.log('find it! the required details item. the id is' + item.request_id)
          this.$bus.$emit("itemDetailsRecvInfoFromAll", item)
          return
        }
      }
    }
  },
  mounted() {
    this.downloadHistory(this.currentPage, this.pageSize);
    this.$bus.$on("itemDetailsNeedInfo", (id) => {
      console.log(id+'号寻味道需要信息')
      this.sendItemInfo(id)
    })
    this.$bus.$on("searchTasteQuery", (data) => {
      if (data.filterType == 0) {
        this.type = ''
      } else {
        this.type = data.filterType;
      }
      this.keyWord = data.filterValue;
      this.downloadHistoryWithFilter(this.currentPage, this.pageSize);
    })
  },
  activated() {
    this.downloadHistory(1, this.pageSize);
  },
  beforeDestroy() {
    this.$bus.$off("itemDetailsNeedInfo")
    this.$bus.$off("searchTasteQuery")
  },
}
</script>

<style scoped>
.seek-container {
  margin: 0 auto;
  width: 1310px;
  height: 1050px;
  position: relative;
  border: #ee6a8b 2px solid;
  /* border-top: none; */
  border-bottom: none;
}

.seek-header {
  margin: 0 auto;
  width: 1310px;
  height: 41px;
  border-radius: 20px 20px 0 0;
  border: #ee6a8b 2px solid;
  border-bottom: none;
  background-image: linear-gradient(10deg, #ee6a8b 30%, #ffa9a5 100%);
}

.seek-header > .el-button {
  border-radius: 4px 20px 4px 4px;
}

.left-style {
  display: block;
  float: left;
  margin-left: 28px;
  color: #fff;
  font-size: 20px;
  font-weight: bold;
  line-height: 41px;
  font-family: Helvetica, 'Hiragino Sans GB', 'Microsoft Yahei', '微软雅黑', Arial, sans-serif;
}

.right-style {
  float: right;
  margin-right: 1px;
}

.el-button--danger {
    color: #F56C6C;
    background-color: #fff;
    border-color: #fff;
}

.el-button--danger:hover {
    color: #fff;
    background-color: #F56C6C;
    border-color: #F56C6C;
}

.el-button--danger:focus {
    color: rgb(189, 189, 189);
    background-color: #a82d33;
    border-color: #a82d33;
}

.item-display {
  width: 100%;
  height: 990px;
}

.item-display ul {
  display: block;
  list-style: none;
  list-style-type: disc;
  margin-block-start: 1em;
  margin-block-end: 1em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  margin-top: 0px;
  padding-inline-start: 20px;
}

.item-display li {
  width: 170px;
  height: 315px;
  float: left;
  position: relative;
  z-index: 1;
  display: inline;
  margin-right: 12px;
  margin-top: 10px;
}

.pagination {
  margin-top: 8px;
}

.el-pagination >>> .el-pager li.active, .el-pagination >>> .el-pager li:hover {
  color: #ee6a8b;
}

.el-pagination >>> button:hover {
  color: #ee6a8b;
}

.el-pagination >>> .el-input__inner:focus {
  border-color: #ee6a8b;
}

.drawer-content .el-input {
  width: 305px;
}

.el-upload__tip {
  margin-top: -20px;
  margin-left: 45px;
}
</style>