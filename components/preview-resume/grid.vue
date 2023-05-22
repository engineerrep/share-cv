<template>
  <div class="grid">
    <div class="grid-title relative"
         :class="`grid-title-${type}`,props.gridTitle==='Photos' ?'mt-90px':'mt-70px'">
      <div class="grid-title-content">
        <div v-if="type==='orange'">
          <img v-if="props.gridTitle==='Photos'" class="grid-title-img"
               src='@/assets/img/personal-center/preview-resume/orange-photos.png'>
          <img v-if="props.gridTitle==='Videos'" class="grid-title-img"
               src='@/assets/img/personal-center/preview-resume/orange-videos.png'>
        </div>
        <span v-else>{{ props.gridTitle }}</span>
      </div>
      <div class="label" :class="`label-${type}`"></div>
    </div>
    <div class="grid-main" :class="`grid-main-${props.gridTitle}-${type}`">
      <div
          v-for="(item,index) in gridList"
          :key="index"
          class="grid-item relative"
      >

        <img v-if="props.gridTitle==='Photos'"
             class="grid-img"
             :src="infoOssBucketUrl+item.fileUrl">
        <img v-if="props.gridTitle==='Videos'"
             class="grid-img grid-img-videos"
             :src="infoOssBucketUrl+item.coverUrl">

        <span v-if="props.gridTitle==='Videos'" class="video-time">{{ filterTime(item.playTime) }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
let props = defineProps({
  type: String,
  gridList: Array,
  gridTitle: String,
  titleImg: String,
})
let infoOssBucketUrl = localStorage.getItem('infoOssBucketUrl')
const filterTime = (second) => {
  let timeStr = ''
  if (second < 60) {
    if (second < 10) {
      timeStr = "00:0" + second;
    } else {
      timeStr = "00:" + second;
    }
  } else {
    var min_total = Math.floor(second / 60);	// 分钟
    var sec = Math.floor(second % 60);	// 余秒
    if (min_total < 60) {
      if (min_total < 10) {
        if (sec < 10) {
          timeStr = "0" + min_total + ":0" + sec;
        } else {
          timeStr = "0" + min_total + ":" + sec;
        }
      } else {
        if (sec < 10) {
          timeStr = "00:" + min_total + ":0" + sec;
        } else {
          timeStr = "00:" + min_total + ":" + sec;
        }
      }
    } else {
      var hour_total = Math.floor(min_total / 60);	// 小时数
      if (hour_total < 10) {
        hour_total = "0" + hour_total;
      }
      var min = Math.floor(min_total % 60);	// 余分钟
      if (min < 10) {
        min = "0" + min;
      }
      if (sec < 10) {
        sec = "0" + sec;
      }
      timeStr = hour_total + ":" + min + ":" + sec;
    }
  }

  return timeStr

}
</script>

<style lang="scss" scoped>
.video-time {
  position: absolute;
  bottom: 10px;
  left: 20px;
  width: 89px;
  height: 46px;
  background: rgba(33, 33, 33, 0.5);
  border-radius: 10px 10px 10px 10px;
  opacity: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 25px;
  font-family: Plus Jakarta Sans-SemiBold, Plus Jakarta Sans;
  font-weight: 600;
  color: #FFFFFF;
  line-height: 29px;
}

.grid {

  .grid-main {
    display: flex;
    flex-wrap: wrap;
    margin-top: 32px;

    > :nth-child(3n) {
      margin-right: 0;
    }
  }

  .grid-item {
    width: 277px;
    height: 277px;
    margin-right: 30px;
    margin-bottom: 30px;
  }


  .grid-img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    border-radius: 10px;
  }
}

.grid-title {
  position: relative;
  width: 164px;
  height: 58px;
  font-size: 48px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: normal;
  color: #000000;
  line-height: 56px;

  .grid-title-content {
    position: relative;
    z-index: 10;
  }
}

.label {
  width: 230px;
  height: 22px;
  position: absolute;
  bottom: 0px;
}

.label-pink {
  background: #BF0C3C;
}

.label-purple {
  background: #6D81E2;
}

.grid-title-orange {
  color: #fff;
  height: 78px;
  background-size: 100% 100%;
  line-height: 78px;
  width: fit-content;
  padding: 0 113px 0 67px;
}

.grid-title-img {
  transform: scale(0.5);
  position: relative;
  top: -50px;
  left: -240px;
}

.grid-title-blue {

  height: 78px;
  line-height: 78px;
  width: fit-content;

  .grid-title-content {
    height: 58px;
    font-size: 48px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: 700;
    color: #000000;
    line-height: 56px;
    padding-left: 38px;
  }

  .label-blue {
    position: relative;
    left: 38px;
    width: 54px;
    height: 6px;
    background: #333333;
    border-radius: 6px;
    opacity: 1;
    margin-top: 14px;
    margin-bottom: 30px;
  }
}

.grid-main-Videos-blue {
  .grid-img-videos {
    &:nth-child(1), &:nth-child(2), &:nth-child(3), &:nth-child(5) {
      border-top-right-radius: 33%;
    }

    &:nth-child(4) {
      border-bottom-right-radius: 33%;
    }
  }
}

</style>
<style scoped lang="scss">
@media (max-width: 400px) {
  .grid-title-img {
    top: -10px;
    left: -100px;
    transform: scale(0.6);
  }
  .grid-title {
    //width: auto!important;
    height: 30px !important;
    margin-top: 45px !important;

    .grid-title-content {
      font-size: 22px !important;
      line-height: 30px !important;

    }


    .label {
      height: 13px !important;
      width: 130px !important;
    }

  }
  .grid-item {
    width: calc(33vw - 14px - 15px) !important; //105px !important;
    height: 105px !important;
    margin-right: 15px !important;
    margin-bottom: 15px !important;

  }
  .video-time {
    font-size: 12px !important;
    height: 15px !important;
    padding: 3px !important;
    width: auto !important;
    left: 7px !important;

  }
  .grid-title-orange {
    padding: 5px 55px 0 38px !important;
    height: 36px !important;
    line-height: 36px !important;

    .grid-title-img {
      left: -89px !important;
    }
  }
  .grid-title-blue {

    height: 78px;
    line-height: 78px;
    width: fit-content;

    .grid-title-content {
      height: 27px !important;
      font-size: 22px !important;
      font-weight: 700 !important;
      color: #000000 !important;
      line-height: 27px !important;
      padding-left: 10px !important;
    }

    .label-blue {
      left: 10px !important;
      width: 20px !important;
      height: 2px !important;
      background: #333333 !important;
      border-radius: 6px !important;
      opacity: 1 !important;
      margin-top: 5px !important;
      margin-bottom: 10px !important;
    }
  }
}


.grid-main {
  margin-top: 20px !important;
}

//.grid-main-Photos-blue{
//    .grid-img{
//        border-top-right-radius: 30px;
//    }
//
//}
</style>
