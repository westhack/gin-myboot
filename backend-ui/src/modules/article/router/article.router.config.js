// eslint-disable-next-line
import { RouteView } from '@/layouts'

const articleRouter =
  {
    path: 'article',
    name: 'article',
    component: () => import('@/modules/article/views/article/table.vue'),
    meta: { title: 'article', icon: '', keepAlive: true, permission: [] }
  }

export default articleRouter
