<template>
    <div class="chat-item">
        <div class="avatar">
            <el-avatar v-if="item.avatar" :size="40" icon="user" />
            <span class="sys-avatar"></span>
        </div>
        <div class="info">
            <div class="name">
                <span>{{ item.name }}</span>
                <span>{{ item.updateTime }}</span>
            </div>
            <div class="latest-message">
                <span>{{  }}</span>
                <span>{{ item.latestLogin }}</span>
            </div>
        </div>
    </div>
</template>

<script>

export default {
    name: 'chat-item',
    props: {
        item: {
            type: Object,
            required: true,
        }
    },
}
</script>

<style lang="scss" scoped>
    .chat-item {
        display: flex;
        padding: 8px 0 0 8px;
        cursor: pointer;
        .avatar {
            width: 56px;
            height: 56px;
            margin-right: 8px;
            text-align: center;

            .sys-avatar {
                display: inline-block;
                width: 48px;
                height: 48px;
                border-radius: 50%;
                background: #ccc;
            }
        }
        .info {
            padding: 6px 0;
            flex: auto;
            border-bottom: 1px solid #ddd;
            .name {
                display: flex;
                justify-content: space-between;
            }
            .signature {
                font-size: 14px;
                color: #999;
            }
        }
    }
</style>