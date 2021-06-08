<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">

        <el-form-item label="用户名">
          <el-input placeholder="用户名" v-model="searchInfo.username" clearable></el-input>
        </el-form-item>
        <!--        关键字模糊搜索-->
        <el-form-item label="关键字">
          <el-input placeholder="关键字" v-model="searchInfo.keywords" clearable></el-input>
        </el-form-item>
        <!--       星期下拉框-->
        <el-form-item label="星期">
          <el-select placeholder="请选择" v-model="searchInfo.week">
            <el-option
                :key="item.value"
                :label="item.label"
                :value="item.value"
                v-for="item in weekOptions">
            </el-option>
          </el-select>
        </el-form-item>
        <!--  时间查询  -->
        <el-form-item label="时间">
          <el-date-picker placeholder="选择日期" v-model="searchInfo.CreatedAt" clearable></el-date-picker>
        </el-form-item>
        <!--        时间区间查询-->
        <el-form-item label="时间区间">
<!--          区间限制:picker-options="pickerOptions"-->
          <el-date-picker
              type="datetimerange"
              v-model="searchInfo.startdate"
              range-separator="至"
              unlink-panels
              :clearable="false"
              start-placeholder="开始日期"
              end-placeholder="结束日期">
          </el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增</el-button>
        </el-form-item>
        <!--          顶部新增功能按钮区域-->
        <el-form-item>
          <el-button class="table-button" @click="updateTopExaExampleTest" size="small" type="primary"
                     icon="el-icon-edit">变更
          </el-button>
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
        @row-click="onSelect"
        @selection-change="handleSelectionChange"
        border
        ref="multipleTable"
        stripe
        style="width: 100%"
        tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55"></el-table-column>
      <!--      单击链接跳转-->
      <el-table-column label="管理员ID" prop="adminId" width="120">
      </el-table-column>

      <el-table-column label="用户名称:" prop="username">
        <template slot-scope="scope">
          <router-link :to="{name: 'usUser',query:{username: scope.row.username}}">
            <span style="color:#06d;">{{ scope.row.username }}</span>
          </router-link>
          <!--          原始写法-->
          <!--          <span style="color:#06d;" @click="usUserInfo(scope.row.username)">{{ scope.row.username }}</span>-->
        </template>
      </el-table-column>

      <el-table-column label="分类ID(单选)" prop="categoryId" width="120"></el-table-column>

      <el-table-column label="分类ID(多选)" prop="categoryIds" width="120"></el-table-column>

      <el-table-column label="星期(单选):" prop="week" width="120">
        <template slot-scope="scope">
          <div>
            {{ scope.row.week }}
            <el-tag
                :key="scope.row.week"
                :type="scope.row.week|weekTagTypeFilte"
                effect="dark"
                size="mini"
            >
              {{ scope.row.week|weekFilte }}
            </el-tag>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="标志(多选):hot=热门,index=首页,recommend=推荐" prop="flag" width="120"></el-table-column>

      <el-table-column label="性别(单选):" prop="genderdata" width="120">
        <template slot-scope="scope">
          {{ scope.row.genderdata|genderFormat}}
        </template>
      </el-table-column>

      <el-table-column label="爱好(多选):music=音乐,reading=读书,swimming=游泳" prop="hobbydata" width="120"></el-table-column>

      <el-table-column label="标题" prop="title" width="120"></el-table-column>

      <el-table-column label="内容" prop="content" width="120"></el-table-column>

      <el-table-column label="图片" prop="image" width="120">
<!--        <template slot-scope="scope">-->
          <img src="@/assets/github.png" style="width:100px;height: 100px">
          <!--          <CustomPic picType="file" :picSrc="scope.row.image" />-->
