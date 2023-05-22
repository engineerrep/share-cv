<template>
  <div class="personal-img">
    <!-- 头像和背景 -->
    <div class="relative">
      <div class="photo-ba h-260px">

        <img v-if="userInfoStore.photoWallImg" class="photo-wall-img h-260px w-full" :src="userInfoStore.photoWallImg" alt=""/>

        <img
            src="@/assets/img/personal-center/icon/camera-img.png"
            class="camera-img cursor-pointer"
            @click="showAddToBulletBox('baEdit')"
        />
      </div>
      <div
          class="avatar-dom cursor-pointer"
          @click="showAddToBulletBox('avatarEdit')"
      >
        <img
            v-if="avatarImg"
            :src="avatarImg"
            alt=""
            class="yes-avatar-dom-img pr"
        />
        <img
            v-else
            src="@/assets/img/personal-center/icon/change-avatar.png"
            alt=""
            class="avatar-dom-img pr"
        />
      </div>
    </div>

    <!-- 用户信息-->
    <div class="personal-wditor-center">
      <!-- 个人信息 -->
      <div class="personal-information relative flex ">
        <div class="h5-w w-1/2  mr-8px">
          <p class="family-name">{{ userInfoStore.personaiInformation?.firstName }}&nbsp;
            {{ userInfoStore.personaiInformation?.lastName }}</p>
          <!--          <p class="occupation">{{ userInfoStore.personaiInformation.nickName }}&nbsp;</p>-->

          <div class="flex personal-information-address mt-6px mb-4px">

            <div class=" ">
              {{ userInfoStore.personaiInformation?.countryName }}
              {{ userInfoStore.personaiInformation?.city ? ',' + userInfoStore.personaiInformation?.city : "" }}
            </div>

          </div>
          <div class="personal-information-age">
            {{ userInfoStore.ageInformation || 0 }}yearsold
            <span v-if="userInfoStore.personaiInformation?.birthday">,</span>
            {{ userInfoStore.personaiInformation?.birthday }}
          </div>

          <img
              src="@/assets/img/personal-center/edit-square-img.png"
              alt=""
              class="edit-square-img cursor-pointer"
              @click="showAddToBulletBox('information')"
          />
        </div>

        <div class="contact-information w-1/2 ml-8px flex flex-col-reverse">
          <div class="flex contact-information-item items-center">
            <img class="mr-16px" src="@/assets/img/personal-center/icon/phone-iocn.png">
            <span @click="setUpContact('phone')" class="cursor-pointer">{{
                userInformation?.phone || 'Add phone'
              }}</span>
          </div>
          <div class="flex mb-8px contact-information-item items-center">
            <img class="mr-16px" src="@/assets/img/personal-center/icon/mail-iocn.png">
            <span @click="setUpContact('mail')" class="cursor-pointer">{{ userInformation?.email }}</span>
          </div>
        </div>
      </div>

      <div v-if="!pcSide" class="h5-share-resume">
        <img class="h5-share" @click="goShareResume()" src="@/assets/img/personal-center/h5-share.png">
        <img class="h5-template" @click="choseTemplate()" src="@/assets/img/personal-center/h5-template.png">
      </div>
      <!-- 照片和视频 -->
      <div class="photos-videos flex mt-32px ">
        <div class="photos-videos-item w-1/2 mr-8px cursor-pointer" @click="showUpload('photos')">
          <div class=" plate-four">
            <img class="photos-videos-img mr-16px" src="@/assets/img/personal-center/icon/photos-iocn.png">
            <div class="plate-four-title">Photos</div>
            <img class="union" src="@/assets/img/personal-center/union.png" alt="">
          </div>
        </div>
        <div class="photos-videos-item  w-1/2 ml-8px cursor-pointer" @click="showUpload('videos')">
          <div class=" plate-four">
            <img class="photos-videos-img mr-16px" src="@/assets/img/personal-center/icon/videos-iocn.png">
            <div class="plate-four-title">Videos</div>
            <img class="union" src="@/assets/img/personal-center/union.png" alt="">
          </div>
        </div>
      </div>
    </div>
  </div>

  <Upload
      v-if="showHideUpload"
      :uploadType="uploadType"
  ></Upload>
  <SelectResume v-if="templates"></SelectResume>
