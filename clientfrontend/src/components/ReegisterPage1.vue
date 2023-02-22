<!-- //see ei saada midagi,aga vues errorit ka pole -->

<template>
    <div>
      <form @submit.prevent="handleSubmit">
        <div>
          <label for="nickname">Nickname:</label>
          <input type="text" id="nickname" v-model="nickname" required />
        </div>
        <div>
          <label for="age">Age:</label>
          <input type="number" id="age" v-model="age" required />
        </div>
        <div>
          <label for="gender">Gender:</label>
          <input type="radio" id="gender" value="male" v-model="gender" required />Male
          <input type="radio" id="gender" value="female" v-model="gender" required />Female
          <input type="radio" id="gender" value="other" v-model="gender" required />Other
        </div>
        <div>
          <label for="firstName">First Name:</label>
          <input type="text" id="firstName" v-model="firstName" required />
        </div>
        <div>
          <label for="lastName">Last Name:</label>
          <input type="text" id="lastName" v-model="lastName" required />
        </div>
        <div>
          <label for="email">E-mail:</label>
          <input type="email" id="email" v-model="email" required />
        </div>
        <div>
          <label for="password">Password:</label>
          <input type="password" id="password" v-model="password" required />
        </div>
        <button type="submit">Register</button>
      </form>
    </div>
  </template>
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        nickname: '',
        age: '',
        gender: '',
        firstName: '',
        lastName: '',
        email: '',
        password: ''
      };
    },
    methods: {
      async handleSubmit() {
        try {
          const formData = {
            nickname: this.nickname,
            age: this.age,
            gender: this.gender,
            firstName: this.firstName,
            lastName: this.lastName,
            email: this.email,
            password: this.password
          };
  
          const response = await axios.post('/register', formData);
  
          if (response.data.success) {
            // handle success response
          } else {
            // handle error response
          }
        } catch (error) {
          // handle the error
        }
      },
    },
  };
  </script>