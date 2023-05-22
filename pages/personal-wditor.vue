<template>
    <n-message-provider>
        <!-- 主体 -->
        <div class="personal-wditor ba" ref="personalWditorRef">
            <div class="personal-wditor-main df">
                <!-- 个人中心左侧 -->
                <div class="personal-wditor-right">
                    <!-- 个人中心左侧-个人信息-->
                    <PersonalInformation></PersonalInformation>

                    <div v-if="isMobile">
                        <div class="h5-space"></div>

                        <H5IsPublic></H5IsPublic>
                    </div>

                    <div class="h5-space"></div>

                    <!-- 个人实力-->
                    <div class="personal-strength h5-personal-strength ba-02 hover-operation-edit-delete "
                         id="PersonalStrength">
                        <div class="personal-wditor-center-02 pr">
                            <div class="personal-strength-title pr">
                                <div>About me</div>
                                <div
                                        class="edit-img cursor-pointer"
                                        @click="showAddToBulletBox('aboutMe')"
                                >
                                    <img src="@/assets/img/personal-center/edit-img.png" alt=""/>
                                </div>
                            </div>
                            <div class="i-m-john-smith whitespace-pre-wrap break-words "
                                 v-html="personalStrengthTxt?personalStrengthTxt.replace(/\n/g, '<br>'):''">
                            </div>
                        </div>
                    </div>

                    <div class="h5-space"></div>

                    <!-- 分析 -->
                    <div class="analytics ba-02" id="Analytics">
                        <div class="personal-wditor-center-02">
                            <p class="analytics-title">Analytics</p>
                            <div class="analytics-private-to-you df">
                                <img
                                        src="@/assets/img/personal-center/private-to-you-img.png"
                                        alt=""
                                        class="private-to-you-img"
                                />
                                <span>Private to you</span>
                            </div>
                            <div class="df">
                                <img
                                        src="@/assets/img/personal-center/person.png"
                                        alt=""
                                        class="person-img"
                                />
                                <div>
                                    <div class="profile-views">{{ accessVal }} profile views</div>
                                    <div class="update-profile">
                                        Update your profile to attract viewers.
                                    </div>
                                    <div class="h5-update-profile">
                                        Update your profile to attract viewers.
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="h5-space"></div>

                    <!-- 资源 -->
                    <div class="resources ba-02" id="Resources">
                        <div class="personal-wditor-center-02">
                            <p class="analytics-title">Resources</p>
                            <div class="analytics-private-to-you df">
                                <img
                                        src="@/assets/img/personal-center/private-to-you-img.png"
                                        alt=""
                                        class="private-to-you-img"
                                />
                                <span>Private to you</span>
                            </div>
                            <div class="df">
                                <img
                                        src="@/assets/img/personal-center/notice-img.png"
                                        alt=""
                                        class="person-img"
                                />
                                <div>
                                    <div class="profile-views">
                                        <span>Creator mode</span>
                                        <el-switch
                                                class="is-it-public"
                                                v-model="isItPublic"
                                                @change="getIsItPublic"
                                        />
                                    </div>
                                    <div class="update-profile">
                                        Your profile will be found by more people.
                                    </div>
                                    <div class="h5-update-profile">
                                        Your profile will be found by more people.
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="h5-space"></div>

                    <!-- 语言 Languages -->
                    <div class="skill-dom ba-02 plate-item" id="Languages">
                        <div class="personal-wditor-center-02 hover-operation-edit-delete pr">
                            <div class="plate-title-two pr">
                                <div>Languages</div>
                                <div class="union-img" @click="showAddToBulletBox('languages')">
                                    <img src="@/assets/img/personal-center/union.png" alt=""/>
                                </div>
                            </div>
                            <n-space class="personal-strength-label">
                                <div
                                        v-for="(item,index) in  userInfoStore.languagesList"
                                        :key="index"
                                        class="skills-item-btn">
                                    <span>{{ item }}</span>
                                    <span class="label-operation-delete">
                    <img @click="deleteLanguagesItem(index)" class="bullet-box-add-icon"
                         src="@/assets/img/delete-icon.png">
                  </span>
                                </div>
                            </n-space>
                        </div>
                    </div>

                    <div class="h5-space"></div>

                    <!--   兴趣dom Interests    -->
                    <div class="interests ba-02 plate-item">
                        <div class="personal-wditor-center-02 hover-operation-edit-delete pr">
                            <div class="plate-title-two pr">
                                <div>Interests</div>

                                <div class="union-img" @click="showAddToBulletBox('interests')">
                                    <img src="@/assets/img/personal-center/union.png" alt=""/>
                                </div>
                            </div>
                            <n-space class="personal-strength-label">
                                <div
                                        v-for="(item, index) in userInfoStore.interestsItem"
                                        :key="index"
                                        class="skills-item-btn"
                                >
                                    <span>{{ item }}</span>
                                    <span class="label-operation-delete">
                    <img @click="deleteInterestsItem(index)" class="bullet-box-add-icon"
                         src="@/assets/img/delete-icon.png">
                  </span>
                                </div>
                            </n-space>
                        </div>
                    </div>

                    <div class="h5-space"></div>

                    <!-- 技能dom Skills-->
                    <div class="skill-dom ba-02 plate-item">
                        <div class="personal-wditor-center-02 hover-operation-edit-delete pr">
                            <div class="plate-title-two pr">
                                <div>Skills</div>
                                <div class="union-img" @click="showAddToBulletBox('skills')">
                                    <img src="@/assets/img/personal-center/union.png" alt=""/>
                                </div>
                            </div>
                            <n-space class="personal-strength-label">
                                <div class="skills-item-btn"
                                     v-for="(item,index) in userInfoStore.skillItem" :key="index">
                                    <span>{{ item }}</span>
                                    <span class="label-operation-delete">
                    <img @click="deleteSkillItem(index)" class="bullet-box-add-icon" src="@/assets/img/delete-icon.png">
                  </span>
                                </div>
                            </n-space>
                        </div>
                    </div>

                    <div class="h5-space"></div>

                    <!-- 教育dom -->
                    <div class="personal-strength education education-dom  ba-02 plate-item" id="Education">
                        <div class=" pr">
                            <div class="plate-item-title pr">
                                <p>Education</p>
                                <div class="union-img" @click="showAddToBulletBox('education','add')">
                                    <img src="@/assets/img/personal-center/union.png" alt=""/>
                                </div>
                            </div>

                            <div
                                    class="personal-strength-item pr"
                                    v-for="(item, index) in userInfoStore?.educationList"
                                    :key="index"
                                    @mouseenter="operationMouseenter"
                                    @mouseleave="operationMouseleave"
                            >
                                <div class="plate-five">
                                    <div class="four-haed">
                                        <div class="four-title plate-title-five">{{ item.name }}</div>
                                        <div class="education-space df content-two">
                                            <p class="">Jan {{ item.startTime }}</p>
                                            <span class="parting" v-if="item.endTime"> &emsp;-&emsp;</span>
                                            <p v-if="item.endTime">Jan {{ item.endTime }}</p>
                                        </div>
                                    </div>
                                    <div>
                                        <div class="four-subtitle">Content</div>
                                        <p class="four-secondary-content"
                                           v-html="item.content.replace(/\n/g, '<br>')"></p>
                                    </div>
                                </div>

                                <div class="operation">
                                    <img @click="editorialEducation(item.id)"
                                         src="@/assets/img/personal-center/edit-img.png" alt=""/>
                                    <img @click="deleteEducation(item.id,index)"
                                         src="@/assets/img/personal-center/delete-img.png"
                                         alt=""/>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="h5-space"></div>

                    <!-- 经验dom -->
                    <div class="personal-strength education ba-02 plate-item" id="Experience">
                        <div class=" pr">
                            <div class="plate-item-title pr">
                                <p>Experience</p>
                                <div class="union-img" @click="showAddToBulletBox('experience','add')">
                                    <img src="@/assets/img/personal-center/union.png" alt=""/>
                                </div>
                            </div>

                            <div
                                    class="personal-strength-item pr"
                                    v-for="(item, index) in userInfoStore.experienceList"
                                    :key="index"
                                    @mouseenter="operationMouseenter"
                                    @mouseleave="operationMouseleave"
                            >
                                <div class="plate-five">
                                    <div class="four-haed ">
                                        <div class="four-title plate-title-five">
                                            <div>{{ item.organization }}</div>
                                        </div>
                                        <div class="education-space df">
                                            <p class="content-two">Jan {{ item.startTime }}</p>
                                            <span class="parting content-two" v-if="item.endTime">&emsp;-&emsp;</span>
                                            <p v-if="item.endTime" class="parting content-two">Jan {{
                                                item.endTime
                                                }}</p>
                                        </div>
                                    </div>
                                    <div class="plate-title-three mb-38px">{{ item.position }}</div>

                                    <div class="plate-title-four">{{ item.title }}</div>

                                    <div class="mt-14px mb-26px">
                                        <div class="four-subtitle">Content</div>
                                        <p class="four-secondary-content"
                                           v-html="item.content.replace(/\n/g, '<br>')"></p>
                                    </div>
                                    <div>
                                        <div class="plate-title-four">Summary</div>
                                        <div class="four-secondary-content"
                                             v-html="item.summary.replace(/\n/g, '<br>')"></div>
                                    </div>
                                </div>

                                <div class="operation">
                                    <img @click="editorExperience(item.id)"
                                         src="@/assets/img/personal-center/edit-img.png" alt=""/>
                                    <img @click="deleteExperience(item.id,index)"
                                         src="@/assets/img/personal-center/delete-img.png"
                                         alt=""/>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="h5-space"></div>

                    <!-- 工作经验dom -->
                    <div class="personal-strength ba-02 plate-item" id="JobExperience">
                        <div class=" pr">
                            <div class="plate-item-title pr">
                                <p>Job experience</p>
                                <div
                                        class="union-img"
                                        @click="showAddToBulletBox('jobExperience','add')"
                                >
                                    <img src="@/assets/img/personal-center/union.png" alt=""/>
                                </div>
                            </div>

                            <div v-for="(item,index) in userInfoStore.companyList" :key="item.id" class=" pr"
                                 @mouseenter="operationMouseenter"
                                 @mouseleave="operationMouseleave">
                                <div class="plate-five">
                                    <div class="four-haed">
                                        <div class="four-title ">
                                            <div class="plate-title-five">{{ item.companyName }}</div>
                                            <div class="plate-title-three" v-if="item.companyIndustry ">{{
                                                item.companyIndustry
                                                }}
                                            </div>
                                        </div>
                                        <div class="content-two df">
                                            <p>Jan {{ item.startTime }}</p>
                                            <p v-if="item.endTime"> &emsp;-&emsp;Jan {{ item.endTime }}</p>
                                        </div>
                                    </div>

                                    <div class="mb-26px">
                                        <div class="four-title-three">
                                            <div class="four-title ">Job content</div>
                                            <div class="content-two">{{ item.jobPosition }}</div>
                                        </div>
                                        <p class="four-secondary-content"
                                           v-html="item.jobContent.replace(/\n/g, '<br>')"></p>
                                    </div>
                                </div>

                                <div class="operation">
                                    <img @click="editJobExperience(item.id)"
                                         src="@/assets/img/personal-center/edit-img.png" alt=""/>
                                    <img @click="deleteJobExperience(item.id,index)"
                                         src="@/assets/img/personal-center/delete-img.png"
                                         alt=""/>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="h5-space"></div>

                    <!-- 项目经验dom   -->
                    <div class="personal-strength ba-02 plate-item project-experience" id="projectExperience">
                        <div class=" pr">
                            <div class="plate-item-title pr">
                                <p>Project experience</p>
                                <div
                                        class="union-img"
                                        @click="showAddToBulletBox('projectExperience','add')"
                                >
                                    <img src="@/assets/img/personal-center/union.png" alt=""/>
                                </div>
                            </div>

                            <div v-for="(item,index) in userInfoStore.projectExperienceList" :key="item.id"
                                 class="plate-five pr"
                                 @mouseenter="operationMouseenter"
                                 @mouseleave="operationMouseleave">
                                <div class="four-haed ">
                                    <div class="four-title plate-title-five">
                                        {{ item.projectName }}
                                    </div>
                                    <div class="content-two df">
                                        <p>Jan {{ item.startTime }}</p>
                                        <p v-if="item.endTime"> &emsp;-&emsp;Jan {{ item.endTime }}</p>
                                    </div>
                                </div>

                                <div class="four-title-three">
                                    <div class="four-title">
                                        Project content
                                    </div>
                                    <div class="content-two">
                                        {{ item.jobPosition }}
                                    </div>
                                </div>
                                <div v-html="item.projectContent .replace(/\n/g, '<br>')"
                                     class="four-time-content"></div>

                                <div class="four-title-three">
                                    <div class="four-title">
                                        Job content
                                    </div>

                                </div>
                                <div v-html="item.jobContent .replace(/\n/g, '<br>')" class="four-time-content"></div>

                                <div class="four-title mb-14px seven-title ">Skills:</div>
                                <n-space class="personal-strength-label">
                                    <div v-for="(sonItem,sonIndex) in item.projectSkills"
                                         :key="sonIndex"
                                         class="skills-item-btn"
                                    >
                                        <span>{{ sonItem }}</span>
                                        <img @click="deleteProjectSkills(item,sonIndex)" class="bullet-box-add-icon"
                                             src="@/assets/img/delete-icon.png">
                                    </div>
                                </n-space>
                                <div class="four-title-four" v-if="item.projectLink">
                                    <div>{{ item.projectLinkPlatform }}:{{ item.projectLink }}</div>
                                </div>
                                <div class="operation">
                                    <img @click="editProjectExperience(item.id)"
                                         src="@/assets/img/personal-center/edit-img.png" alt=""/>
                                    <img @click="deleteProjectExperience(item.id,index)"
                                         src="@/assets/img/personal-center/delete-img.png"
                                         alt=""/>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="h5-space"></div>

                    <!-- 志愿者经历dom -->
                    <div class="personal-strength volunteer-experience ba-02 plate-item" id="volunteerExperience">
                        <div class=" pr">
                            <div class="plate-item-title pr">
                                <p>Volunteer experience</p>
                                <div class="union-img" @click="showAddToBulletBox('volunteerExperience','add')">
                                    <img src="@/assets/img/personal-center/union.png" alt=""/>
                                </div>
                            </div>

                            <div
                                    class="personal-strength-item pr"
                                    v-for="(item, index) in userInfoStore.volunteerExperienceList"
                                    :key="index"
                                    @mouseenter="operationMouseenter"
                                    @mouseleave="operationMouseleave"
                            >
                                <div class="plate-five">
                                    <div class="four-haed">
                                        <div class="four-title plate-title-five">{{ item.title }}</div>
                                        <div class="education-space df content-two">
                                            <p>Jan {{ item.startTime }}</p>
                                            <span class="parting" v-if="item.endTime">&emsp;-&emsp;</span>
                                            <p v-if="item.endTime" class="parting">Jan {{ item.endTime }}</p>
                                        </div>
                                    </div>
                                    <div>
                                        <div class="four-subtitle">Content</div>
                                        <p class="four-secondary-content"
                                           v-html="item.content.replace(/\n/g, '<br>')"></p>
                                    </div>
                                </div>

                                <div class="operation">
                                    <img @click="editVolunteerExperience(item.id)"
                                         src="@/assets/img/personal-center/edit-img.png"
                                         alt=""/>
                                    <img @click="deleteVolunteerExperience(item.id,index)"
                                         src="@/assets/img/personal-center/delete-img.png"
                                         alt=""/>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="h5-space"></div>

                    <!-- 自定义属性dom -->
                    <div class="personal-strength custom-item ba-02 plate-item" id="CustomItem">
                        <div class=" pr">
                            <div class="plate-subtitle pr">
                                <p>
                                    <span class=" ">Custom item</span>
                                    <span class="plate-subtitle-right"
                                    >Describe personalized content</span
                                    >
                                </p>

                                <div
                                        class="union-img"
                                        @click="showAddToBulletBox('customItem','add')"
                                >
                                    <img src="@/assets/img/personal-center/union.png" alt=""/>
                                </div>
                            </div>

                            <div
                                    class="plate-five pr"
                                    v-for="item in userInfoStore.customItemList"
                                    :key="item.id"
                                    @mouseenter="operationMouseenter"
                                    @mouseleave="operationMouseleave"
                            >
                                <div class="four-haed">
                                    <div class="four-title plate-title-five">{{ item.title }}</div>
                                    <div class="education-space df content-two">
                                        <p>Jan {{ item.startTime }}</p>
                                        <span v-if="item.endTime">&emsp;-&emsp;</span>
                                        <p v-if="item.endTime">Jan {{ item.endTime }}</p>
                                    </div>
                                </div>
                                <div class="four-subtitle ">Content</div>
                                <p class="four-secondary-content four-seven"
                                   v-html="item.content?item.content.replace(/\n/g, '<br>'):''"></p>
                                <div class="four-subtitle">Summary</div>
                                <p class="four-secondary-content"
                                   v-html="item.summary?item.summary.replace(/\n/g, '<br>'):''"></p>


                                <div class="operation">
                                    <img @click="editCustomItem(item.id)"
                                         src="@/assets/img/personal-center/edit-img.png" alt=""/>
                                    <img @click="deleteCustomItem(item.id)"
                                         src="@/assets/img/personal-center/delete-img.png" alt=""/>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <!--  是否公开   -->
                <IsPublic></IsPublic>

            </div>

            <div class="dom-clone" ref="domClone"></div>
        </div>
        <client-only>
            <n-dialog-provider>
                <PersonalWditorBulletBox
                        v-if="dialogBox"
                        :getId="getId"
                        :personaiInformation="userInfoStore.personaiInformation"
                        @ChangeGetId="ChangeGetId"
                        @changePersonaiInformation="changePersonaiInformation"
                ></PersonalWditorBulletBox>
            </n-dialog-provider>
        </client-only>
    </n-message-provider>