<!--          <img :src="scope.row.image" style="width: 65px">-->
<!--        </template>-->
      </el-table-column>

      <el-table-column label="图片组" prop="images" width="120">
      </el-table-column>

      <el-table-column label="附件" prop="attachfile" width="120"></el-table-column>

      <el-table-column label="关键字" prop="keywords" width="120"></el-table-column>

      <el-table-column label="描述" prop="description" width="120"></el-table-column>

      <el-table-column label="省市" prop="city" width="120"></el-table-column>

      <el-table-column label="配置:key=名称,value=值" prop="json" width="120"></el-table-column>

      <el-table-column label="价格" prop="price" width="120"></el-table-column>

      <el-table-column label="点击" prop="views" width="120"></el-table-column>

      <el-table-column label="开始日期" width="120">
        <template slot-scope="scope">{{ scope.row.startdate|formatDate }}</template>
      </el-table-column>

      <el-table-column label="活动时间(datetime)" prop="activitytime" width="120">
        <template slot-scope="scope">{{ scope.row.activitytime|formatDate }}</template>
      </el-table-column>

      <el-table-column label="创建日期" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDateYear }}</template>
      </el-table-column>
      <el-table-column label="年" prop="year" width="120"></el-table-column>

      <el-table-column label="时间" prop="times" width="120">
        <template slot-scope="scope">{{ scope.row.times|formatDate }}</template>
      </el-table-column>

      <el-table-column label="刷新时间(int)" prop="refreshtime" width="120">
        <template slot-scope="scope">{{ scope.row.refreshtime|formatDate }}</template>
      </el-table-column>

      <el-table-column label="权重" prop="weigh" width="120"></el-table-column>

      <el-table-column label="开关" prop="switch" width="120">
        <template slot-scope="scope">{{ scope.row.switch|formatBoolean }}</template>
      </el-table-column>

      <el-table-column label="状态" prop="status" width="120">
        <template slot-scope="scope">
          <el-switch active-color="#13ce66" inactive-color="#ff4949" :active-value="1" :inactive-value="0"
                     v-model="scope.row.status" @change="statusSwitch(scope.row)"></el-switch>
        </template>
      </el-table-column>

      <!--          获取对应的数组 -->
      <el-table-column label="状态值:0=禁用,1=正常,2=推荐" prop="state" width="120">
        <template slot-scope="scope">
          <el-tag type="danger" v-if="scope.row.state===0">禁用</el-tag>
          <el-tag type="primary" v-if="scope.row.state===1">正常</el-tag>
          <el-tag type="success" v-if="scope.row.state===2">推荐</el-tag>
        </template>
      </el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateExaExampleTest(scope.row)" size="small" type="primary"
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
        <el-form-item label="管理员ID:">
          <el-input v-model.number="formData.adminId" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="用户名称:">
          <el-input v-model.number="formData.username" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="分类ID(单选):">
          <el-input v-model.number="formData.categoryId" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="分类ID(多选):">
          <el-input v-model="formData.categoryIds" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="星期(单选):monday=星期一,tuesday=星期二,wednesday=星期三:">
          <el-select placeholder="请选择" v-model="formData.week" style="width: 100%">
            <el-option
                :key="item.value"
                :label="item.label"
                :value="item.value"
                v-for="item in weekOptions">
            </el-option>
          </el-select>
          <!--          <el-input v-model="formData.week" clearable placeholder="请输入"></el-input>-->
        </el-form-item>

        <el-form-item label="标志(多选):hot=热门,index=首页,recommend=推荐:">
          <el-input v-model="formData.flag" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="性别(单选):male=男,female=女:">
          <el-radio-group v-model="formData.genderdata">
            <el-radio label="male">男</el-radio>
            <el-radio label="famale">女</el-radio>
          </el-radio-group>
