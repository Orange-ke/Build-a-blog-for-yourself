<template>
    <div>
      <Editor :init="init" v-model="content"></Editor>
    </div>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'
import tinymce from './tinymce.min'
import './icons/default/icons.min'
import './themes/silver/theme.min'
import './langs/zh_CN'
// 注册插件
import './plugins/preview/plugin.min'
import './plugins/paste/plugin.min'
import './plugins/code/plugin.min'
import './plugins/wordcount/plugin.min'
import './plugins/image/plugin.min'
import './plugins/imagetools/plugin.min'
export default {
  components: { Editor },
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      init: {
        language: 'zh_CN',
        height: 1000,
        margin: '0',
        padding: '0',
        plugins: 'preview paste code wordcount image imagetools',
        branding: false,
        toolbar: ['undo redo | formselect | alignleft aligncenter alignright alignjustify', 'preview paste code | image imagetools',],
        // 上传图片
        images_upload_handler: async (blobInfo, succFun, failFun) => {
          let formData = new FormData()
          formData.append('file', blobInfo.blob(), blobInfo.name)
          const { data: res } = await this.$http.post('upload', formData)
          if (res.status !== 200) {
            failFun(res.message)
            return
          }
          succFun(res.url)
        },
        imagetools_cors_hosts: ['*']
      },
      content: this.value
    }
  },
  watch: {
    value (newV) {
      this.content = newV
    },
    content (newV) {
      this.$emit('input', newV)
    }
  }
}
</script>

<style scoped>
@import url('./skins/ui/oxide/skin.min.css');
</style>
