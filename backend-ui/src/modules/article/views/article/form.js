import context from '@/main'

export const defaultSearchFormData = [
  {
    name: 'content',
    label: '内容',
    type: 'input',
    value: null,
    operator: '='
  },
  {
    name: 'title',
    label: '标题',
    type: 'input',
    value: null,
    operator: '='
  },
]

export const defaultFormData = [
  {
    name: 'content',
    label: '内容',
    type: 'input',
    value: null,
    rules: [{ required: true, message: '不能为空' }]
  },
  {
    name: 'title',
    label: '标题',
    type: 'input',
    value: null,
    rules: [{ required: true, message: '不能为空' }]
  },
]

export const columns = [
  {
    title: '#',
    scopedSlots: { customRender: 'serial' },
    width: '50px',
    align: 'center',
    dataIndex: 'no',
    type: 'hidden'
  },
  {
    title: '内容',
    dataIndex: 'content',
    align: 'false',
    width: 'false',
    type: 'input',
    value: '',
    editable: true,
    isSearch: true,
    isForm: true,
    hiddenPopover: false
  },
  {
    title: '标题',
    dataIndex: 'title',
    align: 'false',
    width: 'false',
    type: 'input',
    value: '',
    editable: true,
    isSearch: true,
    isForm: true,
    hiddenPopover: false
  },
  {
    title: '操作',
    width: '100px',
    dataIndex: 'action',
    scopedSlots: { customRender: 'action' },
    align: 'center',
    fixed: 'right'
  }
]
