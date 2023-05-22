<template>
    <div class="edit">
        <!--  修改头像  -->
        <div class="edit-avatar center two-plate edit-avatar-plate "
             v-if="bulletBoxType==='avatarEdit'"
             v-loading="loading"
        >
            <div class="popup-title edit-avatar-head df pl-16px">
                <div class="edit-avatar-title">Edit</div>
                <div
                        class="popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>
                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="bulletBoxEditAvatar = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>
            <div class="edit-avatar-main z-100">
                <div class="avatar-img">
                    <YCutOut type="avatar"></YCutOut>
                </div>
            </div>
        </div>

        <!--  修改照片墙  --->
        <div class="edit-avatar center two-plate edit-avatar-plate "
             v-if="bulletBoxType==='baEdit'"
             v-loading="loading"
        >
            <div class="popup-title edit-avatar-head df pl-16px">
                <div class="edit-avatar-title">Edit</div>
                <div
                        class="popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>
                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="bulletBoxEditAvatar = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>
            <div class="edit-avatar-main z-100">
                <div class="avatar-img">
                    <YCutOut type="baEdit"></YCutOut>
                </div>
            </div>
        </div>

        <!--  修改个人信息  -->
        <div
                class="edit-in-your-profile center two-plate"
                v-else-if="bulletBoxType==='information'"
                v-loading="loading"
        >
            <div class="popup-title add-toyour-profile-head plate-border-bottom">
                <p class="edit-avatar-title">Edit in your profile</p>
                <div class="popup-closed  edit-avatar-closure cursor-pointer" @click="dialogBox = false">
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>
                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="bulletBoxeditInYourProfile = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>
            <div class="edit-form ">
                <el-form
                        ref="personaiInformationRef"
                        :model="personaiInformation"
                        :rules="rulesPersonaiInformation"
                        class="form"
                        require-asterisk-position="right"
                        label-position="top"
                >
                    <el-form-item class="h5-el-form-item">
                        <el-col :span="11">
                            <el-form-item prop="firstName" label="First Name">
                                <el-input v-model="personaiInformation.firstName"
                                          maxlength="50"
                                          placeholder="Please enter up to 50 characters"/>
                            </el-form-item>
                        </el-col>
                        <el-col class="text-center interval" :span="2">
                            <span class="text-gray-500 "></span>
                        </el-col>
                        <el-col :span="11">
                            <el-form-item prop="lastName" label="Last Name">
                                <el-input v-model="personaiInformation.lastName" maxlength="50"
                                          placeholder="Please enter up to 50 characters"/>
                            </el-form-item>
                        </el-col>
                    </el-form-item>

                    <el-form-item class="h5-el-form-item">
                        <el-col :span="11">
                            <el-form-item prop="countryName" label="Country">
                                <el-select
                                        filterable
                                        v-model="personaiInformation.countryName"
                                        placeholder="Please enter country"
                                        style="width: 100%;"
                                >
                                    <el-option
                                            v-for="item in optionsRegion"
                                            :key="item.code"
                                            :label="item.name"
                                            :value="item.name"
                                            @click="blurselect(item.num)"
                                    />
                                </el-select>
                            </el-form-item>
                        </el-col>
                        <el-col class="text-center interval" :span="2">
                            <span class="text-gray-500 "></span>
                        </el-col>
                        <el-col :span="11">
                            <el-form-item prop="city" label="City">
                                <el-input v-model="personaiInformation.city" maxlength="100"
                                          placeholder="Please enter up to 100 characters"/>
                            </el-form-item>
                        </el-col>
                    </el-form-item>

                    <el-form-item class="birthday">
                        <el-col :span="11">
                            <el-form-item prop="birthday" label="Birthday">
                                <el-date-picker
                                        v-model="personaiInformation.birthday"
                                        type="date"
                                        size="large"
                                        placeholder="Please select birthday"
                                        format="YYYY-MM-DD"
                                        value-format="YYYY-MM-DD"
                                        :clearable="false"
                                        :disabled-date="disabledDate"
                                />
                            </el-form-item>
                        </el-col>
                    </el-form-item>
                </el-form>
            </div>
            <n-space justify="end" class="tail operate-btn">
                <CVButton @click="submitUserEdit">
                    <template #main>
              <span>
              Save
              </span>
                    </template>
                </CVButton>
            </n-space>
        </div>

        <!-- 增加个人实力   -->
        <div
                class="add-toyour-profile add-to-about-me add-to-interests center two-plate "
                v-else-if="bulletBoxType==='aboutMe'"
                v-loading="loadingPersonalStrength"
        >
            <div class="popup-title add-toyour-profile-head " justify="space-between">
                <p class="edit-avatar-title">Add to About me</p>
                <div
                        class="popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>
                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="bulletBoxAddToPersonalStrength = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>

            <div class="edit-form">
                <p class="title">Content</p>
                <el-input
                        v-model="addToPersonalStrengthVal"
                        type="textarea"
                        placeholder="Please enter up to 500 characters"
                        maxlength="500"
                        show-word-limit
                        rows="16"
                        resize='none'
                />
            </div>

            <div>
                <n-space justify="end" class="tail operate-btn">
                    <CVButton @click="submitaAddToPersonalStrength">
                        <template #main>
              <span>
              Save
              </span>
                        </template>
                    </CVButton>
                </n-space>
            </div>
        </div>

        <!-- 添加到语言   -->
        <div
                class="add-toyour-profile add-to-interests center two-plate"
                v-if="bulletBoxType==='languages'"
                v-loading="loadingInterests"
        >
            <div class="popup-title add-toyour-profile-head">
                <p class="edit-avatar-title ">Add to Languages</p>

                <div class="popup-closed  edit-avatar-closure cursor-pointer" @click="dialogBox = false">
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>
                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>
            <div class="edit-form">
                <div>
                    <p class="title">Title</p>
                    <el-select
                            filterable
                            v-model="languageVal"
                            class="w-full"
                            placeholder="Please select the language you speak"
                            size="large"
                            @change="changeLanguage"
                    >
                        <el-option
                                v-for="item in getLanguagesList"
                                :key="item.name"
                                :label="item.name"
                                :value="item.name"
                        />
                    </el-select>
                    <n-space class="personal-strength-label">
                        <div
                                v-for="(item,index) in  userInfoStore.languagesList"
                                :key="index"
                                class="skills-item-btn">
                            <span>{{ item }}</span>
                            <span class="label-operation-delete">
                    <img @click="deleteLanguagesItem(index)" class="bullet-box-add-icon"
                         src="@/assets/img/delete-icon.png">
                  </span>
                        </div>
                    </n-space>
                </div>
            </div>

            <n-space justify="end" class="tail operate-btn">
                <CVButton @click="dialogBox = false">
                    <template #main>
              <span>
              Save
              </span>
                    </template>
                </CVButton>
            </n-space>
        </div>

        <!-- 添加到兴趣   -->
        <div
                class="add-toyour-profile add-to-interests center two-plate"
                v-else-if="bulletBoxType==='interests'"
                v-loading="loadingInterests"
        >
            <div class="popup-title add-toyour-profile-head">
                <p class="edit-avatar-title ">Add to Interests</p>

                <div class="popup-closed  edit-avatar-closure cursor-pointer" @click="dialogBox = false">
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>
                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>
            <div class="edit-form">
                <div>
                    <p class="title">Title</p>
                    <n-input v-model:value="addToInterests" maxlength="50" @keyup.enter="submitaddToInterests"
                             placeholder="Please enter up to 50 characters"></n-input>

                    <n-space class="personal-strength-label">
                        <div
                                v-for="(item, index) in userInfoStore.interestsItem"
                                :key="index"
                                class="skills-item-btn"
                        >
                            <span>{{ item }}</span>
                            <img @click="deleteInterestsItem(index)" class="bullet-box-add-icon ml-6px"
                                 src="@/assets/img/delete-icon.png">
                        </div>
                    </n-space>
                </div>
            </div>

            <n-space justify="end" class="tail operate-btn">
                <CVButton @click="submitaddToInterests">
                    <template #main>
              <span>
              Save
              </span>
                    </template>
                </CVButton>
            </n-space>
        </div>

        <!-- 添加到技能 Skills  -->
        <div
                class="add-toyour-profile add-to-interests center two-plate"
                v-else-if="bulletBoxType==='skills'"
                v-loading="loadingSkills"
        >
            <div class="popup-title add-toyour-profile-head">
                <p class="edit-avatar-title">Add to Skills</p>
                <div class="popup-closed edit-avatar-closure cursor-pointer" @click="dialogBox = false">
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>
                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="bulletBoxAddToSkills = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>
            <div class="edit-form">
                <div>
                    <p class="title">Title</p>
                    <n-input v-model:value="addToSkillsInput" maxlength="30" @keyup.enter="submitaAddToSkills"
                             placeholder="Please enter up to 30 characters"></n-input>
                    <n-space class="personal-strength-label">
                        <div
                                v-for="(item, index) in userInfoStore.skillItem"
                                :key="index"
                                class="skills-item-btn"
                        >
                            <div class="pop-ups-skills">{{ item }}</div>
                            <img @click="deleteSkillItemBox(index)" class="bullet-box-add-icon ml-6px"
                                 src="@/assets/img/delete-icon.png">
                        </div>
                    </n-space>
                </div>
            </div>

            <n-space justify="end" class="tail operate-btn">
                <CVButton @click="submitaAddToSkills">
                    <template #main>
              <span>
              Save
              </span>
                    </template>
                </CVButton>
            </n-space>
        </div>

        <!-- 添加到教育 Education -->
        <div
                class="edit-avatar add-to-education center two-plate "
                v-else-if="bulletBoxType==='education'"
                v-loading="loading"
        >
            <div class="popup-title edit-avatar-head df">
                <div class="edit-avatar-title">Add to Education</div>
                <div
                        class="popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>

                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="bulletBoxAddToEducation = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>

            <div class="share-youry-profile-main pr">
                <el-form
                        ref="addToEducationRef"
                        :model="addToEducation"
                        :rules="rulesbulletBoxAddToEducation"
                        require-asterisk-position="right"
                        label-position="top"
                        class="mb-36px"
                >
                    <el-form-item label="School/Organization" prop="name">
                        <el-input
                                v-model="addToEducation.name"
                                type="text"
                                maxlength="100"
                                placeholder="Please enter up to 100 characters"
                        />
                    </el-form-item>

                    <el-form-item class="h5-el-form-item">
                        <el-col :span="11">
                            <el-form-item
                                    label="Start time"
                                    prop="startTime"
                            >
                                <el-date-picker
                                        v-model="addToEducation.startTime"
                                        type="month"
                                        placeholder="Please select time"
                                        :clearable="false"
                                        format="YYYY/MM"
                                        value-format="YYYY-MM"
                                />
                            </el-form-item>
                        </el-col>
                        <el-col class="text-center interval" :span="2">
                            <span class="text-gray-500 ">to</span>
                        </el-col>
                        <el-col :span="11">
                            <el-form-item class="end-time-item" label="End time"
                                          prop="endTime"
                                          style="width: 100%"
                            >
                                <el-date-picker
                                        v-model="addToEducation.endTime"
                                        type="month"
                                        placeholder="Please select time"
                                        :clearable="false"
                                        format="YYYY/MM"
                                        value-format="YYYY-MM"
                                />
                            </el-form-item>
                        </el-col>
                    </el-form-item>

                    <el-form-item label="Content" prop="content">
                        <el-input
                                v-model="addToEducation.content"
                                type="textarea"
                                placeholder="Please enter up to 500 characters"
                                maxlength="500"
                                show-word-limit
                                rows="8"
                                resize='none'
                        />
                    </el-form-item>
                </el-form>
            </div>

            <n-space class="edit-avatar-tail tail operate-btn border-t-1 border-solid  border-gray-200" justify="end">
                <CVButton @click="submitaddToEducation">
                    <template #main>
              <span>
              Save
              </span>
                    </template>
                </CVButton>
            </n-space>
        </div>

        <!-- 添加-经验 Experience -->
        <div
                class="edit-avatar add-to-education center two-plate  min-height"
                v-else-if="bulletBoxType==='experience'"
                v-loading="loading"
                id="experience"
        >
            <div class="popup-title edit-avatar-head df">
                <div class="edit-avatar-title">Add to Experience</div>
                <div
                        class="popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>

                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>

            <div class="share-youry-profile-main pr">
                <el-form
                        ref="experienceFormDataRef"
                        :model="experienceFormData"
                        :rules="rulesExperienceFormData"
                        require-asterisk-position="right"
                        label-position="top"
                        class="mb-36px"
                >
                    <el-form-item label="Organization" prop="organization">
                        <el-input
                                v-model="experienceFormData.organization"
                                type="text"
                                maxlength="100"
                                placeholder="Please enter up to 100 characters"
                        />
                    </el-form-item>
                    <el-form-item label="Position" prop="position">
                        <el-input
                                v-model="experienceFormData.position"
                                type="text"
                                maxlength="100"
                                placeholder="Please enter up to 100 characters"
                        />
                    </el-form-item>
                    <el-form-item label="Title" prop="title">
                        <el-input
                                v-model="experienceFormData.title"
                                type="text"
                                maxlength="100"
                                placeholder="Please enter up to 100 characters"
                        />
                    </el-form-item>
                    <el-form-item label="Content" prop="content">
                        <el-input
                                v-model="experienceFormData.content"
                                type="textarea"
                                placeholder="Please enter up to 1000 characters"
                                maxlength="1000"
                                show-word-limit
                                rows="8"
                                resize='none'
                        />
                    </el-form-item>
                    <el-form-item label="Summary" prop="summary">
                        <el-input
                                v-model="experienceFormData.summary"
                                type="textarea"
                                placeholder="Please enter up to 200 characters"
                                maxlength="200"
                                show-word-limit
                                rows="8"
                                resize='none'
                        />
                    </el-form-item>

                    <el-form-item class="h5-el-form-item">
                        <el-col :span="11">
                            <el-form-item
                                    label="Start time"
                                    prop="startTime"
                            >
                                <el-date-picker
                                        v-model="experienceFormData.startTime"
                                        type="month"
                                        placeholder="Please select time"
                                        :clearable="false"
                                        format="YYYY/MM"
                                        value-format="YYYY-MM"
                                />
                            </el-form-item>
                        </el-col>
                        <el-col class="text-center interval" :span="2">
                            <span class="text-gray-500 ">to</span>
                        </el-col>
                        <el-col :span="11">
                            <el-form-item class="end-time-item" label="End time"
                                          prop="endTime"
                                          style="width: 100%"
                            >
                                <el-date-picker
                                        v-model="experienceFormData.endTime"
                                        type="month"
                                        placeholder="Please select time"
                                        :clearable="false"
                                        format="YYYY/MM"
                                        value-format="YYYY-MM"
                                />
                            </el-form-item>
                        </el-col>
                    </el-form-item>


                </el-form>
            </div>

            <n-space class="edit-avatar-tail tail operate-btn border-t-1 border-solid  border-gray-200" justify="end">
                <CVButton @click="submitaddToxperience">
                    <template #main>
              <span>
              Save
              </span>
                    </template>
                </CVButton>
            </n-space>
        </div>

        <!-- 添加到-工作经验  -->
        <div
                class="edit-avatar add-to-Job-experience center two-plate "
                v-else-if="bulletBoxType==='jobExperience'"
                v-loading="loading"
        >
            <div class="popup-title edit-avatar-head df ">
                <div class="edit-avatar-title">Add to Job Experience</div>
                <div
                        class="popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>
                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="bulletBoxAddToJobExperience = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>

            <div class="share-youry-profile-main pr">
                <el-form
                        ref="addToJobExperienceFormRef"
                        label-position="top"
                        :model="addToJobExperienceForm"
                        :rules="rulesAddToJobExperience"
                        require-asterisk-position="right"
                >
                    <el-form-item label="Company Name" prop="companyName">
                        <el-input
                                v-model="addToJobExperienceForm.companyName"
                                type="text"
                                maxlength="200"
                                placeholder="Please enter 2-200 characters"
                        />
                    </el-form-item>

                    <el-form-item label="Company Industry" class="pr" style="width: 100%">
                        <el-input
                                v-model="addToJobExperienceForm.companyIndustry"
                                type="text"
                                maxlength="200"
                                placeholder="Please enter 1-200 characters"
                        />
                    </el-form-item>

                    <el-form-item label="Job Position" prop="jobPosition" class="pr" style="width: 100%">
                        <el-input
                                v-model="addToJobExperienceForm.jobPosition"
                                type="text"
                                maxlength="100"
                                placeholder="Please enter 1-100 characters"
                        />
                    </el-form-item>

                    <el-form-item label="Job Content" prop="jobContent">
                        <el-input
                                v-model="addToJobExperienceForm.jobContent"
                                type="textarea"
                                placeholder="Please enter 10-1000 characters"
                                maxlength="1000"
                                show-word-limit
                                rows="4"
                                resize='none'
                        />
                    </el-form-item>

                    <el-form-item class="h5-el-form-item el-form-item-last">
                        <el-col :span="11">
                            <el-form-item
                                    label="Start time"
                                    prop="startTime"
                                    class="pr" style="width: 100%"
                            >
                                <el-date-picker
                                        style="width: 100%"
                                        v-model="addToJobExperienceForm.startTime"
                                        type="month"
                                        placeholder="Please select time"
                                        :clearable="false"
                                        format="YYYY/MM"
                                        value-format="YYYY-MM"
                                />
                            </el-form-item>
                        </el-col>

                        <el-col class="text-center interval" :span="2">
                            <span class="text-gray-500">to</span>
                        </el-col>
                        <el-col :span="11">
                            <el-form-item label="End time" class="pr"
                                          prop="endTime"
                                          style="width: 100%">
                                <el-date-picker
                                        style="width: 100%"
                                        v-model="addToJobExperienceForm.endTime"
                                        type="month"
                                        placeholder="Please select time"
                                        :clearable="false"
                                        format="YYYY/MM"
                                        value-format="YYYY-MM"
                                        :disabled-date="changePicker"
                                />
                            </el-form-item>
                        </el-col>
                    </el-form-item>

                </el-form>
            </div>

            <n-space class="edit-avatar-tail tail operate-btn border-t-1" justify="end">
                <CVButton @click="submitAddToJobExperience">
                    <template #main>
              <span>
              Save
              </span>
                    </template>
                </CVButton>
            </n-space>
        </div>

        <!-- 添加到-项目经验  -->
        <div
                class="edit-avatar add-to-project-experience center two-plate"
                id=" ProjectExperience"
                v-else-if="bulletBoxType==='projectExperience'"
                v-loading="loading"
        >
            <div class="popup-title edit-avatar-head df ">
                <div class="edit-avatar-title">Add to Project experience</div>

                <div
                        class="popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>

                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>

            <div class="share-youry-profile-main pr">
                <el-form
                        ref="addToprojectExperienceRef"
                        label-position="top"
                        :model="addToprojectExperience"
                        :rules="rulesAddToprojectExperience"
                        require-asterisk-position="right"
                >
                    <el-form-item label="Project name" prop="projectName">
                        <el-input
                                v-model="addToprojectExperience.projectName"
                                type="text"
                                maxlength="100"
                                placeholder="Please enter up to 100 characters"
                        />
                    </el-form-item>

                    <el-form-item label="Project content" prop="projectContent">
                        <el-input
                                v-model="addToprojectExperience.projectContent"
                                type="textarea"
                                placeholder="Please enter up to 1000 characters"
                                maxlength="1000" show-word-limit
                                :rows="4.5"
                                resize='none'
                        />
                    </el-form-item>

                    <el-form-item label="Job position">
                        <el-input
                                v-model="addToprojectExperience.jobPosition"
                                type="text"
                                maxlength="100"
                                placeholder="Please enter up to 100 characters"
                        />
                    </el-form-item>

                    <el-form-item label="Job content">
                        <el-input
                                v-model="addToprojectExperience.jobContent"
                                type="textarea"
                                maxlength="1000" show-word-limit
                                :rows="4.5"
                                placeholder="Please enter up to 1000 characters"
                                resize='none'
                        />
                    </el-form-item>

                    <el-form-item label="Have skills" prop="projectSkills">
                        <el-select
                                style="width: 100%;"
                                v-model="addToprojectExperience.projectSkills"
                                multiple
                                filterable
                                allow-create
                                default-first-option
                                :reserve-keyword="false"
                                placeholder="Please enter up to 30 characters"
                        >
                            <el-option
                                    v-for="item in skillItemBulletBox"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value"
                            />
                        </el-select>
                    </el-form-item>

                    <el-form-item class="h5-el-form-item ">
                        <el-col :span="11">
                            <el-form-item prop="startTime" label="Start time">
                                <el-date-picker
                                        v-model="addToprojectExperience.startTime"
                                        type="month"
                                        placeholder="Please select time"
                                        :clearable="false"
                                        format="YYYY/MM"
                                        value-format="YYYY-MM"
                                />
                            </el-form-item>
                        </el-col>
                        <el-col class="text-center interval" :span="2">
                            <span class="text-gray-500">to</span>
                        </el-col>
                        <el-col :span="11">
                            <el-form-item prop="endTime" label="End time">
                                <el-date-picker
                                        v-model="addToprojectExperience.endTime"
                                        type="month"
                                        placeholder="Please select time"
                                        :clearable="false"
                                        format="YYYY/MM"
                                        value-format="YYYY-MM"
                                />
                            </el-form-item>
                        </el-col>
                    </el-form-item>

                    <el-form-item class="h5-el-form-item ">
                        <el-col :span="11">
                            <el-form-item label="Project Links Platform">
                                <el-input
                                        v-model="addToprojectExperience.projectLinkPlatform"
                                        type="text"
                                        placeholder="Please enter up to 30 characters"
                                        maxlength="30"
                                />
                            </el-form-item>
                        </el-col>
                        <el-col class="text-center interval" :span="2">
                            <span class="text-gray-500"></span>
                        </el-col>
                        <el-col :span="11">
                            <el-form-item label="Project link">
                                <el-input
                                        v-model="addToprojectExperience.projectLink"
                                        type="text"
                                        maxlength="500"
                                        placeholder="Please enter up to 500 characters"
                                />
                            </el-form-item>
                        </el-col>
                    </el-form-item>
                </el-form>
            </div>

            <n-space class="edit-avatar-tail tail operate-btn border-t-1" justify="end">
                <CVButton @click="submitAddToProjectExperience">
                    <template #main>
              <span>
              Save
              </span>
                    </template>
                </CVButton>
            </n-space>
        </div>

        <!-- 志愿者经历-->
        <div
                class="edit-avatar add-to-education center two-plate "
                v-else-if="bulletBoxType==='volunteerExperience'"
                v-loading="loading"
        >
            <div class="popup-title edit-avatar-head df">
                <div class="edit-avatar-title">Add to Volunteer Experience</div>
                <div
                        class="popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>

                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="bulletBoxAddToEducation = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>

            <div class="share-youry-profile-main pr">
                <el-form
                        ref="volunteerExperienceFormRef"
                        :model="volunteerExperienceForm"
                        :rules="rulesVolunteerExperienceForm"
                        require-asterisk-position="right"
                        label-position="top"
                        class="mb-36px"
                        @submit.prevent
                >
                    <el-form-item label="Title" prop="title">
                        <el-input
                                v-model="volunteerExperienceForm.title"
                                type="text"
                                maxlength="100"
                                placeholder="Please enter up to 100 characters"
                        />
                    </el-form-item>

                    <el-form-item label="Content" prop="content">
                        <el-input
                                v-model="volunteerExperienceForm.content"
                                type="textarea"
                                placeholder="Please enter up to 1000 characters"
                                maxlength="1000"
                                show-word-limit
                                rows="8"
                                resize='none'
                        />
                    </el-form-item>

                    <el-form-item class="h5-el-form-item">
                        <el-col :span="11">
                            <el-form-item
                                    label="Start time"
                                    prop="startTime"
                            >
                                <el-date-picker
                                        v-model="volunteerExperienceForm.startTime"
                                        type="month"
                                        placeholder="Please select time"
                                        :clearable="false"
                                        format="YYYY/MM"
                                        value-format="YYYY-MM"
                                />
                            </el-form-item>
                        </el-col>
                        <el-col class="text-center interval" :span="2">
                            <span class="text-gray-500 ">to</span>
                        </el-col>
                        <el-col :span="11">
                            <el-form-item class="end-time-item"
                                          prop="entTime"
                                          label="End time"
                                          style="width: 100%"
                            >
                                <el-date-picker
                                        v-model="volunteerExperienceForm.endTime"
                                        type="month"
                                        placeholder="Please select time"
                                        :clearable="false"
                                        format="YYYY/MM"
                                        value-format="YYYY-MM"
                                />
                            </el-form-item>
                        </el-col>
                    </el-form-item>
                </el-form>
            </div>
            <n-space class="edit-avatar-tail  tail operate-btn border-t-1 border-solid  border-gray-200" justify="end">
                <CVButton @click="submitaVolunteerExperience(volunteerExperienceFormRef,1)">
                    <template #main>
              <span>
              Save
              </span>
                    </template>
                </CVButton>
            </n-space>
        </div>

        <!-- 添加到自定义属性  -->
        <div
                class="edit-avatar center two-plate add-to-custom-item" id="customItem"
                v-else-if="bulletBoxType==='customItem'"
                v-loading="loading"
        >
            <div class="popup-title add-toyour-profile-head df">
                <div class="edit-avatar-title">Add to Custom item</div>
                <div
                        class="popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="add-img" src="@/assets/img/delete-img.png">
                </div>

                <div
                        class="h5-popup-closed edit-avatar-closure cursor-pointer"
                        @click="dialogBox = false"
                >
                    <img class="" src="@/assets/img/h5-delete-img.png">
                </div>
            </div>
            <div class="share-youry-profile-main pr">
                <el-form
                        ref="customAttributesRef"
                        :model="addToCustomAttributes"
                        :rules="rulesAddToCustomAttributes"
                        require-asterisk-position="right"
                        label-position="top"
                >
                    <el-form-item label="Title" prop="title">
                        <el-input
                                v-model="addToCustomAttributes.title"
                                type="text"
                                placeholder="Please enter 3-30 characters"
                                maxlength="30"
                        />
                    </el-form-item>

                    <el-form-item
                            first
                            prop="content"
                            label="Content"
                    >
                        <el-input
                                v-model="addToCustomAttributes.content"
                                type="textarea"
                                maxlength="1000" show-word-limit
                                :rows="6"
                                resize='none'
                                placeholder="Please enter 5-1000 characters"
                        />
                    </el-form-item>

                    <el-form-item label="Summary">
                        <el-input
                                v-model="addToCustomAttributes.summary"
                                type="textarea"
                                maxlength="200" show-word-limit
                                :rows="6"
                                resize='none'
                                placeholder="Please enter up to 200 characters"
                        />
                    </el-form-item>

                    <el-form-item class="h5-el-form-item">
                        <el-col :span="11">
                            <el-form-item
                                    label="Start time"
                                    prop="startTime"
                            >
                                <el-date-picker
                                        v-model="addToCustomAttributes.startTime"
                                        type="month"
                                        placeholder="Please select time"
                                        :clearable="false"
                                        format="YYYY/MM"
                                        value-format="YYYY-MM"
                                />
                            </el-form-item>
                        </el-col>
                        <el-col class="text-center interval" :span="2">
                            <span class="text-gray-500 ">to</span>
                        </el-col>
                        <el-col :span="11">
                            <el-form-item class="end-time-item"
                                          prop="endTime"
                                          label="End time"
                                          style="width: 100%"
                            >
                                <el-date-picker
                                        v-model="addToCustomAttributes.endTime"
                                        type="month"
                                        placeholder="Please select time"
                                        :clearable="false"
                                        format="YYYY/MM"
                                        value-format="YYYY-MM"
                                />
                            </el-form-item>
                        </el-col>
                    </el-form-item>
                </el-form>
            </div>

            <n-space class="edit-avatar-tail tail operate-btn border-t-1" justify="end">
                <CVButton @click="submitAddToCustomAttributes">
                    <template #main>
              <span>
              Save
              </span>
                    </template>
                </CVButton>
            </n-space>
        </div>
    </div>
