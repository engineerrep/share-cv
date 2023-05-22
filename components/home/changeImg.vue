<template>
  <div style="width: 100%; height: 500px">
    <vue-cropper
      ref="cropper"
      :img="option.img"
      :outputSize="option.size"
      :outputType="option.outputType"
      :centerBox="option.centerBox"
      :autoCrop="option.autoCrop"
      :autoCropWidth="option.autoCropWidth"
      :autoCropHeight="option.autoCropHeight"
      :fixedBox="option.fixedBox"
      :canMove="option.canMove"
      :original="option.original"
      :high="option.high"
      :max="option.max"
    />
  </div>
  <div @click="handleClick">111</div>
</template>

<script setup>
import { VueCropper } from "vue-cropper";
import { NCard, NForm, NFormItem, NInput, NButton } from "naive-ui";
import { useI18n } from "vue-i18n";
const { t } = useI18n();
const emit = defineEmits(["setAnchorpoint"]);
const props = defineProps({
  formData: Object,
});
const option = {
  img: "https://avatars2.githubusercontent.com/u/15681693?s=460&v=4",
  size: 1,
  full: false,
  outputType: "png",
  canMove: true,
  fixedBox: false,
  original: false,
  canMoveBox: true,
  autoCrop: true,
  // 只有自动截图开启 宽度高度才生效
  autoCropWidth: 160,
  autoCropHeight: 150,
  centerBox: false,
  high: true,
  max: 99999,
};
const cropper = ref(null);
/**
 * 点击继续
 * @param {*} e
 */
const handleClick = () => {
  // emit("setAnchorpoint", 1);
  cropper.value.getCropData((data) => {
    console.log(data);
  });
};
</script>

<style lang="scss" scoped></style>
