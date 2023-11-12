<template>
    <div :style="{backgroundImage: `url(${config.public.headerImage})`}" class="headerText text-neutral-content">{{ config.public.siteName }}</div>
    
</template>

<script setup>
    const config = useRuntimeConfig()
</script>