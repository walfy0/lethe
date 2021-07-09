<template>
<html>
<body id="login_div">
  <div id="div">
  <el-form>
  <el-form-item id="user">
    <el-input v-model="user" placeholder="user" id="user"></el-input>
  </el-form-item>
  <el-form-item id="password">
    <el-input v-model="password" placeholder="password"></el-input>
  </el-form-item>
  <el-form-item>
    <el-button type="success" @click="Register">register</el-button>
    <el-button type="primary" @click="Login">login</el-button>
  </el-form-item>
  </el-form>
  </div>
</body>
</html>
</template>

<script>
export default {
  name: 'login',
  data () {
    return {
      user: '',
      password: ''
    }
  },
  methods: {
    Login: function () {
      this.axios.post('/lethe/login', {
        name: this.user,
        password: this.password
      })
        .then((response) => {
          if (response.data.code === '0') {
            this.$router.push({
              path: '/home',
              query: {
                user: this.user
              }
            })
          } else {
            console.log(response)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    Register: function () {
      this.$router.push({
        path: '/register'
      })
    }
  }
}
</script>

<style scoped>
#login_div{
  text-align: center;
  background-image: url('../assets/login.jpg');
  background-size: cover;
  position: absolute;
  width: 100%;
  height: 100%;
  margin: 0px;
}
#div {
  text-align: center;
  position: absolute;
  width: 90%;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
#user{
    margin: auto;
    width: 50%;
    border: 3px;
    padding: 20px;
}
#password{
  margin: auto;
    width: 50%;
    border: 3px;
    padding: 30px;
}
#el-button {
  margin: auto;
  width: 80%;
}
</style>
