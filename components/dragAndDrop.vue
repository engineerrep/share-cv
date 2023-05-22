<template>
  <!--  <input ref="fileInput" type="file" style="display: none;" @change="handleFileChange">-->
  <div class="move" :style="fileList.length===0?'overflow: hidden':'overflow: scroll;'">
    <div class="drag-and-drop-main" :class="`${uploadType}-main`">
      <div v-if="!isPlayVideo">
        <div class="drag-and-drop">
          <div class="container "
               :class="fileList.length?'':'null-data',
           uploadType==='videos'?'videos-upload':'',
            containerCorner,
            `${uploadType}-container`
            "
               :style="{height:containerheight,overflow:overflow}"
          >
            <div
                v-for="(item, index) in fileList"
                :key="index"
                class="relative"
            >
              <div class="cls-default photos-item "
                   :class="item.state?'border_2':'',uploadType==='videos'?'border_1':''"
                   @click="item.state = !item.state"
                   v-if="uploadType==='videos'?index<6:index<9"
              >
                <div class="relative">
                  <div v-if="item.raw.type.indexOf('image/')>=0" class="relative">
                    <img :src="item.url" class="img" :data-index="index" :data-upload="item.upload"
                         :data-orderBy="index"
                         draggable="true"
                         @dragstart="onDragstart($event)"
                         @dragend="onDragend(item, $event)"
                         @dragover="onDragover($event)"
                         @drop="onDrop($event)"
                    >
                  </div>

                  <div v-else-if="item.raw.type.indexOf('video/')>=0"
                       class="relative"
                  >
                    <video v-if="item.type==='video/mp4'" class="upload-item video-upload-item"
                           :src="item.url" draggable="true"
                           @dragstart="onDragstart($event)"
                           @dragend="onDragend(item, $event)"
                           @dragover="onDragover($event)"
                           @drop="onDrop($event)"
                           ref="videoRef" data-upload="false" :data-index="index"
                           :data-orderBy="index">
                    </video>
                    <video v-else class="upload-item video-upload-item" poster=" "
                           :src="item.url" draggable="true"
                           @dragstart="onDragstart($event)"
                           @dragend="onDragend(item, $event)"
                           @dragover="onDragover($event)"
                           @drop="onDrop($event)"
                           ref="videoRef" :data-upload="item.upload" :data-index="index"
                           :data-orderBy="index">
                    </video>
                    <img @click.stop="playVideo(item)" class="play-iocn"
                         src="@/assets/img/personal-center/icon/play-iocn.png">
                  </div>
                  <img
                      v-if="item.state"
                      @click="deleteUploadRequest(item)"
                      class="delete-upload"
                      src="@/assets/img/personal-center/icon/delete-upload-icon.png">
                  <img
                      v-if="item.state"
                      class="item-file"
                      src="@/assets/img/personal-center/preview-resume/item-file.png">
                </div>
              </div>
            </div>

            <el-upload
                class="select-upload have-data"
                :class="isfileList?'litte-select-upload':''"
                id="haveData"
                list-type="picture-card"
                :auto-upload="false"
                :before-upload="hanleBeforeUpload"
                :on-change="showChangeData"
                :accept="uploadType === 'photos'?'image/*':'video/*'"
                :limit="uploadType === 'photos'?50:20"
                drag
                multiple
            >
              <el-icon v-if="isfileList">
                <Plus/>
              </el-icon>

              <div v-else class="upload-demo-main">
                <div>
                  <img v-if="uploadType=== 'photos' "
                       src="@/assets/img/personal-center/icon/upload-photos2.png">
                  <img v-else-if="uploadType === 'videos' "
                       src="@/assets/img/personal-center/icon/upload-videos2.png">
                </div>

                <div v-if="uploadType === 'photos' ">
                  <div class="upload-main-title">
                    {{ photosUploadPromptTxt }}
                  </div>
                  <div class="upload-main-subtitle">
                    Set show, Show up to 9 pictures
                  </div>
                </div>
                <div v-else-if="uploadType=== 'videos' ">
                  <div class="upload-main-title">
                    {{ videosUploadPromptTxt }}
                  </div>
                  <div class="upload-main-subtitle">
                    Set show, Show up to 6 videos
                  </div>
                </div>

              </div>

            </el-upload>
          </div>
        </div>

      </div>
    </div>
    <div class="greater-than8-main  relative">
      <div class="greater-than8-title  " v-if="fileList02.length">Drag and drop to replace and be viewed by
        others
      </div>

      <div class="greater-than8" v-if="fileList02.length">
        <div v-for="(item, index) in fileList"
             :key="index"
             :data-index="index"
        >
          <div
              class="cls-default photos-item "
              :class="item.state?'border_2':'',uploadType==='videos'?'border_1':''"
              @click="item.state = !item.state"
              v-if="uploadType==='videos'?index>5:index>8"
          >
            <div class="relative">
              <div v-if="item.raw.type==='image/jpg'||item.type=='image/jpg'" class="relative">
                <img :src="item.url" class="img" :data-index="index" :data-upload="item.upload"
                     :data-orderBy="index"
                     draggable="true"
                     @dragstart="onDragstart($event)"
                     @dragend="onDragend(item, $event)"
                     @dragover="onDragover($event)"
                     @drop="onDrop($event)"
                >
              </div>

              <div v-else-if="item.raw.type==='video/mp4'||item.type==='video/mp4'" class="relative"
              >
                <video v-if="item.type==='video/mp4'" class="upload-item video-upload-item"
                       :src="item.url"
                       draggable="true"
                       @dragstart="onDragstart($event)"
                       @dragend="onDragend(item, $event)"
                       @dragover="onDragover($event)"
                       @drop="onDrop($event)"
                       ref="videoRef" data-upload="false" :data-index="index" :data-orderBy="index">
                </video>
                <video v-else class="upload-item video-upload-item" poster=" "
                       :src="item.url"
                       draggable="true"
                       @dragstart="onDragstart($event)"
                       @dragend="onDragend(item, $event)"
                       @dragover="onDragover($event)"
                       @drop="onDrop($event)"
                       ref="videoRef" :data-upload="item.upload" :data-index="index"
                       :data-orderBy="index">
                </video>
                <img @click.stop="playVideo(item)" class="play-iocn"
                     src="@/assets/img/personal-center/icon/play-iocn.png">
              </div>
              <img
                  v-if="item.state"
                  @click="deleteUploadRequest(item)"
                  class="delete-upload"
                  src="@/assets/img/personal-center/icon/delete-upload-icon.png">
              <img
                  v-if="item.state"
                  class="item-file"
                  src="@/assets/img/personal-center/preview-resume/item-file.png">
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="play-video" v-if="isPlayVideo">
    <div class="play-video-main">
      <video controls autoplay="" name="media">
        <source
            :src="urlVideo"
            type="video/mp4">
      </video>
      <img @click="closureVideo" class="delete-img"
           src="@/assets/img/personal-center/icon/closure-video-iocn.png">
    </div>
  </div>
