<template>
    <suspense timeout="0">
        <template #default>
            <MessagesComponent />
        </template>
        <template #fallback>
            <div>Loading...</div>
        </template>
    </suspense>
</template>

<script setup>
import MessagesComponent from "../components/Messages.vue"
</script>

<script>
export default {
    name: "MessagesHistory",
    components: {
        MessagesComponent
    }
}
</script>