</template>

<script setup>
import {useRouter} from "vue-router";
import {useState} from "#app";
import {calculateHowOld, showAddToBulletBox} from "~/utils/util";
import Upload from "@/components/personal-wditor/upload.vue";
import SelectResume from "@/components/preview-resume/select-resume.vue";
import {userStore} from "@/composables/store/user.js"
let userInfoStore = userStore()

const router = useRouter();

let userInformation = useCookie('_userInformation', () => {
  return {
    phone: null,
    email: null,
  }
})

let templates = useState("_templates", () => false)
let pcSide = useState('_pcSide');
let userId = useState("_userId");
onMounted(() => {
  let inter = setInterval(() => {
    userId.value = window.localStorage.getItem('userId')
    if (userId.value) {
      clearInterval(inter)
    }
  }, 300)
})
const choseTemplate = () => {
  templates.value = true
}
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
// 头像
let avatarImg = useState('_avatarImg')





// if (userInfoStore.personaiInformation?.birthday) userInfoStore.ageInformation =  calculateHowOld(userInfoStore.personaiInformation?.birthday)

/**
 * 设置联系方式
 * @param {type} phone-电话 mail-邮箱
 * */
const setUpContact = (type) => {
  navigateTo({path: '/forgot-password', query: {type}})
  document.body.style.width = '100%'
  document.body.style.minWidth = '100%'
}

// 显示隐藏上传组件
let showHideUpload = useState("_showHideUpload", () => false)
// showHideUpload = true
// 上传状态
let uploadType = ref(null)
uploadType.value = 'photos'
/**
 * 显示上传照片和视频操作
 * @param {type} photos-照片 videos-视频
 * */
const showUpload = (type) => {
  showHideUpload.value = true
  uploadType.value = type
}

</script>

<style lang="scss">


@media (max-width: 420px) {
  .el-dialog__headerbtn {
    display: none;
  }
}

.null-data {
  .el-upload-list {
    top: 0;
  }

  .select-upload {
    width: 673px;
    height: 420px;
  }

  .el-upload-dragger {
    position: relative;
    z-index: 100;
    //left: -2px;
    width: 674px;
    height: 460px;
    overflow: hidden;
    border: none;
    box-shadow: inset 0px 0px 6px 0px rgba(0, 0, 0, 0.1);
    border-radius: 10px;
  }

  .el-upload {
    position: relative;
    top: -26px;
    left: -30px;;
    width: 673px;
    height: 420px;
  }
}
</style>

<style lang="scss" scoped>
.personal-img {
  background-color: #fff;
  border-radius: 10px;
  padding-bottom: 32px;
  overflow: hidden;
}

.camera-img {
  position: absolute;
  top: 27px;
  right: 10px;
  width: 54px;
  height: 54px;
}

.avatar-dom {
  position: absolute;
  top: 110px;
  left: 31px;
  width: 214px;
  height: 214px;
  background: #E2E2E2;
  border-radius: 105px 105px 105px 105px;
  opacity: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;

  img {
    height: 100%;
    width: 100%;
    object-fit: cover;
  }
}

.personal-wditor-center {
  padding: 0 31px;
}

.personal-information {
  padding-top: 81px;
}

.family-name {
  font-size: 30px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: 600;
  color: #1F1F1F;
  line-height: 44px;
  word-wrap: break-word;
}

.occupation {
  height: 30px;
  line-height: 30px;
  font-size: 21px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #000000;
  margin-top: 6px;
  margin-bottom: 6px;
}

