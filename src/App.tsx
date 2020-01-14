import { BrowserRouter as Router, Link } from "react-router-dom"
import React, { Suspense } from "react"
import { Spin, Layout, Menu } from "antd"
const { Header, Content, Footer } = Layout
import RouterCofig from "./router/index"
import "./global.less"
const App = () => (
  <Router>
    <Layout className="layout" style={{ background: "#fff" }}>
      <Header
        style={{ padding: 0, backgroundColor: "#fff"}}
      >
        <img src={require("@/assets/img/top_bg.png")} alt="" style={{width:'100%'}}/>
      </Header>
      <Content style={{ padding: "0 50px" }}>
        <div style={{ padding: 24, minHeight: 280 }}>
          <Suspense
            fallback={
              <Spin
                size="large"
                style={{
                  position: "fixed",
                  top: "50%",
                  left: "50%",
                  transform: "translate(-50%, -50%)"
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
