import Vue from "vue";
import App from "./components/App/App";

const app = new Vue(App).$mount("#test");

app.text = "Electron Forge with Vue.js!";
