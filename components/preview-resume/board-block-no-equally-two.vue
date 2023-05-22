<template>
  <div class="board-block-no-equally-two mt-70px"
       :class="`board-two-${type}`,`board-two-${data.class}`">
    <div class="title">
      <div class="title-content">
        <div v-if="type==='orange'">
          <img v-if="data.title==='Education'" class="orange-img"
               src="@/assets/img/personal-center/preview-resume/orange-education.png">
          <img v-if="data.title==='Experience'" class="orange-img"
               src="@/assets/img/personal-center/preview-resume/orange-experience.png">
          <img v-if="data.title==='Job Experience'" class="orange-img"
               src="@/assets/img/personal-center/preview-resume/orange-job-experience.png">
          <img v-if="data.title==='Project Experience'" class="orange-img"
               src="@/assets/img/personal-center/preview-resume/orange-project.png">
          <img v-if="data.title==='Volunteer Experience'" class="orange-img"
               src="@/assets/img/personal-center/preview-resume/orange-volunteer.png">
          <img v-if="data.title==='Custom item'" class="orange-img"
               src="@/assets/img/personal-center/preview-resume/orange-custom.png">

        </div>
        <span v-else>
          {{ data.title }}
        </span>

      </div>
      <div class="label" :style="{ width: data.title==='Volunteer Experience'?'538px':''}"></div>
    </div>
    <slot name="head"></slot>
    <div v-if="data.title==='Education'">
      <div class="board-two-card-main" :class="`board-two-main-${type}`">
        <img v-if="type==='purple'" class="board-two-main-icon-purple"
             src="@/assets/img/personal-center/preview-resume/board-block-icon-purple.png">
        <div
            v-for="(item,index) in props.list.list"
            :key="index"
            class="item"
        >
          <div class="item-left"
               :class="(type==='purple'&&index!==(props.list.list.length-1))?'is-last-education':''">
            <div class="education-name">
              {{ item?.name }}
            </div>
            <div class="content-two left-time">
              <span> Jan {{ item?.startTime }}</span>
              <span v-if="item?.endTime"> - Jan {{ item?.endTime }}</span>
            </div>
          </div>
          <div class="item-rigth"
               :class="((type==='purple'||type==='orange')&&index!==(props.list.list.length-1))?'has-bottom-line':''">
            <div class="content-three">Content</div>
            <div class="content-four" v-html="item?.content?.replace(/\n/g, '<br>')||''"></div>
          </div>
        </div>
      </div>
    </div>
    <div v-else-if="data.title==='Experience'">
      <div class="board-two-card-main" :class="`board-two-main-${type} board-two-main-${data.title}`">
        <img v-if="type==='purple'" class="board-two-main-icon-purple"
             src="@/assets/img/personal-center/preview-resume/board-block-icon-purple.png">
        <div
            v-for="(item,index) in props.list.list"
            :key="index"
            class="item"
        >
          <div class="item-left"
               :class="(type==='purple'&&index!==(props.list.list.length-1))?'is-last-education':''">
            <div class="left-name">
              {{ item?.title }}
            </div>

            <div class="content-five describe">{{ item?.address }}</div>
            <div class="content-two left-time">
              <span> Jan {{ item?.startTime }}</span>
              <span v-if="item?.endTime"> - Jan {{ item?.endTime }}</span>
            </div>

          </div>
          <div class="item-rigth"
               :class="(type==='purple'&&index!==(props.list.list.length-1))?'has-bottom-line':''">

            <div class="left-title">
              {{ item?.title }}
            </div>
            <div class="content-three">Content</div>
            <div class="experience-content"
                 v-html="item?.content?.replace(/\n/g, '<br>')||''"
            ></div>
            <div class="content-seven ">Summary</div>
            <div class="content-four " v-html="item?.summary?.replace(/\n/g, '<br>')||''"></div>
          </div>

        </div>
      </div>
    </div>
    <div v-else-if="data.title==='Job Experience'">
      <div class="board-two-card-main" :class="`board-two-main-${type} board-two-main-${data.title}`">
        <img v-if="type==='purple'" class="board-two-main-icon-purple"
             src="@/assets/img/personal-center/preview-resume/board-block-icon-purple.png">
        <div
            v-for="(item,index) in props.list.list"
            :key="index"
            class="item"
        >
          <div class="item-left"
               :class="(type==='purple'&&index!==(props.list.list.length-1))?'is-last-education':''">
            <div class="left-name">
              {{ item?.companyName }}
            </div>

            <div class="content-five describe">{{ item?.companyIndustry }}</div>
            <div class="content-two left-time">
              <span> Jan {{ item?.startTime }}</span>
              <span v-if="item?.endTime"> - Jan {{ item?.endTime }}</span>
            </div>

          </div>

          <div class="item-rigth"
               :class="(type==='purple'&&index!==(props.list.list.length-1))?'has-bottom-line':''">
            <div class="content-three">Job content</div>
            <div class="job-position">{{ item?.jobPosition }}</div>
            <div class="experience-content"
                 v-html="item?.jobContent?.replace(/\n/g, '<br>')||''"
            ></div>

          </div>

        </div>
      </div>
    </div>
    <div v-else-if="data.title==='Project Experience'">

      <div class="board-two-card-main" :class="`board-two-main-${type} board-two-main-${data.title}`">
        <img v-if="type==='purple'" class="board-two-main-icon-purple"
             src="@/assets/img/personal-center/preview-resume/board-block-icon-purple.png">
        <div
            v-for="(item,index) in props.list.list"
            :key="index"
            class="item"
        >
          <div class="item-left"
               :class="(type==='purple'&&index!==(props.list.list.length-1))?'is-last-education':''">
            <div class="left-name">
              {{ item?.projectName }}
            </div>

            <div class="content-five describe">{{ item?.address }}</div>
            <div class="content-two left-time">
              <span> Jan {{ item?.startTime }}</span>
              <span v-if="item?.endTime"> - Jan {{ item?.endTime }}</span>
            </div>
            <div class="project-info-left project-info-left-web">
              <div class="skills-title" v-if="item?.Skills.length>0">
                Skills:
              </div>
              <div class="skill-list" v-if="item?.Skills.length>0">
                <div class="skill-item" v-for="(skill,index) in item?.Skills" :key="index">
                  {{ skill }}
                </div>
              </div>
              <div class="project-link" v-if="item?.projectLink">{{ item?.projectLinkPlatform }}:
                {{ item?.projectLink }}
              </div>

            </div>

          </div>

          <div class="item-rigth"
               :class="((type==='purple'||type==='orange')&&index!==(props.list.list.length-1))?'has-bottom-line':''">
            <div class="content-three">Project content</div>
            <div class="experience-content"
                 v-html="item?.projectContent.replace(/\n/g, '<br>')||''"
            ></div>
            <div class="content-seven ">Job content</div>
            <div class="job-position">{{ item?.jobPosition }}</div>
            <div class="content-four " v-html="item?.jobContent.replace(/\n/g, '<br>')||''"></div>

            <div class="project-info-left project-info-left-h5">
              <div class="skills-title">
                Skills:
              </div>
              <div class="skill-list">
                <div class="skill-item" v-for="(skill,index) in item?.Skills" :key="index">
                  {{ skill }}
                </div>
              </div>
              <div class="project-link">{{ item?.projectLinkPlatform }}:{{ item?.projectLink }}</div>

            </div>
          </div>

        </div>
      </div>
    </div>
    <div v-if="data.title==='Volunteer Experience'">
      <div class="board-two-card-main" :class="`board-two-main-${type}`">
        <img v-if="type==='purple'" class="board-two-main-icon-purple"
             src="@/assets/img/personal-center/preview-resume/board-block-icon-purple.png">
        <div
            v-for="(item,index) in props.list.list"
            :key="index"
            class="item"
        >
          <div class="item-left"
               :class="(type==='purple'&&index!==(props.list.list.length-1))?'is-last-education':''">
            <div class="education-name">
              {{ item?.title }}
            </div>
            <div class="content-two left-time">
              <span> Jan {{ item?.startTime }}</span>
              <span v-if="item?.endTime"> - Jan {{ item?.endTime }}</span>
            </div>
          </div>
          <div class="item-rigth"
               :class="(type==='purple'&&index!==(props.list.list.length-1))?'has-bottom-line':''">
            <div class="content-three">Content</div>
            <div class="content-four" v-html="item?.content?.replace(/\n/g, '<br>')||''"></div>
          </div>
        </div>
      </div>
    </div>
    <div v-else-if="data.title==='Custom item'">
      <div class="board-two-card-main" :class="`board-two-main-${type} board-two-main-${data.title}`">
        <img v-if="type==='purple'" class="board-two-main-icon-purple"
             src="@/assets/img/personal-center/preview-resume/board-block-icon-purple.png">
        <div
            v-for="(item,index) in props.list.list"
            :key="index"
            class="item"
        >
          <div class="item-left"
               :class="((type==='purple'||type==='orange')&&index!==(props.list.list.length-1))?'is-last-education':''">
            <div class="left-name">
              {{ item?.title }}
            </div>
            <div class="content-two left-time">
              <span> Jan {{ item?.startTime }}</span>
              <span v-if="item?.endTime"> - Jan {{ item?.endTime }}</span>
            </div>

          </div>
          <div class="item-rigth"
               :class="(type==='purple'&&index!==(props.list.list.length-1))?'has-bottom-line':''">
            <div class="content-three">Content</div>
            <div class="experience-content"
                 v-html="item?.content?.replace(/\n/g, '<br>')||''"
            ></div>
            <div class="content-seven " v-if="item?.summary">Summary</div>
            <div class="content-four " v-if="item?.summary"
                 v-html="item?.summary?.replace(/\n/g, '<br>')||''"></div>
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

