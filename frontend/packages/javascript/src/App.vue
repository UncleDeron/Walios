<template>
  <div id="walios-app" :class="platform">
    <div id="title-bar" v-if="platform === 'mac'" data-wails-drag>Walios</div>
    <!-- 系统通知栏 -->
    <Notification v-bind="notify"></Notification>
    <!-- 页面 -->
    <div class="view" :class="{'with-notify': notify.active}">
      <router-view v-slot="{ Component, route }">
        <transition :name="route.meta.transition || 'fade'">
          <component :is="Component" />
        </transition>
      </router-view>
    </div>
  </div>
</template>

<script>
import  Notification from "@/components/Notification.vue";
export default {
  setup() {
    const platform = window.navigator.platform.startsWith('Win') ? 'windows' : 'mac';
    return {
      platform
    }
  },
  provide() {
    return {
      app: this,
    };
  },
  components: {
    Notification,
  },
  data() {
    return {
      notify: {
        active: false,
        msg: "服务器连接成功~",
        type: "info",
      },
      notifyTimer: null,
    }
  },
  methods: {
    /**
     * 系统提示
     * @param {String} msg 提示内容
     * @param {String} type 提示类型 success/info/warning/danger
     * @param {Boolean} autoClose 是否自动关闭 默认 true
     * @param {Number} duration 自动关闭时间 默认 3000
     */
    showSystemNotification(msg, type, autoClose = true, duration = 3000) {
      clearTimeout(this.notifyTimer);
      this.notify = {
        active: true,
        msg,
        type,
      }
      if (autoClose) {
        this.notifyTimer = setTimeout(() => {
          this.notify.active = false;
        }, duration);
      }
    },
    addWailsEventListener() {
      runtime.EventsOn("notify", data => {
        let {msg, type, autoClose, duration} = data
        this.showSystemNotification(msg, type, autoClose, duration);
      });

      runtime.EventsOn("closeNotify", data => {
        console.log(data);
        let {msg, type, autoClose, duration} = data
        this.showSystemNotification(msg, type, autoClose, duration);
      });
    }
  },
  mounted() {
    if (window.runtime) {
      this.addWailsEventListener();
    }
  },
  components: { Notification }
};
</script>

<style lang="scss">
@import url("./assets/css/reset.css");
@import url("./assets/css/font.css");
@import url("./assets/css/common.css");

html {
  width: 100%;
  height: 100%;
}

body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-family: "Alibaba";
  background-color: transparent;
}

#app {
  user-select: none;
  position: relative;
  // width: 900px;
  // height: 520px;
  height: 100%;
  background-color: #DFDBCB;
  overflow: hidden;
}

#title-bar {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 28px;
  line-height: 28px;
  background-color: #C7C5B6;
  z-index: 1;
  text-align: center;
  font-family: "PingFangSC-Medium";
  color: #333;
  font-size: 14px;
}

.view {
  position: absolute;
  top: 28px;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: hidden;
  transition: all 0.3s ease-in-out;
  &.with-notify {
    top: 56px;
  }
}

.windows .view {
  top: 0;
  &.with-notify {
    top: 28px;
  }
}
</style>
