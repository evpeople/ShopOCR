<template>
  <div class="main">
    <div class="home-container">
      <div class="navi-module">
        <div class="seek-module">
          <div class="module-title">
            寻味道
          </div>
          <div class="module-content seek">
            <div class="module-table seek" v-show="!isLogin">登录后即可查看</div>
            <div class="module-table seek" v-show="isLogin">中午好！开始今天的寻味道之旅吧！</div>
          </div>
        </div>
        <div class="offer-module">
          <div class="module-title">
            请品鉴
          </div>
          <div class="module-content offer">
            <div class="module-table offer" v-show="!isLogin">登录后即可查看</div>
            <div class="module-table offer" v-show="isLogin">中午好！不要吝啬分享囊中的好味道！</div>
          </div>
        </div>
        <div class="user-module">
          <div class="user-default">
            <div class="user-avatar">
              <img :src="homepageUserAvatar" alt="">
            </div>
            <p class="user-name" v-show="!isLogin">
              Hi! 你好
            </p>
            <p class="user-name" v-show="isLogin">
              Hi! {{ userName }}
            </p>
            <el-button round class="first-button" @click="toPassport1" v-show="!isLogin">立即登录</el-button>
            <el-button round class="second-button" @click="toPassport2" v-show="!isLogin">注册</el-button>
            <el-button round class="first-button" @click="toUser" v-show="isLogin">个人资料</el-button>
            <el-button round class="second-button" @click="userLogout" v-show="isLogin">登出</el-button>
          </div>
          <div class="user-details" @click="toMySeek" v-show="!adminTunnel">
            <img src="@/assets/icon-details.png" alt="">
            <p>用户历史发布详细页面<i class="el-icon-caret-right"></i></p>
          </div>
          <div class="user-details" @click="toStat" v-show="adminTunnel">
            <img src="@/assets/static-icon.png" alt="">
            <p>业务明细及统计信息页面<i class="el-icon-caret-right"></i></p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
    name: 'ComMain',
    data() {
      return {
        isLogin: false,
        userName: '未知用户',
      }
    },
    computed: {
      homepageUserAvatar(){
        if (!this.isLogin) {
          return require('@/assets/icon-nologin.png');
        } else {
          return require('@/assets/icon-user.png');
        }
      },
      adminTunnel() {
        if(this.$cookies.get("sk_uid") == 10) {
          return true
        } else {
          return false
        }
      }
    },
    methods: {
      toPassport1() {
          this.$router.push({
              name: 'login'
          })
      },
      toPassport2() {
          this.$router.push({
              name: 'register'
          })
      },
      toMySeek() {
          this.$router.push({
              path: '/my/seek'
          })
          this.$bus.$emit("MySwitch",1)
      },
      toStat() {
          this.$router.push({
              name: 'stat'
          })
      },
      toUser() {
          this.$router.push({
              name: 'user'
          })
      },
      userLogout(){
        // 注销
        localStorage.removeItem("sk_token")
        this.isLogin = false;
        // let logoutForm = new FormData();
        // logoutForm.append("token", this.getToken());
        // logoutForm.append("uid", this.getUID());
        // this.$axios.post("/NBI/User/logout/",logoutForm, {
        //   headers: {'Content-Type': 'multipart/form-data'}
        // }).then((response)=>{
        //   console.log(response)
        //   if (response.data === 1){
        //     this.$message({
        //       showClose: true,
        //       message: '注销失败，您的账号状态错误！',
        //       type: 'error'
        //     });
        //     this.$bus.$emit("changeStatus",{status: false, uname:''});
        //   }
        //   else if (response.data === 2){
        //     this.$message({
        //       showClose: true,
        //       message: '您已注销',
        //       type: 'warning',
        //     });
        //     this.isLogin = false;
        //     this.$bus.$emit('changeStatus', {status: false, uname: ""});
        //   }
        // })
        this.$bus.$emit("logStatus", true)
        location.reload()
      },
    },
    activated() {
      if(localStorage.getItem('sk_token') && this.isLogin) {
        let uname = this.$cookies.get("sk_uname");
        if (uname !== '') {
          this.userName = uname;
        }
      }
      else {
        this.isLogin = false;
        this.userName = '匿名用户'
      }
    },
    mounted() {
      this.$bus.$on("logStatus", (data) => {
        this.isLogin = data;
      })
    },
    beforeDestroy() {
      this.$bus.$off("logStatus")
    },
}
</script>

