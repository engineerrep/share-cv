<template>
  <n-card title="">
    <div class="registerTitle mb-1">
      {{ $t("forget.选择新密码。") }}
    </div>
    <div class="smallTitle">
      {{ $t("forget.创建至少8个字符长的新密码。") }}
    </div>
    <n-form ref="formRef" :label-width="80" :model="formValue" :rules="rules">
      <n-form-item path="password">
        <n-input
          class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
          type="password"
          show-password-on="mousedown"
          v-model:value="formValue.password"
          :placeholder="$t('forget.请输入新密码')"
        />
      </n-form-item>
      <n-form-item ref="rPasswordFormItemRef" first path="reenteredPassword">
        <n-input
          v-model:value="formValue.reenteredPassword"
          class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
          :disabled="!formValue.password"
          show-password-on="mousedown"
          type="password"
          :placeholder="$t('forget.请再次输入新密码')"
          @keydown.enter.prevent
        />
      </n-form-item>
      <div>
        <n-checkbox v-model:checked="formValue.checkboxValue">
          {{ $t("forget.要求所有设备使用新密码登录") }}
        </n-checkbox>
      </div>
      <n-form-item>
        <n-button
          @click="handleValidateClick"
          type="primary"
          round
          class="w-full h-50px px-4 bg-[#219653]"
        >
          <span class="text-18px text-[#FDFDFD]">
            {{ $t("forget.重置密码") }}
          </span>
        </n-button>
      </n-form-item>
    </n-form>
  </n-card>
</template>

<script setup>
import { NCard, NForm, NFormItem, NInput, NButton, NCheckbox } from "naive-ui";
import { useI18n } from "vue-i18n";
const { t } = useI18n();
const emit = defineEmits(["setUpShowIndex"]);
const props = defineProps({
  formData: Object,
});

const formRef = ref(null);
const rPasswordFormItemRef = ref(null);
// const message = useMessage();
/**
 * 表单数据
 */
const formValue = ref({
  password: "",
  reenteredPassword: "",
});
const validatePasswordStartWith = (rule, value) => {
  return (
    !!formValue.value.password &&
    formValue.value.password.startsWith(value) &&
    formValue.value.password.length >= value.length
  );
};
const validatePasswordSame = (rule, value) => {
  return value === formValue.value.password;
};
const validatePasswordLength = (rule, value) => {
  return value.length >= 8;
};
/**
 * 登录规则
 */
const rules = {
  password: [
    {
      required: true,
      message: t("forget.请输入新密码"),
      trigger: "input",
    },
    {
      validator: validatePasswordLength,
      message: t("forget.密码长度不足8位"),
      trigger: "blur",
    },
  ],
  reenteredPassword: [
    {
      required: true,
      message: t("forget.再次输入新密码"),
      trigger: ["input", "blur"],
    },
    {
      validator: validatePasswordStartWith,
      message: t("forget.两次密码不相同"),
      trigger: "input",
    },
    {
      validator: validatePasswordSame,
      message: t("forget.两次密码不相同"),
      trigger: ["blur", "password-input"],
    },
  ],
};
/**
 * 点击继续
 * @param {*} e
 */
const handleValidateClick = (e) => {
  e.preventDefault();
  formRef.value?.validate((errors) => {
    if (!errors) {
      console.log("这里调接口");
      // emit("setUpShowIndex", 1, formValue.value);
    }
  });
};
</script>

<style lang="scss" scoped>
.n-card {
  width: 600px;
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
<style lang="scss" scoped>
.registerTitle {
  font-size: 28px;
  font-family: Inter-Medium, Inter;
  font-weight: 500;
  color: #020300;
  -webkit-background-clip: text;
}
.smallTitle {
  font-size: 20px;
  font-family: Inter-Medium, Inter;
  font-weight: 500;
  color: #333333;
  -webkit-background-clip: text;
  margin-bottom: 10px;
}
@media (max-width: 1020px) {
  .smallTitle {
    width: 70%;
  }
}
@media (max-width: 750px) {
  .smallTitle {
    width: 90%;
  }
}
</style>
