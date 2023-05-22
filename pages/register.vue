<template>
  <div class=" center02 ">
    <template v-if="shwoIndex === 0 || shwoIndex === 1 || shwoIndex === 2">
      <div class="registerTitle mb-6 text-center">
        {{ $t("register.title01") }}
      </div>
    </template>
    <template v-if="shwoIndex === 3">
      <div class="registerTitle mb-6 text-center">
        <div>
          {{ $t("register.welcome") }} &nbsp;
          {{ userMsg?.firstName }}!
        </div>
        <div class="registerTitle-subtitle">
          Let's start your profile, connect to people you know, and engage with them on
          <div>topics you care about.</div>
        </div>
      </div>
    </template>
    <template v-if="shwoIndex === 4">
      <div class="registerTitle mb-6 text-center">
        {{ $t("register.personal information") }}
      </div>
    </template>
    <template v-if="shwoIndex === 5">
      <div class="registerTitle mb-6 text-center">
        {{ $t("register.Country/Region") }}
      </div>
      <div class="smallTitle">
        {{ $t("register.send code") }} jsmith.mobbin@gmail.com
      </div>
    </template>
    <template v-if="shwoIndex === 6">
      <div class="registerTitle mb-6 text-center">
        {{ $t("register.are you looking for a new job") }}
      </div>
      <div class="smallTitle">
        {{ $t("register.personal search") }}
      </div>
    </template>
    <template v-if="shwoIndex === 7">
      <div class="registerTitle mb-6 text-center">
        {{ $t("register.find a job") }}
      </div>
      <div class="smallTitle">
        {{ $t("register.label position") }}
      </div>
    </template>

    <div v-if="shwoIndex === 0">
      <OneSinup
          @setUpShowIndex="setUpShowIndex"
          :formData="formData"
      ></OneSinup>
    </div>
    <div v-if="shwoIndex === 1">
      <TwoSetLastName
      ></TwoSetLastName>
    </div>
    <div v-if="shwoIndex === 3">
      <FourSetAddress
          @setUpShowIndex="setUpShowIndex"
          :formData="formData"
      ></FourSetAddress>
    </div>
  </div>
</template>

<script setup>
import OneSinup from "@/components/register/oneSinup.vue";
import TwoSetLastName from "@/components/register/twoSetLastName.vue";
import FourSetAddress from "@/components/register/fourSetAddress.vue";
import {useI18n} from "vue-i18n";
import {useState} from "nuxt/app";

let userMsg = useState('_userMsg')
const {t, locale} = useI18n();


onMounted(() => {
  // userMsg.value = JSON.parse(localStorage.getItem('userMsg'))
  // console.log(111,JSON.parse(localStorage.getItem('userMsg')).firstName)
  // console.log( userMsg.value.firstName )Ï
})


// 基础数据
const formData = ref({});
// 步骤下标
let shwoIndex = useState(('_shwoIndex'), () => 0)
let _login = useState(('_login'), () => 0)
if (_login) {
  shwoIndex.value = _login.value
}

const setUpShowIndex = (index, data) => {
  shwoIndex.value = index;
  console.log(index)
  formData.value = {...data, ...formData.value};
};
// 改变验证提示
watch(locale, (newVal, oldVal) => {
  // reviseDom('.name-input .n-form-item-feedback__line', t('register.please enter a name'))
  // reviseDom('.surname-input .n-form-item-feedback__line', t('register.Please enter a last name'))
})

useHead({
  title: t("seo.home.title"),
  meta: {
    keywords: t("seo.home.keywords"),
    description: t("seo.home.description"),
  },
});
</script>

<style>
@media (max-width: 420px) {
  .cv-button {
    height: 40px !important;
    font-size: 18px !important;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    line-height: 21px;
  }
}
</style>
<style lang="scss" scoped>
@import "assets/styles/register.scss";
</style>
