<template>
  <div class="mb-20px">
    <div class="flex items-center py-5px mb-10px pl-15px">
      <img :src="ARectangle" alt="" width="30" />
      <div class="ml-15px text-24px font-bold">
        {{ $t('home.设置') }}
      </div>
    </div>
  </div>
  <div v-for="(item, index) in menuList" :key="item.id">
    <div
      class="flex items-center py-10px mb-10px px-15px cursor-pointer menuItem"
      :class="index === activeIndex ? 'menuActive' : ''"
      @click="handleMenuClick(index)"
    >
      <img :src="index === activeIndex ? item.activeIcon:item.icon" alt="" class="w-28px h-28px flex-shrink-0" />
      <div class="ml-15px">
        {{ item.name }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { NCard, NForm, NFormItem, NInput, NButton } from 'naive-ui'
import { useI18n } from 'vue-i18n'
import ARectangle from '@/assets/img/home/ARectangle.png'
import Suggestedforyou from '@/assets/img/home/Suggestedforyou.png'
import SuggestedforyouActive from '@/assets/img/home/SuggestedforyouActive.png'
import Edit from '@/assets/img/home/Edit.png'
import EditActive from '@/assets/img/home/EditActive.png'
import Analytics from '@/assets/img/home/Analytics.png'
import AnalyticsActive from '@/assets/img/home/AnalyticsActive.png'
import Resources from '@/assets/img/home/Resources.png'
import ResourcesActive from '@/assets/img/home/ResourcesActive.png'
import Personalstrength from '@/assets/img/home/Personalstrength.png'
import PersonalstrengthActive from '@/assets/img/home/PersonalstrengthActive.png'
import Jobexperience from '@/assets/img/home/Jobexperience.png'
import JobexperienceActive from '@/assets/img/home/JobexperienceActive.png'
import Projectexperience from '@/assets/img/home/Projectexperience.png'
import ProjectexperienceActive from '@/assets/img/home/ProjectexperienceActive.png'
import Education from '@/assets/img/home/Education.png'
import EducationActive from '@/assets/img/home/EducationActive.png'
import Skills from '@/assets/img/home/Skills.png'
import SkillsActive from '@/assets/img/home/SkillsActive.png'
import Interests from '@/assets/img/home/Interests.png'
import InterestsActive from '@/assets/img/home/InterestsActive.png'
import Dreamposition from '@/assets/img/home/Dreamposition.png'
import DreampositionActive from '@/assets/img/home/DreampositionActive.png'
import Customitem from '@/assets/img/home/Customitem.png'
import CustomitemActive from '@/assets/img/home/CustomitemActive.png'
const { t } = useI18n()
const emit = defineEmits(['setAnchorpoint'])
const activeIndex = ref(0)
const language = useCookie('lang')
const menuList = ref([])
// 监听切换语言
watch(
  () => {
    if (language) {
      menuList.value = [
        {
          name: t('home.帐户首选项'),
          id: 1,
          icon: Suggestedforyou,
          activeIcon:SuggestedforyouActive,
        },
        {
          name: t('home.建议您使用'),
          id: 2,
          icon: Edit,
          activeIcon:EditActive,

        },
        {
          name: t('home.分析'),
          id: 3,
          icon: Analytics,
          activeIcon:AnalyticsActive,

        },
        {
          name: t('home.资源'),
          id: 4,
          icon: Resources,
          activeIcon:ResourcesActive,

        },
        {
          name: t('home.个人力量'),
          id: 5,
          icon: Personalstrength,
          activeIcon:PersonalstrengthActive,

        },
        {
          name: t('home.工作经历'),
          id: 6,
          icon: Jobexperience,
          activeIcon:JobexperienceActive,

        },
        {
          name: t('home.项目经验'),
          id: 7,
          icon: Projectexperience,
          activeIcon:ProjectexperienceActive,

        },
        {
          name: t('home.教育'),
          id: 8,
          icon: Education,
          activeIcon:EducationActive,

        },
        {
          name: t('home.技能'),
          id: 9,
          icon: Skills,
          activeIcon:SkillsActive,

        },
        {
          name: t('home.权益'),
          id: 10,
          icon: Interests,
          activeIcon:InterestsActive,

        },
        {
          name: t('home.梦想的位置'),
          id: 11,
          icon: Dreamposition,
          activeIcon:DreampositionActive,

        },
        {
          name: t('home.自定义项目'),
          id: 12,
          icon: Customitem,
          activeIcon:CustomitemActive,

        },
      ]
    }
  },

  { immediate: true, deep: true },
)

/**
 * 点击切换菜单
 * @param {*} e
 */
const handleMenuClick = (e) => {
  console.log(e)
  if (activeIndex.value === e) return
  activeIndex.value = e
  emit('setAnchorpoint', e)
}
</script>

<style lang="scss" scoped>
.menuItem {
  position: relative;
}

.menuActive {
  background: rgba(51, 95, 62, 0.1);
  color: #219653;
  border-radius: 0 3px 3px 0;
}
.menuActive::after {
  position: absolute;
  content: '';
  width: 4px;
  height: 100%;
  left: -2px;
  top: 0;
  background: #335f3e;
  border-radius: 0px 0px 0px 0px;
}
</style>
