<template>
    <div class="chat-item">
        <div class="avatar">
            <el-avatar v-if="item.avatar" :size="40" icon="user" />
            <span class="sys-avatar"></span>
        </div>
        <div class="info">
            <div class="name">
                <span>{{ item.name }}</span>
                <span class="update-time">21:33</span>
            </div>
            <div class="message">
                <div class="latest-message">
                    <span class="user">111</span>:<span class="msg">{{ item.latestLogin }}</span>
                </div>
                <span class="unread-count">123</span>
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
        box-sizing: border-box;
        width: calc(100% - 56px);
        padding: 6px 8px 0 0;
        flex: auto;
        border-bottom: 1px solid #ddd;

        .name {
            display: flex;
            justify-content: space-between;
            margin-bottom: 4px;

            .update-time {
                color: #898777;
                font-size: 12px;
            }
        }

        .message {
            display: flex;
            justify-content: space-between;
            align-items: center;

            .latest-message {
                width: calc(100% - 38px);
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
                .user {
                    color: #38697e;
                    font-size: 14px;
                }

                .msg {
                    color: #898777;
                    font-size: 14px;
                }
            }

            .unread-count {
                display: inline-block;
                padding: 0 4px;
                max-width: 30px;
                height: 18px;
                border-radius: 9px;
                background: #38697e;
                color: #fff;
                font-size: 12px;
                text-align: center;
                line-height: 18px;
                margin-left: 8px;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }
        }


    }
}
</style>