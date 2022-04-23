<template>
    <div class="main-page">
        <div class="side-panel">
            <el-scrollbar height="calc(100% - 48px)">
                
                
            </el-scrollbar>
            <div class="tabs">
                    <el-row>
                        <el-col :span="8">
                            <el-icon><user-filled /></el-icon>
                        </el-col>
                        <el-col :span="8">
                            <el-icon><promotion /></el-icon>
                        </el-col>
                        <el-col :span="8">
                            <el-icon><tools /></el-icon>
                        </el-col>
                    </el-row>
                </div>
        </div>
        <div class="main-panel"></div>
    </div>
</template>

<script>

export default {
    name: "Main",
    components: {

    },
    inject: ["app"],
    data() {
        return {

        };
    },
    methods: {

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
        background-color: #D4D2C1;
        position: relative;

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
            }
        }
    }

    .main-panel {
        flex: auto;
        background-color: #DFDBCB;
    }
}
</style>
