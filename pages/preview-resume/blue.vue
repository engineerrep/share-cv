<template>
  <div class="dom-center blue">
    <div class="resum-preview">
      <DownloadOperate type="blue" class="DownloadOperate"></DownloadOperate>
      <div id="resumPreviewMain">
        <div class="resum-preview-main" ref="myDivRef">
          <!-- 个人实力 -->
          <div class="about-me relative" v-if="showInfo['About me']">
            <div class="about-me-hint">
              <img class="decorate-three"
                   src="@/assets/img/personal-center/preview-resume/blue-decorate-three.png">
            </div>
            <div class="about-me-content">
              <div class="about-me-border"></div>
              <div class="break-word">{{ myInfo?.personalStrength }}</div>
            </div>
            <img class="about-me-img" src="@/assets/img/personal-center/about-me-img.png">
          </div>
          <!-- 个人信息 -->
          <div class="personal-information flex">
            <!-- 头像 -->
            <div class="avatar">
              <img
                  v-if="myInfo?.avatar"
                  :src="myInfo?.avatar"
                  alt=""
                  class="avatar-img"
              />
              <img
                  v-else
                  src="@/assets/img/no-head.png"
                  alt=""
                  class="avatar-img"
              />
            </div>
            <!-- 基本信息 -->
            <div class="basic-information">
              <div class="name break-words">{{ myInfo?.firstName }}</div>
              <div class="name break-words">{{ myInfo?.lastName }}</div>
              <div class="interval-two"></div>
              <div class="address">{{ myInfo?.city || '--' }} {{ myInfo?.province || '--' }}
                {{ myInfo?.countryName || '--' }}
              </div>
              <div class="yearsold" v-if="myInfo?.birthday">{{
                  calculateHowOld(myInfo?.birthday)
                }}yearsold,
                {{ myInfo?.birthday }}
              </div>
              <div class="contact-information">
                <div class="flex" v-if="showInfo.Phone">
                  <img class="mr-16px mb-6px" src="@/assets/img/personal-center/icon/phone-iocn.png">
                  <span>{{ myInfo?.phone }}</span>
                </div>
                <div class="flex" v-if="showInfo.Email">
                  <img class="mr-16px" src="@/assets/img/personal-center/icon/mail-iocn.png">
                  <div class="orange-center ">{{ myInfo?.email || '--' }}</div>
                </div>
              </div>
            </div>
          </div>

          <!-- 语言-爱好-技能 -->
          <div class="flex  plate-twelve plate-seven ">
            <!-- 语言 -->
            <div class="plate-six languages" v-if="showInfo['Languages']">
              <div class="six-title">
                <img class="title-img"
                     src="@/assets/img/personal-center/preview-resume/blue-languages-iocn.png">
              </div>
              <div class="six-content">
                             <span v-for="(item,index) in myInfo?.languages?.list"
                                   :key="index"
                                   class="six-item"
                                   :style="myInfo?.languages">{{ item.txt }}
                             </span>
              </div>
            </div>
            <!-- 爱好 -->
            <div class="plate-six six-border interests" v-if="showInfo['Interests']">
              <div class="six-title">
                <img class="title-img"
                     src="@/assets/img/personal-center/preview-resume/blue-interests-iocn.png">
              </div>
              <div class="six-content">
              <span
                  v-for="(item,index) in Interests.list"
                  :key="index"
                  class="six-item"
                  :style="Interests.style"
              >{{ item.txt }}</span>
              </div>
            </div>
            <!-- 技能 -->
            <div class="plate-six six-border skills" v-if="showInfo['Skills']">
              <div class="six-title">
                <img class="title-img skills-img"
                     src="@/assets/img/personal-center/preview-resume/blue-skills-iocn.png">
              </div>
              <div class="six-content">

              <span

                  v-for="(item,index) in Skills.list"
                  :key="index"
                  class="six-item"
                  :style="Skills.style"
              >{{ item.txt }}</span>
              </div>
            </div>
          </div>
          <Grid :gridList="photosList" v-if="photosList.length>0" :type="type" gridTitle="Photos"/>

          <Grid :gridList="videosList" v-if="videosList.length>0" :type="type" gridTitle="Videos"/>


          <!-- 教育 -->
          <div class="plate-seven" v-if="educationList.list.length>0">
            <div class="seven-title">Education</div>
            <div class="interval-three"></div>
            <div class="education-main seven-main relative">
              <div>
                <div class="mark-two"></div>

                <div v-for="(item,index) in educationList.list" :key="index">
                  <div class="relative">
                    <div class="seven-subtitle"> {{ item.name }}</div>
                    <div class="mark"></div>
                  </div>
                  <div class="seven-content-two">
                    Jan {{ item.startTime }}
                    <span v-if='item.endTime'>- Jan {{ item.endTime }}</span>
                  </div>
                  <div class="seven-title-two interval-five">Content</div>
                  <div class="seven-content-three">{{ item.content }}</div>
                </div>
                <img class="education-img" src="@/assets/img/personal-center/icon/education-img.png">
              </div>
            </div>
          </div>

          <!-- 经验 -->
          <div class="preview-experience plate-seven" v-if="experienceDaList.list.length>0">
            <div class="seven-title">Experience</div>
            <div class="interval-three"></div>
            <div class=" seven-main relative">
              <div>
                <div class="mark-two"></div>

                <div v-for="(item,index) in experienceDaList.list" :key="index">
                  <div class="relative">
                    <div class="seven-subtitle"> {{ item.organization }}</div>
                    <div class="mark"></div>
                  </div>
                  <div class="seven-content-four">{{ item.address }}</div>
                  <div class="seven-content-two">
                    Jan {{ item.startTime }}
                    <span v-if='item.endTime'>- Jan {{ item.endTime }}</span>
                  </div>
                  <div class="seven-content-five interval-four">{{ item.title }}</div>
                  <div class="seven-title-two">Content</div>
                  <div class="seven-content-three interval-six">{{ item.content }}</div>
                  <div v-if="item.summary">
                    <div class="seven-title-two">Summary</div>
                    <div class="seven-content-three">{{ item.summary }}</div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- 工作经验 -->
          <div class="preview-job-experience plate-seven" id="JobExperience"
               v-if="testJobExperience.list.length>0">
            <div class="seven-title">Job Experience</div>
            <div class="interval-three"></div>

            <div class="job-eperience-main seven-main ">
              <div>
                <div class="mark-two"></div>

                <div v-for="(item,index) in testJobExperience.list" :key="index">
                  <div class="relative">
                    <div class="seven-subtitle"> {{ item.companyName }}</div>
                    <div class="mark"></div>
                  </div>
                  <div class="seven-content-two">
                    <span v-if="item.startTime">Jan {{ item.startTime }}</span>
                    <span v-if='item.endTime'>- Jan {{ item.endTime }}</span>
                  </div>
                  <div class="seven-title-two interval-five">Content</div>
                  <div class="seven-content-three"
                       v-html="item.jobContent.replace(/\n/g, '<br>')||''"></div>
                </div>
              </div>
            </div>
            <img class="job-experience-img" src="@/assets/img/personal-center/job-experience-img.png">

          </div>

          <!-- 项目经验 -->
          <div class="preview-experience plate-seven" id="ProjectExperience"
               v-if="ProjectExperiencelist.list.length>0">
            <div class="seven-title">Project Experience</div>
            <div class="interval-three"></div>
            <div class=" seven-main relative">
              <div>
                <div class="mark-two"></div>

                <div v-for="(item,index) in ProjectExperiencelist.list" :key="index">
                  <div class="relative">
                    <div class="seven-subtitle"> {{ item.projectName }}</div>
                    <div class="mark"></div>
                  </div>
                  <div class="seven-title-two interval-nine">Project content</div>
                  <div class="seven-content-two">
                    Jan {{ item.startTime }}
                    <span v-if='item.endTime'>- Jan {{ item.endTime }}</span>
                  </div>
                  <div class="seven-title-two">Content</div>
                  <div class="seven-content-three interval-six">{{ item.jobContent }}</div>

                  <div v-if="item.summary">
                    <div class="seven-title-two">Job content</div>
                    <div class="seven-content-two">Job content</div>
                    <div class="seven-content-three">{{ item.summary }}</div>
                  </div>

                  <div class="seven-title-two interval-eight " v-if="item.Skills.length>0">Skills:</div>
                  <div class="flex" style="flex-wrap: wrap;" v-if="item.Skills.length>0">
                    <div v-for="(item,index) in item.Skills" :key="index"
                         class="skills-item"
                    >{{ item }}
                    </div>
                  </div>
                  <div class="seven-content-four interval-nine" v-if="item.projectLink">
                    {{ item.projectLinkPlatform }}: {{ item.projectLink }}
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- 志愿者 -->
          <div class="preview-volunteer-experience plate-seven" id="VolunteerExperience"
               v-if="VolunteerExperiencelist.list.length>0">
            <div class="seven-title">Volunteer Experience</div>
            <div class="interval-three"></div>

            <div class=" seven-main ">
              <div>
                <div class="mark-two"></div>

                <div v-for="(item,index) in VolunteerExperiencelist.list" :key="index">
                  <div class="relative">
                    <div class="seven-subtitle"> {{ item.title }}</div>
                    <div class="mark"></div>
                  </div>
                  <div class="seven-content-two">
                    Jan {{ item.startTime }}
                    <span v-if='item.endTime'>- Jan {{ item.endTime }}</span>
                  </div>
                  <div class="seven-title-two interval-five">Content</div>
                  <div class="seven-content-three seven-content-six">{{ item.content }}</div>
                </div>
              </div>
            </div>
          </div>

          <!-- 自定义属性 -->
          <div class="preview-custom-tem plate-seven" id="customItem" v-if="customItemlist.list.length>0">
            <div class="seven-title">Custom item</div>
            <div class="interval-three"></div>
            <div class=" seven-main relative">
              <div>
                <div class="mark-two"></div>

                <div v-for="(item,index) in customItemlist.list" :key="index">
                  <div class="relative">
                    <div class="seven-subtitle"> {{ item.title }}</div>
                    <div class="mark"></div>
                  </div>
                  <div class="seven-content-two">
                    Jan {{ item.startTime }}
                    <span v-if='item.endTime'>- Jan {{ item.endTime }}</span>
                  </div>
                  <div class="seven-title-two interval-five">Content</div>
                  <div class="seven-content-three interval-six seven-content-six">{{
                      item.content
                    }}
                  </div>
                  <div v-if="item.summary">
                    <div class="seven-title-two">Summary</div>
                    <div class="seven-content-three seven-content-six">{{ item.summary }}</div>
                  </div>
                </div>
              </div>
            </div>
          </div>


        </div>
        <!-- 尾部 -->
        <Footer :type="type"></Footer>
      </div>
    </div>
  </div>
