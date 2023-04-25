<template>
<div>
  <el-card class="box-card" :body-style="{ padding: '20px' }" shadow="hover">
    <el-form :model="searchForm" :rules="rules" ref="searchForm" label-width="100px" inline>
        <el-form-item>
          <el-input placeholder="用户名" v-model="searchForm.username" ></el-input>
        </el-form-item>
      <el-button type="primary" icon="el-icon-search" @click="search">查询</el-button>
      <el-button @click="reset">清空</el-button>
    </el-form>
    <div style="margin: 20px 0">
      <el-button type="primary" @click="addUserShowDialog">添 加</el-button>
      <el-button
        type="danger"
        @click="deleteAll"
      >批量删除</el-button
      >
    </div>
    <el-table
      :data="tableData"
      border
      style="width: 100%"
      @selection-change="handleSelectable"
    >
      <el-table-column type="selection" align="center"/>
      <el-table-column
        prop="username"
        label="用户名"
        width="180" align="center"/>
      <el-table-column
        prop="nickName"
        label="用户昵称"
        align="center"/>
      <el-table-column
        prop="roleName"
        label="角色列表"
        align="center"/>
      <el-table-column
        prop="gmtCreate"
        label="创建时间"
        align="center"/>
      <el-table-column
        prop="gmtModified"
        label="更新时间"
        align="center"/>
      <el-table-column
        prop="address"
        label="操作"
        align="center">
        <template v-slot="{row}">
          <el-button type="info" icon="el-icon-user-solid" size="mini" @click="getUserRoleList(row.username,row.id)"></el-button>
          <el-button type="primary" icon="el-icon-edit-outline" size="mini" @click="addUserShowDialog(row)"></el-button>
          <el-popconfirm
            :title="`确定要删除${row.username}吗？`"
            @onConfirm="deleteUser(row.id)"
          >
            <template #reference>
              <el-button type="danger" icon="el-icon-delete" size="mini"></el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      background
      layout="prev, pager, next,jumper,sizes,total"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :page-sizes="[ 5, 7,10]"
      :page-size="limit"
      :current-page="current"
      :total="total">
    </el-pagination>
  </el-card>
  <el-dialog
    :title="userForm.id? '修改用户':'添加用户'"
    :visible.sync="dialogVisible"
    width="50%">
    <el-form :model="userForm" :rules="userRules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
      <el-form-item label="用户名称" prop="username">
        <el-input v-model="userForm.username"></el-input>
      </el-form-item>
      <el-form-item label="用户昵称" prop="nickName">
        <el-input v-model="userForm.nickName"></el-input>
      </el-form-item>
      <el-form-item label="用户密码" prop="password">
        <el-input type="password" v-model="userForm.password"></el-input>
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
    <el-button @click="dialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="updateOrAddUser">确 定</el-button>
  </span>
  </el-dialog>
  <el-dialog
    title="设置角色"
    :visible.sync="setRoleVisible"
    width="50%">
    <el-form :model="roleForm" ref="ruleForm" label-width="100px">
      <el-form-item label="用户名">
        <el-input v-model="roleForm.username" disabled></el-input>
      </el-form-item>
      <el-form-item label="角色列表">
        <el-checkbox :indeterminate="isIndeterminate" v-model="checkAll" @change="handleCheckAllChange">全选</el-checkbox>
        <div style="margin: 15px 0;"></div>
        <el-checkbox-group v-model="allIds" @change="handleCheckedCitiesChange">
          <el-checkbox v-for="role in roleForm.allRolesList" :label="role.id" :key="role.id">{{role.roleName}}</el-checkbox>
        </el-checkbox-group>
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
    <el-button @click="cancle">取 消</el-button>
    <el-button type="primary" @click="submit">确 定</el-button>
  </span>
  </el-dialog>
</div>
</template>

