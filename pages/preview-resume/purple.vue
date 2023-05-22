<template>
    <div class="preview-resume-container purple">
        <div class="resum-preview">
            <DownloadOperate type="purple" class="DownloadOperate"></DownloadOperate>
            <div id="resumPreviewMain">
                <div class="resum-preview">
                    <img class="head-purple-img"
                         src="@/assets/img/personal-center/preview-resume/preview-resume-head-purple.png">
                    <div class="main" id="resumPreviewMain">
                        <!-- 个人信息 -->
                        <ResumePersonalInformation :showInfo="showInfo" :myInfo="myInfo" :type="type"/>

                        <BoardBlock :type="type" v-if="showInfo['Interests']"
                                    :list="Interests"
                        />

                        <BoardBlock :type="type" v-if="showInfo['Skills']"
                                    :list="Skills"
                        />

                        <Grid :gridList="photosList" v-if="photosList.length>0" :type="type" gridTitle="Photos"/>

                        <Grid :gridList="videosList" v-if="videosList.length>0" :type="type" gridTitle="Videos"/>

                        <div class="relative">
                            <BoardBlockNoEquallyTwo v-if="educationList.list.length>0"
                                                    type="purple"
                                                    :data="testducationList"
                                                    :list="educationList"
                            />
                        </div>
                        <BoardBlockNoEquallyTwo v-if="experienceDaList.list.length>0"
                                                type="purple"
                                                :data="experienceData"
                                                :list="experienceDaList"
                        />

                        <BoardBlockNoEquallyTwo v-if="testJobExperience.list.length>0"
                                                type="purple"
                                                :data="testJobExperienceMsg"
                                                :list="testJobExperience"
                        />

                        <BoardBlockNoEquallyTwo v-if="ProjectExperiencelist.list.length>0"
                                                type="purple"
                                                :data="testProjectExperienceData"
                                                :list="ProjectExperiencelist"
                        />
                        <BoardBlockNoEquallyTwo v-if="VolunteerExperiencelist.list.length>0"
                                                type="purple"
                                                :data="VolunteerExperienceData"
                                                :list="VolunteerExperiencelist"
                        />
                        <BoardBlockNoEquallyTwo
                                v-if="customItemlist.list.length>0"
                                type="purple"
                                :data="customItemData"
                                :list="customItemlist"
                        />
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
import BoardBlockNoEquallyTwo from "~/components/preview-resume/board-block-no-equally-two.vue";
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

let type = 'purple'
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
    style: "border-radius: 10px;margin-top:20px",
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

#resumPreviewMain {
  background: #F6EEE2;
}

// 预览
.resum-preview {
  position: relative;
  left: 50%;
  z-index: 20;
  width: 993px;
  transform: translateX(-50%);
  //padding-bottom: 72px;
  background: #F6EEE2;


  .main {
    padding: 0 50px;
  }
}

.head-purple {
  width: 993px;
  height: 1402px;
}

.purple-decorate-img {
  width: 224px;
  height: 370px;
  position: absolute;
  top: -300px;
  left: -50px;
}

:deep(.head-preview ) {
  height: 0;
}
</style>
<style scoped lang="scss">
@media (max-width: 400px) {
  :deep(.head-preview ) {
    .h5-main {
      position: fixed;
      top: 30px;
    }
  }
  .DownloadOperate {
    //margin: 0 20px;
    height: 41px;
    position: absolute;
  }
  .resum-preview {
    width: 100%;
    padding-top: 10px;
    padding-bottom: 0 !important;
    top: -20px;
    height: calc(100% + 20px);

    .main {
      padding: 0 20px;
    }
  }
}
</style>