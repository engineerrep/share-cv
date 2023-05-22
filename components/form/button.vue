<template>
  <button class="cv-button" type="button">
    <slot name="main"></slot>
  </button>
</template>

<script setup>
let props = defineProps({
  round: Boolean,
  type: String,
})
//  2023-4-18按钮都有圆角
let borderRadius = ref('20px')
// 主体颜色
let themeColor = useState("_themeColor")
// 主体颜色二
let themeColorTwo = useState("_themeColorTwo")
</script>

<style lang="scss" scoped>
.cv-button {
  border-radius: v-bind(borderRadius);
  background-color: v-bind(themeColor);
  border: 0;
  font-size: 20px;
  color: #FDFDFD;

  &:hover {
    background-color: v-bind(themeColorTwo);
  }

  &:active {
    background-color: #0b1f33;
  }
}
</style>