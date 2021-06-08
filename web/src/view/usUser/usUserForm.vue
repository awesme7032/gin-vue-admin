<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="email字段:">
                <el-input v-model="formData.email" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="gender字段:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.gender" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="icon字段:">
                <el-input v-model="formData.icon" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="idCard字段:">
                <el-input v-model="formData.idCard" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="invitationId字段:">
                <el-input v-model="formData.invitationId" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="joinIp字段:">
                <el-input v-model="formData.joinIp" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="lastLoginIp字段:">
                <el-input v-model="formData.lastLoginIp" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="lastLoginTime字段:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.lastLoginTime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="levelId字段:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.levelId" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="nickname字段:">
                <el-input v-model="formData.nickname" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="password字段:">
                <el-input v-model="formData.password" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="phone字段:">
                <el-input v-model="formData.phone" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="pid字段:"><el-input v-model.number="formData.pid" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="realname字段:">
                <el-input v-model="formData.realname" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="state字段:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.state" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="type字段:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.type" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="username字段:">
                <el-input v-model="formData.username" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createUsUser,
    updateUsUser,
    findUsUser
} from "@/api/usUser";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "UsUser",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            email:"",
            gender:false,
            icon:"",
            idCard:"",
            invitationId:"",
            joinIp:"",
            lastLoginIp:"",
            lastLoginTime:new Date(),
            levelId:false,
            nickname:"",
            password:"",
            phone:"",
            pid:0,
            realname:"",
            state:false,
            type:false,
            username:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createUsUser(this.formData);
          break;
        case "update":
          res = await updateUsUser(this.formData);
          break;
        default:
          res = await createUsUser(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findUsUser({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reusUser
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
}
};
</script>

<style>
</style>