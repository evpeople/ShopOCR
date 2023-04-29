<template>
  <div>
    <div class="ceiling"></div>
    <div class="item-display" v-show="totalItem > 0">
      <ul v-if="itemVisible">
        <li v-for="(item, index) in this.historyList.slice(this.offset, this.offend)" :key="index">
          <OfferResponseItem :text="item.offer_msg" :responser="item.offer_people_id" :createDate="item.offer_begin_time" 
          :editDate="item.offer_change_time" :status="item.offer_state" :requester="item.request_id"></OfferResponseItem>
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
import OfferResponseItem from '@/components/OfferResponseItem.vue'
import HistoryLoad from '@/components/Load.vue'

export default {
  name: 'DetailsSeek',
  components: {
    OfferResponseItem,
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
      // this.downloadHistory(this.currentPage,this.pageSize);
      this.reset = true;
    },
    //下载当前页面历史数据
    downloadHistory(currentPage, pageCount){
      console.log('into the my offer download')
      let config = {
          headers: {'Content-Type': 'application/json'}
      };
      // let token = localStorage.getItem("sk_token")
      let currentOffset = (currentPage-1) * pageCount;
      // let request_id = ""
      this.$axios.get(`/v1/response/please`, { params: { offset: currentOffset, limit: pageCount }}, config).then((response) => {
        if (response.data === 1){
          this.$message({
            showClose: true,
            message: '登录状态错误！请重新登录。',
            type: 'error'
          });
          // this.$bus.$emit("changeStatus",{status: false, uname:''});
        }
        else {
          // console.log(response.data);
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
          response_id: item.response_id,
          request_id: item.request_id,
          offer_people_id: item.offer_people_id,
          offer_msg: item.offer_msg,
          offer_begin_time: item.offer_begin_time,
          offer_change_time: item.offer_change_time,
          offer_state: item.offer_state,
        })
      }
      console.log(this.historyList);
      this.reset = true;//重建组件
      // console.log(this.itemVisible+'后');
    }
  },
  mounted() {
    this.downloadHistory(this.currentPage, this.pageSize);
  },
  activated() {
    this.downloadHistory(1, this.pageSize);
  },
}
</script>

<style scoped>
.ceiling {
  background-image: linear-gradient(135deg, #f9bf71 0%, #ff9987 65%, #ff8742 100%);
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
</style>