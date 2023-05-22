<template>
  <n-card title="" v-loading="loading">
    <div class="confirm-your-email">Confirm your email</div>
    <div class="send-to">
      <p class="">Type in the code we sent to </p>
      <p class="">{{ props._email }}</p>
    </div>

    <div
        class="border rounded-6px border-[#616161] px-15px py-15px mb-30px w-full text-center"
    >
      <codeInput :value="codeInputVal" @finish="codeInputFinish"></codeInput>
      <!--      <input-->
      <!--          v-for="(item, index) in inputList"-->
      <!--          type="text"-->
      <!--          maxlength="1"-->
      <!--          v-model="item.val"-->
      <!--          @keyup="inputNextFocus($event, index)"-->
      <!--          @keydown="cancelValue(index)"-->
      <!--          class="border-bottom text-center inputValue"-->
      <!--          :class="index === 0 ? '' : 'ml-10px'"-->
      <!--          :key="index"-->
      <!--      />-->
    </div>

    <div class="border rounded-6px border-[#ECECEC] px-15px py-15px mb-25px">
      <div class="flex items-start mb-10px">
        <img
            src="@/assets/img/register/privacy-img.png"
            alt=""
            class="h-24px w-24px mr-15px"
        />
        <div>
          <div class="text-18px mb-10px font-bold your-privacy">
            {{ $t("register.Your privacy is important") }}

          </div>
          <div class="text-16px leading-21px promotional-information">
            <div> {{ $t('register.promotional information01') }}</div>
            <div> {{ $t('register.promotional information011') }}</div>
            <div> {{ $t('register.promotional information012') }}
              <span class="preferences"> {{ $t('register.promotional information02') }}</span>
            </div>
            <div> {{ $t('register.promotional information03') }}</div>
          </div>
        </div>
      </div>
      <div></div>
    </div>
    <div class="send-again flex items-center justify-center">
      <div class="mr-20px font-bold font-20px">
        {{ $t("register.Didnt receive the code") }}
      </div>
      <div class="send-again-btn font-18px cursor-pointer c-theme" @click="handleAgain">
        {{ $t("register.Send again") }}
      </div>
    </div>
    <div class="my-30px">
      <CVButton @click="handleValidateClick">
        <template #main>
              <span class="text-18px text-[#FDFDFD]">
                {{ $t("register.Agree & Confirm") }}
              </span>
        </template>
      </CVButton>
    </div>
  </n-card>
</template>

<script setup>
import {NCard, useMessage} from "naive-ui";
import {useI18n} from "vue-i18n";
import {apiCheckEmailCode} from "@/api/personal-wditor";
import {useState} from "nuxt/app";
import {apiEmailRegister, sendCode} from "~/api/login";
import CVButton from '@/components/form/button.vue';
import {debounce} from "~/utils/util";
import codeInput from "./codeInput.vue"

let loading = useState("_loading")
const message = useMessage()
const {t} = useI18n();
const props = defineProps({
  formData: Object,
  _email: String,

});
let codeInputVal = ref('')
const emit = defineEmits(["setUpShowIndex", 'isModalEmit']);

const inputList = ref([
  {val: ""},
  {val: ""},
  {val: ""},
  {val: ""},
  {val: ""},
  {val: ""},
]);

const isCanSubMit = computed(() => {
  return inputList.value.filter((item) => !item.val).length === 0;
});
let pcSide = useState('_pcSide')
// input自动切换到下一个、
const inputNextFocus = debounce((el, index) => {
  const dom = document.getElementsByClassName("inputValue"), // 引入inputValue
      currInput = dom[index],
      nextInput = dom[index + 1],
      lastInput = dom[index - 1];
  if (el.keyCode != 8) {
    if (index < inputList.value.length - 1) {
      if (!pcSide.value) {
        let tiemout = setTimeout(() => {
          nextInput.focus();
          clearTimeout(tiemout)
        }, 300)
      } else nextInput.focus();
    } else {
      currInput.blur();
    }
  } else {
    if (index != 0) {
      if (!pcSide.value) {
        let tiemout = setTimeout(() => {
          lastInput.focus();
          clearTimeout(tiemout)
        }, 300)
      } else lastInput.focus();


    }
  }
}, 10);
const cancelValue = (index) => {
  inputList.value[index].val = "";
};
onMounted(() => {
  // let dom = document.querySelectorAll('.inputValue')
  document.addEventListener('paste', function (event) {
    let clipboardData = event.clipboardData;
    let pastedText = clipboardData.getData('text') + ''
    let val = Number(pastedText)
    if (val > 99999) {
      codeInputVal.value = String(val).substring(0, 6)
      // for (let i = 0; i < pastedText.length; i++) {
      //   inputList.value[i].val = pastedText[i]
      //   dom[i].value = pastedText[i]
      // }
    }
  })
})
let codeVal = ref('')
let codeValRef = ref(null)

