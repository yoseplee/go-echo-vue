<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
    askdfjkajsdlkfj
    <br>
    {{ helloText }}
    <br>
    {{ test }}
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
// import Api from '@/middleware/Api.js'
const axios = require('axios')

export default {
  name: 'home',
  components: {
    HelloWorld
  },
  async created() {
    this.helloText = await this.getTest()
  },
  data () {
    return {
      helloText: "",
      test: "This is test"
    }
  },
  methods: {
    async getTest() {
      return await axios.get('http://localhost:1323/api')
    }
  }
}
</script>
