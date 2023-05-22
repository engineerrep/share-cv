<template>
  <div class="edit" v-if="bulletBoxEdit">
    <div class="edit-avatar center" v-if="bulletBoxEditAvatar">
      <div class="edit-avatar-head df">
        <div class="edit-avatar-title">Edit</div>
        <div
            class="edit-avatar-closure cursor-pointer"
            @click="bulletBoxEditAvatar = false"
        >
          X
        </div>
      </div>
      <div class="edit-avatar-main">
        <div class="avatar-img">
          <CutOut></CutOut>
        </div>
        <p>Drag to reposition photo</p>
        <img
            src="@/assets/img/personal-center/to-refresh-img.png"
            class="to-refresh-img"
        />
      </div>
      <n-space class="edit-avatar-tail" justify="end">
        <n-button round>Default</n-button>
        <n-button type="info" round> Info</n-button>
      </n-space>
    </div>

    <!-- 在你的个人资料中编辑 -->
    <div
        class="edit-in-your-profile center"
        v-else-if="bulletBoxeditInYourProfile"
    >
      <n-space justify="space-between" class="edit-in-your-profile-title">
        <p class="">Edit in your profile</p>
        <p @click="bulletBoxeditInYourProfile = false" class="cursor-pointer">
          X
        </p>
      </n-space>
      <div class="edit-in-your-profile-main">
        <div class="indicates-required">*Indicates required</div>
        <n-form ref="formRef" :model="model" :rules="rules">
          <n-form-item path="age" label="fullName">
            <n-input v-model:value="modelRef.fullName" @keydown.enter.prevent/>
          </n-form-item>
          <n-form-item path="password" label="nickname">
            <n-input
                v-model:value="modelRef.nickname"
                type="password"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-form-item path="headline" label="headline">
            <n-input
                v-model:value="modelRef.headline"
                type="password"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-form-item path="positio" label="positio">
            <n-input
                v-model:value="modelRef.positio"
                type="password"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <p class="your-location">Let's confirm your location</p>
          <p class="recruiters-note">Get noticed by recruiters in your area.</p>
          <n-form-item path="countryRegion" label="country/Region">
            <n-input
                v-model:value="modelRef.countryRegion"
                type="password"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-form-item path="city" label="city">
            <n-input
                v-model:value="modelRef.city"
                type="password"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
        </n-form>
      </div>
      <n-space justify="end" class="save-btn">
        <n-button type="info" round @click="submitEdit">Save</n-button>
      </n-space>
    </div>

    <!-- 添加到您的个人资料   -->
    <div
        class="add-toyour-profile center"
        v-else-if="bulletBoxAddToyourProfile"
    >
      <n-space class="add-toyour-profile-head" justify="space-between">
        <p>Add to your profile</p>
        <p class="cursor-pointer" @click="bulletBoxAddToyourProfile = false">
          X
        </p>
      </n-space>
      <div class="add-toyour-profile-main">
        <div v-if="addToyourProfilePage == 1">
          <div>
            <p class="recent-Job-experience">
              Is this your most re cent Job experience?
            </p>
            <p class="match-the-right-job">
              Make sure recruiters match you with the right jobs.
            </p>
            <!---->
          </div>
          <n-space class="space">
            <!--            <p class="img"><img src="\"/></p>-->
            <div>
              <p>UX Analyst</p>
              <p>Mobbinss</p>
            </div>
          </n-space>
          <n-radio-group
              v-model:value="isRecentJobExperienceVal"
              :default-value="isRecentJobExperienceVal"
              name="radiogroup"
          >
            <div>
              <n-radio value="0"> Yes</n-radio>
            </div>
            <div>
              <n-radio value="1"> NO</n-radio>
            </div>
          </n-radio-group>
        </div>

        <div v-else-if="addToyourProfilePage == 2">
          <div>
            <p class="skill-title">
              Let's add your skills to UX Analyst at Mobbin
            </p>
            <p class="skill-top5">
              We recommend adding your top 5 used in this role. They'll also
              appear in your Skills section.
            </p>
            <p class="recommended-skills">
              Suggested skills based on your title and the skills recruiters are
              looking for
            </p>
            <n-input></n-input>
          </div>
        </div>
        <div v-else-if="addToyourProfilePage == 3">
          <div>
            <p class="skill-title">Let's confirm your location</p>
            <p class="skill-top5">Get noticed by recruiters in your area.</p>
            <n-form ref="formRef" :model="modelRef" :rules="rules">
              <n-form-item path="countryRegion" label="country/Region">
                <n-input
                    v-model:value="modelRef.countryRegion"
                    type="password"
                    @input="handlePasswordInput"
                    @keydown.enter.prevent
                />
              </n-form-item>
              <n-form-item path="city" label="city">
                <n-input
                    v-model:value="modelRef.city"
                    type="password"
                    @input="handlePasswordInput"
                    @keydown.enter.prevent
                />
              </n-form-item>
            </n-form>
          </div>
        </div>
      </div>

      <n-space justify="space-between" class="tail">
        <p>{{ addToyourProfilePage }} of 3</p>
        <n-button
            type="info"
            round
            @click="++addToyourProfilePage"
            v-if="addToyourProfilePage < 3"
        >Continue
        </n-button
        >
        <n-button
            type="info"
            round
            @click="submitAddToyourProfile"
            v-else-if="addToyourProfilePage > 2"
        >Done
        </n-button>
      </n-space>
    </div>

    <!--  背景编辑  -->
    <div
        class="edit-avatar background-editor center"
        v-else-if="bulletBoBackgroundEditorr"
    >
      <div class="edit-avatar-head df">
        <div class="edit-avatar-title">Edit</div>
        <div
            class="edit-avatar-closure cursor-pointer"
            @click="bulletBoBackgroundEditorr = false"
        >
          X
        </div>
      </div>
      <div class="background-editor-main pr">
        <div class="background-editor-content"><img src=""/></div>
        <img
            src="@/assets/img/personal-center/to-refresh-img.png"
            class="to-refresh-img cursor-pointer"
        />
      </div>
      <n-space class="edit-avatar-tail" justify="end">
        <n-button round>Change photo</n-button>
        <n-button type="info" round @click="submitBackgroundEditorr">
          Apply
        </n-button>
      </n-space>
    </div>

    <!-- 分享您的个人资料  -->
    <div
        class="edit-avatar share-youry-profile center"
        v-else-if="bulletBoxShareYourProfile"
    >
      <div class="edit-avatar-head df">
        <div class="edit-avatar-title">Share your profile</div>
        <div
            class="edit-avatar-closure cursor-pointer"
            @click="bulletBoxShareYourProfile = false"
        >
          X
        </div>
      </div>
      <div class="share-youry-profile-main pr">
        <n-space>
          <!--          <p class="img"><img src="\"/></p>-->
          <div>
            <p>John Smith</p>
            <p>UX Analyst at Mobbin</p>
            <p>Menlo Park, CA</p>
          </div>
        </n-space>
      </div>
      <n-space class="edit-avatar-tail" justify="end">
        <n-button type="info" round @click="submitShareYourProfile">
          Share
        </n-button>
      </n-space>
    </div>

    <!-- 添加到自定义属性  -->
    <div
        class="edit-avatar add-to-custom-attributes center"
        v-else-if="bulletBoxAddToCustomAttributes"
    >
      <div class="edit-avatar-head df">
        <div class="edit-avatar-title">Add to custom attributes</div>
        <div
            class="edit-avatar-closure cursor-pointer"
            @click="bulletBoxAddToCustomAttributes = false"
        >
          X
        </div>
      </div>
      <div class="share-youry-profile-main pr">
        <n-form
            ref="formRef"
            :model="addToCustomAttributes"
            :rules="rulesAddToCustomAttributes"
        >
          <n-form-item label="Title">
            <n-input
                v-model:value="addToCustomAttributes.title"
                type="password"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-form-item
              ref="rPasswordFormItemRef"
              first
              path="content"
              label="Content"
          >
            <n-input
                v-model:value="addToCustomAttributes.content"
                type="textarea"
                maxlength="2000"
                show-count
                @keydown.enter.prevent
                style="height: 342px"
            />
          </n-form-item>
        </n-form>
      </div>
      <n-space class="edit-avatar-tail" justify="end">
        <n-button type="info" round @click="submitAddToCustomAttributes">
          Share
        </n-button>
      </n-space>
    </div>

    <!-- 添加到教育  -->
    <div
        class="edit-avatar add-to-education center"
        v-else-if="bulletBoxAddToEducation"
    >
      <div class="edit-avatar-head df">
        <div class="edit-avatar-title">Add to Education</div>
        <div
            class="edit-avatar-closure cursor-pointer"
            @click="bulletBoxAddToEducation = false"
        >
          X
        </div>
      </div>
      <div class="share-youry-profile-main pr">
        <n-form
            ref="formRef"
            :model="addToEducation"
            :rules="rulesbulletBoxAddToEducation"
        >
          <n-form-item label="University/School*" path="universitySchool">
            <n-input
                v-model:value="addToEducation.universitySchool"
                type="text"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-form-item label="Title" path="degree">
            <n-input
                v-model:value="addToEducation.degree"
                type="text"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-space align="center" justify="space-between">
            <n-form-item
                label="Start-time*"
                path="startTime"
                style="width: 298px"
            >
              <n-input
                  v-model:value="addToEducation.startTime"
                  type="text"
                  @input="handlePasswordInput"
                  @keydown.enter.prevent
              />
            </n-form-item>
            <span>to</span>
            <n-form-item label="End-time*" path="EndTime" style="width: 298px">
              <n-input
                  v-model:value="addToEducation.EndTime"
                  type="text"
                  @input="handlePasswordInput"
                  @keydown.enter.prevent
              />
            </n-form-item>
          </n-space>
        </n-form>
      </div>
      <n-space class="edit-avatar-tail" justify="end">
        <n-button type="info" round @click="submitaddToEducation">
          Share
        </n-button>
      </n-space>
    </div>

    <!-- 添加到兴趣   -->
    <div
        class="add-toyour-profile add-to-interests center"
        v-else-if="bulletBoxAddToInterests"
    >
      <n-space class="add-toyour-profile-head" justify="space-between">
        <p>Add to Interests</p>
        <p class="cursor-pointer" @click="bulletBoxAddToInterests = false">X</p>
      </n-space>
      <div class="add-toyour-profile-main">
        <div>
          <p class="title">title</p>
          <n-input v-model:value="addToInterests"></n-input>
          <div>
            <n-button
                type="primary"
                round
                v-for="(item, index) in addToInterestsList"
                :key="index"
                class="add-to-interests-item"
            >
              {{ item.txt }} X
            </n-button>
          </div>
        </div>
      </div>

      <n-space justify="end" class="tail">
        <n-button type="info" round @click="submitaddToInterests"
        >Save
        </n-button>
      </n-space>
    </div>

    <!-- 添加到项目经验  -->
    <div
        class="edit-avatar add-to-project-experience center"
        v-else-if="bulletBoxAddToProjectExperience"
    >
      <div class="edit-avatar-head df">
        <div class="edit-avatar-title">Add to Project experience</div>
        <div
            class="edit-avatar-closure cursor-pointer"
            @click="bulletBoxAddToProjectExperience = false"
        >
          X
        </div>
      </div>
      <div class="share-youry-profile-main pr">
        <n-form
            ref="formRef"
            :model="addToEducation"
            :rules="rulesbulletBoxAddToEducation"
        >
          <n-form-item label="Project name" path="universitySchool">
            <n-input
                v-model:value="addToEducation.universitySchool"
                type="text"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-form-item
              label="Project description"
              path="degree"
              style="height: 145px"
          >
            <n-input
                v-model:value="addToEducation.degree"
                type="textarea"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-form-item
              label="Project Job content"
              path="degree"
              style="height: 145px"
          >
            <n-input
                v-model:value="addToEducation.degree"
                type="textarea"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-space align="center" justify="space-between">
            <n-form-item
                label="Start-time*"
                path="startTime"
                style="width: 298px"
            >
              <n-input
                  v-model:value="addToEducation.startTime"
                  type="text"
                  @input="handlePasswordInput"
                  @keydown.enter.prevent
              />
            </n-form-item>
            <span>to</span>
            <n-form-item label="End-time*" path="EndTime" style="width: 298px">
              <n-input
                  v-model:value="addToEducation.EndTime"
                  type="text"
                  @input="handlePasswordInput"
                  @keydown.enter.prevent
              />
            </n-form-item>
          </n-space>

          <n-form-item label="Project link">
            <n-input
                v-model:value="addToEducation.degree"
                type="text"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
        </n-form>
      </div>
      <n-space class="edit-avatar-tail" justify="end">
        <n-button type="info" round @click="submitaddToEducation">
          Share
        </n-button>
      </n-space>
    </div>

    <!-- 添加到梦想位置  -->
    <div
        class="edit-avatar add-to-dream-position center"
        v-else-if="bulletBoxAddToDreamPosition"
    >
      <div class="edit-avatar-head df">
        <div class="edit-avatar-title">Add to dream position</div>
        <div
            class="edit-avatar-closure cursor-pointer"
            @click="bulletBoxAddToDreamPosition = false"
        >
          X
        </div>
      </div>
      <div class="share-youry-profile-main pr">
        <n-form
            ref="formRef"
            :model="addToDreamPosition"
            :rules="rulesbulletBoxAddToDreamPosition"
        >
          <n-form-item label="Job type" path="universitySchool">
            <n-radio-group
                v-model:value="addToDreamPosition.JobType"
                :default-value="addToDreamPosition.JobType"
                name="radiobuttongroup1"
            >
              <n-radio-button value="Full time" label="Full time" round/>
              <n-radio-button value="Part time" label="Part time" round/>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="Project name" path="JobCity">
            <n-input
                v-model:value="addToDreamPosition.JobCity"
                type="text"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>

          <n-form-item label="Project link" path="dreamPosition">
            <n-input
                v-model:value="addToDreamPosition.dreamPosition"
                type="text"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-form-item label="Project link" path="expectIndustry">
            <n-input
                v-model:value="addToDreamPosition.expectIndustry"
                type="text"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>

          <n-space align="center">
            <n-form-item label="Start-time*" path="salaryRequirements">
              <n-date-picker
                  v-model:value="addToDreamPosition.startTime"
                  type="date"
              />
            </n-form-item>
            <span>to</span>
            <n-form-item label="End-time" path="endTime" style="width: 298px">
              <n-date-picker
                  v-model:value="addToDreamPosition.endTime"
                  type="date"
              />
            </n-form-item>
          </n-space>
        </n-form>
      </div>
      <n-space class="edit-avatar-tail" justify="end">
        <n-button type="info" round @click="submitAddToDreamPosition">
          Share
        </n-button>
      </n-space>
    </div>

    <!-- 添加到技能   -->
    <div
        class="add-toyour-profile add-to-interests center"
        v-else-if="bulletBoxAddToSkills"
    >
      <n-space class="add-toyour-profile-head" justify="space-between">
        <p>Add to Skills</p>
        <p class="cursor-pointer" @click="bulletBoxAddToSkills = false">X</p>
      </n-space>
      <div class="add-toyour-profile-main">
        <div>
          <p class="title">title</p>
          <n-input v-model:value="addToSkillsInput"></n-input>
          <div>
            <n-button
                type="primary"
                round
                v-for="(item, index) in addToSkillsList"
                :key="index"
                class="add-to-interests-item"
            >
              {{ item.txt }} X
            </n-button>
          </div>
        </div>
      </div>

      <n-space justify="end" class="tail">
        <n-button type="info" round @click="submitaAddToSkills">Save</n-button>
      </n-space>
    </div>

    <!-- 增加工作经验  -->
    <div
        class="edit-avatar add-to-project-experience-bullet-box center"
        v-else-if="bulletBoxAddToJobExperience"
    >
      <div class="edit-avatar-head df">
        <div class="edit-avatar-title">Add to Job experience</div>
        <div
            class="edit-avatar-closure cursor-pointer"
            @click="bulletBoxAddToJobExperience = false"
        >
          X
        </div>
      </div>
      <div class="share-youry-profile-main pr">
        <n-form
            ref="formRef"
            :model="addToJobExperienceForm"
            :rules="rulesAddToJobExperience"
        >
          <n-form-item label="companyName" path="companyName">
            <n-input
                v-model:value="addToJobExperienceForm.companyName"
                type="text"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-form-item label="city" path="city">
            <n-input
                v-model:value="addToJobExperienceForm.city"
                type="text"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>

          <n-space align="center" justify="space-between">
            <n-form-item label="department" style="width: 298px">
              <n-input
                  v-model:value="addToJobExperienceForm.department"
                  type="text"
                  @input="handlePasswordInput"
                  @keydown.enter.prevent
              />
            </n-form-item>
            <span>to</span>
            <n-form-item label="jobTitle" style="width: 298px">
              <n-input
                  v-model:value="addToJobExperienceForm.jobTitle"
                  type="text"
                  @input="handlePasswordInput"
                  @keydown.enter.prevent
              />
            </n-form-item>
          </n-space>
          <n-space align="center" justify="space-between">
            <n-form-item
                label="Start-tim"
                path="startTime"
                style="width: 298px"
            >
              <n-input
                  v-model:value="addToJobExperienceForm.startTime"
                  type="text"
                  @input="handlePasswordInput"
                  @keydown.enter.prevent
              />
            </n-form-item>
            <span>to</span>
            <n-form-item label="End-time" path="EndTime" style="width: 298px">
              <n-input
                  v-model:value="addToJobExperienceForm.endTime"
                  type="text"
                  @input="handlePasswordInput"
                  @keydown.enter.prevent
              />
            </n-form-item>
          </n-space>

          <n-form-item
              label="Job content"
              path="JobContent"
              style="height: 145px"
          >
            <n-input
                v-model:value="addToJobExperienceForm.JobContent"
                type="textarea"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-form-item label="Job performance">
            <n-input
                v-model:value="addToJobExperienceForm.JobPerformance"
                type="textarea"
                @input="handlePasswordInput"
                @keydown.enter.prevent
            />
          </n-form-item>
          <n-form-item label="haveSkills" path="haveSkills">
            <n-cascader
                v-model:value="addToJobExperienceForm.haveSkills"
                placeholder=""
                :options="optionsHaveSkills"
            />
          </n-form-item>
        </n-form>
      </div>
      <n-space class="edit-avatar-tail" justify="end">
        <n-button type="info" round @click="submitAddToJobExperience">
          Share
        </n-button>
      </n-space>
    </div>

    <!-- 增加个人实力   -->
    <div
        class="add-toyour-profile add-to-interests center"
        v-else-if="bulletBoxAddToPersonalStrength"
    >
      <n-space class="add-toyour-profile-head" justify="space-between">
        <p class="break-words">Add to personal strength</p>
        <p
            class="cursor-pointer"
            @click="bulletBoxAddToPersonalStrength = false"
        >
          X
        </p>
      </n-space>
      <div class="add-toyour-profile-main">
        <div>
          <p class="title">Content</p>
          <n-input
              v-model:value="addToPersonalStrengthVal"
              type="textarea"
              style="height: 342px"
          ></n-input>
        </div>
      </div>
      <div>
        <n-space justify="end" class="tail">
          <n-button type="info" round @click="submitaAddToPersonalStrength"
          >Save
          </n-button>
        </n-space>
      </div>
    </div>

    <!-- 删除消息   -->
    <div v-if="showModal">
      <n-modal v-model:show="showModal">
        <n-card
            style="width: 600px"
            title="Delete message"
            :bordered="false"
            size="huge"
            role="dialog"
            aria-modal="true"
        >
          <template #header-extra>
            <div @click="showModal = false" class="cursor-pointer">X</div>
          </template>
          Cannot be restored after deletion, confirm deletion？
          <template #footer>
            <n-space justify="end">
              <n-button
                  type="info"
                  round
                  @click="onPositiveClick"
                  class="cursor-pointer"
              >
                Delete
              </n-button
              >
            </n-space>
          </template>
        </n-card>
      </n-modal>
    </div>
  </div>
