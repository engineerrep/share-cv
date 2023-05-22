<template>
  <NMessageProvider>
    <div class="forgot-password box-middle-h5 position-center forgot-password" id="forgotPassword">
      <div v-if="showModalPasswordWmail===1" v-loading="loading">
        <div class="confirm-your-email">{{ title }}</div>
        <div class="send-to">
          <img :src="Icon" alt="">

          <p class="">{{ subtitle }}</p>
        </div>

        <div class="input forgot-password-input" :class="mistake?'mistake':''">
          <el-input
              v-if="type==='phone'"
              v-model="userForm.phone"
              placeholder="Phone number"
              class=" h-70px"
              maxlength="11"
          >
            <template #prepend>
              <el-select v-model="userForm.selected"
                         class="input-select mb-26px"
                         style="width: 100px;height: 70px;"
                         placeholder=" "
                         id="codeSelect"
              >
                <el-option
                    v-for="item in phoneNumList"
                    :key="item.code"
                    :label="'+'+item.num"
                    :value="item.code"
                    :data-id="item.id"
                    @click="changeSelect"
                > +{{ item.num }} {{ item.name }}
                </el-option>
              </el-select>
            </template>
          </el-input>

          <div v-else-if="type==='mail'" class="relative">
            <el-input
                class="email-input change-email-input"
                v-model="userForm.Email"
                placeholder=" "
                @blur="blurVal(refEmail)"
                @focus="focuslVal(refEmail)"
                @input="verifyVal"
            />
            <label class="email-label" :class="mistake?'input-error':''" ref="refEmail">Email</label>
            <div class="absolute input-error" v-if="mistake">Please enter a valid email address</div>
          </div>
          <el-input
              v-else
              class="email-input"
              v-model="userForm.Email"
              :placeholder="placeholder"
          />
        </div>

        <div class="my-30px">
          <CVButton @click="submitRequest" v-loading="isLoading">
            <template #main>
              <span>  {{ submitTxt }}</span>
            </template>
          </CVButton>
          <div class="back"><span @click="router.push('/')">Back</span></div>
        </div>
      </div>

      <ForgotPasswordWmail
          v-if="showModalPasswordWmail===2"
          :formData="formData"
          :userForm="userForm"
          :_email="userForm.Email"
      ></ForgotPasswordWmail>
      <ForgotPasswordNewPassword
          v-if="showModalPasswordWmail===3"
          :formData="formData"
          :userForm="userForm"
          :_email="userForm.Email"
      ></ForgotPasswordNewPassword>
    </div>
  </NMessageProvider>

</template>

<script setup>
import {NButton, NMessageProvider} from "naive-ui";
import {useI18n} from "vue-i18n";
import {useState} from "nuxt/app";
import {useRouter} from "vue-router";
import ForgotPasswordWmail from "@/components/forgot-password/forgot-password-email.vue";
import ForgotPasswordNewPassword from "@/components/forgot-password/forgot-password-new-password.vue";
import {apiCommonCountry, apiRegisterCheckEmail, sendCode} from "~/api/login";
import {apiSendPhoneCode} from "~/api/personal-wditor";
import CVButton from '@/components/form/button.vue';

const router = useRouter();

const {t} = useI18n();
const props = defineProps({
  formData: Object,
  _email: String,
});
const formData = ref({});
const emit = defineEmits(["setUpShowIndex"]);
let userInformation = useCookie("_userInformation")
// 表单-数据
let userForm = ref({
  Email: null,
  phoneNum: null,
  selected: 'US',
  phone: null,
  countryId: 238,
})
let {type} = router.currentRoute.value.query
if (type === 'phone') userForm.value.phone = userInformation.value.phone

watch(() => useState('_resetEmail').value, (newVal) => {
  if (newVal) {
    userForm.value.Email = null
  }
}, {immediate: true})
let isLoading = ref(false)
// 标题
let title = ref('Forgot password?')
// 子标题
let subtitle = ref('Reset password in two quick steps')
// 输入框提示
let placeholder = ref('Email')
// 提交按钮文字
let submitTxt = ref('Reset password')
// 手机号前缀-列表
let phoneNumList = ref([])
let refEmail = ref(null)

