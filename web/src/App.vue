<script setup>
import { ref } from 'vue'
import LoginModal from './components/LoginModal.vue'
import RegisterModal from './components/RegisterModal.vue'

// Thông tin tĩnh về dự án
const projectInfo = {
  name: "Volunteer Management Project",
  description: "A basic system to manage volunteers.",
  version: "1.0.0",
  author: "Your Name",
  features: [
    "User registration and login",
    "Volunteer application management",
    "Admin panel for managing volunteers",
    "Email notifications for approvals and rejections"
  ]
}

// State to manage modals
const showLoginModal = ref(false)
const showRegisterModal = ref(false)

const handleLogin = () => {
  showLoginModal.value = true
}

const handleRegister = () => {
  showRegisterModal.value = true
}

const closeLoginModal = () => {
  showLoginModal.value = false
}

const closeRegisterModal = () => {
  showRegisterModal.value = false
}
</script>

<template>
  <header>
    <h1>{{ projectInfo.name }}</h1>
  </header>

  <main>
    <p>{{ projectInfo.description }}</p>
    <ul>
      <li><strong>Version:</strong> {{ projectInfo.version }}</li>
      <li><strong>Author:</strong> {{ projectInfo.author }}</li>
    </ul>
    <h3>Features:</h3>
    <ul>
      <li v-for="feature in projectInfo.features" :key="feature">{{ feature }}</li>
    </ul>

    <div class="buttons">
      <button @click="handleLogin">Login</button>
      <button @click="handleRegister">Register</button>
    </div>

    <!-- Login Modal -->
    <LoginModal :show="showLoginModal" @close="closeLoginModal" />

    <!-- Register Modal -->
    <RegisterModal :show="showRegisterModal" @close="closeRegisterModal" />
  </main>
</template>

<style scoped>
header {
  text-align: center;
  margin: 2rem 0;
}

main {
  text-align: center;
}

ul {
  list-style: none;
  padding: 0;
}

.buttons {
  margin-top: 2rem;
}

button {
  margin: 0 1rem;
  padding: 0.5rem 1rem;
  font-size: 1rem;
  cursor: pointer;
}
</style>
