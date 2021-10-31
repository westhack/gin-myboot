import { axios } from '@/utils/request'

const API_VERSION = 'install'

const api = {
  getSystemConfig: API_VERSION + '/getSystemConfig',
  setSystemConfig: API_VERSION + '/setSystemConfig',
  pong: API_VERSION + '/pong',
  initdb: API_VERSION + '/initdb',
  checkdb: API_VERSION + '/checkdb'
}

export function getSystemConfig (parameter) {
  return axios({ url: api.getSystemConfig, method: 'post', data: parameter })
}

export function setSystemConfig (parameter) {
  return axios({ url: api.setSystemConfig, method: 'post', data: parameter })
}
export function pong (parameter) {
  return axios({ url: api.pong, method: 'get', params: parameter })
}
export function initdb (parameter) {
  return axios({ url: api.initdb, method: 'post', data: parameter })
}
export function checkdb (parameter) {
  return axios({ url: api.checkdb, method: 'post', data: parameter })
}