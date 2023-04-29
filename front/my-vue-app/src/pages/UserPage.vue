<template>
  <div class="user-container">
    <div class="avatar-part user-module">
        <div class="avatar-inner">
            <img src="@/assets/icon-user.png" alt="">
        </div>
        <el-button round class="first-button module-button" @click="isChangePWD = true">修改密码</el-button>
        <el-dialog title="修改密码" :visible.sync="isChangePWD" width="30%">
            <el-form ref="pwdForm" label-width="100px" :model="pwdForm">
                <el-form-item label="原密码" prop="formerPwd">
                    <el-input type="password" v-model="pwdForm.formerPwd" autocomplete="off" show-password></el-input>
                </el-form-item>
                <el-form-item label="新密码" prop="newPwd">
                    <el-input type="password" v-model="pwdForm.newPwd" autocomplete="off" show-password></el-input>
                </el-form-item>
                <el-form-item label="确认密码" prop="checkNewPwd">
                    <el-input type="password" v-model="pwdForm.checkNewPwd" autocomplete="off" show-password></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="resetForm()">重 置</el-button>
                <el-button type="primary" @click="submitForm()">确 定</el-button>
            </div>
        </el-dialog>
        <el-button round class="second-button module-button" @click="userLogout">登 出</el-button>
    </div>
    <div class="profile-part user-module">
        <div class="profile-front">
            <div class="left-title">{{ userName }}&emsp;<el-tag type="info">{{ userLevel }} {{ userType }}</el-tag></div>
            <div class="right-title">最后修改时间：{{ lastChangeTime }}<el-button type="info" plain style="margin-left: 20px;" @click="isChangeInfo = true"><i class="el-icon-edit"></i>修改资料</el-button></div>
            <el-dialog title="修改个人资料" :visible.sync="isChangeInfo" width="30%">
                <el-form ref="infoForm" label-width="100px" :model="infoForm">
                    <el-form-item label="联系电话">
                        <el-input v-model="infoForm.phone" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="个人简介">
                        <el-input v-model="infoForm.intro" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="cancelInfoForm">取 消</el-button>
                    <el-button type="primary" @click="submitInfoForm()">确 定</el-button>
                </div>
            </el-dialog>
        </div>
        <div class="profile-main">
            <el-descriptions class="margin-top" :column="3" border>
                <el-descriptions-item>
                    <template slot="label">
                        <i class="el-icon-user"></i>
                        UID
                    </template>
                    {{ uid }}
                </el-descriptions-item>
                <el-descriptions-item>
                    <template slot="label">
                        <i class="el-icon-office-building"></i>
                        注册时间
                    </template>
                    {{ registerTime }}
                </el-descriptions-item>
                <el-descriptions-item>
                    <template slot="label">
                        <i class="el-icon-location-outline"></i>
                        注册城市
                    </template>
                    {{ registerCity.country }}-{{ registerCity.province }}-{{ registerCity.city }}
                </el-descriptions-item>
                <el-descriptions-item>
                    <template slot="label">
                        <i class="el-icon-tickets"></i>
                        姓名
                    </template>
                    {{ classiName }}
                </el-descriptions-item>
                <el-descriptions-item>
                    <template slot="label">
                        <i class="el-icon-office-building"></i>
                        {{ classiType }}
                    </template>
                    {{ classiID }}
                </el-descriptions-item>
                <el-descriptions-item>
                    <template slot="label">
                        <i class="el-icon-mobile-phone"></i>
                        手机号
                    </template>
                    {{ phone }}
                </el-descriptions-item>
            </el-descriptions>
            <el-descriptions direction="vertical" :column="1" border>
            <el-descriptions-item label="个人简介">
                {{ intro }}
            </el-descriptions-item>
            </el-descriptions>
        </div>
    </div>
  </div>
</template>