<style scoped lang="scss">
.board-block-no-equally-two {
  .board-two-card-main {
    background: #fff;
    margin-top: 21px;
    padding: 66px 25px 73px 82px;
    position: relative;

    > :last-child {
      .item-rigth:after {
        border-bottom: 0 !important;
      }
    }

    .item {
      display: flex;

      .item-left {
        width: 293px;
        word-break: break-all;

        &.is-last-education {
          .education-time {
            margin-bottom: 20px;
          }
        }

        .education-name, .left-name {
          font-size: 30px;
          font-family: Inter-Medium, Inter;
          font-weight: 500;
          color: #000000;
          line-height: 35px;

        }

        .describe {
          font-size: 14px;
          font-family: Inter-Medium, Inter;
          font-weight: 500;
          color: #616161;
          line-height: 28px;
        }

        .left-time {
          font-size: 12px;
          font-family: Inter-Medium, Inter;
          font-weight: 500;
          color: #999999;
          line-height: 28px;
        }

        .project-info-left {
          margin-top: 130px;
          margin-bottom: 37px;


        }
      }

      .item-left, .item-rigth {
        .skill-title {
          font-size: 18px;
          font-family: Inter-Semi Bold, Inter;
          font-weight: normal;
          color: #000000;
          line-height: 28px;
        }

        .skill-list {
          flex: 1;
          display: flex;
          flex-wrap: wrap;

          .skill-item {
            // width: 63px;
            margin-top: 16px;
            padding: 0 20px;
            height: 35px;
            border-radius: 23px 23px 23px 23px;
            opacity: 1;
            font-size: 16px;
            line-height: 35px;
            text-align: center;
            margin-right: 20px;

          }
        }

        .project-link {
          font-size: 16px;
          font-family: Inter-Medium, Inter;
          font-weight: 500;
          color: #616161;
          line-height: 28px;
          margin-top: 34px;
        }
      }

      .item-rigth {
        flex: 1;
        position: relative;
        margin-left: 28px;
        padding-bottom: 40px;

        .content-three {
          font-size: 20px;
          font-family: Inter-Medium, Inter;
          font-weight: 500;
          line-height: 28px;
        }

        .content-four {
          margin-top: 4px;
          font-size: 16px;
          font-family: Inter-Regular, Inter;
          font-weight: 400;
          color: #666666;
          line-height: 28px;

        }

        .experience-content, .content-four {
          font-size: 16px;
          font-family: Inter-Regular, Inter;
          font-weight: 400;
          color: #666666;
          line-height: 28px;
          word-break: break-word;
        }

        .job-position {
          font-size: 12px;
          font-family: Inter-Medium, Inter;
          font-weight: 500;
          color: #999999;
          line-height: 28px;
          margin-bottom: 4px;
        }

        .content-seven {
          font-size: 20px;
          font-family: Inter-Medium, Inter;
          font-weight: 500;
          line-height: 23px;
          margin-top: 26px;
        }

        .left-title {
          font-size: 24px;
          font-family: Inter-Medium, Inter;
          font-weight: 500;
          color: #000000;
          line-height: 28px;
          margin-bottom: 10px;
        }

        &.has-bottom-line {
          .content-four, .experience-content {
            margin-bottom: 40px;
          }


        }

        &.has-bottom-line:after {
          content: '';
          border-bottom: solid 1px #FFC086;
          width: calc(100% - 28px);
          position: absolute;
          bottom: 20px;
        }

        &:before {
          content: '';
          border-left: solid 1px #DFA868;
          position: absolute;
          height: 100%;
          left: -28px;
        }

        &:after {
          content: '';
          border-bottom: solid 1px #DFA868;
          position: absolute;
          bottom: 20px;
          width: calc(100% - 53px);
        }
      }
    }
  }
}

