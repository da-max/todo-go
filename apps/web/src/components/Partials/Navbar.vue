<script lang="ts" setup>
import NavbarUser from './NavbarUser.vue'
import { useUserStore } from '~/stores/user'
import { useToast } from 'flowbite-vue'
import { whenever } from '@vueuse/core'

const userStore = useUserStore()
const { add } = useToast()

userStore.getCurrent()

whenever(
    () => userStore.isAuthenticated,
    () => {
        add({
            type: 'success',
            time: 5_000,
            text: `Bienvenue ${userStore.user?.username}`,
        })
    },
)
</script>

<template>
    <nav
        class="z-10 p-5 bg-primary drop-shadow-lg flex justify-between align-middle"
    >
        <div>
            <h1 class="text-3xl text-white">📄 Todo GO</h1>
        </div>
        <NavbarUser />
    </nav>
</template>
