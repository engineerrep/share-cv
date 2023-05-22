<template>
  <div class="footer" :class="`footer-${props.type}`">
    <div class=" footer-tail-content" :class="`footer-tail-${props.type}`">
      <div v-if="props.type==='orange'">
        <img class="orange-footer" src="@/assets/img/personal-center/preview-resume/orange-footer.png">
      </div>

      <div v-else-if="props.type==='purple'">
        <div class="footer-txt">Irresistible dating resume templates for you to use for free</div>
        <img class="footer-img" src="@/assets/img/personal-center/preview-resume/purple-preview-bottom.png">
      </div>
      <div v-else-if="props.type==='blue'" class="preview-tail">
        <div class="preview-tail-content">Irresistible dating resume templates for you to use for free</div>
        <img src="@/assets/img/personal-center/tail-img.png">
      </div>
      <div v-else-if="props.type==='pink'" class="preview-tail">
        <img src="@/assets/img/personal-center/preview-resume/pink-preview-bottom.png">
      </div>

    </div>
    <div class="preview-footer">
      <div>{{ url }}</div>
      <div class="preview-footer-right">
        <span>Powered by</span>
        <img class="resum-preview-logo cursor-pointer" src="@/assets/img/home_icon.png" alt=""/>
      </div>
    </div>
  </div>
</template>

<script setup>
let url = ref()
onMounted(() => {
  url.value = location.host
})

let props = defineProps({
  type: String,
})

</script>

<style lang="scss" scoped>
.preview-footer {
  font-size: 26px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #000000;
  line-height: 50px;
  -webkit-background-clip: text;
  display: flex;
  justify-content: space-between;
  margin: 0 54px;
  padding-bottom: 73px;
}

.resum-preview-logo {
  width: 140px;
  height: 50px;
}

.preview-footer-right {
  display: flex;
  align-items: center;

  span {
    margin-right: 10px;
  }
}

.footer-tail-content {
  position: relative;
}

.footer-txt {
  font-size: 40px;
  font-family: Almarai-Bold, Almarai;
  font-weight: bold;
  line-height: 60px;
  -webkit-background-clip: text;
  text-align: center;
  position: absolute;

}

.footer-tail-purple {
  height: 900px;

  .footer-txt {
    height: 270px;
    width: 720px;
    color: #6D81E2;
    margin-left: 136px;
    margin-top: 430px;
    margin-right: 136px;
  }

  .footer-img {
    position: absolute;
  }
}

.footer-tail-orange {
  position: relative;
  margin: 72px 50px 30px;
}

.preview-tail {
  position: relative;
  left: -50px;
  width: calc(100% + 100px);
  padding-top: 340px;
  padding-left: 88px;
  margin-bottom: 130px;

  .preview-tail-content {
    position: absolute;
    top: 266px;
    right: 0px;

    width: 432px;
    font-size: 40px;
    font-family: Almarai-Bold, Almarai;
    font-weight: bold;
    color: #000000;
    line-height: 60px;
    -webkit-background-clip: text;
  }
}

.footer-purple {
  //padding: 0 20px;
}

.footer-purple, .footer-pink, .footer-blue {
  top: 0;
  z-index: 20;
  width: 993px;
  margin: 0 auto;
  position: relative;

  .preview-footer {
    //padding-left: 50px;
    //padding-right: 50px;
    //margin: 0
  }
}

.footer-tail-blue {
  margin-top: 90px;
  margin-bottom: 100px;

  .preview-tail {
    width: 100%;
    //height: auto;
    margin-top: 0px;
    margin-bottom: 0;
    left: 0;
    padding-left: 54px;
    padding-right: 54px;

    img {
      width: 603px;
      height: auto;
      padding-top: 0;
    }
  }

  .preview-tail-content {
    right: 119px;
  }
}

.footer-tail-pink {
  .preview-tail {
    width: 100%;
    //height: auto;
    margin-top: -363px;
    margin-bottom: 0;
    left: 0;
    padding-left: 0;

    img {
      width: 100%;
      height: auto;
      padding-top: 0;
    }
  }
}

</style>
<style scoped lang="scss">
@media (max-width: 400px) {
  .footer {
    width: calc(100vw);
    padding-left: 20px;
    padding-right: 20px;
    //position: absolute;
    bottom: 0;
    padding-bottom: 40px;

    .footer-tail-orange {
      top: 30px;
      margin: 0 0 110px 0 !important;

    }
  }
  .preview-tail {
    position: relative;
    margin-top: 50px !important;
    left: 0 !important;
    width: auto !important;
    padding-top: 0px !important;
    padding-left: 0px !important;

    .preview-tail-content {
      font-size: 16px !important;
      font-weight: bold !important;
      color: #000000 !important;
      line-height: 24px !important;
      width: 184px !important;
      top: 0 !important;
      right: 10px !important;
    }

    img {
      width: 230px !important;
      height: 127px !important;
      padding-top: 30px;

    }
  }
  .preview-footer {
    font-size: 12px !important;
    line-height: 20px !important;
    margin: 0 !important;
    padding-bottom: 0 !important;

    .resum-preview-logo {
      width: 57px !important;
      height: 20px !important;
    }
  }
  .tail {
    .logo-min-img {
      width: auto !important;
    }
  }
  .footer-tail-purple {
    height: 365px !important;
    margin-top: -145px;

    .footer-txt {
      height: auto !important;
      width: auto !important;
      color: #6D81E2;
      margin-left: 60px !important;
      margin-top: 163px !important;
      margin-right: 60px !important;
      font-size: 16px !important;
      line-height: 24px !important;

    }
  }
  .footer-purple, .footer-pink {
    padding-left: 0 !important;
    padding-right: 0 !important;

    .preview-footer {
      padding-left: 20px !important;
      padding-right: 20px !important;
    }
  }
  .footer-pink {
    position: relative;
    z-index: 1;
  }

  .footer-tail-pink {
    .preview-tail {
      width: 100% !important;
      //height: auto;
      margin-top: -22px !important;
      margin-bottom: 0 !important;
      position: relative;
      z-index: -1;

      img {
        width: 100% !important;
        height: auto !important;
        padding-top: 0 !important;
      }
    }
  }
}
</style>