import request from '@/utils/request'

//1. 请求品牌分页列表
export const reqTrademarkList = (page, limit) => {
  return request.get(
    `/admin/product/baseTrademark/${page}/${limit}`
  );
};

//2.删除某个品牌
export const reqDeleteTrademark = (id)=>{
  return request.delete(`/admin/product/baseTrademark/remove/${id}`)
}

//3.添加品牌
export const reqAddTrademark = (userInfo)=>{
  return request.post(`/admin/product/baseTrademark/save`,userInfo)
}

//4.修改品牌
export const reqUpdateTrademark = (userInfo)=>{
  return request.put(`/admin/product/baseTrademark/update`,userInfo)
}
