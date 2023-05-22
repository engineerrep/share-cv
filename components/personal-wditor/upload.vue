<template>
  <div class="upload" v-loading="loading">
    <div class="upload-main center " id="uploadMain">
      <div class="upload-head " id="uploadHead">
        <span> Add to {{ uploadTitle }}</span>
        <img @click="showHideUpload=false" class="closure-img" src="@/assets/img/delete-img.png">
      </div>
      <DragAndDrop></DragAndDrop>

      <n-space justify="end" class="tail operate-btn" id="uploadTail">
        <CVButton v-if="props.uploadType === 'photos' " @click="submitUpload">
          <template #main>
              <span>
              Apply
              </span>
          </template>
        </CVButton>
        <CVButton v-if="props.uploadType === 'videos'" @click="submitUploadVideos">
          <template #main>
              <span>
              Apply
              </span>
          </template>
        </CVButton>
      </n-space>
    </div>
  </div>
</template>

<script setup>
import {NButton, NSpace} from "naive-ui";
import axios from "axios";
import DragAndDrop from "@/components/dragAndDrop.vue";
import {getCookie} from "~/utils/util";
import {apiPhotoAdd, apiPhotoDelete, apiPhotoList, apiVideoAdd, apiVideoList} from "~/api/personal-wditor";
import {useState} from "#app";
import CVButton from '@/components/form/button.vue';

onMounted(() => {
  document.body.style.cssText = 'overflow-y:hidden';
  document.documentElement.style.overflow = 'hidden';
})
onUnmounted(() => {
  document.body.style.cssText = 'overflow-y:auto';
  document.documentElement.style.overflow = 'auto';
})

onBeforeUnmount(() => {
  useState("_haveData").value = false
  useState("_fileList").value = []
})

const props = defineProps({
  uploadType: String,
});

let showHideUpload = useState('_showHideUpload')

const config = useRuntimeConfig()
const baseUrl = config.public.VITE_API_URL
// 上传文件-请求地址
let action = ref()
let uploadTitle = ref()
// 上传文件类型
let accept = ref(null)
let lessThan9FileList = ref([])
let greaterThan8FileList = ref([])
// 文件列表
const fileList = useState("_fileList", () => [])
// 筛选小于fileList的索引小于等于8和大于8的数据
let uploadType = useState("_uploadType", () => null)
uploadType.value = props.uploadType

const filterData = (e) => {
  if (e) {
    if (props.uploadType === 'videos') {
      e.type = 'video/mp4'
    } else {
      e.type = 'image/jpg'
    }
    e.state = false
    // if (fileList.value.length < 9) {
    //   lessThan9FileList.value.push(e)
    // } else {
    //   greaterThan8FileList.value.push(e)
    // }
    useState("_fileList").value.push(e)
    haveData.value = true
  } else {
    let lessThanList = []
    let greaterThan = []
    if (useState("_uploadType").value === 'videos') {
      useState("_fileList").value.forEach((item, index) => {
        if (index < 6) {
          lessThanList.push(item);
        } else {
          greaterThan.push(item);
        }
      });
    } else {
      useState("_fileList").value.forEach((item, index) => {
        if (index < 9) {
          lessThanList.push(item);
        } else {
          greaterThan.push(item);
        }
      });
    }

    lessThan9FileList.value = lessThanList
    greaterThan8FileList.value = greaterThan
  }
  useState('_videoAdd').value = true
}
let haveData = useState("_haveData", () => false)