</template>

<script setup>
import {Plus} from '@element-plus/icons-vue'
import {apiPhotoDelete, apiPhotoList, apiVideoDelete, apiVideoList} from "~/api/personal-wditor";
import {useState} from "#app";

let fileList = useState("_fileList")
fileList.value = []
let videoRef = ref(null)
// 上传文件-请求地址
let action = ref()
let stargindex = ref("");
let endIndex = ref("");

let fileList01 = ref([]);
let fileList02 = ref([]);
// 视频-上传提示
let videosUploadPromptTxt = ref("Drag the video here, or click Upload")
// 照片-上传提示
let photosUploadPromptTxt = ref("Drag the image here, or click Uploa")

let accept = ref()
let uploadType = useState("_uploadType")

if (useState("_uploadType").value === 'videos') {
  accept.value = 'video/mp4'
} else {
  accept.value = 'image/png,image/jpg,image/jpeg'
}
let hideDom = ref('')
let containerCorner = ref('')
// 设置边框
const setCorner = () => {
  if (uploadType.value === 'videos') {
    containerCorner.value = fileList.value.length > 6 ? 'greater-than-nine ' : ''
  } else if (uploadType.value === 'photos') {
    containerCorner.value = fileList.value.length > 9 ? 'greater-than-nine ' : ''
  }
}

