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
  <el-form-item id="repeat">
    <el-input v-model="repeat" placeholdr="repeat"></el-input>
  </el-form-item>
  <el-form-item id="email">
    <el-input v-model="email" placeholder="email"></el-input>
  </el-form-item>
  <el-form-item id="code">
    <el-input v-model="code" placeholder="code"></el-input>
  </el-form-item>
  <el-form-item>
    <el-button type="primary" @click="Register">register</el-button>
  </el-form-item>
  </el-form>
  </div>
</body>
</html>
</template>

<script>
export default {
  name: 'register',
  data () {
    return {
      user: '',
      password: '',
      repeat: '',
      email: '',
      code: ''
    }
  },
  methods: {
    Register: function () {
      if (this.password !== this.repeat) {
        this.$message.error('password not same!')
      }
      this.axios.post('/lethe/register', {
        name: this.user,
        password: this.password,
        email: this.email,
        code: this.code
      })
        .then((response) => {
          if (response.data.code === '0') {
            this.$router.push({
              path: '/login',
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
    SendEmail: function () {
      this.axios.post('/lethe/send_mail', {
        email: this.email
      })
        .then((response) => {
          if (response.data.code === '0') {
            this.$message({
              message: 'mail send success',
              type: 'success'
            })
          } else {
            console.log(response)
          }
        }).catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>

<style scoped>
head{
  margin: 0px;
}
#login_div{
  text-align: center;
  background-image: url('../assets/login.jpg');
  position: absolute;
  background-size: cover;
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
#password, #email, #code, #repeat{
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
