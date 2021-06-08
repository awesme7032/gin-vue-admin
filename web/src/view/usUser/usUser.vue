<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="">
          <el-input placeholder="用户名" v-model="username" clearable></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增用户表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
              <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
            </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
        :data="tableData"
        @selection-change="handleSelectionChange"
        border
        ref="multipleTable"
        stripe
        style="width: 100%"
        tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55"></el-table-column>

      <el-table-column label="用户名" prop="username" width="120"></el-table-column>

      <!--      <el-table-column label="邮箱" prop="email" width="120"></el-table-column>-->

      <el-table-column label="性别" prop="gender" width="120">
        <template slot-scope="scope">{{ scope.row.gender|formatBoolean }}</template>
      </el-table-column>

      <el-table-column label="头像" prop="icon" width="120"></el-table-column>

      <!--    <el-table-column label="邀请码ID" prop="invitationId" width="120"></el-table-column>-->

      <el-table-column label="昵称" prop="nickname" width="120"></el-table-column>

      <el-table-column label="手机号" prop="phone" width="120"></el-table-column>

      <el-table-column label="状态" prop="state" width="120">
        <template slot-scope="scope">{{ scope.row.state|formatBoolean }}</template>
      </el-table-column>
      <el-table-column label="加入IP" prop="joinIp" width="120"></el-table-column>

      <el-table-column label="上次登录IP" prop="lastLoginIp" width="120"></el-table-column>

      <el-table-column label="上次登录时间" prop="lastLoginTime" width="120"></el-table-column>

      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>

      <!--    <el-table-column label="type字段" prop="type" width="120">-->
      <!--         <template slot-scope="scope">{{scope.row.type|formatBoolean}}</template>-->
      <!--    </el-table-column>-->


      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateUsUser(scope.row)" size="small" type="primary"
                     icon="el-icon-edit">变更
          </el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :style="{float:'right',padding:'20px'}"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
        layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">

        <el-form-item label="用户名:">
          <el-input v-model="formData.username" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="密码:">
          <el-input v-model="formData.password" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="邮箱:">
          <el-input v-model="formData.email" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="性别:">
          <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="男" inactive-text="女"
                     v-model="formData.gender" clearable></el-switch>
        </el-form-item>

        <el-form-item label="头像:">
          <el-input v-model="formData.icon" clearable placeholder="请输入"></el-input>
        </el-form-item>


        <el-form-item label="加入IP:">
          <el-input v-model="formData.joinIp" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="上次登录IP:">
          <el-input v-model="formData.lastLoginIp" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="上次登录时间:">
          <el-date-picker type="date" placeholder="选择日期" v-model="formData.lastLoginTime" clearable></el-date-picker>
        </el-form-item>

        <el-form-item label="昵称:">
          <el-input v-model="formData.nickname" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="手机号:">
          <el-input v-model="formData.phone" clearable placeholder="请输入"></el-input>
        </el-form-item>


        <el-form-item label="真实名称:">
          <el-input v-model="formData.realname" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="state字段:">
          <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否"
                     v-model="formData.state" clearable></el-switch>
        </el-form-item>


      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {createUsUser, deleteUsUser, deleteUsUserByIds, findUsUser, getUsUserList, updateUsUser} from "@/api/usUser"; //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";

export default {
  name: "UsUser",
  mixins: [infoList],
  data() {
    return {
      listApi: getUsUserList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [], formData: {
        email: "",
        gender: false,
        icon: "",
        idCard: "",
        invitationId: "",
        joinIp: "",
        lastLoginIp: "",
        lastLoginTime: new Date(),
        levelId: false,
        nickname: "",
        password: "",
        phone: "",
        pid: 0,
        realname: "",
        state: false,
        type: false,
        username: "",

      },
      username:"", //跳转后无法输入，新增定义一个变量
    };
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? "男" : "女";
      } else {
        return "";
      }
    }
  },
  methods: {
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10

      if (this.searchInfo.gender == "") {
        this.searchInfo.gender = null
      }
      if (this.searchInfo.levelId == "") {
        this.searchInfo.levelId = null
      }
      if (this.searchInfo.state == "") {
        this.searchInfo.state = null
      }
      if (this.searchInfo.type == "") {
        this.searchInfo.type = null
      }
      this.getTableData()
      this.searchInfo.username = "";

    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteUsUser(row);
      });
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length == 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
      this.multipleSelection.map(item => {
        ids.push(item.ID)
      })
      const res = await deleteUsUserByIds({ids})
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length == ids.length) {
          this.page--;
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateUsUser(row) {
      const res = await findUsUser({ID: row.ID});
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reusUser;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        email: "",
        gender: false,
        icon: "",
        idCard: "",
        invitationId: "",
        joinIp: "",
        lastLoginIp: "",
        lastLoginTime: new Date(),
        levelId: false,
        nickname: "",
        password: "",
        phone: "",
        pid: 0,
        realname: "",
        state: false,
        type: false,
        username: "",

      };
    },
    async deleteUsUser(row) {
      const res = await deleteUsUser({ID: row.ID});
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
          this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
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
          type: "success",
          message: "创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  },
  //监听属性变化
  watch:{
    username(val,from){
      this.searchInfo.username = val;
    }
  },
  mounted() {
    if(this.$route.query.username){
      this.searchInfo.username = this.$route.query.username
    }
    this.onSubmit()
  }
};
</script>

<style>
</style>
