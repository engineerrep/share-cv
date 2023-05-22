import {defineStore} from 'pinia'
import localPhotoWallImg from '@/assets/img/personal-strength-ba.png'

export const userStore = defineStore("userStore", {
    state: () => {
        return {
            emial: '',
            photoWallImg: '',
            personaiInformation: {
                firstName: null,
                lastName: null,
                countryId: null,
                countryName: null,
                city: null,
                birthday: null,
            },
            userInfo:{},
            ageInformation: 0,//更具生日计算的年龄
            addToEducationForm: {},
            educationList:[],
            experienceFormData: {},
            experienceList:[],
            addToJobExperienceForm: {},
            companyList:[],
            addToprojectExperience: {},
            projectExperienceList:[],
            volunteerExperienceForm: {},
            volunteerExperienceList:[],
            addToCustomAttributes: {},
            customItemList:[],
            skillItem: [],
            interestsItem: [],
            idealGoalList: {},
            purposeCityList: {},
            personalStrengthTxt: {},
            languagesList: [],
            avatarImg: {}


        }
    },
    actions: {
        signOut() {
            /****************************************清除个人数据****************************************************/
            //对象清空
            ['personaiInformation', 'addToEducationForm', 'experienceFormData', 'addToJobExperienceForm', 'addToprojectExperience', 'volunteerExperienceForm',
                'addToCustomAttributes', 'idealGoalList', 'purposeCityList', 'personalStrengthTxt',  'avatarImg','userInfo'].forEach(val => {
                this[val] = {}
            });
            //数组清空
            ['educationList','experienceList','companyList','volunteerExperienceList','customItemList','skillItem','interestsItem','languagesList'].forEach(val => {
                this[val] = []
            });
            //字符串
            [].forEach(val => {
                this[val] = ''
            });
            //数字
            ['ageInformation'].forEach(val => {
                this[val] = 0
            });
            this.photoWallImg = localPhotoWallImg;
            /****************************************清除网站数据****************************************************/
            let webInfoStore = webStore();
            webInfoStore.isLogOut = true
        },
    },
    persist: {
        storage: persistedState.localStorage,
    }


})
export const webStore = defineStore("webStore", {
    state: () => {
        return {
            isLogOut: false
        }
    },
    persist: {
        storage: persistedState.localStorage,
    }


})