// 获取数据
const getData = () => {
  let data = []

  if (props.uploadType === 'photos') {

    action.value = `${baseUrl}/photos`
    uploadTitle.value = 'photos'
    accept.value = 'image/png,image/jpg'
    // 照片-数据
    apiPhotoList().then(res => {
      useState("_loading").value = false

      if (res.data.list) {
        isfileList.value = true
        res.data.list.forEach(item => {
          data.push({
            url: localStorage.getItem('infoOssBucketUrl') + item.fileUrl,
            fileUrl: item.fileUrl,
            id: item.id,
            state: false,
            orderBy: item.orderBy,
            raw: {
              type: "image/jpg"
            }
          })
        })
        useState("_fileList").value = data
        filterData()
        haveData.value = true
      } else {
        useState("_haveData").value = false
        useState("_fileList").value = []
      }
    })
  } else if (props.uploadType === 'videos') {

    action.value = `${baseUrl}/videos`
    uploadTitle.value = 'videos'
    accept.value = 'video/mp4'
    // 视频-数据
    apiVideoList().then(res => {
      useState("_loading").value = false
      if (res.data.list) {
        isfileList.value = true
        res.data.list.forEach(item => {
          data.push({
            url: localStorage.getItem('infoOssBucketUrl') + item.fileUrl,
            coverUrl: localStorage.getItem('infoOssBucketUrl') + item.coverUrl,
            coverUrl02: item.coverUrl,
            id: item.id,
            playTime: item.playTime,
            videoSize: item.videoSize,
            orderBy: item.orderBy,
            fileUrl: item.fileUrl,
            upload: true,
            state: false,
            raw: {
              type: "video/mp4"
            }
          })
        })
        haveData.value = true
        useState("_fileList").value = data
        filterData()
      } else {
        useState("_haveData").value = false
        useState("_fileList").value = []
      }
    })
  }
}

getData()
let loading = useState("_loading")
loading.value = true
// 是否有数据
let isfileList = ref(fileList.value.length ? true : false)

// 当用户-添加数据时触发
const showChangeData = (e) => {
  loading.value = true
  if (e.size < 20971521) {
    if (e.raw.type.indexOf('image') !== -1) {
      isfileList.value = true
      loading.value = false
      filterData(e)
    } else if (e.raw.type.indexOf('video') !== -1) {
      isfileList.value = true
      loading.value = false
      filterData(e)
    } else {
      setTimeout(() => {
        fileList.value = fileList.value.filter(item => item.uid !== e.uid);
        setTimeout(() => {
          loading.value = false
        }, 500)
        // 判断是图片和视频吗
        if (props.uploadType === 'videos') {
          ElMessage({
            showClose: true,
            message: 'Please upload a video in MP4 format.',
            type: 'error',
          })
        } else {
          ElMessage({
            showClose: true,
            message: 'Please upload an image in PNG or JPG format.',
            type: 'error',
          })
        }
      })
    }
  } else {

  }
}

// 当删除照片后，是最后一张显示没有数据的上传组件
const handleRemove = () => {
  if (!fileList.value.length) {
    document.querySelector("#haveData").style.display = "none";
    document.querySelector("#nullData").style.display = "block";
  }
}
let videoRef = useState("_videoRef", () => null)
// 提交上传图片
const submitUpload = async () => {
  if (!fileList.value.length) {
    return
  }

  const files = fileList.value
  for (let i = 0; i < files.length; i++) {
    let item = files[i]
    if (item.status) {//=== "ready" || item.status === "success"
      loading.value = true
      let formData = new FormData()
      await compress(item.raw, 0.2).then(async compressFile => {
        formData.append("file", compressFile.file)
        const config = useRuntimeConfig()
        const baseUrl = config.public.VITE_API_URL
        let res = await axios({
          method: 'post',
          url: `${baseUrl}/api/oss/upload`,
          data: formData,
          headers: {
            'Content-Type': 'multipart/form-data',
            'Authorization': getCookie("token")
          }
        })
        item.fileUrl = res.data.data.filePath
        item.id = 0
      })


    }
  }

  let list = []
  files.forEach(item => {
    list.push({
      id: item.id,
      orderBy: item.orderBy,
      fileUrl: item.fileUrl,
      title: "",
    })
  })

  if (list.length) {
    apiPhotoAdd({list}).then(sonRes => {
      if (sonRes.data && sonRes.data.status) {
        loading.value = false
        getData()
        showHideUpload.value = false
      }
    }).finally(() => {
      loading.value = false
    })
  }
}