setCorner()
watch(() => fileList.value, () => {
  setCorner()
})

let overflow = ref(null)
// 设置大小
const setSize = () => {
  if (document.body.clientWidth > 420) {
    if (useState("_uploadType").value === 'videos') {
      if (fileList.value.length === 0) {
        containerheight.value = '468px'
      } else if (fileList.value.length < 3) {
        containerheight.value = '230px'
      } else if (fileList.value.length < 6) {
        containerheight.value = '450px'
      } else {
        containerheight.value = '670px'
      }
    } else {
      if (fileList.value.length === 0) {
        containerheight.value = '468px'
      } else if (fileList.value.length < 3) {
        containerheight.value = '230px'
      } else if (fileList.value.length < 6) {
        containerheight.value = '450px'
      } else if (fileList.value.length < 9) {
        containerheight.value = '670px'
      } else if (fileList.value.length >= 9) {
        containerheight.value = '890px'
      }
    }
  } else {
    if (useState("_uploadType").value === 'videos') {
      if (fileList.value.length === 0) {
        containerheight.value = '472px';
        overflow.value = 'hidden'
      } else if (fileList.value.length < 3) {
        containerheight.value = '114px'
        overflow.value = 'hidden'
      } else if (fileList.value.length < 6) {
        containerheight.value = '230px'
        overflow.value = 'hidden'
      } else {
        containerheight.value = '324px'
      }
    } else {
      if (fileList.value.length === 0) {
        containerheight.value = '472px'
        nextTick(() => {
          document.querySelector('.drag-and-drop-main').style.height = 'auto'
          document.querySelector('.drag-and-drop-main').style.overflowY = 'hidden'
        })
      } else if (fileList.value.length < 3) {
        containerheight.value = '114px'
        document.querySelector('.drag-and-drop-main').style.height = '146px'
        document.querySelector('.drag-and-drop-main').style.overflowY = 'hidden'
      } else if (fileList.value.length < 6) {
        containerheight.value = '220px'
        document.querySelector('.drag-and-drop-main').style.height = '250px'
        document.querySelector('.drag-and-drop-main').style.overflowY = 'hidden'
      } else if (fileList.value.length < 9) {
        containerheight.value = '330px'
        document.querySelector('.drag-and-drop-main').style.height = '360px'
        document.querySelector('.drag-and-drop-main').style.overflowY = 'hidden'
      } else if (fileList.value.length >= 9) {
        containerheight.value = '470px'
        document.querySelector('.drag-and-drop-main').style.height = '470px'
        document.querySelector('.drag-and-drop-main').style.overflowY = 'hidden'
      }
    }
  }
}
// 是否有数据
let isfileList = ref(fileList.value.length ? true : false)
let containerheight = ref()

const filterData = (e, data) => {
  let lessThanList = [];
  let greaterThan = [];
  if (e) {
    if (e === 'add') {
      if (useState("_uploadType").value === 'videos') {
        data.forEach((item, index) => {
          if (index < 6) {
            lessThanList.push(item);
          } else {
            greaterThan.push(item);
          }
        });
      } else {
        data.forEach((item, index) => {
          if (index < 9) {
            lessThanList.push(item);
          } else {
            greaterThan.push(item);
          }
        });
      }

      fileList01.value = lessThanList;
      fileList02.value = greaterThan;
      loading.value = false
    } else {
      if (useState("_uploadType").value === 'videos') {
        e.type = 'video/mp4'
      } else {
        e.type = 'image/jpg'
      }

      if (useState("_uploadType").value === 'videos') {
        if (useState("_fileList").value.length < 6) {
          fileList01.value.push(e);
        } else {
          fileList02.value.push(e);
        }
      } else {
        if (useState("_fileList").value.length < 9) {
          fileList01.value.push(e);
        } else {
          fileList02.value.push(e);
        }
      }
      fileList.value.push(e)
    }
  } else {
    if (useState("_uploadType").value === 'videos') {
      fileList.value.forEach((item, index) => {
        if (index < 6) {
          lessThanList.push(item);
        } else {
          greaterThan.push(item);
        }
      });
    } else {
      fileList.value.forEach((item, index) => {
        if (index < 9) {
          lessThanList.push(item);
        } else {
          greaterThan.push(item);
        }
      });
    }
    fileList01.value = lessThanList;
    fileList02.value = greaterThan;
  }
  useState('_videoAdd').value = false
  setSize()
  if (fileList.value.length) isfileList.value = true
  else isfileList.value = false
};
filterData()

