<template>
  <div class="notification-bar" :class="[type + '-background' ,{'is-active': active}]">
    <span>{{msg}}</span>
  </div>
</template>

<script>
export default {
  name: "Notification",
  props: {
    msg: String,
    active: Boolean,
    type: {
      type: String,
      default: "info"
    }
  },
  setup(props) {
    
  },
};
</script>

<style lang="scss">
.notification-bar {
  width: 100%;
  height: 24px;
  line-height: 24px;
  font-size: 14px;
  color: #fff;
  text-align: center;
  transform: translate(0, -100%);
  transition: all 0.3s ease-in-out;

  &.is-active {
    transform: translate(0, 0);
  }
}
</style>
