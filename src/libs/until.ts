export const deepClone = (target:any) => {
  if (typeof target === 'object') {
    let cloneTarget:any = Array.isArray(target) ? []:{}
    for (const key in target) {
      if (target.hasOwnProperty(key)) {
        cloneTarget[key] = deepClone(target[key])
      }
      return cloneTarget
    }  
  }else {
    return target
  }
}

export const  parseQueryString = (url:string) => {
  let obj:any = {}
  let keyvalue = []
  let key = "",
    value = ""
  let paraString = url.substring(url.indexOf("?") + 1, url.length).split("&")
  for (let i in paraString) {
    keyvalue = paraString[i].split("=")
    key = keyvalue[0]
    value = keyvalue[1]
    obj[key] = value
  }
  return obj
}