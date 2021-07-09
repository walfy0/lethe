<template>
  <el-container>
  <el-header>
    <div id="lethe">lethe</div>
    <el-dropdown>
      <span class="el-dropdown-link">
      {{ user }}<i class="el-icon-arrow-down el-icon--right"></i>
      </span>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item>info</el-dropdown-item>
        <el-dropdown-item @click.native="logout">logout</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
  </el-header>
  <el-divider></el-divider>
  <el-container>
    <el-aside>
      <el-menu
      default-active="2"
      class="el-menu-vertical-demo"
      @open="handleOpen"
      @close="handleClose">
      <el-submenu index="1">
        <template slot="title">
          <i class="el-icon-location"></i>
          <span>导航一</span>
        </template>
        <el-menu-item-group>
          <template slot="title">分组一</template>
          <el-menu-item index="1-1">选项1</el-menu-item>
          <el-menu-item index="1-2">选项2</el-menu-item>
        </el-menu-item-group>
        <el-menu-item-group title="分组2">
          <el-menu-item index="1-3">选项3</el-menu-item>
        </el-menu-item-group>
        <el-submenu index="1-4">
          <template slot="title">选项4</template>
          <el-menu-item index="1-4-1">选项1</el-menu-item>
        </el-submenu>
      </el-submenu>
      <el-menu-item index="2" @click.native="blog">
        <i class="el-icon-menu"></i>
        <span slot="title">blog</span>
      </el-menu-item>
      <el-menu-item index="3" disabled>
        <i class="el-icon-document"></i>
        <span slot="title">导航三</span>
      </el-menu-item>
      <el-menu-item index="4">
        <i class="el-icon-setting"></i>
        <span slot="title">导航四</span>
      </el-menu-item>
    </el-menu>
    </el-aside>
    <el-main>Main</el-main>
  </el-container>
</el-container>
</template>

<script>
export default {
  name: 'Home',
  data () {
    return {
      user: 'welcome',
      password: ''
    }
  },
  mounted: function () {
    this.user = this.$route.query.user
  },
  methods: {
    logout () {
      this.axios.get('/lethe/logout').then().catch(error => {
        console.log(error)
      })
      this.$router.push({
        path: 'login'
      })
    },
    blog () {
      this.$router.push({
        path: 'blog',
        query: {
          user: this.user
        }
      })
    }
  }
}
</script>

<style scope>
.el-header, .el-footer {
  color: #333;
  text-align: center;
  line-height: 70px;
}

.el-aside, .el-menu {
  color: #333;
  text-align: center;
  line-height: 800px;
  overflow-y: auto;
}

.el-main {
  color: #333;
  text-align: center;
  line-height: 110px;
  overflow-y: auto;
}

#lethe {
  float: left;
  font-size: 28px;
  color:rgb(69, 130, 243)
}

#header_left {
  margin: auto;
  border-radius: 4px;
}

.el-col{
  width: 10px;
  height: 10px;
}
.el-dropdown, .el-dropdown-menu, .el-dropdown-menu {
  float: right;
  font-size: 18px;
}
</style>
