<template>
    <n-card title="">
        <n-form ref="formRef" :label-width="80" :model="formValue" :rules="rules">
            <n-form-item :label='$t("register.Email") '
                         class="email-input  verify-one-sinup-input"
                         path="Email"
            >
                <n-input
                        class="h-44px rounded-6px   text-18px text-[#4B4B4B] leading-44px"
                        v-model:value="formValue.Email"
                        :placeholder="$t('register.Please enter a valid email address')"
                        @input="formValue.Email=formValue.Email?.trim()"
                        @blur="verifyEmailDoesItExist"
                />

            </n-form-item>

            <n-form-item
                    :label="$t('register.Password') + '（' + $t('register.6 digits or more') + '）'"
                    path="Password"
                    class="password-input verify-one-sinup-input"
            >
                <n-input
                        class=" h-44px rounded-6px text-18px text-[#4B4B4B] leading-44px"
                        v-model:value="formValue.Password"
                        type="password"
                        :show-password-on="pcSide?'mousedown':'click'"
                        show-password-on="mousedown"
                        :placeholder="$t('register.please enter password')"
                >
                    <template #password-visible-icon>
                        <img class="pwd-state" src="@/assets/img/pwd-show.png">
                    </template>
                    <template #password-invisible-icon>
                        <img class="pwd-state" src="@/assets/img/pwd-hide.png">
                    </template>
                </n-input>
            </n-form-item>
            <div class="text-center  centerStyle mt-2">
                <!--        -->
                <div v-if="locale==='en'">
                    <div>
                        {{ $t("register.agree and join") }}
                        <span> CV </span>
                        <nuxt-link class="link" to="/legal/user-agreement">
                            {{ $t("register.agree and join03") }}
                        </nuxt-link>
                        <nuxt-link class="link" to="/legal/privacy-policy">
                            {{ $t("register.agree and join02") }}
                        </nuxt-link>
                        <nuxt-link class="link" to="/legal/privacy-policy">{{ }}</nuxt-link>
                        {{ $t("register.and") }}
                        <nuxt-link class="link" to="/legal/cookie-policy">
                            Cookie {{ $t("register.policy") }}
                        </nuxt-link>
                    </div>
                </div>
                <div v-else-if="locale==='cn'">
                    {{ $t("register.agree and join") }}
                    <nuxt-link class="link" to="/legal/user-agreement">
                        {{ $t("register.agree and join02") }}
                    </nuxt-link
                    >
                    <nuxt-link class="link" to="/legal/privacy-policy">{{ }}
                    </nuxt-link>
                    {{ $t("register.and") }}
                    <nuxt-link class="link" to="/legal/cookie-policy">Cookie{{ $t("register.policy") }}
                    </nuxt-link>
                </div>
            </div>
            <n-form-item>
                <button
                        v-loading="isLoading"
                        @click="handleValidateClick"
                        type="button"
                        class="ch-button w-full h-50px px-4 "
                        :class="isDisabled?'is-disabled':'button-one '"
                >
                    {{ $t("register.Agree & Join") }}
                </button>
            </n-form-item>
        </n-form>
        <!--    <div class="flex items-center px-18px mb-22px" style="z-index: 100">-->
        <!--      <div class="flex-1 h-2px bg-[#D9D9D9]"></div>-->
        <!--      <div class="px-20px text-[#1C1C1C] text-20px">or</div>-->
        <!--      <div class="flex-1 h-2px bg-[#D9D9D9]"></div>-->
        <!--    </div>-->

        <!--    <div-->
        <!--        @click="handleGoogleClick"-->
        <!--        class="w-full cursor-pointer h-50px rounded-56px border-solid border-1 border-[#DADADA] flex items-center justify-center"-->
        <!--    >-->
        <!--      <img-->
        <!--          src="@/assets/img/btn_google.png"-->
        <!--          class="w-28px h-28px flex-shrink-0"-->
        <!--          alt=""-->
        <!--      />-->
        <!--      <span class="ml-31px text-[#616363] text-18px">-->
        <!--        {{ $t("register.Sign in with Google") }}</span-->
        <!--      >-->
        <!--    </div>-->
        <div
                class="signin-tail mt-50px mb-20px text-center font-500 cursor-pointer"
                @click="goSign"
        >
            {{ $t("register.Already on Realme") }} ?
            <span class="signin">{{ $t("register.Signin") }}</span>
        </div>
    </n-card>
    <el-dialog
            v-model="showModal"
            title="Tips"
            width="30%"
    >
        <span>This is a message</span>
        <template #footer>
      <span class="dialog-footer">
      </span>
        </template>
    </el-dialog>
    <NMessageProvider>
        <n-modal v-model:show="isModal" style="width: 660px;">
            <SixSendCode
                    :formData="formValue"
                    @isModalEmit="isModalEmit"
                    :_email="_email"
            ></SixSendCode>
        </n-modal>
        <n-modal v-model:show="showModalForgotPassword" style="width: 577px;">
            <ForgotPassword
                    :formData="formData"
                    :_email="_email"
            ></ForgotPassword>
        </n-modal>
    </NMessageProvider>


