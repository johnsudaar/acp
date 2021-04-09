<template>
    <div class="container">
        <div class="messages">
        <p v-for="(msg) in $store.state.chat.chats[elem.chatId]" v-bind:key="msg.id" :style="styleFor(elem)">
            <span class="user"> {{msg.from}}: </span>
            <span class="message"> {{msg.message}} </span>
        </p>
        </div>
    </div>
</template>

<style scoped>
.container {
    overflow: hidden;
}
.messages {
    height: 100%;
    margin: 10px;
    border: solid 1px white;
}
.user {
    font-weight: bold;
    color: cyan;
}

p {
    padding: 0px;
    margin: 0px;
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