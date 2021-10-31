<template>
  <a-card :bordered="false">

    <edit-table
      ref="editTable"
      rowKey="{{.RowKey}}"
      :columns="columns"
      :table-action-bars="tableActionBars"
      :tableHeaderButtons="tableHeaderButtons"
      :create-api-url="api.create"
      :update-api-url="api.update"
      :list-api-url="api.getList"
      :delete-api-url="api.delete"
      :delete-by-ids-api-url="api.deleteByIds"
      :scroll="tableScroll"
      :searchFormData="searchFormData"
      :formData="formData"
      :isDownload="{{.IsDownload}}"
      :defaultFormWidth="'600px'"
      :showPopover="{{.ShowPopover}}"
      :isFormCreate="{{.IsFormCreateUpdate}}"
      :isFormUpdate="{{.IsFormCreateUpdate}}"
      :isBatchDelete="{{.IsBatchDelete}}"
      :isTableDelete="{{.IsTableDelete}}"
      :isSearch="{{.IsSearch}}"
      :isTableCreate="{{.IsTableCreateUpdate}}"
      :isTableUpdate="{{.IsTableCreateUpdate}}"
      :isDblclickUpdate="{{.IsDblclickUpdate}}"
      @change="onChange"
    ></edit-table>

  </a-card>
</template>

<script>

import { api } from '@/modules/{{.ModuleName}}/api/{{.Abbreviation}}'
import httpResponse from '@/mixins/httpResponse'
import _ from 'lodash'
import { defaultFormData, defaultSearchFormData, columns } from './form'

export default {
  name: '{{.StructName}}',
  components: {},
  mixins: [httpResponse],
  data () {
    const vm = this

    return {
      multiple: true,
      tableScroll: { x: 1200 },
      columns: columns,
      optionAlertShow: false,
      pageParam: 1,
      formData: defaultFormData,
      searchFormData: defaultSearchFormData,
      api: api,
      selectedRowKeys: [],
      selectedRows: [],
      tableHeaderButtons: [],
      tableActionBars: []
    }
  },
  created () {
    this.searchFormData = _.cloneDeep(defaultSearchFormData)
  },
  mounted () {},
  methods: {
    onChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },

    handleSubmit (e) {
    },

    handleReset () {
    },

    handleSearchReset () {
    }

  }
}
</script>

<style scoped>
</style>
