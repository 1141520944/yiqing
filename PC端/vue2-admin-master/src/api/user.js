import request from '@/utils/request'

export function login({username, password}) {
  return request.post('/admin/acl/index/login',{username, password})
}

export function getInfo(token) {
  return request.get('/admin/acl/index/info',{params:{token}})
}

export function logout() {
  return request.post('/admin/acl/index/logout')
}