// 上传视频
const submitUploadVideos = async () => {
  if (!fileList.value.length) {
    return
  }
  const files = fileList.value
  for (let i = 0; i < files.length; i++) {
    let item = files[i]
    if (item.status === "ready") {
      loading.value = true
      break
    } else {
      loading.value = false
    }
  }

  const config = useRuntimeConfig()
  const baseUrl = config.public.VITE_API_URL

  let list = []
  let videoDom = document.querySelectorAll('.upload-item ')

  async function printVideoInfo(res, url, index, dom, name) {
    let scale = 0.3
    let canvas = document.createElement('canvas')
    let width = dom.videoWidth * scale
    let height = dom.videoHeight * scale
    const MIN_WIDTH = 375
    if (width < MIN_WIDTH) {
      height = height * MIN_WIDTH / width
      width = MIN_WIDTH
    }
    canvas.width = width
    canvas.height = height
    canvas.getContext('2d').drawImage(dom, 0, 0, width, height)
    let image = canvas
        .toDataURL('image/jpeg')
        .replace('image/jpeg', 'image/octet-stream')

    let video = videoDom[index];
    video.src = url
    video.style.display = "block";
    let playTime = null
    await new Promise(resolve => {
      video.addEventListener("loadedmetadata", function () {
        playTime = video.duration
        resolve();
      });
    });

    function dataURLtoFile(dataurl, filename) {
      var arr = dataurl.split(','), mime = arr[0].match(/:(.*?);/)[1],
          bstr = atob(arr[1]), n = bstr.length, u8arr = new Uint8Array(n);
      while (n--) {
        u8arr[n] = bstr.charCodeAt(n);
      }
      return new File([u8arr], filename, {type: mime});
    }

    const base64String = image;

    const blob = dataURLtoFile(base64String, name + '.png');
    let formData = new FormData()
    formData.append("file", blob)
    let resSon = await axios({
      method: 'post',
      url: `${baseUrl}/api/oss/upload`,
      data: formData,
      headers: {
        'Content-Type': 'multipart/form-data',
        'Authorization': getCookie("token")
      }
    })

    list.push({
      id: 0,
      orderBy: index,
      fileUrl: res.data.data.filePath,
      videoSize: Math.ceil(files[index].size / 1024),
      playTime: Math.ceil(playTime),
      coverUrl: resSon.data.data.filePath,
      title: ""
    })

    // console.log(list)

  }

  for (let index = 0; index < videoDom.length; index++) {
    let uploadItem = videoDom[index]
    if (uploadItem.getAttribute("data-upload") === 'false') {
      let item = files[index]
      if (item.status) {
        let formData = new FormData()
        formData.append("file", item.raw)
        let res = await axios({
          method: 'post',
          url: `${baseUrl}/api/oss/upload`,
          data: formData,
          headers: {
            'Content-Type': 'multipart/form-data',
            'Authorization': getCookie("token")
          }
        })
        await printVideoInfo(res, item.url, index, uploadItem, item.uid)
      }
    }
  }

  if (isAnyUploadTrue(videoDom)) {
    apiVideoAdd({list}).then(resVideo => {
      getData()
      loading.value = false
      showHideUpload.value = false
      for (let i = 0; i < videoDom.length; i++) {
        videoDom[i].setAttribute('data-upload', 'true')
      }
    })
  } else {
    showHideUpload.value = false
  }

}

function isAnyUploadTrue(videoDom) {
  for (let i = 0; i < videoDom.length; i++) {
    const isUpload = videoDom[i].getAttribute('data-upload') === 'true';
    if (!isUpload) {
      return true;
    }
  }
  return false;
}

const onProgress = (e) => {
  console.log(e)
}

// 删除图片和视频
const deleteUploadRequest = (file) => {
  console.log(file)
  if (file.id) {
    loading.value = true
    apiPhotoDelete({id: file.id}).then(res => {
      fileList.value = fileList.value.filter(item => item.uid !== file.uid);
      loading.value = false
      filterData()
      // getData()
      if (!fileList.value.length) {

      }
    })
  } else {
    if (!fileList.value.length) {
      isfileList.value = false
    }
    fileList.value = fileList.value.filter(item => item.uid !== file.uid);
    console.log(fileList.value)
    filterData()
  }
}

</script>

<style>

