<template>
  <div class="select-resume">
    <client-only>
      <el-dialog
          v-model="templates"
          class="chose-template-dialog"
          :fullscreen="true"
          :lock-scroll="true"
      >
        <template #header>
          <div class="template-title" v-if="pcSide">
            Choose the resume template you want
          </div>
        </template>
        <img @click="templates=false" class="return-img" src="@/assets/img/return-img.png">
        <el-carousel class="chose-template-carousel"
                     arrow="always" :initial-index="defaultCarouseIndex"
                     @change="carouselChange"
                     :interval="4000" type="card" :autoplay="false" height="473px" indicator-position="none">
          <el-carousel-item class="template-item">
            <img class="temp-orange" src="@/assets/img/personal-center/temp-orange.png">
          </el-carousel-item>
          <el-carousel-item class="template-item">
            <img class="temp-blue" src="@/assets/img/personal-center/temp-blue.png">
          </el-carousel-item>
          <el-carousel-item class="template-item">
            <img class="temp-purple" src="@/assets/img/personal-center/temp-purple.png">
          </el-carousel-item>
          <el-carousel-item class="template-item">
            <img class="temp-pink" src="@/assets/img/personal-center/temp-pink.png">
          </el-carousel-item>
        </el-carousel>
        <el-button @click="choseThisTemplate()" class="chose-button">
          Use this template
        </el-button>
      </el-dialog>
    </client-only>
  </div>
</template>

<script setup>
import {useState} from "#app";

let pcSide = useState('_pcSide');
const templates = useState("_templates")

let defaultCarouseIndex = ref(0)

const carouselChange = (index) => {
  defaultCarouseIndex.value = index
}
onMounted(() => {
  document.body.style.cssText = 'overflow-y:hidden';
  document.documentElement.style.cssText = 'overflow-y:hidden';
})

onUnmounted(() => {
  document.body.style.cssText = 'overflow-y:auto';
  document.documentElement.style.cssText = 'overflow-y:auto';
  templates.value = false
})
let userId = useState("_userId");
const choseThisTemplate = () => {
  let type = ['orange', 'blue', 'purple', 'pink'][defaultCarouseIndex.value]
  let inter = setInterval(() => {
    if (userId.value) {
      window.localStorage.setItem(`${userId}_template`, type)
      clearInterval(inter);
      navigateTo(`/preview-resume/${type}`)
    }
  }, 300)
}
</script>


<style lang="scss" scoped>
.template-title {
  font-size: 20px;
  font-weight: normal;
  color: #171725;
  line-height: 48px;
  text-align: center;
  margin-top: 20px;
}

.chose-template-dialog {
  position: relative;

  .el-dialog__headerbtn {
    background: url("@/assets/img/delete-img.png") no-repeat;
    width: 25px;
    height: 24px;
    top: 17px;
    right: 17px;
    background-size: 100% 100%;

    .el-dialog__close {
      display: none;
    }
  }

  &:before {
    position: absolute;
    content: '';
    width: 100%;
    height: 60%;
    background: url("@/assets/img/personal-center/template-chose-back.png") no-repeat;
    background-size: 100% 100%;
    pointer-events: none
  }
}

.chose-template-carousel {
  .el-carousel__arrow--left {
    background-color: #0A66C2;
  }

  .el-carousel__arrow--right {
    background-color: #0A66C2;
  }

  .is-active {
    border: 3px solid #0A66C2;
  }
}

@media (max-width: 420px) {
  .return-img {
    width: 24px;
    position: absolute;
    top: 18px;
    left: 18px;
    z-index: 10;
  }
  .select-resume {
    :deep(.el-dialog__body) {
      padding: 0;
    }

    :deep(.el-carousel__item--card) {
      width: 100%;
      transform: translateX(0) scale(1) !important;
    }
  }

  .chose-button {
    width: calc(100% - 48px);
    height: 40px;
    display: block;
    margin: 30px auto;
    font-size: 18px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #FFFFFF;
    line-height: 21px;
    -webkit-background-clip: text;
    background: #0A66C2;
    border-radius: 20px;
    position: absolute;
    left: 24px;
    bottom: 0;
  }

  .chose-template-carousel {
    .is-active {
      border: 0;
    }
  }

  :deep(.el-dialog__header) {
    display: none;
  }
  .temp-orange {
    max-width: calc(100% + 2px);
  }
  :deep(.el-carousel ) {
    height: calc(100% - 60px);
  }
  :deep(.el-carousel__container ) {
    height: 100% !important;
  }
  :deep(.el-dialog__body) {
    height: calc(100% - 50px);
  }
  :deep(.el-carousel__arrow) {
    width: 40px;
    height: 40px;

    .el-icon {
      display: none;
    }
  }
  :deep(.el-carousel__arrow--right) {
    background-image: url("@/assets/img/personal-center/preview-resume/carousel-arrow-right.png");
    background-size: 40px;
  }
  :deep(.el-carousel__arrow--left) {
    background-image: url("@/assets/img/personal-center/preview-resume/carousel-arrow-left.png");
    background-size: 40px;
  }

  :deep(.el-carousel__item) {
    overflow-y: scroll;
  }
}
</style>