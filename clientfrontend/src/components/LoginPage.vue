<template>
  <div>
    <form @submit.prevent="handleSubmit">
      <h3>Login</h3>
      <span class="error badge bg-secondary" v-if="error">
        Incorrect Nickname/Email or Password
      </span>
      <div>
        <label for="username">Nickname/Email:</label>
        <input type="text" id="username" v-model="username" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" required />
      </div>
      <button type="submit">Login</button>
    </form>
    <router-link to="/register" class="nav-link">Sign up</router-link>
  </div>
</template>

<script>
import axios from 'axios';


export default {
  data() {
    return {
      username: '',
      password: '',
      error: false,      
    };
  },
  
  methods: {
    async handleSubmit() {
      try {
        const formData = {
          username: this.username,
          password: this.password
        };
        console.log('joyson', JSON.stringify(formData))
        const response = await axios.post('http://localhost:8082/login', JSON.stringify(formData));
        
        if (response.data.success) {          
            this.handleSuccessResponse(response);
            console.log('joyson2', JSON.stringify(formData))
          
        } else {
          this.error = true;
        }
      } catch (error) {
        // Handle the error
      }
    },
    handleSuccessResponse(response) {
      localStorage.setItem('sessionID', response.data.sessionID);
      console.log(response)
      console.log('joudsin siia')

      this.$router.push('/loggedUser' );
      console.log('joudsin siia peale routerit')
    }
  }
};
</script>