</template>

<script setup>
import {
  NButton,
  NCard,
  NCascader,
  NDatePicker,
  NForm,
  NFormItem,
  NInput,
  NModal,
  NRadio,
  NRadioButton,
  NRadioGroup,
  NSpace,
  useDialog,
} from "naive-ui";

import CutOut from "@/components/personal-wditor/cut-out.vue";

// if(process.browser) {
//   vueCropper = require('vue-cropper')
//   Vue.use(vueCropper.default)
// }
let bulletBoxEdit = useState("_bulletBoxEdit", () => true);
let boxVueCropper = ref(false);
// 编辑头像-弹框
let bulletBoxEditAvatar = useState("_bulletBoxEditAvatar");
// 在你的个人资料中编辑-弹框
let bulletBoxeditInYourProfile = useState("_bulletBoxeditInYourProfile");
let fileList = ref([
  {
    id: "a",
    name: "我是上传出错的普通文件.png",
    status: "error",
  },
]);


let previews = ref({});
//实时预览函数
const imgLoad = (msg) => {
  console.log("工具初始化函数=====" + msg);
};
//实时预览函数
const realTime = (data) => {
  previews.value = data;
};
// 添加到您的个人资料-弹框
let bulletBoxAddToyourProfile = useState("_bulletBoxAddToyourProfile");