</template>

<script setup>
import {NInput, NSpace, useMessage} from "naive-ui";
import YCutOut from "@/components/personal-wditor/cut-out.vue";
import {useState} from "nuxt/app";
import {
    apiCommonLanguage,
    apiCustomAttributeGetId,
    apiCustomAttributeList,
    apiExperienceAdd,
    apiExperienceGetId,
    apiExperienceList,
    apiUpdateBase,
    apiUpdateInterests,
    apiUpdateLanguages,
    apiUpdatePersonalStrength,
    apiUpdateSkills,
    apiUserCompanyAdd,
    apiUserCompanyGet,
    apiUserCompanyList,
    apiUserEducationAdd,
    apiUserEducationGet,
    apiUserEducationList,
    apiUserProjectAdd,
    apiUserProjectGet,
    apiUserProjectList,
    apiVolunteerAdd,
    apiVolunteerGetId,
    apiVolunteerList,
    userCustomAttributeChange
} from "~/api/personal-wditor";
import {useI18n} from "vue-i18n";
import {ElMessage, ElMessageBox} from 'element-plus'
import {apiCommonCountry} from "~/api/login";
import {calculateHowOld} from "~/utils/util";
import CVButton from '@/components/form/button.vue';
import {storeToRefs} from 'pinia'
import {userStore} from "@/composables/store/user.js";