</template>

<script setup>
import {NDialogProvider, NMessageProvider, NSpace} from "naive-ui";
import {useRouter} from "vue-router";
import {ElMessage, ElMessageBox} from 'element-plus'
import localPhotoWallImg from '@/assets/img/personal-strength-ba.png'
import PersonalWditorBulletBox from "@/components/personal-wditor/personal-wditor-bullet-box.vue";
import H5IsPublic from "@/components/personal-wditor/h5-is-public.vue";
import {
    apiCustomAttributeDelete,
    apiCustomAttributeList,
    apiExperienceDelete,
    apiExperienceList,
    apiUpdateInterests,
    apiUpdateLanguages,
    apiUpdateMode,
    apiUpdateSkills,
    apiUserCompanyDelete,
    apiUserCompanyList,
    apiUserEducationDelete,
    apiUserEducationList,
    apiUserProjectDelete,
    apiUserProjectList,
    apiUserProjectUpdSkills,
    apiVolunteerDelete,
    apiVolunteerList
} from "~/api/personal-wditor";
import {useI18n} from "vue-i18n";
import {apUserInfo} from "~/api/user";
import {getCookie, showAddToBulletBox, calculateHowOld} from "~/utils/util";
import IsPublic from "@/components/personal-wditor/is-public.vue";
import PersonalInformation from "@/components/personal-wditor/personal-information.vue";
import axios from 'axios'
import {useState} from "#app";
import {userStore, webStore} from "@/composables/store/user.js"

