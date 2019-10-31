<template lang="pug">
  .container 
    .row.left left
      el-button(@click="join") join
    .row.middle 
      .item(v-for="(msg,i) in msgBox" :key="i") {{msg}}
      el-input(@submit="handleMessageSubmit")
    .row.right 
      .person(:key="item.id" v-for="item in onlineList") {{item.name}}

</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import io from 'socket.io-client'
@Component({})
export default class MyComponent extends Vue {
  public msgBox: string[] = [];
  public socket: io.socket; 
  public onlineList: object[] = [
    { name: "peng", id: "1" },
    { name: "heng", id: "2" },
    { name: "ling", id: "3" },
    { name: "bing", id: "4" }
  ];
  public join() {
    this.socket.emit('join', {data: 'I\'m connected!'});
  }
  public handleMessageSubmit(val) {
    this.socket.emit('message', val);
  }
  mounted() {
    this.socket = io.connect("http://" + document.domain + ":" + 5000);
    this.socket.on("connect", (data) => {
      console.log("connect")
    });
    this.socket.on("message", msg => {
      this.msgBox.push(msg)
      // this.$message({
      //   type: "success",
      //   showClose: true,
      //   message: msg
      // });
    });
  }
}
</script>

<style lang="scss" scoped>
.container {
  color: red;
  display: flex;
  flex-direction: row;
  height: 100%;
  .row {
    flex: 1;
  }
}
</style>