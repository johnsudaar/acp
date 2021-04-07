<template>
<div>
    <p v-for="(msg) in $store.state.chat.chats[elem.chatId]" v-bind:key="msg.id" :style="styleFor(elem)">
        <span class="user"> {{msg.from}}: </span>
        <span class="message"> {{msg.message}} </span>
    </p>
</div>
</template>

<style>
.user {
    font-weight: bold;
    color: cyan;
}
</style>

<script>
export default {
    props: {
        elem: Object
    },
    methods: {
        styleFor(elem) {
            let style = "";
            if(elem.chatFontSize) {
                style = `font-size: ${elem.chatFontSize};`;
            }
            return style;
        }
    }
}
</script>