const formRef = ref(null);
const rPasswordFormItemRef = ref(null);
const modelRef = ref({
  fullName: 1,
  nickname: null,
  headline: null,
  positio: null,
  countryRegion: null,
  city: null,
});
console.log(modelRef.value.age);

function validatePasswordStartWith(rule, value) {
  return (
      !!modelRef.value.password &&
      modelRef.value.password.startsWith(value) &&
      modelRef.value.password.length >= value.length
  );
}

function validatePasswordSame(rule, value) {
  return value === modelRef.value.password;
}

const rules = {
  headline: [
    {
      required: true,
      message: "请输入密码",
      trigger: ["input", "blur"],
    },
  ],
  positio: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("需要年龄");
        } else if (!/^\d*$/.test(value)) {
          return new Error("年龄应该为整数");
        } else if (Number(value) < 18) {
          return new Error("年龄应该超过十八岁");
        }
        return true;
      },
      trigger: ["input", "blur"],
    },
  ],
  countryRegion: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("需要年龄");
        } else if (!/^\d*$/.test(value)) {
          return new Error("年龄应该为整数");
        } else if (Number(value) < 18) {
          return new Error("年龄应该超过十八岁");
        }
        return true;
      },
      trigger: ["input", "blur"],
    },
  ],
  city: [
    {
      required: true,
      message: "请再次输入密码",
      trigger: ["input", "blur"],
    },
    {
      validator: validatePasswordStartWith,
      message: "两次密码输入不一致",
      trigger: "input",
    },
    {
      validator: validatePasswordSame,
      message: "两次密码输入不一致",
      trigger: ["blur", "password-input"],
    },
  ],
};
const handlePasswordInput = () => {
  if (modelRef.value.reenteredPassword) {
    rPasswordFormItemRef.value?.validate({trigger: "password-input"});
  }
};

