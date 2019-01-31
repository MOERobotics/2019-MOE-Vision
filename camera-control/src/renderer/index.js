import Vue from "vue";
import Buefy from "buefy";

Vue.use(Buefy);
import App from "./components/App/App";

const app = new Vue(App).$mount("#VueApp");

app.text = "MOE 365 Camera Control App!";
