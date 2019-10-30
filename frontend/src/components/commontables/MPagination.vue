<template>
    <div class="common_pagination_content">
        <button
            v-if="showToFristEnd"
            @click="toFrist"
            :class="[currentPage === 1 || total === 0 ? 'no_disable' : '']"
        >«</button>
        <button @click="prev" :class="[currentPage === 1 || total === 0 ? 'no_disable' : '']">‹</button>
        <span class="page_number_content">{{currentPage}}/{{Math.ceil(total/pageSize) === 0 ? 1 : Math.ceil(total/pageSize)}}</span>
        <button @click="after" :class="[forMaxRange() || total === 0 ? 'no_disable' : '']">›</button>
        <button
            v-if="showToFristEnd"
            @click="toEnd"
            :class="[currentPage === totalPages || total === 0 ? 'no_disable' : '']"
        >»</button>
    </div>
</template>

<script>
export default {
    name: "MPagination",
    props: {
        page: {
            type: Number,
            default: 1
        },
        total: {
            type: Number,
            default: 0
        },
        pageSize: {
            type: Number,
            default: 30
        },
        pageChange: {
            type: Function,
            default: function() {}
        },
        range: {
            type: Array,
            default: function() {
                return [];
            }
        },
        showToFristEnd: {
            type: Boolean,
            default: true
        },
	    showInfo:{
        	type: Boolean,
            default: true
        },
        ascending: {
            //rang升序
            type: Boolean,
            default: true
        }
    },
    data() {
        return {
            currentPage: 1,
            currentRange: []
        };
    },
    computed: {
        totalPages() {
            return Math.ceil(this.total / this.pageSize);
        },
        from() {
            return (this.currentPage - 1) * this.pageSize + 1;
        },
        to() {
            let to = this.currentPage * this.pageSize;
            if (to < this.total) {
                return to;
            } else {
                return this.total;
            }
        }
    },
    methods: {
        toFrist() {
            this.currentPage = 1;
        },
        toEnd() {
            this.currentPage = this.totalPages;
        },
        prev() {
            if (this.currentPage > 1) {
                this.currentPage = this.currentPage - 1;
            }
        },
        after() {
            if (this.currentPage < this.totalPages && this.currentRange[1]) {
                if (!this.ascending) {
                    if (this.currentRange[1] > 1) {
                        this.currentPage = this.currentPage + 1;
                    }
                } else {
                    if (this.currentRange[1] < this.total) {
                        this.currentPage = this.currentPage + 1;
                    }
                }
            }
        },
        forRange() {
            if (this.ascending) {
                let after = this.currentPage * this.pageSize;
                this.currentRange = [
                    (this.currentPage - 1) * this.pageSize + 1,
                    after > this.total ? this.total : after
                ];
            } else {
                this.currentRange = this.range;
            }
        },
        forMaxRange() {
            if (this.ascending && this.currentRange[1] >= this.total) {
                return true;
            } else if (this.currentRange[1] === 1) {
                return true;
            }
            return false;
        }
    },
    watch: {
        page(newVal) {
            this.currentPage = newVal;
        },
        currentPage(newVal) {
            this.pageChange(newVal);
            this.forRange();
        },
        range(newVal) {
            this.currentRange = newVal;
        },
        total(newVal) {
            this.forRange();
        }
    },
    mounted() {
	    this.currentPage = this.page;
        this.forRange();
    }
};
</script>

<style lang="scss" scoped>
.info {
    color: var(--contentColor);
    vertical-align: middle;
    margin-right: 0.1rem;
}
.common_pagination_content{
    border: 0.01rem solid #dee2e6;
}
.page_number_content{
    padding: 0.05rem 0.05rem;
    font-size: 0.14rem;
    line-height: 1;
    background: #fff;
}
button {
    outline: none;
    border: none;
    padding: 5px 7.5px;
    color: var(--bgColor);
    line-height: 17px;
    background-color: #ffffff;
    cursor: pointer;
    &:hover {
        background-color: #e9ecef;
    }
    &:focus {
        outline: none;
    }
}
button:nth-child(1) {
    border-top-left-radius: 0.025rem;
    border-bottom-left-radius: 0.025rem;
    border-right: 0.01rem solid #dee2e6;
}
button:nth-of-type(2){
    border-right: 0.01rem solid #dee2e6;
}
button:nth-of-type(3){
    border-left: 0.01rem solid #dee2e6;
}
button:nth-last-child(1) {
    border-left: 0.01rem solid #dee2e6;
    border-top-right-radius: 0.025rem;
    border-bottom-right-radius: 0.025rem;
}
button.no_disable {
    cursor: auto;
    color: #c0c4cc;
    &:hover {
        background-color: #ffffff;
    }
}
</style>