</template>

<script setup>
import Footer from "~/components/preview-resume/footer.vue";
import {NButton, NSpace} from "naive-ui";
import {useRouter} from "vue-router";
import {calculateHowOld} from "~/utils/util";
import domToImage from "dom-to-image";
import DownloadOperate from '@/components/preview-resume/download-operate.vue'
import {
  apiCustomAttributeList,
  apiExperienceList,
  apiPhotoList,
  apiUserCompanyList,
  apiUserEducationList, apiUserProjectList,
  apiVideoList, apiVolunteerList
} from "~/api/personal-wditor";
import {apUserInfo} from "~/api/user";
import Grid from "~/components/preview-resume/grid.vue";

let type = 'blue'
const router = useRouter();
let url = ref()
let gridList = ref([
  {txt: 'English',},
  {txt: 'English',},
  {txt: 'English',},
  {txt: 'English',},
  {txt: 'English',},
  {txt: 'English',},
  {txt: 'English',},
])
let photosList = ref([])
let videosList = ref([])

let Skills = ref({
  title: "Skills",
  style: "margin-right: 14px;",
  list: []
})

let Interests = ref({
  title: "Interests",
  style: "margin-right: 14px;",
  list: []
})
let testducationList = ref({
  title: "Education",
  class: "education",
})
// 教育-数据
let educationList = ref({
  type: 'Education',
  list: []
})
let experienceData = ref({
  title: "Experience",
  class: "experience",
})
let experienceDaList = ref({
  type: 'Experience',
  list: []
})
// 教育-数据
let testexperienceList = {
  type: 'Education',
  list: []
}
let testJobExperienceMsg = ref({
  title: "Job Experience",
  class: "job-experience",
})

