<template>
  <div id="mainContainer">
    <div id="mainContainerInner">
      <div id="switchBtnContainer">
				<button class="switchBtn" @click="switchPage(1);" title="登录您的账号" id="user_loginBtn">登&emsp;&emsp;录</button>
				<button class="switchBtn" @click="switchPage(2);" title="注册新的账号" id="user_registerBtn">注&emsp;&emsp;册</button>
			</div>
      <div style="width: 94%;margin-left: 3%;margin-top: 6px;height: 3px;border-radius: 5%;background-color: #d01d40;"></div>
      <div id="mainContainer_switch">
        <div :class="sliderType">
          <!--账号登录-->
          <div id="user_subPage_login">
            <el-form class="loginForm" :rules="loginFormRule" ref="loginForm" :model="loginForm" label-width="50px">
              <el-form-item label="账户" prop="account">
                <el-input v-model="loginForm.account" placeholder="请输入用户名"></el-input>
              </el-form-item>
              <el-form-item label="密码" prop="password">
                <el-input type="password" v-model="loginForm.password" autocomplete="off" @keyup.enter.native="loginCheck()" placeholder="请输入密码"></el-input>
              </el-form-item>
              <el-form-item>
                <div style="width: 100%;display: flex;flex-direction: row;align-items: center;justify-content: center">
                  <div style="width: 50%;height: 100%">
                    <input ref="rememAC" type="checkbox" id="rememAC" name="rememAC"/>
                    <label for="rememAC" style="cursor: pointer;margin-right:9rem;color: #d01d40;font-size: 14px;line-height: 14px;">记住我</label>
                  </div>
                  <div style="width: 45%;height: 100%;display: flex;justify-content: end;align-items: center;margin-right: 5%;">
                    <a class="fontLink" style="margin-left: 10%" @click="switchPage(2)">没有账号？立即注册</a>
                  </div>
                </div>
              </el-form-item>
              <input type="button" id="loginBtn" @click="loginCheck()" value="登      录"/>
            </el-form>
          </div>

          <!--账号注册-->
          <div id="user_subPage_register">
            <el-form class="registerForm" :rules="registerFormRule" ref="registerForm" :model="registerForm" label-width="18%">
              <el-form-item label="账户" prop="account">
                <el-input v-model="registerForm.account" placeholder="您的用户名"></el-input>
              </el-form-item>
              <el-form-item label="密码" prop="pass">
                <el-input type="password" v-model="registerForm.password" autocomplete="off" placeholder="您的密码"></el-input>
              </el-form-item>
              <el-form-item label="确认密码" prop="checkPass">
                <el-input type="password" v-model="registerForm.checkPass" autocomplete="off" placeholder="再次确认您的密码"></el-input>
              </el-form-item>
              <el-form-item label="手机号" prop="telephone">
                <el-input v-model="registerForm.telephone" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="注册城市" prop="location">
                <!-- <el-input v-model="registerForm.location" autocomplete="off"></el-input> -->
                <el-cascader
                  :options="areaSelectData"
                  @change="handleChange"
                  class="full-width"
                  size="large"
                  placeholder="请选择您所在的城市（默认北京市海淀区）"
                />
              </el-form-item>
              <el-form-item label="姓名" prop="name">
                <el-input v-model="registerForm.name" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="证件类型" prop="cred_type">
                <el-input v-model="registerForm.cred_type" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="证件号码" prop="cred_id">
                <el-input v-model="registerForm.cred_id" autocomplete="off"></el-input>
              </el-form-item>
              <input type="button" id="registerBtn" @click="registerCheck()" value="注      册"/>
              <p style="color: gray;font-size: small;text-align: center;font-family: 幼圆,serif;">点击即代表同意 <a @click="showInfo()" class="fontLink">好味道交易中介平台 用户协议</a></p>
            </el-form>
          </div>
        </div>
      </div>
    </div>
    <!-- slogan -->
    <div id="slogan">
        <p>-&emsp;登录好味道 解锁完整功能&emsp;-</p>
    </div>
  </div>

</template>

<script>
import { regionData, CodeToText, TextToCode } from "element-china-area-data";
import jwtDecode from 'jwt-decode'