let userInfoStore = userStore()
let webInfoStore = webStore()
const router = useRouter();


const {t, locale} = useI18n();
let userMsg = ref({})

useState("_isJoinNow").value = false
webInfoStore.isLogOut = true
let personalWditorRef = ref()

let isMobile = ref(false)
onMounted(async () => {
    userMsg.value = JSON.parse(localStorage.getItem("userMsg"))
    if (getCookie('token')) {
        const config = useRuntimeConfig()
        const baseUrl = config.public.VITE_API_URL;
        await axios({
            method: 'post',
            url: `${baseUrl}/api/config/info`,
            headers: {
                'Content-Type': 'application/json',
                'Authorization': getCookie("token")
            }
        }).then(res => {
            let infoOssBucketUrl = res.data.data.ossBucketUrl
            localStorage.setItem("infoOssBucketUrl", infoOssBucketUrl)
        })
    }
    getUser()

    // if (document.body.clientWidth <= 1500 && document.body.clientWidth >= 1368) {
    //   personalWditorRef.value.style.width = '1730px'
    //   personalWditorRef.value.style.minWidth = '1730px'
    // }
    // if (document.body.clientWidth <= 1367 && document.body.clientWidth >= 1026) {
    //   personalWditorRef.value.style.width = '1500px'
    //   personalWditorRef.value.style.minWidth = '1500px'
    // }
    // if (document.body.clientWidth <= 1025 && document.body.clientWidth >= 410) {
    //   personalWditorRef.value.style.width = '1367px'
    //   personalWditorRef.value.style.minWidth = '1367px'
    // }

    if (document.body.clientWidth < 420) {
        isMobile.value = true
    }
})

