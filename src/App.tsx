import { BrowserRouter as Router } from 'react-router-dom'
import React, { Suspense } from 'react'
import { Spin, Layout } from 'antd'
import RouterCofig from './router/index'
import './global.less'

const { Content } = Layout

const App = () => (
  <Router>
    <Layout className="layout" style={{ background: '#fff' }}>
      <Content>
        <div>
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
            <RouterCofig></RouterCofig>
          </Suspense>
        </div>
      </Content>
    </Layout>
  </Router>
)
export default App
