<template>
  <div class="my-10">
    <OneForgot
      v-if="shwoIndex === 0"
      @setUpShowIndex="setUpShowIndex"
    ></OneForgot>
    <TwoSenCode
      v-if="shwoIndex === 1"
      @setUpShowIndex="setUpShowIndex"
      :formData="formData"
    ></TwoSenCode>
    <TreSetpassword
      v-if="shwoIndex === 2"
      @setUpShowIndex="setUpShowIndex"
      :formData="formData"
    ></TreSetpassword>
  </div>
</template>

<script setup>
import OneForgot from "@/components/forget/oneForgot.vue";
import TwoSenCode from "@/components/forget/twoSenCode.vue";
import TreSetpassword from "@/components/forget/treSetpassword.vue";

import { useI18n } from "vue-i18n";
const { t } = useI18n();
// 基础数据
const formData = ref({});
// 步骤下标
const shwoIndex = ref(2);
const setUpShowIndex = (index, data) => {
  shwoIndex.value = index;
  formData.value = { ...data, ...formData.value };
  console.log(formData.value);
};

useHead({
  title: t("seo.home.title"),
  meta: {
    keywords: t("seo.home.keywords"),
    description: t("seo.home.description"),
  },
});
</script>

<style lang="scss" scoped></style>
