<template>
    <div class="modal" v-if="show">
      <div class="modal-content">
        <span class="close" @click="close">&times;</span>
        <h2>Login</h2>
        <form @submit.prevent="login">
          <div>
            <label for="email">Email:</label>
            <input type="email" id="email" v-model="email" required />
          </div>
          <div>
            <label for="password">Password:</label>
            <input type="password" id="password" v-model="password" required />
          </div>
          <button type="submit">Login</button>
        </form>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  
  const props = defineProps({
    show: Boolean,
  })
  
  const emit = defineEmits(['close', 'login'])
  
  const email = ref('')
  const password = ref('')
  
  const close = () => {
    emit('close')
  }
  
  const login = () => {
    // Add login logic here
    console.log('Logging in with', email.value, password.value)
    emit('login', { email: email.value, password: password.value })
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
  