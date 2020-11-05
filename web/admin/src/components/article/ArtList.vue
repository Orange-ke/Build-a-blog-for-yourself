<template>
  <div>
    <h3>文章列表</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search placeholder="输入查找" enter-button v-model="queryArtName" @search="getArtList" allowClear></a-input-search>
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="addArt">新增</a-button>
        </a-col>
        <a-col :span="6">
          <a-select allowClear :value="defaultValue" style="width: 120px" @change="showCateArticles">
            <a-select-option v-for="item in cateList" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
          </a-select>
        </a-col>
      </a-row>
      <a-table
        rowKey="ID"
        :columns="columns"
        :pagination="paginationOption"
        :dataSource="articleList"
        bordered
      >
        <span slot="img" slot-scope="img">
          <img :src="img" alt="缩略图" style="width: 100px; height: 80px;">
        </span>
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button size="small" type="primary" icon="edit" style="margin-right: 15px" @click="editArt(data.ID)">编辑</a-button>
            <a-button size="small" type="danger" icon="delete" style="margin-right: 15px" @click="deleteArt(data.ID)">删除</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '5%',
    key: 'id'
  },
  {
    title: '文章分类',
    dataIndex: 'Category.name',
    width: '10%',
    key: 'cid'
  },
  {
    title: '文章标题',
    dataIndex: 'title',
    width: '20%',
    key: 'title'
  },
  {
    title: '描述',
    dataIndex: 'desc',
    width: '20%',
    key: 'desc'
  },
  {
    title: '缩略图',
    dataIndex: 'img',
    width: '20%',
    key: 'img',
    align: 'center',
    scopedSlots: { customRender: 'img' }
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    scopedSlots: { customRender: 'action' }
  }
]
export default {
  data () {
    return {
      defaultValue: '选择分类',
      cid: 0,
      queryArtName: '',
      paginationOption: {
        pageSizeOptions: ['5', '10', '20'],
        defaultCurrent: 1,
        defaultPageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
        onChange: (page, pageSize) => {
          this.getArtList(pageSize, page)
          this.paginationOption.defaultCurrent = page
          this.paginationOption.defaultPageSize = pageSize
        },
        onShowSizeChange: (current, size) => {
          this.getArtList(size, current)
          this.paginationOption.defaultCurrent = current
          this.paginationOption.defaultPageSize = size
        }
      },
      cateList: [],
      articleList: [],
      columns
    }
  },
  created () {
    this.getArtList(this.queryArtName, this.cid, this.paginationOption.defaultPageSize, this.paginationOption.defaultCurrent)
    this.getCateList()
  },
  methods: {
    showCateArticles (value) {
      if (value === undefined) {
        this.defaultValue = '选择分类'
        this.cid = 0
      } else {
        this.defaultValue = value
        this.cid = value
      }
      this.getArtList(this.queryArtName, this.cid, this.paginationOption.defaultPageSize, this.paginationOption.defaultCurrent)
    },
    async getCateList (name, pageSize, pageNum) {
      const { data: res } = await this.$http.get('categories', {
        params: {
          name: this.queryCateName,
          pagesize: pageSize,
          pagenum: pageNum
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.cateList = res.data
    },
    async getArtList (pageSize, pageNum) {
      const { data: res } = await this.$http.get('articles', {
        params: {
          title: this.queryArtName,
          cid: this.cid,
          pagesize: pageSize,
          pagenum: pageNum
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.articleList = res.data
      this.paginationOption.total = res.total
    },
    // 删除文章
    async deleteArt (id) {
      this.$confirm({
        title: '确定要删除该文章码？',
        content: '一旦删除，无法恢复',
        onOk: async () => {
          const res = await this.$http.delete(`article/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          await this.getArtList()
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        }
      })
    },
    // 新增文章
    addArt () {
      this.$router.push('/addart')
    },
    // 编辑用户
    editArt (id) {
      this.$router.push(`/addart/${id}`)
    }
  }
}
</script>

<style scoped>
  .actionSlot {
    display: flex;
    justify-content: center;
  }
</style>
