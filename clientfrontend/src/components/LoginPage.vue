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
    <router-link to="register" class="nav-link">Sign up</router-link>
  </div>
</template>
<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      password: '',
      error: false
    };
  },
  methods: {
    async handleSubmit() {
      try {
        const formData = {
          username: this.username,
          password: this.password
        };
        console.log(JSON.stringify(formData))
        const response = await axios.post('http://localhost:8082/login', JSON.stringify(formData));

        if (response.data.success) {
          // handle success response
          this.handleSuccessResponse(response);
        } else {
          this.error = true;
        }
      } catch (error) {
        // handle the error
      }
    },
    handleSuccessResponse(response) {
      // handle success response
      // Store the user data in the local storage
      localStorage.setItem("sessionID", response.data.sessionID);
      // e.g. redirect the user to the dashboard or home page
     
      this.$router.push({ name: '/loggedUser' });
    },
  },
};
</script>