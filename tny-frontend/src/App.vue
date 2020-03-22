<template>
  <div id="app">
    <navBar />
    <img alt="Vue logo" src="./assets/logo.png" />
    <router-view></router-view>
    <Footer />
  </div>
</template>

<script>
import Footer from "./components/Footer";
import navBar from "./components/nav";
export default {
  name: "App",
  components: {
    Footer,
    navBar
  }
};
</script>

<style>
navBar {
  margin-bottom: 60px;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
  /* padding-top: 60px; */
}

html,
body {
  margin: 0;
  height: 100%;
}
</style>
