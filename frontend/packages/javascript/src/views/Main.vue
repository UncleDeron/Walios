<template>
    <div class="main-page">
        <div class="side-panel">
            <left-panel :activeTab="currentTab" @switchTab="switchTab"></left-panel>
        </div>
        <div class="main-panel"></div>
    </div>
</template>

<script>
import LeftPanel from "../components/main/LeftPanel/LeftPanel.vue";
export default {
    name: "Main",
    components: {
        LeftPanel
    },
    inject: ["app"],
    data() {
        return {
            currentTab: 'chats'
        };
    },
    methods: {
        switchTab(tab) {
            this.currentTab = tab;
        }
    }
};
</script>
<style scoped lang="scss">
.main-page {
    height: 100%;
    display: flex;

    &>div {
        height: 100%;
    }

    .side-panel {
        width: 300px;
    }

    .main-panel {
        flex: auto;
        background-color: #DFDBCB;
    }
}
</style>
