<template>
    <h1>Form to submit new message and notify users</h1>
    <div class="content">
        <form>
            <div>
                <label>Select the categories of the message</label>
                <div v-for="option in options" :key="option.id" class="options">
                    <input type="checkbox" v-model="selectedOptions" :value="option.value" />
                    {{ option.text }}
                </div>
            </div>
            <div>
                <label>Write the message content</label>
                <textarea cols="50" rows="10" v-model="message"></textarea>
            </div>
            <button @click.prevent="submitForm">Submit</button>
        </form>
    </div>
</template>

<script setup>
import { ref, shallowRef } from 'vue'

const categories = await (await fetch(process.env.VUE_APP_SERVER_URL + "/categories")).json()

let options = []

for (const e of Object.entries(categories["categories"])) {
    options.push({ id: e[1], text: e[0], value: e[1] },)
}

const selectedOptions = shallowRef([])
const message = ref('')

const submitForm = async () => {


    const body = {
        "categories": selectedOptions.value,
        "content": message.value
    };

    const requestOptions = {
        method: "POST",
        body: JSON.stringify(body)
    };

    await fetch(process.env.VUE_APP_SERVER_URL + "/message", requestOptions);
    selectedOptions.value = []
    message.value = ''
}
</script >

<script>
export default {
    name: 'FormComponent',
}
</script>

<style scoped>
.options {
    display: inline-block;
    padding: 25px;
}

.content {
    margin-top: 20px;
    padding-left: 50px;
    padding-right: 50px;
}

form {
    width: 500px;
    margin: 50px auto;
    text-align: center;
}

input[type="submit"] {
    width: 100%;
    padding: 10px;
    background-color: #00AEEF;
    color: white;
    border: none;
    border-radius: 4px;
    font-size: 16px;
    cursor: pointer;
}

input[type="submit"]:hover {
    background-color: #0091C9;
}

label {
    font-size: 1.25rem;
    font-weight: bold;
    margin-bottom: 10px;
    display: block;
}
</style>