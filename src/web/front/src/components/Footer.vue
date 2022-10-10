<template>
  <v-footer padless color="indigo darken-2">
    <v-row justify="center" no-gutters>

      <v-col class="py-4 text-center white--text" cols="12">{{ new Date().getFullYear() }} - myBlog</v-col>

      <v-col class="py-4 text-center white--text" cols="12">网站总访问量为：{{visitCount}}</v-col>

      <div class="text-center white--text">
        <a class="text-center white--text" href="https://beian.miit.gov.cn/">{{icp_record}}</a>
      </div>

    </v-row>
  </v-footer>
</template>

<script>
export default {
  data() {
    return {

      icp_record: '',

      visitCount: 0,
    }
  },
  mounted() {

    this.$root.$on('msg', (msg) => {
      this.icp_record = msg
    }),

        this.getTotalCount();
  },

  methods: {

    // 查询站点总访问信息
    async getTotalCount() {

      // (1) 发送http get请求
      const { data: res } = await this.$http.get('data/total')
      this.visitCount = res.total

    }
  }
}
</script>

<style>
</style>