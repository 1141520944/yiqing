import request from '@/utils/request'

//获取用户列表
export const reqGetUserList = (page,limit,searchParams)=>{
  return request.get(`/admin/acl/user/${page}/${limit}`,{params:searchParams})
}

//删除用户列表
export const reqDeleteUserList = ids=>{
  return request.delete(`/admin/acl/user/batchRemove`,{ data:ids })
}

//新增用户
export const reqAddUser = userInfo=>{
  return request.post(`/admin/acl/user/save`,userInfo)
}

//修改用户
export const reqModifyUser = userInfo=>{
  return request.put(`/admin/acl/user/update`,userInfo)
}

//删除用户
export const reqDeleteUser = id=>{
  return request.delete(`/admin/acl/user/remove/${id}`)
}

//获取用户角色列表
export  const reqGetRoleList = userId=>{
  return request.get(`/admin/acl/user/toAssign/${userId}`)
}

//分配角色
export const reqToAssign = (userId,roleId)=>{
  return request.post(`/admin/acl/user/doAssign`,null,{
    params:{
      roleId,
      userId
    }
  })
}
