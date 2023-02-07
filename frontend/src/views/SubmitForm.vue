<template>
    <suspense timeout="0">
        <template #default>
            <FormComponent />
        </template>
        <template #fallback>
            <div>Loading...</div>
        </template>
    </suspense>
</template>

<script setup>
import FormComponent from "../components/Form.vue";
</script>

<script>
export default {
    name: "SubmitForm",
    components: {
        FormComponent
    }
};
</script>
