import { axios } from '@/utils/request'

const API_VERSION = '/article/article/'

export const api = {
  getList: API_VERSION + 'getList',
  create: API_VERSION + 'create',
  update: API_VERSION + 'update',
  delete: API_VERSION + 'delete',
  deleteByIds: API_VERSION + 'deleteByIds',
  find: API_VERSION + 'find'
}

export default api

export const createSysArticle = (data) => {
  return axios({
    url: api.create,
    method: 'post',
    data: data
  })
}

export const deleteSysArticle = (data) => {
  return axios({
    url: api.delete,
    method: 'post',
    data: data
  })
}

export const deleteSysArticleByIds = (data) => {
  return axios({
    url: api.deleteByIds,
    method: 'post',
    data: data
  })
}

export const updateSysArticle = (data) => {
  return axios({
    url: api.update,
    method: 'post',
    data: data
  })
}

export const findSysArticle = (params) => {
  return axios({
    url: api.find,
    method: 'get',
    params: params
  })
}

export const getSysArticleList = (data) => {
  return axios({
    url: api.getList,
    method: 'post',
    data: data
  })
}