const focusCodeInput = () => {
  codeValRef.value.focus()
}
/**
 * 再次发送
 * @param {*} e
 */
const handleAgain = () => {
  useState('_verifyEmail').value = false
  loading.value = true
  grecaptcha.enterprise.ready(function () {
    grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
        .then(function (token) {
          let params = {
            email: props._email,
            googleCaptcha: token
          }
          // showModal.value = true
          sendCode(params).then((res) => {
            loading.value = false
            if (res.data.status == 1) {
              showModal.value = true
              emit("isModalEmit", false)
            }
          })
        });
  })
};
let userInformation = useCookie('_userInformation')
/**
 * 验证码输入完成
 * @param val
 */
const codeInputFinish = (val) => {
  codeInputVal.value = val
}
/**
 * 点击继续
 * @param {*} e
 */
const handleValidateClick = (e) => {
  e.preventDefault();
  let code = JSON.parse(JSON.stringify(codeInputVal.value)) + '';
  // for (let i = 0; i < inputList.value.length; i++) {
  //   code += `${inputList.value[i].val}`
  // }
  if (code.length !== 6) {
    return
  }
  loading.value = true;
  let prams = {
    code,
    email: props._email,
  }
  let params = {
    code,
    email: props._email,
    password: props.formData.Password,
  }
  grecaptcha.enterprise.ready(function () {
    grecaptcha.enterprise.execute('6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv', {action: 'login'})
        .then(function (token) {
          prams.googleCaptcha = token
          apiCheckEmailCode(prams).then((res) => {
            if (res.data && res.data.status) {
              apiEmailRegister(params).then(res => {
                if (res.data.accessToken && res.data) {
                  useCookie('token').value = res.data.accessToken
                  document.cookie = "token=" + res.data.accessToken;
                  emit("isModalEmit", false)
                  useState('_loading').value = false
                  useState("_showModal").value = false;
                  useState('_shwoIndex').value = 1
                  useCookie('_userInformation').value = {
                    phone: null,
                    email: props._email,
                  }
                } else if (res.code == 1020) {
                  ElMessage({
                    showClose: true,
                    message: 'Password reset successful',
                    type: 'error',
                  })
                }
              })
            }
          }).catch(err => {
            loading.value = false;
          })
        });
  })

};
</script>

<style lang="scss" scoped>
.n-card {
  padding-top: 30px;
  margin: 0 auto;
  border-radius: 10px;
}

.centerStyle {
  font-size: 14px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #333333;
  line-height: 19px;
}

.inputValue {
  border-bottom: 1px solid #4d4d4d;
  width: 50px;
  outline: none;
}

.preferences {
  //color: #005bb6;
}

.promotional-information {
  width: 526px;
  height: 84px;
  font-size: 16px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #5C5C5C;
  line-height: 21px;
}

.confirm-your-email {
  height: 39px;
  font-size: 32px;
  font-family: Inter-Semi Bold, Inter;
  font-weight: normal;
  color: #030100;
  line-height: 38px;
  text-align: center;
}

.send-to {
  height: 24px;
  font-size: 20px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #505050;
  line-height: 23px;
  text-align: center;
  margin-top: 21px;
  margin-bottom: 68px;
  justify-content: center;
}

.cv-button {
  width: 100%;
  height: 69px;
  border-radius: 40px !important;
}

@media (max-width: 400px) {
  .confirm-your-email {
    font-size: 26px;
  }
  .send-to {
    font-size: 16px;
    margin-top: 6px;
    margin-bottom: 48px;
  }
  .your-privacy {
    font-size: 14px;
  }
  .promotional-information {
    font-size: 14px;
    width: 240px;
    height: auto;
  }
  .send-again {
    font-size: 14px;
  }
  .cv-button {
    height: 40px;
    font-size: 18px;
    background: #0A66C2;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    border-radius: 39px 39px 39px 39px;
  }
}

.send-again-btn:hover {
  color: #003466;
}

.send-again-btn:active {
  color: #0d1d2c
}

@media (max-width: 420px) {

  .h5-code-item {
    //position: absolute;
    width: 20px;
    height: 20px;
    border-bottom: 2px solid #4d4d4d;
    margin-right: 10px;
  }
}

@media (max-width: 480px) {
  .n-card {
    width: 90% !important;
    margin: 0 auto;
    border-radius: 10px;
  }
  .inputValue {
    border-bottom: 2px solid #4d4d4d;
    width: 20px;
    outline: none;
  }
}
</style>
