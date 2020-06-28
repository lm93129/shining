import React, { useEffect, useState, useContext } from 'react'
import style from '../style.less'
import { Layout, Menu, Avatar } from 'antd'
import {
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  UserOutlined,
} from '@ant-design/icons'
import { Link, useLocation, useHistory } from 'react-router-dom'
import router from '../../router/consoleRoutes'
const { Header, Sider, Content } = Layout

const ConsoleFrame = (props: any) => {
  const [collapsed, setCollapsed] = useState(true)
  const location = useLocation()
  let history = useHistory()

  let selectedItemIndex = router.findIndex((e: any) => {
    return location.pathname.indexOf(e.path) !== -1
  })
  const toggle = () => {
    setCollapsed(!collapsed)
  }
  useEffect(() => {}, [])
  return (
    <Layout>
      <Sider
        trigger={null}
        collapsible
        collapsed={collapsed}
        style={{
          overflow: 'auto',
          height: '100vh',
          position: 'fixed',
          left: 0,
          zIndex: 2,
        }}
      >
        <div className={style.logo} />
        <Menu
          theme="dark"
          mode="inline"
          defaultSelectedKeys={[String(selectedItemIndex)]}
        >
          {router.map((e: any, index: any) => {
            return !e.hidden ? (
              <Menu.Item key={index}>
                {React.createElement(e.icon)}
                <span>{e.name}</span>
                <Link
                  to={{
                    pathname: e.path,
                  }}
                ></Link>
              </Menu.Item>
            ) : (
              ''
            )
          })}
        </Menu>
      </Sider>
      <Layout style={{ marginLeft: collapsed ? 80 : 200 }}>
        <Header style={{ padding: '0 20px 0px 0px', background: '#fff' }}>
          {React.createElement(
            collapsed ? MenuUnfoldOutlined : MenuFoldOutlined,
            {
              className: style.trigger,
              onClick: toggle,
            }
          )}
          <div style={{ float: 'right' }}>
            <Avatar icon={<UserOutlined />} style={{ margin: 10 }} />
          </div>
        </Header>
        <Content
          style={{
            margin: '24px 16px 0',
            overflow: 'initial',
            backgroundColor: '#fff',
            padding: 30,
          }}
        >
          {props.children}
        </Content>
      </Layout>
    </Layout>
  )
}
export default ConsoleFrame
