<template>
  <div class="is-public ml-112px">
    <div class="is-public-main">
      <div class="share-resume cursor-pointer"
           @click="goShareResume()"
      >
        <img src="@/assets/img/personal-center/icon/share-resume-iocn.png">
        <span>Share resume</span>
      </div>
      <div class="share-resume-template cursor-pointer"
           @click="choseTemplate()"
      >
        <img src="@/assets/img/personal-center/icon/share-resume-template-iocn.png">
        <span>Template</span>
      </div>
      <div class="plate-three mt-28px">
        <div>
          Basic (required)
        </div>
        <div class="plate-three-content flex  break-words mb-50px ">
          <img class="radio-disabled-img mr-12px"
               src="@/assets/img/personal-center/icon/selected-public-img.png">
          <p class="break-words plate-three-txt ">Name,avatar,country,city,borthday,age,languagrs</p>
        </div>
      </div>

      <div class="is-public-subject plate-three ">
        <div>
          Public
        </div>

        <div class="plate-three-content flex text-17px break-words mb-50px ">
          <img
              v-if="status"
              @click="setDisplay"
              class="radio-disabled-img mr-12px cursor-pointer"
              src="@/assets/img/personal-center/icon/choose-public.png">
          <img
              v-else
              @click="setDisplay"
              class="radio-disabled-img mr-12px cursor-pointer"
              src="@/assets/img/personal-center/not-selected-public-img.png">
          <p class="break-words plate-three-txt w-282px">Your Settings can be seen in the resume preview</p>
        </div>

        <div class="font-family flex justify-between" v-for="item in publicList" :key="item.txt">
          <div>{{ item.txt }}</div>
          <div class="flex items-center flex-col pt-12px">
            <el-checkbox v-model="item.state" @change="changePublic" size="large"/>
            <span class="mt-8px">{{ item.state ? 'Show' : 'Hide' }}</span>
          </div>
        </div>
      </div>
    </div>
    <client-only>
      <el-dialog
          v-model="templates"
          class="chose-template-dialog"
          :fullscreen="true"
      >
        <template #header>
          <div class="template-title">
            Choose the resume template you want
          </div>
        </template>
        <el-carousel class="chose-template-carousel"
                     arrow="always" :initial-index="defaultCarouseIndex"
                     @change="carouselChange"
                     :interval="4000" type="card" :autoplay="false" height="473px" indicator-position="none">
          <el-carousel-item class="template-item">
            <img src="@/assets/img/personal-center/temp-orange.png">
            <!--                        <div class=" template-orange">-->

            <!--                        </div>-->
          </el-carousel-item>
          <el-carousel-item class="template-item">
            <img src="@/assets/img/personal-center/temp-blue.png">

          </el-carousel-item>
          <el-carousel-item class="template-item">
            <img src="@/assets/img/personal-center/temp-purple.png">

          </el-carousel-item>
          <el-carousel-item class="template-item">
            <img src="@/assets/img/personal-center/temp-pink.png">
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
import {useRouter} from "vue-router";
import Button from "~/components/form/button.vue";

let themeColor = useState("_themeColor")
const templates = ref(false)
let defaultCarouseIndex = ref(0)
const router = useRouter();
// 预览页面可以展示的内容
let publicList = ref([])
let status = ref(true)
let isShareResume = useState("_isShareResume", () => false)
let userId = null;
onMounted(() => {
  let inter = setInterval(() => {
    userId = window.localStorage.getItem('userId')
    if (userId) {
      setOrGetpublicListInfo(userId)
      clearInterval(inter)
    }
  }, 300)
})
const setOrGetpublicListInfo = (userId) => {
  let localPublicList = window.localStorage.getItem(`${userId}_publicList`)
  if (localPublicList) {
    publicList.value = JSON.parse(localPublicList)
  } else {
    let tempPublicList = [
      {txt: "Email", state: true},
      {txt: "Phone", state: true},
      {txt: "About me", state: true},
      {txt: "Languages", state: true},
      {txt: "Interests", state: true},
      {txt: "Skills", state: true},
      {txt: "Photos", state: true},
      {txt: "Videos", state: true},
      {txt: "Education", state: true},
      {txt: "Experience", state: true},
      {txt: "Job Experience", state: true},
      {txt: "Project Experience", state: true},
      {txt: "Volunteer Experience", state: true},
      {txt: "Custom item", state: true},
    ]
    window.localStorage.setItem(`${userId}_publicList`, JSON.stringify(tempPublicList))
    publicList.value = tempPublicList
  }
}
// 设置-预览页面显示-全选-取消全选
const setDisplay = () => {
  let state = publicList.value.every(item => item.state === true);
  if (state) {
    publicList.value.forEach(item => {
      item.state = false
    })
    status.value = false
  } else {
    publicList.value.forEach(item => {
      item.state = true
    })
    status.value = true
  }
  window.localStorage.setItem(`${userId}_publicList`, JSON.stringify(publicList.value))

}

// 设置-预览页面显示-单选
const changePublic = () => {
  let state = publicList.value.some(item => item.state === true);
  status.value = state;
  window.localStorage.setItem(`${userId}_publicList`, JSON.stringify(publicList.value))
}
/**
 * 去简历预览页面
 */
