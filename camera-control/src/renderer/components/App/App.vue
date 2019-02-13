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

            <b-dropdown-item>Exposure</b-dropdown-item>
            <b-dropdown-item>Zoom</b-dropdown-item>
            <b-dropdown-item>Pan Tilt</b-dropdown-item>
            <b-dropdown-item>Focus</b-dropdown-item>
            <b-dropdown-item>Brightness</b-dropdown-item>
            <b-dropdown-item>Contrast</b-dropdown-item>
            <b-dropdown-item>Saturation</b-dropdown-item>
            <b-dropdown-item>Sharpness</b-dropdown-item>
            <b-dropdown-item>White Balance Temp.</b-dropdown-item>
            <b-dropdown-item>Backlight Contrast</b-dropdown-item>
          </b-dropdown>
        </template>

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
            <b>Zoom control</b>
          </h3>
          <div class="slidecontainer">
            <input
              type="range"
              min="1"
              max="100"
              v-model="zoomVal"
              class="slider"
              id="Zoom"
              @change="zoom"
            >
          </div>
          <h3>
            <b>Pan Tilt control</b>
          </h3>
          <div class="slidecontainer">
            <input
              type="range"
              min="1"
              max="100"
              v-model="panTiltval"
              class="slider"
              id="panTilt"
              @change="pantilt"
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
              @change="backlight"
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
      panTiltval: 50,
      zoomVal: 50,
      exposureVal: 50
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
    zoom(event) {
      console.log(event.srcElement.id + " set to: " + event.srcElement.value);

      this.camera.set("absoluteZoom", event.srcElement.value, function(error) {
        if (!error) {
          console.log("Zoom Set OK!");
        }
      });
    },
    pantilt(event) {
      console.log(event.srcElement.id + " set to: " + event.srcElement.value);

      this.camera.set("absolutePanTilt", event.srcElement.value, function(
        error
      ) {
        if (!error) {
          console.log("Pan Tilt Set OK!");
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
    backlight(event) {
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

      this.backliteval = 50;
      this.whiteBalanceVal = 50;
      this.sharpnessVal = 50;
      this.saturationVal = 50;
      this.contrastVal = 50;
      this.brightnessVal = 50;
      this.focusVal = 50;
      this.panTiltval = 50;
      this.zoomVal = 50;
      this.exposureVal = 50;
    }
  }
};
</script>
