<template>
    <div style="height: 100%;">
        <div class="line_tab_container">
            <div
                v-for="(v, i) in data"
                :key="i"
                :class="{'active': v.isActive}"
                @click="choseTab(v, i)"
            >{{v.title}}</div>
        </div>
    </div>
</template>

<script>
export default {
    name: "MTabs",
    props: {
        data: {
            type: Array,
            default: []
        },
        chose: {
            type: Function,
            default: function() {}
        }
    },
    methods: {
        choseTab(v, i) {
            this.data.forEach(v => {
                v.isActive = false;
            });
            v.isActive = true;
            this.chose && this.chose(i, v);
        }
    }
};
</script>

<style lang="scss" scoped>
.line_tab_container {
    width: auto;
    height: 100%;
    display: flex;
    font-size: 14px;
    text-align: center;
    cursor: pointer;
    div {
        width: 120px;
        display: flex;
        align-items: center;
        justify-content: center;
        box-sizing: border-box;
        border: 1px solid #d7dce0;
        background-color: #ffffff;
        &:nth-of-type(1) {
            border-right-width: 0;
            border-top-left-radius: 1px;
            border-bottom-left-radius: 1px;
        }
        &:last-child {
            border-left-width: 0;
            border-top-right-radius: 1px;
            border-bottom-right-radius: 1px;
        }
    }
    .active {
        background-color: var(--bgColor);
        border-color: var(--bgColor);
        color: #ffffff;
    }
}
</style>