.board-two-purple {
  .title {
    position: relative;

    height: 58px;
    font-size: 48px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #000000;
    line-height: 56px;

    .title-content {
      position: relative;
      z-index: 10;

      & span:after {
        content: '';
        width: fit-content;
        left: 0;
        position: absolute;
        bottom: 0;
        background: #6D81E2;
        height: 22px;
      }
    }

    .label {
      width: 463px;
      height: 22px;
      position: absolute;
      bottom: 0px;
      background: #6D81E2;
    }
  }

  .board-two-card-main {
    .board-two-main-icon-purple {
      position: absolute;
      top: 40px;
      left: 33px;
    }

    .item {
      .item-left, .item-rigth {
        .project-info-left {

        }

        .skill-list {
          .skill-item {
            background: #6D81E2;
            color: #FEFEFC;
          }
        }

        .project-link {
          color: #616161;
        }
      }

      .item-rigth {
        .content-three {
          color: #000000;
        }

        .content-four {
          color: #666666;
        }

        .experience-content, .content-four {
          color: #666666;
        }

        .job-position {
          color: #999999;
        }

        .content-seven {
          color: #000000;
        }

        .left-title {
          color: #000000;
        }

        &.has-bottom-line {
          .content-four, .experience-content {
          }
        }

        &.has-bottom-line:after {
          content: '';
          border-bottom: solid 1px #000;
          bottom: 44px !important;
        }

        &:before {
          content: '';
          border-left: solid 1px #000;
        }
      }

      .title-content {
        position: relative;
        z-index: 10;
      }

      .label {
        width: 411px;
        height: 22px;
        position: absolute;
        bottom: 0px;
        background: #6D81E2;
      }
    }

    .board-two-card-main {
      .board-two-main-icon-purple {
        position: absolute;
        top: 40px;
        left: 33px;
        width: 26px;
        height: 26px;
      }

      .item {
        .item-left, .item-rigth {
          .project-info-left {

          }

          .skill-list {
            .skill-item {
              background: #6D81E2;
              color: #FEFEFC;
            }
          }

          .project-link {
            color: #616161;
          }
        }

        .item-rigth {
          .content-three {
            color: #000000;
          }

          .content-four {
            color: #666666;
          }

          .experience-content, .content-four {
            color: #666666;
          }

          .job-position {
            color: #999999;
          }

          .content-seven {
            color: #000000;
          }

          .left-title {
            color: #000000;
          }

          &.has-bottom-line {
            .content-four, .experience-content {
            }
          }

          &.has-bottom-line:after {
            content: '';
            border-bottom: solid 1px #000;
          }

          &:before {
            content: '';
            border-left: solid 1px #000;
          }
        }


      }
    }

  }

  .board-two-orange {
    .title {
      color: #fff;
      background: url("@/assets/img/personal-center/preview-resume/orange-title-back-long.png") no-repeat;
      height: 78px;
      background-size: 100% 100%;
      line-height: 78px;
      width: fit-content;
      padding: 0 113px 0 67px;
      font-size: 48px
    }

    .board-two-card-main {
      box-shadow: 0px 6px 22px 0px rgba(114, 114, 114, 0.25);
      border-radius: 10px 10px 10px 10px;
      opacity: 1;
      border: 2px solid #E77F1F;

      .item {
        .item-left, .item-rigth {
          .project-info-left {
            .skill-list {
              .skill-item {
                background: #E77F1F;
                color: #FEFEFC;
              }
            }

            .project-link {
              color: #616161;
            }
          }
        }

        .item-rigth {
          .content-three {
            color: #DFA868;
          }

          .content-four {
            color: #666666;
          }

          .experience-content, .content-four {
            color: #666666;
          }

          .job-position {
            color: #999999;
          }

          .content-seven {
            color: #DFA868;
          }

          .left-title {
            color: #000000;
          }

          &.has-bottom-line {
            .content-four, .experience-content {
            }
          }

          &.has-bottom-line:after {
            content: '';
            border-bottom: solid 1px #DFA868;
          }

          &:before {
            content: '';
            border-left: solid 1px #DFA868;
          }
        }
      }
    }

    .board-two-orange {
      .title {
        color: #fff;
        background: url("@/assets/img/personal-center/preview-resume/orange-title-back-long.png") no-repeat;
        height: 78px;
        background-size: 100% 100%;
        line-height: 78px;
        width: 627px; //fit-content;
        padding: 0 113px 0 67px;
        font-size: 48px
      }

      .board-two-card-main {
        box-shadow: 0px 6px 22px 0px rgba(114, 114, 114, 0.25);
        border-radius: 10px 10px 10px 10px;
        opacity: 1;
        border: 2px solid #E77F1F;

        .item {
          .item-left, .item-right {
            .project-info-left {
              .skill-list {
                .skill-item {
                  background: #E77F1F;
                  color: #FEFEFC;
                }
              }

              .project-link {
                color: #616161;
              }
            }
          }

          .item-rigth {
            .content-three {
              color: #DFA868;
            }

            .content-four {
              color: #666666;
            }

            .experience-content, .content-four {
              color: #666666;
            }

            .job-position {
              color: #999999;
            }

            .content-seven {
              color: #DFA868;
            }

            .left-title {
              color: #000000;
            }

            &.has-bottom-line {
              .content-four, .experience-content {
              }
            }

            &.has-bottom-line:after {
              content: '';
              border-bottom: solid 1px #DFA868;
            }

            &:before {
              content: '';
              border-left: solid 1px #DFA868;
            }
          }
        }

      }
    }
  }
}

