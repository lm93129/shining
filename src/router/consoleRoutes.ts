import Manage from '../view/manage'
import {
  FundFilled,
} from '@ant-design/icons'
export default [
  {
    path: '/appdl/manage',
    component: Manage,
    exact: true,
    name: 'APP列表',
    icon: FundFilled,
  },
]
