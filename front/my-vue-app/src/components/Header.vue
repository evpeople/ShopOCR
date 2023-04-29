<template>
  <div class="header">
    <div class="header-bar">
        <el-row :gutter="10">
            <el-col :span="3">
                <div class="header-bar-position" style="float: right;">
                    <el-popover
                        placement="bottom"
                        width="400"
                        trigger="hover">
                        <el-button @click="locationChange('北京')">北京</el-button>
                        <el-button @click="locationChange('上海')">上海</el-button>
                        <el-button @click="locationChange('广州')">广州</el-button>
                        <el-button @click="locationChange('深圳')">深圳</el-button>
                        <el-button @click="locationChange('天津')">天津</el-button>
                        <el-button slot="reference"><i class="el-icon-location"></i>{{ location}}</el-button>
                    </el-popover>
                </div>
            </el-col>
            <el-col :span="14">
                <div class="header-bar-nav" style="float: left;">
                    <el-button slot="reference" @click="toHome" autofocus>首页</el-button>
                    <el-button slot="reference" @click="toSeek">寻味道</el-button>
                </div>
            </el-col>
            <el-col :span="7">
                <div class="header-bar-user">
                    <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
                        <el-menu-item index="1" @click="toMySeek">我的寻味道</el-menu-item>
                        <el-menu-item index="2" @click="toMyOffer">我的请品鉴</el-menu-item>
                        <el-menu-item index="3" @click="toUser">个人中心</el-menu-item>
                    </el-menu>
                </div>
            </el-col>
        </el-row>
    </div>
  </div>
</template>

<script>
export default {
    name: 'ComHeader',

    data() {
        return {
            location: '北京',
        }
    },
    methods: {
        locationChange(newLocation) {
            this.location = newLocation;
        },
        toHome() {
            this.$router.push({
                name: 'home'
            })
        },
        toUser() {
            this.$router.push({
                name: 'user'
            })
        },
        toSeek() {
            this.$router.push({
                name: 'seek-taste'
            })
        },
        toMySeek() {
            this.$router.push({
                path: '/my/seek'
            })
            this.$bus.$emit("MySwitch",1)
        },
        toMyOffer() {
            this.$router.push({
                path: '/my/offer'
            })
            this.$bus.$emit("MySwitch",2)
        }
    },
}
</script>

<style scoped>
a {
    text-decoration:none;
}

.header {
    position: relative;
}

.header-bar {
    width: 100%;
    height: 40px;
    background-color: #F8F8F8;
    overflow: hidden;
    z-index: 100;
}

.header-bar-user {
    height: 40px;
}

.el-button {
    height: 40px;
    font-size: 12px;
    color: #909399;
    background: transparent;
    border: none;
}

.el-button:hover, .el-button:focus {
    color: #000;
    background: #fff;
}

.el-menu {
    background: transparent;
}

.el-menu-item {
    height: 40px !important;
    line-height: 40px !important;
    font-size: 12px;
}

.el-submenu >>> .el-submenu__title {
    height: 40px !important;
    line-height: 40px !important;
    font-size: 12px !important;
}

.el-menu--horizontal>.el-menu-item.is-active {
    border-bottom: 2px solid #ff6e8b;
    color: #303133;
}
</style>