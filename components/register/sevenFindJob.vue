<template>
  <div class="fourSetAddress">
    <n-button
      @click="handleValidateClick"
      type="primary"
      ghost
      round
      class="w-full h-50px px-4 mb-30px"
    >
      <span class="text-18px text-[#219653]">
        {{ $t("login.是的") }}
      </span>
    </n-button>
    <n-button
      @click="handleValidateClick"
      type="primary"
      ghost
      round
      class="w-full h-50px px-4"
    >
      <span class="text-18px text-[#219653]">
        {{ $t("login.不是现在") }}
      </span>
    </n-button>
  </div>
</template>

<script setup>
import { NButton } from "naive-ui";
import { useI18n } from "vue-i18n";
const { t } = useI18n();
const props = defineProps({
  formData: Object,
});
const emit = defineEmits(["setUpShowIndex"]);
/**
 * 点击继续
 * @param {*} e
 */
const handleValidateClick = (e) => {
  emit("setUpShowIndex", 7);
};
</script>
<style lang="scss" scoped>
.fourSetAddress {
  width: 600px !important;
  padding-top: 30px;
  margin: 0 auto;
  border-radius: 10px;
}

@media (max-width: 480px) {
  .fourSetAddress {
    width: 90% !important;
    margin: 0 auto;
    border-radius: 10px;
  }
}
</style>
