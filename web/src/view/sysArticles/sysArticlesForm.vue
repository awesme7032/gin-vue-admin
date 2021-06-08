<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="作者:">
                <el-input v-model="formData.username" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="文章分类外键:"><el-input v-model.number="formData.categoryId" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="评论次数
评论次数:"><el-input v-model.number="formData.commentCount" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="文章内容:">
                <el-input v-model="formData.content" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="状态,正常为1:"><el-input v-model.number="formData.status" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="文章摘要:">
                <el-input v-model="formData.summary" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="阅读次数:"><el-input v-model.number="formData.viewCount" clearable placeholder="请输入"></el-input>
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
    createSysArticles,
    updateSysArticles,
    findSysArticles
} from "@/api/sysArticles";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "SysArticles",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            username:"",
            categoryId:0,
            commentCount:0,
            content:"",
            status:0,
            summary:"",
            viewCount:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createSysArticles(this.formData);
          break;
        case "update":
          res = await updateSysArticles(this.formData);
          break;
        default:
          res = await createSysArticles(this.formData);
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
    const res = await findSysArticles({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.resysArticles
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