/* 提交编辑 */
const submitEdit = () => {
  let data = {
    fullName: modelRef.value.fullName,
    nickname: modelRef.value.nickname,
    headline: modelRef.value.headline,
    positio: modelRef.value.positio,
    countryRegion: modelRef.value.countryRegion,
    city: modelRef.value.city,
  };
  bulletBoxeditInYourProfile.value = false;
  console.log("提交编辑-请求", data);
};

// 添加到您的个人资料-页
let addToyourProfilePage = ref(1);

// 是否是最近的工作经历
let isRecentJobExperienceVal = ref("0");

// 添加到您的个人资料
const submitAddToyourProfile = () => {
  bulletBoxAddToyourProfile.value = false;
};
// 背景编辑-弹框
let bulletBoBackgroundEditorr = useState("_bulletBoBackgroundEditorr");

// 提交-背景编辑
const submitBackgroundEditorr = () => {
  bulletBoBackgroundEditorr.value = false;
};
// 分享您的个人资料-弹框
let bulletBoxShareYourProfile = useState("_bulletBoxShareYourProfile");

// 分享您的个人资料-操作
const submitShareYourProfile = () => {
  console.log("分享您的个人资料-操作");
};

// 添加到自定义属性-弹框
let bulletBoxAddToCustomAttributes = useState(
    "_bulletBoxAddToCustomAttributes"
);
// 添加到自定义属性-表单
let addToCustomAttributes = ref({
  title: "",
  content: "",
});
// 添加到自定义属性-表单-验证规则
let rulesAddToCustomAttributes = {
  content: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("content");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
};

