<template>
    <div class=" center02 ">
        <n-card title="Feedback" v-loading="loading">
            <hr style="background-color: #EBEBEB;width: 100%;height: 1px;margin-bottom: 22px">
            <el-form
                    require-asterisk-position="right"
                    ref="feedbackModelRef"
                    label-position="top"
                    :model="feedbackModel"
                    :rules="rulesFeedbackModel">
                <el-form-item label="Conetent" prop="content" required>
                    <el-input
                            class="input-tarea"
                            v-model="feedbackModel.content"
                            type="textarea"
                            placeholder="Please enter up to 1000 characters"
                            maxlength="1000"
                            show-word-limit
                            rows="6"
                            resize='none'
                    />
                </el-form-item>
                <el-form-item :label="`Upload picture(${imgeList.length}/3)`">
                    <div class="img-list">
                        <img :src="infoOssBucketUrl+item" class="img" v-for="(item, index) in imgeList"
                             :key="index"
                             :data-index="index">

                    </div>
                    <el-upload
                            class="select-upload have-data"
                            v-model:file-list="feedbackModel.imgeList"
                            list-type="picture-card"
                            :auto-upload="false"
                            :on-change="handleChange"
                            :accept="'image/*'"
                            :limit="3"
                            v-if="feedbackModel.imgeList.length<3"
                    >
                        <el-icon>
                            <Plus/>
                        </el-icon>
                    </el-upload>
                </el-form-item>
            </el-form>
            <hr style="background-color: #EBEBEB;width: 100%;height: 1px">
            <template #footer>
                <CVButton @click="feedback" style="float: right;width: 104px;height: 42px;">
                    <template #main>
              <span>
              Submit
              </span>
                    </template>
                </CVButton>
            </template>
        </n-card>

    </div>
</template>

<script setup>
import {NCard} from "naive-ui";
import {Plus} from "@element-plus/icons-vue";
import {ElLoading, ElMessage} from 'element-plus'
import axios from "axios";
import {getCookie} from "~/utils/util";
import CVButton from "~/components/form/button.vue";
import {apiAccountChangeEmailCode, apiFeedback} from "~/api/user";
import {useRouter} from "vue-router";

const router = useRouter();

let loading = ref(false)

let feedbackModel = ref({
    content: '',
    imgeList: []
})
let rulesFeedbackModel = ref({
    content: [
        {required: true, min: 10, max: 1000, message: 'Please enter 10-1000 characters', trigger: 'blur'},
    ],
});
let content = ref('');
let imgeList = ref([])
let infoOssBucketUrl = ref('')
onMounted(() => {
    infoOssBucketUrl.value = localStorage.getItem('infoOssBucketUrl')
})
const handleExceed = (files, fileList) => {
    ElMessage.warning(`最多只能上传 3 张照片`);
}
const handleChange = async (file, fileList) => {
    // file.url =
    let formData = new FormData()

    await compress(file.raw, 0.2).then(async compressFile => {
        formData.append("file", compressFile.file)
        const config = useRuntimeConfig()
        const baseUrl = config.public.VITE_API_URL
        loading.value = true
        let res = await axios({
            method: 'post',
            url: `${baseUrl}/api/oss/upload`,
            data: formData,
            headers: {
                'Content-Type': 'multipart/form-data',
                'Authorization': getCookie("token")
            }
        }).finally(() => {
            loading.value = false
        })
        imgeList.value.push(res.data.data.filePath)

    })


}
let feedbackModelRef = ref()
/**
 * 提交反馈
 */
const feedback = () => {

    feedbackModelRef.value.validate((valid, fields) => {
        if (valid) {
            let postData = {
                content: feedbackModel.value.content,
                fileUrls: imgeList.value.join(',')
            }
            loading.value = true
            apiFeedback(postData).then(res => {
                if (res.data.status == 1) {
                    ElMessage.warning(`反馈成功`);
                    feedbackModelRef.value.resetFields()
                    imgeList.value = []
                    router.push('/personal-wditor')
                }
            }).finally(() => {
                loading.value = false
            })
        }
    })

}
</script>

<style scoped lang="scss">
.n-card {
  border-radius: 10px;
  box-shadow: 0 4px 4px #00000040;
  margin: 0 auto;
  padding-top: 30px;
  width: 655px;
  margin-top: 66px;

}

:deep(.el-form) {
  .el-form-item__label {
    font-size: 18px;
    font-weight: 400;
    color: #333333;
  }
}


.img-list {
  display: flex;
  justify-content: flex-start;

  img {
    width: 190px;
    height: 190px;
    margin-right: 10px;
    border-radius: 6px;
  }
}

.n-card--bordered {
  padding-top: 0 !important;
}

:deep(.input-tarea) {
  border: none;

  .el-textarea__inner {
    //border-radius: 6px 6px 6px 6px;
    //opacity: 1;
    //border: 1px solid #616161;
  }

}

:deep(.el-upload-list) {
  .el-upload--picture-card {
    width: 190px;
    height: 190px;
    background: #E0E0E0;
    border-radius: 6px 6px 6px 6px;
  }
}

@media (max-width: 420px) {
  .n-card {
    width: calc(100vw - 32px) !important;
  }
  .img-list {
    img {
      width: calc(33vw - 11px - 28px);
      height: calc(33vw - 11px - 20px);
      margin-right: 15px;
    }

  }
  :deep(.select-upload) {
    width: calc(33vw - 11px - 28px);
    height: calc(33vw - 11px - 20px);

    .el-upload-list {
      width: calc(33vw - 11px - 28px);
      height: calc(33vw - 11px - 20px);
    }

    .el-upload--picture-card {
      width: calc(33vw - 11px - 28px);
      height: calc(33vw - 11px - 20px);
      margin-right: 0 !important;

    }
  }
  :deep(.el-form) {
    .el-form-item__label {
      font-size: 12px;
      font-weight: 400;
      color: #333333;
    }
  }
}
</style>