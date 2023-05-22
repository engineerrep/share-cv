<template>
    <div class="head-preview" :class="`head-preview-${type}`">
        <div class="main">
            <p @click="router.push('/personal-wditor')" class="hide-resum-preview cursor-pointer">
                <img class="return-img" src="@/assets/img/return-img.png">
            </p>
            <div class="download-btn" @click="downloadPreview">
                <img class="download-img" src="@/assets/img/download-img.png">
                <span class=" download-img-txt">Download</span>
            </div>
        </div>
        <div class="h5-main">
            <img src="@/assets/img/personal-center/preview-resume/download-operate.png">
            <div class="h5-download-operate">
                <div @click="router.push('/personal-wditor')" class="h5-preview-return"></div>
                <div @click="downloadPreview" class="h5-preview-download">
                </div>
            </div>
        </div>
        <div class="ba" :class="`ba-${type}`"></div>
    </div>
</template>

<script setup>
import {useRouter} from "vue-router";
import domToImage from "dom-to-image";

const router = useRouter();
let props = defineProps({
  type: String,
})

// 下载预览简历
const downloadPreview = () => {
  let dom = document.querySelector("#resumPreviewMain")
  domToImage.toPng(dom).then(function (dataUrl) {
    let link = document.createElement('a');
    link.download = 'Share-CV.jpeg';
    link.href = dataUrl;
    link.click();
    link.remove()
  })
};
</script>

<style lang="scss" scoped>
.head-preview {
  position: sticky;
  top: 0px;
  z-index: 50;
  height: 130px;

  .return-img {
    transform: scale(0.5);
  }

  .main {
    width: 100%;
    position: relative;
    z-index: 50;
    height: 130px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 54px;

  }

  .ba {
    position: absolute;
    top: 0;
    height: 130px;
    width: 994px;
    opacity: 80%;
  }

  .ba-blue {
    background-image: url("@/assets/img/personal-center/preview-resume/blue-head.png");
  }

  .ba-pink {
    background-image: url("@/assets/img/personal-center/preview-resume/pink-head.png");
  }

  .ba-purple {
    background-image: url("@/assets/img/personal-center/preview-resume/purple-head.png");
  }

  .ba-orange {
    background-image: url("@/assets/img/personal-center/preview-resume/head-preview-orange.png");
  }

  .download-img-txt {
    height: 25px;
    font-size: 21px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #FFFFFF;
    line-height: 25px;
    -webkit-background-clip: text;
    margin-left: 10px;
  }

  .download-img {
    display: inline-block;
    width: 28px;
    height: 28px;
  }

  .download-btn {
    padding: 8px 30px;
    height: auto;
    background-color: #005db6;
    border-radius: 42px;
    cursor: pointer;
  }
}

.head-preview-pink {
  height: 0;
}

.h5-main {
  display: none;
}

@media (max-width: 420px) {
  .head-preview {
    height: 41px;
    top: 10px;

    .main, .ba {
      display: none;
    }

    .h5-main {
      display: block;
      //margin: 0 20px;
    }
  }
  .head-preview-orange {
    .h5-main {
      display: block;
      margin: 0;

      img {
        padding: 0 20px;
      }
    }
  }
  .head-preview-purple {
    .h5-main {
      img {
        padding: 20px;
      }
    }
  }
  .h5-download-operate {
    position: absolute;
    top: 0;
    width: 100%;
    height: 41px;
    display: flex;
    justify-content: space-between;

    .h5-preview-return {
      width: 41px;
      height: 41px;
    }

    .h5-preview-download {
      position: absolute;
      right: 44px;
      width: 166px;
      height: 41px;
    }
  }
}
</style>