<template>
  <div id="nav">
    <router-link to="/">Home</router-link> |
    <router-link to="/submit">Submit Form</router-link> |
    <router-link to="/messages">Messages History</router-link>
  </div>
  <router-view v-slot="{ Component }">
    <component :is="Component" :key="$route.path"></component>
  </router-view>
</template>


<script setup>
import { Suspense } from 'vue';
</script >

<script>
export default {
  name: 'App',
  components: {
    Suspense,
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
