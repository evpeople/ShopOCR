<template>
  <div>
    <div class="ceiling"></div>
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
    <el-empty description="暂无记录" v-show="totalItem == 0" style="height: 650px;"></el-empty>
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
      pageSize: 12,
      historyList: [],
      reset: true,
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
      console.log(this.offset)
      console.log(this.offend)
      // this.downloadHistory(this.currentPage,this.pageSize);
      this.reset = true;
    },
    //下载当前页面历史数据
    downloadHistory(currentPage, pageCount){
      let config = {
          headers: {'Content-Type': 'application/json'}
      };
      // let token = localStorage.getItem("sk_token")
      let currentOffset = (currentPage-1) * pageCount;
      this.$axios.get("/v1/user/request", { params: { offset: currentOffset, limit: pageCount } }, config).then((response) => {
        if (response.data === 1){
          this.$message({
            showClose: true,
            message: '登录状态错误！请重新登录。',
            type: 'error'
          });
          // this.$bus.$emit("changeStatus",{status: false, uname:''});
        }
        else {
          console.log('我的寻味道：');
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
      // console.log(totalItem)
      // this.totalItem = 24;
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
          this.$bus.$emit("itemDetailsRecvInfo", item)
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
  },
  activated() {
    this.downloadHistory(1, this.pageSize);
    console.log(this.offset)
    console.log(this.offend)
    console.log(this.historyList)
    console.log(this.historyList.slice(this.offset, this.offend))
  },
}
</script>

<style scoped>
.ceiling {
  background-image: linear-gradient(135deg, #ed678a 0%, #ffc6a5 65%, #e63e7e 100%);
  width: 100%;
  height: 20px;
}

.item-display {
  width: 1160px;
  height: 650px;
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
  margin-right: 15px;
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

.load-container {
  height: 429px;
  margin: 0 0 5px 0;
}
</style>