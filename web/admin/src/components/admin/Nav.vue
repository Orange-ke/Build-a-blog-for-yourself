<template>
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="log">
      <span>{{ collapsed ? 'blog' : 'My Blog' }}</span>
    </div>
    <a-menu theme="dark" mode="inline" @click="gotoPage">
      <a-menu-item key="index">
        <a-icon type="dashboard"></a-icon>
        <span>仪表盘</span>
      </a-menu-item>
      <a-sub-menu>
        <span slot="title"><a-icon type="file"></a-icon><span>文章管理</span></span>
        <a-menu-item key="addart">
          <a-icon type="form"></a-icon>
          <span>写文章</span>
        </a-menu-item>
        <a-menu-item key="artlist">
          <a-icon type="snippets"></a-icon>
          <span>文章列表</span>
        </a-menu-item>
      </a-sub-menu>
      <a-menu-item key="catelist">
        <a-icon type="book"></a-icon>
        <span>分类列表</span>
      </a-menu-item>
      <a-menu-item key="userlist">
        <a-icon type="user"></a-icon>
        <span>用户列表</span>
      </a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data () {
    return {
      collapsed: false
    }
  },
  methods: {
    gotoPage (item) {
      this.$router.push('/' + item.key).catch(() => {})
    }
  }
}
</script>

<style scoped>
  .log {
    height: 32px;
    margin: 16px;
    background-color: rgba(255, 255, 255, .75);
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 17px;
    border-radius: 10px;
  }
</style>