// 工作经验-数据
let testJobExperience = ref({
  type: 'JobExperience',
  list: []
})
// 项目经验
let testProjectExperienceData = ref({
  title: "Project Experience",
  class: "project-experience",
})

// 项目经验-数据
let ProjectExperiencelist = ref({
  type: 'ProjectExperience',
  list: []
})

// 志愿者
let VolunteerExperienceData = ref({
  title: "Volunteer Experience",
  class: "volunteer-experience",
})

// 志愿者-数据
let VolunteerExperiencelist = ref({
  type: 'VolunteerExperience',
  list: []
})

// 志愿者
let customItemData = ref({
  title: "Custom item",
  class: "custom-item",
})

// 志愿者-数据
let customItemlist = ref({
  type: 'CustomItem',
  list: []
})
let personalInformation = ref({})
let publicList = []
let getInfoFuncObj = {
  "Photos": async () => {
    await apiPhotoList().then(res => {
      if (res.data.list?.length > 9) {
        res.data.list.length = 9
      }
      photosList.value = res.data.list || []
    })
  },
  "Videos": async () => {
    await apiVideoList().then(res => {
      if (res.data.list?.length > 6) {
        res.data.list.length = 6
      }
      videosList.value = res.data.list || []
    })
  },
  "Education": async () => {
    await apiUserEducationList().then(res => {
      educationList.value.list = res.data.list || []
    })
  },
  'Experience': async () => {
    await apiExperienceList().then(res => {
      experienceDaList.value.list = res.data.list || []
    })

  },
  "Job Experience": async () => {
    await apiUserCompanyList().then(res => {
      testJobExperience.value.list = res.data.list || []
    })
  },
  "Project Experience": async () => {
    await apiUserProjectList().then(res => {
      ProjectExperiencelist.value.list = res.data.list?.map(val => {
        return {...val, Skills: val.projectSkills ? val.projectSkills?.split(',') : []}
      }) || []
    })
  },
  "Custom item": async () => {
    await apiCustomAttributeList().then(res => {
      customItemlist.value.list = res.data.list || []
    })
  },
  "Volunteer Experience": async () => {
    await apiVolunteerList().then(res => {
      VolunteerExperiencelist.value.list = res.data.list || []
    })
  },


};
let showInfo = ref({})
let userId = null;
onMounted(() => {
  let inter = setInterval(() => {
    userId = window.localStorage.getItem('userId')
    if (userId) {
      getpublicListInfo(userId)
      clearInterval(inter)
    }
  }, 300)
})
const getpublicListInfo = (userId) => {
  let localPublicList = window.localStorage.getItem(`${userId}_publicList`)
  if (localPublicList) {
    publicList = JSON.parse(localPublicList)
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
    publicList = tempPublicList
  }

  publicList.forEach(val => {
    showInfo.value[val.txt] = val.state
    if (val.state) {
      if (getInfoFuncObj[val.txt]) {
        let timeout = setTimeout(() => {
          getInfoFuncObj[val.txt]()
          clearTimeout(timeout)
        }, 300)
      }
    }
  })
}