<script>
import {
  reqAddUser,
  reqDeleteUser,
  reqDeleteUserList,
  reqGetRoleList,
  reqGetUserList,
  reqModifyUser, reqToAssign
} from '@/api/acl/user'
import {mapState} from 'vuex'
import { Message } from 'element-ui'
import {cloneDeep} from 'lodash'
export default {
  name: 'User',
  data(){
    return{
      searchName:'',
      current:1,
      limit:5,
      pages:1,
      tableData:[],
      total:0,
      searchForm:{username:''},
      rules:{
        username: [
          { required: true, message: '用户名必须输入' },
          { min: 3, message: '用户名不能小于3位' },
        ],
      },
      userList:[],
      dialogVisible:false,
      userForm:{
        username:'',
        nickName:'',
        password:''
      },
      userRules:{
        username:[
      { required: true, message: '请输入用户名称', trigger: 'blur' },
      { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
    ],
        nickName:[
          { required: true, message: '请输入用户名称', trigger: 'blur' },
          { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入用户密码', trigger: 'blur' },
          { min: 3, max: 12, message: '长度在 3 到12 个字符', trigger: 'blur' }
        ]
      },
      setRoleVisible:false,
      roleForm:{
        username:'',
        assignRoles:[],
        allRolesList:[]
      },
      isIndeterminate:true,
      checkAll:true,
      allIds:[],
      userId:''
    }
  },
  computed:{
    ...mapState('user',['name'])
  },
  watch:{
    roleForm:{
      handler(){
        this.allIds = this.roleForm.assignRoles.map(item=>item.id)
      },
      deep:true,
      immediate:false
    }
  },
  created() {
    this.getUserList()
  },
  methods:{
    //获取用户列表
    async getUserList(name){
      const res = await reqGetUserList(this.current,this.limit,name===null? this.name:name)
      this.total = res.data.total
      this.tableData = res.data.items
    },
    //点击分配角色总按钮
    handleCheckAllChange(val) {
      this.roleForm.assignRoles = val ? this.roleForm.allRolesList : [];
      this.isIndeterminate = false;
    },
    //点击单个分配角色按钮
    handleCheckedCitiesChange(value) {
      let checkedCount = value.length;
      this.checkAll = checkedCount === this.roleForm.allRolesList.length;
      this.isIndeterminate = checkedCount > 0 && checkedCount < this.roleForm.allRolesList.length;
    },
    //分配角色提交
    async submit(){
      try {
        await reqToAssign(this.userId,this.allIds.join(','))
        Message.success('分配权限成功~')
      }catch (e) {
        Message.error(e.message)
      }
      this.setRoleVisible = false
    },
    //点击取消角色分配按钮
    cancle(){
      this.setRoleVisible = false
    },
    //获取用户角色列表
    async getUserRoleList(username,id){
      this.userId = id
      this.setRoleVisible = true
      this.roleForm.username = username
      try {
        const res = await  reqGetRoleList(id)
        console.log(res)
        this.roleForm.assignRoles = res.data.assignRoles
        this.roleForm.allRolesList = res.data.allRolesList
      }catch (e) {
        Message.error(e.message)
      }
    },
  //  搜索
    async search(){
      if(!this.searchForm.username)return
      await this.getUserList({username:this.searchForm.username})
    },
  //  内容条数改变
    handleSizeChange(val){
      this.limit = val
      this.current = 1
      this.getUserList()
    },
  //  当前页发生变化
    handleCurrentChange(val){
      this.current = val
      this.getUserList()
    },
  //  表单重置
    reset(){
      this.searchForm.username=''
      this.getUserList()
    },
  //  多选框
    handleSelectable(selection){
      this.userList = selection
    },
  //  批量删除
    async deleteAll(){
      const ids = this.userList.map(item=>item.id)
      try {
        await reqDeleteUserList(ids)
        await this.getUserList()
        Message.success('删除成功~')
      }catch (e) {
        Message.error(e.message)
      }
    },
  //  添加用户框的显示
    addUserShowDialog(row={}){
      if(row.id){
        this.userForm = cloneDeep(row)
      }else {
        this.userForm={
          username:'',
          nickName:'',
          password:''
        }
      }
      this.dialogVisible=true
    },
  //  添加用户或修改用户
    async updateOrAddUser(){
      try {
        if(this.userForm.id){
          await reqModifyUser(this.userForm)
          Message.success('更新用户成功~')
        }else{
          await reqAddUser(this.userForm)
          Message.success('添加用户成功~')
        }
      }catch (e){
        Message.error(e.message)
      }
      await this.getUserList()
      this.dialogVisible = false
    },
  //  删除用户
    async deleteUser(id){
      try {
        await reqDeleteUser(id)
        Message.success('删除用户成功~')
      }catch (e){
        Message.error(e.message)
      }
      await this.getUserList()
    }
  }
}
</script>

<style scoped>
.box-card{
  margin-top: 20px;
}
.el-pagination{
  text-align: center;
  margin-top: 20px;
}

</style>
