<!-- Header.vue -st tuleb siis template header component veel panna -->


<template>
  <div id="app">
    <!-- <Header /> -->
    <router-view />
  </div>
</template>

<script>
// import Header from "./components/Header.vue";

export default {
  name: "App",
  mounted () {
    this.socket = new WebSocket("ws://localhost:8082/socket")
  },
  // components: {
  //   Header,
  // },
};
</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>


<!-- selline oli ene header.vue lisamist
<template>
  <div id="app">
    <router-view />
  </div>
</template>
<script>


export default {
  name: 'App',
  
  mounted () {
    this.socket = new WebSocket("ws://localhost:8082/socket")
  }
}
</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
 -->



<!-- //see on tagavara enne websocketi kulgepookimist


<template>
  <div id="app">
    <router-view />
  </div>
</template>
<script>
export default {
  name: 'App',
}
</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>



-->


<!-- see ei toota

<template>
  <div id="app">
    <router-view />
  </div>
</template>
<script>
import { Ws, WsProvider } from '@/plugins/vue-gorilla-websocket.js'

export default {
  name: 'App',
  provide () {
    return {
      $ws: new Ws('ws://localhost:8000/ws')
    }
  },
  mounted () {
    this.$ws.connect()
  }
}
</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
-->