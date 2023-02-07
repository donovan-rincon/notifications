<template>
    <h1>These are the mock users for the application and theis configuration </h1>
    <suspense timeout="0">
        <template #default>
            <UsersComponent />
        </template>
        <template #fallback>
            <div>Loading...</div>
        </template>
    </suspense>
</template>

<script>
import UsersComponent from '@/components/Users';

export default {
    name: "HomePage",
    components: { UsersComponent }
}
</script>
