<template>
  <div class="stat-container">
    <div class="stat-header">
        <el-cascader
            :options="areaSelectData"
            @change="handleChange"
            class="full-width"
            size="large"
            placeholder="选择城市"
        />
        <div style="width: 40px;line-height: 40px;">|</div>
        <el-date-picker
            v-model="date1"
            type="month"
            placeholder="选择起始月">
        </el-date-picker>
        <div style="width: 40px;line-height: 40px;">至</div>
        <el-date-picker
            v-model="date2"
            type="month"
            placeholder="选择截止月">
        </el-date-picker>
        <div style="width: 40px;line-height: 40px;">|</div>
        <button class="filter-last" @click="filterLastEvent()">
            <div class="filter-starter">
                <i class="el-icon-search" style="vertical-align: middle"></i>
                <span>  查询</span>
            </div>
        </button>
    </div>
    <div class="stat-main">
        <ve-line :data="chartData" :settings="chartSettings" height="635px" style="background: #fff;"></ve-line>
        <el-table
            :data="recvData"
            stripe
            style="width: 100%"
            >
            <el-table-column
            prop="ID"
            label="序列号"
            width="200">
            </el-table-column>
            <el-table-column
            prop="request_people_id"
            label="发布人"
            width="200">
            </el-table-column>
            <el-table-column
            prop="response_people_id"
            label="响应人"
            width="200">
            </el-table-column>
            <el-table-column
            prop="end_date"
            label="达成年月"
            width="200">
            </el-table-column>
            <el-table-column
            prop="request_people_fee"
            label="发布人中介费"
            width="200">
            </el-table-column>
            <el-table-column
            prop="response_people_fee"
            label="响应人中介费"
            width="200">
            </el-table-column>
            <el-table-column
            prop="total_fee"
            label="总中介费"
            width="200"
            sortable>
            </el-table-column>
        </el-table>
    </div>
  </div>
</template>

<script>
import { regionData, CodeToText, TextToCode } from "element-china-area-data";

export default {
    name: 'StatisticPage',
    data() {
        this.chartSettings = {
            area: true
        }
        return {
            areaSelectData: regionData,
            date1: '',
            date2: '',
            searchForm: {
                location: {
                    country: "北京市",
                    province: "市辖区",
                    city: "海淀区"
                },
            },
            recvData:[],
            chartData:{
                columns: ['Date', 'TotalDeals', 'TotalFees'],
                rows: []
            },
            tempRows:[]
        }
    },
    methods: {
        handleChange(value) {
            //value为省市区code数组
            if (value) {
                var provinceCode = CodeToText[value[0]];//code转为省
                var cityCode = CodeToText[value[1]];//市
                var orgion = CodeToText[value[2]];//区
                this.searchForm.location.country = provinceCode;
                this.searchForm.location.province = cityCode;
                this.searchForm.location.city = orgion;
                // this.form.selectedOptions = provinceCode + cityCode + orgion;
            }
        },
        compare(property){
            return function(a,b){
                var value1 = a[property];
                var value2 = b[property];
                return value1 - value2;
            }
        },
        showEditDialog() {//showEditDialog为点击编辑数据显示的按钮
            this.getList();
            this.dialogConsignee = true;
            this.form.selectedOptions = [
                TextToCode[this.list.province].code,
                TextToCode[this.list.province][this.list.city].code,
                TextToCode[this.list.province][this.list.city][this.list.area].code,
            ];
        },
        classifyYear(date) {
            var year = date.getFullYear(); //获取时间，年
            return parseInt(year)
        },
        classifyMonth(date) {
            var month = date.getMonth() + 1; //获取时间 月 月是0-11，正常显示月份要+1
            return parseInt(month)
        },
        filterLastEvent() {
            var begin_date_year = this.classifyYear(this.date1)
            var begin_date_month = this.classifyMonth(this.date1)
            var end_date_year = this.classifyYear(this.date2)
            var end_date_month = this.classifyMonth(this.date2)
            let config = {
                headers: {'Content-Type': 'application/json'}
            };
            this.$axios.get('/v1/admin/money',
            { params: { begin_date_year: begin_date_year, begin_date_month: begin_date_month, end_date_year: end_date_year, end_date_month: end_date_month,
                        country: this.searchForm.location.country, province: this.searchForm.location.province, city: this.searchForm.location.city }}, config)
            .then((response)=>{
                console.log(response.data)
                if (response.data===1){
                this.isLogin = false;
                this.$message({
                    showClose: true,
                    message: '没有找到对应信息',
                    type: 'error'
                });
                }else{
                console.log('管理员操作成功')
                var data = response.data;
                this.recvData.splice(0);
                for (var key in data) {
                    var item = data[key];
                    // console.log(item)
                    this.recvData.push({
                    ID: item.ID,
                    request_people_id: item.request_people_id,
                    response_people_id: item.response_people_id,
                    end_date: String(item.end_date_year)+'/'+String(item.end_date_month),
                    request_people_fee: item.request_people_fee*6,
                    response_people_fee: item.response_people_fee,
                    total_fee: item.request_people_fee*6+item.response_people_fee,                   
                    })
                    var end = ''
                    if (item.end_date_month<10) {
                        end = '0'+String(item.end_date_month)
                    } else {
                        end = String(item.end_date_month)
                    }
                    var dateID = String(item.end_date_year)+'/'+end;
                    // console.log('BBBBBBOOOORRRDERLINE')
                    // console.log(this.chartData.rows.find(object => object.Date === dateID))
                    if (typeof(this.tempRows.find(object => object.Date === dateID)) == "undefined") {
                        this.tempRows.push({
                            Date: dateID,
                            TotalFees: item.request_people_fee*6+item.response_people_fee,
                            TotalDeals: 1
                        })
                    } else {
                        this.tempRows.find(object => object.Date === dateID).TotalFees += item.request_people_fee*6+item.response_people_fee;
                        this.tempRows.find(object => object.Date === dateID).TotalDeals += 1;
                    }
                }
                console.log('准备输出chartData！')
                this.chartData.rows = (this.tempRows.sort(this.compare('Date')))
                console.log(this.chartData.rows)
                }
            });
        }
    },
}
</script>

<style scoped>
.stat-header button {
    outline: none;
    border: none;
}

.stat-container {
  width: 100%;
  min-height: 680px;
  background: linear-gradient(180deg,#fd7b90,#d01d40 100%);
  color: #fff;
}

.stat-header {
    padding-top: 40px;
    display: flex;
    justify-content: center;
}

.filter-last {
  width: 120px;
  height: 40px;
  justify-content: center;
  align-items: center;
  display: flex;
  border-radius: 5px 40px 40px 5px;
  cursor: pointer;
  transition: 0.15s ease;
  background-image: linear-gradient(135deg, #c71c5f 0%, #d31e37 100%);
}
.filter-last:hover {
  opacity: 0.8;
}
.filter-last:hover .filter-starter,.filter-closer{
  color: #e5e3e3;
}
.filter-starter,.filter-closer {
  color: #ffffff;
  transition: 0.15s ease;
  font-size: 14px;
  font-weight: bold;
}

.stat-main {
    width: 1310px;
    height: auto;
    margin: auto;
    padding-top: 40px;
}
</style>