<style scoped>
a {
  text-decoration: none;
}

.main {
  width: 100%;
  height: 560px;
  background: linear-gradient(180deg,#F8F8F8,rgba(255,255,255,0) 100%);
}

.home-container {
  margin: 0 auto;
  width: 1310px;
  position: relative;
  padding-top: 30px;
}

.navi-module {
  position: relative;
}

.seek-module {
  width: 480px;
  height: 480px;
  border:1px solid transparent;
  border-right: none;
  border-bottom: none;
  border-radius: 30px 0 0 30px;
  background-color: #FF99AC;
  background-image: linear-gradient(90deg, #FF99AC 0%, #FF6A88 55%, #FF9A8B 100%);
  float: left;
  display: flex;
  flex-direction: column;
}

.module-title {
  padding-left: 30px;
  padding-top: 8px;
  height: 50px;
  line-height: 50px;
  font-size: 24px;
  font-weight: bold;
  color: #fff;
  text-align: left;
}

.module-content.seek {
  margin-left: 30px;
  flex: 1;
  background: rgba(255, 255, 255, 0.8);
  border: none;
}

.module-content.offer {
  margin-right: 30px;
  flex: 1;
  background: rgba(255, 255, 255, 0.8);
  border: none;
}

.offer-module {
  margin-left: 20px;
  width: 480px;
  height: 480px;
  border:1px solid transparent;
  border-left: none;
  border-bottom: none;
  border-radius: 0 30px 30px 0;
  background-color: #FBAB7E;
  background-image: linear-gradient(62deg, #FBAB7E 0%, #F7CE68 100%);
  float: left;
  display: flex;
  flex-direction: column;
}

.module-table {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: bold;
}

.module-table.seek {
  color: #FF6A88;
}

.module-table.offer {
  color: #FBAB7E;
}

.user-module {
  margin-left: 40px;
  width: 280px;
  height: 480px;
  float: left;
}

.user-default {
  border:1px solid #C0C4CC;
  background-color: #fff;
  border-radius: 3px;
  width: 280px;
  height: 240px;
  padding-top: 40px;
}

.user-avatar {
  margin: 0px auto 4px auto;
}

.user-avatar img{
  width: 47px;
  height: 47px;
  border-radius: 50%;
  border: 4px solid #E5E5E5;
}

.user-name {
  font-size: 18px;
  color: #222;
  text-align: center;
  font-weight: 500;
  white-space: nowrap;
  width: 6em;
  overflow: hidden;
  text-overflow: ellipsis;
  margin: 0 auto;
}

.el-button {
  width: 120px;
  height: 40px;
  text-align: center;
  cursor: pointer;
  transition: 0.15s ease;
  margin: 15px auto 20px auto;
  background: #fff;
  border: 1px solid #e5e5e5;
  border-radius: 40px;
  font-size: 14px;
  color: #333;
  display: block;
}
.el-button--default.is-round.first-button {
    color: #fff;
    background: #ed2d78;
    border-color: #ed2d78;
}
.el-button--default.is-round.second-button {
    color: #333;
    background: #fff;
    border-color: #e5e5e5;
}
.el-button.is-round.first-button:focus, .el-button.is-round.first-button:hover {
    background: #d31e37;
    border-color: #d31e37;
    color: #fff;
}
.el-button.is-round.second-button:focus, .el-button.is-round.second-button:hover {
    background: #fff4f6;
    border-color: #FF99AC;
    color: #FF99AC;
}

.user-details {
  margin-top: 20px;
  border:1px solid #C0C4CC;
  background-image: linear-gradient(135deg, #fff 0%, #fff 85%, #ffaec0 100%);
  border-radius: 3px;
  width: 280px;
  height: 150px;
  padding-top: 30px;
  transition:  0.15s ease;
  cursor: pointer;
}

.user-details:hover {
  opacity: 0.7;
}

.user-details img {
  width: 87px;
  height: 87px;
  border-radius: 50%;
  border: 4px solid #e66a76;
}

.user-details p {
  font-weight: bold;
  color: #e66a76;
}
</style>