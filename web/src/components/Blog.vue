<template>
  <el-container>
</el-container>
</template>

<script>
export default {
  name: 'Blog',
  data () {
    return {
      user: 'welcome',
      doc_list: ''
    }
  },
  mounted: function () {
    this.user = this.$route.query.user
    this.axios.post('/lethe/list', {
      params: {
        page: 1,
        page_size: 20
      }
    }).then(response => {
      if (response.data.code === '0') {
        this.doc_list = response.data.data
      } else {
        console.log(response)
      }
    }).catch(error => {
      console.log(error)
    })
  },
  methods: {
    logout () {
      this.axios.get('/lethe/logout').then().catch(error => {
        console.log(error)
      })
      this.$router.push({
        path: 'login'
      })
    }
  }
}
</script>

<style scope>

</style>
