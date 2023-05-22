<template>
  <n-card title="">
    <div class="registerTitle mb-1">
      {{ $t("forget.我们向您的电子邮件发送了代码") }}
    </div>
    <div class="smallTitle">
      {{ $t("forget.输入发送至的6位验证码") }}j*****@gmail.com.<span
        class="color-219653"
        >{{ $t("forget.改变") }}</span
      >
    </div>
    <n-form ref="formRef" :label-width="80" :model="formValue" :rules="rules">
      <n-form-item path="Code">
        <div class="w-full">
          <n-input
            class="h-44px w-full rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
            v-model:value="formValue.Code"
            :placeholder="$t('forget.请输入验证码')"
          />
          <div
            class="text-right mt-2 cursor-pointer"
            @click="obtain"
            :class="countDown.isGrey ? '' : 'color-219653'"
          >
            {{ $t("forget.重新发送代码") }}
            <span v-if="!countDown.show">{{ countDown.count }}s</span>
          </div>
        </div>
      </n-form-item>
      <n-form-item>
        <n-button
          @click="handleValidateClick"
          type="primary"
          round
          class="w-full h-50px px-4 bg-[#219653]"
        >
          <span class="text-18px text-[#FDFDFD]">
            {{ $t("forget.提交") }}
          </span>
        </n-button>
      </n-form-item>
      <div class="cursor-pointer w-full mb-3">
        {{
          $t("forget.如果在收件箱中没有看到电子邮件，请检查垃圾邮件文件夹。")
        }}
      </div>
      <div class="cursor-pointer w-full mb-3 color-219653">
        {{ $t("forget.无法访问此电子邮件？") }}
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
const countDown = reactive({
  dis: false,
  show: true,
  isGrey: false, //按钮样式
  timer: null, //设置计时器,
  count: "",
});
//点击获取验证码：

const obtain = () => {
  if (countDown.dis) return;
  if (!countDown.timer) {
    let mins = 60
    countDown.count = mins;
    countDown.isGrey = true;
    countDown.show = false;
    countDown.dis = true;
    countDown.timer = setInterval(() => {
      if (countDown.count > 0 && countDown.count <= mins) {
        countDown.count--;
      } else {
        countDown.dis = false;
        countDown.isGrey = false;
        countDown.show = true;
        clearInterval(countDown.timer);
        countDown.timer = null;
      }
    }, 1000);
  }
};
const formRef = ref(null);
// const message = useMessage();
/**
 * 表单数据
 */
const formValue = ref({
  Code: "",
});
/**
 * 登录规则
 */
const rules = {
  Code: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("forget.请输入验证码"));
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
      emit("setUpShowIndex", 2, formValue.value);
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