export default {
  name: "P_AccountPage",
  props: ['w'],
  data() {
    const validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        if (this.registerForm.checkPass !== '') {
          this.$refs.registerForm.validateField('checkPass');
        }
        callback();
      }
    };
    const validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.registerForm.password) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };
    return {
      areaSelectData: regionData,
      sliderType: "mainContainer_switch_slider_left",
      loginForm: {
        account: "",
        password: "",
      },
      loginFormRule: {
        account: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ]
      },
      registerForm:{
        account: "",
        password:"",
        checkPass: "",
        type: 1,
        level: 1,
        name: "",
        cred_type: "",
        cred_id: "",
        telephone: "",
        location: {
          country: "北京市",
          province: "市辖区",
          city: "海淀区"
        },
        intro: "该用户很懒，还没有个人简介哦~"
      },
      registerFormRule: {
        account: [
          { required: true, message: '请输入用户名', trigger: 'change' }
        ],
        password: [
          { required: true,  validator: validatePass, trigger: 'blur' }
        ],
        checkPass: [
          { required: true,  validator: validatePass2, trigger: 'blur' }
        ],
        name: [
          { required: true, message: '请输入姓名', trigger: 'change' }
        ],
        cred_type: [
          { required: true, message: '请输入证件类型', trigger: 'change' }
        ],
        cred_id: [
          { required: true, message: '请输入证件号码', trigger: 'change' }
        ],
        telephone: [
          { required: true, message: '请输入联系电话', trigger: 'change' }
        ],
      },
    }
  },
  methods:{
    showInfo(){
      alert("敬请期待！")
    },
    switchPage(w){
      if (w===1){
        this.sliderType = "mainContainer_switch_slider_left";
      }
      else if(w===2){
        this.sliderType = "mainContainer_switch_slider_right";
      }
    },
    // getCookie(objname){//获取指定名称的cookie的值
    //   const arrstr = document.cookie.split("; ");
    //   for(let i = 0; i < arrstr.length; i ++){
    //     const temp = arrstr[i].split("=");
    //     if(temp[0] === objname) return temp[1];
    //   }
    //   return null;
    // },
    // setCookie(name, value, hours, path) {
    //   const expires = new Date();
    //   expires.setTime(expires.getTime() + hours * 3600000);
    //   path = path === "" ? "" : ";path=" + path;
    //   const _expires = (typeof hours) == "string" ? "" : ";expires=" + expires.toUTCString();
    //   document.cookie = name + "=" + value + _expires + path;
    // },
    loginCheck(){
      if (this.loginForm.account === "" || this.loginForm.password===""){
        this.$message({
          showClose: true,
          message: '请首先输入账号和密码',
          type: 'error'
        });
        return;
      }
      // console.log('about to get!')
      // this.$axios.get('/test/ping').then(function(response){
      //   alert(response.data);
      // }).catch(function (err) {
      //   console.log(err);
      // })
      console.log('about to post!')
      let config = {
          headers: {'Content-Type': 'application/json'}
      }; 
      this.$axios.post('/v1/login', {
        "username": this.loginForm.account,
        "password": this.loginForm.password
      }, config).then((response) => {
        // console.log(response.data)
        if (response.status === 401){
          this.$message({
            showClose: true,
            message: '该用户不存在或密码错误！',
            type: 'error'
          });
        }
        else{
          // 登录成功
          console.log(response.data.token)
          localStorage.setItem("sk_token", response.data.token)
          // this.$cookies.set("sk_token", response.data.token, {expires:"1d"})
          this.$cookies.set("sk_uname", this.loginForm.account, {expires:"7d"})
          if (this.$refs.rememAC.checked){
            //记录pwd
            this.$cookies.set("sk_pwd", this.loginForm.password, {expires:"7d"})
          }
          // this.setCookie("NBI_token",response.data.token, 2, "/NBI");
          // this.setCookie("NBI_UID", response.data.uid, 72, "/NBI");
          // this.$bus.$emit('changeStatus', {status: true, uname: response.data.uname});
          // if (this.$refs.rememAC.checked){
          //   //记录pwd
          //   this.setCookie("NBI_pwd", this.loginForm.password, 72, "/NBI");
          // }
          const code = jwtDecode(response.data.token)
          console.log('这是token信息')
          console.log(code)//解析成功
          this.$cookies.set("sk_uid", code.ID, {expires:"7d"})
          this.$bus.$emit("logStatus", true)
          this.$message({
            showClose: true,
            message: '登录成功！页面将会自动跳转',
            type: 'success'
          });
          setTimeout(()=>{
            this.$router.push({path: "/user"});
          }, 2000);

          console.log("success")
        }
      }).catch(err => {
          console.log(err)
          if (err.message === 'Request failed with status code 401') {
              this.$message({
              showClose: true,
              message: '该用户不存在或密码错误！',
              type: 'error'
            });
          }      
      });
    },
    registerCheck(){
      if (this.registerForm.checkPass !== this.registerForm.password){
        this.$message({
          showClose: true,
          message: '两次密码输入不一致',
          type: 'error'
        });
        return;
      }
      if (this.registerForm.account === null || this.registerForm.password === null){
        this.$message({
          showClose: true,
          message: '注册关键信息不能为空',
          type: 'error'
        });
        return;
      }
      let config = {
          headers: {'Content-Type': 'application/json'}
      }; 
      this.$axios.post('/v1/sign', {
        "username": this.registerForm.account,
        "password": this.registerForm.password,
        "type": this.registerForm.type,
        "level": this.registerForm.level,
        "name": this.registerForm.name,
        "cred_type": this.registerForm.cred_type,
        "cred_id": this.registerForm.cred_id,
        "telephone": this.registerForm.telephone,
        "intro": this.registerForm.intro,
        "location": this.registerForm.location
      }, config).then((response) => {
        if (response.data === 0){
          this.$message({
            showClose: true,
            message: '注册失败，该用户名已被注册！',
            type: 'error'
          });
          this.registerForm.account = "";
          this.registerForm.password = "";
        }
        else {
          this.$message({
            showClose: true,
            message: '注册成功！请登录。',
            type: 'success'
          });
          setTimeout(()=>{
            this.switchPage(1);
          }, 3000);
        }
      });
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
    handleChange(value) {
      //value为省市区code数组
      if (value) {
        var provinceCode = CodeToText[value[0]];//code转为省
        var cityCode = CodeToText[value[1]];//市
        var orgion = CodeToText[value[2]];//区
        this.registerForm.location.country = provinceCode;
        this.registerForm.location.province = cityCode;
        this.registerForm.location.city = orgion;
        // this.form.selectedOptions = provinceCode + cityCode + orgion;
      }
    },
  },
  mounted(){
    const pwd = this.$cookies.get("sk_pwd")
    console.log(pwd)
    if (pwd != null){
      this.loginForm.password = pwd;
      this.loginForm.account = this.$cookies.get("sk_uname");
      this.$refs.rememAC.checked = true;
    }
  },
  activated() {
    this.switchPage(this.w);
    // console.log(this.w)
  },
}
</script>