<!--          <el-vadio v-model="formData.genderdata"></el-vadio>-->
        </el-form-item>

        <el-form-item label="爱好(多选):music=音乐,reading=读书,swimming=游泳:">
          <el-input v-model="formData.hobbydata" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="标题:">
          <el-input v-model="formData.title" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="内容:">
          <quill-editor :options="editorOption" v-model="formData.content"></quill-editor>
        </el-form-item>

        <el-form-item label="图片:">
          <el-input v-model="formData.image" clearable placeholder="请输入"></el-input>
          <el-upload :action="uploadUrl" :on-success="uploadSuccess">
            <el-button>上传</el-button>
          </el-upload>
        </el-form-item>

        <el-form-item label="图片组:">
          <el-input v-model="formData.images" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="附件:">
          <el-input v-model="formData.attachfile" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="关键字:">
          <el-input v-model="formData.keywords" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="描述:">
          <el-input v-model="formData.description" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="省市:">
          <el-input v-model="formData.city" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="配置:key=名称,value=值:">
          <el-input v-model="formData.json" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="价格:">
          <el-input-number v-model="formData.price" :precision="2" clearable></el-input-number>
        </el-form-item>

        <el-form-item label="点击:">
          <el-input v-model.number="formData.views" clearable placeholder="请输入"></el-input>
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

        <el-form-item label="刷新时间(int):">
          <el-input v-model.number="formData.refreshtime" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="权重:">
          <el-input v-model.number="formData.weigh" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="开关:">
          <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否"
                     v-model="formData.switch" clearable></el-switch>
        </el-form-item>

        <el-form-item label="状态:">
          <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="开启" inactive-text="关闭"
                     :active-value="1" :inactive-value="0"
                     v-model="formData.status" clearable></el-switch>
        </el-form-item>

        <el-form-item label="状态值:0=禁用,1=正常,2=推荐:">
          <el-input-number v-model="formData.state" clearable placeholder="请输入"></el-input-number>
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
import {
  createExaExampleTest,
  deleteExaExampleTest,
  deleteExaExampleTestByIds,
  findExaExampleTest,
  getExaExampleTestList,
  updateExaExampleTest
} from "@/api/exaExampleTest"; //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";
// 格式化的需要再这里引入
import {formatEnumList} from "@/utils/common";
import infoList from "@/mixins/infoList";
// 富文本编辑器
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
import {quillEditor} from 'vue-quill-editor'
import {quillRedefine} from 'vue-quill-editor-upload'
// import CustomPic from "@/components/customPic";
//  星期下拉框固定数据源
const weekOptions = [
  {
    value: "Monday",
    label: "星期一",
    type: ""
  },
  {
    value: "Tuesday",
    label: "星期二",
    type: ""
  }, {
    value: "Wednesday",
    label: "星期三",
    type: ""
  }, {
    value: "Thursday",
    label: "星期四",
    type: ""
  }, {
    value: "Friday",
    label: "星期五",
    type: ""
  }, {
    value: "Saturday",
    label: "星期六",
    type: "danger"
  }, {
    value: "Sunday",
    label: "星期日",
    type: "danger"
  }
];

