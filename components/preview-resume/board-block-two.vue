<template>
  <div class="board-block-two"
       :class="`board-two-${type}`,`board-two-${data.class}`,data.title==='Education'?'mt-270px':'mt-70px'">
    <div class="title">
      <div class="title-content">{{ data.title }}</div>
      <div class="label"></div>
    </div>
    <slot name="head"></slot>
    <div v-if="data.title==='Education'">
      <div :class="`board-two-main-${type}`">
        <div
            v-for="(item,index) in props.list.list"
            :key="item.index"
            class="item"
        >
          <div class="name">
            {{ item.name }}
          </div>
          <div class="content-two">
            <span> Jan {{ item.startTime }}</span>
            <span v-if="item.endTime"> -<span v-if='item.endTime'> Jan {{ item.endTime }}</span></span>
          </div>
          <div class="content-three">Content</div>
          <div class="content-four" v-html="item.content.replace(/\n/g, '<br>')||''"></div>
        </div>
      </div>
    </div>
    <div v-else-if="data.title==='Experience'">
      <div :class="`board-two-main-${type} board-two-main-${data.title}`">
        <div
            v-for="(item,index) in props.list.list"
            :key="item.index"
            class="item"
        >
          <div class="right">
            <div class="name">
              {{ item.organization }}
            </div>
            <div class="content-two">
              <span> Jan {{ item.startTime }}</span>
              <span v-if="item.endTime"> -<span v-if='item.endTime'> Jan {{ item.endTime }}</span></span>
            </div>
            <div class="content-five">{{ item.title }}</div>

            <div class="content-three">Content</div>
            <div class="content-four"
                 v-html="item.content.replace(/\n/g, '<br>')||''"
            ></div>
          </div>

          <div class="left">
            <div class="content-six">{{ item.position }}</div>
            <div class="content-seven ">Summary</div>
            <div class="content-four " v-html="item.summary?.replace(/\n/g, '<br>')||''"></div>
          </div>
        </div>
      </div>
    </div>
    <slot name="main"></slot>
  </div>
</template>

<script setup>
let props = defineProps({
  list: Object,
  data: {
    type: Object,
    default: {
      title: '',
    },
  },
  type: String,
})
</script>

<style lang="scss" scoped>
.board-two-pink {
  background-color: #fff;
  border-radius: 10px;
  padding: 30px 44px 60px;
  box-shadow: 0px 6px 22px 0px rgba(114, 114, 114, 0.25);
  border-radius: 10px 10px 10px 10px;

  .title {
    position: relative;
    text-align: center;
    height: 58px;
    font-size: 48px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #000000;
    line-height: 56px;
    -webkit-background-clip: text;
  }

  .title-content {
    position: relative;
    z-index: 10;
    margin-bottom: 50px;
  }

  .board-two-main-pink {
    display: flex;
    flex-wrap: wrap;
  }

  .item {
    width: 50%;
    margin-top: 40px;
  }

  .label {
    height: 14px;
    background: #FFE4E4;
    border-radius: 0px 0px 0px 0px;
    opacity: 1;
    position: absolute;
    bottom: 12px;
  }

  .name {
    font-size: 30px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #000000;
    line-height: 35px;
    -webkit-background-clip: text;
  }

  .content-two {
    height: 28px;
    font-size: 12px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #999999;
    line-height: 28px;
    -webkit-background-clip: text;
  }

  .content-three {
    margin-top: 20px;
    height: 28px;
    font-size: 20px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #ECB875;
    line-height: 28px;
    -webkit-background-clip: text;
  }

  .content-four {
    //width: 381px;
    font-size: 16px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #666666;
    line-height: 28px;
    -webkit-background-clip: text;
  }

  .content-five {
    font-size: 24px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #ECB875;
    line-height: 28px;
    -webkit-background-clip: text;
  }

  .content-six {
    height: 28px;
    font-size: 14px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #616161;
    line-height: 28px;
    -webkit-background-clip: text;
  }

  .content-seven {
    height: 24px;
    font-size: 20px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #ECB875;
    line-height: 23px;
    -webkit-background-clip: text;
  }
}

.board-two-experience {
  .content-two {
    margin-bottom: 26px;
  }

  .content-six {
    margin-bottom: 26px;
  }

  //.item {
  //  width: 100%;
  //  display: flex;
  //  flex-wrap: wrap;
  //  .right{
  //    width: 50%;
  //  }
  //}
}

.board-two-education {
  position: relative;
  z-index: 10;

  .label {
    left: 282px;
    width: 281px;
  }
}

.board-two-experience {
  .label {
    left: 268px;
    width: 314px;
  }
}

.board-two-main-experience {
  display: block;

  .item {
    width: 100%;
    display: flex;
  }

  .content-seven {
    margin-top: 42px;
  }

  .content-three {
    margin-top: 32px;
  }
}

.board-two-job-experience {
  .label {
    left: 215px;
    width: 411px;
  }
}

.board-two-project-experience {
  position: relative;
  z-index: 2;

  .label {
    left: 178px;
    width: 483px;
  }
}

.board-two-volunteer-experience {
  .label {
    left: 146px;
    width: 557px;
  }
}

.board-two-custom-item {
  position: relative;
  z-index: 1;

  .label {
    width: 336px;
    left: 248px;
  }
}

.board-two-pink {
  .label {
    left: 163px !important;
    width: 483px !important;
  }
}
</style>
<style lang="scss" scoped>
@media (max-width: 400px) {
  .board-two-pink.board-two-education {
    top: -138px;
    margin-bottom: -138px;
  }
  .board-two-pink {
    padding: 15px;

    .title-content {
      font-size: 22px !important;
    }

    .label {
      left: -50% !important;
      margin-left: 187.5px !important;
      height: 10px !important;
      bottom: 16px !important;
      width: 272px !important;
      display: flex !important;
      justify-content: center;
    }

    .board-two-custom-item {
      .label {
        width: 176px !important;
      }
    }

    .board-two-main-pink {
      flex-direction: column !important;

      .item {
        margin-top: 20px;
        width: 100% !important;
      }

      .name {
        font-size: 20px;
      }

      .content-two {
        font-size: 12px !important;
        line-height: 13px !important;
        height: auto !important;
      }

      .content-three {
        margin-top: 0px !important;
        font-size: 18px !important;
      }

      .content-four {
        font-size: 14px !important;
      }
    }
  }

}
</style>