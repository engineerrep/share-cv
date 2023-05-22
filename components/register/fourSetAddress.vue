<template>
  <div class="fourSetAddress">
    <n-form ref="formRef" :label-width="80" :model="formValue" :rules="rules">
      <n-form-item :label="$t('register.country region')" path="Region" class="region-input">
        <!-- <n-input
          class="h-44px rounded-6px border-[#696969] text-18px text-[#4B4B4B] leading-44px"
          v-model:value="formValue.Region"
          :placeholder="$t('login.请输入国家/地区')"
        /> -->
        <n-cascader
            v-model:value="formValue.Region"
            :placeholder="$t('register.Please enter a country')"
            expand-trigger="click"
            :options="optionsRegion"
            class="h-44px rounded-6px border-[#616161] text-18px text-[#4B4B4B] leading-44px"
            @update:value="handleUpdateValue"
        />
      </n-form-item>
      <n-form-item :label="$t('register.Address')" path="area" class="area-input">
        <n-input
            class="h-44px rounded-6px border-[#616161] text-18px text-[#4B4B4B] leading-44px"
            v-model:value="formValue.area"
            :placeholder="$t('register.Please enter detailed address')"
        />
      </n-form-item>
      <n-form-item>
        <CVButton @click="handleValidateClick">
          <template #main>
              <span class="text-20px text-[#FDFDFD]">
              Next
              </span>
          </template>
        </CVButton>
      </n-form-item>
    </n-form>
  </div>
</template>

<script setup>
import {NButton, NCascader, NForm, NFormItem, NInput} from "naive-ui";
import {useI18n} from "vue-i18n";
import {apiCommonCountry, apiUpdateCity, apiUpdateCountry} from "~/api/login";
import {reviseDom} from "~/utils/util";
import {useRouter} from "vue-router";
import CVButton from '@/components/form/button.vue';

let router = useRouter()
const {t, locale} = useI18n();
const props = defineProps({
  formData: Object,
});
const emit = defineEmits(["setUpShowIndex"]);
// 国家列表
const optionsRegion = ref([]);
const formRef = ref(null);
const loading = ref(false);

apiCommonCountry().then(res => {
  // console.log('获取国家列表', res)
  let data = []
  if (res.data.list) {
    res.data.list.forEach(item => {
      data.push(
          {
            value: item.id, label: item.name,
            code: item.code, num: item.num
          }
      )
    });
    optionsRegion.value = data
  }
})

// const message = useMessage();
/**
 * 表单数据
 */
const formValue = ref({
  Region: null,
  area: "",
});

// 改变验证提示
watch(locale, (newVal, oldVal) => {
  reviseDom('.region-input .n-form-item-feedback__line', t('register.Please enter a country'))
  reviseDom('.area-input .n-form-item-feedback__line', t('register.Please enter detailed address'))
})
/**
 * 登录规则
 */
const rules = {
  Region: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("register.Please enter a country"));
      }
      return true;
    },
  },
  area: {
    required: true,
    trigger: ["input", "blur"],
    validator(rule, value) {
      if (!value) {
        return new Error(t("register.Please enter detailed address"));
      }
      return true;
    },
  },
};
/**
 * 设置国家-选择框事件
 * @param {*} e
 */
const handleUpdateValue = (value, option) => {
  apiUpdateCountry({countryId: option.value, countryName: option.label}).then(res => {
    console.log(res)
    if (res.data && res.data.status) {
      //   设置国家成功
    }
  })
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
      let params = {
        city: formValue.value.area,
      }
      apiUpdateCity(params).then(res => {
        if (res.data.status && res.data) {
          router.push('/personal-wditor')
        }
      })
    }
  });
};
</script>

<style lang="scss">
.fourSetAddress {
  .n-base-selection,
  .n-base-selection-label {
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

.cv-button {
  width: 100%;
  height: 69px;
  border-radius: 30px;
}

@media (max-width: 400px) {
  .n-input {
    font-size: 14px;
  }
}

@media (max-width: 480px) {
  .fourSetAddress {
    width: 90% !important;
    margin: 0 auto;
    border-radius: 10px;
  }
}
</style>
