<template>
    <div v-for="(value, key) in data.categories">
        <div v-if='!key.startsWith("_")' class="card w-96 bg-base-100 shadow-xl">
            <figure class="h-64">
                <NuxtLink :to="key">
                    <nuxt-img quality="10" format="webp" class="object-fit" :src="baseUrl + '/' + value.randomAlbumCover" :alt=key />
                </NuxtLink>
            </figure>
            <NuxtLink :to="key">
                <div class="card-body">
                    <h2 class="card-title">{{ useCaptialiseString(key) }}</h2>
                    <p class="break-words">This is the {{ key }} category.</p>
                </div>
            </NuxtLink>
        </div>
    </div>
    <div v-if="error">{{ error.message }}</div>
    <div v-else-if="!data">Loading...</div>
</template>

<script setup>
    const config = useRuntimeConfig()
    const baseUrl = config.public.APIBaseUrl
    const dataUrl = config.public.APIBaseUrl + "/categories"

    const { data, error } = await useFetch(dataUrl, {
        pick: ['categories'],
    })

    const categories = data._rawValue.categories
    var category_key_array = Object.keys(categories)
    var randomCategory = category_key_array[Math.floor(Math.random()*category_key_array.length)];
    var imageLocation = baseUrl+'/'+categories[randomCategory].randomAlbumCover

    useSeoMeta({
        title: `Home | ${config.siteName}`,
        ogTitle: `Home | ${config.siteName}`,
        description: config.siteDescription,
        ogDescription: config.siteDescription,
        ogImage: imageLocation,
        ogSiteName: config.siteName,
        ogLocale: 'en_GB',
        ogType: 'website',
        ogUrl: config.siteBaseUrl,

        twitterCard: 'summary',
        twitterTitle: `Home | ${config.siteName}`,
        twitterDescription: config.siteDescription,
        twitterImage: imageLocation,
    })
</script>