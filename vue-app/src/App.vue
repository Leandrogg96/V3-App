<template>
  <Header />
  <div>
    <router-view/>
  </div>
  <Footer />
</template>

<script>
import Header from "./components/App-Header";
import Footer from "./components/App-Footer";
import { store } from './components/store.js';

export default {
  name: 'App',
  components: {
    Header,
    Footer,
  },
  data() {
    return {
      store
    }
  }
}
</script>

<style>
 
</style>