let userInfo = ref({})
let myInfo = ref()
/**
 * 获取用户信息
 * @returns {Promise<void>}
 */
const getUser = async () => {
  await apUserInfo().then(res => {
    userInfo.value = res.data
    myInfo.value = {
      firstName: res.data.firstName,
      lastName: res.data.lastName,
      countryId: res.data.countryId || null,
      countryName: res.data.countryName,
      city: res.data.city,
      birthday: res.data.birthday,
      // avatar: res.data.avatar,
      avatar: (userInfo.value.avatar.indexOf("https://") !== -1) ? userInfo.value.avatar : (userInfo.value.avatar ? localStorage.getItem('infoOssBucketUrl') + userInfo.value.avatar : ''),

      phone: useCookie('_userInformation').value.phone,
      email: useCookie('_userInformation').value.email,
      languages: {
        title: "Languages",
        style: 'margin-right: 14px;',
        list: []
      }
    }

    if (!userInfo.value.personalSkills) userInfo.value.personalSkills = ''
    if (userInfo.value.personalSkills.indexOf(",") !== -1 ||
        userInfo.value.personalSkills.length) {
      // skillItem.value = userInfo.value.personalSkills.split(',')
      Skills.value.list = userInfo.value.personalSkills.split(',').map(val => {
        return {txt: val}
      }) || []

    }
    if (!userInfo.value.interests) userInfo.value.interests = ''
    if (userInfo.value.interests.indexOf(",") !== -1 ||
        userInfo.value.interests.length) {
      // interestsItem.value = userInfo.value.interests.split(',')
      Interests.value.list = userInfo.value.interests.split(',').map(val => {
        return {txt: val}
      }) || []
    }
    if (!userInfo.value.datingTargets) userInfo.value.datingTargets = ''
    if (userInfo.value.datingTargets.indexOf(",") !== -1 ||
        userInfo.value.datingTargets.length) {
      // idealGoalList.value = userInfo.value.datingTargets.split(',')
    }
    if (!userInfo.value.purposeCity) userInfo.value.purposeCity = ''
    if (userInfo.value.purposeCity.indexOf(",") !== -1 ||
        userInfo.value.purposeCity.length) {
      // purposeCityList.value = userInfo.value.purposeCity.split(',')
    }
    if (!userInfo.value.personalStrength) myInfo.value.personalStrength = ''
    if (userInfo.value.personalStrength.length) {
      myInfo.value.personalStrength = userInfo.value.personalStrength
    }
    if (res.data.languages) {
      myInfo.value.languages.list = res.data.languages.split(',').map(val => {
        return {txt: val}
      })


    }

  })
}
let timeout = setTimeout(() => {
  getUser()
  clearTimeout(timeout)
}, 300)


