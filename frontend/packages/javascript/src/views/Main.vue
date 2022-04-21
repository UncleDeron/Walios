<template>
    <div class="main-page">
        <div class="side-panel"></div>
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
    }

    .main-panel {
        flex: auto;
        background-color: #DFDBCB;
    }
}
</style>
