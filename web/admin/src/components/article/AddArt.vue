<template>
  <div>
    <a-card>
      <h3>{{ id ? '编辑文章' : '新增文章' }}</h3>
      <a-form-model :model="artInfo" ref="artInfoRef" :rules="artInfoRules">
        <a-form-model-item label="文章标题" prop="title">
          <a-input style="width: 300px;" v-model="artInfo.title"></a-input>
        </a-form-model-item>
        <a-form-model-item label="文章分类" prop="cid">
          <a-select v-model="artInfo.cid" style="width: 200px;" placeholder="请选择分类" @change="cateChange">
            <a-select-option v-for="item in cateList" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="文章描述" prop="desc">
          <a-input type="textarea" v-model="artInfo.desc"></a-input>
        </a-form-model-item>
        <a-form-model-item label="文章缩略图" prop="img">
          <a-upload
            :multiple="false"
            :action="upUrl"
            :headers="headers"
            list-type="picture"
            :defaultFileList="fileList"
            @change="upChange"
          >
            <a-button>
              <a-icon type="upload"></a-icon>点击上传
            </a-button>
          </a-upload>
          <br />
          <template v-if="id && artInfo.img">
            <img :src="artInfo.img" alt="缩略图" style="width: 200px; height: 180px; display: inline-block; float: left">
          </template>
        </a-form-model-item>
        <a-form-model-item label="文章内容" prop="content">
          <Editor v-model="artInfo.content"></Editor>
        </a-form-model-item>
        <a-form-model-item>
          <a-button type="danger" style="margin-right: 20px;" @click="artOk(artInfo.id)">{{ artInfo.id ? '更新' : '提交' }}</a-button>
          <a-button type="primary" @click="cancel">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>
<script>
import { Url } from '../../plugin/http'
import Editor from '../editor/index'
export default {
  components: { Editor },
  props: ['id'],
  data () {
    return {
      artInfo: {
        id: 0,
        title: '',
        cid: undefined,
        desc: '',
        content: '',
        img: ''
      },
      cateList: [],
      upUrl: Url + '/upload',
      headers: {},
      fileList: [],
      artInfoRules: {
        title: [{ required: true, message: '请输入文章标题', trigger: 'blur' }],
        cid: [{ required: true, message: '请输入文章分类', trigger: 'change' }],
        desc: [{ required: true, message: '请输入文章描述', trigger: 'blur' }, { max: 120, message: '描述最多120个字符', trigger: 'change' }],
        img: [{ required: true, message: '请选择文章缩略图', trigger: 'blur' }],
        content: [{ required: true, message: '请输入文章内容', trigger: 'blur' }]
      }
    }
  },
  created () {
    this.getCateList()
    this.headers = {
      Authorization: `Bearer ${window.sessionStorage.getItem('token')}`
    }
    if (this.id) {
      this.getArtInfo(this.id)
    }
  },
  methods: {
    async getArtInfo (id) {
      const { data: res } = await this.$http.get(`article/${id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.artInfo = res.data
      this.artInfo.id = res.data.ID
    },
    async getCateList () {
      const { data: res } = await this.$http.get('categories')
      if (res.status !== 200) return this.$message.error(res.message)
      this.cateList = res.data
    },
    // 选择分类
    cateChange (id) {
      this.artInfo.cid = id
    },
    // 上传文件
    upChange (info) {
      if (info.file.status !== 'uploading') {
      }
      if (info.file.status === 'done') {
        this.$message.success(`${info.file.name} 上传成功`)
        this.artInfo.img = info.file.response.url
      } else if (info.file.status === 'error') {
        this.$message.error(`${info.file.name} 上传失败`)
      }
    },
    // 提交或更新文章
    artOk (id) {
      this.$refs.artInfoRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('参数不符合要求')
        }
        if (id === 0) {
          // 新增文章
          const { data: res } = await this.$http.post('article/add', this.artInfo)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('添加文章成功')
          await this.$router.push('/artlist')
        } else {
          const { data: res } = await this.$http.put(`article/${id}`, this.artInfo)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('更新文章成功')
          await this.$router.push('/artlist')
        }
      })
    },
    cancel () {
      this.$refs.artInfoRef.resetFields()
    }
  }
}
</script>
