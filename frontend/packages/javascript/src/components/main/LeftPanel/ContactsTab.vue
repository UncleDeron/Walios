<template>
    <div class="contacts-tab">
        <header>
            <div class="search-wrap">
                <el-input
                    style="--el-input-bg-color: #c7c5b7; --el-input-border-color: transparent; --el-input-placeholder-color: #999"
                    placeholder="搜索" prefix-icon="Search" class="search-input" v-model="searchText"></el-input>
            </div>
            <el-icon>
                <chat-line-round />
            </el-icon>
        </header>
        <div class="user-list">
            <el-scrollbar height="100%">
                <div class="add-contact">
                    <el-icon>
                        <User></User>
                    </el-icon>
                    <span>添加联系人</span>
                </div>
                <user-item v-for="item in 40" :key="item" :item="{ name: item, latestLogin: '2022-04-24' }"
                    class="scrollbar-demo-item">{{ item }}</user-item>
            </el-scrollbar>
        </div>
    </div>
</template>

<script>
import UserItem from './UserItem.vue'
export default {
    name: 'contacts-tab',
    components: {
        UserItem,
    },
    data() {
        return {
            searchText: '',
        }
    },

}
</script>

<style lang="scss" scoped>
.contacts-tab {
    height: 100%;
    position: relative;

    header {
        display: flex;
        height: 48px;
        line-height: 48px;
        padding: 0 0 0 8px;

        .search-wrap {
            flex: auto;

            .el-input {
                width: 100%;
            }
        }

        .el-icon {
            width: 48px;
            height: 48px;
            line-height: 48px;
            text-align: right;
            font-size: 24px;
            color: #38697e;
            cursor: pointer;
        }
    }

    .user-list {
        height: calc(100% - 48px);

        .add-contact {
            height: 48px;
            line-height: 48px;
            font-size: 15px;
            color: #38697e;
            cursor: pointer;
            padding-left: 8px;
            border-bottom: 1px solid #ddd;

            .el-icon {
                width: 48px;
                float: left;
                margin-right: 8px;
                height: 48px;
                line-height: 48px;
                font-size: 24px;

                &::before {
                    content: "+";
                    position: absolute;
                    width: 9px;
                    height: 18px;
                    line-height: 12px;
                    left: 8px;
                    bottom: 5px;
                    font-size: 18px;
                    background: #D4D2C1;
                }
            }
        }
    }
}

</style>