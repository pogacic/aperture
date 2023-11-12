export default defineNuxtConfig({
    css: ["~/assets/css/css.css", "~/assets/css/photoswipe-caption.css"],
    modules: [
        '@nuxtjs/tailwindcss',
        'nuxt-simple-robots',
        'nuxt-simple-sitemap',
        '@nuxt/image-edge',
        ['nuxt3-lazy-load', {
            loadingClass: 'isLoading',
            loadedClass: 'isLoaded',
            appendClass: 'lazyLoad',            
        }]
    ],
    sitemap: {
        // SET THIS FOR THE SITEMAP
        hostname: 'https://pogacic.com',
    },
    nitro: {
        prerender: {
            crawlLinks: true,
            routes: [
                '/',
            ]
        }
    },
    app: {
        pageTransition: { name: 'page', mode: 'out-in' },
        head: {
            charset: 'utf-8',
            viewport: 'width=device-width, initial-scale=1',
          }
      },
    head: {
        htmlAttrs: {
            lang: 'en'
        },
        link: [
            { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
        ]
    },
    runtimeConfig: {
        public: {
            // CHANGE THESE
            siteName: 'Aperture',
            siteDescription: 'Photography website',
            siteBaseUrl: 'https://pogacic.com',
            headerImage: 'https://api.pogacic.com/photos/header.jpg',
            APIBaseUrl: 'http://127.0.0.1:8080',

            // Don't touch unless you know what you're doing
            APIPhotos: 'photos'
        }
      },
});
