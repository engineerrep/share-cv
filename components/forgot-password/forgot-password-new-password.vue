<template>

  <div title="">
    <div class="confirm-your-email">Choose a new password.</div>
    <div class="input-sent-to">
      <p class="">Create a new password that is at least 8 characters long. </p>
      <div>
        <!--        <span>long.</span>-->
        <!--        <span class="strong-passwor"> What makes a strong password?</span>-->

      </div>
    </div>

    <!-- 邮箱，密码，登录 -->
    <div class="sign-in-form pr">
      <div class="pr form-item">
        <input
            v-model="passwordVal.newPassword"
            class="email-password-input"
            required="required"
            :type="pwdInputType"
            @input="checkTwoPasswords"
        />
        <label class="email-label">New password</label>
        <!--          <label class="email-label">{{ $t("home.New password") }}</label>-->
        <img
            v-if="pwdState"
            alt=""
            class="pwd-state"
            src="@/assets/img/pwd-hide.png"
            @click="changePwdState"
        />
        <img
            v-else
            alt=""
            class="pwd-state"
            src="@/assets/img/pwd-show.png"
            @click="changePwdState"
        />
      </div>

      <div class="pr form-item password-item">
        <div class="">
          <input
              :type="newPwdInputType"
              v-model="passwordVal.retypeNewPassword"
              class="email-password-input"
              required="required"
              @input="checkTwoPasswords"
          />
          <label class="email-label">Retype new password</label>
          <img
              v-if="retypeNewPwdState"
              alt=""
              class="pwd-state"
              src="@/assets/img/pwd-hide.png"
              @click="changenewPwdState"
          />
          <img
              v-else
              alt=""
              class="pwd-state"
              src="@/assets/img/pwd-show.png"
              @click="changenewPwdState"
          />
        </div>
      </div>
    </div>

    <div class="my-30px pr pt-6px">
      <div class="pa top-n16 text-red-600" v-if="twoDifferentPasswords">

        {{ $t('register.password prompt twice') }}
      </div>
      <div class="mb-44px">
        <!--        <el-checkbox v-model="agree" size="large" class="mr-6px " ref="refAgree"/>-->
        <!--        <span class="" :style="{color:agreeColor}">-->
        <!--       Require all devices to sign in with new password-->
        <!--        </span>-->
      </div>

      <CVButton @click="submitNewPassword" v-loading="isLoading">
        <template #main>
          <span> Submit</span>
        </template>
      </CVButton>

    </div>

  </div>
</template>

<script setup>
import {NButton, useMessage} from "naive-ui";
import {useI18n} from "vue-i18n";
import {apiAccountForgetPassword} from "~/api/user";
import CVButton from '@/components/form/button.vue';
import {useRouter} from "vue-router";

const router = useRouter();

const message = useMessage()
const {t} = useI18n();

const props = defineProps({
  formData: Object,
  _email: String,
});
let isLoading = ref(false)
const formData = ref({});
const emit = defineEmits(["setUpShowIndex"]);
let user = ref({
  Email: null
})
let showModalSixSendCode = ref(false)
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
// 密码状态
let retypeNewPwdState = ref(true);
// 密码输入框类型
let newPwdInputType = ref("password");
// 改变密码状态
const changenewPwdState = () => {
  if (retypeNewPwdState.value) {
    retypeNewPwdState.value = false;
    newPwdInputType.value = "txt";
  } else {
    retypeNewPwdState.value = true;
    newPwdInputType.value = "password";
  }
};

// 是否同意协议
let agree = ref(false)
let refAgree = ref()
// 重新设置密码
let passwordVal = ref({
  newPassword: null,
  retypeNewPassword: null,
})

let verifyPassword = ref({
  newPassword: false,
  retypeNewPassword: false,
})

let agreeColor = ref('')
let twoDifferentPasswords = ref(false)

const VerifyPasswordTwice = (obj, greaterOrEqualTo) => {
  if (obj.newPassword === obj.retypeNewPassword && (obj.newPassword + '').length >= greaterOrEqualTo) {
    return true;
  } else {
    return false;
  }
}
/**
 * 检查两次输入的密码是否一致
 */
const checkTwoPasswords = () => {
  let newPassword = passwordVal.value.newPassword
  let retypeNewPassword = passwordVal.value.retypeNewPassword
  twoDifferentPasswords.value = (newPassword && retypeNewPassword && (newPassword === retypeNewPassword)) ? false : true
}
// 提交
const submitNewPassword = () => {
  let newPassword = passwordVal.value.newPassword
  let retypeNewPassword = passwordVal.value.retypeNewPassword
  if (newPassword.length >= 6 && retypeNewPassword.length >= 6 & retypeNewPassword === newPassword) {
    twoDifferentPasswords.value = false
    let data = {
      "email": props._email,
      "password": passwordVal.value.newPassword,
      "code": useState("_userCode").value,
    }
    isLoading.value = true
    apiAccountForgetPassword(data).then(res => {
      if (res.data && res.data.status) {
        router.push('/')
      }
    }).finally(() => {
      isLoading.value = false
    })
  } else {
    twoDifferentPasswords.value = true
  }

}
</script>


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

.strong-passwor {
  color: #0051a9;
}

.pr {
  position: relative;
}

.sign-in-form {
  position: relative;

  label {
    position: absolute;
    top: 34px;
    left: 16px;
    min-width: 545px;
    color: #000;
    font-size: 16px;
    line-height: 48px;
    pointer-events: none;
    transition: all 0.5s;
  }

  .pwd-label {
    top: 24px;
  }
}

.pwd-state {
  position: absolute;
  top: 42px;
  right: 10px;
  font-family: Inter-Medium, Inter;
  font-weight: 500;
  color: #4b4b4b;
  cursor: pointer;
  width: 40px;
}

.form-item input:focus + label,
.form-item input:valid + label {
  font-size: 18px;
  top: 30px;
}

.sign-in-with-google {
  color: #616363;
  border: 1px solid #000000;
  position: relative;
}

.email-password-input {
  width: 100%;
  height: 68px;
  padding-top: 34px;
  padding-left: 16px;
  border-radius: 6px;
  opacity: 1;
  border: 1px solid #E0E0E0;
  margin-top: 25px;
  font-size: 16px;

  &:hover {
    border: 2px solid #60ace5;
  }
}


.email-password-input02 {
  margin-top: 16px;
}

.resend-code {
  height: 25px;
  font-size: 19px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: normal;
  color: #0A66C2;
  line-height: 25px;
  margin-top: 16px;
}

.sunmit-btn {
  background: #005bb6;
  font-size: 24px;
}

.email-input {
  height: 68px;
  font-size: 22px;
  letter-spacing: 30px;

}

.email-input-hint {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.email-input-err {
  display: none;
  color: red;
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
  height: 69px;
  border-radius: 40px !important;
}
</style>