// 添加到自定义属性-操作
const submitAddToCustomAttributes = () => {
  console.log("分享您的个人资料-操作");
};
// 添加到教育-弹框
let bulletBoxAddToEducation = useState("_bulletBoxAddToEducation");
// 添加到教育-表单
let addToEducation = ref({
  universitySchool: "",
  degree: "",
  startTime: "",
  EndTime: "",
});
// 添加到教育-表单-验证规则
let rulesbulletBoxAddToEducation = {
  universitySchool: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("universitySchool");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
  degree: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("universitySchool");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
  startTime: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("universitySchool");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
  EndTime: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("universitySchool");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
};

// 添加到教育-操作
const submitaddToEducation = () => {
  console.log("分享您的个人资料-操作");
};

// 添加到兴趣-弹框
let bulletBoxAddToInterests = useState("_bulletBoxAddToInterests");
let addToInterests = ref();
let addToInterestsList = ref([{txt: "sport"}]);
// 添加到兴趣-操作
const submitaddToInterests = () => {
  // console.log(isRecentJobExperienceVal.value)
  if (addToInterests.value) {
    addToInterestsList.value.push({txt: addToInterests.value});
    addToInterests.value = null;
  }

  console.log("添加到兴趣");
};

// 添加到项目经验-弹框
let bulletBoxAddToProjectExperience = useState(
    "_bulletBoxAddToProjectExperience"
);
// let addToInterests = ref()
// let addToInterestsList = ref([
//   {txt: "sport"}
// ])
// // 添加到兴趣-操作
// const submitaddToInterests = () => {
//   // console.log(isRecentJobExperienceVal.value)
//   if (addToInterests.value) {
//     addToInterestsList.value.push({txt: addToInterests.value})
//     addToInterests.value = null
//   }
//
//   console.log('添加到兴趣')
// }

