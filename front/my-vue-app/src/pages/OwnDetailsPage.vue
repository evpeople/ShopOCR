<template>
  <div class="my-details-container">
    <div class="nav-part details-module" :style="switchNavBackground">
        <button class="first-button" @click="switchToSeek" id="firstButton">我的寻味道</button>
        <button class="second-button" @click="switchToOffer" id="secondButton">我的请品鉴</button>
    </div>
    <div class="display-part details-module">
        <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
    name: 'OwnDetailsPage',
    data() {
        return {
            switchDisplay: 1,
        }
    },
    computed: {
        switchNavBackground() {
            if (this.switchDisplay === 1) {
                return 'background-image: linear-gradient(135deg, #e63e7e 0%, #ffc6a5 65%, #e63e7e 100%);'
            } else {
                return 'background-image: linear-gradient(135deg, #F7CE68 0%, #ff9987 65%, #ff8742 100%);'
            }
        }
    },
    methods: {
        switchToSeek() {
            this.switchDisplay = 1;
            this.$router.push({
                name: 'seek'
            })
        },
        switchToOffer() {
            this.switchDisplay = 2
            this.$router.push({
                name: 'offer'
            })
        }
    },
    mounted() {
        this.$bus.$on("MySwitch",(data) => {
            if (data===1) {
                this.switchToSeek();
            } else {
                this.switchToOffer();
            }
        })
    },
    beforeDestroy() {
        this.$bus.$off("MySwitch")
    },
}
</script>

<style scoped>
.my-details-container {
  margin: 0 auto;
  width: 1310px;
  height: 730px;
  position: relative;
  background: #fff;
  box-shadow: 0 20px 50px rgb(65 62 101 / 15%);
}

.details-module {
  float: left;
}

.nav-part {
  width: 150px;
  height: 730px;
}

.nav-part button {
    width: 100%;
    height: 40px;
    outline: none;
    font-size: 14px;
    color: #fff;
    cursor: pointer;
}

.first-button {
    margin-top: 20px;
    border-top: rgb(255, 156, 169) 1px solid;
    border-bottom: none;
    border-left: none;
    border-right: none;
    background-image: linear-gradient(62deg, #e467aa 0%, #f65185 85%);
}

.first-button:hover, .first-button:focus {
    background-image: linear-gradient(62deg, #e33391 0%, #cf1f57 85%);
}

.second-button {
    border-top: none;
    border-bottom: #ffe8af 1px solid;
    border-left: none;
    border-right: none;
    background-image: linear-gradient(62deg, #FBAB7E 0%, #F7CE68 100%);
}

.second-button:hover, .second-button:focus {
    background-image: linear-gradient(62deg, #ff8742 0%, #d5a734 100%);
}

.display-part {
  width: 1160px;
}
</style>