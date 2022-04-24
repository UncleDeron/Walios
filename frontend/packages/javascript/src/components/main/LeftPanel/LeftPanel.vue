<template>
    <div class="left-panel">
        <div class="tab-content">
            <component :is="activeTabComponent"></component>
        </div>
        <div class="tabs">
            <el-row>
                <el-col :span="8" @click="switchTab('contacts')">
                    <el-icon class="active">
                        <user-filled />
                    </el-icon>
                </el-col>
                <el-col :span="8" @click="switchTab('charts')">
                    <el-icon>
                        <promotion />
                    </el-icon>
                </el-col>
                <el-col :span="8" @click="switchTab('settings')">
                    <el-icon>
                        <tools />
                    </el-icon>
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script>
import ContactsTab from './ContactsTab.vue';

export default {
    name: "leftPanel",
    components: {
        ContactsTab
    },
    data() {
        return {
            activeTab: 'contacts',
        };
    },
    setup() {

    },
    methods: {
        switchTab(tab) {
            this.$emit('switchTab', tab);
        }
    },
    computed: {
        activeTabComponent() {
            return this.activeTab +　'-tab';
        }
    }
}
</script>

<style lang="scss" scoped>
.left-panel {
    width: 100%;
    height: 100%;
    background-color: #D4D2C1;
    position: relative;

    .tab-content {
        height: calc(100% - 48px);
        overflow: hidden;
        background-color: #D4D2C1;
    }

    .tabs {
        position: absolute;
        bottom: 0;
        left: 0;
        width: 100%;
        height: 48px;
        border-top: 1px solid #ddd;

        .el-col {
            height: 48px;
            line-height: 48px;
            text-align: center;
            font-size: 24px;
            color: #898777;
            cursor: pointer;

            .el-icon {
                height: 48px;
                line-height: 48px;
                &.active {
                    color: #38697e;
                }
            }
        }
    }
}
</style>