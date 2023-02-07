<template>
    <div class="content" v-for="user in users" v-bind:key="user">
        <table class="mdl-data-table mdl-js-data-table">
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Email</th>
                    <th>Phone Number</th>
                    <th>Categories subscribed</th>
                    <th>Channels Subscribed</th>
                </tr>
            </thead>
            <tbody id="table-body">
                <td>{{ user["name"]}}</td>
                <td>{{ user["email"]}}</td>
                <td>{{ user["phone-number"] }}</td>
                <td>{{ user["subscribed"] }}</td>
                <td>{{ user["channels"] }}</td>
            </tbody>
        </table>
    </div>
</template>
<script setup>
import { ref } from 'vue'

const response = await (await fetch(process.env.VUE_APP_SERVER_URL + "/users")).json()

const users = ref([])
users.value = response["users"]
</script >

<script>
export default {
    name: 'UsersComponent',
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