let userInfoStore = userStore()

let message = useMessage()
const {t, locale} = useI18n();
const emit = defineEmits(["ChangeGetId", 'changeUserInfo']);

const props = defineProps({
    getId: String,
    personaiInformation: Object,
})

let personaiInformation = ref({})
personaiInformation.value = JSON.parse(JSON.stringify(userInfoStore.personaiInformation))


onMounted(() => {
    document.body.style.cssText = 'overflow-y:hidden';
})
onUnmounted(() => {
    document.body.style.cssText = 'overflow-y:auto';
})
const filterTime = (aimFormData, callback, form, type = '') => {
    if (aimFormData.value.startTime && aimFormData.value.endTime) {
        if (new Date(aimFormData.value.startTime).getTime() > new Date(aimFormData.value.endTime).getTime()) {
            callback(new Error(""));
        } else {
            if (type == 'startTime') {
                form.value.validateField('endTime')
            }
            if (type == 'endTime') {
                // form.value.validateField('startTime')
            }
            callback();
        }
    } else {
        callback();
    }
}
/**
 * 判断时间选择
 */
const selectTimeChangeEducation = (rule, value, callback, source, options) => {
    console.log(rule, value, source, options)
    return filterTime(addToEducation, callback, addToEducationRef, Object.keys(source)[0])

}
const selectTimeChangeExperience = (rule, value, callback, source, options) => {
    return filterTime(experienceFormData, callback, experienceFormDataRef, Object.keys(source)[0])

}
const selectTimeChangeJobExperience = (rule, value, callback, source, options) => {
    return filterTime(addToJobExperienceForm, callback, addToJobExperienceFormRef, Object.keys(source)[0])

}
const selectTimeChangeprojectExperience = (rule, value, callback, source, options) => {
    return filterTime(addToprojectExperience, callback, addToprojectExperienceRef, Object.keys(source)[0])

}
const selectTimeChangeVolunteerExperience = (rule, value, callback, source, options) => {
    return filterTime(volunteerExperienceForm, callback, volunteerExperienceFormRef, Object.keys(source)[0])

}
const selectTimeChangeCustom = (rule, value, callback, source, options) => {
    return filterTime(addToCustomAttributes, callback, customAttributesRef, Object.keys(source)[0])

}