let videoAdd = useState('_videoAdd', () => false)
watch(() => videoAdd.value, (newVal) => {
      if (newVal) {
        filterData()
      }
    },
    {immediate: true, deep: true}
)

onMounted(() => {
  if (document.body.clientWidth <= 420) {
    videosUploadPromptTxt.value = 'Click to upload video'
    photosUploadPromptTxt.value = 'Click to upload image'
  }
})

const getData = () => {

  let data = []
  if (useState("_uploadType").value === 'photos') {
    // 照片-数据
    apiPhotoList().then(res => {
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
        fileList.value = data
        filterData()
      } else {
        fileList.value = data
        filterData()
      }

    })
  } else if (useState("_uploadType").value === 'videos') {
    // 视频-数据
    apiVideoList().then(res => {
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
        fileList.value = data
        filterData()
      } else {
        fileList.value = data
        filterData()
      }
    })
  }


  setTimeout(() => {
    useState("_loading").value = false
  }, 1000)
}
// getData()
const onDragstart = (event) => {
  stargindex.value = event.target.getAttribute("data-index");
};
const onDrop = (event) => {
  if (event.target.getAttribute("data-index")) {
    endIndex.value = event.target.getAttribute("data-index")
  }
};
// 拖拽
const onDragend = (item, event) => {
  let dataList = fileList.value

  let dataListClone = JSON.parse(JSON.stringify(dataList))
  // 拖拽的索引
  let dataListCloneIndex = JSON.parse(JSON.stringify(dataList)).findIndex(items => items.orderBy === item.orderBy);
  // 被替换的索引
  let endIndexClone = Number(endIndex.value)
  let aimObj = dataListClone.splice(dataListCloneIndex, 1)[0]
  dataListClone.splice(endIndexClone, 0, aimObj)
  dataListClone.forEach((value, index) => {
    value.orderBy = index + 1
  })
  fileList.value = dataListClone

  // dataList[dataListCloneIndex].orderBy = dataListClone[endIndexClone].orderBy;
  // dataList[endIndexClone].orderBy = dataListClone[dataListCloneIndex].orderBy;
  // [fileList.value[dataListCloneIndex], fileList.value[endIndexClone]] =
  //   [fileList.value[endIndexClone], fileList.value[dataListCloneIndex]];

  // if (state) {
  //
  // } else {
  //   ElMessage({
  //     showClose: true,
  //     message: 'There are unuploaded files',
  //     type: 'error',
  //   })
  // }
};

const onDragover = (event) => {
  event.preventDefault();
};

let loading = useState("_loading")

const hanleBeforeUpload = (e) => {
  if (uploadType.value === 'videos') {

  }
}

// 当用户-添加数据时触发
const showChangeData = async (e) => {
  loading.value = true
  useState("_videoRef").value = videoRef.value
  if (uploadType.value === 'videos') {
    if (e.size < 20971521 && e.raw.type.indexOf('video/') >= 0) {
      let uploadMain = document.querySelector('#uploadMain')
      uploadMain.style.overflow = 'scroll'
      if (window.createObjectURL != undefined) {
        e.url = window.createObjectURL(e.raw);
      } else if (window.URL != undefined) {
        e.url = window.URL.createObjectURL(e.raw);
      } else if (window.webkitURL != undefined) {
        e.url = window.webkitURL.createObjectURL(e.raw);
      }
      if (e.isChange) {
        loading.value = false
        return
      }
      e.isChange = true
      addFileList(e)
    } else {
      loading.value = false
      ElMessage({
        showClose: true,
        message: 'Video selection type not supported',
        type: 'error',
      })
    }
  } else if (uploadType.value === 'photos') {

    if (e.raw.type.indexOf('image/') >= 0) {//e.size < 1000000
      // addFileList(e)
      // let uploadMain = document.querySelector('#uploadMain')
      // uploadMain.style.overflow = 'hidden'
      if (e.isCompress) {
        loading.value = false
        return
      }
      await compress(e.raw, 0.2).then(async compressFile => {
        e.url = compressFile.afterSrc
        e.raw = compressFile.file
        e.isCompress = true
        addFileList(e)
        let uploadMain = document.querySelector('#uploadMain')
        uploadMain.style.overflow = 'hidden'
        loading.value = false
      })
    } else {
      loading.value = false
      ElMessage({
        showClose: true,
        message: 'Photo selection type not supported',
        type: 'error',
      })
    }
  } else {
    //   没有视频和照片类型
  }
}

