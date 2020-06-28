import React, { createContext, useState } from 'react'
export const GlobalContext = createContext({
  globalData: {
    appId: ''
  },
  setGlobalData: undefined,
})

export const GlobalProvider = (props: any) => {
  const [globalData, setGlobalData] = useState({
    appId: ''
  })
  return (
    <GlobalContext.Provider value={{ globalData, setGlobalData }}>
      {props.children}
    </GlobalContext.Provider>
  )
}

export const GlobalConsumer = GlobalContext.Consumer
