<template>
  <div>

    <div id="userLayout" :class="['user-layout-wrapper', device]">
      <div class="container">
        <div class="top">
          <div class="header">
            <a href="/">
              <img src="~@/assets/logo.png" class="logo" alt="logo">
              <span class="title">Myboot Vue</span>
            </a>
          </div>
          <div class="desc">
            Myboot + Vue 敏捷后台脚手架
          </div>
        </div>

        <div style="width: 80%; margin: auto">
          <a-steps :current="current">
            <a-step v-for="item in steps" :key="item.title" :title="item.title" />
          </a-steps>
          <div class="steps-content">
            <div style="padding: 10px" v-show="current == 0">
              <FormGenerator ref="form" :fields="formData">
                <template slot="footer">
                  <div style="text-align: center">
                    <a-button v-if="current < steps.length - 1" type="primary" @click="next">
                      下一步
                    </a-button>
                  </div>
                </template>
              </FormGenerator>
            </div>
            <div style="padding: 10px" v-show="current == 1">
              <div>
                <a-spin tip="Loading...">
                  <div class="spin-content">
                    系统配置中。。。
                  </div>
                </a-spin>
              </div>

              <div class="steps-action">
              </div>
            </div>
            <div style="padding: 10px" v-show="current == 2">
              <a-result
                status="success"
                title="系统安装成功!"
                sub-title=""
              >
                <template #extra>
                  <a-button key="home" type="primary" @click="goHome">
                    返回首页
                  </a-button>
                  <a-button key="help">
                    查看帮助
                  </a-button>
                </template>
              </a-result>
            </div>
            <div style="padding: 10px" v-show="current == 3">
              <a-result
                status="error"
                title="系统安装失败!"
                sub-title=""
              >
                <template #extra>
                  <a-button key="home" type="primary" @click="goHome">
                    返回首页
                  </a-button>
                  <a-button key="help">
                    查看帮助
                  </a-button>
                </template>
              </a-result>
            </div>
          </div>
        </div>

        <div class="footer">
          <div class="links">
            <a href="https://github.com/westhack/gin-myboot">github</a>
            <a href="http://docs.limaopu.com/">在线文档</a>
            <a href="http://demo.limaopu.com/">演示环境</a>
          </div>
          <div class="copyright">
            Copyright &copy; 2021 @westhack
          </div>
        </div>
      </div>
    </div>

  </div>
</template>
<script>
import { getSystemConfig, setSystemConfig, pong } from '../api/api'
import FormGenerator from '@/components/FormGenerator/FormGenerator'
import { mixinDevice } from '@/utils/mixin'
import { httpResponseCode } from '@/constants/httpResponseCode'

export default {
  name: 'Install',
  components: {
    FormGenerator
  },
  mixins: [mixinDevice],
  data () {
    return {
      formLoading: false,
      formData: {},
      current: 0,
      isDemo: false,
      config: {},
      steps: [
        {
          title: '配置系统',
          content: '配置系统'
        },
        {
          title: '导入数据',
          content: '导入数据'
        },
        {
          title: '完成',
          content: '完成'
        }
      ]
    }
  },
  async beforeCreate () {
    const res = await getSystemConfig()
    if (res.code == httpResponseCode.SUCCESS) {
      this.formData = res.data.config
    } else {
      this.$message.error(res.message)
      if (res.data.install == 'ok') {
        this.current = 2
      }
    }
  },
  methods: {
    next () {
      this.current++
      if (this.current == 1) {
        const values = this.$refs['form'].getFieldsValue()
        this.config = values
        this.setSystemConfig()
      } else if (this.current == 2) {

      }
    },
    async setSystemConfig () {
      const interval = window.setInterval(() => {
        pong().then(res => {
          if (res.code == 200) {
            this.current = 2
            window.clearInterval(interval)
            this.$message.success(res.message)
          }
        })
      }, 2000)

      const values = this.$refs['form'].getFieldsValue()
      const res = await setSystemConfig(values)

      if (res.data && res.data.install == 'ok') {
      } else {
        this.$message.error(res.message)
        this.current = 3
      }

      console.log(res)
    },
    prev () {
      this.current--
    },
    goHome () {
      this.$router.push('/')
    }
  }
}
</script>
<style scoped lang="less">
.steps-content {
  margin-top: 16px;
  border: 1px dashed #e9e9e9;
  border-radius: 6px;
  background-color: #fafafa;
  min-height: 200px;
}

.steps-action {
  margin-top: 24px;
}

#userLayout.user-layout-wrapper {
  height: 100%;

  &.mobile {
    .container {
      .main {
        max-width: 368px;
        width: 98%;
      }
    }
  }

  .container {
    width: 100%;
    min-height: 100%;
    background: #f0f2f5 url(~@/assets/background.svg) no-repeat 50%;
    background-size: 100%;
    padding: 50px 0 54px;
    position: relative;

    a {
      text-decoration: none;
    }

    .top {
      text-align: center;

      .header {
        height: 44px;
        line-height: 44px;

        .badge {
          position: absolute;
          display: inline-block;
          line-height: 1;
          vertical-align: middle;
          margin-left: -12px;
          margin-top: -10px;
          opacity: 0.8;
        }

        .logo {
          height: 44px;
          vertical-align: top;
          margin-right: 16px;
          border-style: none;
        }

        .title {
          font-size: 23px;
          color: rgba(0, 0, 0, .85);
          font-family: Avenir, 'Helvetica Neue', Arial, Helvetica, sans-serif;
          font-weight: 600;
          position: relative;
          top: 2px;
        }
      }
      .desc {
        font-size: 14px;
        color: rgba(0, 0, 0, 0.45);
        margin-top: 12px;
        margin-bottom: 20px;
      }
    }

    .main {
      min-width: 260px;
      width: 368px;
      margin: 0 auto;
    }

    .footer {
      width: 100%;
      bottom: 0;
      padding: 0 16px;
      margin: 48px 0 24px;
      text-align: center;

      .links {
        margin-bottom: 8px;
        font-size: 14px;
        a {
          color: rgba(0, 0, 0, 0.45);
          transition: all 0.3s;
          &:not(:last-child) {
            margin-right: 40px;
          }
        }
      }
      .copyright {
        color: rgba(0, 0, 0, 0.45);
        font-size: 14px;
      }
    }
  }
}
.spin-content {
  border: 1px solid #91d5ff;
  background-color: #e6f7ff;
  padding: 30px;
}
</style>
