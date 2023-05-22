<template>
  <div class="h5-is-public">
    <div class="basic-required">
      <div class="title">Basic (required)</div>
      <div class="flex h5-basic-content">
        <img class="radio-disabled-img mr-12px" src="@/assets/img/personal-center/icon/selected-public-img.png">
        <div class="hint-txt">Name,avatar,country,city,borthday,age</div>
      </div>
    </div>

    <div class="h5-space "></div>

    <div class="public">
      <div class="title">
        Public
      </div>
      <div class="public-main">
        <div class="public-list">
          <div class="flex " @click="setDisplay">
            <img
                v-if="status"
                class="radio-disabled-img mr-6px cursor-pointer"
                src="@/assets/img/personal-center/icon/choose-public.png">
            <img
                v-else
                class="radio-disabled-img mr-6px cursor-pointer"
                src="@/assets/img/personal-center/not-selected-public-img.png">
            <div class="public-plate">Your Settings can be seen in the resume preview</div>
          </div>
          <div>
            <img v-if="showHidePublic"
                 @click.stop="showHidePublic=false"
                 class="public-on-iocn" src="@/assets/img/personal-center/public-on-iocn.png">
            <img v-else class="public-on-iocn"
                 @click.stop="showHidePublic=true"
                 src="@/assets/img/personal-center/public-off-iocn.png">
          </div>
        </div>

        <div class="public-content" v-if="showHidePublic">
          <div class="public-item" v-for="item in publicList" :key="item.txt">
            <el-checkbox v-model="item.state" :label="item.txt "
                         @change="changePublic" size="large"/>

          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script setup>
let status = ref(true)
let showHidePublic = ref(true)
let themeColor = useState("_themeColor")
// 预览页面可以展示的内容
let publicList = ref([])
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

</script>

<style>
.el-checkbox__input.is-checked + .el-checkbox__label {
  color: #000;
}

.el-checkbox__input {
  margin-left: 2px;
  margin-right: 3px;
}

.el-checkbox__inner {
  border: 1px solid #616161;
}

.el-checkbox__label {
  color: #000;
  padding-left: 6px;
}
</style>
<style lang="scss" scoped>
:deep(.el-checkbox__input.is-checked ) .el-checkbox__inner {
  background-color: v-bind(themeColor);
  border-color: v-bind(themeColor);
}

.title {
  height: 36px;
  font-size: 22px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: normal;
  color: #000000 !important;
  line-height: 26px;
  -webkit-background-clip: text;
  margin-bottom: 12px;
}

.basic-required {
  padding: 25px 0;

  .radio-disabled-img {
    width: 20px;
    height: 20px;
  }

  .hint-txt {
    height: 13px;
    font-size: 15px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #333333;
    line-height: 14px;
    -webkit-background-clip: text;
  }
}

.public {
  padding-top: 25px;
  margin-bottom: 32px;
}

.radio-disabled-img {
  width: 20px;
  height: 20px;
}

.public-main {
  .public-plate {
    width: 273px;
    height: 27px;
    font-size: 15px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #333333;
    line-height: 14px;
    -webkit-background-clip: text;
  }

  .public-on-iocn {
    position: relative;
    right: 2px;
    width: 28px;
    height: 28px;
  }

  .public-list {
    display: flex;
    justify-content: space-between;
  }
}


.public-content {
  display: flex;
  flex-wrap: wrap;
  font-size: 18px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #000000;
  line-height: 21px;
  position: relative;
  z-index: 1;
  padding-top: 12px;

  .public-item {
    display: flex;
    align-items: center;
    height: 22px;
    margin-top: 16px;
  }

  > :nth-child(2n-1) {
    width: 250px;
  }
}


@media (max-width: 420px) {
  .h5-space {
    position: relative;
    left: -20px;
    width: calc(100% + 40px);
    border-top: 9px solid #F4F2F0;
    margin-top: 16px;
  }
  .basic-required {
    .radio-disabled-img {
      width: 24px;
      height: 24px;
    }

    .h5-basic-content {
      display: flex;
      align-items: center;
    }
  }
  .radio-disabled-img {
    width: 24px;
    height: 24px;
  }
}

</style>