<style scoped>
#mainContainer{
    width: 100%;
    height: 710px;
    padding-top: 60px;
    overflow: hidden;
    /* display: flex; */
    align-items: center;
    /* justify-content: center; */
    background-color: #fbced5;
    background-image: url('@/assets/passport-background.jpg');
    background-repeat: no-repeat;
    background-size: 1280px 800px;
    background-position: 0 -120px;
}
#mainContainerInner{
    width: 446px;
    background-color: rgba(255, 255, 255, 0.99);
    transition: 0.2s ease;
    overflow: hidden;
    margin: auto;
}
#switchBtnContainer{
	width: 100%;
	padding-top: 20px;
	display: flex;
	justify-content: center;
}
.switchBtn{
	width: 204px;
	height: 40px;
	margin: 5px;
	border: #f4f4f4 solid 1px;
	background: #FFF;
    color: #666;
	cursor: pointer;
	transition: all 0.2s ease;
}
.switchBtn:hover, .switchBtn:focus {
	background: #fbced5;
    color: #fff;
    font-weight: bold;
}
#mainContainer_switch{
    overflow: hidden;
}
.mainContainer_switch_slider_left{
    margin-top: 3%;
    margin-left: 0;
    width: 200%;
    height: 285px;
    display: flex;
    flex-direction: row;
    transition: all 0.2s ease;
}
.mainContainer_switch_slider_right{
    margin-top: 3%;
    margin-left: -100%;
    width: 200%;
    height: 610px;
    display: flex;
    flex-direction: row;
    transition: all 0.2s ease;
}
#user_subPage_login{
  width: 100%;
  /*height: 100%;*/
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  transition: all 0.2s ease;
}
#user_subPage_register{
  width: 100%;
  /*height: 100%;*/
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  transition: all 0.2s ease;
}
.loginForm{
	width: 90%;
	white-space: nowrap;
}
.registerForm{
  width: 90%;
	white-space: nowrap;
}
input[type="checkbox"] {
    height: 16px;
    width: 16px;
    visibility: hidden;
}
input[type=checkbox]:after {
  position: absolute;
  width: 11px;
  height: 16px;
  top: 12px;
  content: " ";
  background-color: #d01d40;
  color: #fff;
  display: inline-block;
  visibility: visible;
  padding: 0px 3px;
  border-radius: 3px;
}
input[type=checkbox]:checked:after {
  content: "✓";
  font-size: 12px;
  font-weight: bold;
}
#rememAC{
  margin-left: -30px;
  cursor: pointer;
  margin-top: 8px;
}
#loginBtn{
	width: 100%;
	height: 45px;
	border: #f4f4f4 solid 1px;
	background: #d01d40;
    color: #fff;
    font-size: 18px;
	cursor: pointer;
	transition: all 0.2s ease;
}
#loginBtn:hover{
    color: rgb(221, 221, 221);
    background-color: #a81935;
}
input[type="checkbox"] {
    height: 18px;
    width: 18px;
    margin-top: 18px;
    margin-left: 10px;
}
#registerBtn{
	width: 100%;
	height: 45px;
    margin-top: 2px;
	border: #f4f4f4 solid 1px;
	background: #d01d40;
    color: #fff;
    font-size: 18px;
	cursor: pointer;
	transition: all 0.2s ease;
}
#registerBtn:hover{
	color: rgb(221, 221, 221);
    background-color: #a81935;
}
#slogan{
	width: 100%;
	margin-top: 20px;
	display: flex;
	justify-content: center;
	align-items: center;
	font-family: Arial, Helvetica, sans-serif;
	color: #fff;
}
.fontLink{
  margin-left: 4px;
  margin-right: 4px;
  color: grey;
  cursor: pointer;
  font-size: small;
  font-family: 幼圆,serif;
}
.fontLink:hover{
  color: #d01d40;
}

.el-input >>> .el-input__inner:focus {
    border-color: #fbced5;
}

.el-cascader >>> .el-input {
  width: 330px;
}

.el-cascader >>> .el-input__inner:focus {
  border-color: #fbced5;
}
</style>