<template>
  <div class="text-[#333333]">
    <!-- 用户设置 -->
    <div class="mb-20px bg-white rounded-10px border border-[#E2E0DF]" id="0">
      <div class="relative">
        <img
          src="@/assets/img/home/camera.png"
          class="w-54px h-54px absolute top-27px right-25px cursor-pointer"
          alt=""
        />
        <img
          src="@/assets/img/home/userBack.png"
          class="w-full h-261px"
          alt=""
        />
        <div
          class="w-214px h-214px bg-[#E2E2E2] rounded-105px flex items-center justify-center absolute bottom-[-63px] left-31px cursor-pointer"
        >
          <img
            src="@/assets/img/home/userFrame.png"
            class="w-67px h-67px"
            alt=""
          />
        </div>
      </div>
      <div class="p-32px">
        <div class="flex flex-row-reverse mb-23px">
          <img
            src="@/assets/img/home/userEdit.png"
            class="w-31px h-31px cursor-pointer"
            alt=""
          />
        </div>
        <div class="text-30px text-[#000000] mb-4px">{{ 'John Smith' }}</div>
        <div class="text-21px font-400 text-[#000000] mb-6px">
          {{ 'UX Analyst at Mobbin' }}
        </div>
        <div class="text-18px font-400 text-[#828282] mb-30px">
          <span>{{ 'Menlo Park, California,United States·Contact info' }}</span>
          <span class="text-[#219653]">{{ 'UX Analyst at Mobbin' }}</span>
        </div>
        <div class="flex items-center mb-31px">
          <span
            class="truncate px-19px leading-43px bg-[#219653] rounded-21.5px h-43px text-[#ffffff] text-22px mr-10px cursor-pointer"
          >
            {{ $t('home.打开至') }}
          </span>
          <span
            class="truncate px-19px leading-43px rounded-21.5px border border-[#219653] h-43px text-[#219653] text-22px mr-10px cursor-pointer"
          >
            {{ $t('home.添加配置文件部分') }}
          </span>
          <span
            class="truncate px-19px leading-43px rounded-21.5px border border-[#535353] h-43px text-[#000000] font-bold text-22px mr-10px cursor-pointer"
          >
            {{ $t('home.更多') }}
          </span>
          <span
            class="truncate px-19px leading-43px bg-[#219653] rounded-21.5px h-43px text-[#ffffff] text-22px cursor-pointer"
          >
            {{ $t('home.开放工作') }}
          </span>
        </div>

        <div class="flex items-center">
          <div
            class="w-auto h-104px bg-[#EAE6E0] rounded-8px p-16px box-border mr-10px relative"
          >
            <img
              src="@/assets/img/home/textEdit.png"
              class="w-28px h-28px absolute top-12px right-14px cursor-pointer"
              alt=""
            />
            <div class="font-500 text-[#000000] text-20px mb-5px">
              {{ $t('home.开放工作') }}
            </div>
            <div class="w-90% text-18px font-400 truncate mb-3px">
              User Experience Analy st and Product Designe...
            </div>
            <div class="text-18px text-[#219653]">
              {{ $t('home.查看所有详细信息') }}
            </div>
          </div>
          <div
            class="w-auto h-104px border boder-[#ECECEC] flex-col justify-between rounded-8px ml-10px p-16px box-border relative"
          >
            <div class="font-500 text-[#000000] text-20px mb-25px">
              {{ $t('home.共享您的个人资料') }}
            </div>
            <div class="text-18px text-[#219653] cursor-pointer">{{ $t('home.开始') }}</div>
          </div>
        </div>
      </div>
    </div>
    <!-- 建议您使用 -->
    <div
      class="mb-20px bg-white rounded-8px border border-[#E2E0DF] px-31px pt-33px pb-32px"
      id="1"
    >
      <div class="text-[#000000] text-26px leading-30px mb-7px">
        {{ $t('home.建议您使用') }}
      </div>
      <div class="flex items-center mb-28px">
        <img
          src="@/assets/img/home/see.png"
          class="w-23px h-23px cursor-pointer flex-shrink-0 mr-4px"
          alt=""
        />
        <span class="text-[#828282] text-18px font-400">
          {{ $t('home.对你来说是私人的') }}
        </span>
      </div>
      <div class="text-[#000000] text-21px font-500 mb-5px">
        {{ $t('home.中间的') }}
      </div>
      <div class="flex items-center justify-between mb-16px">
        <div
          class="flex-1 relative bg-[#DBDBDB] h-12px rounded-6px overflow-hidden"
        >
          <div
            class="h-12 bg-[#333333] absolute top-0 left-0 z-10"
            :style="{
              width: Number(nowNumber / allNumber).toFixed(2) * 100 + '%',
            }"
          ></div>
        </div>
        <div class="flex-shrink-0 ml-5px text-[#060606]">
          {{ nowNumber }}/{{ allNumber }}
        </div>
      </div>
      <div class="text-19px mb-36px">
        <span>{{ $t('home.完成2个步骤以实现') }}</span>
        <span class="text-[#219653] cursor-pointer">{{ $t('home.全明星') }}</span>
      </div>
      <div class="flex items-center mb-10px">
        <div
          class="flex-1 border border-[#ECECEC] rounded-10px py-17px px-20px pr-32px box-border h-226px mr-10px"
        >
          <div class="flex items-center">
            <img
              src="@/assets/img/home/school.png"
              class="w-62px h-62px flex-shrink-0 mr-9px"
              alt=""
            />
            <div class="text-[#000000] text-20px">
              {{ $t('home.你上过哪所大学或学校？') }}
            </div>
          </div>
          <div class="text-[#828282] text-17px leading-28px mb-10px">
            {{ $t('home.包括学校在内的会员可获得') }}
          </div>
          <div
            class="h-43px border text-center border-[#535353] rounded-22px px-19px text-22px leading-43px cursor-pointer"
          >
            {{ $t('home.添加教育') }}
          </div>
        </div>

        <div
          class="flex-1 border border-[#ECECEC] rounded-10px py-17px px-20px pr-32px box-border h-226px ml-10px"
        >
          <div class="flex items-center">
            <img
              src="@/assets/img/home/idCard.png"
              class="w-62px h-62px flex-shrink-0 mr-9px"
              alt=""
            />
            <div class="text-[#000000] text-20px">
              {{ $t('home.写一篇总结，突出你的个性或工作经验') }}
            </div>
          </div>
          <div class="text-[#828282] text-17px leading-28px mb-10px">
            {{ $t('home.包含摘要的成员可获得多达') }}
          </div>
          <div
            class="h-43px border text-center border-[#535353] rounded-22px px-19px text-22px leading-43px cursor-pointer"
          >
            {{ $t('home.添加摘要') }}
          </div>
        </div>
      </div>
    </div>
    <!-- 分析 -->
    <div
      class="mb-20px bg-white rounded-8px px-31px py-32px border border-[#E2E0DF]"
      id="2"
    >
      <div class="text-[#000000] text-26px leading-30px mb-7px">
        {{ $t('home.分析') }}
      </div>
      <div class="flex items-center mb-15px">
        <img
          src="@/assets/img/home/see.png"
          class="w-23px h-23px cursor-pointer flex-shrink-0 mr-4px"
          alt=""
        />
        <span class="text-[#828282] text-18px font-400">
          {{ $t('home.对你来说是私人的') }}
        </span>
      </div>
      <div class="flex">
        <img
          src="@/assets/img/home/AnalyticsUser.png"
          class="w-25px h-27px flex-shrink-0 mr-14px"
          alt=""
        />
        <div>
          <div class="text-[#000000] text-21px leading-25px mb-3px">
            0{{ $t('home.个配置文件视图') }}
          </div>
          <div class="text-17px leading-28px">
            {{ $t('home.更新您的个人资料以吸引观众。') }}
          </div>
        </div>
      </div>
    </div>
    <div
      class="mb-20px bg-white px-31px py-32px rounded-8px border border-[#E2E0DF]"
      id="3"
    >
      <div class="text-[#000000] text-26px leading-30px mb-7px">
        {{ $t('home.资源') }}
      </div>
      <div class="flex items-center mb-15px">
        <img
          src="@/assets/img/home/see.png"
          class="w-23px h-23px cursor-pointer flex-shrink-0 mr-4px"
          alt=""
        />
        <span class="text-[#828282] text-18px font-400">
          {{ $t('home.对你来说是私人的') }}
        </span>
      </div>
      <div class="flex">
        <img
          src="@/assets/img/home/notice.png"
          class="w-32px h-32px flex-shrink-0 mr-10px"
          alt=""
        />
        <div>
          <div class="flex items-center">
            <span class="text-[#000000] text-21px leading-25px mb-3px mr-10px">
              0{{ $t('home.创建者模式') }}
            </span>
            <div
              class="bg-[#efefef] rounded-7px px-10px h-28px leading-28px text-[#000000]"
            >
              off
            </div>
          </div>
          <div class="text-17px leading-28px">
            {{
              $t('home.获得发现，在您的个人资料中展示内容，并访问创建者工具')
            }}
          </div>
        </div>
      </div>
    </div>
    <div
      class="mb-20px bg-white rounded-8px px-31px py-32px border border-[#E2E0DF] relative"
      id="4"
    >
      <div class="absolute top-26px right-29px flex items-center">
        <img
          src="@/assets/img/home/ContentEdit.png"
          class="w-44px h-44px flex-shrink-0 cursor-pointer mr-24px"
          alt=""
        />
        <img
          src="@/assets/img/home/ContentDelete.png"
          class="w-44px h-44px flex-shrink-0 cursor-pointer"
          alt=""
        />
      </div>
      <div class="text-[#000000] text-26px leading-30px mb-32px">
        {{ $t('home.个人力量') }}
      </div>
      <div class="text-19px font-400">{{ $t('我是') }}John Smith.</div>
    </div>
    <div
      class="mb-20px bg-white rounded-8px border border-[#E2E0DF] px-31px pt-37px pb-30px"
      id="5"
    >
      <div class="flex items-center justify-between mb-37px">
        <div class="text-[#000000] text-26px leading-30px">
          {{ $t('home.工作经历') }}
        </div>
        <img src="@/assets/img/home/userAdd.png" class="w-23px h-23px cursor-pointer" alt="" />
      </div>
      <div class="flex items-center justify-center mb-14px">
        <div class="flex-1">
          <div class="text-[#000000] text-21px font-500 mb-3px">
            UX Analyst（产品部门）
          </div>
          <div class="text-18px font-500">
            Mobbin·Menlo Park, California,United States
          </div>
        </div>
        <div class="flex items-center">
          <img
            src="@/assets/img/home/ContentEdit.png"
            class="w-44px h-44px flex-shrink-0 cursor-pointer mr-24px"
            alt=""
          />
          <img
            src="@/assets/img/home/ContentDelete.png"
            class="w-44px h-44px flex-shrink-0 cursor-pointer"
            alt=""
          />
        </div>
      </div>

      <div class="text-[#828282] text-17px mb-40px">Jan 2018 - Jan 2022</div>
      <div class="text-[#000000] text-21px leading-25px">
        {{ $t('home.工作内容') }}
      </div>
      <div class="text-18px leading-28px font-500 mb-26px">
        <div v-for="it in jionList">{{ it }}</div>
      </div>
      <div class="text-[#000000] text-21px leading-25px">
        {{ $t('home.工作绩效') }}
      </div>
      <div class="text-18px leading-28px font-500 mb-34px">
        <div v-for="it in jionList">{{ it }}</div>
      </div>

      <div class="text-[#000000] text-19px mb-14px">{{ $t('home.技能') }}</div>

      <div class="flex items-center">
        <div
          v-for="i in 2"
          class="h-41px bg-[#404040] rounded-23px px-15px text-[#ffffff] text-21px leading-41px mr-20px"
        >
          UX X
        </div>
      </div>
    </div>
    <div
      class="mb-20px bg-white rounded-8px border px-31px pt-37px pb-30px border-[#E2E0DF]"
      id="6"
    >
      <div class="flex items-center justify-between mb-37px">
        <div class="text-[#000000] text-26px leading-30px">
          {{ $t('home.项目经验') }}
        </div>
        <img src="@/assets/img/home/userAdd.png" class="w-23px h-23px cursor-pointer" alt="" />
      </div>
      <div class="flex items-center justify-center mb-14px">
        <div class="flex-1">
          <div class="text-[#000000] text-21px font-500 mb-6px">
            {{ $t('home.项目名称') }}
          </div>
          <div class="text-[#828282] text-17px mb-40px">
            Jan 2018 - Jan 2022
          </div>
        </div>
        <div class="flex items-center">
          <img
            src="@/assets/img/home/ContentEdit.png"
            class="w-44px h-44px flex-shrink-0 cursor-pointer mr-24px"
            alt=""
          />
          <img
            src="@/assets/img/home/ContentDelete.png"
            class="w-44px h-44px flex-shrink-0 cursor-pointer"
            alt=""
          />
        </div>
      </div>

      <div class="text-[#000000] text-21px leading-25px">
        {{ $t('home.项目描述') }}
      </div>
      <div class="text-18px leading-28px font-500 mb-26px">
        <div v-for="it in jionList">{{ it }}</div>
      </div>
      <div class="text-[#000000] text-21px leading-25px">
        {{ $t('home.项目工作内容') }}
      </div>
      <div class="text-18px leading-28px font-500 mb-34px">
        <div v-for="it in jionList">{{ it }}</div>
      </div>

      <div class="text-[#000000] text-19px mb-14px">
        {{ $t('home.项目链接') }}
        <span class="text-[#828282] text-17px">www.baidu.com</span>
      </div>
    </div>
    <div class="mb-20px bg-white rounded-8px py-32px px-31px border border-[#E2E0DF]" id="7">
      <div class="flex items-center justify-between mb-38px">
        <div class="text-[#000000] text-26px leading-30px">
          {{ $t('home.教育') }}
        </div>
        <img src="@/assets/img/home/userAdd.png" class="w-23px h-23px cursor-pointer" alt="" />
      </div>

      <div class="flex items-center justify-center mb-14px">
        <div class="flex-1">
          <div class="text-[#000000] text-21px font-500 mb-3px">
            Boston University
          </div>
          <div class="text-18px leading-28px">
            Bachelor's degree, Business Administration
          </div>
        </div>
        <div class="flex items-center">
          <img
            src="@/assets/img/home/ContentEdit.png"
            class="w-44px h-44px flex-shrink-0 cursor-pointer mr-24px"
            alt=""
          />
          <img
            src="@/assets/img/home/ContentDelete.png"
            class="w-44px h-44px flex-shrink-0 cursor-pointer"
            alt=""
          />
        </div>
      </div>

      <div class="text-[#828282] text-17px mb-40px">Jan 2018 - Jan 2022</div>
    </div>

    <div
      class="mb-20px bg-white rounded-8px py-32px px-31px border border-[#E2E0DF]"
      id="8"
    >
      <div class="flex items-center justify-between mb-45px">
        <div class="text-26px text-[#000000]">{{ $t('home.技能') }}</div>
        <img src="@/assets/img/home/userAdd.png" class="w-23px h-23px cursor-pointer" alt="" />
      </div>

      <div class="flex items-center">
        <div
          v-for="i in 2"
          class="h-41px bg-[#404040] rounded-23px px-15px text-[#ffffff] text-21px leading-41px mr-20px"
        >
          UX X
        </div>
      </div>
    </div>

    <div
      class="mb-20px bg-white px-32px py-31px rounded-8px border border-[#E2E0DF]"
      id="9"
      style="height: 500px;"
    >
      <div class="flex items-center justify-between mb-45px">
        <div class="text-26px text-[#000000]">{{ $t('home.权益') }}</div>
        <img src="@/assets/img/home/userAdd.png" class="w-23px h-23px cursor-pointer" alt="" />
      </div>

      <div class="flex items-center flex-wrap">
        <div
          v-for="i in 10"
          class="h-41px bg-[#404040] rounded-23px px-15px text-[#ffffff] text-21px leading-41px mr-20px mb-30px"
        >
          Sport X
        </div>
      </div>
    </div>

    <div
      class="mb-20px bg-white py-32px px-30px rounded-8px border border-[#E2E0DF]"
      id="10"
    >
      <div class="flex items-center justify-between mb-20px">
        <div class="text-26px text-[#000000]">{{ $t('home.梦想的位置') }}</div>
        <img src="@/assets/img/home/userAdd.png" class="w-23px h-23px cursor-pointer" alt="" />
      </div>

      <div class="flex items-center justify-between mb-10px">
        <div>
          <div class="text-[#000000] text-21px mb-19px">UX Analyst</div>
          <div class="text-17px">
            Job type:
            <span class="text-[#828282]">Full-time</span>
          </div>
        </div>
        <div class="flex items-center">
          <img
            src="@/assets/img/home/ContentEdit.png"
            class="w-44px h-44px flex-shrink-0 cursor-pointer mr-24px"
            alt=""
          />
          <img
            src="@/assets/img/home/ContentDelete.png"
            class="w-44px h-44px flex-shrink-0 cursor-pointer"
            alt=""
          />
        </div>
      </div>
      <div class="text-17px mb-10px">
        {{ $t('home.工作城市') }}:
        <span class="text-[#828282]">Menlo Park, California,United States</span>
      </div>
      <div class="text-17px mb-10px">
        {{ $t('home.预期行业') }}:
        <span class="text-[#828282]">unlimited</span>
      </div>
      <div class="text-17px">
        {{ $t('home.薪水') }}:
        <span class="text-[#828282]">12K-18K</span>
      </div>
    </div>

    <div class="mb-20px bg-white rounded-8px border border-[#E2E0DF]" id="11">
      <div class="flex items-center px-31px py-32px">
        <div class="flex-1">
          <span class="text-[#000000] text-26px font-normal mr-25px">
            {{ $t('home.自定义项目') }}
          </span>
          <span class="text-[#828282] text-20px font-500">
            {{ $t('home.描述个性化内容') }}
          </span>
        </div>
        <img
          src="@/assets/img/home/userAdd.png"
          class="w-23px h-23px cursor-pointer flex-shrink-0"
          alt=""
        />
      </div>

      <div
        class="py-34px px-31px border-[#E2E0DF]"
        style="border-top: 1px solid #eeeeee;"
        v-for="i in 2"
      >
        <div class="flex items-center justify-between mb-31px">
          <div class="text-[#000000] text-26px">自定义标题1</div>
          <div class="flex items-center">
            <img
              src="@/assets/img/home/ContentEdit.png"
              class="w-44px h-44px flex-shrink-0 cursor-pointer mr-24px"
              alt=""
            />
            <img
              src="@/assets/img/home/ContentDelete.png"
              class="w-44px h-44px flex-shrink-0 cursor-pointer"
              alt=""
            />
          </div>
        </div>

        <div class="text-21px leading-25px">
          自定义内容自定义内容自定义内容自定义内容自定义内容自定义内容自定义内容自定义内容自定义
          内容自定义内容
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { NCard, NForm, NFormItem, NInput, NButton } from 'naive-ui'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const emit = defineEmits(['setAnchorpoint'])
const props = defineProps({
  anchorpoint: Number,
})
const activeIndex = ref(0)
const allNumber = ref(7)
const nowNumber = ref(5)
const jionList = ref([
  '1、Project descriptionProject',
  '2、descriptionProject description',
  '3、Project descriptionProject description',
  '4、Project descriptionProject descriptionProject',
  '5、descriptionProject descriptionProject description',
  '6、Project descriptionProject descriptionProject description',
])
watch(
  props,
  () => {
    console.log('11111111111' + props.anchorpoint)
    console.log(document.getElementById(props.anchorpoint))
    document.getElementById(props.anchorpoint).scrollIntoView({
      behavior: 'smooth',
    })
  },
  // { immediate: true }
)
</script>

<style lang="scss" scoped>
.menuItem {
  position: relative;
}

.menuActive {
  background: rgba(51, 95, 62, 0.1);
  color: #219653;
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