// 添加到梦想位置-弹框
let bulletBoxAddToDreamPosition = useState("_bulletBoxAddToDreamPosition");
let addToDreamPosition = ref({
  JobType: "Full time",
  JobCity: "",
  dreamPosition: "",
  expectIndustry: "",
  startTime: 1183135260000,
  endTime: 1183135260000,
});

// 添加到梦想位置-规则
let rulesbulletBoxAddToDreamPosition = {
  universitySchool: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("universitySchool");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
  degree: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("universitySchool");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
  startTime: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("universitySchool");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
  EndTime: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("universitySchool");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
};

// 添加到梦想位置-操作
const submitAddToDreamPosition = () => {
  //
  /*{
      JobType: "Full time",
      JobCity: "",
      dreamPosition: "",
      expectIndustry: "",
      startTime: "",
      endTime:""
    }*/
  let data = addToDreamPosition.value;
  console.log("添加到梦想位置", data);
};

let bulletBoxAddToSkills = useState("_bulletBoxAddToSkills");
// 添加到技能-规则
let addToSkillsInput = ref();
let addToSkillsList = ref([{txt: "UX"}]);

// 添加到技能-操作
const submitaAddToSkills = () => {
  addToSkillsList.value.push({txt: addToSkillsInput.value});
  addToSkillsInput.value = null;
  let data = addToDreamPosition.value;
  console.log("添加到技能", data);
};