// 当移出input的时候触发
const blurVal = (dom) => {
  if (document.body.clientWidth >= 410) {
    if (userForm.value.Email) {
      dom.setAttribute('style', 'top:10px;font-size:16px')
      document.querySelector('.change-email-input .el-input__inner')
          .setAttribute('style', 'font-size:20px')
    } else {
      dom.setAttribute('style', 'left: 20px;top:20px;font-size: 22px;')
    }
  } else {
    if (userForm.value.Email) {
      dom.setAttribute('style', 'font-size:12px;top:0px')
    } else {
      dom.setAttribute('style', 'top: 10px;')
    }
  }
  if (type === 'mail') {
    verifyVal()
  }
}
// 当移入input的时候触发
const focuslVal = (dom) => {
  if (document.body.clientWidth >= 410) {
    dom.setAttribute('style', 'top:10px;font-size:16px')
  } else {
    dom.setAttribute('style', 'font-size:12px;top:0px')
  }
  if (type === 'mail') {
    verifyVal()
  }
}

const verifyVal = () => {
  userForm.value.Email = userForm.value?.Email?.trim()
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (emailRegex.test(userForm.value.Email)) {
    mistake.value = false
  } else {
    mistake.value = true
  }
}

/**
 * 判断是邮箱是手机号还是忘记密码
 * @param {type} phone-更改添加电话
 * @param {type} mail-更改邮箱
 * @param {type} 没有就是设置密码
 * */
if (type === 'phone') {
  title.value = 'Just a quick security check'
  subtitle.value = 'As an extra security step, we’ll need to give you a verification code.'
  placeholder.value = 'Phone'
  submitTxt.value = 'Submit'
} else if (type === 'mail') {
  title.value = 'New email address'
  subtitle.value = 'Replace it with your commonly used email'
  placeholder.value = 'Email'
  submitTxt.value = 'Send Code'
}
if (type === 'phone') {
  apiCommonCountry().then(res => {
    phoneNumList.value = JSON.parse(JSON.stringify(res.data.list))
  })
}

let showModalPasswordWmail = useState('_showModalPasswordWmail', () => 1)
showModalPasswordWmail.value = 1

const changeSelect = (event) => {
  userForm.value.countryId = event.target.getAttribute("data-id")
}
let loading = useState('_loading')
let mistake = ref(false)
/**
 * 重置密码,添加手机号-改变邮箱-提交
 */
const submitRequest = () => {
  if (type) {
    if (type === 'phone') {
      let phoneNum = document.querySelector('#codeSelect').value.replace("+", "");
      userForm.value.phoneNum = phoneNum
      if (userForm.value.phone.length <= 11 && Number(userForm.value.phone)
          && phoneNum) {
        showModalPasswordWmail.value = 2
        grecaptcha.enterprise.ready(function () {
          grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
              .then(function (token) {
                let params = {
                  phone: userForm.value.phone,
                  phoneNum: phoneNum,
                  googleCaptcha: token
                }
                apiSendPhoneCode(params).then((res) => {

                })
              });
        })
      }
    } else if (type === 'mail') {
      console.log(12)
      if (userForm.value.Email) {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(userForm.value.Email)) {
          mistake.value = true
          return
        }
        mistake.value = false
        loading.value = true
        grecaptcha.enterprise.ready(function () {
          grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
              .then(function (token) {
                let params = {
                  email: userForm.value.Email,
                  googleCaptcha: token
                }
                apiRegisterCheckEmail(params).then((res) => {
                  loading.value = false
                  if (res.data.status) {
                    grecaptcha.enterprise.ready(function () {
                      grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
                          .then(function (token) {
                            let params = {
                              email: userForm.value.Email,
                              googleCaptcha: token
                            }
                            sendCode(params).then((resSon) => {
                              showModalPasswordWmail.value = 2
                            })
                          });
                    })
                  }
                }).catch(err => {
                  loading.value = false
                })

              });
        })
      }
    }
  } else {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (emailRegex.test(userForm.value.Email))
      grecaptcha.enterprise.ready(function () {
        grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
            .then(function (token) {
              let params = {
                email: userForm.value.Email,
                googleCaptcha: token
              }
              isLoading.value = true
              sendCode(params).then((res) => {
                if (res.data.status == 1) {
                  showModalPasswordWmail.value = 2
                }
              }).finally(() => {
                isLoading.value = false
              })
            });
      })
  }
}

