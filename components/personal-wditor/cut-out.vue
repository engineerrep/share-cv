<template>
  <div class="cut-out">
    <vue-cropper ref="cropper" v-bind="option"></vue-cropper>
    <div>
      <div class="cut-out-operate">
        <label class="btn hange-photo operate-dom" for="uploads">Change photo</label>

        <CVButton @click="uploadImage" class="apply-img operate-dom">
          <template #main>
      <span>
      Apply
      </span>
          </template>
        </CVButton>
      </div>
      <div>
        <input
            type="file"
            ref="uploadsRef"
            id="uploads"
            style="position: absolute; clip: rect(0 0 0 0)"
            accept="image/*"
            @change="uploadImg($event, 1)"/>
      </div>
    </div>
  </div>
</template>

<script setup>
import {VueCropper} from "vue-cropper";
import {apiUpdateAvatar, apiUpdateBackground} from "~/api/personal-wditor";
import axios from 'axios'
import {getCookie} from "~/utils/util";
import {ElMessage} from "element-plus";
import {useI18n} from "vue-i18n";
import {useState} from "nuxt/app";
import CVButton from '@/components/form/button.vue';
import {userStore} from "@/composables/store/user.js"

let userInfoStore = userStore()

const props = defineProps({
  type: String,
})
// 头像
let avatarImg = useState('_avatarImg', () => '')
// 照片墙


const {t,} = useI18n();
let cropper = ref()
let uploadsRef = ref()

let option = ref({
  img: props.type === 'avatar' ? (avatarImg.value ? avatarImg.value : '') : userInfoStore.photoWallImg, //裁剪图片的地址
  size: 1, //裁剪生成图片的质量(可选0.1 - 1)
  outputType: "jpeg", //裁剪生成图片的格式（jpeg || png || webp）
  info: true, //图片大小信息
  fixed: true, //是否开启截图框宽高固定比例
  fixedNumber: props.type === 'avatar' ? [1, 1] : [4, 1], //截图框的宽高比例
  autoCropWidth: props.type === 'avatar' ? 400 : 993,
  autoCropHeight: props.type === 'avatar' ? 400 : 248,
  canMove: true,// 上传图片是否可以移动
  centerBox: true,// 截图框是否限制在图片里(只有在自动生成截图框时才能生效)
  autoCrop: true,// 是否自动生成截图框
  fixedBox: true,// 是否截图框固定大小
  canMoveBox: false,// 是否拖动截图框
})

let pcSide = useState('_pcSide')
if (!pcSide.value) {
  option.value.autoCropHeight = 296
}

// 重置裁剪框
const recovery = () => {
  cropper.value.reload()
}

function dataURLtoFile(dataurl, filename) {
  let arr = dataurl.split(","),
      mime = arr[0].match(/:(.*?);/)[1],
      bstr = atob(arr[1]),
      n = bstr.length,
      u8arr = new Uint8Array(n);
  while (n--) {
    u8arr[n] = bstr.charCodeAt(n);
  }
  return new File([u8arr], filename, {type: mime});
}

let fileName = ref(null)
let isFile = ref(false)
let loading = useState("_dialog")
// 上传图片
const uploadImage = () => {
  if (isFile.value)
    cropper.value.getCropData((data) => {
      const files = dataURLtoFile(data, `${fileName.value}.png`);
      loading.value = true
      let formData = new FormData()
      formData.append("file", files)
      const config = useRuntimeConfig()
      const baseUrl = config.public.VITE_API_URL

      axios({
        method: 'post',
        url: `${baseUrl}/api/oss/upload`,
        data: formData,
        headers: {
          'Content-Type': 'multipart/form-data',
          'Authorization': getCookie("token")
        }
      }).then(res => {
        if (res.data.data && res.data.data.filePath) {
          if (props.type === 'baEdit') {
            apiUpdateBackground({background: res.data.data.filePath})
                .then(sonRes => {
                  userInfoStore.photoWallImg = localStorage.getItem('infoOssBucketUrl') + res.data.data.filePath
                  useState('_dialogBox').value = false;
                  loading.value = false
                })
          } else {
            apiUpdateAvatar({avatar: res.data.data.filePath})
                .then(sonRes => {
                  useState('_avatarImg').value = localStorage.getItem('infoOssBucketUrl') + res.data.data.filePath
                  useState('_dialogBox').value = false;
                  // localStorage.setItem('_avatarImg', JSON.stringify(localStorage.getItem('infoOssBucketUrl') + res.data.data.filePath))
                  loading.value = false
                })
          }
        }
      })
    });
}
//
// 更新图片
const uploadImg = (e, num) => {
  const form = new FormData();
  let file = e.target.files[0];
  form.append('file', file);
  if (file.type.indexOf('image') === -1) {
    console.log(4)
    ElMessage({
      type: 'error',
      message: t('personal-wditor.image type'),
    })
    return false;
  }
  let reader = new FileReader();
  reader.onload = (e) => {
    let data;
    if (typeof e.target.result === "object") {
      // 把Array Buffer转化为blob 如果是base64不需要
      data = window.URL.createObjectURL(new Blob([e.target.result]));
    } else {
      data = e.target.result;
    }
    option.value.img = data;
  };
  fileName.value = file.name
  reader.readAsArrayBuffer(file);
  isFile.value = true
};

</script>

<style>
@import "vue-cropper/dist/index.css";

.vue-cropper {
  background-color: #1c1c1c;
  background-image: url("assets/img/personal-center/black.png") !important;
}

.cut-out {
  width: 993px;
  height: 400px;
  position: absolute;
  top: 0px;
  left: 0px;
}

.apply-img {
  width: 156px;
  height: 42px;
  background: #0A66C2;
  border-radius: 39px 39px 39px 39px;
  opacity: 1;
  position: absolute;
  bottom: -66px;
  right: 32px;
  z-index: 30;
  font-size: 21px;
}

.hange-photo {
  width: 186px;
  height: 42px;
  border-radius: 39px 39px 39px 39px;
  opacity: 1;
  border: 2px solid #9E9E9E;
  position: absolute;
  bottom: -66px;
  right: 216px;
  z-index: 30;
  color: #828282;
  line-height: 39px;
  font-size: 21px;
}

// h5适配
@media (max-width: 400px) {
  .cut-out {
    width: 100%;
    height: 296px;
  }

  .cut-out-operate {
    display: flex;
    justify-content: space-evenly;
    align-items: center;
    padding: 32px 0;
    position: absolute;
    width: 100%;
  }

  .hange-photo {
    display: inline-block;
  }

  .apply-img {
    float: right;
  }

  .operate-dom {
    width: auto;
    font-size: 18px;
    position: static;
    padding: 0 18px;
    height: 40px;
    line-height: 40px;
  }
}
</style>