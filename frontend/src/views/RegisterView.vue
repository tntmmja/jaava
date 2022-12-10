<script setup>
import axios from 'axios'

let x = 0
function registerButtonHandler(event) {

    event.preventDefault()
    x++
    console.log("you clicked button ", x, " times")


    axios.get("http://localhost:8000/").then((response) => {
        console.log(response.data)
    })
}  


</script>

<template>
  <main>
this is a register page
    <form @submit.prevent="registerButtonHandler">
        <button type="submit">
            Register User
        </button>
    </form>
  </main>
</template>
