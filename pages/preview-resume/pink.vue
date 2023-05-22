<template>
  <div class="preview-resume-container pink">
    <div class="resum-preview ">
      <DownloadOperate type="pink" class="DownloadOperate"></DownloadOperate>
      <div id="resumPreviewMain">
        <div class="resum-preview">
          <img src="@/assets/img/personal-center/preview-resume/preview-resume-head-blue.png">
          <div class="main" id="resumPreviewMain">
            <!-- 个人信息 -->
            <ResumePersonalInformation :showInfo="showInfo" :myInfo="myInfo" :type="type"/>

            <BoardBlock :type="type" v-if="showInfo['Skills']"
                        :list="Skills"
            />

            <BoardBlock :type="type" v-if="showInfo['Interests']"
                        :list="Interests"
            />

            <Grid :gridList="videosList" v-if="videosList.length>0" :type="type" gridTitle="Videos"/>

            <Grid :gridList="photosList" v-if="photosList.length>0" :type="type" gridTitle="Photos"/>

            <div class="relative" v-if="educationList.list.length>0">
              <img class="pink-decorate-img"
                   src="@/assets/img/personal-center/preview-resume/pink-decorate-img.png">
              <BoardBlockTwo
                  type="pink"
                  :data="testducationList"
                  :list="educationList"
              />
            </div>
            <!-- 教育 -->
            <BoardBlockTwo v-if="testexperienceList.list.length>0"
                           type="pink"
                           :data="experienceData"
                           :list="testexperienceList"
            />
            <!-- 经验 -->
            <BoardBlockTwo v-if="experienceDaList.list.length>0"
                           type="pink"
                           :data="experienceData"
                           :list="experienceDaList"
            />
            <BoardBlockTwo v-if="testJobExperience.list.length>0"
                           type="pink"
                           :data="testJobExperienceMsg"
            >

              <template #main>
                <div class="plate-content job-experience two-lines ">
                  <div
                      v-for="(item,index) in testJobExperience.list"
                      :key="index"
                      class="item"
                  >
                    <div>
                      <div class="name">
                        {{ item.companyName }}
                      </div>

                      <div class="content-two">
                        <span> Jan {{ item.startTime }}</span>
                        <span v-if="item.endTime"> -<span
                            v-if='item.endTime'> Jan {{ item.endTime }}</span></span>
                      </div>

                      <div class="content-three">Job content</div>
                      <div class="job-position">{{ item?.jobPosition }}</div>
                      <div class="content-four"
                           v-html="item.jobContent?.replace(/\n/g, '<br>')||''"
                      ></div>
                    </div>
                  </div>
                </div>
              </template>
            </BoardBlockTwo>

            <!-- 项目经验 -->
            <div class="relative" v-if="ProjectExperiencelist.list.length>0">
              <img class="pink-decorate-two"
                   src="@/assets/img/personal-center/preview-resume/pink-decorate-two.png">
              <BoardBlockTwo
                  type="pink"
                  :data="testProjectExperienceData"
              >
                <template #main>
                  <div class="plate-content project-experience">
                    <div
                        v-for="(item,index) in ProjectExperiencelist.list"
                        :key="index"
                        class="item"
                    >
                      <div>
                        <div class="name">
                          {{ item.projectName }}
                        </div>
                        <div>{{ item.address }}</div>
                        <div class="content-two">
                          <span> Jan {{ item.startTime }}</span>
                          <span v-if="item.endTime"> -<span
                              v-if='item.endTime'> Jan {{ item.endTime }}</span></span>
                        </div>

                        <div class="">
                          <div>
                            <div class="content-three">Job content</div>
                            <div class="content-four"
                                 v-html="item.jobContent?.replace(/\n/g, '<br>')||''"
                            ></div>
                          </div>

                          <div>
                            <div class="content-three">Project Content</div>
                            <div class="content-two" v-if="item?.jobPosition">
                              {{ item?.jobPosition }}
                            </div>
                            <div class="content-four"
                                 v-html="item.projectContent?.replace(/\n/g, '<br>')||''"
                            ></div>
                          </div>
                        </div>
                        <div class="content-eight" v-if="item.Skills.length>0">Skills:</div>
                        <div class="flex" v-if="item.Skills.length>0">
                          <div class="content-nine flex">
                            <div v-for="(sonItem,sonIndex) in item.Skills" :key="sonIndex">
                              {{ sonItem }}
                            </div>
                          </div>

                        </div>
                        <div class="content-nine" style="margin-top: 10px">
                          <span v-if="item?.projectLinkPlatform">{{ item?.projectLinkPlatform }}：</span>
                          <span>{{ item?.projectLink }}</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </template>

              </BoardBlockTwo>

            </div>

            <!-- 志愿者经历-->
            <BoardBlockTwo v-if="VolunteerExperiencelist.list.length>0"
                           type="pink"
                           :data="VolunteerExperienceData"
            >
              <template #main>
                <div class="plate-content volunteer-experience two-lines">
                  <div
                      v-for="(item,index) in VolunteerExperiencelist.list"
                      :key="index"
                      class="item"
                  >
                    <div>
                      <div class="name">
                        {{ item.title }}
                      </div>

                      <div class="content-two">
                        <span> Jan {{ item.startTime }}</span>
                        <span v-if="item.endTime"> -<span
                            v-if='item.endTime'> Jan {{ item.endTime }}</span></span>
                      </div>

                      <div class="content-three">Content</div>
                      <div class="content-four"
                           v-html="item.content?.replace(/\n/g, '<br>')||''"
                      ></div>
                    </div>
                  </div>
                </div>
              </template>
            </BoardBlockTwo>

            <!-- 自定义属性 -->
            <BoardBlockTwo v-if="customItemlist.list.length>0"
                           type="pink"
                           :data="customItemData"
            >
              <template #main>
                <div class="custom-item plate-content">
                  <div style="width: calc(50% - 20px);     margin-right: 20px;">
                    <div v-for="(item,index) in customItemlist.list"
                         :key="index"
                         class="item"
                    >
                      <div v-if="index%2==0">
                        <div class="name">
                          {{ item.title }}
                        </div>

                        <div class="content-two">
                          <span> Jan {{ item.startTime }}</span>
                          <span v-if="item.endTime"> -<span
                              v-if='item.endTime'> Jan {{ item.endTime }}</span></span>
                        </div>

                        <div class="content-three">Content</div>
                        <div class="content-four"
                             v-html="item.content?.replace(/\n/g, '<br>')||''"
                        ></div>

                        <div v-if="item.summary">
                          <div class="content-three">Summary</div>
                          <div class="content-four"
                               v-html="item.summary?.replace(/\n/g, '<br>')||''"
                          ></div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div style="width: calc(50% - 20px);     margin-right: 20px;">
                    <div v-for="(item,index) in customItemlist.list"
                         :key="index"
                         class="item"
                    >
                      <div v-if="index%2!==0">
                        <div class="name">
                          {{ item.title }}
                        </div>

                        <div class="content-two">
                          <span> Jan {{ item.startTime }}</span>
                          <span v-if="item.endTime"> -<span
                              v-if='item.endTime'> Jan {{ item.endTime }}</span></span>
                        </div>

                        <div class="content-three">Content</div>
                        <div class="content-four"
                             v-html="item.content?.replace(/\n/g, '<br>')||''"
                        ></div>

                        <div v-if="item.summary">
                          <div class="content-three">Summary</div>
                          <div class="content-four"
                               v-html="item.summary?.replace(/\n/g, '<br>')||''"
                          ></div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </template>
            </BoardBlockTwo>
          </div>
        </div>
        <Footer :type="type"></Footer>
      </div>

    </div>

  </div>