</script>

<style scoped>
.break-word {
  word-break: break-word;
}

.dom-center {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 20;
  width: 100vw;
  height: 100vh;
  background: #fff;
  overflow: auto;
}

#resumPreviewMain {
  background: linear-gradient(180deg, #F5FFFB 0%, #EBEBFF 100%);
}

// 预览
.resum-preview {
  position: relative;
  top: 0;
  left: 50%;
  z-index: 21;
  width: 993px;
  transform: translateX(-50%);
  width: 993px;
  background: linear-gradient(180deg, #F5FFFB 0%, #EBEBFF 100%);
  border-radius: 13px;

  .resum-preview-main {
    width: 994px;
    padding: 120px 50px 72px;
  //background: linear-gradient(180deg, #F5FFFB 0%, #EBEBFF 100%);
  }

  .about-me {
    font-size: 20px;
    height: auto;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #333333;
    line-height: 28px;
    padding: 64px 70px 30px 46px;
    opacity: 1;
    position: relative;
    background-image: url("assets/img/personal-center/about-me-ba.png");
    background-size: 100% 100%;
    border-radius: 10px;
    min-height: 100px !important;

    .about-me-hint {
      position: absolute;
      top: 34px;
      left: 14px;
      width: 47px;
      height: 77px;
      font-size: 64px;
      font-family: Inter-Semi Bold, Inter;
      font-weight: normal;
      color: #000000;
      line-height: 75px;
    }

    .about-me-content {
      position: relative;
      width: 880px;
      font-size: 20px;
      font-family: Inter-Regular, Inter;
      font-weight: 400;
      color: #333333;
      line-height: 28px;
    //direction: rtl; text-align: left; position: relative;
      left: -42px;
      padding-left: 50px;
      padding-right: 58px;
    }

    .about-me-border {
      width: 6px;
      height: 100%;
      position: absolute;
      left: 0;
      border-radius: 6px;
      background-color: #000;
    }
  }


  .basic-information {
    .name {
      width: 508px;
      font-size: 64px;
      font-family: Almarai-Bold, Almarai;
      font-weight: bold;
      color: #000000;
      line-height: 64px;
    }

    .profession {
      height: 30px;
      font-size: 21px;
      font-family: Inter-Regular, Inter;
      font-weight: 400;
      color: #000000;
      line-height: 25px;
      margin-top: 20px;
    }
  }

  .preview-head {
    height: 30px;
    font-size: 25px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #000000;
    line-height: 29px;
    margin-top: 26px;
    margin-bottom: 61px;
    justify-content: space-between;
  }

  .avatar {
    .avatar-img {
      width: 305px;
      height: 306px;
      border-radius: 10px 110px 10px 10px;
      opacity: 1;
      margin-right: 100px;
    }

  }

  .education {
    width: 304px;
    background: linear-gradient(#FAECDF 0%, #F8FADF 100%);
  // box-shadow:  0px 4px 4px 0px rgba(0, 0, 0, 0.1), 0px 4px 4px 0px rgba(0, 0, 0, 0.25);  // border-radius: 40px 0px 40px 0px; opacity: 1; // padding: 50px 40px; // margin-top: 38px;

    .education-title {
      width: 127px;
      height: 31px;
      font-size: 26px;
      font-family: Inter-Semi Bold, Inter;
      font-weight: normal;
      color: #000000;
      line-height: 30px;
      margin-bottom: 52px;
    }
  }

  .Job-experience {
    width: 551px;
    padding-bottom: 20px;
    border-radius: 40px 0px 40px 0px;
    background: linear-gradient(#F8F3EF 0%, #F0F8EF 100%);
  //box-shadow: inset 0px 4px 4px 0px rgba(0, 0, 0, 0.1), 0px 4px 4px 0px rgba(0, 0, 0, 0.25); border-radius: 40px 0px 40px 0px; opacity: 1; padding-top: 50px; padding-left: 48px; margin-top: 38px; margin-left: 26px;

    .Job-experience-title {
      height: 31px;
      font-size: 26px;
      font-family: Inter-Semi Bold, Inter;
      font-weight: normal;
      color: #000000;
      line-height: 30px;
      margin-bottom: 52px;
    }
  }

  .project-experience {
    margin-top: 26px;
    width: 551px;
    border-radius: 40px 0px 40px 0px;
    background: linear-gradient(#FDF7E3 0%, #F5E3FD 100%);
    padding-bottom: 26px;
  }

  .selector {
    &:after {
      content: '';
      width: 53px;
      height: 2px;
      background-color: #cccccc;
      position: absolute;
      left: 0;
      top: 40px;
      box-sizing: border-box;
    }
  }
}

.yearsold {
  margin-top: 4px;
}

// 个人信息
.personal-information {
  margin-top: 84px;

  .address, .yearsold, .contact-information {
    font-size: 18px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #585858;
    line-height: 21px;
  }

  .contact-information {
    height: auto;
    margin-top: 12px;
    color: #0A66C2;

    img {
      width: 24px;
      height: 24px;
    }
  }
}

// 语言-爱好-技能
.plate-six {
  width: 277px;
  min-height: 200px;
  border-radius: 10px 10px 10px 10px;
  opacity: 1;
  padding: 32px 22px;

  .six-content {
    white-space: pre-line;
  }

  .six-title {
    font-size: 30px;
    font-family: Fjalla One-Regular, Fjalla One;
    font-weight: 400;
    color: #000000;
    line-height: 35px;
    height: 58px;
  }

  .six-item {
    padding: 2px 13px;
    font-size: 18px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #327548;
    border-radius: 10px;
    opacity: 1;
    border: 1px solid rgba(50, 117, 72, 0.5);
    display: inline-block;
    margin-bottom: 14px;
    word-break: break-word;
  }
}

.plate-twelve {
  .plate-six {
    margin-right: 30px;
  }
}

.plate-seven {
  position: relative;
  margin-top: 70px;

  .seven-title {
    height: 58px;
    font-size: 48px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: 700;
    color: #000000;
    line-height: 56px;
    padding-left: 38px;
  }

  .seven-content {
    display: flex;
    flex-wrap: wrap;

    .seven-item {
      margin-right: 30px;
      margin-bottom: 30px;
      overflow: hidden;
    }

    & > :nth-child(3n) {
      margin-right: 0;
    }

    & > :nth-child(-n+3) {
      border-radius: 10px 110px 10px 10px;
      margin-bottom: 30px;
    }

    & > :nth-child(4) {
      border-radius: 10px 10px 110px 10px;
    }

    & > :nth-child(5) {
      border-radius: 10px 110px 10px 10px;
    }

    & > :nth-child(6) {
      border-radius: 4px;
    }

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
  }

  .seven-item {
    width: 277px;
    height: 277px;
    opacity: 1;
  }

  .seven-main {
    padding: 50px 80px;
    border-radius: 10px 110px 10px 10px;
    background-color: #fff;
    position: relative;
  }

  .seven-subtitle {
    font-size: 30px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #000000;
    line-height: 35px;
    margin-top: 40px;
  }

  .seven-content-two {
    height: 28px;
    font-size: 14px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #999999;
    line-height: 28px;
  }

  .seven-content-four {
    font-size: 14px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #616161;
    line-height: 28px;
    -webkit-background-clip: text;
  }

  .seven-content-five {
    height: 29px;
    font-size: 24px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #000000;
    line-height: 28px;
    -webkit-background-clip: text;
  }

  .seven-title-two {
    height: 28px;
    font-size: 20px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #000000;
    line-height: 28px;
    -webkit-background-clip: text;

  }

  .seven-content-three {
    /*width: 377px;*/
    font-size: 18px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #666666;
    line-height: 28px;
    -webkit-background-clip: text;
  }

  .seven-content-six {
    width: auto;
    word-break: break-word;
  }

  .mark {
    width: 20px;
    height: 20px;
    border-radius: 10px 10px 10px 10px;
    opacity: 1;
    position: absolute;
    top: 6px;
    left: -38px;
  }

  .mark-two {
    width: 6px;
    height: 110px;
    border-radius: 3px 3px 3px 3px;
    opacity: 1;
    position: absolute;
    top: 48px;
    left: 0;
  }

  .seven-content-four {
    font-size: 18px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #616161;
    line-height: 28px;
    -webkit-background-clip: text;
  }
}

.interval-four {
  margin-top: 30px;
  margin-bottom: 10px;
}

.interval-five {
  margin-top: 20px;
  margin-bottom: 4px;
}

.interval-six {
  margin-bottom: 20px;
}

.interval-eight {
  margin-top: 20px;
  margin-bottom: 14px;
}

.interval-nine {
  margin-top: 20px;
}

.education-main {
  width: 590px;

  .mark, .mark-two {
    background: #ED058C;
  }

  .education-img {
    position: absolute;
    right: -290px;
    bottom: 0;
    width: 187px;
    height: 200px;
  }
}

.job-eperience-main {
  width: 620px;
  background: #FFFFFF;
  border-radius: 10px 110px 10px 10px;
  opacity: 1;

  .mark, .mark-two {
    background: #486AFF;
  }
}

.job-experience-img {
  position: absolute;
  right: 16px;
  bottom: -80px;
  width: 306px;
  height: 296px;
}

.preview-experience {
  .mark, .mark-two {
    background: #F8C81E;
  }

  .seven-content-two {
    margin-bottom: 6px;
  }
}

.preview-custom-tem {
  .mark, .mark-two {
    background: #0599ED;
  }
}

.preview-volunteer-experience {
  .mark, .mark-two {
    background: #D132F8;
  }
}

.photos-plate {
  .preview-photos-item {
    width: 277px;
    margin-right: 30px;
    margin-bottom: 30px;

    .photos-img {
      width: 100%;
      height: 100%;
      object-fit: cover;
      border-radius: 10px;

    }
  }
}

.plate-seven {
  & > :nth-child(3n) {
    margin-right: 0;
  }
}

.interval-two {
  width: 285px;
  height: 6px;
  background: #333333;
  border-radius: 3px 3px 3px 3px;
  opacity: 1;
  margin-top: 10px;
  margin-bottom: 16px;
}

.interval-three {
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

.languages {
  background: #FFFFFF;
}

.title-img {
  transform: scale(0.5);
  position: relative;
  left: -55px;
}

.languages {
  .title-img {
    top: -34px;
  }
}

.interests {
  .title-img {
    top: -36px;
  }
}

.skills {
  .title-img {
    top: -36px;
    left: -30px;
  }
}

.six-border {
  border: 1px dashed #6F9976;;
}

.skills-item {
  padding: 4px 20px;
  background: #327548;
  border-radius: 23px 23px 23px 23px;
  font-size: 16px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: normal;
  color: #FEFEFC;
  line-height: 25px;
  margin-right: 20px;
}

.preview-tail {
  position: relative;
  left: -50px;
  width: calc(100% + 100px);
  padding-top: 340px;
  padding-left: 88px;
  margin-bottom: 130px;

  .preview-tail-content {
    position: absolute;
    top: 266px;
    right: 88px;

    width: 432px;
    font-size: 40px;
    font-family: Almarai-Bold, Almarai;
    font-weight: bold;
    color: #000000;
    line-height: 60px;
    -webkit-background-clip: text;
  }
}

.preview-bottom {
  font-size: 26px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #000000;
  line-height: 30px;
  -webkit-background-clip: text;
  display: flex;
  justify-content: space-between;
  margin-top: 130px;

  .preview-bottom-right {
    display: flex;
    align-items: center;

    span {
      margin-right: 10px;
    }
  }

  .resum-preview-logo {
    width: 120px;
    height: 50px;
    margin-right: 10px;
  }
}

.about-me-img {
  position: absolute;
  bottom: -22px;
  right: 97px;
  width: 70px;
}

.decorate-three {
  width: 40px;
  height: 40px;
  position: absolute;
  top: -6px;
  left: 18px;
}

.seven-content-three {
  word-break: break-word;
}
</style>
<style scoped lang="scss">
@media (max-width: 420px) {
  .DownloadOperate {
    margin: 0 20px;
    height: 41px;
  }
  .resum-preview {
    transform: none !important;
    left: 0 !important;
    width: calc(100vw) !important;

    .resum-preview-main {
      width: 100% !important;
      padding: 20px !important;

      .about-me {
        width: 100% !important;
        padding: 10px !important;

        .about-me-hint {
          top: 10px !important;
          width: 18px !important;
          height: 29px !important;

          .decorate-three {
            width: 18px !important;
            height: 29px !important;
          }
        }

        .about-me-content {
          width: calc(100% - 10px) !important;
          left: 10px !important;
          padding-left: 20px !important;
          padding-right: 0px !important;
          padding-top: 20px !important;
          font-size: 14px !important;
          font-weight: 400 !important;
          color: #333333 !important;
          line-height: 21px !important;
        }

        .about-me-border {
          height: 41px !important;
          width: 4px !important;
        }
      }
    }

    .personal-information {
      .avatar {
        width: 120px !important;

        .avatar-img {
          width: 115px !important;
          height: 115px !important;
          border-radius: 10px 45px 10px 10px !important;
        }
      }

      .basic-information {

        flex: 1;
        width: calc(100vw - 180px);
        margin-left: 20px;

        .name {
          width: auto !important;
          font-size: 30px;
          font-weight: bold;
          color: #000000;
          line-height: 35px;
        }

        .profession {
          font-size: 16px !important;
          line-height: 20px !important;
          height: 20px !important;
          margin-top: 5px !important;

        }

        .interval-two {
          width: 108px !important;
          height: 2px !important;
          margin-top: 5px !important;
          margin-bottom: 8px !important;
        }

        .address {
          font-size: 14px !important;
          font-weight: 400;
          line-height: 1.2 !important;
        }

        .yearsold {
          margin-top: 5px !important;
          font-size: 12px !important;
          font-weight: 400;
          line-height: 1.5 !important;
        }

        .contact-information {
          font-size: 14px !important;
          line-height: 1 !important;

          img {
            height: 24px !important;
            width: 24px !important;
            margin-right: 6px !important;
          }
        }
      }

      .orange-center {
        display: flex;
        align-items: center;
        font-size: 14px;
        word-break: break-all;
      }
    }

    .plate-seven {
      margin-top: 50px !important;

      .plate-six {
        width: calc((100vw - 60px) / 3) !important;
        padding: 9px !important;

        &:nth-child(2) {
          margin-right: 10px;
          margin-left: 10px;
        }

        .six-title {
          height: 35px !important;

          .title-img {
            top: 0 !important;
            width: 70px !important;
            height: 30px !important;
            transform: scale(1) !important;
            left: 0 !important;
          }

          .skills-img {
            width: 45px !important;
            height: 30px !important;
          }
        }

        .six-content {
          margin-top: 8px;
        }

        .six-item {
          font-size: 14px !important;
          padding: 2px 6px !important;
          margin-right: 5px !important;
          word-break: break-all;
        }

      }
    }

    .seven-content-four {
      font-size: 13px !important;
    }
  }
  .seven-title {
    height: 27px !important;
    font-size: 22px !important;
    font-weight: 700 !important;
    color: #000000 !important;
    line-height: 27px !important;
    padding-left: 10px !important;
  }
  .interval-three {
    left: 10px !important;
    width: 20px !important;
    height: 2px !important;
    background: #333333 !important;
    border-radius: 6px !important;
    opacity: 1 !important;
    margin-top: 5px !important;
    margin-bottom: 10px !important;
  }

  .education-main, .seven-main, .job-eperience-main {
    width: auto;
  !important;
    padding: 40px 10px 40px 40px !important;

    .seven-content-two {
      font-size: 12px;
    }

    .seven-content-three {
      width: auto !important;
    }

    .seven-title-two {
      font-size: 18px !important;
    }

    .seven-content-three {
      font-size: 14px !important;
      line-height: 20px !important;
    }

    .mark {
      left: -24px !important;
      width: 10px !important;
      height: 10px !important;
      top: 13px !important;
    }

    .seven-subtitle {
      word-break: break-all;
    }

    .seven-subtitle {
      height: auto;
    !important;
      margin-top: 0 !important;
      font-size: 20px !important;
    }

    .skills-item {
      margin-right: 5px !important;
      margin-bottom: 5px !important;
      padding: 2px 7px !important;
    }
  }
  .education-img, .job-experience-img {
    display: none;
  }
  .preview-tail {
    position: relative;
    margin-top: 50px !important;
    left: 0 !important;

    .preview-tail-content {
      font-size: 16px !important;
      font-weight: bold !important;
      color: #000000 !important;
      line-height: 24px !important;
      width: 184px !important;
      top: 0 !important;
      right: 10px !important;
    }

    img {
      width: 230px !important;
      height: 127px !important;
      padding-top: 30px;

    }
  }
  .preview-bottom {
    font-size: 12px !important;
    line-height: 20px !important;

    .resum-preview-logo {
      width: 57px !important;
      height: 20px !important;
    }
  }

  .about-me-img {
    right: 46px;
    bottom: -10px;
    width: 30px;
  }
  .plate-seven .seven-content-four {
    font-size: 14px;
  }
  .plate-twelve {
    .plate-six {
      margin-right: 0px;
    }
  }
}
</style>