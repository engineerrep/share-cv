<template>
  <n-card title="">
    <n-form ref="formRef" :label-width="80" :model="formValue" :rules="rules">
      <n-form-item :label="$t('login.全名')" path="fullName">
        <n-input
          class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
          v-model:value="formValue.fullName"
          :placeholder="$t('login.请输入全名')"
        />
      </n-form-item>
      <n-form-item :label="$t('login.昵称')" path="nickName">
        <n-input
          class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
          v-model:value="formValue.nickName"
          :placeholder="$t('login.请输入昵称')"
        />
      </n-form-item>
      <n-form-item>
        <n-button
          @click="handleValidateClick"
          type="primary"
          round
          class="w-full h-50px px-4 bg-[#219653]"
        >
          <span class="text-18px text-[#FDFDFD]">
            {{ $t("login.继续") }}
          </span>
        </n-button>
      </n-form-item>
    </n-form>
  </n-card>
</template>

<script setup>
import { NCard, NForm, NFormItem, NInput, NButton } from "naive-ui";
import { useI18n } from "vue-i18n";
const { t } = useI18n();
const props = defineProps({
    formData:Object
})
const emit = defineEmits(["setUpShowIndex"]);

const formRef = ref(null);
const loading = ref(false);
// const message = useMessage();
/**
 * 表单数据
 */
const formValue = ref({
  fullName: "",
  nickName: "",
});
/**
 * 登录规则
 */
const rules = {
  fullName: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("login.请输入全名"));
      }
      return true;
    },
  },
  nickName: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("login.请输入昵称"));
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
  loading.value = true;
  e.preventDefault();
  formRef.value?.validate((errors) => {
    if (!errors) {
      console.log("这里调接口");
      emit("setUpShowIndex", 3, formValue.value);
      setTimeout(() => {
        loading.value = false;
      }, 1500);
    }
  });
};
</script>

<style lang="scss" scoped>
.n-card {
  width: 600px !important;
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

@media (max-width: 480px) {
  .n-card {
    width: 90% !important;
    margin: 0 auto;
    border-radius: 10px;
  }
}
</style>
