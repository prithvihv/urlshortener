<template>
  <div class="hello">
    <input v-model="url" type="text" name="fullURL" id="fullURL" />
    <br />
    <div  :key="key" v-for="(value, key) in tnyStore">
      <div>
        {{ key }}
        <br>
        <a :href="getAppServerURL+value">{{ getAppServerURL+value}}</a>
      </div>
    </div>

    <br />
    <button @click="onSubmit">Get Small url</button>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  data() {
    return {
      url: "https://gobyexample.com/random-numbers"
    };
  },
  computed: {
    ...mapGetters(["tnyStore", "getAppServerURL"])
  },
  methods: {
    async onSubmit(evt) {
      console.log("on click triggered");
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
