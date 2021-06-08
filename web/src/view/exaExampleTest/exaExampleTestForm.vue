<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="管理员ID:"><el-input v-model.number="formData.adminId" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="分类ID(单选):"><el-input v-model.number="formData.categoryId" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="分类ID(多选):">
                <el-input v-model="formData.categoryIds" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="星期(单选):monday=星期一,tuesday=星期二,wednesday=星期三:">
                <el-input v-model="formData.week" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="标志(多选):hot=热门,index=首页,recommend=推荐:">
                <el-input v-model="formData.flag" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="性别(单选):male=男,female=女:">
                <el-input v-model="formData.genderdata" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="爱好(多选):music=音乐,reading=读书,swimming=游泳:">
                <el-input v-model="formData.hobbydata" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="标题:">
                <el-input v-model="formData.title" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="内容:">
                <el-input v-model="formData.content" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="图片:">
                <el-input v-model="formData.image" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="图片组:">
                <el-input v-model="formData.images" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="附件:">
                <el-input v-model="formData.attachfile" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="关键字:">
                <el-input v-model="formData.keywords" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="描述:">
                <el-input v-model="formData.description" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="省市:">
                <el-input v-model="formData.city" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="配置:key=名称,value=值:">
                <el-input v-model="formData.json" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="价格:">
                  <el-input-number v-model="formData.price" :precision="2" clearable></el-input-number>
           </el-form-item>
           
             <el-form-item label="点击:"><el-input v-model.number="formData.views" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="开始日期:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.startdate" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="活动时间(datetime):">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.activitytime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="年:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.year" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.times" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="刷新时间(int):"><el-input v-model.number="formData.refreshtime" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="权重:"><el-input v-model.number="formData.weigh" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="开关:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.switch" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="状态:">
                <el-input v-model="formData.status" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="状态值:0=禁用,1=正常,2=推荐:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.state" clearable ></el-switch>
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
    createExaExampleTest,
    updateExaExampleTest,
    findExaExampleTest
} from "@/api/exaExampleTest";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "ExaExampleTest",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            adminId:0,
            categoryId:0,
            categoryIds:"",
            week:"",
            flag:"",
            genderdata:"",
            hobbydata:"",
            title:"",
            content:"",
            image:"",
            images:"",
            attachfile:"",
            keywords:"",
            description:"",
            city:"",
            json:"",
            price:0,
            views:0,
            startdate:new Date(),
            activitytime:new Date(),
            year:new Date(),
            times:new Date(),
            refreshtime:0,
            weigh:0,
            switch:false,
            status:"",
            state:false,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createExaExampleTest(this.formData);
          break;
        case "update":
          res = await updateExaExampleTest(this.formData);
          break;
        default:
          res = await createExaExampleTest(this.formData);
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
    const res = await findExaExampleTest({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reexaExampleTest
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