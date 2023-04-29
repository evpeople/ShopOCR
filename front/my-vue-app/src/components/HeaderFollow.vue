<template>
  <div class="header-content">
        <el-row :gutter="10">
            <el-col :span="3">
                <div class="header-content-module">
                    <div class="header-content-title">
                        <a href=""><img src="@/assets/logo.png" alt="好味道"></a>
                    </div>
                </div>
            </el-col>
            <el-col :span="21">
                <div class="header-query-module">
                   <el-row class="query-container">
                        <!-- <button class="filter-first" :style="filterFirstBorder" @click="changeFilterState()">
                            <div class="filter-holder">
                                <i class="iconfont icon-shaixuan" style="vertical-align: middle"></i>
                                <span>寻味道</span>
                            </div>
                        </button> -->
                        <el-col :span="3">
                            <div>
                                <el-select v-model="searchValue" placeholder="请选择味道类型" :popper-append-to-body="false" @change="searchText='';isSearch=false">
                                    <el-option
                                    v-for="item in options"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value"
                                    :disabled="item.disabled"
                                    >
                                    </el-option>
                                </el-select>
                            </div>
                        </el-col>
                        <el-col :span="17">
                            <div class="filter-third">
                                <el-input
                                    placeholder="锁定关键词，开始寻找想要的味道~"
                                    v-model="searchText"
                                    ref="search"
                                    clearable
                                    prefix-icon="el-icon-search"
                                    @keyup.enter.native="searchNewHistory()"
                                ></el-input>
                            </div>
                        </el-col>
                        <el-col :span="4">
                            <button class="filter-last" :style="filterLastBackground" @click="filterLastEvent()">
                                <div v-show="!isSearch" class="filter-starter">
                                    <i class="el-icon-search" style="vertical-align: middle"></i>
                                    <span>  寻味道</span>
                                </div>
                                <!-- <div v-show="isSearch" class="filter-closer">
                                    <i class="el-icon-close" style="vertical-align: middle"></i>
                                    <span>  取消</span>
                                </div> -->
                            </button>
                        </el-col>
                    </el-row>
                </div>
            </el-col>
        </el-row>
    </div>
</template>

<script>
export default {
  name: 'ComHeaderQuery',
  data() {
    return {
      options: [{
        value: '0',
        label: '所有',
        disabled: false,
      }, {
        value: '1',
        label: '家乡小吃',
        disabled: false,
      }, {
        value: '2',
        label: '地方特色',
        disabled: false,
      }, {
        value: '3',
        label: '香辣风味',
        disabled: false,
      }, {
        value: '4',
        label: '甜酸风味',
        disabled: false,
      }, {
        value: '5',
        label: '绝一味菜',
        disabled: false,
      }],
      searchValue: '',
      searchText: '',
      isSearch: false,
    }
  },
  computed: {
    filterLastBackground(){
        return { 'background-image': 'linear-gradient(135deg, #c71c5f 0%, #d31e37 100%)' };
    }
  },
  methods:{
    // changeFilterState() {
    //   this.isFiltrate = !this.isFiltrate;
    // },
    filterLastEvent(){
      if (!this.isSearch) {
        this.$router.push({
              name: 'seek-taste'
          })
        this.searchNewHistory();
      } else {
        // this.getAllHistory();
        console.log('正常来说不会进入这个分支。。')
      }
    },
    searchNewHistory(){
      this.$bus.$emit('searchTasteQuery', {"filterType": this.searchValue, "filterValue": this.searchText});
      // this.isSearch = true;
    },
    // getAllHistory(){
    //   this.$bus.$emit('noFilter');
    //   this.$bus.$emit('updateHistoryPage');
    //   this.isSearch = false;
    // }
  }
}
</script>

<style scoped>
.header-content {
    height: 120px;
    background: #fff;
    position: relative;
}

.header-content-module {
    float: right;
}

.header-content-module .header-content-title {
    height: 120px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.header-content-module .header-content-title img {
    width: 76px;
    height: 76px;
    border: 0;
}

.header-query-module {
    height: 120px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.query-container button {
	border: none;
	outline: none;
}
.query-container {
  width: 1300px;
  height: 75px;
  margin: 0 auto;
  padding-bottom: 10px;
  display: flex;
  justify-content: start;
  align-items: center;
  flex-direction: row;
}
.filter-first {
  width: 7%;
  height: 40px;
  justify-content: center;
  align-items: center;
  display: flex;
  background-image: linear-gradient(45deg, #7de1ff 0%, #85FFBD 100%);
  cursor: pointer;
  transition: 0.15s ease;
}
.filter-holder {
  color: #ffffff;
  font-size: 14px;
  font-weight: bold;
  transition: 0.15s ease;
}
.filter-first:hover {
  opacity: .8;
}
.filter-first:hover .filter-holder{
  color: #1d5a7d;
}
.filter-last {
  width: 120px;
  height: 40px;
  justify-content: center;
  align-items: center;
  display: flex;
  border-radius: 0 40px 40px 0;
  cursor: pointer;
  transition: 0.15s ease;
}
.filter-last:hover {
  opacity: 0.8;
}
.filter-starter,.filter-closer {
  color: #ffffff;
  transition: 0.15s ease;
  font-size: 14px;
  font-weight: bold;
}
.filter-last:hover .filter-starter,.filter-closer{
  color: #e5e3e3;
}
.query-container .el-select >>> .el-input__inner {
  border-radius: 0;
  /* border-left: transparent; */
}
.query-container .el-select >>> .el-input.is-focus .el-input__inner {
  border-color: #c71c5f !important;
}
.query-container .el-select >>> .el-input__inner:focus {
  border-color: #c71c5f !important;
}
.query-container .el-select >>> .el-scrollbar .el-select-dropdown__item.selected {
  color: #c71c5f;
}
.query-container .el-input >>> .el-input__inner {
  border-radius: 0;
}
.filter-third .el-input >>> .el-input__inner:focus {
  border-color: #c71c5f !important;
}
.filter-third .el-date-editor.el-input__inner {
  border-radius: 0;
}
.filter-third .el-range-editor.is-active {
  border-color: #c71c5f !important;
}
</style>