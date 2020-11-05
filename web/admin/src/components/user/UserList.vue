<template>
  <div>
    <h3>用户列表</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search placeholder="输入用户名查找" enter-button v-model="queryUserName" @search="getUserList" allowClear></a-input-search>
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="addUser">新增</a-button>
        </a-col>
      </a-row>
      <a-table
        rowKey="username"
        :columns="columns"
        :pagination="paginationOption"
        :dataSource="userlist"
        bordered
      >
        <span slot="role" slot-scope="role">{{ role === 1 ? '管理者' : '订阅者' }}</span>
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" style="margin-right: 15px" @click="editUser(data.ID)">编辑</a-button>
            <a-button type="danger" icon="delete" style="margin-right: 15px" @click="deleteUser(data.ID, data.role)">删除</a-button>
            <a-button type="info" icon="info" @click="resetPassword(data.ID)">密码重置</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
    <!--  新增用户区域  -->
    <a-modal
      title="新增用户"
      closable
      :visible="addUserVisible"
      @ok="addUserHandleOk"
      @cancel="addUserHandleCancel"
    >
      <a-form-model :model="newUserInfo" :rules="newUserRules" ref="addUserRef">
        <a-form-model-item label="用户名" prop="username">
          <a-input v-model="newUserInfo.username"></a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback label="密码" prop="password">
          <a-input-password v-model="newUserInfo.password"></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkPassword">
          <a-input-password v-model="newUserInfo.checkPassword"></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!--  编辑用户区域  -->
    <a-modal
      title="编辑用户"
      closable
      :visible="editUserVisible"
      @ok="editUserHandleOk"
      @cancel="editUserHandleCancel"
    >
      <a-form-model :model="userInfo" :rules="userRules" ref="editUserRef">
        <a-form-model-item label="用户名" prop="username">
          <a-input v-model="userInfo.username"></a-input>
        </a-form-model-item>
        <a-form-model-item label="是否为管理员" prop="role">
          <a-switch :checked="userInfo.role === 1" @change="adminChange"></a-switch>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!--  重置密码区域  -->
    <a-modal
      title="重置密码"
      closable
      :visible="resetPasswordVisible"
      @ok="resetPasswordHandleOk"
      @cancel="resetPasswordHandleCancel"
    >
      <a-form-model :model="resetPasswordInfo" :rules="resetPasswordRules" ref="resetPasswordRef">
        <a-form-model-item has-feedback label="新密码" prop="password">
          <a-input-password v-model="resetPasswordInfo.password"></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkPassword">
          <a-input-password v-model="resetPasswordInfo.checkPassword"></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id'
  },
  {
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    key: 'username'
  },
  {
    title: '角色',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    scopedSlots: { customRender: 'role' }
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
      queryUserName: '',
      paginationOption: {
        pageSizeOptions: ['5', '10', '20'],
        defaultCurrent: 1,
        defaultPageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
        onChange: (page, pageSize) => {
          this.getUserList(pageSize, page)
          this.paginationOption.defaultCurrent = page
          this.paginationOption.defaultPageSize = pageSize
        },
        onShowSizeChange: (current, size) => {
          this.getUserList(size, current)
          this.paginationOption.defaultCurrent = current
          this.paginationOption.defaultPageSize = size
        }
      },
      userlist: [],
      columns,
      // add user
      addUserVisible: false,
      newUserInfo: {
        id: 0,
        username: '',
        password: '',
        checkPassword: '',
        role: 2
      },
      newUserRules: {
        username: {
          validator: (rule, value, callback) => {
            if (this.newUserInfo.username === '') {
              callback(new Error('请输入用户名'))
            }
            if (this.newUserInfo.username.length < 4 || this.newUserInfo.username.length > 12) {
              callback(new Error('用户名应当在4到12之间'))
            } else {
              callback()
            }
          },
          trigger: 'blur'
        },
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.newUserInfo.password === '') {
                callback(new Error('请输入密码'))
              }
              if (this.newUserInfo.password.length < 6 || this.newUserInfo.password.length > 20) {
                callback(new Error('密码应当在6到20之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkPassword: [
          {
            validator: (rule, value, callback) => {
              if (this.newUserInfo.password === '') {
                callback(new Error('请输入确认密码'))
              }
              if (this.newUserInfo.password !== this.newUserInfo.checkPassword) {
                callback(new Error('两次输入密码不一致'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      // edit user
      editUserVisible: false,
      userInfo: {
        id: 0,
        username: '',
        role: 2
      },
      userRules: {
        username: {
          validator: (rule, value, callback) => {
            if (this.userInfo.username === '') {
              callback(new Error('请输入用户名'))
            }
            if (this.userInfo.username.length < 4 || this.userInfo.username.length > 12) {
              callback(new Error('用户名应当在4到12之间'))
            } else {
              callback()
            }
          },
          trigger: 'blur'
        }
      },
      resetPasswordVisible: false,
      resetPasswordInfo: {
        id: 0,
        password: '',
        checkPassword: ''
      },
      resetPasswordRules: {
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.resetPasswordInfo.password === '') {
                callback(new Error('请输入密码'))
              }
              if (this.resetPasswordInfo.password.length < 6 || this.resetPasswordInfo.password.length > 20) {
                callback(new Error('密码应当在6到20之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkPassword: [
          {
            validator: (rule, value, callback) => {
              if (this.resetPasswordInfo.password === '') {
                callback(new Error('请输入确认密码'))
              }
              if (this.resetPasswordInfo.password !== this.resetPasswordInfo.checkPassword) {
                callback(new Error('两次输入密码不一致'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      }
    }
  },
  created () {
    this.getUserList(this.queryUserName, this.paginationOption.defaultPageSize, this.paginationOption.defaultCurrent)
  },
  methods: {
    async getUserList (pageSize, pageNum) {
      const { data: res } = await this.$http.get('users', {
        params: {
          username: this.queryUserName,
          pagesize: pageSize,
          pagenum: pageNum
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.userlist = res.data
      this.paginationOption.total = res.total
    },
    // 删除用户
    async deleteUser (id, role) {
      this.$confirm({
        title: '确定要删除该用户码？',
        content: '一旦删除，无法恢复',
        onOk: async () => {
          if (role === 1) {
            this.$message.error('无删除权限')
            return
          }
          const res = await this.$http.delete(`user/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          await this.getUserList()
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        }
      })
    },
    // 新增用户
    addUser () {
      this.addUserVisible = true
    },
    addUserHandleOk () {
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('参数不符合要求')
        }
        const { data: res } = await this.$http.post('user/add',
          {
            username: this.userInfo.username,
            password: this.userInfo.password,
            role: this.userInfo.role
          }
        )
        if (res.status !== 200) return this.$message.error(res.message)
        this.addUserVisible = false
        this.$message.success('添加用户成功')
        await this.getUserList()
      })
    },
    addUserHandleCancel () {
      this.$refs.addUserRef.resetFields()
      this.addUserVisible = false
    },
    adminChange (check) {
      if (check) {
        this.userInfo.role = 1
      } else {
        this.userInfo.role = 2
      }
    },
    // 编辑用户
    async editUser (id) {
      this.editUserVisible = true
      const { data: res } = await this.$http.get(`user/${id}`)
      this.userInfo = res.data
      this.userInfo.id = id
    },
    editUserHandleOk () {
      this.$refs.editUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求')
        const { data: res } = await this.$http.put(`user/${this.userInfo.id}`, {
          username: this.userInfo.username,
          role: this.userInfo.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.editUserVisible = false
        this.$message.success('更新用户成功')
        await this.getUserList()
      })
    },
    editUserHandleCancel () {
      this.$refs.editUserRef.resetFields()
      this.editUserVisible = false
    },
    // 重置密码
    resetPassword (id) {
      this.resetPasswordInfo.id = id
      this.resetPasswordVisible = true
    },
    resetPasswordHandleOk () {
      this.$refs.resetPasswordRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求')
        const { data: res } = await this.$http.put(`reset/${this.resetPasswordInfo.id}`, {
          password: this.resetPasswordInfo.password
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.resetPasswordVisible = false
        this.$message.success('更新密码成功')
        await this.getUserList()
      })
    },
    resetPasswordHandleCancel () {
      this.$refs.resetPasswordRef.resetFields()
      this.resetPasswordVisible = false
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
