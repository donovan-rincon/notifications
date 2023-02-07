<template>
    <h1>Messages history </h1>
    <div class="content" v-for="msg in msgs" v-bind:key="msg">
        <table class="mdl-data-table mdl-js-data-table">
            <thead>
                <tr>
                    <th>NotificationType</th>
                    <th>User</th>
                    <th>Timestamp</th>
                    <th>Message</th>
                </tr>
            </thead>
            <tbody id="table-body">
                <td>{{ msg["notification-type"]}}</td>
                <td>{{ msg["user"]}}</td>
                <td>{{ msg["timestamp"]}}</td>
                <td>{{ msg["message"] }}</td>
            </tbody>
        </table>
    </div>
</template>



<script setup>
import { ref } from 'vue'

const messages = await (await fetch(process.env.VUE_APP_SERVER_URL + "/logs/messages")).json()
console.log(messages["messages"])

const msgs = ref([])
msgs.value = messages["messages"]
</script >

<script>
export default {
    name: 'MessagesComponent',
}
</script>

<style scoped>
.content {
    margin-top: 20px;
    padding-left: 50px;
    padding-right: 50px;
}

table {
    width: 100%;
    border-collapse: collapse;
    padding: 30px;
}

th {
    background-color: #00AEEF;
    color: white;
    padding: 15px;
    text-align: left;
    font-weight: bold;
}

td {
    border: 1px solid #dddddd;
    padding: 15px;
    text-align: left;
}

tr:nth-child(even) {
    background-color: #f2f2f2;
}
</style>