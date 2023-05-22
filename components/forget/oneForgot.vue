<template>
  <n-card title="">
    <div class="registerTitle mb-1">
      {{ $t("forget.忘记密码？") }}
    </div>
    <div class="smallTitle">
      {{ $t("forget.通过两个快速步骤重置密码") }}
    </div>
    <n-form ref="formRef" :label-width="80" :model="formValue" :rules="rules">
      <n-form-item path="Email">
        <n-input
          class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
          v-model:value="formValue.Email"
          :placeholder="$t('forget.请输入邮箱')"
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
            {{ $t("forget.重置密码") }}
          </span>
        </n-button>
      </n-form-item>
      <div class="text-center cursor-pointer w-full mb-3">
        {{ $t("forget.返回") }}
      </div>
    </n-form>
  </n-card>
</template>

<script setup>
import { NCard, NForm, NFormItem, NInput, NButton } from "naive-ui";
import { useI18n } from "vue-i18n";
const { t } = useI18n();
const emit = defineEmits(["setUpShowIndex"]);
const props = defineProps({
  formData: Object,
});

const formRef = ref(null);
// const message = useMessage();
/**
 * 表单数据
 */
const formValue = ref({
  Email: "",
});
/**
 * 登录规则
 */
const rules = {
  Email: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("forget.请输入邮箱"));
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
      console.log("这里调接口");
      emit("setUpShowIndex", 1, formValue.value);
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