.board-two-main-icon-purple {
  width: 26px;
  height: 26px;
}

.orange-img {
  transform: scale(0.5);
  position: relative;
  left: -206px;
  top: 30px;
}

.board-two-main-orange {
  box-shadow: 0px 6px 22px 0px rgba(114, 114, 114, 0.25);
  border-radius: 10px 10px 10px 10px;
  opacity: 1;
  border: 1px solid #E77F1F;
}

.board-two-main-purple {
  box-shadow: 0px 6px 22px 0px rgba(114, 114, 114, 0.25);
  border-radius: 10px 10px 10px 10px;
  opacity: 1;
  border-right: 2px solid #6D81E2;
  border-bottom: 2px solid #6D81E2;
}

.board-two-orange {
  .skill-item {
    background: #E77F1F;
    color: #fff;
  }
}

.project-info-left-h5 {
  display: none;
}
</style>
<style scoped lang="scss">
@media (max-width: 400px) {
  .project-info-left-web {
    display: none;
  }
  .project-info-left-h5 {
    display: block;
  }
  .orange-img {
    left: -88px !important;
    top: 0 !important;
    height: 52px;
    transform: scale(0.6);
  }
  .board-two-education {
    .orange-img {
      left: -38px !important;
      transform: scale(1);
      width: 183px;
      height: 36px;
    }
  }
  .board-two-experience {
    .orange-img {
      left: -38px !important;
      transform: scale(1);
      width: 197px;
      height: 36px;
    }
  }
  .board-two-job-experience {
    .orange-img {
      left: -38px !important;
      transform: scale(1);
      width: 234px;
      height: 36px;
    }
  }
  .board-two-project-experience {
    .orange-img {
      left: -38px !important;
      transform: scale(1);
      width: 274px;
      height: 36px;
    }
  }
  .board-two-volunteer-experience {
    .orange-img {
      left: -38px !important;
      transform: scale(1);
      width: 303px;
      height: 36px;
    }
  }
  .board-two-custom-item {
    .orange-img {
      left: -38px !important;
      transform: scale(1);
      width: 211px;
      height: 36px;
    }
  }
  .board-block-no-equally-two {
    margin-top: 60px;

    .title {
      width: fit-content !important;
      height: 30px !important;

      .title-content {
        font-size: 22px !important;
        line-height: 30px !important;
      }

      .label {
        height: 13px !important;
        //width: 270px !important;
        width: auto !important;
      }
    }

    &.board-two-purple {
      .label {
        width: 270px !important;
      }
    }

    .board-two-card-main {
      padding: 20px 20px 20px 35px !important;

      img {
        left: 12px !important;
        top: 12px !important;
        width: 12px !important;
        height: 12px !important;
      }

      .education-name, .left-name {

        font-size: 20px !important;
        font-weight: 500 !important;
        color: #000000 !important;
        line-height: 23px !important;
      }
    }

    .item {
      flex-wrap: wrap !important;

      .item-left, .item-rigth {
        margin-top: 30px;

        .left-time {
          line-height: 15px !important;
        }

        .describe {
          font-size: 12px !important;
        }

        .project-info-left {
          margin-top: 10px !important;
          margin-bottom: 5px !important;


        }

        .skill-item {
          font-size: 12px !important;
          margin-right: 5px !important;
          height: 25px !important;
          line-height: 25px !important;
          padding: 0 5px !important;
        }

        .project-link {
          margin-top: 15px !important;
          font-size: 14px !important;
        }
      }

      .item-rigth {
        margin-left: 0 !important;

        &:before {
          content: '' !important;
          display: none !important;
          border: none !important;
        }

        .content-three {
          margin-top: 15px;
          font-size: 18px !important;
        }

        .left-title {
          font-size: 19px !important;
          line-height: 21px !important;
          margin-top: 20px;
          font-weight: 500;
        }
      }

      .has-bottom-line {
        .experience-content, .content-four {
          margin-bottom: 30px !important;
          font-size: 14px !important;
          line-height: 21px !important;
        }
      }
    }

    &.board-two-orange {
      .title {
        padding: 5px 55px 0 38px !important;
        height: 36px !important;
        line-height: 36px !important;

      }
    }

  }
  .board-two-main-Job.board-two-main-orange {
    .item {
      margin-bottom: 30px;
      position: relative;

      &:after {
        content: "";
        border-bottom: solid 1px #FFC086;
        width: calc(100% - 28px);
        position: absolute;
        bottom: -25px;
      }
    }

    > :last-child {
      &:after {
        border-bottom: 0;
      }
    }
  }
}
</style>