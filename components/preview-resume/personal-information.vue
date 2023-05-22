<template>
  <div class="personal-information " :class="`personal-information-${props.type}`">
    <!-- 基本信息 -->
    <div v-if="props.type==='blue'" class=" flex-1">
      <div class="name break-words">{{ props.myInfo?.firstName }}</div>
      <div class="name break-words">{{ props.myInfo?.lastName }}</div>
      <div class="interval-two"></div>
      <div class="address"> {{ props.myInfo?.city || '--' }}
        <!--                {{ props.myInfo?.province || '&#45;&#45;' }}-->
        {{ props.myInfo?.countryName || '--' }}
      </div>
      <div class="yearsold" v-if="props.myInfo?.birthday">{{ calculateHowOld(props.myInfo?.birthday) }}yearsold,
        {{ props.myInfo?.birthday }}
      </div>
      <div class="contact-information">
        <div class="flex" v-if="userCookieInfo.email&&props.showInfo?.Email">
          <img class="mr-16px mb-6px" src="@/assets/img/personal-center/icon/phone-iocn.png">
          <span class="email">{{ userCookieInfo.email }}</span>
        </div>
        <div class="flex" v-if="userCookieInfo.phone&&props.showInfo?.Phone">
          <img class="mr-16px" src="@/assets/img/personal-center/icon/mail-iocn.png">
          <span>{{ userCookieInfo.phone }}</span>
        </div>
      </div>
    </div>
    <div v-else-if="props.type==='pink'" class="" style="width: 100%">
      <div class="basic-information flex ">
        <div class="basic-information-div">
          <div class="name break-words">{{ props.myInfo?.firstName }}</div>
          <div class="name break-words">{{ props.myInfo?.lastName }}</div>
          <div class="address"> {{ props.myInfo?.city || '--' }}
            <!--                        {{props.myInfo?.province || '&#45;&#45;'}}-->
            {{ props.myInfo?.countryName || '--' }}
          </div>
          <div class="yearsold" v-if="props.myInfo?.birthday">{{
              calculateHowOld(props.myInfo?.birthday)
            }}yearsold, {{
              props.myInfo?.birthday
            }}
          </div>
          <div class="contact-information">
            <div class="flex" v-if="userCookieInfo.email&&props.showInfo?.Email">
              <img class="mr-16px mb-6px"
                   src="@/assets/img/personal-center/preview-resume/phone-iocn-pink.png">
              <span>{{ userCookieInfo.phone }}</span>
            </div>
            <div class="flex">
              <img class="mr-16px mb-6px"
                   src="@/assets/img/personal-center/preview-resume/mail-iocn-pink.png">
              <span v-if="userCookieInfo?.email">{{ userCookieInfo.email }}</span>
            </div>
          </div>
        </div>
        <!-- 头像 -->
        <img class="avatar-img" v-if="props.myInfo?.avatar" :src="props.myInfo?.avatar">
        <img class="avatar-img" v-else src="@/assets/img/no-head.png">
      </div>
      <div class="about-me-content">
        {{
          props.myInfo?.personalStrength
        }}
      </div>

      <div class="board-block-list">
        <BoardBlock :type="props.type" class="Languages-BoardBlock" v-if="props.showInfo?.Languages"
                    :list="Languages"
        ></BoardBlock>
      </div>
    </div>
    <div v-else-if="props.type==='purple'" class="purple">
      <div class="about-me-content">
        <div class="about-me-ornament1"></div>
        <div class="about-me-ornament2">
          <!--                    “-->
          <img class="decorate-three"
               src="@/assets/img/personal-center/preview-resume/blue-decorate-three.png">
        </div>
        <div class="about-me-world">
          {{
            props.myInfo?.personalStrength
          }}
        </div>
      </div>

            <div class="basic-information flex">
                <div>
                    <img class="avatar-img" v-if="props.myInfo?.avatar" :src="props.myInfo?.avatar">
                    <img class="avatar-img" v-else src="@/assets/img/no-head.png">
                </div>
                <div class="contact-information">
                    <div class="mr-6px">
                        <div>
                            <img class="mr-16px mb-6px"
                                 src="@/assets/img/personal-center/icon/phone-iocn.png">
                        </div>
                        <div v-if="userCookieInfo.phone&&props.showInfo?.Phone"
                             style=" writing-mode: tb-rl;padding-top: 5px">
                            {{ userCookieInfo.phone || '--' }}
                        </div>
                    </div>
                    <div class="">
                        <div><img class="mr-16px mb-6px"
                                  src="@/assets/img/personal-center/icon/mail-iocn.png"></div>
                        <div v-if="userCookieInfo.email&&props.showInfo?.Email"
                             style=" writing-mode: tb-rl;padding-top: 5px">
                            {{ userCookieInfo.email }}
                        </div>
                    </div>

        </div>
        <div class="information-left">
          <div class="name break-words" style="text-align: right">{{ props.myInfo?.firstName }}</div>
          <div class="name break-words purple-last-name" style="text-align: right">{{
              props.myInfo?.lastName
            }}
          </div>
          <div class="flex"
               style="margin-top: 18px;justify-content: space-between;">
            <div class="diagonal-line"></div>
            <div style="width: 235px;word-break: break-all;">
              <div class="address-item">
                {{ props.myInfo?.city || '--' }}
              </div>
              <!--                            <div class="address-item">-->

              <!--                                {{ props.myInfo?.province || '&#45;&#45;' }}-->
              <!--                            </div>-->
              <div class="address-item">
                {{ props.myInfo?.countryName || '--' }}

              </div>
              <div class="yearsold" v-if="props.myInfo?.birthday">
                {{ calculateHowOld(props.myInfo?.birthday) }}yearsold,
                {{
                  props.myInfo?.birthday
                }}
              </div>
            </div>
          </div>

        </div>
      </div>

      <BoardBlock :type="props.type" style="margin-top: 60px"
                  v-if="props.showInfo?.Languages&&props.showInfo['Languages']"
                  :list="props.myInfo?.languages"
      ></BoardBlock>
    </div>
    <div v-else-if="props.type==='orange'">
      <div class="basic-information flex">
        <div>
          <img class="avatar-img" v-if="props.myInfo?.avatar" :src="props.myInfo?.avatar">
          <img class="avatar-img" v-else src="@/assets/img/no-head.png">
        </div>

        <div class="information-right">
          <div class="name break-words">{{ props.myInfo?.firstName }}</div>
          <div class="name break-words">{{ props.myInfo?.lastName }}</div>
          <div class="address">
            {{ props.myInfo?.city || '--' }}
            <!--                {{ props.myInfo?.province || '&#45;&#45;' }}-->
            {{ props.myInfo?.countryName || '--' }}
          </div>
          <div class="yearsold" v-if="props.myInfo?.birthday">{{
              calculateHowOld(props.myInfo?.birthday)
            }}yearsold,
            {{ props.myInfo?.birthday }}
          </div>

          <div class="flex phone-email" style="justify-content: flex-start"
          >
            <div class="orange-center">
              <img class="mr-16px mb-6px img"
                   src="@/assets/img/personal-center/preview-resume/phone-iocn-orange.png">
            </div>
            <div class="orange-center" v-if="userCookieInfo.phone&&props.showInfo?.Phone">
              {{ userCookieInfo.phone || '--' }}
            </div>
          </div>
          <div class="flex phone-email" v-if="userCookieInfo.email&&props.showInfo?.Email"
               style="justify-content: flex-start">
            <div class="orange-img"><img class="mr-16px mb-6px img"
                                         src="@/assets/img/personal-center/preview-resume/mail-iocn-orange.png">
            </div>
            <div class="orange-center">{{ userCookieInfo.email }}</div>
          </div>

        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import BoardBlock from "@/components/preview-resume/board-block.vue";
