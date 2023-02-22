<template>
  <div>
    <h3>Logged in User</h3>
    <div>
      <label>Nickname:</label>
      <span>{{ user.nickname }}</span>
    </div>
    <div>
      <label>Email:</label>
      <span>{{ user.email }}</span>
    </div>
    <button @click="handleLogout">Logout</button>
  </div>
</template>
<script>
import axios from 'axios'
export default {
  data() {
    return {
      user: null,
    };
  },
  created() {
    const sessionID = localStorage.getItem("sessionID");
    if (!sessionID) {
      this.$router.push({ name: 'login' });
    } else {
      // Make an API call to fetch the user data using the session ID
      axios.get(`http://localhost:8082/login/${sessionID}`)
        .then(response => {
          this.user = response.data;
        })
        .catch(error => {
          console.error(error);
        });
    }
  },
  methods: {
    handleLogout() {
      localStorage.removeItem("sessionID");
      this.$router.push({ name: 'login' });
    },
  },
};
</script>
