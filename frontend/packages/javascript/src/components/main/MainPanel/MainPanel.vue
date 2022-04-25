<template>
    <div class="main-panel">
        <div class="main-content">
            <div class="blank">
                <div class="bg"></div>
            </div>
            <component :is="activeTabComponent"></component>
        </div>

    </div>
</template>

<script>


export default {
    name: "main-panel",
    components: {

    },
    props: {
        activeTab: {
            type: String,
            default: 'chats'
        }
    },
    data() {
        return {
            // activeTab: 'contacts',
        };
    },
    setup(props) {
        // return {
        //     activeTabComponent: props.activeTab + '-tab'
        // }
    },
    methods: {

    },
    computed: {
        activeTabComponent() {
            return this.activeTab + '-tab';
        }
    }
}
</script>

<style lang="scss" scoped>
.main-panel {
    width: 100%;
    height: 100%;
    background-color: #D4D2C1;
    position: relative;

    .main-content {
        height: 100%;
        overflow: hidden;
        background-color: #D4D2C1;

        .blank {

            width: 100%;
            height: 100%;
            background-color: #D4D2C1;

            .bg {
                position: absolute;
                top: 50%;
                left: 50%;
                width: 178px;
                height: 178px;
                background: url(@/assets/images/blank_bg.png) no-repeat;
                background-size: cover;
                opacity: .4;
                transform: translate(-50%, -50%);
            }
        }
    }
}
</style>