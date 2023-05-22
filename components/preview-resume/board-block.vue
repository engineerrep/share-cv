<template>
  <div>
    <div class="board-block" :class="`board-block-${type}`,props.list?.title"
         :style="props.list?.style"
    >
      <div class="board-title">
        <img v-if="type==='purple'" class="board-block-icon-purple"
             src="@/assets/img/personal-center/preview-resume/board-block-icon-purple.png">
        {{ props.list?.title }}
      </div>
      <div class="board-content">
        <div
            class="board-block-item"
            v-for="(item,index) in props.list?.list" :key="index">
          {{ item.txt }}
        </div>
        <div v-if="type=='purple'" class="purple-style-line"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
let props = defineProps({
  type: String,
  list: Object,
})

</script>

<style lang="scss" scoped>
.board-block {
  padding: 48px 42px;
  display: flex;

  .board-title {
    width: 270px;
    height: 36px;
    font-size: 30px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    line-height: 35px;
  }

  .board-content {
    flex: 1;
    display: flex;
    flex-wrap: wrap;
    margin-top: 16px;

  }

  .board-block-item {
    padding: 6px 16px;
    border-radius: 22px;
    margin-right: 20px;
    font-size: 20px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    line-height: 23px;
    margin-top: 20px;
  }
}

.board-block-pink.Interests {
  margin-top: 20px;
  position: relative;
  z-index: 10;
}

.personal-information-pink {
  .Languages-BoardBlock {
    width: 893px;
    position: relative;
  }
}

.board-block-pink {
  background-color: #fff;
  box-shadow: 0px 6px 22px 0px rgba(114, 114, 114, 0.25);
  border-radius: 10px 10px 10px 10px;

  .board-title {
    color: #BF0C3C;
  }

  .board-content {
  }

  .board-block-item {
    background: #BF0C3C;
    color: #FFFFFF;
  }
}

.board-block-purple {
  background-color: #fff;
  box-shadow: 0px 6px 22px 0px rgba(114, 114, 114, 0.25);
  border-radius: 10px 10px 10px 10px;
  opacity: 1;
  border-right: 2px solid #6D81E2;
  border-bottom: 2px solid #6D81E2;

  .board-title {
    color: #6D81E2;
    width: 223px;

    .board-block-icon-purple {
      width: 26px;
      height: 26px;
      position: relative;
      left: -21px;
      top: -10px;

    }
  }

  .board-content {


    height: 100%;
    width: 1px;
    padding-left: 47px;
    position: relative;

    &:before {
      content: "";
      border-left: 1px solid #000000;
      height: calc(100% - 48px);
      position: absolute;
      left: 0;

    }

    .purple-style-line {
      width: 100%;
      border-bottom: solid 1px #000;
      margin-top: 49px;
    }
  }

  .board-block-item {
    background: #6D81E2;
    color: #FFFFFF;
  }
}

.about-me-content + div {
  margin-top: 60px;
}

.board-block-orange {
  background-color: #fff;
  box-shadow: 0px 6px 22px 0px rgba(114, 114, 114, 0.25);
  border-radius: 10px 10px 10px 10px;
  opacity: 1;
  border: 1px solid #E77F1F;
  margin-top: 20px;

  .board-title {
    color: #E77F1F;
  }

  .board-content {
  }

  .board-block-item {
    background: #E77F1F;
    color: #FFFFFF;
    word-break: break-word;
  }
}

.board-block-item {
  word-break: break-word;
}


</style>
<style scoped lang="scss">
@media (max-width: 400px) {
  .board-block {
    padding: 20px !important;

    .board-title {
      font-size: 14px !important;
      line-height: 38px !important;
      width: 30vw;
      font-weight: 700;

      .board-block-icon-purple {
        left: -7px !important;
        top: -4px !important;
        width: 12px !important;
        height: 12px !important;
      }
    }

    .board-content {
      padding-left: 12px !important;
      margin-bottom: 16px;

      .purple-style-line {
        margin-top: 12px !important;
      }

      &:before {
        content: '';
        height: calc(100% - 12px) !important;
      }
    }

    .board-block-item {
      padding: 6px 16px;
      border-radius: 22px;
      margin-right: 5px;
      font-size: 14px;
      line-height: 16px;
      margin-top: 5px;
      display: inline-block;
    }
  }
  .board-content {
    display: block !important;
  }
  .board-block-orange {
    padding: 10px !important;

    .board-title {
      width: calc(30vw - 20px) !important;
      font-size: 16px !important;
      font-weight: bold;
    }

    .board-content {
      padding-left: 0 !important;

      &:before {
        display: none;
      }
    }
  }
}
</style>