</template>

<script setup>
import {NCard, NForm, NFormItem, NInput, NMessageProvider, NModal,} from "naive-ui";
import {useRouter} from "vue-router";
import {useI18n} from "vue-i18n";
import {apiRegisterCheckEmail, googleAuthLogin, sendCode} from "@/api/login";
import {googleAuthUrl} from "@/api/home";
import {getBrowserInfo, googleDataError, reviseDom, uuid} from "@/utils/util";
import SixSendCode from "@/components/register/sixSendCode.vue";
import ForgotPassword from "~/pages/forgot-password.vue";
import {useState} from "nuxt/app";
import {useToken} from "~/composables/user";

const {t, locale, i18n} = useI18n();
const router = useRouter();
let showModal = useState("_showModal", () => false)
let isModal = ref(false)
let isLoading = ref(false)
// 忘记密码-弹框
let showModalForgotPassword = useState('_showModalForgotPassword')

const isModalEmit = (value) => {
    isModal.value = value
}

let themeColor = useState("_themeColor")
let themeColorTwo = useState("_themeColorTwo")

const requesta = (accessToken) => {
    // token.value = accessToken
    // document.cookie = "token" + "=" + accessToken;
    useToken().value = accessToken
}
let pcSide = useState("_pcSide")
// 谷歌登录-窗口关闭了-去个人中心
onMounted(() => {
    let _googleCode = ref(useState("_googleCode", () => null))

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
            console.log('sss', res);
            let {code, data} = res
            if (code == 200) {
                let {accessToken, firstName, lastName} = data
                requesta(accessToken)

                // firstName = null
                if (firstName && lastName) {
                    router.push('/personal-wditor')
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
            // console.log('41111',code,state,error)
            if (error) {
                googleDataError(error);
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

// 谷歌登录
const handleGoogleClick = async (e) => {

    let url = window.location.href.replace(/\/$/, "");
    let params = {callbackURL: url};
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
            window.screenTop +
            ",left=" +
            window.screenLeft;
        window.open(googleAuth, "newwin", features);
    }
};
const emit = defineEmits(["setUpShowIndex"]);
const goSign = () => {
    router.push("/");
};
const formRef = ref(null);

/**
 * 表单数据
 */
const formValue = ref({
    Email: "",
    Password: "",
});

// 改变验证提示
watch(locale, (newVal, oldVal) => {
    reviseDom('.email-input .n-form-item-feedback__line', t('register.Please enter a valid email address'))
    reviseDom('.password-input .n-form-item-feedback__line', t('register.please enter password'))
})

/**
 * 登录规则
 */
const rules = {
    Password: {
        required: true,
        trigger: ["input", "blur"],
        validator(rule, value) {
            if (!value) {
                return new Error(t("register.please enter password"));
            } else if (value.length < 6) {
                return new Error(t("register.6 digits or more"));
            }
            return true;
        },
    },
    Email: {
        required: true,
        trigger: ["input", "blur"],
        validator(rule, value) {
            const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
            if (!emailRegex.test(value)) {
                return new Error(t("register.Please enter a valid email address"));
            }
            return true;
        },
    },
};
let isDisabled = ref(true)
let _email = ref()
let verifyEmail = useState('_verifyEmail', () => false)
let emailCode = useState("_emailCode", () => false)
// 验证邮箱提示
let isRrulesEmailInputHint = ref(false)
// 验证邮箱提示
let rulesEmailInputHint = ref()

// 验证邮箱是否存在
const verifyEmailDoesItExist = () => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(formValue.value.Email)) {
        rulesEmailInputHint.value = 'Please enter a valid email address'
        isRrulesEmailInputHint.value = true
        return
    }
    rulesEmailInputHint.value = ''
    isRrulesEmailInputHint.value = false
    isDisabled.value = true
    grecaptcha.enterprise.ready(function () {
        grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
            .then(function (token) {
                let data = {
                    email: formValue.value.Email,
                    googleCaptcha: token
                }
                apiRegisterCheckEmail(data).then(res => {
                    if (res.data.status) {
                        isDisabled.value = false
                    }
                }).catch(err => {
                    isDisabled.value = true
                })
            });
    })
}

watch(() => formValue.value.Email, (newVal) => {
        verifyEmail.value = false
    },
    {immediate: true, deep: true}
)

/**
 * 注册邮箱
 * @param {*} e
 */
const handleValidateClick = (e) => {
    if (!isDisabled.value) {
        e.preventDefault();
        formRef.value.validate(err => {
            if (!err) {
                if (verifyEmail.value) {
                    console.log('邮箱验证成功')
                } else {
                    formRef.value?.validate((errors) => {
                        if (!errors) {
                            console.log("这里调接口-验证邮箱是否注册");
                            _email.value = formValue.value.Email
                            grecaptcha.enterprise.ready(function () {
                                grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
                                    .then(function (token) {
                                        let params = {
                                            email: _email.value,
                                            googleCaptcha: token
                                        }
                                        isLoading.value = true
                                        sendCode(params).then((res) => {
                                            if (res.data.status == 1) {
                                                useState('_loading').value = false
                                                isModal.value = true
                                            }
                                        }).finally(() => {
                                            isLoading.value = false
                                        })
                                    });
                            })
                            // showModal.value = true
                        }
                    })
                }
            }
        })
    } else {
        formRef.value?.validate(() => {
        })
    }
}

</script>

<style lang="scss" scoped>

.rules-email-input-hint {
  position: absolute;
  top: 40px;
  color: #d03050;
  font-size: 14px;
}

.mark-location {
  position: absolute;
  top: -30px;
  left: 46px;
  color: #d03050;
  font-size: 14px;
}

.link {
  color: #005bb6;
  font-weight: 600;
  text-decoration: none;
}

.n-card {
  width: 630px;
  padding-top: 30px;
  margin: 0 auto;
  border-radius: 10px;
  box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0.25);
}

.centerStyle {
  height: 38px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #333333;
  line-height: 19px;
  font-size: 16px;
}

.signin {
  color: #005bb6;
}

.pwd-state {
  width: 26px;
  height: 26px;
}

.button-one {
  background-color: v-bind(themeColor);
  border: 1px solid v-bind(themeColor);

  &:hover {
    background-color: v-bind(themeColorTwo);
    border: 0;
  }
}

@media (max-width: 420px) {
  :deep(.n-form-item-label__text) {
    font-size: 18px;
    margin-top: 6px;
  }
  .pwd-state {
    width: 26px;
    height: 26px;
  }
  .centerStyle {
    height: 60px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #333333;
    line-height: 19px;
    font-size: 12px;
  }
  .n-input {
    font-size: 14px;
  }
  .signin-tail {
    font-size: 16px;
  }
  .n-card.n-card--bordered {
    border: 0;
  }
  .n-card {
    width: 100% !important;
    margin: 0;
    border-radius: 6px;
    box-shadow: 0 0 0 0 rgba(0, 0, 0, 0);
  }
}
</style>
