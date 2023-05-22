<template>

  <di class="center forgot-password-email">
    <div class="confirm-your-email">{{ title }}</div>
    <div class="input-sent-to">
      <p v-if="type==='phone'">
        <span>Please enter the verification code we gave you,It might take a few minutes to receive your code.</span>
        <span @click="showModalPasswordWmail=1" class="cursor-pointer text-theme-color">Change</span>
      </p>
      <div v-else-if="type==='mail'" class="forgot-mail">
        <div>Enter the 6-digit verification code sent to</div>
        <div>
          {{ props._email }}
          <span @click="changeEmail" class="cursor-pointer text-theme-color"> Change</span>

        </div>
      </div>
      <p v-else>
        <span>Enter the 6-digit verification code sent to {{ props._email }} </span>
        <span @click="showModalPasswordWmail=1" class="cursor-pointer text-theme-color"> Change</span>
      </p>
    </div>

    <div>
      <el-input @input="inputComplete" class="email-input email-verification"
                :class="codeError?'input-error':''"
                v-model="user.code" maxlength="6"
                placeholder=" "/>
      <div class="email-input-hint">
        <n-button class="resend-code cursor-pointer"
                  :class="isDisabledResendCode?'resend-code-disabled':''"
                  :disabled="isDisabledResendCode"
                  text
                  color="#828282"
                  @click="resendCode">Resend code
          <span class="ml-6px">{{ seconds }}S</span>
        </n-button>
        <div class="email-input-err" v-if="codeError">Verification code error</div>
      </div>
    </div>


    <div class="my-30px">
      <CVButton @click="verificationCodeVerification" v-loading="isLoading">
        <template #main>
          <span> Submit</span>
        </template>
      </CVButton>
      <div class="back"><span @click="router.push('/')">Cancel</span></div>
      <div class="see-this-email" v-if="type==='mail'">
        If you don't see the email in your inbox, check your spam folder.
      </div>
    </div>

  </di>
</template>

<script setup>
import {NButton, useMessage} from "naive-ui";
import {useI18n} from "vue-i18n";
import {useState} from "nuxt/app";
import {apiAccountChangeEmailCode, apiAccountChangeEmailWithPassword} from "~/api/user";
import {sendCode} from "~/api/login";
import {useRouter} from "vue-router";
import {
  apiChangeeMailWithCode, apiCheckEmailCode,
  apiCheckPhoneCode,
  apiSendPhoneCode,
  apiUserbaseUpdatePhone
} from "~/api/personal-wditor";
import CVButton from '@/components/form/button.vue';


const router = useRouter();
const message = useMessage()
const {t} = useI18n();
const props = defineProps({
  formData: Object,
  _email: String,
  userForm: Object,
});
let isLoading = ref(false)
const formData = ref({});
let _resetEmail = useState('_resetEmail', () => false)
const emit = defineEmits(["setUpShowIndex"]);
let user = ref({
  code: null
})
let codeError = ref(false)
let showModalPasswordWmail = useState('_showModalPasswordWmail')

let title = ref('We sent a code to your email')
/**
 * 判断是邮箱还是手机号
 * @param {type} phone 电话 mail 邮箱 没有就是设置密码
 * @param {type}  mail 邮箱 没有就是设置密码
 * @param {*}  没有就是设置密码
 * */
let {type} = router.currentRoute.value.query
if (type === 'phone') {
  title.value = 'Enter the code that was sent to your mobile phone.'
} else if (type === 'mail') {
  title.value = 'We sent a code to your email'
}
const changeEmail = () => {
  _resetEmail.value = true
  showModalPasswordWmail.value = 1
}
// 60秒倒计时
const seconds = ref(60)
// 测试code
let userCode = useState("_userCode", () => 0)
// 禁止按钮
const isDisabled = ref(true)
// 禁止按钮-重新发送验证码
const isDisabledResendCode = ref(false)
// 输入完成
const inputComplete = () => {
  if (user.value.code.length >= 6) isDisabled.value = false
  else isDisabled.value = true
}
// 60秒倒计时
const startTimer = () => {
  isDisabled.value = true
  isDisabledResendCode.value = true
  const timer = setInterval(() => {
    seconds.value--
    if (seconds.value === 0) {
      clearInterval(timer)
      isDisabledResendCode.value = false
      seconds.value = 60
    }
  }, 1000)
}
startTimer()