// 添加图片和视频
const addFileList = (e) => {
  const maxOrderBy = Math.max(...fileList.value.map(item => item.orderBy));
  if (fileList.value.length) {
    e.orderBy = maxOrderBy + 1
  } else {
    e.orderBy = 1
  }
  filterData(e)
  if (e.raw.type.indexOf('image') !== -1) {
    isfileList.value = true
    loading.value = false
  } else if (e.raw.type.indexOf('video') !== -1) {
    isfileList.value = true
    loading.value = false
  }
}

// 删除图片和视频
const deleteUploadRequest = (file) => {
  loading.value = true
  if (file.id) {
    if (useState("_uploadType").value === 'photos') {
      apiPhotoDelete({id: file.id}).then(res => {
        getData()
      })
    } else if (useState("_uploadType").value === 'videos') {
      apiVideoDelete({id: file.id}).then(res => {
        getData()
      })
    }

  } else {
    fileList.value = fileList.value.filter(item => item.uid !== file.uid);
    filterData()
    setTimeout(() => {
      loading.value = false
    }, 1000)
  }
  if (!useState("_fileList").value.length) {
    let uploadMain = document.querySelector('#uploadMain')
    useState("_haveData").value = false
    uploadMain.style.overflow = 'hidden'
  }
}

let isPlayVideo = useState("_isPlayVideo", () => false)
let urlVideo = ref('')
// 播放视频
const playVideo = (item) => {
  let dom = document.querySelector("#uploadMain")
  dom.style.height = '558px'
  dom.style.width = '990px'
  dom.style.maxWidth = '990px'
  if (item?.type) {
    urlVideo.value = reloadVideo(urlVideo.value, item.raw)
  } else {
    urlVideo.value = item.url
  }
  isPlayVideo.value = true
}

/**
 * 重新加载视频
 * @param val 要重新加载的值
 * @param file  要重新加载的视频
 */
const reloadVideo = (val, file) => {
  if (window.createObjectURL != undefined) {
    val = window.createObjectURL(file);
  } else if (window.URL != undefined) {
    val = window.URL.createObjectURL(file);
  } else if (window.webkitURL != undefined) {
    val = window.webkitURL.createObjectURL(file);
  }
  return val
}

// 关闭视频
const closureVideo = () => {
  let dom = document.querySelector("#uploadMain")
  dom.style.width = '736px'
  dom.style.maxWidth = '990px'
  dom.style.height = 'auto'
  // 重新加载视频
  fileList.value.forEach(item => {
    if (item.status === "ready") {
      item.url = reloadVideo(item.url, item.raw)
    }
  })
  isPlayVideo.value = false
}

</script>

<style lang="scss">
.el-upload--picture-card {
  border: 0;
  background-color: transparent;
}

.el-upload {
  width: 20px;
  height: 20px;
  position: relative;
  left: 0;
}


.el-upload-dragger {
  left: 6px;
  background: #E0E0E0;

  &:hover {
    border: 0;
  }
}

.el-upload:focus .el-upload-dragger {
  border: 0;
}

.null-data {
  .el-upload-dragger {
    background: #f2f0ee;
  }
}

.container {
  display: flex;
  flex-wrap: wrap;
}