let loadingList = ref(false)
//  用户详情
let userInformation = useState("_userInformation")
// 获取个人信息
const getUser = async () => {
    await apUserInfo().then(res => {
        userInfoStore.userInfo = res.data
        accessVal.value = res.data.views
        userInfoStore.personaiInformation = {
            firstName: res.data.firstName,
            lastName: res.data.lastName,
            countryId: res.data.countryId || null,
            countryName: res.data.countryName,
            city: res.data.city,
            birthday: res.data.birthday,
        }
        userInfoStore.ageInformation = calculateHowOld(res.data.birthday)
        isItPublic.value = res.data.modeType === 1
        useState("_cloePersonaiInformation").value = userInfoStore.personaiInformation
        window.localStorage.setItem('userId', res.data.userId)
        isPersonaiInformation.value = true
        if (!userInfoStore.userInfo.personalSkills) userInfoStore.userInfo.personalSkills = ''
        if (userInfoStore.userInfo.personalSkills.indexOf(",") !== -1 ||
            userInfoStore.userInfo.personalSkills.length) {
            userInfoStore.skillItem = userInfoStore.userInfo.personalSkills.split(',')

            userInfoStore.skillItem.forEach(item => {
                skillItemBulletBox.value.push({label: item, value: item})
            })
        }
        if (!userInfoStore.userInfo.interests) userInfoStore.userInfo.interests = ''
        if (userInfoStore.userInfo.interests.indexOf(",") !== -1 ||
            userInfoStore.userInfo.interests.length) {
            userInfoStore.interestsItem = userInfoStore.userInfo.interests.split(',')
        }
        if (!userInfoStore.userInfo.datingTargets) userInfoStore.userInfo.datingTargets = ''
        if (userInfoStore.userInfo.datingTargets.indexOf(",") !== -1 ||
            userInfoStore.userInfo.datingTargets.length) {
            idealGoalList.value = userInfoStore.userInfo.datingTargets.split(',')
        }
        if (!userInfoStore.userInfo.purposeCity) userInfoStore.userInfo.purposeCity = ''
        if (userInfoStore.userInfo.purposeCity.indexOf(",") !== -1 ||
            userInfoStore.userInfo.purposeCity.length) {
            purposeCityList.value = userInfoStore.userInfo.purposeCity.split(',')
        }
        if (!userInfoStore.userInfo.personalStrength) userInfoStore.userInfo.personalStrength = ''
        if (userInfoStore.userInfo.personalStrength.length) {
            personalStrengthTxt.value = userInfoStore.userInfo.personalStrength
        }
        if (userInfoStore.userInfo.languages.length) {
            userInfoStore.languagesList = userInfoStore.userInfo.languages.split(',')
        }
        if (userInfoStore.userInfo.avatar) {
            if (userInfoStore.userInfo.avatar.indexOf("https://") !== -1) avatarImg.value = userInfoStore.userInfo.avatar
            else avatarImg.value = localStorage.getItem('infoOssBucketUrl') + userInfoStore.userInfo.avatar
        } else {
            avatarImg.value = false
        }

        if (userInfoStore.userInfo.backgroundImg) {
            userInfoStore.photoWallImg = localStorage.getItem('infoOssBucketUrl') + userInfoStore.userInfo.backgroundImg
        }
    })
}

