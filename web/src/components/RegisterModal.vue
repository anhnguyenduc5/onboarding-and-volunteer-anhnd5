<template>
    <div class="modal" v-if="show">
      <div class="modal-content">
        <span class="close" @click="close">&times;</span>
        <h2>Register</h2>
        <form @submit.prevent="register">
          <div>
            <label for="email">Email:</label>
            <input type="email" id="email" v-model="email" required />
          </div>
          <div>
            <label for="password">Password:</label>
            <input type="password" id="password" v-model="password" required />
          </div>
          <div>
            <label for="confirmPassword">Confirm Password:</label>
            <input type="password" id="confirmPassword" v-model="confirmPassword" required />
          </div>
          <button type="submit">Register</button>
        </form>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  
  const props = defineProps({
    show: Boolean,
  })
  
  const emit = defineEmits(['close', 'register'])
  
  const email = ref('')
  const password = ref('')
  const confirmPassword = ref('')
  
  const close = () => {
    emit('close')
  }
  
  const register = () => {
    if (password.value !== confirmPassword.value) {
      alert('Passwords do not match')
      return
    }
    // Add register logic here
    console.log('Registering with', email.value, password.value)
    emit('register', { email: email.value, password: password.value })
  }
  </script>
  
  <style scoped>
  .modal {
    display: flex;
    justify-content: center;
    align-items: center;
    position: fixed;
    z-index: 1;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: auto;
    background-color: rgb(0,0,0);
    background-color: rgba(0,0,0,0.4);
    color: black;
  }
  
  .modal-content {
    background-color: #fefefe;
    padding: 20px;
    border: 1px solid #888;
    width: 300px;
    color: black;
  }
  
  .close {
    color: #aaa;
    float: right;
    font-size: 28px;
    font-weight: bold;
  }
  
  .close:hover,
  .close:focus {
    color: black;
    text-decoration: none;
    cursor: pointer;
  }
  </style>
  