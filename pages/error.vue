<template>
  <div class="login_wrap">
    <NConfigProvider inline-theme-disabled :theme-overrides="themeOverrides">
      <n-result status="404" :title="t('layout.出错啦')" description="">
        <template #footer>
          <nuxt-link to="/">
            <n-button>{{ $t("layout.返回首页") }}</n-button>
          </nuxt-link>
        </template>
      </n-result>
    </NConfigProvider>
  </div>
</template>

<script setup>
import { NConfigProvider, NResult, NButton } from 'naive-ui';
import { useI18n } from "vue-i18n"
const { t, locale } = useI18n()

const themeOverrides = {
  common: {
    primaryColor: "#EC028C",
    primaryColorHover: "#EC028C",
    primaryColorPressed: "#EC028C",
    primaryColorSuppl: "#EC028C",
    borderRadius: "4px",
    fontSize: "16px",
  },
}

</script>

<style lang="scss" scoped>

</style>