<template>
    <div v-if="!isExists">
        <The404 />
    </div>
    <div v-for="(value, key) in data.albums[$route.params.category]" class="card w-96 bg-base-100 shadow-xl">
        <figure class="h-64">
            <NuxtLink :to="config.public.siteBaseUrl + '/' + $route.params.category + '/' + key">
                <nuxt-img width="18rem" quality="10" class="object-fit" :src="baseUrl+'/photos/'+$route.params.category+'/'+key+'/album.jpg'" :alt=key />
            </NuxtLink>
        </figure>
        <NuxtLink :to="config.public.siteBaseUrl + '/' + $route.params.category + '/' + key">
            <div class="card-body">
                <h2 class="card-title">{{ useCaptialiseString(value.name) }}</h2>
                <p class="break-words">🌍 {{ value.placeName }}</p>
                <p class="break-words">📅 {{ albumDatetime(value.startDate) }}</p>
            </div>
        </NuxtLink>
    </div>
    <div v-if="error">{{ error.message }}</div>
    <div v-else-if="!data">Loading...</div>
</template>

<script setup>
    const route = useRoute()
    const config = useRuntimeConfig()

    const baseUrl = config.public.APIBaseUrl
    const dataUrl = config.public.APIBaseUrl + "/albums"

    let isExists = ref(false)

    const { data, error } = await useFetch(dataUrl, { pick: ['albums']}).then(res => {
        var albums = res.data._rawValue.albums
        if (route.params.category in albums) {
            isExists = true
        }
        return res
    })

    let routeName = (route.params.category).charAt(0).toUpperCase() + route.fullPath.slice(2);

    const albums = data._rawValue.albums[route.params.category]
    var albums_key_array = Object.keys(albums)
    var randomAlbum = albums_key_array[Math.floor(Math.random()*albums_key_array.length)];
    var imageLocation = baseUrl+'/photos/'+route.params.category+'/'+randomAlbum+'/album.jpg'

    useSeoMeta({
        title: `${routeName} | ${config.siteName}`,
        ogTitle: `${routeName} | ${config.siteName}`,
        description: config.siteDescription,
        ogDescription: config.siteDescription,
        ogImage: imageLocation,
        ogSiteName: config.siteName,
        ogLocale: 'en_GB',
        ogType: 'website',
        ogUrl: config.siteBaseUrl,

        twitterCard: 'summary',
        twitterTitle: `${route.params.category} | ${config.siteName}`,
        twitterDescription: config.siteDescription,
        twitterImage: imageLocation,
    })
</script>