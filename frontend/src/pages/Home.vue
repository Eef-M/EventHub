<template>
  <DefaultLayout>
    <section class="bg-white py-20">
      <div class="max-w-7xl mx-auto px-4 text-center">
        <h1 class="text-5xl font-bold text-gray-900 mb-4">Discover Exciting Events</h1>
        <p class="text-lg text-gray-600 mb-8">
          Explore and manage event easily.
        </p>
        <div class="flex justify-center gap-4">
          <RouterLink to="/events"
            class="px-6 py-3 bg-blue-600 text-white rounded-xl shadow hover:bg-blue-700 transition">
            Browse Events
          </RouterLink>
        </div>
      </div>
    </section>

    <section class="bg-gray-50 py-16">
      <div class="max-w-6xl mx-auto px-4">
        <h2 class="text-3xl font-semibold text-center mb-10">News Events</h2>
        <div class="grid md:grid-cols-3 gap-6">
          <!-- Card -->
          <div class="rounded overflow-hidden shadow-lg flex flex-col">
            <a href="#"></a>
            <div class="relative"><a href="#">
                <img class="w-full"
                  src="https://images.pexels.com/photos/61180/pexels-photo-61180.jpeg?auto=compress&amp;cs=tinysrgb&amp;dpr=1&amp;w=500"
                  alt="Sunset in the mountains">
                <div
                  class="hover:bg-transparent transition duration-300 absolute bottom-0 top-0 right-0 left-0 bg-gray-900 opacity-25">
                </div>
              </a>
              <a href="#!">
                <div
                  class="text-xs absolute top-0 right-0 bg-indigo-600 px-4 py-2 text-white mt-3 mr-3 hover:bg-white hover:text-indigo-600 transition duration-500 ease-in-out">
                  Example
                </div>
              </a>
            </div>
            <div class="px-6 py-4 mb-auto">
              <a href="#"
                class="font-medium text-lg hover:text-indigo-600 transition duration-500 ease-in-out inline-block mb-2">Card
                Example</a>
              <p class="text-gray-500 text-sm">
                Lorem Ipsum is simply dummy text of the printing and typesetting industry.
              </p>
            </div>
            <div class="px-6 py-3 flex flex-row items-center justify-between bg-gray-100">
              <span href="#" class="py-1 text-xs font-regular text-gray-900 mr-1 flex flex-row items-center">
                <svg height="13px" width="13px" version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg"
                  xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 512 512"
                  style="enable-background:new 0 0 512 512;" xml:space="preserve">
                  <g>
                    <g>
                      <path
                        d="M256,0C114.837,0,0,114.837,0,256s114.837,256,256,256s256-114.837,256-256S397.163,0,256,0z M277.333,256 c0,11.797-9.536,21.333-21.333,21.333h-85.333c-11.797,0-21.333-9.536-21.333-21.333s9.536-21.333,21.333-21.333h64v-128 c0-11.797,9.536-21.333,21.333-21.333s21.333,9.536,21.333,21.333V256z">
                      </path>
                    </g>
                  </g>
                </svg>
                <span class="ml-1">6 mins ago</span>
              </span>

              <span href="#" class="py-1 text-xs font-regular text-gray-900 mr-1 flex flex-row items-center">
                <svg class="h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z">
                  </path>
                </svg>
                <span class="ml-1">39 Comments</span>
              </span>
            </div>
          </div>
          <!-- Card -->
        </div>
      </div>
    </section>
  </DefaultLayout>
</template>

<script setup lang="ts">
import DefaultLayout from '../layouts/DefaultLayout.vue';

</script>