// 头像
let avatarImg = useState("_avatarImg", () => null)
// 照片墙

let isPersonaiInformation = useState("_isPersonaiInformation", () => false)

const changePersonaiInformation = (obj) => {
    userInfoStore.personaiInformation = obj
}

// 是否公开
let isItPublic = ref(true)

const getIsItPublic = (modeType) => {
    if (modeType) modeType = 1
    else modeType = 2
    apiUpdateMode({modeType}).then(res => {
        if (res.data && res.data.status) {
            // 设置是否公开
        }
    })
}
// 技能-数据
let skillItemBulletBox = useState("_skillItemBulletBox", () => [])


// 理想目标-数据
let idealGoalList = useState('_idealGoalList', () => [])
// 目的城市-数据
let purposeCityList = useState('_purposeCityList', () => [])
// 个人实力-数据
let personalStrengthTxt = useState('_personalStrengthTxt', () => null)
// 被浏览次数统计
let accessVal = useState("_accessVal", () => 0)
// 筛选后的国家
let countryName = ref(null)
// 弹框
// let dialogBox = useState("_dialogBox", () => false)
let dialogBox = useState("_dialogBox", () => false)
// useState('_dialogBoxType').value = 'aboutMe'
// useState('_dialogBox').value = true
// 加载中
let dialog = useState("_dialog", () => false)

