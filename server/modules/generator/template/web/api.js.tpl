import { axios } from '@/utils/request'

const API_VERSION = '/{{.ModuleName}}/{{.Abbreviation}}/'

export const api = {
  getList: API_VERSION + 'getList',
  create: API_VERSION + 'create',
  update: API_VERSION + 'update',
  delete: API_VERSION + 'delete',
  deleteByIds: API_VERSION + 'deleteByIds',
  find: API_VERSION + 'find'
}

export default api

export const create{{.StructName}} = (data) => {
  return axios({
    url: api.create,
    method: 'post',
    data: data
  })
}

export const delete{{.StructName}} = (data) => {
  return axios({
    url: api.delete,
    method: 'post',
    data: data
  })
}

export const delete{{.StructName}}ByIds = (data) => {
  return axios({
    url: api.deleteByIds,
    method: 'post',
    data: data
  })
}

export const update{{.StructName}} = (data) => {
  return axios({
    url: api.update,
    method: 'post',
    data: data
  })
}

export const find{{.StructName}} = (params) => {
  return axios({
    url: api.find,
    method: 'get',
    params: params
  })
}

export const get{{.StructName}}List = (data) => {
  return axios({
    url: api.getList,
    method: 'post',
    data: data
  })
}
