import React, { useEffect,  Suspense } from 'react'
import style from '../style.less'
import { Layout, Spin } from 'antd'
import { Link, useHistory } from 'react-router-dom'
const { Header, Content } = Layout
const backgroundImage = require('../../assets/img/top_bg.png')

const Main = (props: any) => {
  let history = useHistory()
  useEffect(() => {

  }, [])
  return (
    <Layout style={{ background: '#fff' }}>
      <Header style={{ padding: 0, backgroundColor: '#fff' }}>
        <div
          style={{
            backgroundImage: `url(${backgroundImage}`,
            backgroundSize: 'contain',
            height: 64,
            backgroundRepeat: 'repeat',
          }}
        ></div>
      </Header>
      <Content style={{ padding: '0 50px' }}>
        <div style={{ padding: 24, minHeight: 280 }}>
          <Suspense
            fallback={
              <Spin
                size="large"
                style={{
                  position: 'fixed',
                  top: '50%',
                  left: '50%',
                  transform: 'translate(-50%, -50%)',
                }}
              />
            }
          >
            {props.children}
          </Suspense>
        </div>
      </Content>
    </Layout>
  )
}
export default Main
