<template>
  <div class="fiveSetJob">
    <n-form ref="formRef" :label-width="80" :model="formValue" :rules="rules">
      <template v-if="!isStudent">
        <n-form-item :label="$t('login.最近的职务')" path="job">
          <n-input
              class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
              v-model:value="formValue.job"
              :placeholder="$t('login.请输入最近的职务')"
          />
        </n-form-item>
        <n-form-item :label="$t('login.就业类型')" path="employmentType">
          <n-cascader
              v-model:value="formValue.employmentType"
              :placeholder="$t('login.请输入就业类型')"
              expand-trigger="click"
              :options="options"
              class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
              @update:value="handleUpdateValue"
          />
        </n-form-item>
        <n-form-item :label="$t('login.最近的公司')" path="company">
          <n-input
              class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
              v-model:value="formValue.company"
              :placeholder="$t('login.请输入最近的公司')"
          />
        </n-form-item>
      </template>
      <template v-else>
        <n-form-item :label="$t('login.大学/学校')" path="school">
          <n-input
              class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
              v-model:value="formValue.school"
              :placeholder="$t('login.请输入大学/学校')"
          />
        </n-form-item>
        <n-form-item :label="$t('login.开始时间')" path="startTime">
          <n-date-picker
              class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px w-full"
              v-model:value="formValue.startTime"
              type="date"
              :placeholder="$t('login.请输入开始时间')"
          />
        </n-form-item>
        <n-form-item :label="$t('login.结束时间')" path="endTime">
          <n-date-picker
              class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px w-full"
              v-model:value="formValue.endTime"
              type="date"
              :placeholder="$t('login.请输入结束时间')"
          />
        </n-form-item>
      </template>

      <div
          class="registerTitle mb-6 text-center cursor-pointer"
          @click="handleChange"
      >
        {{ isStudent ? $t("login.我不是学生") : $t("login.我是学生") }}
      </div>
      <n-form-item>
        <n-button
            @click="handleValidateClick"
            type="primary"
            :disabled="
            !formValue.job || !formValue.employmentType || !formValue.company
          "
            round
            class="w-full h-50px px-4 bg-[#219653]"
        >
          <span class="text-18px text-[#FDFDFD]">
            {{ $t("login.确定") }}
          </span>
        </n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<script setup>
import {
  NForm,
  NFormItem,
  NInput,
  NButton,
  NCascader,
  NDatePicker,
} from "naive-ui";
import {useI18n} from "vue-i18n";

const {t} = useI18n();
const props = defineProps({
  formData: Object,
});
const emit = defineEmits(["setUpShowIndex"]);
// 国家列表
const options = ref([
  {value: 1, label: "中国", children: [{value: 3, label: "台湾"}]},
  {value: 2, label: "美国"},
]);
const formRef = ref(null);
const loading = ref(false);
// 是否是学生
const isStudent = ref(false);
/**
 * 表单数据
 */
const formValue = ref({
  employmentType: null,
  job: "",
  company: "",
  school: "",
  startTime: null,
  endTime: null,
});
/**
 * 登录规则
 */
const rules = {
  job: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("login.请输入最近的职务"));
      }
      return true;
    },
  },
  employmentType: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("login.请输入就业类型"));
      }
      return true;
    },
  },
  company: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("login.请输入最近的公司"));
      }
      return true;
    },
  },
  school: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("login.请输入大学/学校"));
      }
      return true;
    },
  },
  startTime: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("login.请输入开始时间"));
      }
      return true;
    },
  },
  endTime: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("login.请输入结束时间"));
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
 * 点击我的学生
 */
const handleChange = () => {
  isStudent.value = !isStudent.value;
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
      emit("setUpShowIndex", 5);
      setTimeout(() => {
        loading.value = false;
      }, 1500);
    }
  });
};
</script>
<style lang="scss">
.fiveSetJob {
  .n-base-selection,
  .n-base-selection-label,
  .n-input--stateful {
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
.registerTitle {
  font-size: 28px;
  font-family: Inter-Medium, Inter;
  font-weight: 500;
  color: #020300;
  -webkit-background-clip: text;
}


.fiveSetJob {
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
  .fiveSetJob {
    width: 90% !important;
    margin: 0 auto;
    border-radius: 10px;
  }
}
</style>
