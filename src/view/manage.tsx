import React, { useEffect, useState, useContext } from 'react'
import ConsoleFrame from './compoents/console-frame'
import { Divider, Row, Col, Table } from 'antd'
import { getAppList } from '../api/api'
// import { GlobalContext } from '../components/global-context'

const Manage = () => {
  const [pageSize, setPageSize] = useState(10)
  const [pageIndex, setPageIndex] = useState(1)
  const [total, setTotal] = useState(0)
  const [list, setList] = useState([])
  // const { globalData, setGlobalData } = useContext(GlobalContext)
  
  const getList = () => {
    getAppList({
      page: pageIndex,
      pageSize: pageSize,
      project_id: 'asoco-app',
      version_type: 'android',
    }).then((res: any) => {
      setTotal(res.data.total)
      setList(res.data.records)
    })
  }
  const onPageChange = (page: any) => {
    setPageIndex(page)
  }
  const columns = [
    {
      title: '订单单号',
      dataIndex: 'id',
      key: 'id',
    },
  ]
  useEffect(() => {
    getList()
  }, [pageIndex])
  return (
    <ConsoleFrame>
      <Row>
        {/* <Col span={8}>
          时间范围：
          <RangePicker />
        </Col> */}
      </Row>
      <Divider />
      <Row>
        <Col span={24}>
          <Table
            rowKey="id"
            dataSource={list}
            columns={columns}
            pagination={{
              current: pageIndex,
              total: total,
              pageSize: pageSize,
              onChange: onPageChange,
            }}
          />
        </Col>
      </Row>
    </ConsoleFrame>
  )
}
export default Manage