// 增加工作经验-弹框
let bulletBoxAddToJobExperience = useState("_bulletBoxAddToJobExperience");
let optionsHaveSkills = ref([
  {
    value: 1,
    label: 1,
  },
]);
let addToJobExperienceForm = ref({
  companyName: "",
  city: "",
  department: "",
  jobTitle: "",
  startTime: "",
  endTime: "",
  JobContent: "",
  JobPerformance: "",
  haveSkills: "",
});
let rulesAddToJobExperience = {
  companyName: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("wcompanyNamew");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
  city: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("city");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
  startTime: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("startTime");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
  endTime: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("endTime");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
  JobContent: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("JobContent");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
  haveSkills: [
    {
      required: true,
      validator(rule, value) {
        if (!value) {
          return new Error("haveSkills");
        }
      },
      trigger: ["input", "blur"],
    },
  ],
};

const submitAddToJobExperience = () => {
  let data = addToJobExperienceForm.value;
  console.log("增加工作经验", data);
};
let showModal = ref(false);
const dialog = useDialog();
// 增加个人实力
let bulletBoxAddToPersonalStrength = useState(
    "_bulletBoxAddToPersonalStrength"
);
let addToPersonalStrengthVal = ref();
const submitaAddToPersonalStrength = () => {
  console.log("增加个人实力", addToPersonalStrengthVal.value);
  showModal.value = true;
};

const onNegativeClick = () => {
  console.log("取消");
  showModal.value = false;
};
const onPositiveClick = () => {
  console.log("确认");
  showModal.value = false;
};

//
</script>

<style lang="scss" scoped>
@import "vue-cropper/dist/index.css";

