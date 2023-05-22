<template>
  <div class="body" id="app">
    <Header/>
    <main class="main_container">
      <slot/>
    </main>
    <Footer/>
  </div>
</template>

<script setup>
import {isMobile} from '@/utils/util';

const route = useRoute();

let pcSide = useState("_pcSide", () => true)
let isGrecaptchaBadge = ref('block')
onMounted(() => {
  // 谷歌验证图标只是在登录和注册页面显示，其他页面不显示
  if (document.body.clientWidth >= 420) {
    pcSide.value = true
  } else {
    pcSide.value = false
  }
})

// 主体颜色
let themeColor = useState("_themeColor", () => '#005bb6')
// 主体颜色二
let themeColorTwo = useState("_themeColorTwo", () => '#003a74')
</script>

<style lang="scss" scoped>
.body {
  // background: #f4f2f0;
  @apply flex flex-col min-h-screen;
}

.main_container {
  min-height: calc(100vh - 188px);
}
</style>
