<!-- Header.vue is to make different components too look unified
in main App.vue component, you include the header component in the template. -->







<template>
    <header>
      <nav>
        <div class="logout-container">
          <button @click="logout">Logout</button>
        </div>
      </nav>
    </header>
  </template>
  
  <script>
  export default {
    methods: {
      logout() {
        // Remove the token from local storage or cookie
        localStorage.removeItem('token');
        // Redirect the user to the login page
        this.$router.push({ name: 'login' });
      }
    }
  };
  </script>
  
  <style>
  header {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    height: 50px;
  }
  
  .logout-container {
    margin-right: 10px;
  }
  </style>
  