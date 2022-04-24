<template>
    <div class="user-item">
        <div class="avatar">
            <el-avatar v-if="item.avatar" :size="40" icon="user" />
            <span class="sys-avatar"></span>
        </div>
        <div class="info">
            <div class="name">
                <span>{{ item.name }}</span>
            </div>
            <div class="signature">
                <span>{{ item.latestLogin }}</span>
            </div>
        </div>
    </div>
</template>

<script>

export default {
    name: 'user-item',
    props: {
        item: {
            type: Object,
            required: true,
        }
    },
}
</script>

<style lang="scss" scoped>
    .user-item {
        display: flex;
        padding: 8px 0 0 8px;
        cursor: pointer;
        .avatar {
            width: 48px;
            height: 48px;
            margin-right: 8px;
            text-align: center;

            .sys-avatar {
                display: inline-block;
                width: 40px;
                height: 40px;
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