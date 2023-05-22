<template>
  <n-card title="">
    <n-form ref="formRef" :label-width="80" :model="formValue" :rules="rules">
      <n-form-item :label="$t('register.Name')" path="firstName" class="name-input">
        <n-input
            class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
            v-model:value="formValue.firstName"
            :placeholder="$t('register.please enter a name')"
        />
      </n-form-item>

      <n-form-item :label="$t('register.surname')" path="lastName" class="surname-input">
        <n-input
            class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
            v-model:value="formValue.lastName"
            :placeholder="$t('register.Please enter a last name')"
        />
      </n-form-item>
      <n-form-item>
        <CVButton @click="handleValidateClick">
          <template #main>
              <span class="text-18px text-[#FDFDFD]">
                {{ $t("register.Agree & Confirm") }}
              </span>
          </template>
        </CVButton>
      </n-form-item>
    </n-form>
  </n-card>
</template>

<script setup>
import {NButton, NCard, NForm, NFormItem, NInput} from "naive-ui";
import {useI18n} from "vue-i18n";
import {reviseDom} from "@/utils/util";
import {apiUpdateName} from "@/api/personal-wditor";
import {ElMessage} from 'element-plus'
import CVButton from '@/components/form/button.vue';
import {useState} from "nuxt/app";

const {t, locale} = useI18n();
const formRef = ref(null);
// const message = useMessage();
/**
 * 表单数据
 */
const formValue = ref({
  firstName: "",
  lastName: "",
});
// 改变验证提示
watch(locale, (newVal, oldVal) => {
  reviseDom('.name-input .n-form-item-feedback__line', t('register.please enter a name'))
  reviseDom('.surname-input .n-form-item-feedback__line', t('register.Please enter a last name'))
})
/**
 * 登录规则
 */
const rules = {
  firstName: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("register.please enter a name"));
      }
      return true;
    },
  },
  lastName: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("register.Please enter a last name"));
      }
      return true;
    },
  },
};
/**
 * 点击继续
 * @param {*} e
 */
const handleValidateClick = (e) => {
  e.preventDefault();
  formRef.value?.validate((errors) => {
    if (!errors) {
      // console.log('填写姓名-这里调接口')

      let params = {
        firstName: formValue.value.firstName,
        lastName: formValue.value.lastName,
      }

      apiUpdateName(params).then(res => {
        console.log('填写姓名-res', res.data)
        if (res.data.status) {
          useState('_userMsg').value = params
          useState('_shwoIndex').value = 3
        }
      })
    }
  });
};
</script>

<style lang="scss" scoped>
.cv-button {
  width: 100%;
  height: 62px;
  border-radius: 30px;
}

.n-card {
  width: 600px;
  padding-top: 30px;
  margin: 0 auto;
  border-radius: 10px;
  box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0.25);
}

.centerStyle {
  font-size: 14px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #333333;
  line-height: 19px;
}

@media (max-width: 400px) {
  .n-input {
    font-size: 14px;
  }
  .n-card.n-card--bordered {
    border: 0;
  }
  .n-card {
    box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0);
  }
  .cv-button {
    height: 40px;
  }
}

@media (max-width: 480px) {
  .n-card {
    width: 90% !important;
    margin: 0 auto;
    border-radius: 10px;
  }
}
</style>