// 弹框80% 最小高度866

const changeVal = (e) => {
    e.preventDefault()
}

// 添加到教育-表单
let {addToEducationForm: addToEducation} = storeToRefs(userInfoStore);
// 教育-根据ID获取数据
const getEducation = () => {
    apiUserEducationGet({id: props.getId}).then(res => {
        loading.value = false
        if (res.data && res.data.id) {
            addToEducation.value = res.data
        }
    }).catch(err => {
        loading.value = false
    })
}
/**
 * 禁止日期选择
 */
const disabledDate = (date) => {
    return date && date.valueOf() > Date.now() || ((Date.now() - date.valueOf()) / (1000 * 60 * 60 * 24 * 365) > 120);
}
// const changePickerEndTime = (date) => {
//   return date && date.valueOf()
// }
let dialogBox = useState('_dialogBox')
/**
 * 弹框
 * param {*} 根据传的字段显示对应的弹框
 */
let bulletBoxType = computed(() => {
    return useState('_dialogBoxType').value
})
/**
 * 弹框是什么操作
 * param {*} add 添加操作
 * param {*} edit 编辑操作
 */
let formActionType = useState('_formActionType', () => null)
if (formActionType.value === 'add') {
    emit("ChangeGetId", 0);
}
// 加载中
let loading = useState('_dialog')
// 国家-下拉数据
const optionsRegion = ref([]);
// 语言-下拉数据
let getLanguagesList = ref([])
// 语言-弹框-值
let languageVal = ref(null);
/**
 * 判断是那个弹框
 * @param {dialogBoxType}  Languages 添加语言
 * @param {dialogBoxType}  等等
 */