// 编辑操作-隐藏
const operationMouseleave = (e) => {
    let dom = e.target.lastChild;
    dom.style.setProperty("display", "none");
};
// 编辑操作-显示
const operationMouseenter = (e) => {
    let dom = e.target.lastChild;
    dom.style.setProperty("display", "flex");
};

let preview = ref();
let domClone = ref();

let callback = function (imgName, blob) {
    var triggerDownload = $("<a>")
        .attr("href", URL.createObjectURL(blob))
        .attr("download", imgName)
        .appendTo("body")
        .on("click", function () {
            if (navigator.msSaveBlob) {
                return navigator.msSaveBlob(blob, imgName);
            }
        });
    triggerDownload[0].click();
    triggerDownload.remove();
};

// 编辑操作的父级弹框
let bulletBoxEditDom = useState("_bulletBoxEditDom")

let bulletBox = ref([
    useState("_bulletBoxEditAvatar"),
    useState("_bulletBoxeditInYourProfile"),
    useState("_bulletBoBackgroundEditorr"),
    useState("_bulletBoxShareYourProfile"),
    useState("_bulletBoxAddToCustomAttributes"),
    useState("_bulletBoxAddToEducation"),
    useState("_bulletBoxAddToInterests"),
    useState("_bulletBoxAddIdealGoal"),
    useState("_bulletBoxAddPurposeCity"),
    useState("_bulletBoxAddToProjectExperience"),
    useState("_bulletBoxAddToDreamPosition"),
    useState("_bulletBoxAddToSkills"),
    useState("_bulletBoxAddToJobExperience"),
    useState("_bulletBoxAddToPersonalStrength"),
    useState("_dialogBox"),
]);


watch(bulletBox.value, (newVal) => {
        bulletBoxEditDom.value = newVal.some((element) => element.value);
    },
    {immediate: true, deep: true}
);
// 显示弹框的类型
let dialogBoxType = useState('_dialogBoxType', () => null)

// 编辑头像-弹框
let bulletBoxEditAvatar = useState("_bulletBoxEditAvatar", () => false);
// 在你的个人资料中编辑-弹框
let bulletBoxeditInYourProfile = useState(
    "_bulletBoxeditInYourProfile",
    () => false
);
// 背景编辑-弹框
let bulletBoBackgroundEditorr = useState(
    "_bulletBoBackgroundEditorr",
    () => false
);
// 添加到自定义属性-弹框
let bulletBoxAddToCustomAttributes = useState(
    "_bulletBoxAddToCustomAttributes",
    () => false
);


// 获取教育数据
const geteducationList = () => {
    apiUserEducationList().then(res => {
        if (res.data && res.data.list) {
            userInfoStore.educationList = res.data.list
        } else {
            userInfoStore.educationList = []
        }
    })
}
geteducationList()

// 添加-显示-教育-弹框
const showAddToEducationbulletBox = () => {
    userInfoStore.addToEducationForm = {orderBy: 0}
    getId.value = 0
    useState('_bulletBoxAddToEducation').value = true
}
let getId = ref()
let bulletBoxType = ref()
let frameType = useState("_frameType", () => false)
const ChangeGetId = (id) => {
    getId.value = id
    bulletBoxType.value = null
}

// 教育-编辑
const editorialEducation = (id) => {
    getId.value = id
    showAddToBulletBox('education', 'edit')
}
// 教育-删除
const deleteEducation = (id, index) => {
    ElMessageBox({
        title: 'Delete',
        confirmButtonText: 'Delete',
        roundButton: "true",
        message: h('p', null, [
            h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                'Are you sure you want to delete the current item?'),
        ]),
    }).then(() => {
        apiUserEducationDelete({id}).then(res => {
            // 教育-删除-请求
            if (res.data.status == 1) {
                userInfoStore.educationList.splice(index, 1)
            }
        })
    })
}