</template>

<script setup>
import Grid from '@/components/preview-resume/grid.vue'
import DownloadOperate from '@/components/preview-resume/download-operate.vue'
import ResumePersonalInformation from '@/components/preview-resume/personal-information.vue'
import BoardBlock from "~/components/preview-resume/board-block.vue";
import BoardBlockTwo from "~/components/preview-resume/board-block-two.vue";
import Footer from "~/components/preview-resume/footer.vue";
import {apUserInfo} from "~/api/user";
import {
  apiCustomAttributeList,
  apiExperienceList,
  apiPhotoList,
  apiUserCompanyList,
  apiUserEducationList,
  apiUserProjectList,
  apiVideoList,
  apiVolunteerList
} from "~/api/personal-wditor";

let type = 'pink'
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
  style: "border-radius:10px;margin:20px 0 ;",
  list: []
})

let Interests = ref({
  title: "Interests",
  style: "border-radius: 10px;",
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
      languages: {
        title: "Languages",
        style: ['pink', 'blue'].indexOf(type) > -1 ? "border-radius: 0px 0px 10px 10px;" : 'border-radius: 10px 10px 10px 10px;margin:20px 0;',
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
      // personalStrengthTxt.value = userInfo.value.personalStrength
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


<style lang="scss" scoped>
.preview-resume-container {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 30;
  width: 100vw;
  height: 100vh;
  overflow: auto;
  background: #fff;
}

.pink-back {
  background: linear-gradient(180deg, #FFE9E9 0%, #FFD6D6 100%);
}

#resumPreviewMain {
  background: linear-gradient(180deg, #FFE9E9 0%, #FFD6D6 100%);
}

// 预览
.resum-preview {
  position: relative;
  top: 0;
  left: 50%;
  z-index: 20;
  width: 993px;
  transform: translateX(-50%);


  .main {
    padding: 0 50px;
  }
}

.pink-decorate-img {
  width: 224px;
  height: 370px;
  position: absolute;
  top: -300px;
  left: -50px;
}

.plate-content {

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

  .content-four, .job-position {
    width: 381px;
    font-size: 16px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #666666;
    line-height: 28px;
    -webkit-background-clip: text;
  }

  .content-five {
    height: 29px;
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

  .content-eight {
    height: 28px;
    font-size: 18px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #000000;
    line-height: 28px;
    -webkit-background-clip: text;
  }

  .content-nine {
    width: 50%;
    font-size: 16px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #616161;
    line-height: 28px;
    -webkit-background-clip: text;
  }

  .item {
    margin-top: 60px;
  }
}

.two-lines {
  display: flex;
  flex-wrap: wrap;

  .item {
    margin-right: 21px;
    width: calc(50% - 21px);
  }

  > :nth-child(2n) {
    margin-right: 0;
    margin-left: 21px;
  }

}

.pink-decorate-two {
  position: absolute;
  top: -70px;
  right: -50px;
}

.project-experience {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;

  .item {
    width: calc(50% - 20px);
    margin-right: 20px;
  }

  .content-eight {
    margin-top: 26px;
    margin-bottom: 14px;
  }

  .content-nine {
    width: 100%;

    div {
      font-size: 16px;
      font-family: Inter-Semi Bold, Inter;
      font-weight: normal;
      color: #FEFEFC;
      line-height: 19px;
      -webkit-background-clip: text;
      padding: 8px 20px;
      border-radius: 23px;
      background: #BF0C3C;
      margin-right: 20px;
    }
  }
}

.volunteer-experience {

}

.custom-item {
  display: flex;
  flex-wrap: wrap;
}

.preview-tail {
  position: relative;
  top: -40px;
  left: -50px;
  width: calc(100% + 100px);

  .preview-tail-content {
    position: absolute;
    top: 354px;
    transform: translateX(50%);
    width: 530px;
    height: 270px;
    font-size: 40px;
    font-family: Almarai-Bold, Almarai;
    font-weight: bold;
    color: #BF0C3C;
    line-height: 60px;
    -webkit-background-clip: text;
    text-align: center;
  }
}

.preview-bottom {
  height: 36px;
  font-size: 26px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #000000;
  line-height: 30px;
  -webkit-background-clip: text;
  display: flex;
  justify-content: space-between;
  margin: 0 54px;
  padding-bottom: 72px;
}

.resum-preview-logo {
  width: 120px;
  height: 50px;
}

.preview-bottom-right {
  display: flex;
  align-items: center;

  span {
    margin-right: 10px;
  }
}

.content-four {
  word-break: break-word;
}
</style>
<style scoped lang="scss">
@media (max-width: 400px) {
  .pink-decorate-two {
    right: -20px !important;
    width: 119px;
    height: 119px;
  }
  :deep(.head-preview) {
    height: 61px;

    .h5-main {
      margin: 0;
    }
  }
  :deep(.preview-resume-container) {
    overflow: hidden;
  }
  :deep(.Languages-BoardBlock) {
    width: 100% !important;
  }
  :deep(.contact-information) {
    img {
      margin-right: 6px;
    }

    div {
      display: flex;
      align-items: center;
    }
  }
  .DownloadOperate {
    margin: 0 20px;
    height: 41px;
    position: absolute;
  }
  .resum-preview {
    width: 100%;

    .main {
      padding: 0 20px;
    }
  }
  .pink-decorate-img {
    width: 84px;
    height: 140px;
    top: -267px !important;
    left: -20px !important;
  }
  .title-content {

  }
  .plate-content {
    .item {
      margin-top: 15px !important;

      .name {
        font-size: 20px !important;
      }
    }
  }
  .content-two {
    font-size: 12px !important;
    line-height: 13px !important;
  }
  .content-three {
    margin-top: 10px !important;
    font-size: 18px !important;
  }
  .content-four {
    font-size: 14px !important;
  }
  .content-eight {
    font-size: 14px !important;
    font-weight: bold !important;

  }
  .content-nine {
    margin-bottom: 15px !important;

    div {
      margin-right: 5px !important;
      margin-bottom: 5px !important;
      padding: 2px 5px !important;
    }
  }
  .custom-item {
    width: 100% !important;
  }
  .two-lines {
    > :nth-child(2n) {
      margin-left: 0 !important;
    }

    .item {
      width: 100% !important;
      margin-right: 0 !important;

      .name {
        word-break: break-all;
      }
    }
  }
  .plate-content {
    .content-four, .job-position {
      width: 100% !important;
    }
  }
  .ht-flex, .custom-item {
    display: flex;
    flex-direction: column;
  }
  .project-experience {
    .item {
      width: 100% !important;
    }
  }

}
</style>