.personal-information-address {
  font-size: 18px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #585858;
  line-height: 21px;
  padding-right: 8px;
  word-wrap: break-word;

  .personal-information-city {
    width: 100%;
  }
}

.personal-information-age {
  height: 27px;
  font-size: 18px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #585858;
  line-height: 21px;
}

.contact-information {
  font-size: 18px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #0A66C2;

  .contact-information-item {
    height: 27px;

    img {
      width: 24px;
      height: 24px;
    }
  }
}

.edit-square-img {
  position: absolute;
  top: 27px;
  right: 0px;
  width: 31px;
  height: 31px;
}

.plate-four {
  height: 100px;
  background: #FFFFFF;
  border-radius: 10px;
  opacity: 1;
  border: 1px solid #E2E0DF;
  display: flex;
  align-items: center;
  //justify-content: ;
  padding: 0 26px;

  .plate-four-title {
    width: 100%;
    height: 31px;
    font-size: 26px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #030000;
    line-height: 30px;
  }
}

.avatar-dom {
  img {
    width: 68px;
    height: 68px;
  }

  .yes-avatar-dom-img {
    width: 216px;
    height: 216px;
    border-radius: 216px;
  }
}

.photos-videos-img {
  width: 40px;
  height: 40px;
}

.union {
  width: 44px;
  height: 44px;
}

.photo-wall-img {
  object-fit: cover;
}


// h5适配
@media (max-width: 420px) {
  .photos-videos-img {
    width: 34px;
    height: 34px;
  }
  .photos-videos {
    margin-top: 20px !important;
  }
  .h5-share-resume {
    display: flex;
    justify-content: space-between;
    margin-top: 16px;
  }
  .h5-template {
    width: 34px;
    height: 34px;
  }
  .h5-share {
    width: 293px;
    height: 34px;
  }
  .photo-ba {
    width: 100%;
    height: 120px;
    margin-top: 16px;

    .photo-wall-img {
      height: 120px;
    }

    .camera-img {
      width: 40px;
      height: 40px;
      top: 4px;
      right: 4px;
    }

  }
  .avatar-dom {
    width: 118px;
    height: 118px;
    top: 22px;
    left: 16px;
    border: 2px solid #fff;

    .avatar-dom-img {
      width: 42px;
      height: 42px;
    }

    .yes-avatar-dom-img {
      width: 118px;
      height: 118px;
    }
  }

  .personal-wditor-center {
    padding: 0;
  }
  .personal-information {
    padding-top: 60px;
    display: block;
  }
  .family-name {
    margin-top: 16px;
    font-size: 24px;
    line-height: 28px;
  }
  .edit-square-img {
    width: 24px;
    height: 24px;
    top: 40px;
    right: 6px;
  }
  .occupation {
    display: none;
  }
  .personal-information-age {
    font-size: 14px;
    height: 27px;
  }
  .contact-information {
    font-size: 14px;
    margin: 0;

    .contact-information-item {
      height: 24px;
      margin-bottom: 6px;

      img {
        width: 24px;
        height: 24px;
        margin-right: 6px;
      }
    }

  }

  .photos-videos {
    //width: 520px;
    display: block;
    margin-top: 44px;
    overflow: auto;
    white-space: nowrap;

    &::-webkit-scrollbar {
      width: 0 !important
    }

    .plate-four {
      height: 73px;
      background: #FFFFFF;
      border-radius: 4px 4px 4px 4px;
      opacity: 1;
      border: 1px solid #E2E0DF;
      padding: 0 30px;
    }

    .photos-videos-item {
      display: inline-block;
      width: 250px;
    }

    .union {
      width: 28px;
      height: 28px;
    }

    .plate-four-title {
      height: 27px;
      font-size: 18px;
      font-family: Inter-Semi Bold, Inter;
      font-weight: normal;
      color: #030000;
      line-height: 26px;
    }

  }

  .h5-w {
    width: 100%;
  }
}
</style>