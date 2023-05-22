<template>
  <div class=" footer plate-ten">
    <div>
      <img src="@/assets/img/logo_min.png" alt="" class="log-bottom"/>
      <!--      <span>© 2023</span>-->
    </div>
    <div class="tail-right df">
      <!--      <div class="mr-65px">Share-CV Corporation © 2023</div>-->
      <p v-for="item in tailItem"
         :key="item"
         class="tail-item cursor-pointer"
         @click="router.push(`/legal/${item.href}`)"
      >
        {{ item.txt }}
      </p>
    </div>
  </div>
</template>

<script setup>
import {useRouter} from "vue-router";

const router = useRouter();
onMounted(() => {
  window.addEventListener('scroll', onScroll, true)
})
let plateTenLeft = ref()
const onScroll = () => {
  plateTenLeft.value = document.documentElement.scrollLeft + 'px'
}
// 尾部
const tailItem = [
  {txt: "User Agreement", href: "user-agreement"},
  {txt: "Privacy Policy", href: "privacy-policy"},
  {txt: "Cookie Policy", href: "cookie-policy"},
]

</script>

<style lang="scss" scoped>
@import "@/assets/styles/head-tail.scss";

.plate-ten {
  position: relative;
  left: v-bind(plateTenLeft);
}
</style>
