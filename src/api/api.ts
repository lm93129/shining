import axios from '../libs/api.request'

export const getAppList = (params: any) => {
  return axios.request({
    url: `/appFile/appInfos`,
    method: 'get',
    params,
  })
}