// 经验-获取数据
const getExperienceList = () => {
    apiExperienceList().then(res => {
        if (res.data && res.data.list) {
            userInfoStore.experienceList = res.data.list
        } else {
            userInfoStore.experienceList = []
        }
    })
}
getExperienceList()

// 经验-编辑
const editorExperience = (id) => {
    getId.value = id
    showAddToBulletBox('experience', 'edit')
}
// 经验-删除
const deleteExperience = (id, index) => {
    ElMessageBox({
        title: 'Delete',
        confirmButtonText: 'Delete',
        roundButton: "true",
        message: h('p', null, [
            h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                'Are you sure you want to delete the current item?'),
        ]),
    }).then(() => {
        apiExperienceDelete({id}).then(res => {
            // 经验-删除-请求
            if (res.data.status == 1) {
                userInfoStore.experienceList.splice(index, 1)
            }
        })
    })
}


// 获取-公司-数据
const getCompanyList = () => {
    apiUserCompanyList().then(res => {
        if (res.data && res.data.list) {
            let data = JSON.parse(JSON.stringify(res.data.list))
            data.forEach(item => {
                if (item.jobSkills.indexOf(",") !== -1 ||
                    item.jobSkills.length) {
                    item.jobSkills = item.jobSkills.split(',')
                }
            })
            userInfoStore.companyList = data
        } else {
            userInfoStore.companyList = []
        }
    })
}
getCompanyList()

// 公司-编辑
const editJobExperience = (index) => {
    getId.value = index
    showAddToBulletBox('jobExperience', 'edit')
}

// 公司-删除
const deleteJobExperience = (id, index) => {
    ElMessageBox({
        title: 'Delete',
        confirmButtonText: 'Delete',
        roundButton: "true",
        message: h('p', null, [
            h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                'Are you sure you want to delete the current item?'),
        ]),
    }).then(() => {
        apiUserCompanyDelete({id}).then(res => {
            // 公司-删除-请求
            if (res.data.status == 1) {
                userInfoStore.companyList.splice(index, 1)
            }

        })
    })
}

// 项目经验-获取数据
const getProjectExperienceList = () => {
    apiUserProjectList().then(res => {
        if (res.data && res.data.list) {
            let data = JSON.parse(JSON.stringify(res.data.list))
            data.forEach(item => {
                if (item.projectSkills.indexOf(",") !== -1 ||
                    item.projectSkills.length) {
                    item.projectSkills = item.projectSkills.split(',')
                }
            })
            userInfoStore.projectExperienceList = data
        } else {
            userInfoStore.projectExperienceList = []
        }
    })
}
getProjectExperienceList()

// 项目经验-编辑
const editProjectExperience = (id) => {
    getId.value = id
    showAddToBulletBox('projectExperience', 'edit')
}

// 项目经验-删除
const deleteProjectExperience = (id, index) => {
    ElMessageBox({
        title: 'Delete',
        confirmButtonText: 'Delete',
        roundButton: "true",
        message: h('p', null, [
            h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                'Are you sure you want to delete the current item?'),
        ]),
    }).then(() => {
        apiUserProjectDelete({id}).then(res => {
            // 项目经验-删除-请求
            if (res.data.status == 1) {
                userInfoStore.projectExperienceList.splice(index, 1)
            }

        })
    })
}

// 项目经历-技能-删除
const deleteProjectSkills = (item, index) => {
    if (item.projectSkills.length > 1) {
        ElMessageBox({
            title: 'Delete',
            confirmButtonText: 'Delete',
            roundButton: "true",
            message: h('p', null, [
                h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                    'Are you sure you want to delete the current item?'),
            ]),
        }).then(() => {
            let data = {
                id: 0,
                projectSkills: ''
            }
            item.projectSkills.splice(index, 1)
            item.projectSkills.forEach((itemforEach, indexforEach) => {
                if (indexforEach) {
                    data.projectSkills += `,${itemforEach}`
                } else {
                    data.projectSkills += `${itemforEach}`
                }
            })
            data.id = item.id

            apiUserProjectUpdSkills(data).then(res => {
                // 项目经历-技能-删除
            })
        })
    } else {
        ElMessage({
            showClose: true,
            message: t('personal-wditor.skills need to keep one'),
            type: 'error',
        })
    }
}


// 志愿者经历-获取数据
const getVolunteerExperienceList = () => {
    apiVolunteerList().then(res => {
        if (res.data && res.data.list) {
            let data = JSON.parse(JSON.stringify(res.data.list))
            userInfoStore.volunteerExperienceList = data
        } else {
            userInfoStore.volunteerExperienceList = []
        }
    })
}
getVolunteerExperienceList()

// 志愿者经历-编辑
const editVolunteerExperience = (id) => {
    getId.value = id
    showAddToBulletBox('volunteerExperience', 'edit')
}

