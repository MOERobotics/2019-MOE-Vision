<template>
  <div>
    <div class="container is-centered">
      <div class="notification">
        <h1>
          <center>
            <big>
              <strong>
                <u>{{text}}</u>
              </strong>
            </big>
          </center>
        </h1>

        <template>
          <a class="button is-primary is-rounded" @click="handleResetAll">Reset all</a>
          <b-dropdown>
            <button class="button is-primary is-rounded" slot="trigger">
              <span>Reset select:</span>
              <b-icon icon="menu-down"></b-icon>
            </button>

            <b-dropdown-item @click="selecty('exposureClick')">Exposure</b-dropdown-item>
            <b-dropdown-item @click="selecty('focusClick')">Focus</b-dropdown-item>
            <b-dropdown-item @click="selecty('brightnessClick')">Brightness</b-dropdown-item>
            <b-dropdown-item @click="selecty('contrastClick')">Contrast</b-dropdown-item>
            <b-dropdown-item @click="selecty('saturationClick')">Saturation</b-dropdown-item>
            <b-dropdown-item @click="selecty('sharpnessClick')">Sharpness</b-dropdown-item>
            <b-dropdown-item @click="selecty('whiteBalClick')">White Balance Temp.</b-dropdown-item>
            <b-dropdown-item @click="selecty('backlightClick')">Backlight Compensation</b-dropdown-item>
          </b-dropdown>
        </template>
        <div>
          <button type="button" @click="kennethButton">Print current values</button>
        </div>
        <center>
          <h3>
            <b>Exposure Control</b>
          </h3>
          <div class="slidecontainer">
            <input
              type="range"
              min="1"
              max="100"
              v-model="exposureVal"
              class="slider"
              id="Exposure"
              @change="exposure"
            >
          </div>
          <h3>
            <b>Focus Control</b>
          </h3>
          <div class="slidecontainer">
            <input
              type="range"
              min="1"
              max="100"
              v-model="focusVal"
              class="slider"
              id="Focus"
              @change="focus"
            >
          </div>
          <h3>
            <b>Brightness control</b>
          </h3>
          <div class="slidecontainer">
            <input
              type="range"
              min="1"
              max="100"
              v-model="brightnessVal"
              class="slider"
              id="Brightness"
              @change="brightness"
            >
          </div>
          <h3>
            <b>Contrast control</b>
          </h3>
          <div class="slidecontainer">
            <input
              type="range"
              min="1"
              max="100"
              v-model="contrastVal"
              class="slider"
              id="Contrast"
              @change="contrast"
            >
          </div>
          <h3>
            <b>Saturation control</b>
          </h3>
          <div class="slidecontainer">
            <input
              type="range"
              min="1"
              max="100"
              v-model="saturationVal"
              class="slider"
              id="Saturation"
              @change="saturation"
            >
          </div>
          <h3>
            <b>Sharpness control</b>
          </h3>
          <div class="slidecontainer">
            <input
              type="range"
              min="1"
              max="100"
              v-model="sharpnessVal"
              class="slider"
              id="Sharpness"
              @change="sharpness"
            >
          </div>
          <h3>
            <b>White Balance Temperature control</b>
          </h3>
          <div class="slidecontainer">
            <input
              type="range"
              min="1"
              max="100"
              v-model="whiteBalanceVal"
              class="slider"
              id="whiteBalanceTemperature"
              @change="whiteTemp"
            >
          </div>
          <h3>
            <b>Backlight Compensation control</b>
          </h3>
          <div class="slidecontainer">
            <input
              type="range"
              min="1"
              max="100"
              v-model="backliteVal"
              class="slider"
              id="backlightCompensation"
              @change="backlite"
            >
          </div>
        </center>
      </div>
    </div>
  </div>
</template>


<script>
import UVCControl from "uvc-control";
import Buefy from "buefy";

