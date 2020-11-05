<template>
  <div>
    <h3>分类列表</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search placeholder="输入分类名查找" enter-button v-model="queryCateName" @search="getCateList" allowClear></a-input-search>
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="addCate">新增</a-button>
        </a-col>
      </a-row>
      <a-table
        rowKey="id"
        :columns="columns"
        :pagination="paginationOption"
        :dataSource="cateList"
        bordered
      >
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" style="margin-right: 15px" @click="editCate(data.id)">编辑</a-button>
            <a-button type="danger" icon="delete" style="margin-right: 15px" @click="deleteCate(data.id)">删除</a-button>
          </div>
        </template>
      </a-table>
      <!--  新增用户区域  -->
      <a-modal
        title="新增分类"
        closable
        :visible="addCateVisible"
        @ok="addCateHandleOk"
        @cancel="addCateHandleCancel"
      >
        <a-form-model :model="newCateInfo" :rules="newCateRules" ref="addCateRef">
          <a-form-model-item label="分类名" prop="name">
            <a-input v-model="newCateInfo.name"></a-input>
          </a-form-model-item>
        </a-form-model>
      </a-modal>
      <!--  编辑用户区域  -->
      <a-modal
        title="编辑分类"
        closable
        :visible="editCateVisible"
        @ok="editCateHandleOk"
        @cancel="editCateHandleCancel"
      >
        <a-form-model :model="cateInfo" :rules="cateRules" ref="editCateRef">
          <a-form-model-item label="分类名" prop="name">
            <a-input v-model="cateInfo.name"></a-input>
          </a-form-model-item>
        </a-form-model>
      </a-modal>
    </a-card>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    width: '10%',
    key: 'id'
  },
  {
    title: '分类名',
    dataIndex: 'name',
    width: '20%',
    key: 'cateName'
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
      queryCateName: '',
      paginationOption: {
        pageSizeOptions: ['5', '10', '20'],
        defaultCurrent: 1,
        defaultPageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
        onChange: (page, pageSize) => {
          this.getCateList(this.queryCateName, pageSize, page)
          this.paginationOption.defaultCurrent = page
          this.paginationOption.defaultPageSize = pageSize
        },
        onShowSizeChange: (current, size) => {
          this.getCateList(this.queryCateName, size, current)
          this.paginationOption.defaultCurrent = current
          this.paginationOption.defaultPageSize = size
        }
      },
      cateList: [],
      columns,
      addCateVisible: false,
      newCateInfo: {
        id: 0,
        name: ''
      },
      newCateRules: {
        name: {
          validator: (rule, value, callback) => {
            if (this.newCateInfo.name === '') {
              callback(new Error('请输入分类名'))
            }
            callback()
          },
          trigger: 'blur'
        }
      },
      // edit cate
      editCateVisible: false,
      cateInfo: {
        id: 0,
        name: ''
      },
      cateRules: {
        username: {
          validator: (rule, value, callback) => {
            if (this.cateInfo.name === '') {
              callback(new Error('请输入分类名'))
            }
            callback()
          },
          trigger: 'blur'
        }
      }
    }
  },
  created () {
    this.getCateList(this.queryCateName, this.paginationOption.defaultPageSize, this.paginationOption.defaultCurrent)
  },
  methods: {
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
      this.paginationOption.total = res.total
    },
    // 删除分类
    async deleteCate (id) {
      this.$confirm({
        title: '确定要删除该分类码？',
        content: '一旦删除，无法恢复',
        onOk: async () => {
          const res = await this.$http.delete(`category/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          await this.getCateList()
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        }
      })
    },
    // 新增用户
    addCate () {
      this.addCateVisible = true
    },
    addCateHandleOk () {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('参数不符合要求')
        }
        const { data: res } = await this.$http.post('category/add',
          {
            name: this.newCateInfo.name
          }
        )
        if (res.status !== 200) return this.$message.error(res.message)
        this.addCateVisible = false
        this.$message.success('添加分类成功')
        await this.getCateList()
      })
    },
    addCateHandleCancel () {
      this.$refs.addCateRef.resetFields()
      this.addCateVisible = false
    },
    // 编辑用户
    async editCate (id) {
      this.editCateVisible = true
      const { data: res } = await this.$http.get(`category/${id}`)
      this.cateInfo = res.data
      this.cateInfo.id = id
    },
    editCateHandleOk () {
      this.$refs.editCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求')
        const { data: res } = await this.$http.put(`category/${this.cateInfo.id}`, {
          name: this.cateInfo.name
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.editCateVisible = false
        this.$message.success('更新分类成功')
        await this.getCateList()
      })
    },
    editCateHandleCancel () {
      this.$refs.editCateRef.resetFields()
      this.editCateVisible = false
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
