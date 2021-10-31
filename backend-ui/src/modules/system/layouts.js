const layouts = {
  'headerCenter': {
    ReloadSystem: () => import('@/modules/system/components/ReloadSystem.vue'),
    Message: () => import('@/modules/websocket/components/Message.vue')
  }
}

export default layouts
