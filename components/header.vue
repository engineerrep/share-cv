<template>
  <!-- 头部导航 -->
  <div class="header header_bg">
    <div class="header-nav center center02 h5-center plate-ten">
      <div>
        <img @click="router.push('/')" class="logo-img cursor-pointer" src="@/assets/img/home_icon.png" alt=""/>
      </div>
      <n-space class="join-now">
        <n-button v-if="isJoinNow" @click="router.push('/register')"
                  type="tertiary" round ghost color="#0A66C2"
                  class="text-20">
          {{ $t("layout.JoinNow") }}
        </n-button
        >
        <n-dropdown
            v-if="webInfoStore.isLogOut"
            trigger="click"
            :options="logOutOptions"
            @select="optionSelect"
        >
          <span type="primary" round ghost color="#0A66C2" class="cursor-pointer">
                 <img
                     v-if="avatarImg"
                     :src="avatarImg"
                     alt=""
                     class="head-avatarImg pr"
                 />
                <img
                    v-else
                    src="@/assets/img/personal-center/icon/head-image-iocn.png"
                    class="head-avatarImg pr"
                />
            <div class="text-center me">Me</div>
          </span>
        </n-dropdown>

        <!--        <n-dropdown-->
        <!--            trigger="click"-->
        <!--            :options="lang"-->
        <!--            @select="setLang"-->
        <!--        >-->
        <!--          <n-button type="primary" round ghost color="#0A66C2">{{-->
        <!--              curLang[0].label-->
        <!--            }}-->
        <!--          </n-button>-->
        <!--        </n-dropdown>-->
      </n-space>
    </div>

    <div ref="scrollToTopButton" v-if="showScrollToTopButton"
         class="back-to-the-top cursor-pointer  fixed bottom-25 right-30px  p-2 rounded-full z-90">
      <img @click="backToTheTop" class="w-66px h-66px" src="@/assets/img/home/back-to-the-top.png">
    </div>
  </div>
</template>

<script setup>
import {NButton, NDropdown, NSpace} from "naive-ui";
import {langConfig} from "@/locals/config";
import {useI18n} from "vue-i18n";
import {useRouter} from "vue-router";
import {backToTheTop} from "@/utils/util";
import {apiLoginLogout} from "~/api/personal-wditor";
import {ElLoading} from 'element-plus'
import localPhotoWallImg from '@/assets/img/personal-strength-ba.png'
import {userLanguages} from '@/composables/languages.js'
import {useToken} from "~/composables/user";
import {userStore, webStore} from "@/composables/store/user.js"

let userInfoStore = userStore()
let webInfoStore = webStore()

const router = useRouter();

// 回到顶部-按钮
const scrollToTopButton = ref(null);

// 监听滚动事件
const scrollTop = ref(0);
const showScrollToTopButton = ref(false);

const handleScroll = () => {
  scrollTop.value = document.documentElement.scrollTop;
  showScrollToTopButton.value = scrollTop.value > 100;
}
onMounted(() => {
  if (document.body.clientWidth >= 410) {
    handleScroll()
    window.addEventListener('scroll', handleScroll);
  } else {
    showScrollToTopButton.value = false
  }
  // if (document.body.clientWidth <= 1500 && document.body.clientWidth >= 1367) {
  //   document.body.style.width = '1730px'
  //   document.body.style.minWidth = '1730px'
  // }
  // if (document.body.clientWidth <= 1367 && document.body.clientWidth >= 1025) {
  //   document.body.style.width = '1500px'
  //   document.body.style.minWidth = '1500px'
  // }
  // if (document.body.clientWidth <= 1025 && document.body.clientWidth >= 410) {
  //   document.body.style.width = '1167px'
  //   document.body.style.minWidth = '1167px'
  // }
  window.addEventListener('scroll', onScroll, true)
})
let plateTenLeft = ref()
const onScroll = () => {
  plateTenLeft.value = document.documentElement.scrollLeft + 'px'
}

const {t, locale} = useI18n();
const language = useCookie("lang");
const userLang = userLanguages();
const lang = ref(langConfig);
const curLang = computed(() => {
  return langConfig.filter((item) => {
    if (item.key === locale.value) {
      return item.label;
    }
  });
});
let avatarImg = useState("_avatarImg")

webInfoStore.isLogOut = false
let logOutOptions = [{
  label: 'Log out',
  key: 'Log out'
},
  {
    label: 'feedback',
    key: 'feedback'
  }
]

const token = useToken();
// 是否显示加入我们
let isJoinNow = useState("_isJoinNow", () => true)
watch(() => router.currentRoute.value.fullPath, (newVal) => {
      if (newVal === '/') isJoinNow.value = true
      else isJoinNow.value = false

      if (newVal === '/personal-wditor') webInfoStore.isLogOut = true
      else webInfoStore.isLogOut = false
    },
    {immediate: true, deep: true}
)


const setLang = (val) => {
  locale.value = val; // 文字显示切换
  language.value = val; //更新cookie
  userLang.value = val; //更新state
};
const bgColor = ref("bg-white");

const cookieToken = useCookie("token");
const isToken = computed(() => {
  return token.value;
});
watch(
    () => router.currentRoute.value.path,
    (toPath) => {
      // 改变顶部背景颜色
      if (toPath === "/") bgColor.value = "";
      else bgColor.value = "bg-white";
    },
    {immediate: true, deep: true}
);
const optionSelect = (key) => {
  switch (key) {
    case "Log out":
      signOut()
      break
    case "feedback":
      router.push('/feedback')
      break


  }
}
// 退出登录
const signOut = () => {
  const loadingInstance = ElLoading.service({fullscreen: true})
  apiLoginLogout().then(res => {
    if (res.data && res.data.status) {
      isJoinNow.value = true

      // document.cookie = 'token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;'
      useToken().value = ''
      useCookie('_userInformation').value = {
        phone: '', email: ''
      }
      userInfoStore.signOut()

      userInfoStore.photoWallImg = localPhotoWallImg;
      //清除保存的部分数据
      ['_experienceFormData', '_addToJobExperienceForm', '_addToprojectExperience', '_volunteerExperienceForm', '_addToCustomAttributes', '_skillItem', '_interestsItem', '_idealGoalList', '_purposeCityList', '_personalStrengthTxt', '_languagesList', '_avatarImg'].forEach(val => {
        useState(val).value = ''
      });
      router.push('/')


    }
  }).finally(() => {
    loadingInstance.close()
  })
}

</script>

<style lang="scss" scoped>
@import "@/assets/styles/header.scss";
@import "@/assets/styles/head-tail.scss";

.back-to-the-top {
  bottom: 130px;
}

.plate-ten {
  position: relative;
  left: v-bind(plateTenLeft);
}

</style>