<script>
export default {
    name: 'UserPage',
    data() {
      return {
        uid: '00',
        userName: '匿名用户',
        type: 1,
        level: 1,
        classiName: '王小明',
        classiType: '身份证',
        classiID: '450101200001010000',
        phone: '00000000000',
        intro: '该用户很懒，还没有个人简介噢~',
        registerCity: {
          country: "北京市",
          province: "市辖区",
          city: "海淀区"
        },
        registerTime: '2022/12/26',
        lastChangeTime: '2022/12/26',
        pwdForm: {
          formerPwd: '',
          newPwd: '',
          checkNewPwd: ''
        },
        infoForm: {
            phone: '',
            intro: ''
        },
        isChangePWD: false,
        isChangeInfo: false,
        islogin: false
      }
    },
    computed: {
        userType() {
            if (this.type == 1) {
                return '普通用户'
            } else {
                return '管理员'
            }
        },
        userLevel() {
            if (this.type == 1) {
                return '一般'
            } else {
                return 'VIP'
            }
        },
    },
    methods: {
        resetForm() {
            this.pwdForm.formerPwd = "";
            this.pwdForm.newPwd = "";
            this.pwdForm.checkNewPwd = "";
            this.$refs.pwdForm.clearValidate();
        },
        submitForm() {
            if (this.pwdForm.newPwd === "" || this.pwdForm.checkNewPwd === "" || this.pwdForm.formerPwd === ""){
                this.$message({
                showClose: true,
                message: '请完成输入原密码、新密码及确认密码',
                type: 'error',
                });
                return;
            }
            if (this.pwdForm.newPwd !== this.pwdForm.checkNewPwd){
                this.$message({
                showClose: true,
                message: '新密码与确认密码不一致',
                type: 'error',
                });
                return;
            }
            this.$axios.post(`/v1/change`, {
                "id": this.uid,
                "password": this.pwdForm.checkNewPwd,
                "intro": this.intro,
                "telephone": this.phone
            }, {
                headers: {'Content-Type': 'application/json'}
            }).then((response)=>{
                if (response.data === 1){
                    this.$message({
                        showClose: true,
                        message: '您的账号状态错误！',
                        type: 'error',
                    });
                    this.$bus.$emit("changeStatus",{status: false, uname:''});
                }
                else if(response.data === 2){
                    this.$message({
                        showClose: true,
                        message: '原始密码错误，修改失败！',
                        type: 'error',
                    });
                    this.resetForm();
                }
                else {
                    console.log(response.data)
                    this.$message({
                        showClose: true,
                        message: '设置成功',
                        type: 'success',
                    });
                    this.isChangePWD = false;
                    this.resetForm();
                }
            });
        },
        getUserInfo() {
            let config = {
                headers: {'Content-Type': 'application/json'}
            };
            // let token = localStorage.getItem("sk_token")
            this.$axios.get('/v1/info', config).then((response)=>{
                console.log(response.data)
                if (response.data===1){
                this.isLogin = false;
                this.$message({
                    showClose: true,
                    message: '没有找到您的用户信息，您可能是没有登录或通信不畅',
                    type: 'error'
                });
                }else{
                this.isLogin = true;
                this.userName = response.data.user_name;
                this.phone = response.data.telephone;
                this.registerCity = response.data.location;
                this.intro = response.data.intro;
                this.level = response.data.level;
                this.uid = response.data.people_id;
                this.classiType = response.data.cred_type;
                this.classiID = response.data.cred_id;
                this.classiName = response.data.name;
                this.infoForm.intro = response.data.intro;
                this.infoForm.phone = response.data.telephone;
                }
            });
            localStorage.setItem("sk_locate1", this.registerCity.country)
            localStorage.setItem("sk_locate2", this.registerCity.province)
            localStorage.setItem("sk_locate3", this.registerCity.city)
        },
        submitInfoForm() {
            this.$axios.post(`/v1/change`, {
                "id": this.uid,
                "intro": this.infoForm.intro,
                "telephone": this.infoForm.phone
            }, {
                headers: {'Content-Type': 'application/json'}
            }).then((response)=>{
                if (response.data === 1){
                    this.$message({
                        showClose: true,
                        message: '您的账号状态错误！',
                        type: 'error',
                    });
                    this.$bus.$emit("changeStatus",{status: false, uname:''});
                }
                else {
                    console.log(response.data)
                    this.$message({
                        showClose: true,
                        message: '设置成功',
                        type: 'success',
                    });
                    this.isChangeInfo = false;
                }
            });
        },
        cancelInfoForm() {
            this.isChangeInfo = false;
        },
        userLogout() {
            localStorage.removeItem("sk_token")
            this.$bus.$emit("logStatus", false)
            this.$router.push({
              name: 'home'
          })
        }
    },
    mounted() {
        this.getUserInfo();
        this.$bus.$on("logStatus", (data) => {
            this.isLogin = data;
        })
    },
    activated() {
      if(!localStorage.getItem('sk_token') || !this.isLogin) {
        this.isLogin = false;
        this.userName = '匿名用户'
      } else {
        this.getUserInfo();
      }
    },
    beforeDestroy() {
      this.$bus.$off("logStatus")
    },
}
</script>

<style scoped>
.user-container {
  margin: 0 auto;
  width: 1310px;
  height: 680px;
  position: relative;
  background: #fff;
  box-shadow: 0 20px 50px rgb(65 62 101 / 15%);
}

.user-module {
  float: left;
}

.avatar-part {
  width: 450px;
  height: 680px;
  background-image: linear-gradient(135deg, #FF99AC 0%, #FF6A88 55%);
}

.avatar-inner {
  padding-top: 100px;
}

.avatar-inner img{
  width: 300px;
  height: 300px;
  border-radius: 50%;
  border: 20px solid #E5E5E5;
}

.el-dialog .el-input {
  width: 300px;
}

.module-button.el-button {
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
  font-weight: bold;
}
.module-button.el-button--default.is-round.first-button {
    color: #333;
    background: #FF99AC;
    border-color: #FF99AC;
}
.module-button.el-button--default.is-round.second-button {
    color: #333;
    background: #fff;
    border-color: #e5e5e5;
}
.module-button.el-button.is-round.first-button:focus, .el-button.is-round.first-button:hover {
    background: #d31e37;
    border-color: #d31e37;
    color: #fff;
}
.module-button.el-button.is-round.second-button:focus, .el-button.is-round.second-button:hover {
    background: #fff4f6;
    border-color: #FF99AC;
    color: #FF99AC;
}

.profile-part {
    width: 860px;
}

.profile-front {
    padding-top: 50px;
    padding-left: 40px;
    font-size: 24px;
    text-align: left;
}

.profile-front > div {
    display: inline-block;
}

.profile-front .right-title {
    font-size: 18px;
    float: right;
    padding-right: 40px;
}

.profile-main {
    padding: 30px 40px 0 40px;
}
</style>