if (bulletBoxType.value === 'languages') {
    // 获取语言
    apiCommonLanguage().then(res => {
        getLanguagesList.value = JSON.parse(JSON.stringify(res.data.list))
    })
} else if (bulletBoxType.value === 'information') {
    // 国家列表
    apiCommonCountry().then(res => {
        if (res.data.list) {
            optionsRegion.value = JSON.parse(JSON.stringify(res.data.list))
        }
    })
} else if (bulletBoxType.value === 'education') {
    addToEducation.value = {orderBy: 0}
}

let bulletBoxEditDom = useState("_bulletBoxEditDom", () => false);

let loadingImgBulletBoxloading = useState("_loadingImgBulletBoxloading", () => false)
// 编辑头像-弹框
let bulletBoxEditAvatar = useState("_bulletBoxEditAvatar");
// 在你的个人资料中编辑-弹框
let bulletBoxeditInYourProfile = useState("_bulletBoxeditInYourProfile");

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

const blurselect = (val) => {
    personaiInformation.value.countryId = Number(val)
}

let personaiInformationRef = ref()
let rulesPersonaiInformation = ref({
    firstName: [{required: true, message: 'Please enter up to 50 characters', trigger: 'blur'},],
    lastName: [{required: true, message: 'Please enter up to 50 characters', trigger: 'blur'},],
    countryName: [{required: true, message: 'Please enter country', trigger: 'blur'},],
    city: [{required: true, message: 'Please enter up to 100 characters', trigger: 'blur'},],
    birthday: [{required: true, message: 'Please select birthday', trigger: 'blur'},],
})
// 用户信息-弹框-提交修改
const submitUserEdit = (e) => {
    personaiInformationRef.value.validate((valid, fields) => {
        if (valid) {
            let params = personaiInformation
            loading.value = true

            apiUpdateBase(params).then((res) => {
                dialogBox.value = false
                loading.value = false
                if (res.data && res.data.status) {
                    let data = JSON.parse(JSON.stringify(params))
                    userInfoStore.personaiInformation = data._value
                    userInfoStore.ageInformation = calculateHowOld(params.value.birthday)
                    // useState('_ageInformation').value = calculateHowOld(params.value.birthday)
                } else {
                    loading.value = false
                    return
                }
            }).catch(err => {
                loading.value = false
                return
            });
        }
    })
}

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


