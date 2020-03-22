<template>
  <div class="hello">
    <input v-model="url" type="text" name="fullURL" id="fullURL" />
    <br />
    {{ shortURL }}
    <br />
    <button @submit="onSubmit">Get Small url</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      url: "please enter url here",
      shortURL: ""
    };
  },
  methods: {
    async onSubmit(evt) {
      evt.preventDefault();
      await this.$store.dispatch("fetchShotURL", this.url);
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