import {useState} from "#app";
import {calculateHowOld} from "~/utils/util";

let props = defineProps({
    type: String,
    myInfo: Object,
    showInfo: Object
})


let Languages = {
    title: "Languages",
    style: ['pink', 'blue'].indexOf(props.type) > -1 ? "border-radius: 0px 0px 10px 10px;" : 'border-radius: 10px 10px 10px 10px;margin:20px 0;',
    list: [
        {txt: 'English'},
        {txt: 'English'},
        {txt: 'English'},
    ]
}

let userCookieInfo = useCookie('_userInformation')
</script>

<style lang="scss" scoped>
.Languages-BoardBlock {
  margin-left: -90px;
}

.personal-information {
  opacity: 1;

  .avatar-img {
    width: 228px;
    height: 228px;
    border-radius: 50%;
    opacity: 1;
  }

  .name {
    font-size: 64px;
    font-family: Almarai-Bold, Almarai;
    font-weight: bold;
    color: #000000;
    line-height: 60px;
    -webkit-background-clip: text;
  }
}

.personal-information-pink {
  display: flex;
  padding: 0px 90px 0 90px;
  background: linear-gradient(180deg, #FFFFFF 0%, rgba(255, 255, 255, 0) 100%);

  .name {
    width: 450px;
  }

  .basic-information {
    position: relative;
    left: -90px;
    padding: 144px 80px 0;
    width: calc(100% + 90px);
    border-radius: 40px 40px 0px 0px;
  }

  .profession {
    height: 60px;
    font-size: 42px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #000000;
    line-height: 49px;
    -webkit-background-clip: text;
    margin-bottom: 20px;
  }

  .address {
    font-size: 27px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #BF0C3C;
    line-height: 32px;
    -webkit-background-clip: text;
    margin-top: 26px;
  }

  .yearsold {
    height: 27px;
    font-size: 18px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #BF0C3C;
    line-height: 21px;
    -webkit-background-clip: text;
  }

  .contact-information {
    font-size: 22px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #D55A72;
    line-height: 25px;
    -webkit-background-clip: text;

    img {
      width: 30px;
      height: 30px;
    }
  }

  .about-me-content {
    position: relative;
    left: -46px;
    width: 801px;
    padding: 42px 90px 92px 90px;
    margin-top: 73px;
    background: linear-gradient(180deg, #FFE7E7 0%, rgba(254, 238, 238, 0) 100%);
    border-radius: 40px 40px 0px 0px;
    opacity: 1;
    font-size: 20px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #000000;
    word-break: break-word;
  }

  .board-block-list {
    > :nth-child(2) {
      margin-top: 40px;
      margin-bottom: 40px;
    }
  }

  .board-block {
    position: relative;
    left: -90px;
    width: calc(100% + 92px);
  }
}

.personal-information-purple {

  margin-top: -362px;

  .basic-information {
    position: relative;
    margin-top: 245px;
    justify-content: space-between;

    .contact-information {

      font-size: 22px;
      font-family: Inter-Regular, Inter;
      font-weight: 400;
      width: 70px;
      display: flex;
      //letter-spacing:10px;

      //color: #0A66C2;
      //line-height: 25px;
      //-webkit-background-clip: text;
      //transform: rotate(90deg);
      //margin-top: 245px;
      //position: absolute;
      //top: calc(-50% + 32px);
      //margin-left: 64px;
    }
  }

  .information-left {
    width: 430px;

    .diagonal-line {
      width: 4px;
      height: 220px;
      background: #333333;
      border-radius: 3px 3px 3px 3px;
      opacity: 1;
      transform: rotate(60deg);
      top: -56.7px;
      left: 96.3px;
      position: relative;
    }

    .address-item {
      font-size: 27px;
      font-family: Inter-Regular, Inter;
      font-weight: 400;
      color: #6D81E2;
      line-height: 32px;
      -webkit-background-clip: text;
      display: flex;
      justify-content: flex-end;
    }

    .yearsold {
      font-size: 18px;
      font-family: Inter-Regular, Inter;
      font-weight: 400;
      color: #6D81E2;
      line-height: 21px;
      display: flex;
      justify-content: flex-end;
      margin-top: 14px;
    }

    .profession {
      position: absolute;
      right: 0;
      font-size: 39px;
      font-family: Inter-Regular, Inter;
      font-weight: 400;
      color: #000000;
      line-height: 49px;
      display: flex;
      justify-content: flex-end;
    }
  }

  .about-me-content {
    width: 100%;
    min-height: 204px;
    height: fit-content;
    padding-bottom: 10px;
    background: #CCDAF2;
    border-radius: 10px 10px 10px 10px;
    opacity: 1;

    .about-me-ornament1 {
      width: 54px;
      height: 4px;
      background: #000000;
      border-radius: 3px 3px 3px 3px;
      opacity: 1;
      transform: rotate(90deg);
      position: relative;
      top: 107px;

    }

    .about-me-ornament2 {
      width: 47px;
      height: 77px;
      font-size: 64px;
      font-family: Inter-Semi Bold, Inter;
      font-weight: normal;
      color: #000000;
      line-height: 75px;
      -webkit-background-clip: text;
      position: relative;
      top: 45px;
      margin-left: 33px;
    }

    .about-me-world {
      width: calc(100% - 65px);
      min-height: 56px;
      font-size: 20px;
      font-family: Inter-Regular, Inter;
      font-weight: 400;
      color: #333333;
      line-height: 28px;
      -webkit-background-clip: text;
      margin-left: 65px;
    }
  }
}

.personal-information-orange {
  .basic-information {
    justify-content: space-around;

    .information-right {
      text-align: left;
      width: 526px;

      .name {
        line-height: 60px;
        margin-bottom: 10px;
      }

      .profession {
        font-size: 42px;
        font-family: Inter-Regular, Inter;
        font-weight: 400;
        color: #000000;
        line-height: 49px;
      }

      .address {
        font-size: 27px;
        font-family: Inter-Regular, Inter;
        font-weight: 400;
        color: #E77F1F;
        line-height: 32px;
        margin-top: 4px;
      }

      .phone-email {
        font-size: 22px;
        font-family: Inter-Regular, Inter;
        font-weight: 400;
        color: #DFA868;
        line-height: 25px;
        margin-top: 14px;
      }
    }
  }
}

.personal-information-pink {
  .profession {
    font-size: 40px;
  }
}

.img {
  width: 30px;
  height: 30px;
}

.purple {
  .contact-information {
    color: #6D81E2;
  }
}

.name {
  word-break: break-word;
}

.orange-img {
  display: contents;
}

.orange-center {
  line-height: 30px;
}

</style>
<style scoped lang="scss">
@media (max-width: 400px) {
  .purple-last-name {
    margin-bottom: 26px !important;
  }
  .personal-information-purple {
    margin-top: -132px !important;

    .about-me-content {
      min-height: 132px !important;
    }

    .about-me-ornament1 {
      top: 56px !important;
      left: -24px !important;
    }

    .about-me-ornament2 {
      top: -0px !important;
      left: -21px !important;
      height: 20px !important;

      img {
        height: 30px !important;
        width: 30px !important;
      }
    }

    .about-me-world {
      font-size: 14px !important;
      line-height: 21px !important;
      margin-left: 34px !important;
    }

    .basic-information {
      margin-top: 80px !important;

    }

    .profession {
      margin-top: -50px;
    }
  }
  .personal-information-orange {
    margin-top: 0 !important;

    .information-right {
      width: calc(100vw - 40px - 125px) !important;
      margin-left: 15px;
    }

    .address {
      font-size: 12px !important;
      line-height: 15px !important;
    }


  }
  .personal-information-pink {
    width: calc(100vw - 40px);
    padding: 0px !important;
    border-radius: 10px;

    .basic-information {
      left: 0 !important;
      width: 100%;
      padding: 60px 15px 15px !important;

      .basic-information-div {

        flex: 1;

        .name {
          width: auto !important;
          line-height: 35px;
          margin-bottom: 5px;
        }

        .address {
          margin-top: 0px !important;
          font-size: 14px !important;
          line-height: 13px;
        }

        .yearsold {
          font-size: 12px !important;
          line-height: 13px;
        }

        .profession {
          font-size: 16px !important;
          height: auto !important;
        }

        .contact-information {
          font-size: 12px;
          line-height: 12px;
        }
      }

    }

    .about-me-content {
      width: calc(100% - 28px);
      left: 0 !important;
      padding: 20px !important;
      margin-top: 10px !important;
      margin-left: 12px;
      margin-right: 12px;
      font-size: 14px !important;
      line-height: 21px !important;
      border-radius: 14px 14px 0px 0px;
    }

    .board-block-list {
      width: 100%;

      .Languages-BoardBlock {
        margin-left: 0 !important;
      }

      .board-content {
        margin-bottom: 6px;

        .board-block-item {
          margin-top: 0 !important;
        }
      }


    }
  }
  .personal-information {
    .avatar-img {
      width: 110px !important;
      height: 110px !important;
    }
  }
  .contact-information {
    font-size: 14px !important;
    color: #0A66C2;

    img {
      width: 24px !important;
      height: 24px !important;
      margin-right: 0;
      margin-bottom: 6px;
    }

  }
  .information-left {
    width: 45vw !important;

  }
  .name {
    font-size: 26px !important;
    color: #000000 !important;
    line-height: 21px !important;
    letter-spacing: 2px !important;
    margin-bottom: 6px;
  }
  .diagonal-line {
    height: 20vw !important;
    top: -5vw !important;
    left: 10vw !important;
  }
  .address-item, .yearsold {
    font-size: 12px !important;
    line-height: 13px !important;
  }
  .profession {
    font-size: 16px !important;
    line-height: 18px !important;
  }

  .orange-center {
    display: flex;
    align-items: center;
    line-height: 24px !important;
    max-width: 204px;
    word-break: break-word;
  }

  .phone-email {
    font-size: 14px !important;
    line-height: 12px !important;
    margin-top: 10px !important;

    img {
      width: 24px !important;
      height: 24px !important;
      margin-right: 6px;
      margin-bottom: 0;
    }
  }

  .personal-information-pink {
    .contact-information {
      img {
        margin-bottom: 0;
      }

      .flex {
        margin-bottom: 6px;
      }
    }
  }
}
</style>