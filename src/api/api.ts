import axios from '../libs/api.request'

export const getData = (params:any) => {
  return axios.request({
    url: `appfile/getappinfos`,
    method: "get",
    params
  })
}