// 语言-添加
const changeLanguage = (val) => {
    if (!userInfoStore.languagesList.includes(val)) {
        userInfoStore.languagesList.push(val)
        let data = ''
        userInfoStore.languagesList.forEach((item, index) => {
            if (index) {
                data += `,${item}`
            } else {
                data += item
            }
        })
        apiUpdateLanguages({languages: data}).then(res => {
            // console.log(res)
        })
    }
}
// 语言-删除
const deleteLanguagesItem = (index) => {
    if (userInfoStore.languagesList.length === 1) {
        ElMessage({
            type: 'error',
            message: t('personal-wditor.Interest needs to keep one'),
        })
        return
    }
    ElMessageBox({
        title: 'Delete',
        confirmButtonText: 'Delete',
        roundButton: "true",
        message: h('p', null, [
            h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                'Are you sure you want to delete the current item?'),
        ]),
    }).then(() => {
        userInfoStore.languagesList.splice(index, 1)
        let data = ''
        userInfoStore.languagesList.forEach((item, index) => {
            if (index) {
                data += `,${item}`
            } else {
                data += item
            }
        })
        apiUpdateLanguages({languages: data}).then(res => {
            // 语言-删除-请求
        })
    })
}

// 背景编辑-弹框
let bulletBoBackgroundEditorr = useState("_bulletBoBackgroundEditorr");
// 添加到自定义属性-弹框
let bulletBoxAddToCustomAttributes = useState(
    "_bulletBoxAddToCustomAttributes"
);
// 添加到自定义属性-表单
let {addToCustomAttributes} = storeToRefs(userInfoStore);
// 添加到自定义属性-表单-验证规则
let rulesAddToCustomAttributes = ref({
    title: [
        {required: true, min: 3, max: 30, message: 'Please enter 3-30 characters', trigger: 'blur'},
    ],
    content: [
        {required: true, min: 5, max: 1000, message: 'Please enter 5-1000 characters', trigger: 'blur'},
    ],
    startTime: [
        {
            required: true,
            message: 'Please select time',
            trigger: 'blur'
        }, {message: 'End time must be greater than start time', trigger: 'blur', validator: selectTimeChangeCustom}],
    endTime: {message: 'End time must be greater than start time', trigger: 'blur', validator: selectTimeChangeCustom},
});
let customAttributesRef = ref()
let loadingCustomAttributes = ref(false)
// 添加到自定义属性-操作
const submitAddToCustomAttributes = () => {
    customAttributesRef.value.validate(valid => {
        if (valid) {
            let params = addToCustomAttributes.value

            for (const dataKey in params) {
                if (params[dataKey] === '') delete params[dataKey]
            }
            if (props.getId) params.id = props.getId
            else params.id = 0
            params.orderBy = 0

            loading.value = true

            userCustomAttributeChange(params).then((res) => {
                loading.value = false
                let {code, data} = JSON.parse(JSON.stringify(res))
                let {id} = data
                dialogBox.value = false
                if (id) {
                    useState("_reacquireData").value = 1
                    apiCustomAttributeList().then(res => {
                        if (res.data.list) {
                            userInfoStore.customItemList = res.data.list
                        }
                    })
                }
            }).catch(err => {
                loadingCustomAttributes.value = false;
                loading.value = false
            });
        }
    })

};
let frameType = useState("_frameType")

// 公司（工作经验）-获取数据
const getJobExperience = () => {
    apiUserCompanyGet({id: props.getId}).then(res => {
        loading.value = false
        if (res.data && res.data.id) {
            let data = JSON.parse(JSON.stringify(res.data))
            addToJobExperienceForm.value = data
            useState("_frameType").value = ''
        }
    }).catch(err => {
        loading.value = false
    })
}

// 项目经验-获取数据
const getProjectExperienceList = () => {
    apiUserProjectGet({id: props.getId}).then(res => {
        loading.value = false
        if (res.data && res.data.id) {
            let data = JSON.parse(JSON.stringify(res.data))
            data.projectSkills = data.projectSkills.split(',')
            addToprojectExperience.value = data
            useState("_frameType").value = ''
        }
    }).catch(err => {
        loading.value = false
    })
}

// 经验-根据ID获取数据
const getApiExperienceGetId = () => {
    apiExperienceGetId({id: props.getId}).then(res => {
        loading.value = false
        if (res.data && res.data.id) {
            let data = JSON.parse(JSON.stringify(res.data))
            experienceFormData.value = data
            useState("_frameType").value = null
        }
    }).catch(err => {
        loading.value = false
    })
}

// 志愿者经历-根据id获取数据
const getApiVolunteerGetId = () => {
    apiVolunteerGetId({id: props.getId}).then(res => {
        loading.value = false
        if (res.data && res.data.id) {
            let data = JSON.parse(JSON.stringify(res.data))
            volunteerExperienceForm.value = data
            useState("_frameType").value = null
        }
    }).catch(err => {
        loading.value = false
    })
}
// 自定义属性-根据id获取数据
const getApiCustomAttributeGetId = () => {
    apiCustomAttributeGetId({id: props.getId}).then(res => {
        loading.value = false
        if (res.data && res.data.id) {
            let data = JSON.parse(JSON.stringify(res.data))
            addToCustomAttributes.value = data
            useState("_frameType").value = null
        }
    }).catch(err => {
        loading.value = false
    })
}

// 判断传入的弹框类型，请求对应的数据
watch(() => props.getId, (newVal) => {
        if (bulletBoxType.value === 'education') {
            if (formActionType.value === 'edit') {
                getEducation()
            }
        } else if (bulletBoxType.value === 'experience') {
            if (formActionType.value === 'edit') {
                getApiExperienceGetId()
            }
        } else if (bulletBoxType.value === 'jobExperience') {
            if (formActionType.value === 'edit') {
                getJobExperience()
            }
        } else if (bulletBoxType.value === 'projectExperience') {
            if (formActionType.value === 'edit') {
                getProjectExperienceList()
            }
        } else if (bulletBoxType.value === 'volunteerExperience') {
            if (formActionType.value === 'edit') {
                getApiVolunteerGetId()
            }
        } else if (bulletBoxType.value === 'customItem') {
            if (formActionType.value === 'edit') {
                getApiCustomAttributeGetId()
            }
        }
    },
    {immediate: true, deep: true}
);

// 添加到教育-弹框
let bulletBoxAddToEducation = useState("_bulletBoxAddToEducation");

// 添加到教育-表单-教育-表单-验证规则
let rulesbulletBoxAddToEducation = ref({
    name: [
        {required: true, message: 'Please enter up to 100 characters', trigger: 'blur'},
        {required: true, message: 'Please enter up to 100 characters', trigger: 'change'},
    ],
    startTime: [{
        required: true,
        message: 'Please select time',
        trigger: 'blur'
    }, {message: 'End time must be greater than start time', trigger: 'blur', validator: selectTimeChangeEducation}],
    endTime: {
        message: 'End time must be greater than start time',
        trigger: 'blur',
        validator: selectTimeChangeEducation
    },
    content: [{required: true, message: 'Please enter up to 500 characters', trigger: 'blur'},],
})

