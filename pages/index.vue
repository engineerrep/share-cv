<template>
  <div>
    <!-- 欢迎来到真正的简历社区 -->
    <div class="greet-main center center02 h5-center">
      <div class="df show-greet-main">
        <div class="greet">
          <div
              class="resume-community-title"
              v-html="$t('home.title')"
          ></div>
          <!-- 邮箱，密码，登录 -->
          <el-form class="sign-in-form" v-loading="loadingIndxForm">
            <el-form
                ref="formDataRef"
                :model="formData"
                :rules="rulesformData"
                require-asterisk-position="right"
                label-position="top"
                class="index-form"
            >
              <el-form-item prop="emailVal">
                <el-input
                    v-model="formData.emailVal"
                    required="required"
                    type="text"
                    @blur="blurVal(emailLabel)"
                    @focus="focuslVal(emailLabel)"
                    @input="formData.emailVal=formData.emailVal?.trim()"
                />
                <label class="email-label" ref="emailLabel">{{ $t("home.Email") }}</label>

              </el-form-item>

              <el-form-item prop="pwdVal">
                <el-input
                    :type="pwdInputType"
                    v-model="formData.pwdVal"
                    required="required"
                    @blur="blurVal(pwdValLabel)"
                    @focus="focuslVal(pwdValLabel)"
                />
                <label class="email-label" ref="pwdValLabel">{{ $t("home.Password") }}</label>

                <img
                    v-if="pwdState"
                    alt=""
                    class="pwd-state"
                    src="@/assets/img/pwd-hide.png"
                    @click="changePwdState"
                    @mouseenter="setPwdState"
                />

                <img
                    v-else
                    alt=""
                    class="pwd-state"
                    src="@/assets/img/pwd-show.png"
                    @click="changePwdState"
                    @mouseleave="setPwdState"
                />
              </el-form-item>
            </el-form>

            <div class="forgot-password" @click="router.push('/forgot-password')">
              {{ $t("home.ForgotPassword") }}
            </div>

            <div>
              <CVButton @click="emailPasswordLogin">
                <template #main>
              <span>
               {{ $t("home.Signin") }}
              </span>
                </template>
              </CVButton>
            </div>

            <div class="dividing-line">
              <div class="division-or">
                <div class="or">{{ $t("home.or") }}</div>
              </div>
            </div>

            <div>

              <button :disabled="!isReady"
                      class="btn sign-in-with-google"
                      type="button"

              >
                <a href="javascript:;" class="sign-in-google" @click="login()">
                  <img
                      alt=""
                      class="google-symbol-img"
                      src="@/assets/img/google-symbol-1.png"
                  />
                  <span>{{ $t("home.Sign in with Google") }}</span>
                </a>
              </button>

              <!--                <GoogleSignInButton-->
              <!--                        @success="handleLoginSuccess"-->
              <!--                        @error="handleLoginError"-->
              <!--                ></GoogleSignInButton>-->
            </div>
          </el-form>
        </div>

        <div class="h5-greet-img" style="flex: 1">
          <img
              alt=""
              class="frame-img"
              src="@/assets/img/jobplace-community.png"
          />
          <img
              alt=""
              class="h5-frame-img"
              src="@/assets/img/jobplace-community.png"
          />
        </div>
      </div>
    </div>

    <!-- 完整的模版 -->
    <div class="full-resume ba">
      <div class="h5-full-resume-main center02 h5-center df">
        <div class="flex1 complete-resume">
          {{ $t("home.Module") }}
        </div>
        <div class="flex1 component">
          <div class="component-title">{{ $t("home.Modules") }}</div>
          <div class="df df-wrap component-item">
            <div
                v-for="item in componentItem"
                :key="item"
                class="complete-item"
            >
              {{ $t("home." + item) }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 分享你想要分享的目标 -->
    <div class="share">
      <div class="h5-full-resume-main center02 h5-center df">
        <div class="flex1 share-title">
          {{ $t("home.ShareGoalsTip") }}
        </div>
        <div class="h5-share-title">
          {{ $t("home.ShareGoalsTip") }}
        </div>
        <div class="flex1 component">
          <div class="component-title">{{ $t("home.ShareGoals") }}</div>
          <div class="df df-wrap">
            <div
                v-for="(item, index) in targetstItem"
                :key="index"
                :style="{ fontSize: item.fs }"
                class="complete-item"
            >
              {{ $t("home." + item.txt) }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 分享你的简历，让更多人了解你  -->
    <div class="ba share-resume">
      <div class="share-resume-main center02 h5-center df">
        <div class="share-esume-title share_esume_title flex1">
          <div>Share your resume and let</div>
          <div>more people know you</div>
        </div>
        <div class="h5-share-esume-title share_esume_title flex1">
          {{ $t("home.Share") }}
        </div>
        <div class="flex1 share-resume-dom">
          <button class="share-resume-btn" type="button"

                  @click="backToTheTop">
            {{ $t("home.Share resume") }}
          </button>
        </div>
      </div>
    </div>

    <!-- 个人介绍 -->
    <div class="ba01">
      <div
          class="h5-self-introduction self-introduction-main df center02 h5-center"
      >
        <div class="h5-self-introduction-main self-introduction">
          <div class="self_introduction_title01 self-introduction-title01">
            {{ $t("home.ShareTitle") }}
          </div>
          <div class="h5-self_introduction_title01 self-introduction-title01">
            {{ $t("home.ShareContent") }}
          </div>
          <div class="self_introduction_title02 self-introduction-title02">
            {{ $t("home.ShareContent") }}
          </div>
          <div class="h5-self_introduction_title02 whitespace-pre-wrap self-introduction-title02">
            {{ $t("home.ShareContent") }}
          </div>
        </div>

        <div class="personal-display">
          <img
              alt=""
              class="personal-display-img"
              src="@/assets/img/personal-display-img.png"
          />
          <img
              alt=""
              class="h5-personal-display-img"
              src="@/assets/img/h5-personal-display-img.png"
          />
        </div>
      </div>
    </div>

    <!--  图片展示 -->
    <div class="img-exhibit"></div>
  </div>
</template>

<script setup>
import {googleAuthUrl} from "@/api/home";
import {useRouter} from "vue-router";
import {useI18n} from "vue-i18n";
import {apiRegisterCheckEmail, emailLogin, googleAuthLogin} from "@/api/login.js";
import {googleDataError, uuid} from "@/utils/util.js";
import {ElLoading, ElMessage} from 'element-plus'
import {backToTheTop} from "~/utils/util";
import {apUserInfo} from "~/api/user";
import {
  useTokenClient,
} from "vue3-google-signin";
import CVButton from '@/components/form/button.vue';
import {useToken} from "~/composables/user";
import {userStore} from "@/composables/store/user.js"

let userInfoStore = userStore()
const router = useRouter();
const {t, locale} = useI18n();

let _googleCode = ref()
// 监听路由变化
// let token = ref(useState('_token', () => false));

const requesta = (accessToken) => {
  // token.value = accessToken
  // document.cookie = "token" + "=" + accessToken;
  useToken().value = accessToken
}
let _login = useState('_login')
const _loginFn = (loadingInstance) => {
  _login.value = 2
  router.push('/register');
  loadingInstance?.close()
}
let emailLabel = ref()
let pwdValLabel = ref()
// 当移出input的时候触发
const blurVal = (dom) => {
  if (document.body.clientWidth >= 410) {
    dom.setAttribute('style', 'font-size:16px;top:3px')
  } else {
    dom.setAttribute('style', 'font-size:14px;top:-8px')
  }
}
// 当移入input的时候触发
const focuslVal = (dom) => {
  if (document.body.clientWidth >= 410) {
    dom.setAttribute('style', 'font-size:12px;top:-3px')
  } else {
    dom.setAttribute('style', 'font-size:12px;top:-8px')
  }
}

onMounted(() => {
  let _googleCode = ref(useState("_googleCode", () => null))
  // 谷歌授权失败提示
  window.googleAuthorizationError = () => {
    ElMessage({
      showClose: true,
      message: 'Google authorization login failed during translation',
      type: 'error',
    })
  }
  window.request = function (code, state) {
    _googleCode.value = code
    let url = window.location.origin;
    // let url = window.location.href.replace(/\/$/, '');
    let params = {
      code,
      state,
      callbackURL: url,
      deviceId: uuid(),
      browserName: getBrowserInfo()[0].replace('/', 'V'),
    }

    googleAuthLogin(params).then((res) => {
      let {code, data} = res
      if (code == 200) {
        let {accessToken, firstName, lastName, email, phone} = data
        requesta(accessToken)
        if (firstName && lastName) {
          setTimeout(() => {
            getUser()
            userInformation.value = {
              phone,
              email,
            }
          }, 1000)
        } else {
          _loginFn()
        }
      }

    });
  }

  watch(
      () => router.currentRoute.value.path,
      () => {
        const {code, state, error} = router.currentRoute.value.query;
        if (error) {
          googleDataError(error)
          window.opener.googleAuthorizationError()
        }
        if (code && state) {
          document.cookie = "googleCode" + "=" + code;
          _googleCode.value = code
          window.opener.request(code, state)
          window.close()
        }
      },
      {immediate: true, deep: true}
  );
})

// 组件
const componentItem = [
  "Languages",
  "Skills",
  "Photos",
  "Passions",
  "Summary",
  "Strengths",
  "Experience",
  "Projects",
  "TrainingCourses",
  "Industry",
  "Expertise",
  "Education",
  "Custom",
];
// 目标
const targetstItem = [
  {txt: "Favorite person"},
  {txt: "Looking for a job"},
  {txt: "Find a partner"},
  {txt: "Find a founder"},
  {txt: "Find someone to travel with"},
  {txt: "Find someone to climb the mountain together", fs: "15px"},
  {txt: "Chat with someone"},
  {txt: "Find appointments"},
  {txt: "Find social activities"},
  {txt: "To be recognized by someone"},
  {txt: "Leave behind talent"},
];

// 登录-数据
let formData = ref({
  emailVal: null,
  pwdVal: null,
})

// 密码状态
let pwdState = ref(true);
// 密码输入框类型
let pwdInputType = ref("password");
// 改变密码状态
const changePwdState = () => {
  if (pwdState.value) {
    pwdState.value = false;
    pwdInputType.value = "txt";
  } else {
    pwdState.value = true;
    pwdInputType.value = "password";
  }
};
let pcSide = useState("_pcSide")
const setPwdState = () => {
  if (pcSide.value) {
    pwdState.value = !pwdState.value;
    if (pwdState.value) pwdInputType.value = "password";
    else pwdInputType.value = "txt";
  }
}

/**
 *谷歌登錄成功回調
 * @param response
 */
const handleOnSuccess = (response) => {
  let params = {
    "code": response.access_token,
    "callbackURL": window.location.href.replace(/\/$/, ""),
    "deviceId": uuid(),
    "browserName": getBrowserInfo()[0].replace('/', 'V')// 浏览器名称
  }
  const loadingInstance = ElLoading.service({fullscreen: true})
  googleAuthLogin(params).then((res) => {
    let {code, data} = res
    if (code == 200) {
      let {accessToken, firstName, lastName, email, phone} = data
      requesta(accessToken)
      if (firstName && lastName) {
        setTimeout(() => {
          getUser(loadingInstance)
          userInformation.value = {
            phone,
            email,
          }
        }, 1000)
      } else {
        _loginFn(loadingInstance)
      }
    }
  }).catch(() => {
    loadingInstance.close()
  });
};

/**
 * 谷歌登錄失败回調
 */
const handleOnError = () => {
  ElMessage({
    showClose: true,
    message: 'Google authorization login failed during translation',
    type: 'error',
  })
};

const {isReady, login} = useTokenClient({
  onSuccess: handleOnSuccess,
  onError: handleOnError,
  // other options
});

// 谷歌登录
const signInWithGoogle = async () => {
  let url = window.location.href.replace(/\/$/, "");
  let params = {callbackURL: url};
  const width = 600;
  const height = 600;
  let left = (screen.width / 2) - (width / 2);
  let top = (screen.height / 2) - (height / 2);
  let res = await googleAuthUrl(params);
  if (res.code === 200) {
    let googleUrl = res.data.url;
    let clientId = res.data.clientId;
    let scope = res.data.scope;
    let state = res.data.state;
    let googleAuth =
        googleUrl +
        "?client_id=" +
        clientId +
        "&redirect_uri=" +
        url +
        "&response_type=code&scope=" +
        scope +
        "&state=" +
        state;
    let features =
        "toolbar=no,scrollbars=yes,menubar=no,width=558,height=600,resizable=no,top=" +
        top +
        ",left=" +
        left;

    window.open(googleAuth, "newwin", features);

  }
};

useHead({
  title: t("seo.home.title"),
  meta: {
    keywords: t("seo.home.keywords"),
    description: t("seo.home.description"),
  },
});

let accessVal = useState("_accessVal")
let isPersonaiInformation = useState("_isPersonaiInformation")
let idealGoalList = useState("_idealGoalList")
let purposeCityList = useState("_purposeCityList")
let personalStrengthTxt = useState("_personalStrengthTxt")

let avatarImg = useState("_avatarImg")
let userInformation = useCookie('_userInformation')
// 获取个人信息
const getUser = (loadingInstance) => {
  apUserInfo().then(res => {
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
    isPersonaiInformation.value = true
    if (!userInfoStore.userInfo.personalSkills) userInfoStore.userInfo.personalSkills = ''
    if (userInfoStore.userInfo.personalSkills.indexOf(",") !== -1 ||
        userInfoStore.userInfo.personalSkills.length) {
      userInfoStore.skillItem = userInfoStore.userInfo.personalSkills.split(',')
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
    }

    if (userInfoStore.userInfo.backgroundImg) {
      userInfoStore.photoWallImg = localStorage.getItem('infoOssBucketUrl') + userInfoStore.userInfo.backgroundImg
    }
    router.push('/personal-wditor')
  }).finally(() => {
    loadingInstance?.close()
  })
}


let userMsg = useState("_userMsg", () => null)
let loadingIndxForm = ref(false)

const rulesEmailVal = (rule, value, callback) => {
  if (!value) {
    return callback(new Error('Please enter a valid email address'))
  }
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  setTimeout(() => {
    if (!emailRegex.test(value)) {
      callback(new Error('Please enter a valid email address'))
    } else {
      callback()
    }
  })
}
const rulesPwdVal = (rule, value, callback) => {
  if (!value) {
    return callback(new Error('Password must be 6 characters or more.'))
  }
  setTimeout(() => {
    if (value.length < 6) {
      callback(new Error('Password must be 6 characters or more.'))
    } else {
      callback()
    }
  })
}

// 进行验证
const rulesformData = ref({
  emailVal: [
    {required: true, validator: rulesEmailVal, trigger: 'blur'},
    {validator: rulesEmailVal, trigger: 'change'},
  ],
  pwdVal: [
    {required: true, validator: rulesPwdVal, trigger: 'blur'},
    {validator: rulesPwdVal, trigger: 'change'},
  ]
})

let formDataRef = ref()
// 邮箱密码登录
const emailPasswordLogin = async () => {
  formDataRef.value.validate((valid) => {
    if (valid) {
      grecaptcha.enterprise.ready(function () {
        grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
            .then(async function (token) {

              let params = {
                "email": formData.value.emailVal,
                "password": formData.value.pwdVal,
                "googleCaptcha": token,// 获取的谷歌id
                "deviceId": uuid(),// 设备id
                "browserName": getBrowserInfo()[0].replace('/', 'V')// 浏览器名称
              }
              loadingIndxForm.value = true
              emailLogin(params)
                  .then(res => {
                    loadingIndxForm.value = false
                    if (res.data.accessToken) {
                      userMsg.value = res.data
                      let {firstName, lastName, accessExpire, phone, phoneNum, email} = res.data
                      // document.cookie = "token=" + res.data.accessToken;
                      useToken().value = res.data.accessToken
                      if (firstName && lastName) {
                        useCookie('_userInformation').value = {
                          phone, phoneNum, email
                        }
                        localStorage.setItem("_email", email)
                        router.push('/personal-wditor')
                      } else {
                        useState("_login").value = 1
                        router.push('/register')
                      }
                    }
                  })
                  .finally(() => {
                    loadingIndxForm.value = false
                  })

            })
      })

    }
  })
}

</script>
<style>
@media (max-width: 420px) {
  .index-form .el-input__inner {
    top: 24px;
    font-size: 14px;
    height: 14px;
    width: 86%;
    vertical-align: middle;
    line-height: 1;
  }
}
</style>
<style lang="scss" scoped>
@import "@/assets/styles/index.scss";
@import "@/assets/styles/index-h5.scss";
</style>
