<template>
    <div v-if="!data">
        <The404 />
    </div>
    <div v-else>
        <div class="pt-4 break-words">
            <h1 class="text-5xl font-bold">{{data.name}}</h1>
            <p class="text-2xl mb-5 ">{{ data.subtitle }}</p>
            <p class="">🌍 {{ data.placeName }}</p>
            <p class="mb-5">📅 {{ parseDatetime(data.startDate, data.endDate) }}</p>
            <p class="">{{ data.description }}</p>
            <p class="mb-5">
                <NuxtLink class="link" :to="data.placeURL" target="_blank">
                    Find information about {{ data.name }} here
                </NuxtLink>
            </p>
        </div>
        <div id="the-gallery">
            <masonry-wall :items=genGalleryData() :ssr-columns="1" :column-width="300" :gap="16">
                <template #default="{ item }">
                    <a class="smooth-load"
                        :href="item.largeURL"
                        :data-pswp-width="item.width"
                        :data-pswp-height="item.height"
                        rel="noreferrer">

                        <span class="pswp-caption-content">
                            <p>Gear: {{ item.metadata.model }}, {{ item.metadata.lens_model }}</p>
                            <p>Settings: {{ item.metadata.exposure }}, f{{item.metadata.f_stop}}, iso{{ item.metadata.iso }}, {{ item.metadata.focal_length }}mm</p>
                        </span>

                        <img style="w-72" class="object-cover" :src="item.thumbnailURL" :alt="data.name + ' ' + data.subtitle + ' ' + data.description" />
                        
                    </a>
                </template>
            </masonry-wall>
        </div>
    </div>
</template>

<script setup>
import PhotoSwipeDynamicCaption from 'photoswipe-dynamic-caption-plugin';
import PhotoSwipeLightbox from 'photoswipe/lightbox';
import 'photoswipe/style.css';

const props = defineProps({
  album: String,
  category: String
})

const config = useRuntimeConfig()
const dataUrl = config.public.APIBaseUrl + "/" + props.category + '/' + props.album
const photoUrl = config.public.APIBaseUrl + "/" + config.public.APIPhotos + '/' + props.category + '/' + props.album

const { data } = await useFetch(dataUrl)

const genGalleryData = () => {
    const images = []
    for (const [key, value] of Object.entries(data.value['images'])) {
        images.push({
            largeURL: photoUrl + '/large/' + key,
            thumbnailURL: photoUrl + '/thumbnail/' + value.thumbnail,
            metadata: value.metadata,
            width: value.metadata.width,
            height: value.metadata.height,
        })
    }
    images.sort((a, b) => a.largeURL.localeCompare(b.largeURL, undefined, { numeric: true, sensitivity: 'base' }))
    return images
}

const lightbox = new PhotoSwipeLightbox({ 
    gallery: '#the-gallery',
    children: 'a',
    bgOpacity: 0.95,
    pswpModule: () => import('photoswipe'),
});

new PhotoSwipeDynamicCaption(lightbox, {
    type: 'auto',
});

onMounted( () => {
    lightbox.init();
});

useSeoMeta({
    title: `${data.value?.name} | ${config.siteName}`,
    ogTitle: `${data.value?.name} | ${config.siteName}`,
    description: `${data.value?.placeName} - ${data.value?.subtitle} - ${data.value?.description}`,
    ogDescription: `${data.value?.placeName} - ${data.value?.subtitle} - ${data.value?.description}`,
    ogImage: `${config.public.APIBaseUrl}/photos/${props.category}/${props.album}/album.jpg`,
    ogSiteName: config.siteName,
    ogLocale: 'en_GB',
    ogType: 'website',
    ogUrl: config.siteBaseUrl,

    twitterCard: 'summary',
    twitterTitle: `${data.value?.name} | ${config.siteName}`,
    twitterDescription: `${data.value?.placeName} - ${data.value?.subtitle} - ${data.value?.description}`,
    twitterImage: `${config.public.APIBaseUrl}/photos/${props.category}/${props.album}/album.jpg`,
})
</script>