</script>

<style lang="scss">
.mistake {
  .el-input__wrapper {
    box-shadow: 0 0 0 1px #C42A30 inset !important;
  }
}

.change-email-input {
  .el-input__inner {
    position: absolute;
    top: 30px;
    left: 20px;

    &:focus {
      font-size: 20px;
    }
  }
}

#forgotPassword {
  .input-select {
    .el-select__icon {
      svg {
        color: #616161;
      }
    }

    .el-input .el-input__wrapper {
      height: 70px;
      position: relative;
      top: 0px;
      box-shadow: 1px 0 0 0 #616161 inset,
      0 1px 0 0 #616161 inset,
      0 -1px 0 0 #616161 inset !important;
    }
  }

  .el-input.is-focus .el-input__wrapper {
    box-shadow: 1px 0 0 0 #616161 inset,
    1px 0 0 0 #616161,
    0 1px 0 0 #616161 inset,
    0 -1px 0 0 #616161 inset !important;
  }
}

@media (max-width: 420px) {
  .input-select .el-input__wrapper, .el-input, .el-select {
    height: 44px !important;
  }

  .el-input__inner {
    font-size: 14px;
  }

  .change-email-input {
    .el-input__inner {
      top: 10px;
      left: 12px;

      &:focus {
        font-size: 14px;
      }
    }
  }

}

</style>

<style lang="scss" scoped>
.forgot-password {
  width: 534px;
  padding: 30px;
  margin: auto;
  background-color: #fff;
  border: 1px solid #eeeeee;
  border-radius: 14px 14px 14px 14px;
  box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0.25);
  opacity: 1;
}

.centerStyle {
  font-size: 14px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #333333;
  line-height: 19px;
}

.email-input {
  height: 69px;
  border-radius: 6px;

  font-size: 22px;
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

.reset-password-btn {
  font-size: 21px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: normal;
  color: #FDFDFD;
  line-height: 24px;
}

.inputValue {
  border-bottom: 2px solid #4d4d4d;
  width: 50px;
  outline: none;
}

.confirm-your-email {
  height: 39px;
  font-size: 32px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: normal;
  color: #030100;
  line-height: 38px;
}

.send-to {
  height: 24px;
  font-size: 20px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #505050;
  line-height: 23px;
  margin-top: 21px;
  margin-bottom: 44px;
}

.el-select, .select-trigger {
  height: 100px;

  .el-input__wrapper {
    text-align: center;

  }
}

.input {
  .input-select {
    background-color: #fff;
  }
}

.input .input-select .select-trigger {
  background-color: #ddd;
}

.email-label {
  position: absolute;
  left: 20px;
  top: 20px;
  height: 27px;
  font-size: 22px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #525152;
  line-height: 26px;
  -webkit-background-clip: text;
}

.input-error {
  color: #C42A30;
  font-size: 14px;
  line-height: 1;
  padding-top: 2px;
}

.cv-button {
  height: 69px;
  width: 100%;
  border-radius: 40px !important;
}

@media (max-width: 420px) {
  .confirm-your-email {
    font-size: 24px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #000000;
    line-height: 28px;
  }
  .el-input {
    font-size: 14px;
  }

  .forgot-password {
    box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0);
  }
  .send-to {
    font-size: 16px;
    margin-top: 14px;
    margin-bottom: 28px;
  }
  .email-input,
  .reset-password-btn,
  .back {
    height: 46px;
    font-size: 16px;
  }
  .email-label {
    top: 10px;
    left: 12px;
    font-size: 16px;
  }
}
</style>
