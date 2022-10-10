<template>
  <div>

    <div id="visitInfo" :style="{width: '900px', height: '300px'}"></div>

    <h3  :style="{'color':'red','width': '900px', 'height': '300px','font-size': '20px'}">用户总访问量为：{{visitCount}}</h3>


  </div>

</template>

<script>

export default {

  name: 'hello',

  data () {
    return {

      visitCount: 0,

      dates:[],
      counts:[],

      // 用户阅读排名，用户信息和阅读信息
      userReadCount:[],
      userName:[],
      userInfos:[],

    }
  },
  mounted(){
    this.getVisitInfos();
    this.getReadDetailsInfos();
    this.getTotalCount();
  },
  methods: {

    // 1：查询用户访问信息
    async getVisitInfos() {

      // (1) 发送http get请求
      const { data: res } = await this.$http.get(`admin/data/infos`)
      if (res.status >= 400) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$message.error(res.message)
          this.$router.push('/login')
          return
        }
        this.$message.error(res.message)
        return
      }

      // 数组遍历和组装
      res.forEach((item,index) => {
        this.dates.push(item.date)
        this.counts.push(item.count)
      })

      // 基于刚刚准备好的 DOM 容器，初始化 EChart 实例
      let myChart = this.$echarts.init(document.getElementById('visitInfo'))

      // 绘制图表
      myChart.setOption({
        title: { text: '一月用户访问情况' },
        tooltip: {},
        xAxis: {
          data:this.dates,
        },
        yAxis: {},
        series: [{
          name: '数量',
          type: 'bar',
          data: this.counts
        }]
      });
    },

    // 2：展示访问最多的ip数和用户的关系
    async getReadDetailsInfos() {

      let myChart = this.$echarts.init(document.getElementById('userInfo'))

      // 绘制图表
      myChart.setOption({
        title: { text: '用户访问信息表' },
        tooltip: {},
        xAxis: {
          data:["111"],
        },
        yAxis: {},
        series: [{
          name: '数量',
          type: 'bar',
          data: [222]
        }]
      })
    },

    // 3：查询站点总访问信息
    async getTotalCount() {

      // (1) 发送http get请求
      const { data: res } = await this.$http.get(`admin/data/total`)
      if (res.status >= 400) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$message.error(res.message)
          this.$router.push('/login')
          return
        }
        this.$message.error(res.message)
        return
      }

      this.visitCount = res.total
    },
  }
}
</script>