</style>
<style lang="scss" scoped>
.upload {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  z-index: 100;
  background-color: rgba(54, 54, 54, 0.4);

  .upload-main {
    width: 736px;
    max-width: 736px;
    max-height: 80%;
    background: #FFFFFF;
    border-radius: 12px 12px 12px 12px;
    opacity: 1;
    display: flex;
    flex-direction: column;
  }

  .upload-head {
    position: sticky;
    top: 0;
    background-color: #fff;
    z-index: 50;
    padding: 24px 32px;
    font-size: 25px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #000000;
    line-height: 29px;
    border-bottom: 1px solid #eaeaea;
    display: flex;
    justify-content: space-between;
  }

  .tail {
    margin-top: 32px;
    border-top: 1px solid #eaeaea;
    padding: 22px 32px;
    position: sticky;
    bottom: 0;
    background-color: #fff;
    z-index: 50;
  }

  .save-tail-btn {
    padding: 22px 36px;
    font-size: 21px;
    font-family: Inter-Semi Bold, Inter;
    font-weight: normal;
    color: #FFFFFF;
    line-height: 25px;
    background: #0A66C2;
  }
}

.closure-img {
  width: 24px;
  height: 24px;
  cursor: pointer;
}

.plate-six {
  padding: 32px;
}

.upload-demo-main {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 100%;

  .upload-main-title {
    height: 24px;
    font-size: 20px;
    font-family: Inter-Medium, Inter;
    font-weight: 500;
    color: #0A66C2;
    line-height: 23px;
  }

  .upload-main-subtitle {
    height: 22px;
    font-size: 18px;
    font-family: Inter-Regular, Inter;
    font-weight: 400;
    color: #1A1A1A;
    line-height: 21px;
    margin-top: 11px;
  }
}

.center {
  position: absolute;
  top: 50%;
  left: 50%;
  z-index: 100;
  max-height: 100%;
  overflow: scroll;
  transform: translate(-50%, -50%);
  background-color: #fff;
}

.have-data {
  display: none;
  padding: 24px;
  background: #F4F2F0;
  border-radius: 6px;
  flex-wrap: wrap;
}

.file-center {
  display: flex;
  align-items: center;
  width: 190px;
  height: 190px;
  position: relative;
  border-radius: 6px;
  border: 2px solid transparent;
  overflow: hidden;

  .upload-item {
    width: 190px;
    height: 190px;
  }

  .delete-upload {
    position: absolute;
    top: 0;
    right: 0;
    width: 36px;
    height: 36px;
    cursor: pointer;
    z-index: 40;
  }
}

.border_2 {
  border: 2px solid #0A66C2;
}

.photos-item {
  width: 190px;
  height: 190px;
  margin-right: 26px;
  margin-bottom: 20px;
  border-radius: 6px;
  overflow: hidden;

  .img {
    width: 190px;
    height: 190px;
  }

  .delete-upload {
    position: absolute;
    top: 0;
    right: 0;
    width: 36px;
    height: 36px;
    cursor: pointer;
    z-index: 40;
  }
}

.photos-item:nth-child(3n),
.greater-than8-item:nth-child(3n) {
  margin-right: 0;
}


//.greater-than8 {
//  width: calc(100% + 48px);
//  position: relative;
//  left: -24px;
//  padding: 86px 23px 0;
//  margin-top: 30px;
//  box-shadow: inset 0px 0px 6px 0px rgba(0, 0, 0, 5%);
//  border-radius: 6px;
//  background: linear-gradient(#f2f0ee 0%, rgba(0, 0, 0, 10%));
//}

.greater-than8-title {
  position: absolute;
  top: 30px;
  left: 184px;
  z-index: 20;
  height: 22px;
  font-size: 18px;
  font-family: Inter-Regular, Inter;
  font-weight: 400;
  color: #1A1A1A;
  line-height: 21px;
}

.video-upload-item {
  width: 190px;
  height: 190px;
  background-size: cover;
  background-position: center center;
}

.select-upload {
  display: inline-block;
}

.unselected {
  width: 100%;
}

.null-data {
  height: 420px;
}

.cv-button {
  width: 104px;
  height: 42px
}

@media (max-width: 420px) {
  .upload {
    .upload-main {
      width: calc(100% - 32px) !important;
      overflow-y: hidden;
    }
  }
  .upload {
    .save-tail-btn {
      padding: 9px 16px;
      font-size: 18px;
      font-family: Inter-Semi Bold, Inter;
      font-weight: normal;
      line-height: 21px;
    }

    .tail {
      margin-top: 0;
      padding: 12px 22px;
    }
  }
}
</style>
