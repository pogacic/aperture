<template>
    <div class="text-sm breadcrumbs pt-6">
        <ul class="flex flex-wrap">
            <li>
                <NuxtLink to='/'>
                    <svg class="w-4 h-4 mr-2 stroke-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2 8l9.732-4.866a.6.6 0 0 1 .536 0L22 8m-2 3v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"/></svg>
                    home
                </NuxtLink>
            </li>
            <li v-if=$route.params.category>
                <NuxtLink :to="'/' + $route.params.category">
                    <svg class="w-4 h-4 mr-2 stroke-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M5 19.5V5a2 2 0 0 1 2-2h11.4a.6.6 0 0 1 .6.6V21M9 7h6m-8.5 8H19M6.5 18H19M6.5 21H19"/><path stroke-linejoin="round" d="M6.5 18c-1 0-1.5-.672-1.5-1.5S5.5 15 6.5 15m0 6c-1 0-1.5-.672-1.5-1.5S5.5 18 6.5 18"/></g></svg>
                    {{ $route.params.category }}
                </NuxtLink>
            </li>
        </ul>
    </div>
</template>