// 重置密码-重新发送验证码
const resendCode = () => {
  startTimer()
  if (seconds.value < 60) {
    return;
  }
  if (type === 'phone') {
    grecaptcha.enterprise.ready(function () {
      grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
          .then(function (token) {
            let params = {
              phone: props.userForm.phone,
              phoneNum: props.userForm.phoneNum,
              googleCaptcha: token
            }
            apiSendPhoneCode(params).then((res) => {

            })
          });
    })
  } else if (type === 'mail') {
    grecaptcha.enterprise.ready(function () {
      grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
          .then(function (token) {
            let params = {
              email: props._email,
              googleCaptcha: token,
            }
            sendCode(params).then((res) => {

            })
          });
    })
  } else {
    grecaptcha.enterprise.ready(function () {
      grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
          .then(function (token) {
            let params = {
              email: props._email,
              googleCaptcha: token
            }
            sendCode(params).then((res) => {

            })
          });
    })
  }

}

// 忘记密码-验证码验证 邮箱-手机号修改
const verificationCodeVerification = () => {

  let data = {code: user.value.code}
  let params = {code: user.value.code}
  if (type === 'phone') {
    if (user.value.code > 5) {
      data.phone = props.userForm.phone
      data.phoneNum = props.userForm.phoneNum
      params.phone = props.userForm.phone
      params.phoneNum = props.userForm.phoneNum
      params.countryId = props.userForm.countryId
      grecaptcha.enterprise.ready(function () {
        grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
            .then(function (token) {
              data.googleCaptcha = token
              isLoading.value = true
              apiCheckPhoneCode(data).then(res => {
                if (res.data.status == 1) {
                  apiUserbaseUpdatePhone(params).then(res => {
                    router.push('/personal-wditor')
                  })
                }
              }).finally(() => {
                isLoading.value = false
              })
            });
      })
    }
  } else if (type === 'mail') {
    if (user.value.code > 5)
      grecaptcha.enterprise.ready(function () {
        grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
            .then(function (token) {
              let params = {
                email: props._email,
                googleCaptcha: token,
                code: user.value.code,
              }
              let param = {
                newEmail: props._email,
                code: user.value.code,
              }
              isLoading.value = true
              apiCheckEmailCode(params).then((res) => {
                codeError.value = false
                apiChangeeMailWithCode(param).then(sonRes => {
                  if (sonRes.data.status == 1) {
                    router.push('/personal-wditor')
                  }
                })
              }).catch(err => {
                codeError.value = true
              }).finally(() => {
                isLoading.value = false
              })
            });
      })
  } else {
    if (user.value.code?.length > 5) {
      data.email = props._email
      grecaptcha.enterprise.ready(function () {
        grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
            .then(function (token) {
              data.googleCaptcha = token
              isLoading.value = true
              apiAccountChangeEmailCode(data).then(res => {
                if (res.data.status == 1) {
                  useState('_showModalPasswordWmail').value = 3
                  userCode.value = user.value.code
                }
              }).finally(() => {
                isLoading.value = false
              })
            })
      })
    }
  }
}

</script>

<style>
.input-error {
  .el-input__wrapper {
    box-shadow: 0 0 0 1px #C42A30 inset !important;
  }
}
</style>

<style lang="scss" scoped>

.confirm-your-email {
  font-size: 30px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: normal;
  color: #030100;
  line-height: 35px;
}

.input-sent-to {
  font-size: 19px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #010101;
  line-height: 25px;
  margin-top: 15px;
  margin-bottom: 24px;
}

.resend-code {
  height: 25px;
  font-size: 19px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: normal;
  line-height: 25px;
  margin-top: 16px;
  color: #0A66C2;
}

.resend-code-disabled {
  color: #828282;
}

.sunmit-btn {
  font-size: 24px;
  border: 0;
}

.email-input {
  height: 68px;
  font-size: 22px;
}

.email-input-hint {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.email-input-err {
  height: 25px;
  font-size: 19px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: normal;
  color: #C42A30;
  line-height: 25px;
  -webkit-background-clip: text;
  margin-top: 16px;
}

.see-this-email {
  height: 52px;
  font-size: 18px;
  font-family: Inter-Medium, Inter;
  font-weight: 500;
  color: #5E5E5E;
  line-height: 24px;
  margin-top: 34px;
  margin-bottom: 16px;
}

.inaccessible-email {
  height: 24px;
  font-size: 20px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: normal;
  color: #0A66C2;
  line-height: 24px;
}

.cv-button {
  width: 100%;
  height: 70px;
  border-radius: 40px !important;
}

.back {
  margin-top: 27px;
  text-align: center;
  height: 24px;
  font-size: 21px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: normal;
  color: #575757;
  line-height: 24px;
  cursor: pointer;
}

@media (max-width: 420px) {
  .back {
    height: 46px;
    font-size: 16px;
    margin-top: 13px;
  }
}

</style>
