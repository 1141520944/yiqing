<template>
  <div>
    <el-card class="box-card" :body-style="{ padding: '20px' }" shadow="hover">
      <div slot="header">
        <el-button icon="el-icon-plus" type="primary" @click="addTrademark">添加品牌</el-button>
      </div>
      <el-table
        :data="tableData"
        border
        style="width: 100%">
        <el-table-column
          type="index"
          label="序号"
          width="80" align="center"/>
        <el-table-column
          prop="tmName"
          label="品牌名称"
          align="center"/>
        <el-table-column
          prop="logoUrl"
          label="品牌logo" align="center">
          <template v-slot="{row}">
            <el-image
              style="height: 100px"
              :src="row.logoUrl"
              draggable="false"
              :preview-src-list="[row.logoUrl]">
            </el-image>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
        align="center"
        >
          <template v-slot="{row}">
            <el-button type="warning" icon="el-icon-edit" size="mini" @click="editTrademark(row)">修改</el-button>
            <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteTrademarkItem(row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        background
        layout="prev, pager, next,sizes"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page.sync="current"
        :page-sizes="[3, 5, 7]"
        :page-size="limit"
        :page-count="pages">
      </el-pagination>
    </el-card>
<!--    弹出框-->
    <el-dialog
      :title="ruleForm.id? '修改品牌' : '添加品牌'"
      :visible.sync="dialogVisible"
      width="30%">
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px">
        <el-form-item label="品牌名称" prop="tmName">
          <el-input v-model="ruleForm.tmName"></el-input>
        </el-form-item>
        <el-form-item label="品牌logo" prop="logoUrl">
          <el-upload
            class="avatar-uploader"
            action="http://gmall-h5-api.atguigu.cn/admin/product/upload"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload">
            <img v-if="ruleForm.logoUrl" :src="ruleForm.logoUrl" class="avatar">
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="dialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="confirm('ruleForm')">确 定</el-button>
  </span>
    </el-dialog>
  </div>
</template>

<script>

import { reqAddTrademark, reqDeleteTrademark, reqTrademarkList, reqUpdateTrademark } from '@/api/trademark'
import { Message } from 'element-ui'
import {cloneDeep} from 'lodash'

export default {
  name: 'Trademark',
  data(){
    return{
      current:1,//当前页
      limit:5,//每页几条
      pages:1,//总页数
      tableData: [],
      dialogVisible:false,//控制弹出框的显隐
      ruleForm:{
        tmName:'',
        logoUrl:''
      },
      rules: {
        tmName: [
          { required: true, message: '请输入品牌名称', trigger: 'blur' },
          { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
        ],
        logoUrl: [
          { required: true, message: '请选择图片', trigger: 'change' }
        ],
      }
    }
  },
  mounted() {
    this.initTrademark()
  },
  methods:{
    //初始化trademark
    async initTrademark(){
      const res = await reqTrademarkList(this.current,this.limit)
      this.limit = res.data.size
      this.pages = res.data.pages
      this.tableData = res.data.records
    },
  //  页数变化时
    handleSizeChange(val){
      this.current = 1
      this.limit = val
      this.initTrademark()
    },
  //  当前页改变时
    handleCurrentChange(val){
      this.current = val
      this.initTrademark()
    },
  //  删除trademark
    async deleteTrademarkItem(id){
      try {
        await reqDeleteTrademark(id)
        await this.initTrademark()
      }catch (e) {
        Message.error(e.message)
      }
    },
  //  添加品牌
    addTrademark(){
      this.ruleForm={
        tmName:'',
        logoUrl:''
      }
      this.dialogVisible = true
    },
  //  编辑trademark
    editTrademark(row){
      this.dialogVisible = true
      const newRow = cloneDeep(row)
      this.ruleForm = newRow
    },
  //  提交信息
    confirm(formName){
      this.$refs[formName].validate(async valid => {
        if (valid) {
          try {
            if(this.ruleForm.id){
              await reqUpdateTrademark(this.ruleForm)
              Message.success('修改成功')
            }else {
              await reqAddTrademark(this.ruleForm)
              Message.success('新增成功~')
            }
            this.dialogVisible = false
            await this.initTrademark()
          }catch (e) {
            this.dialogVisible = false
            Message.error(e.message)
          }
        } else {
          console.log('error submit!!');
          return false;
        }
      });

    },
  //上传成功
    handleAvatarSuccess(res, file){
      this.ruleForm.logoUrl = URL.createObjectURL(file.raw);
    },
    //上传之前
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!');
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
      }
      return isJPG && isLt2M;
    }
  }
}
</script>

<style>
.el-pagination{
  text-align: center;
  margin-top: 20px;
}
.avatar-uploader .el-upload {
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}


.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>
<style scoped>
.box-card{
  margin-top: 20px;
}
.avatar-uploader{
  width: 178px;
  height: 178px;
  display: block;
}

.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}

</style>