// 志愿者经历-删除
const deleteVolunteerExperience = (id, index) => {
    ElMessageBox({
        title: 'Delete',
        confirmButtonText: 'Delete',
        roundButton: "true",
        message: h('p', null, [
            h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                'Are you sure you want to delete the current item?'),
        ]),
    }).then(() => {
        apiVolunteerDelete({id}).then(res => {
            // 志愿者经历-删除-请求
            if (res.data.status == 1) {
                userInfoStore.volunteerExperienceList.splice(index, 1)
            }

        })
    })
}

// 语言-删除
const deleteLanguagesItem = (index) => {
    if (userInfoStore.languagesList.length === 1) {
        ElMessage({
            type: 'error',
            message: "languages needs to keep one",//t('personal-wditor.Interest needs to keep one'),
        })
        return
    }
    ElMessageBox({
        title: 'Delete',
        confirmButtonText: 'Delete',
        roundButton: "true",
        message: h('p', null, [
            h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                'Are you sure you want to delete the current item?'),
        ]),
    }).then(() => {
        userInfoStore.languagesList.splice(index, 1)
        let data = ''
        userInfoStore.languagesList.forEach((item, index) => {
            if (index) {
                data += `,${item}`
            } else {
                data += item
            }
        })
        apiUpdateLanguages({languages: data}).then(res => {
            // 语言-删除-请求
        })
    })
}

// 删除-兴趣
const deleteInterestsItem = (index) => {
    if (userInfoStore.interestsItem.length === 1) {
        ElMessage({
            type: 'error',
            message: t('personal-wditor.Interest needs to keep one'),
        })
        return
    }
    ElMessageBox({
        title: 'Delete',
        confirmButtonText: 'Delete',
        roundButton: "true",
        message: h('p', null, [
            h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                'Are you sure you want to delete the current item?'),
        ]),
    }).then(() => {
        userInfoStore.interestsItem.splice(index, 1)
        let data = ''
        userInfoStore.interestsItem.forEach((item, index) => {
            if (index) {
                data += `,${item}`
            } else {
                data += item
            }
        })
        apiUpdateInterests({interests: data}).then(res => {
            // 兴趣-删除-请求
        })
    })
}

// 添加到项目经验-弹框
let bulletBoxAddToProjectExperience = useState(
    "_bulletBoxAddToProjectExperience",
    () => false
);
// 添加到梦想位置-弹框
let bulletBoxAddToDreamPosition = useState(
    "_bulletBoxAddToDreamPosition",
    () => false
);

// 删除-技能
const deleteSkillItem = (index) => {
    if (userInfoStore.skillItem.length === 1) {
        ElMessage({
            type: 'error',
            message: t('personal-wditor.skills need to keep one'),
        })
        return
    }
    ElMessageBox({
        title: 'Delete',
        confirmButtonText: 'Delete',
        roundButton: "true",
        message: h('p', null, [
            h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                'Are you sure you want to delete the current item?'),
        ]),
    }).then(() => {
        userInfoStore.skillItem.splice(index, 1)
        let data = ''
        userInfoStore.skillItem.forEach((item, index) => {
            if (index) {
                data += `,${item}`
            } else {
                data += item
            }
        })
        apiUpdateSkills({personalSkills: data}).then(res => {
            // 技能-删除-请求
        })
    })
}

// 重新获取数据
let reacquireData = useState("_reacquireData", () => 0)
// 修改自定义-数据
let reacquireDataItem = useState("_reacquireDataItem", () => 0)


// 获取自定义标题
const getApiCustomAttributeList = () => {

    apiCustomAttributeList().then(res => {
        if (res.data.list) {
            userInfoStore.customItemList = res.data.list
        } else {
            userInfoStore.customItemList = []
        }
    })
}
getApiCustomAttributeList()

watch(() => reacquireData.value, (newVal) => {
        if (newVal === 1) {
            getApiCustomAttributeList()
        }
    },
    {immediate: true, deep: true}
);

const editCustomItem = (id) => {
    getId.value = id
    showAddToBulletBox('customItem', 'edit')
}

// 自定义标题-删除
const deleteCustomItem = (id) => {
    ElMessageBox({
        title: 'Delete',
        confirmButtonText: 'Delete',
        roundButton: "true",
        message: h('p', null, [
            h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                'Are you sure you want to delete the current item?'),
        ]),
    }).then(() => {
        apiCustomAttributeDelete({id}).then(res => {
            if (res.data && res.data.status) {
                // 自定义标题-删除-请求
                // 获取-自定义标题
                getApiCustomAttributeList()
            }
        })
    })
}
// 自定义标题-修改
const reviseCustomItem = (item) => {
    reacquireDataItem.value = item
    userInfoStore.addToCustomAttributes = item
    useState('_bulletBoxAddToCustomAttributes').value = true
}

</script>

<style lang="scss" scoped>
@import "assets/styles/personal-wditor.scss";
@import "assets/styles/h5-personal-wditor.scss";
</style>