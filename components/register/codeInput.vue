<template>
    <view class="u-code-input">
        <view
                class="u-code-input__item"
                :style="[itemStyle(index)]"
                v-for="(item, index) in codeLength"
                :key="index"
        >
            <!--      <view-->
            <!--          class="u-code-input__item__dot"-->
            <!--          v-if="dot && codeArray.length > index"-->
            <!--      ></view>-->
            <text
                    :style="{
					fontSize: '20px',
					fontWeight: '400',
					color: color
				}"
            >{{ codeArray[index] }}
            </text>
            <view
                    class="u-code-input__item__line"
                    v-if="mode === 'line'"
                    :style="[lineStyle]"
            ></view>
            <!-- #ifndef APP-PLUS -->
            <view v-if="isFocus && codeArray.length === index" :style="{backgroundColor: color}"
                  class="u-code-input__item__cursor"></view>
            <!-- #endif -->
        </view>
        <input
                ref="inputValueRef"
                :disabled="false"
                type="number"
                :focus="focus"
                v-model="inputValue"
                :maxlength="maxlength"
                :adjustPosition="false"
                class="u-code-input__input"
                @input="inputHandler"
                :style="{
				height: '33px',opacity:'0',
				position:'absolute',
				width:'100%'
			}"
                @focus="isFocus = true"
                @blur="isFocus = false"
        />
    </view>
</template>

<script setup>
import {computed} from 'vue'

/**
 * CodeInput 验证码输入
 * @description 该组件一般用于验证用户短信验证码的场景，也可以结合uView的键盘组件使用
 * @tutorial https://www.uviewui.com/components/codeInput.html
 * @property {String | Number}  maxlength      最大输入长度 （默认 6 ）
 * @property {Boolean}      dot          是否用圆点填充 （默认 false ）
 * @property {String}      mode        显示模式，box-盒子模式，line-底部横线模式 （默认 'box' ）
 * @property {Boolean}      hairline      是否细边框 （默认 false ）
 * @property {String | Number}  space        字符间的距离 （默认 10 ）
 * @property {String | Number}  value        预置值
 * @property {Boolean}      focus        是否自动获取焦点 （默认 false ）
 * @property {Boolean}      bold        字体和输入横线是否加粗 （默认 false ）
 * @property {String}      color        字体颜色 （默认 '#606266' ）
 * @property {String | Number}  fontSize      字体大小，单位px （默认 18 ）
 * @property {String | Number}  size        输入框的大小，宽等于高 （默认 35 ）
 * @property {Boolean}      disabledKeyboard  是否隐藏原生键盘，如果想用自定义键盘的话，需设置此参数为true （默认 false ）
 * @property {String}      borderColor      边框和线条颜色 （默认 '#c9cacc' ）
 * @property {Boolean}      disabledDot      是否禁止输入"."符号 （默认 true ）
 *
 * @event {Function}  change  输入内容发生改变时触发，具体见上方说明      value：当前输入的值
 * @event {Function}  finish  输入字符个数达maxlength值时触发，见上方说明  value：当前输入的值
 * @example  <u-code-input v-model="value4" :focus="true"></u-code-input>
 */
let emit = defineEmits(['change', 'input', 'finish'])
let maxlength = ref(6)
let mode = ref('line')
let hairline = ref(false)
let inputValue = ref('')
let isFocus = ref(false)
let color = ref('rgb(51, 54, 57)')
let borderColor = ref('#4d4d4d')
let focus = ref(false)
let inputValueRef = ref()
const props = defineProps({
    value: String | Number
});

watch(() => props.value, (val) => {
    inputValue.value = String(val).substring(0, maxlength.value)
}, {immediate: true, deep: true})
let codeLength = computed(() => {
    return new Array(Number(maxlength.value))
});
let itemStyle = computed(() => {
    return index => {

        const style = {
            width: '16.66%',
            height: '33px'
        }
        if (index !== codeLength.value.length - 1) {
            // 设置验证码字符之间的距离，通过margin-right设置，最后一个字符，无需右边框
            style.marginRight = '10px'
        } else {
            // 最后一个盒子的有边框需要保留
            style.marginRight = 0
        }

        return style
    }
});
let codeArray = computed(() => {
    return String(inputValue.value).split('')
});

let lineStyle = computed(() => {
    const style = {}
    style.height = '1px'
    style.width = '100%'
    // 线条模式下，背景色即为边框颜色
    style.backgroundColor = borderColor.value
    return style
});

