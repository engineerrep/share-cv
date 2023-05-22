<template>
  <div class="fourSetAddress">
    <n-form ref="formRef" :label-width="80" :model="formValue" :rules="rules">
      <n-form-item :label="$t('login.职务名称')" path="Job">
        <n-select
          v-model:value="formValue.Job"
          multiple
          @update:value="handleUpdateValue"
          :options="optionsJob"
          class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
        />
      </n-form-item>
      <n-form-item :label="$t('login.工作地点')" path="area">
        <n-select
          v-model:value="formValue.area"
          multiple
          @update:value="handleUpdateValue"
          :options="optionsJob"
          class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
        />
      </n-form-item>
      <n-form-item>
        <n-button
          @click="handleValidateClick"
          type="primary"
          :disabled="formValue.area.Job === 0 || formValue.area.length === 0"
          round
          class="w-full h-50px px-4 bg-[#219653]"
        >
          <span class="text-18px text-[#FDFDFD]">
            {{ $t("login.继续") }}
          </span>
        </n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<script setup>
import { NForm, NFormItem, NButton, NSelect } from "naive-ui";
import { useI18n } from "vue-i18n";
const { t } = useI18n();
const props = defineProps({
  formData: Object,
});
const emit = defineEmits(["setUpShowIndex"]);
// 国家列表
const optionsJob = ref([
  { value: 1, label: "中国" },
  { value: 2, label: "美国" },
]);
const formRef = ref(null);
const loading = ref(false);
// const message = useMessage();
/**
 * 表单数据
 */
const formValue = ref({
  Job: [],
  area: [],
});
/**
 * 登录规则
 */
const rules = {
  Job: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (value.length === 0) {
        return new Error(t("login.请输入职务名称"));
      }
      return true;
    },
  },
  area: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (value.length === 0) {
        return new Error(t("login.请输入工作地点"));
      }
      return true;
    },
  },
};
/**
 * 选择框事件
 * @param {*} e
 */
const handleUpdateValue = (value, option) => {
  console.log(value, option);
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
      // emit("setUpShowIndex", 8);
      setTimeout(() => {
        loading.value = false;
      }, 1500);
    }
  });
};
</script>
<style lang="scss">
.fourSetAddress {
  .n-base-selection,
  .n-base-selection-label,
  .n-base-selection-tags {
    height: 44px !important;
    line-height: 44px !important;
    color: rgba(75, 75, 75, 1) !important;
    font-size: 18px !important;
    border-radius: 6px !important;
    border-color: rgba(105, 105, 105, 1) !important;
  }
}
</style>
<style lang="scss" scoped>
.fourSetAddress {
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
  .fourSetAddress {
    width: 90% !important;
    margin: 0 auto;
    border-radius: 10px;
  }
}
</style>