@media (max-width: 420px) {
  .greater-than8-main {
    position: relative;
    top: -24px;
  }
  .el-upload-dragger {
    top: -2px;
  }
  .el-upload {
    position: relative;
    top: 2px;
    left: 3px;
    //height: 94px !important;
  }
  .null-data {
    .el-upload-dragger {
      left: 40px;
    }
  }
  .el-upload-dragger {
    left: 0;
  }
  .null-data {
    width: 100% !important;

    .upload-demo-main, .el-upload-dragger {
      width: 100% !important;
      padding: 10px;
    }

    .el-upload-dragger {
      border: 0;
    }

    .el-upload-dragger {
      width: 100%;
    }
  }

  .select-upload .el-upload-dragger, .el-upload-list, .select-upload {
    width: 86px;
    height: 88px;
  }
  .have-data .el-upload-list .el-upload--picture-card {
    width: 100%;
    height: 100%;
  }
  .container {
    padding: 16px 10px 0px 10px !important;
    overflow: hidden;

    > :nth-child(3n) {
      .photos-item {
        margin-right: 0 !important;
      }
    }

    .photos-item {
      margin-right: 9px !important;
      margin-bottom: -11px !important;
    }

    .select-upload {
      .el-upload-dragger {
        box-shadow: 0 0 0 0 transparent;
        padding: 0 !important;
      }
    }
  }
  .null-data .select-upload {
    width: 100% !important;
  }
  .cls-default {
    right: 0 !important;
  }
  .select-upload .el-upload--picture-card .el-upload-dragger .el-icon {
    top: -10px;
  }
  .videos-main {
    overflow-x: hidden;

    .null-data {
      .el-upload-dragger {
        width: 100% !important;
      }
    }

    .el-upload-list {
      > :nth-child(2) {
        //top: 86px;
        //left: -112px;
      }

      > :nth-child(n+8) {
        top: 64px;
        left: -7px;
      }
    }

    .select-upload .el-upload--picture-card .el-upload-dragger .el-icon {
      top: 22px;
    }
  }
}

</style>

<style lang="scss" scoped>
.move {
  overflow: scroll;
  padding: 0 36px;
}

.play-video {
  position: absolute;
  top: 50%;
  left: 50%;
  z-index: 100;
  max-height: 100%;
  transform: translate(-50%, -50%);
  width: 990px;
  height: 558px;
  display: flex;
  align-items: center;
  background-color: #fff;

  .play-video-main {
    width: 990px;
    height: 558px;
    background-color: rgba(0, 0, 0, 0.8);
  }

  video {
    width: 100%;
    height: 100%;
    //object-fit: cover;
  }

  .delete-img {
    position: absolute;
    top: 22px;
    right: 17px;
    width: 40px;
    height: 40px;
    cursor: pointer;
  }
}

.have-data {
  display: none;
}

.drag-and-drop-main {
  margin-top: 32px;
}

.container {
  padding: 24px;
  background: #F4F2F0;
  border-radius: 12px;
}

.cls-default {
  display: inline-block;
  width: 188px;
  height: 188px;
  box-sizing: border-box;
  overflow: hidden;
  margin-right: 26px;
  margin-bottom: 10px;
  overflow: hidden;
  border-radius: 10px;
  position: relative;
  top: -2px;
  right: -6px;
  z-index: 10;
  border: 2px solid transparent;
}

.videos-main {
  .border_1 {
    //border: 2px solid #dcdbdc;
  }

  .border_2 {
    border: 2px solid transparent;
  }
}


.photos-item {
  width: 188px;
  height: 190px;
  margin-right: 17px;
  margin-bottom: 17px;
  border-radius: 8px;
  overflow: hidden;

  .img {
    width: 190px;
    height: 190px;
    object-fit: cover;
  }

  .delete-upload {
    position: absolute;
    top: -2px;
    right: -2px;
    width: 36px;
    height: 36px;
    cursor: pointer;
    z-index: 40;
  }
}

.dom-item {
  width: 190px;
  height: 190px;
}

