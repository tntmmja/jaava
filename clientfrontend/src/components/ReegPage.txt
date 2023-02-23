<template>

    <div class="index">

        <h1>{{ msg }}</h1>

    </div>
	<div class="auth-wrapper">
		<div class="auth-inner">
			<form @submit.prevent="handleSubmit">
				<h3>Register</h3>	

				<span class="error badge bg-secondary" v-if="error">
					Incorrect Nickname/Email or Password
				</span>
			
				<div class="form-group">
					<label>Nickname or Email</label>
					<input class="form-control"
					v-model="nickname"
					autocomplete="username"
					placeholder="Nickname or Email"/>
				</div>

				<div class="form-group">
					<label>Password</label>
					<input class="form-control"
					type="password"
					v-model="password"
					autocomplete="password"
					placeholder="Password"/>
				</div>

				<button class="btn btn-primary btn-block">Sign up</button>
			</form>	
			
		</div>
	</div>	
</template>

<script>
	export default {
        // component data, methods, etc.
		name: 'RegisterPage',
        props: {
            msg: String
        },
		data() {
			return {
				nickname: '',
				password: '',
				error: false,
			}
		},
		methods: {
			handleSubmit() {
				this.$store.dispatch('logInUser', {
					nickname: this.nickname, 
					password: this.password
				})
				.then(() => {
					this.$router.push({name: 'feed'})
				})
				.catch(() => {
					this.error = true
				})
			},
		},
	}
</script>


<style>
  /* component styles */
</style>