let addToEducationRef = ref()
// 添加到教育-操作
const submitaddToEducation = () => {

    addToEducationRef.value.validate((valid, fields) => {
        if (valid) {
            loading.value = true
            let data = {
                id: addToEducation.value.id || 0,
                name: addToEducation.value.name,
                startTime: addToEducation.value.startTime,
                endTime: addToEducation.value.endTime,
                content: addToEducation.value.content,
            }

            for (const dataKey in data) {
                if (data[dataKey] === 'undefined' || data[dataKey] === '') delete data[dataKey]
            }
            apiUserEducationAdd(data).then(res => {
                loading.value = false
                // console.log('res-添加到教育', res)
                if (res.data && res.data.id) {
                    if (data.id) {
                        data.id = res.data.id
                        apiUserEducationList().then(res => {
                            // console.log('获取教育数据', res)
                            if (res.data && res.data.list.length) {
                                userInfoStore.educationList = res.data.list
                                dialogBox.value = false
                            }
                        })
                    } else {
                        addToEducation.value = {orderBy: 0}
                        data.id = res.data.id
                        userInfoStore.educationList.push(data)
                        dialogBox.value = false
                    }
                } else {
                    loading.value = false
                }
            }).catch(err => {
                loading.value = false
            })
        }
    })
};

let experienceFormDataRef = ref()
// 经验-表单数据
let {experienceFormData} = storeToRefs(userInfoStore);
// 经验-表单-验证规则
let rulesExperienceFormData = ref({
    organization: [{required: true, message: 'Please enter up to 100 characters', trigger: 'blur'},],
    position: [{required: true, message: 'Please enter up to 100 characters', trigger: 'blur'},],
    title: [{required: true, message: 'Please enter up to 100 characters', trigger: 'blur'},],
    content: [{required: true, message: 'Please enter up to 1000 characters', trigger: 'blur'},],
    startTime: [{required: true, message: 'Please select time', trigger: 'blur'},
        {message: 'End time must be greater than start time', trigger: 'blur', validator: selectTimeChangeExperience}],
    endTime: {
        message: 'End time must be greater than start time',
        trigger: 'blur',
        validator: selectTimeChangeExperience
    }
})
// 经验-添加-操作
const submitaddToxperience = () => {

    experienceFormDataRef.value.validate((valid, fields) => {
        console.log(valid)
        if (valid) {
            loading.value = true
            let data = {}
            data = JSON.parse(JSON.stringify(experienceFormData.value))
            if (!data.id) {
                data.id = 0
            }

            for (const dataKey in data) {
                if (data[dataKey] === '' || data[dataKey] === null) delete data[dataKey]
            }

            apiExperienceAdd(data).then(res => {
                loading.value = false
                if (res.data && res.data.id) {
                    dialogBox.value = false
                    apiExperienceList().then(res => {
                        if (res.data && res.data.list.length) {
                            userInfoStore.experienceList = res.data.list
                        }
                    })
                } else {
                    loading.value = false
                }
            }).catch(err => {
                loading.value = false
            })
        }
    })
};

let addIdealGoalVal = ref();
let idealGoalList = useState('_idealGoalList');

let loadingIdealGoal = ref(false)

let addPurposeCityVal = ref();
let loadingPurposeCity = ref(false);
let addPurposeCityList = useState('_purposeCityList');

let addToInterests = ref();

let loadingInterests = ref(false)
// 添加到兴趣-操作
const submitaddToInterests = () => {
    // dialogBox.value = false
    if (addToInterests.value) {
        userInfoStore.interestsItem.push(addToInterests.value);
        let data = ''
        userInfoStore.interestsItem.forEach((item, index) => {
            if (index) {
                data += `,${item}`
            } else {
                data += item
            }
        })

        apiUpdateInterests({interests: data}).then(res => {
            if (res.data && res.data.status) {
                addToInterests.value = null;
            }
        })
    }
};

// 删除-兴趣
const deleteInterestsItem = (index) => {
    if (userInfoStore.interestsItem.length === 1) {
        ElMessage({
            type: 'error',
            message: t('personal-wditor.At least keep one'),
        })
        return
    }

    ElMessageBox({
        title: 'Delete',
        confirmButtonText: 'Delete',
        roundButton: "true",
        message: h('p', null, [
            h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                'Are you sure you want to delete the current item?'),
        ]),
    }).then(() => {
        userInfoStore.interestsItem.splice(index, 1)
        let data = ''
        userInfoStore.interestsItem.forEach((item, index) => {
            if (index) {
                data += `,${item}`
            } else {
                data += item
            }
        })
        apiUpdateInterests({interests: data}).then(res => {
        })
    })
}

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
    // let data = addToDreamPosition.value;
    // console.log("添加到梦想位置", data);
};

let bulletBoxAddToSkills = useState("_bulletBoxAddToSkills");
// 添加到技能-规则
let addToSkillsInput = ref();

let loadingSkills = ref(false)
// 添加到技能-操作
const submitaAddToSkills = () => {
    if (addToSkillsInput.value) {
        userInfoStore.skillItem.push(addToSkillsInput.value);
        let data = ''
        userInfoStore.skillItem.forEach((item, index) => {
            if (index) {
                data += `,${item}`
            } else {
                data += item
            }
        })

        apiUpdateSkills({personalSkills: data}).then(res => {
            if (res.data && res.data.status) {
                addToSkillsInput.value = null;
            }
        })
    }
};
// 技能-删除
const deleteSkillItemBox = (index) => {
    if (userInfoStore.skillItem.length === 1) {
        ElMessage({
            type: 'error',
            message: t('personal-wditor.skills need to keep one'),
        })
        return
    }

    ElMessageBox({
        title: 'Delete',
        confirmButtonText: 'Delete',
        roundButton: "true",
        message: h('p', null, [
            h('div', {style: 'border-top: 1px solid #e9e9e9;width: calc(100% + 30px); padding-left: 15px;position: relative;right: 15px;'},
                'Are you sure you want to delete the current item?'),
        ]),
    }).then(() => {
        userInfoStore.skillItem.splice(index, 1)
        let data = ''
        userInfoStore.skillItem.forEach((item, index) => {
            if (index) {
                data += `,${item}`
            } else {
                data += item
            }
        })
        apiUpdateSkills({personalSkills: data}).then(res => {
        })
    })
}

// 增加工作经验-弹框
let bulletBoxAddToJobExperience = useState("_bulletBoxAddToJobExperience");
// 公司（工作经验）-表单数据

let {addToJobExperienceForm} = storeToRefs(userInfoStore);

let changePicker = ref()
changePicker.value = (date) => {
    return date && date.valueOf() > addToJobExperienceForm.value?.startTime;
}
let startTime = addToJobExperienceForm.value?.startTime
let endTime = addToJobExperienceForm.value?.endTime
watch(() => startTime, () => {
    console.log(startTime)
    changePicker.value = endTime && endTime.valueOf() > addToJobExperienceForm.value?.startTime;
})
const checkJobExperience = (rule, value, callback) => {
    if (typeof value == 'undefined' || value < 10) {
        callback(new Error('Please enter 2-100 characters'))
    } else {
        callback()
    }
}

