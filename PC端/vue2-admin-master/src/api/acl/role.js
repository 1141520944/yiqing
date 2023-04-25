import request from '@/utils/request'

//请求角色列表
export const reqRoleList = (page,limit)=>{
  return request.get(`/admin/acl/role/${page}/${limit}`)
}
