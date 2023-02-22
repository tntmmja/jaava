import { createApp } from 'vue'
import App from './App.vue'
//import Toast from 'vue-toastification';
//import "vue-toastification/dist/index.css";
//import Vuelidate from 'vuelidate'; // see oli mingi aeg valja kommenteeritud ja vuelidate core oli alles jaetud
/////////////import VueRouter from 'vue-router'
import { createRouter, createWebHistory} from 'vue-router'
import LoginPage from './components/LoginPage.vue'
import RegisterPage from './components/RegisterPage.vue'
//import { Vuelidate } from '@vuelidate/core'
import LoggedUser from './components/LoggedUser.vue'

const router = createRouter({ /////enne oli VueRouter
  routes: [
    { path: '/login', component: LoginPage },
    { path: '/register', component: RegisterPage },
    { path: '/loggedUser', component: LoggedUser }
  ],
  history: createWebHistory(),
})

// Toast options
//const options = {  };

createApp(App)
  .use(router)
 // .use(Vuelidate)
  //.use(Toast, options)
  .mount('#app')


/*
see oli ka mingi variant mida soovitati
import Vue from 'vue'
import VueSocketIO from 'vue-socket.io'

Vue.use(new VueSocketIO({
  debug: true,
  connection: 'ws://localhost:8000/ws'
}))*/