// 工作经验-规则
let rulesAddToJobExperience = ref({
    companyName: [
        {required: true, min: 2, max: 200, message: 'Please enter 2-200 characters', trigger: 'blur'},
    ],
    jobPosition: [
        {required: true, min: 1, max: 100, message: 'Please enter 1-100 characters', trigger: 'blur'},
    ],
    startTime: [
        {
            required: true,
            message: 'Please select time',
            trigger: 'blur'
        }, {
            message: 'End time must be greater than start time',
            trigger: 'blur',
            validator: selectTimeChangeJobExperience
        }],
    endTime: {
        message: 'End time must be greater than start time',
        trigger: 'blur',
        validator: selectTimeChangeJobExperience
    },
    jobContent: [
        {required: true, min: 10, max: 1000, message: 'Please enter 10-1000 characters', trigger: 'blur'},
    ],
});
// 技能选择
let skillItemBulletBox = useState('_skillItemBulletBox')

let addToJobExperienceFormRef = ref()
let loadingJobExperience = ref(false)
// 公司（工作经验）-提交
const submitAddToJobExperience = () => {
    addToJobExperienceFormRef.value.validate((valid, fields) => {
        if (valid) {
            loading.value = true
            let data = addToJobExperienceForm.value;

            for (const dataKey in data) {
                if (data[dataKey] === '') delete data[dataKey]
            }
            if (props.getId) data.id = props.getId
            else data.id = 0
            data.orderBy = 0
            apiUserCompanyAdd(data).then(res => {
                loading.value = false
                if (res.data && res.data.id) {
                    apiUserCompanyList().then(res => {
                        loading.value = false
                        if (res.data && res.data.list.length) {
                            let data = JSON.parse(JSON.stringify(res.data.list))
                            dialogBox.value = false
                            userInfoStore.companyList = data
                            emit("ChangeGetId", 0);
                        }
                    }).catch(err => {
                        loading.value = false
                    })
                }
            }).finally(() => {
                loading.value = false
            })
        }
    })

};


let {addToprojectExperience} = storeToRefs(userInfoStore);
// 项目经验-规则
let rulesAddToprojectExperience = ref(
    {
        projectName: [
            {required: true, min: 1, max: 100, message: 'Please enter up to 100 characters', trigger: 'blur'},
            {required: true, min: 1, max: 100, message: 'Please enter up to 100 characters', trigger: 'input'},
        ],
        projectContent: [
            {required: true, min: 1, max: 1000, message: 'Please enter up to 100 characters', trigger: 'blur'},
            {required: true, min: 1, max: 1000, message: 'Please enter up to 100 characters', trigger: 'input'},
        ],
        startTime: [
            {
                required: true,
                message: 'Please select time',
                trigger: 'blur'
            }, {
                message: 'End time must be greater than start time',
                trigger: 'blur',
                validator: selectTimeChangeprojectExperience
            }],
        endTime: {
            message: 'End time must be greater than start time',
            trigger: 'blur',
            validator: selectTimeChangeprojectExperience
        },
    }
)

let addToprojectExperienceRef = ref()
let loadingProjectExperience = ref(false)
// 项目经验-提交
const submitAddToProjectExperience = () => {
    addToprojectExperienceRef.value.validate(async (valid, fields) => {
        if (valid) {
            loading.value = true
            let data = addToprojectExperience.value;

            let projectSkills = ''

            if (data.projectSkills.length) {
                data.projectSkills.forEach((item, index) => {
                    if (index) projectSkills += `,${item}`
                    else projectSkills += `${item}`
                })
            } else {
                delete data.projectSkills
            }

            for (const dataKey in data) {
                if (data[dataKey] === '') delete data[dataKey]
            }

            data.projectSkills = projectSkills

            if (props.getId) data.id = props.getId
            else data.id = 0

            let res = await apiUserProjectAdd(data)
            if (res.data && res.data.id) {
                apiUserProjectList().then(sonRes => {

                    loading.value = false
                    dialogBox.value = false
                    if (sonRes.data && sonRes.data.list.length) {
                        let data = JSON.parse(JSON.stringify(sonRes.data.list))
                        data.forEach(item => {
                            if (item.projectSkills.indexOf(",") !== -1 || item.projectSkills.length) {
                                item.projectSkills = item.projectSkills.split(',')
                            }
                        })
                        userInfoStore.projectExperienceList = data
                    }
                }).catch(err => {
                    loading.value = false
                })
            }
        }
    })
};


let {volunteerExperienceForm} = storeToRefs(userInfoStore);
// 志愿者-表单规则验证
let rulesVolunteerExperienceForm = ref({
    title: [
        {required: true, min: 1, max: 100, message: 'Please enter up to 100 characters', trigger: 'blur'},
    ],
    content: [
        {required: true, min: 1, max: 1000, message: 'Please enter up to 1000 characters', trigger: 'blur'},
    ],
    startTime: [
        {
            required: true,
            message: 'Please select time',
            trigger: 'blur'
        }, {
            message: 'End time must be greater than start time',
            trigger: 'blur',
            validator: selectTimeChangeVolunteerExperience
        }],
    endTime: {
        message: 'End time must be greater than start time',
        trigger: 'blur',
        validator: selectTimeChangeVolunteerExperience
    },
})
let volunteerExperienceFormRef = ref()
// 志愿者经历-操作
const submitaVolunteerExperience = (formEl, num) => {
    if (!num) return;
    formEl.validate(async (valid, fields) => {
        if (valid) {
            loading.value = true
            let data = volunteerExperienceForm.value;

            for (const dataKey in data) {
                if (data[dataKey] === '') delete data[dataKey]
            }

            if (props.getId) data.id = props.getId
            else data.id = 0
            apiVolunteerAdd(data).then(res => {
                if (res.data && res.data.id) {
                    apiVolunteerList().then(sonRes => {
                        loading.value = false
                        dialogBox.value = false
                        if (sonRes.data && sonRes.data.list.length) {
                            let data = JSON.parse(JSON.stringify(sonRes.data.list))
                            userInfoStore.volunteerExperienceList = data
                        } else {
                            loading.value = false
                        }
                    }).catch(err => {
                        loading.value = false
                    })
                }
            }).catch(function (err) {
                loading.value = false
            });

        }
    })
};

// 增加个人实力
let bulletBoxAddToPersonalStrength = useState("_bulletBoxAddToPersonalStrength");
let addToPersonalStrengthVal = ref(JSON.parse(JSON.stringify(useState('_personalStrengthTxt').value)));
let loadingPersonalStrength = ref(false)
// 个人实力-操作
const submitaAddToPersonalStrength = () => {
    if (addToPersonalStrengthVal.value) {
        loadingPersonalStrength.value = true
        let params = {personalStrength: addToPersonalStrengthVal.value}
        apiUpdatePersonalStrength(params).then(res => {
            loadingPersonalStrength.value = false
            if (res.data && res.data.status) {
                dialogBox.value = false
                useState('_personalStrengthTxt').value = addToPersonalStrengthVal.value
            }
        }).catch(err => {
            loadingPersonalStrength.value = false
        })
    }
};
//
</script>

<style lang="scss">
.birthday {
  .el-form-item__content {
    display: block;
  }
}

@media (max-width: 420px) {
  .birthday {
    .el-col-11 {
      width: 100%;
      max-width: 100%;
    }
  }
}
</style>

<style lang="scss" scoped>
@import "vue-cropper/dist/index.css";
@import "assets/styles/personal-wditor-bullet-box.scss";
@import "assets/styles/h5-personal-wditor-bullet-box.scss";
</style>