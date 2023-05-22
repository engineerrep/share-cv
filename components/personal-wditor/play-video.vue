<template>
  <div class="video">
    <div class="video-main">
      <video controls autoplay="" name="media">
        <source
            :src="urlVideo"
            type="video/mp4">
      </video>
      <img @click="isPlayVideo=false" class="delete-img" src="@/assets/img/personal-center/icon/closure-video-iocn.png">
    </div>
  </div>
</template>

<script setup>
import {useState} from "#app";

let props = defineProps({
  urlVideo: String,
})
console.log('------', props.urlVideo)

let isPlayVideo = useState("_isPlayVideo")
onMounted(() => {
  let dom = document.querySelector('#uploadMain')
  let uploadHead = document.querySelector('#uploadHead')
  let uploadTail = document.querySelector('#uploadTail')
  console.log(uploadHead)
  if (document.body.clientWidth > 420) {
    dom.style.height = '558px';
    dom.style.width = '990px';
    dom.style.maxWidth = '990px';
  } else {
    dom.style.height = '320px';
    dom.style.width = 'calc(100% - 16px)';
    dom.style.maxWidth = ' calc(100% - 16px)';
    dom.style.background = 'transparent';
    uploadHead.style.display = 'none'
    uploadTail.style.display = 'none'
  }

})
onBeforeUnmount(() => {
  let dom = document.querySelector('#uploadMain')
  let uploadHead = document.querySelector('#uploadHead')
  let uploadTail = document.querySelector('#uploadTail')
  if (document.body.clientWidth > 420) {
    dom.style.height = 'auto';
    dom.style.width = '736px';
    dom.style.maxWidth = '736px';
  } else {
    dom.style.height = 'auto';
    dom.style.width = 'calc(100% - 32px);';
    dom.style.maxWidth = '736px';
    dom.style.background = '#FFFFFF';
    uploadHead.style.display = 'flex'
    uploadTail.style.display = 'flex'
  }
})
</script>

<style lang="scss" scoped>
.video {
  position: absolute;
  top: 50%;
  left: 50%;
  z-index: 100;
  max-height: 100%;
  transform: translate(-50%, -50%);
  width: 990px;
  height: 558px;
  display: flex;
  align-items: center;
  background-color: #fff;

  .video-main {
    width: 990px;
    height: 558px;
    background-color: rgba(0, 0, 0, 0.8);
  }

  video {
    width: 100%;
    height: 100%;
    //object-fit: cover;
  }
}

.delete-img {
  position: absolute;
  top: 22px;
  right: 17px;
  width: 40px;
  height: 40px;
  cursor: pointer;
}

@media (max-width: 420px) {
  .video {
    position: relative;
    width: calc(100% - 4px);
    height: 320px;
    z-index: 100;

    video {
      width: 100%;
      height: 100%;
    }

    .video-main {
      height: 320px;
      overflow: hidden;
    }

    .delete-img {
      top: 52px;
    }
  }
}
</style>