export default {
  name: "App",
  data() {
    return {
      text: "",
      camera: new UVCControl(0x046d, 0x082d),
      backliteval: 50,
      whiteBalanceVal: 50,
      sharpnessVal: 50,
      saturationVal: 50,
      contrastVal: 50,
      brightnessVal: 50,
      focusVal: 50,
      exposureVal: 50,
      dropdown: null
    };
  },
  methods: {
    exposure(event) {
      console.log(event.srcElement.id + " set to: " + event.srcElement.value);

      this.camera.set("absoluteExposureTime", event.srcElement.value, function(
        error
      ) {
        if (!error) {
          console.log("Exposure Set OK!");
        }
      });
    },
    focus(event) {
      console.log(event.srcElement.id + " set to: " + event.srcElement.value);

      this.camera.set("absoluteFocus", event.srcElement.value, function(error) {
        if (!error) {
          console.log("Focus Set OK!");
        }
      });
    },
    brightness(event) {
      console.log(event.srcElement.id + " set to: " + event.srcElement.value);

      this.camera.set("brightness", event.srcElement.value, function(error) {
        if (!error) {
          console.log("Brightness Set OK!");
        }
      });
    },
    contrast(event) {
      console.log(event.srcElement.id + " set to: " + event.srcElement.value);

      this.camera.set("contrast", event.srcElement.value, function(error) {
        if (!error) {
          console.log("Contrast Set OK!");
        }
      });
    },
    saturation(event) {
      console.log(event.srcElement.id + " set to: " + event.srcElement.value);

      this.camera.set("saturation", event.srcElement.value, function(error) {
        if (!error) {
          console.log("Saturation Set OK!");
        }
      });
    },
    sharpness(event) {
      console.log(event.srcElement.id + " set to: " + event.srcElement.value);

      this.camera.set("sharpness", event.srcElement.value, function(error) {
        if (!error) {
          console.log("Sharpness Set OK!");
        }
      });
    },
    whiteTemp(event) {
      console.log(event.srcElement.id + " set to: " + event.srcElement.value);

      this.camera.set(
        "whiteBalanceTemperature",
        event.srcElement.value,
        function(error) {
          if (!error) {
            console.log("White Balance Temperature Set OK!");
          }
        }
      );
    },
    backlite(event) {
      console.log(event.srcElement.id + " set to: " + event.srcElement.value);

      this.camera.set("backlightCompensation", event.srcElement.value, function(
        error
      ) {
        if (!error) {
          console.log("Backlight Compensation Set OK!");
        }
      });
    },
    handleResetAll(event) {
      console.log("Resetting all values...");

      this.backliteVal = 50;
      this.whiteBalanceVal = 50;
      this.sharpnessVal = 50;
      this.saturationVal = 50;
      this.contrastVal = 50;
      this.brightnessVal = 50;
      this.focusVal = 50;
      this.exposureVal = 50;
    },
    selecty(event) {
      if (event == "exposureClick") {
        this.exposureVal = 50;
        console.log("Exposure RESET OK!");
      } else if (event == "backlightClick") {
        this.backliteVal = 50;
        console.log("Backlight RESET OK!");
      } else if (event == "whiteBalClick") {
        this.whiteBalanceVal = 50;
        console.log("White Balance RESET OK!");
      } else if (event == "contrastClick") {
        this.contrastVal = 50;
        console.log("Contrast RESET OK!");
      } else if (event == "sharpnessClick") {
        this.sharpnessVal = 50;
        console.log("Sharpness RESET OK!");
      } else if (event == "saturationClick") {
        this.saturationVal = 50;
        console.log("Saturation RESET OK!");
      } else if (event == "brightnessClick") {
        this.brightnessVal = 50;
        console.log("Brightness RESET OK!");
      } else if (event == "focusClick") {
        this.focusVal = 50;
        console.log("Focus RESET OK!");
      } else {
      }
    },
    kennethButton(event) {
      this.camera.get("brightness", function(error, value) {
        console.log("Brightness setting:", value);
      });
    }
  }
};
</script>
