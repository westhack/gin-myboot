// eslint-disable-next-line
import { RouteView } from '@/layouts'

const {{.Abbreviation}}Router =
  {
    path: '{{.Abbreviation}}',
    name: '{{.Abbreviation}}',
    component: () => import('@/modules/{{.ModuleName}}/views/{{.Abbreviation}}/table.vue'),
    meta: { title: '{{.Description}}', icon: 'menu', keepAlive: true, permission: [] }
  }

export default {{.Abbreviation}}Router