.edit {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  z-index: 40;
  background-color: rgba(#363636, 0.4);
}

// 裁剪
.vue-cropper {
  width: 600px;
  height: 230px;
}

.cropper-content {
  display: flex;
  display: -webkit-flex;
  justify-content: flex-end;
}

.cropper-box {
  flex: 1;
  width: 100%;
}

.show-preview {
  flex: 1;
  -webkit-flex: 1;
  display: flex;
  display: -webkit-flex;
  justify-content: center;
}

.cropper {
  width: auto;
  height: 300px;
}

.preview {
  overflow: hidden;
  border: 1px solid #67c23a;
  background: #cccccc;
}

.cropper-view-box {
  display: block;
  overflow: hidden;
  width: 100%;
  height: 100%;
  outline: 1px solid #39f;
  outline-color: rgba(51, 153, 255, .75);
  user-select: none;
}

.img {
  width: 69px;
  height: 69px;
  background: #d9d9d9;
  border-radius: 0px 0px 0px 0px;
  opacity: 1;
}


.edit-avatar {
  width: 993px;
  height: 681px;
  background: #ffffff;
  border-radius: 13px 13px 13px 13px;
  opacity: 1;

  .edit-avatar-head {
    justify-content: space-between;
  }

  .edit-avatar-title {
    font-size: 25px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #000000;
    line-height: 82px;
    padding-left: 32px;
  }


  .edit-avatar-main {
    width: 993px;
    height: 493px;
    background: #000000;
    opacity: 1;
    color: #fff;
    padding-top: 20px;
    text-align: center;
    position: relative;
  }

  .to-refresh-img {
    position: absolute;
    bottom: 50px;
    right: 26px;
  }

  .avatar-img {
    width: 307px;
    height: 307px;
    border: 2px solid #fff;
    border-radius: 154px 154px 154px 154px;
    opacity: 1;
    margin: 0px auto 18px;

    .n-upload {
      width: 307px;
      height: 307px;
      border: 2px solid #fff;
      border-radius: 154px 154px 154px 154px;
    }
  }

  .edit-avatar-tail {
    padding-top: 26px;
    padding-right: 32px;
  }
}

.edit-avatar-closure {
  width: 24px;
  line-height: 82px;
  padding-right: 32px;
}

.bullet-box-add-img {
  width: 20px;
  height: 20px;
}

.edit-in-your-profile {
  width: 997px;
  //height: 986px;
  padding-bottom: 10px;
  background: #ffffff;
  border-radius: 12px 12px 12px 12px;
  opacity: 1;
  border: 1px solid #e0e0e0;

  .edit-in-your-profile-main {
    padding: 11px 38px 48px 20px;

    .indicates-required {
      height: 19px;
      font-size: 16px;
      font-family: Inter-Medium, Inter;
      font-weight: 500;
      color: #828282;
      line-height: 19px;
      margin-bottom: 20px;
    }
  }

  .edit-in-your-profile-title {
    font-size: 25px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #000000;
    line-height: 29px;
    padding: 25px 38px 26px 30px;
    border-bottom: 1px solid #ededed;
  }

  .your-location {
    height: 36px;
    font-size: 30px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #000000;
    line-height: 35px;
    margin-top: 25px;
  }

  .recruiters-note {
    height: 22px;
    font-size: 18px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #828282;
    line-height: 21px;
    margin-bottom: 23px;
  }

  .save-btn {
    opacity: 1;
    border-top: 1px solid #eeeeee;

    .n-button {
      width: 91px;
      height: 42px;
      margin-top: 17px;
      margin-right: 39px;
    }
  }
}

.add-toyour-profile {
  width: 736px;
  background: #ffffff;
  border-radius: 12px 12px 12px 12px;
  opacity: 1;

  .add-toyour-profile-head {
    padding: 24px 34px 25px 32px;
    font-size: 25px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #000000;
    line-height: 29px;
    border-bottom: 1px solid #eaeaea;
  }

  .recent-Job-experience {
    height: 36px;
    font-size: 30px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #000000;
    line-height: 35px;
  }

  .add-toyour-profile-main {
    padding: 22px 32px 29px 32px;
    border-bottom: 1px solid #f1f1f1;
  }

  .match-the-right-job {
    height: 22px;
    font-size: 18px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #333333;
    line-height: 21px;
  }

  .space {
    margin-top: 24px;
    margin-bottom: 30px;
  }

  .ux-analystss {
    height: 25px;
    font-size: 21px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #000000;
    line-height: 25px;
  }

  .mobbinss {
    height: 22px;
    font-size: 18px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #333333;
    line-height: 21px;
  }

  .tail {
    padding: 22px 32px;
  }

  .skill-title {
    height: 36px;
    font-size: 30px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #000000;
    line-height: 35px;
  }

  .skill-top5 {
    width: 682px;
    height: 49px;
    font-size: 18px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #585858;
    line-height: 21px;
  }

  .recommended-skills {
    width: 632px;
    height: 22px;
    font-size: 18px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #828282;
    line-height: 21px;
    margin-top: 24px;
    margin-bottom: 10px;
  }
}

.background-editor {
  .background-editor-main {
    width: 993px;
    height: 493px;
    background: #000000;
    opacity: 1;
    padding: 92px 33px 0;
  }

  .background-editor-content {
    width: 928px;
    height: 232px;
    background-color: #0a66c2;
    opacity: 1;
  }
}

.share-youry-profile {
  height: 328px;

  .share-youry-profile-main {
    padding: 25px 32px 48px;
    border-top: 1px solid #f5f5f5;
    border-bottom: 1px solid #f5f5f5;
  }
}

.add-to-custom-attributes {
  height: 738px;
}

.add-to-education {
  width: 736px;
  height: 500px;
}

.add-to-project-experience {
  width: 736px;
  height: 780px;
}

.center {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: #fff;
}

.share-youry-profile-main {
  padding: 25px 32px 48px;
  border-top: 1px solid #f5f5f5;
  border-bottom: 1px solid #f5f5f5;
}

.add-to-interests {
  .title {
    height: 22px;
    font-size: 18px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #333333;
    line-height: 21px;
  }

  .add-to-interests-item {
    margin-top: 15px;
    margin-right: 15px;
  }
}

.add-to-project-experience-bullet-box {
  width: 736px;
  height: 880px;

  .share-youry-profile-main {
    padding: 25px 32px 0;
  }
}

.add-to-dream-position {
  height: 660px;
  width: 736px;
}

.df {
  display: flex;
}
</style>