.greater {
  margin: 0 32px;
  background: linear-gradient(#f2f0ee 0%, rgba(0, 0, 0, 0.1));
}

.photos-item:nth-child(3n),
.greater-than8-item:nth-child(3n) {
  margin-right: 0;
}


.greater-than8 {
  width: 100%;
  position: relative;
  left: 0px;
  padding: 86px 23px 0;
  box-shadow: inset 0px 0px 6px 0px rgba(0, 0, 0, 5%);
  background: linear-gradient(#f2f0ee 0%, rgba(0, 0, 0, 10%));
  display: flex;
  flex-wrap: wrap;
  border-radius: 0 0 12px 12px;
}

.greater-than8-title {
  width: 100%;
  text-align: center;
  position: absolute;
  top: 30px;
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
  background-color: #dcdbdc;
}

.play-iocn {
  position: absolute;
  top: 74px;
  left: 64px;
  width: 50px;
  height: 50px;
  z-index: 20;
  cursor: pointer;
}

.select-upload {
  display: inline-block;
}

.unselected {
  width: 100%;
}

.upload-demo-main {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 100%;

  img {
    width: 80px;
    height: 80px;
  }

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

.null-data {
  width: 673px;
  height: 420px;
  background-color: transparent;
}

.greater-than-nine {
  border-radius: 12px 12px 0 0;
}

.item-file {
  width: 190px;
  height: 186px;
  position: absolute;
  top: 0;
  z-index: 10;
}

.display-none {
  display: none;
}
</style>

<style ang="scss" scoped>

@media (max-width: 420px) {
  .move {
    padding: 0 0px !important;
  }

  :deep(.photos-container), :deep(.videos-container) {
    margin-bottom: 16px;

    .select-upload {
      height: 100%;

      .el-upload-list {
        height: 100%;

        .el-upload {
          margin-top: 27px;

          .el-upload-dragger {
            height: 100% !important;
          }

          .el-upload-dragger {
            height: 100% !important;
          }
        }
      }
    }

    > .relative {
      width: 33%;

      .cls-default {
        width: 100% !important;

        .img {
          width: 100% !important;
        }
      }
    }

    > :nth-child(n+10) {
      display: none;
    }

  }

  :deep(.select-upload ) {
    display: block !important;
  }

  .videos-main {
    .el-upload-dragger {
      .el-icon {
        top: -10px !important;
      }
    }
  }

  .item-file {
    height: 94px;
  }

  .drag-and-drop-main {
    margin-top: 0;
  }

  .upload-demo-main {
    .upload-main-title {
      height: auto;
    }
  }

  .drag-and-drop-main {
    position: relative;
    left: -8px;
    padding: 0 0px 0 16px;
    overflow-y: auto;

    &::-webkit-scrollbar-thumb {
      height: calc(100% - 64px);
    }
  }

  .drag-and-drop {
    padding: 0 8px;
  }

  .photos-item {
    margin-right: 9px;
    margin-bottom: 0;
  }

  .greater-than8 > :nth-child(3n) .photos-item {
    margin-right: 0 !important;
  }

  .greater-than8 {
    width: calc(100% - 32px);
    top: 10px;
    left: 16px;
    padding: 86px 10px 20px;
    margin: 0;
  }

  .greater-than8-title {
    width: 80%;
    left: 50%;
    transform: translateX(-50%);
    font-size: 14px;
    margin: 0 auto;
  }

  .cls-default {
    width: 90px;
    height: 98px;
  }

  .photos-item .img {
    width: 98px;
    height: 98px;
  }

  .container {
    margin-top: 16px;
    width: 100%;
    padding: 7px 0 7px 5px !important;
  }

  .dom-item, .select-upload {
    width: 86px;
    height: 96px;
  }

  .videos-main {
    /*margin-bottom: 16px;*/

    .video-upload-item {
      width: 100%;
      height: 98px;
    }

    .play-iocn {
      width: 20px;
      height: 20px;
      top: 38px;
      left: 38px;
    }
  }

  :deep(.litte-select-upload) {
    width: 33%;
    height: 98px !important;

    .el-upload-list {
      width: 100% !important;

      .el-upload-dragger {
        width: 100%;
      }
    }

    .el-upload {
      margin-top: 0 !important;
      margin-right: 0 !important;
      height: 98px;
      width: 100%;

      .el-icon {
        top: 30px !important;
      }
    }

    .photos-container, .videos-container {
      padding-top: 10px;
      height: 220px;

      > .relative {
        height: 98px;
      }

      .video-upload-item {
        height: 98px !important;
      }
    }

  }
}

.play-video {
  width: 100% !important;
  height: 100% !important;

  .play-video-main {
    width: 100% !important;
    height: 100% !important;
  }
}
</style>