const inputHandler = (e) => {
    const value = e.target.value
    //inputValue.value = value
    // 是否允许输入“.”符号
    if (true) {
        nextTick(() => {
            inputValue.value = value.replace('.', '')
        })
    }
    // 未达到maxlength之前，发送change事件，达到后发送finish事件

    emit('change', value)
    // 修改通过v-model双向绑定的值
    emit('input', value)
    // 达到用户指定输入长度时，发出完成事件
    if (String(value).length >= Number(maxlength.value)) {
        focus.value = false
        nextTick(() => {
            inputValue.value = String(value).substring(0, maxlength.value)
            // inputValueRef.value.value = inputValue.value
        })
        emit('finish', inputValue.value)
    }
}

</script>

<style lang="scss" scoped>
// 超出行数，自动显示行尾省略号，最多5行
// 来自uView的温馨提示：当您在控制台看到此报错，说明需要在App.vue的style标签加上【lang="scss"】
//@for $i from 1 through 5 {
//  .u-line-#{$i} {
//    /* #ifndef APP-NVUE */
//    // vue下，单行和多行显示省略号需要单独处理
//    @if $i == '1' {
//      overflow: hidden;
//      white-space: nowrap;
//      text-overflow: ellipsis;
//    } @else {
//      display: -webkit-box!important;
//      overflow: hidden;
//      text-overflow: ellipsis;
//      word-break: break-all;
//      -webkit-line-clamp: $i;
//      -webkit-box-orient: vertical!important;
//    }
//    /* #endif */
//  }
//}


// 此处加上!important并非随意乱用，而是因为目前*.nvue页面编译到H5时，
// App.vue的样式会被uni-app的view元素的自带border属性覆盖，导致无效
// 综上，这是uni-app的缺陷导致我们为了多端兼容，而必须要加上!important
// 移动端兼容性较好，直接使用0.5px去实现细边框，不使用伪元素形式实现
.u-border {
  border-width: 0.5px !important;
  border-color: #4d4d4d !important;
  border-style: solid;
}

.u-border-top {
  border-top-width: 0.5px !important;
  border-color: #4d4d4d !important;
  border-top-style: solid;
}

.u-border-left {
  border-left-width: 0.5px !important;
  border-color: #4d4d4d !important;
  border-left-style: solid;
}

.u-border-right {
  border-right-width: 0.5px !important;
  border-color: #4d4d4d !important;
  border-right-style: solid;
}

.u-border-bottom {
  border-bottom-width: 0.5px !important;
  border-color: #4d4d4d !important;
  border-bottom-style: solid;
}

.u-border-top-bottom {
  border-top-width: 0.5px !important;
  border-bottom-width: 0.5px !important;
  border-color: #4d4d4d !important;
  border-top-style: solid;
  border-bottom-style: solid;
}

// 去除button的所有默认样式，让其表现跟普通的view、text元素一样
.u-reset-button {
  padding: 0;
  background-color: transparent;
  /* #ifndef APP-PLUS */
  font-size: inherit;
  line-height: inherit;
  color: inherit;
  /* #endif */
  /* #ifdef APP-NVUE */
  border-width: 0;
  /* #endif */
}

/* #ifndef APP-NVUE */
.u-reset-button::after {
  border: none;
}

/* #endif */

.u-hover-class {
  opacity: 0.7;
}

$u-code-input-cursor-width: 1px;
$u-code-input-cursor-height: 40%;
$u-code-input-cursor-animation-duration: 1s;
$u-code-input-cursor-animation-name: u-cursor-flicker;

.u-code-input {
  display: flex;
  flex-direction: row;
  position: relative;
  overflow: hidden;
  justify-content: center;

  &__item {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    position: relative;

    &__text {
      font-size: 15px;
      color: rgb(51, 54, 57);
    }

    &__dot {
      width: 7px;
      height: 7px;
      border-radius: 100px;
      background-color: rgb(51, 54, 57);
    }

    &__line {
      position: absolute;
      bottom: 0;
      height: 4px;
      border-radius: 100px;
      width: 40px;
      background-color: rgb(51, 54, 57);
    }

    /* #ifndef APP-PLUS */
    &__cursor {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 1px;
      height: 40%;
      animation: 1s u-cursor-flicker infinite;
    }

    /* #endif */

  }

  &__input {
    // 之所以需要input输入框，是因为有它才能唤起键盘
    // 这里将它设置为两倍的屏幕宽度，再将左边的一半移出屏幕，为了不让用户看到输入的内容
    position: absolute;
    left: -750rpx;
    width: 1500 rpx;
    top: 0;
    background-color: transparent;
    text-align: left;
  }
}

/* #ifndef APP-PLUS */
@keyframes u-cursor-flicker {
  0% {
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}

/* #endif */

</style>