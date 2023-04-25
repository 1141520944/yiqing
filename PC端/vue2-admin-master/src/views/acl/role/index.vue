<template>
<div>
  <el-card class="box-card" :body-style="{ padding: '20px' }" shadow="hover">
    <el-form  ref="searchForm" label-width="100px" inline>
      <el-form-item>
        <el-input placeholder="用户名"  ></el-input>
      </el-form-item>
      <el-button type="primary" icon="el-icon-search" >查询</el-button>
      <el-button >清空</el-button>
    </el-form>
    <div style="margin: 20px 0">
      <el-button type="primary" >添 加</el-button>
      <el-button
        type="danger"
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
    </el-table>
  </el-card>
</div>
</template>

<script>
export default {
  name: 'Role'
}
</script>

<style scoped>

</style>