const goShareResume = () => {
  let inter = setInterval(() => {
    if (userId) {
      let type = window.localStorage.getItem(`${userId}_template`)
      clearInterval(inter);
      if (type) {
        navigateTo(`/preview-resume/${type}`)
      } else {
        templates.value = true
      }
    }
  }, 300)
  // navigateTo('/preview-resume/purple')
}
/**
 * 选择模板
 */
const choseTemplate = () => {
  templates.value = true
}
const carouselChange = (index) => {
  defaultCarouseIndex.value = index
}
/**
 * 选择当前模板
 */
const choseThisTemplate = () => {
  let type = ['orange', 'blue', 'purple', 'pink'][defaultCarouseIndex.value]
  console.log(defaultCarouseIndex.value)
  let inter = setInterval(() => {
    if (userId) {

      window.localStorage.setItem(`${userId}_template`, type)
      clearInterval(inter);
      router.push(`/preview-resume/${type}`)
    }
  }, 300)

}
</script>

<style>
.is-public {
  .el-checkbox__input.is-checked .el-checkbox__inner {
    background-color: v-bind(themeColor);
    border-color: v-bind(themeColor);
  }
}


</style>
<style lang="scss" scoped>
.is-public {
  width: 376px;
  overflow: hidden;

  .is-public-head {
    height: 110px;
    line-height: 30px;
    background: #E3E9EE;
    padding: 25px 28px 16px;
    font-size: 26px;
    border-radius: 10px 10px 0px 0px;
    border: 1px solid #E2E0DF;
  }

  .share-resume, .share-resume-template {
    width: 376px;
    height: 70px;
    line-height: 70px;
    background: #DFE6EB;
    border-radius: 10px 10px 0px 0px;
    opacity: 1;
    font-size: 25px;
    font-weight: 500;
    color: #0A66C2;

    img {
      display: inline-block;
      margin-left: 27px;
      margin-right: 12px;
      width: 28px;
      height: 28px;
    }
  }

  .share-resume-template {
    color: #616161;
    border-radius: 0;
  }

  .is-public-main {
    height: auto;
    background-color: #fff;
    border-radius: 10px;
    opacity: 1;
    border: 1px solid #E2E0DF;
    font-size: 21px;
    padding-bottom: 22px;
  }

  .radio-disabled-img {
    width: 28px;
    height: 28px;
  }

  .is-public-subject {
    border-radius: 0px 0px 10px 10px;
    opacity: 1;
    border-top: 1px solid #E2E0DF;
    padding-top: 16px;
  }

  .font-family {
    font-family: Inter-Medium, Inter;
  }
}

.plate-three {
  padding: 0 28px;

  .plate-three-content {
    margin-top: 22px;
    font-size: 17px;
  }

  .plate-three-txt {
    width: 282px;
    font-size: 17px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #616161;
    line-height: 20px;
  }
}

.template-title {
  font-size: 41px;
  font-weight: normal;
  color: #171725;
  line-height: 48px;
  text-align: center;
  margin-top: 100px;
  position: relative;
}


.template-item {
  width: 331px;
  height: 473px;
  overflow-y: auto;
  box-shadow: 0px 12px 24px 0px rgba(0, 0, 0, 0.25);
}

.chose-template-carousel {
  width: 1100px;
  margin: 0 auto;

  .el-carousel__container {
    .el-carousel__item {
      margin-left: 90px;
    }
  }
}

:deep(.chose-template-carousel) {
  .el-carousel__arrow--left {
    left: 348px;
    font-size: 20px;
    background-color: #0A66C2;
    top: 80%;
    font-weight: bold;
  }

  .el-carousel__arrow--right {
    right: 389px;
    font-size: 20px;
    background-color: #0A66C2;
    top: 80%;
    font-weight: bold;
  }

  .is-active {
    box-shadow: 0px 4px 30px 0px rgba(0, 0, 0, 0.2);
    border-radius: 6px 6px 6px 6px;
    opacity: 1;
    border: 3px solid #0A66C2;

    .template-purple, .template-blue, .template-pink, .template-orange {
      margin-left: -2px;
    }
  }
}

.chose-button {
  width: 259px;
  height: 62px;
  background: #0A66C2;
  border-radius: 39px 39px 39px 39px;
  color: #fff;
  font-size: 21px;
  line-height: 62px;
  display: flex;
  justify-content: center;
  margin: 42px auto 30px auto;
}

:deep(.chose-template-dialog) {
  position: relative;

  .el-dialog__headerbtn {
    background: url("@/assets/img/delete-img.png") no-repeat;
    width: 25px;
    height: 24px;
    top: 17px;
    right: 17px;
    background-size: 100% 100%;
    display: flex;

    .el-dialog__close {
      display: none;
    }
  }

  &:before {
    position: absolute;
    content: '';
    background: url("@/assets/img/personal-center/template-chose-back.png") no-repeat;
    background-size: 100% 100%;
    width: 1100px;
    height: 666px;
    left: calc(50% - 550px);
    pointer-events: none
  }

  .el-carousel__arrow {
    position: absolute;
    top: 50%;
  }

}

@media (max-width: 1367px) {
  .plate-three {
    padding: 0 20px;
  }
}

@media (max-width: 1025px) {
  .plate-three-txt {
    width: 250px !important;
  }
  .chose-template-carousel {
    width: 100%;
  }
  :deep(.chose-template-dialog) {
    &:before {
      width: 100%;
    }
  }
}
</style>