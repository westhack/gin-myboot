import context from '@/main'

export const defaultSearchFormData = [
{{- range .Fields}}{{if .FieldSearchType}}
  {
    name: '{{.FieldJson}}',
    label: '{{.FieldDesc}}',
    type: '{{.InputType}}',
    value: null,
    operator: '{{.FieldSearchType}}'
  },{{end}}
{{- end }}
]

export const defaultFormData = [
{{- range .Fields}}{{if .InputType}}
  {
    name: '{{.FieldJson}}',
    label: '{{.FieldDesc}}',
    type: '{{.InputType}}',
    value: null,
    rules: [{{range $v := .InputRules}}{{$v}},{{- end }}]
  },{{end}}
{{- end }}
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
{{- range .Fields}}
  {
    title: '{{.FieldDesc}}',
    dataIndex: '{{.FieldJson}}',
    align: '{{.TableAlign}}',
    width: '{{.TableWidth}}',
    type: '{{.InputType}}',
    value: null,
    editable: {{if .InputType}}true{{else}}false{{end}},
    isSearch: {{if .FieldSearchType}}true{{else}}false{{end}},
    isForm: {{if .InputType}}true{{else}}false{{end}},
    hiddenPopover: {{.HiddenPopover}}
  },
{{- end }}
  {
    title: '操作',
    width: '100px',
    dataIndex: 'action',
    scopedSlots: { customRender: 'action' },
    align: 'center',
    fixed: 'right'
  }
]
