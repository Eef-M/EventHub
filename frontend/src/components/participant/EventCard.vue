<template>
  <Card class="overflow-hidden bg-white shadow-lg hover:shadow-2xl transition-all duration-300 border-0 group">
    <div class="relative overflow-hidden">
      <img class="w-full h-48 object-cover group-hover:scale-110 transition-transform duration-300"
        :src="event.banner_url" :alt="event.title" />
      <div
        class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300" />
      <Badge class="absolute top-3 right-3 font-medium shadow-lg bg-yellow-600">
        {{ event.category || 'Event' }}
      </Badge>
    </div>

    <CardContent class="p-6">
      <h3 class="text-xl font-bold text-slate-900 mb-3 group-hover:text-purple-600 transition-colors">
        {{ truncateText(event.title, 50) }}
      </h3>
      <div class="flex items-center text-slate-600 mb-2">
        <MapPin class="h-4 w-4 mr-2 text-purple-500" />
        <span class="text-sm font-medium">{{ event.location }}</span>
      </div>
      <div class="flex items-center text-slate-600 mb-4">
        <Calendar class="h-4 w-4 mr-2 text-purple-500" />
        <span class="text-sm font-medium">{{ formatDate(event.date) }}</span>
      </div>
      <p class="text-slate-700 text-sm leading-relaxed mb-4">
        {{ event.description ? truncateText(event.description, 80) : 'No description available.' }}
      </p>
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-2">
          <div class="flex -space-x-2">
            <div class="w-8 h-8 rounded-full bg-purple-500 border-2 border-white flex items-center justify-center">
              <span class="text-xs font-semibold text-white">A</span>
            </div>
            <div class="w-8 h-8 rounded-full bg-blue-500 border-2 border-white flex items-center justify-center">
              <span class="text-xs font-semibold text-white">B</span>
            </div>
            <div class="w-8 h-8 rounded-full bg-green-500 border-2 border-white flex items-center justify-center">
              <span class="text-xs font-semibold text-white">+</span>
            </div>
          </div>
          <span class="text-xs text-slate-500">12+ attending</span>
        </div>
        <RouterLink :to="`/events/${event.id}/detail`">
          <Button variant="ghost" size="sm"
            class="text-purple-600 hover:text-purple-700 hover:bg-purple-50 font-medium cursor-pointer">
            View Details
            <ArrowRight class="ml-1 h-4 w-4" />
          </Button>
        </RouterLink>
      </div>
    </CardContent>
  </Card>
</template>

<script setup lang="ts">
import { formatDate } from '@/utils/format';
import type { EventInterface } from '@/types/event';
import { Card, CardContent } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { MapPin, Calendar, ArrowRight } from 'lucide-vue-next';

defineProps<{
  event: EventInterface
}>()

const truncateText = (text: string, maxLength: number) => {
  if (text.length <= maxLength) return text;
  return text.slice(0, maxLength) + '...';
}
</script>