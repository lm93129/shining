import { BrowserRouter as Router, Link } from 'react-router-dom'
import React, { Suspense } from 'react'
import { Spin, Layout, Menu } from 'antd'
import RouterCofig from './router/index'
import './global.less'

const { Header, Content, Footer } = Layout
const backgroundImage = require('./assets/img/top_bg.png')

const App = () => (
  <Router>
    <Layout style={{ background: '#fff' }}>
      <Header style={{ padding: 0, backgroundColor: '#fff' }}>
        <div
          style={{
            backgroundImage: `url(${backgroundImage}`,
            backgroundSize: 'contain',
            height: 64,
            backgroundRepeat: 'repeat'
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
                  transform: 'translate(-50%, -50%)'
                }}
              />
            }
          >
            <RouterCofig></RouterCofig>
          </Suspense>
        </div>
      </Content>
    </Layout>
  </Router>
)
export default App
