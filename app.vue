<template>
  <!-- naive-ui 默认情况下使用 inline style 作为主题变量的载体，因此每个组件上都会挂载许多 inline CSS。如果你需要 SSR，或者想让开发者工具看起来更干净，可以打开 inline-theme-disabled 属性。 -->
  <NConfigProvider inline-theme-disabled :theme-overrides="themeOverrides">
    <NuxtLayout>
      <NuxtPage/>
    </NuxtLayout>
  </NConfigProvider>

</template>

<script setup>
import {NConfigProvider} from "naive-ui";
import {useI18n} from "vue-i18n";
import {useRouter} from "vue-router";

const router = useRouter();

const {t, locale} = useI18n();

// console.log('routeroute')

/**
 * js 文件下使用这个做类型提示
 * @type import('naive-ui').GlobalThemeOverrides
 */
const themeOverrides = {
    common: {
        primaryColor: "#219653",
        primaryColorHover: "#219653",
        primaryColorPressed: "#219653",
        primaryColorSuppl: "#FFFFFF",
        borderRadius: "4px",
        fontSize: "20px",
    },
};

onMounted(() => {
  let viewport = {}
  if (/iPad|iPhone|iPod/.test(navigator.userAgent)) {
    viewport = {
      name: 'viewport',
      content: 'width=device-width,initial-scale=1, maximum-scale=1,user-scalable=no'
    }
  } else if (/Android/.test(navigator.userAgent)) {
    viewport = {
      name: 'viewport',
      content: 'width=device-width,initial-scale=1, maximum-scale=1,user-scalable=no'
    }
  } else {
    viewport = {
      name: 'viewport',
      content: 'width=device-width, initial-scale=1'
    }
  }
// 设置外部script
  useHead({
    title: "Share-CV",
    script: [
      {
        src: "https://www.google.com/recaptcha/enterprise.js?render=6Lff22ckAAAAAA8xHrmY_BQ_YtTWyp39k-jJuxEv",
        body: true,
        async: true,
        defer: true
      }
    ],
    meta: [
      {name: t("seo.app.description"), content: t("seo.app.description")},
      {name: t("seo.app.keywords"), content: t("seo.app.keywords")},
      {name: 'referrer', content: 'no-referrer'},
      viewport
    ],


  });

})
</script>

<style lang="scss">
@import "assets/styles/community.scss";

</style>
