// kui vue installida terminali siis tekib huk faile nagu rtforumis on sh main.js

import Vue from 'vue'   
import App from './App.vue' 

Vue.config.productionTip = false

Vue.use(MyPlugin)

new Vue({
    render: h => h(App),
}).$mount('#app')