export default {

  components: {
    // eslint-disable-next-line vue/no-unused-components
    quillEditor, quillRedefine,
    // CustomPic,
  },
  name: "ExaExampleTest",
  mixins: [infoList],
  data() {
    return {
      listApi: getExaExampleTestList,
      editorOption: {},
      uploadUrl: this.$store.state.uploadUrl,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [], formData: {
        adminId: 0,
        categoryId: 0,
        categoryIds: "",
        week: "",
        username: "",
        flag: "",
        genderdata: "",
        hobbydata: "",
        title: "",
        content: "",
        image: "",
        images: "",
        attachfile: "",
        keywords: "",
        description: "",
        city: "",
        json: "",
        price: 0,
        views: 0,
        startdate: new Date(),
        activitytime: new Date(),
        year: new Date(),
        times: new Date(),
        refreshtime: 0,
        weigh: 0,
        switch: true,
        status: 1,
        state: 1,
      },
      weekOptions: weekOptions,
      // 时间区间限制
      pickerOptions: {
        disabledDate(time) {
          //小于当月大于当天
          return time.getTime() < Date.now() - 8.64e7 || time.getMonth() != new Date().getMonth()
        }
      },
    };
  },
  filters: {
    genderFormat:function (gender){
      if (gender == "male") {
        return "男"
      }else {
        return "女"
      }
    },
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatDateYear: function (time) {
      if (time != null && time != "") {
        let data = new Date(time)
        return formatTimeToStr(data, "yyyy")
      } else {
        return "";
      }
    },
    // 枚举（enum）类型格式化为下拉框或者单选按钮
    formatEnumForDropDownList: function (val) {
      if (val != null && val != "") {
        return formatEnumList(val, "")
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    },
    weekFilte(value) {
      const target = weekOptions.filter(item => item.value === value)[0];
      return target.label
      // return target && `${target.label}`;
    },
    weekTagTypeFilte(value) {
      const target = weekOptions.filter(item => item.value === value)[0];
      return target && `${target.type}`;
    },
  },
  methods: {
    uploadSuccess(res) {
      if (res.code === 0) {
        this.formData.image = res.data
        this.$message.success("上传成功")
      } else {
        this.$message.error(res.msg)
      }
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.switch == "") {
        this.searchInfo.switch = null
      }
      this.getTableData()
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
        this.deleteExaExampleTest(row);
      });
    },
    // 顶部选中行修改
    async updateTopExaExampleTest() {
      if (this.multipleSelection.length == 0) {
        this.$message({
          type: 'warning',
          message: '请选择要修改的数据'
        })
        return
      }
      // 通过引用获取选中的行
      const selectData = this.$refs.multipleTable.selection
      if (selectData.length > 1) {
        this.$message.error("只允许修改一行")
        return false
      }
      const res = await findExaExampleTest({ID: selectData[0].ID});
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reexaExampleTest;
        this.dialogFormVisible = true;
      }
    },
    // 单击行选中复选框
    onSelect(val) {
      this.$refs.multipleTable.toggleRowSelection(val)
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
      const res = await deleteExaExampleTestByIds({ids})
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
    async updateExaExampleTest(row) {
      const res = await findExaExampleTest({ID: row.ID});
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reexaExampleTest;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        adminId: 0,
        categoryId: 0,
        categoryIds: "",
        week: "",
        flag: "",
        genderdata: "",
        hobbydata: "",
        title: "",
        content: "",
        image: "",
        images: "",
        attachfile: "",
        keywords: "",
        description: "",
        city: "",
        json: "",
        price: 0,
        views: 0,
        startdate: new Date(),
        activitytime: new Date(),
        year: new Date(),
        times: new Date(),
        refreshtime: 0,
        weigh: 0,
        switch: false,
        status: 0,
        state: 0

      };
    },
    async deleteExaExampleTest(row) {
      const res = await deleteExaExampleTest({ID: row.ID});
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
    },
    // // 用户详情跳转
    // usUserInfo(username) {
    //   this.$router.push({
    //     name: 'usUser',
    //     query: {
    //       username: username
    //     }
    //   })
    // },
    // 状态 开关 启用与关闭
    async statusSwitch(row) {
      const res = await updateExaExampleTest(row);
      if (res.code == 0) {
        if (row.status === 1) {
          this.$message.success("已启用")
        } else {
          this.$message.error("已停用")
        }
        // this.$message.success(`${row.status===1 ? '已启用':'已停用'}`)
        await this.getTableData()
      }
    },
  },
  async created() {
    await this.getTableData();
    // 富文本图片上传设置
    this.editorOption = quillRedefine(
        {
          // 图片上传的设置
          uploadConfig: {
            action: this.$store.state.uploadFullUrl,  // 必填参数 图片上传地址
            // 必选参数  res是一个函数，函数接收的response为上传成功时服务器返回的数据
            // 你必须把返回的数据中所包含的图片地址 return 回去
            res: (res) => {
              return this.$store.state.imgPrefix + res.data
            },
            name: 'file'  // 图片上传参数名
          }
        })
  }
};
</script>

<style>
.edit_container {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.ql-editor {
  height: 400px;
}

</style>

